package unittest_test

import (
	"encoding/json"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"asposecellscloud/requests"
)

// requireErrContains fails unless err is non-nil and its message contains substr.
func requireErrContains(t *testing.T, err error, substr string) {
	t.Helper()
	if err == nil {
		t.Fatalf("expected error containing %q, got nil", substr)
	}
	if !strings.Contains(err.Error(), substr) {
		t.Fatalf("expected error containing %q, got %q", substr, err.Error())
	}
}

// requireQuery asserts the value of a single query parameter.
func requireQuery(t *testing.T, v url.Values, key, want string) {
	t.Helper()
	if got := v.Get(key); got != want {
		t.Errorf("query %q: expected %q, got %q", key, want, got)
	}
}

// TestConvertSpreadsheetRequest covers the multipart file-upload shape: a
// required File parameter must be rejected by Validate() until bytes are
// attached, then emit the correct method, path, query and multipart fields.
func TestConvertSpreadsheetRequest(t *testing.T) {
	req := requests.NewConvertSpreadsheetRequest("pdf", "")
	requireErrContains(t, req.Validate(), "Spreadsheet")

	req.SetSpreadsheetBytes([]byte{1, 2, 3}, "Spreadsheet")
	if err := req.Validate(); err != nil {
		t.Fatalf("Validate with attached file should pass: %v", err)
	}

	if got := req.GetMethod(); got != "PUT" {
		t.Errorf("method: expected PUT, got %s", got)
	}
	if got := req.GetHeaderParameters()["Content-Type"]; got != "multipart/form-data" {
		t.Errorf("content-type: expected multipart/form-data, got %q", got)
	}
	if got := req.GetPath(); got != "/cells/convert/spreadsheet" {
		t.Errorf("path: expected /cells/convert/spreadsheet, got %q", got)
	}
	q := req.GetQueryParameters()
	requireQuery(t, q, "format", "pdf")
	if q.Get("outPath") != "" {
		t.Errorf("optional outPath should be omitted when empty, got %q", q.Get("outPath"))
	}

	form := req.GetMultipartForm()
	if len(form) != 1 {
		t.Fatalf("multipart form: expected 1 field, got %d", len(form))
	}
	if _, ok := form["Spreadsheet"].([]byte); !ok {
		t.Errorf("multipart Spreadsheet value should be []byte, got %T", form["Spreadsheet"])
	}
}

// TestExportSpreadsheetAsFormatRequest covers the cloud-file export shape: path
// segments are URL-encoded and optional query params are omitted when empty.
func TestExportSpreadsheetAsFormatRequest(t *testing.T) {
	req := requests.NewExportSpreadsheetAsFormatRequest("csv", "Book 1.xlsx")
	if err := req.Validate(); err != nil {
		t.Fatalf("Validate should pass: %v", err)
	}
	if got := req.GetMethod(); got != "GET" {
		t.Errorf("method: expected GET, got %s", got)
	}
	if got := req.GetHeaderParameters()["Content-Type"]; got != "application/json" {
		t.Errorf("content-type: expected application/json, got %q", got)
	}
	if got := req.GetPath(); got != "/cells/Book%201.xlsx" {
		t.Errorf("path: expected encoded path, got %q", got)
	}
	q := req.GetQueryParameters()
	requireQuery(t, q, "format", "csv")
	if q.Get("folder") != "" {
		t.Errorf("optional folder should be omitted when empty, got %q", q.Get("folder"))
	}

	withOpts := requests.NewExportSpreadsheetAsFormatRequest("csv", "book.xlsx", requests.WithCommonParameter("folder", "TestData/In"))
	q2 := withOpts.GetQueryParameters()
	requireQuery(t, q2, "folder", "TestData/In")
}

// TestPostCellCharactersRequest covers the JSON-body shape: multiple required
// path params are substituted, URL-encoded, and Validate() reports the missing
// parameter by its wire name.
func TestPostCellCharactersRequest(t *testing.T) {
	req := requests.NewPostCellCharactersRequest("A1", "Book 1.xlsx", "Sheet 1")
	if err := req.Validate(); err != nil {
		t.Fatalf("Validate should pass: %v", err)
	}
	want := "/cells/Book%201.xlsx/worksheets/Sheet%201/cells/A1/characters"
	if got := req.GetPath(); got != want {
		t.Errorf("path: expected %q, got %q", want, got)
	}
	if len(req.GetQueryParameters()) != 0 {
		t.Errorf("expected no query params, got %v", req.GetQueryParameters())
	}
	// An unset optional body is a typed-nil slice, which the interface wraps as
	// non-nil; it must serialize to JSON null (i.e. no payload).
	if body := req.GetJSONBody(); body != nil {
		if enc, _ := json.Marshal(body); string(enc) != "null" {
			t.Errorf("expected nil JSON body when options are unset, got %s", enc)
		}
	}
	if got := req.GetMethod(); got != "POST" {
		t.Errorf("method: expected POST, got %s", got)
	}

	missing := requests.NewPostCellCharactersRequest("", "n", "s")
	requireErrContains(t, missing.Validate(), "cellName")
}

// TestNoParamRequest covers the degenerate shape: a request with no required
// params validates cleanly and emits no body, no multipart and no query params.
func TestNoParamRequest(t *testing.T) {
	req := requests.NewGetCellsCloudServiceStatusRequest()
	if err := req.Validate(); err != nil {
		t.Fatalf("Validate should pass: %v", err)
	}
	if got := req.GetPath(); got != "/cells/status/check" {
		t.Errorf("path: expected /cells/status/check, got %q", got)
	}
	if got := req.GetMethod(); got != "GET" {
		t.Errorf("method: expected GET, got %s", got)
	}
	if len(req.GetQueryParameters()) != 0 {
		t.Errorf("expected no query params, got %v", req.GetQueryParameters())
	}
	if req.GetJSONBody() != nil {
		t.Errorf("expected nil JSON body, got %v", req.GetJSONBody())
	}
	if len(req.GetMultipartForm()) != 0 {
		t.Errorf("expected empty multipart form, got %v", req.GetMultipartForm())
	}
}

// TestImportXMLDataIntoSpreadsheetRequest covers the two-file import shape:
// both required File params must be populated, and GetMultipartForm supports
// both the byte-buffer convention and the "@filename" local-path convention.
func TestImportXMLDataIntoSpreadsheetRequest(t *testing.T) {
	req := requests.NewImportXMLDataIntoSpreadsheetRequest("", "", "A1", "Sheet1")
	requireErrContains(t, req.Validate(), "datafile")

	req.SetDatafileBytes([]byte("xml"), "datafile")
	req.SetSpreadsheetBytes([]byte("tpl"), "Spreadsheet")
	if err := req.Validate(); err != nil {
		t.Fatalf("Validate should pass once both files are attached: %v", err)
	}
	if got := req.GetPath(); got != "/cells/import/data/xml" {
		t.Errorf("path: expected /cells/import/data/xml, got %q", got)
	}
	q := req.GetQueryParameters()
	requireQuery(t, q, "worksheet", "Sheet1")
	requireQuery(t, q, "startcell", "A1")
	form := req.GetMultipartForm()
	if _, ok := form["datafile"].([]byte); !ok {
		t.Errorf("expected byte value under key datafile, got %v", form)
	}
	if _, ok := form["Spreadsheet"].([]byte); !ok {
		t.Errorf("expected byte value under key Spreadsheet, got %v", form)
	}

	// Local path variant: Validate passes on non-empty paths and the form keys
	// use the "@filename" convention the client maps to real file uploads.
	pathReq := requests.NewImportXMLDataIntoSpreadsheetRequest("data.xml", "tpl.xlsx", "A1", "Sheet1")
	if err := pathReq.Validate(); err != nil {
		t.Fatalf("Validate should pass with local file paths: %v", err)
	}
	pform := pathReq.GetMultipartForm()
	if _, ok := pform["@data.xml"]; !ok {
		t.Errorf("expected key @data.xml, got %v", pform)
	}
	if _, ok := pform["@tpl.xlsx"]; !ok {
		t.Errorf("expected key @tpl.xlsx, got %v", pform)
	}
}

// TestRequiredFileParamValidation is the contract for the required-File check:
// every request with a required upload param rejects construction until either
// bytes are attached (SetXxxBytes) or a non-empty local path is supplied.
func TestRequiredFileParamValidation(t *testing.T) {
	type validater interface {
		Validate() error
	}
	cases := []struct {
		name   string
		build  func() validater
		attach func(validater)
	}{
		{"add_worksheet",
			func() validater { return requests.NewAddWorksheetToSpreadsheetRequest("") },
			func(r validater) {
				r.(*requests.AddWorksheetToSpreadsheetRequest).SetSpreadsheetBytes([]byte{1}, "Spreadsheet")
			}},
		{"upload_file",
			func() validater { return requests.NewUploadFileRequest("out.xlsx", "") },
			func(r validater) { r.(*requests.UploadFileRequest).SetUploadFilesBytes([]byte{1}, "UploadFiles") }},
		{"remove_blank_rows",
			func() validater { return requests.NewRemoveSpreadsheetBlankRowsRequest("") },
			func(r validater) {
				r.(*requests.RemoveSpreadsheetBlankRowsRequest).SetSpreadsheetBytes([]byte{1}, "Spreadsheet")
			}},
		{"convert_spreadsheet",
			func() validater { return requests.NewConvertSpreadsheetRequest("pdf", "") },
			func(r validater) {
				r.(*requests.ConvertSpreadsheetRequest).SetSpreadsheetBytes([]byte{1}, "Spreadsheet")
			}},
		{"import_csv",
			func() validater { return requests.NewImportCSVDataIntoSpreadsheetRequest("", "", "A1", "Sheet1") },
			func(r validater) {
				r.(*requests.ImportCSVDataIntoSpreadsheetRequest).SetDatafileBytes([]byte{1}, "datafile")
				r.(*requests.ImportCSVDataIntoSpreadsheetRequest).SetSpreadsheetBytes([]byte{2}, "Spreadsheet")
			}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req := tc.build()
			if err := req.Validate(); err == nil {
				t.Fatal("expected Validate to reject a missing required file")
			}
			tc.attach(req)
			if err := req.Validate(); err != nil {
				t.Fatalf("Validate should pass after attaching the file: %v", err)
			}
		})
	}
}

// TestAllSpecOperationsHaveConstructors verifies 1:1 coverage between the
// specification operations and the generated request constructors. It catches a
// skipped or renamed operation that compilation alone would not report.
func TestAllSpecOperationsHaveConstructors(t *testing.T) {
	specBytes, err := os.ReadFile(filepath.Join("..", "aspose.cells.cloud.specification.json"))
	if err != nil {
		t.Fatalf("read spec: %v", err)
	}
	var spec struct {
		Operations []struct {
			Name    string `json:"Name"`
			Ignored bool   `json:"Ignored"`
		} `json:"Operations"`
	}
	if err := json.Unmarshal(specBytes, &spec); err != nil {
		t.Fatalf("parse spec: %v", err)
	}

	files, err := filepath.Glob(filepath.Join("..", "requests", "*.go"))
	if err != nil {
		t.Fatal(err)
	}
	constructors := make(map[string]bool)
	ctorRe := regexp.MustCompile(`func New(\w+Request)\(`)
	for _, f := range files {
		if strings.HasSuffix(f, "_test.go") {
			continue
		}
		src, err := os.ReadFile(f)
		if err != nil {
			t.Fatal(err)
		}
		for _, m := range ctorRe.FindAllSubmatch(src, -1) {
			constructors[string(m[1])] = true
		}
	}

	var missing []string
	for _, op := range spec.Operations {
		if op.Ignored {
			continue
		}
		if !constructors[op.Name+"Request"] {
			missing = append(missing, op.Name)
		}
	}
	if len(missing) > 0 {
		t.Errorf("spec operations missing generated constructors: %v", missing)
	}
	if got := len(constructors); got != 457 {
		t.Errorf("expected 457 generated request constructors, got %d", got)
	}
}
