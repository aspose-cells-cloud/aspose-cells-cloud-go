package unittest_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	asposecellscloud "asposecellscloud"
	"asposecellscloud/requests"
)

// TestClientPipeline drives a request through the real client path
// (executeOnce: Validate -> URL build -> OAuth token -> headers -> multipart)
// against an httptest server. It verifies that:
//
//   - a missing required File param is rejected by Validate() before any HTTP
//     call is made,
//   - the versioned path, query params, Bearer token, User-Agent and multipart
//     content-type reach the wire as expected.
func TestClientPipeline(t *testing.T) {
	var gotPath, gotAuth, gotCT, gotUA string

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.HasSuffix(r.URL.Path, "/cells/connect/token"):
			io.WriteString(w, `{"access_token":"tok123","token_type":"Bearer","expires_in":3600}`)
		case r.URL.Path == "/v4.0/cells/convert/spreadsheet":
			gotPath = r.URL.RequestURI()
			gotAuth = r.Header.Get("Authorization")
			gotCT = r.Header.Get("Content-Type")
			gotUA = r.Header.Get("User-Agent")
			io.WriteString(w, "converted")
		default:
			t.Errorf("unexpected request to %s", r.URL.Path)
			http.NotFound(w, r)
		}
	}))
	defer ts.Close()

	client := asposecellscloud.NewAsposeCellsCloudClient("id", "secret", ts.URL)

	// Missing required file: rejected by Validate() before any network call.
	req := requests.NewConvertSpreadsheetRequest("pdf", "")
	if _, err := asposecellscloud.DoChecked(context.Background(), client, req); err == nil {
		t.Fatal("expected Validate to reject a request with no file attached")
	}

	req.SetSpreadsheetBytes([]byte("xlsx-bytes"), "Spreadsheet")
	resp, err := asposecellscloud.DoChecked(context.Background(), client, req)
	if err != nil {
		t.Fatalf("DoChecked failed: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
	if resp.ToString() != "converted" {
		t.Errorf("expected body %q, got %q", "converted", resp.ToString())
	}

	if !strings.HasPrefix(gotPath, "/v4.0/cells/convert/spreadsheet?") {
		t.Errorf("unexpected request URI: %s", gotPath)
	}
	if !strings.Contains(gotPath, "format=pdf") {
		t.Errorf("expected format=pdf in query, got %s", gotPath)
	}
	if gotAuth != "Bearer tok123" {
		t.Errorf("expected Bearer tok123, got %q", gotAuth)
	}
	if !strings.HasPrefix(gotCT, "multipart/form-data") {
		t.Errorf("expected multipart content-type, got %q", gotCT)
	}
	if !strings.Contains(gotUA, "Aspose Cells Cloud SDK") {
		t.Errorf("expected SDK User-Agent, got %q", gotUA)
	}
}

// TestClientPipelinePathEncoding drives a request whose path parameter contains
// special characters through the client and verifies the emitted URL keeps them
// properly encoded.
func TestClientPipelinePathEncoding(t *testing.T) {
	var gotURI string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.HasSuffix(r.URL.Path, "/cells/connect/token"):
			io.WriteString(w, `{"access_token":"tok123","expires_in":3600}`)
		case r.URL.Path == "/v4.0/cells/Book 1.xlsx/worksheets/Sheet 1":
			// r.URL.Path is decoded; the raw request URI preserves the encoding.
			gotURI = r.URL.RequestURI()
			io.WriteString(w, "ok")
		default:
			t.Errorf("unexpected request to %s", r.URL.Path)
			http.NotFound(w, r)
		}
	}))
	defer ts.Close()

	client := asposecellscloud.NewAsposeCellsCloudClient("id", "secret", ts.URL)
	req := requests.NewExportWorksheetAsFormatRequest("pdf", "Book 1.xlsx", "Sheet 1")
	resp, err := asposecellscloud.DoChecked(context.Background(), client, req)
	if err != nil {
		t.Fatalf("DoChecked failed: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
	if strings.Contains(gotURI, "Book 1") || strings.Contains(gotURI, "Sheet 1") {
		t.Errorf("path parameters must be URL-encoded, got %s", gotURI)
	}
	if !strings.Contains(gotURI, "Book%201") || !strings.Contains(gotURI, "Sheet%201") {
		t.Errorf("expected %%-encoded path segments, got %s", gotURI)
	}
}
