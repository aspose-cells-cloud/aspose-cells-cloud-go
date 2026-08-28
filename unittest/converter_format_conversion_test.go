package unittest_test

import (
	"context"
	"errors"
	"strings"
	"testing"

	"asposecellscloud"
	"asposecellscloud/converter"
	"asposecellscloud/datasource"
	"asposecellscloud/internal/testutil"
)

func TestConvert(t *testing.T) {
	client, capture := testutil.NewServer(t, "pdf-bytes")
	sink := &datasource.BytesSink{}

	err := converter.Convert(context.Background(), client, datasource.BytesSource([]byte("xlsx-bytes")), sink, converter.FormatPDF)
	if err != nil {
		t.Fatalf("converter.Convert failed: %v", err)
	}

	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/convert/spreadsheet" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/convert/spreadsheet", c.Method, c.Path)
	}
	if got := c.Query.Get("format"); got != "pdf" {
		t.Errorf("format = %q, want %q", got, "pdf")
	}
	if got := string(c.Files["Spreadsheet"]); got != "xlsx-bytes" {
		t.Errorf("Spreadsheet part = %q, want %q", got, "xlsx-bytes")
	}
	if c.Auth != "Bearer tok123" {
		t.Errorf("Authorization = %q, want Bearer tok123", c.Auth)
	}
	if !strings.HasPrefix(c.CType, "multipart/form-data") {
		t.Errorf("Content-Type = %q, want multipart", c.CType)
	}
	if got := string(sink.Bytes()); got != "pdf-bytes" {
		t.Errorf("sink = %q, want %q", got, "pdf-bytes")
	}
}

func TestConvertWrappers(t *testing.T) {
	tests := []struct {
		name string
		fn   func(context.Context, *asposecellscloud.AsposeCellsCloudClient, datasource.DataSource, datasource.DataSink, ...converter.Option) error
		want string
	}{
		{"converter.ConvertToPDF", converter.ConvertToPDF, converter.FormatPDF},
		{"converter.ConvertToCSV", converter.ConvertToCSV, converter.FormatCSV},
		{"converter.ConvertToJSON", converter.ConvertToJSON, converter.FormatJSON},
		{"converter.ConvertToHTML", converter.ConvertToHTML, converter.FormatHTML},
		{"converter.ConvertToXlsx", converter.ConvertToXlsx, converter.FormatXlsx},
		{"converter.ConvertToPNG", converter.ConvertToPNG, converter.FormatPNG},
		{"converter.ConvertToDocx", converter.ConvertToDocx, converter.FormatDocx},
		{"converter.ConvertToSQL", converter.ConvertToSQL, converter.FormatSQL},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, capture := testutil.NewServer(t, "out")
			sink := &datasource.BytesSink{}
			if err := tt.fn(context.Background(), client, datasource.BytesSource([]byte("x")), sink); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got := capture().Query.Get("format"); got != tt.want {
				t.Errorf("format = %q, want %q", got, tt.want)
			}
			if got := string(sink.Bytes()); got != "out" {
				t.Errorf("sink = %q, want %q", got, "out")
			}
		})
	}
}

func TestWorkbook(t *testing.T) {
	client, capture := testutil.NewServer(t, "csv-bytes")
	sink := &datasource.BytesSink{}

	wf := &asposecellscloud.WorkbookRef{Name: "Book1.xlsx", Folder: "TestData/In", StorageName: "s3"}
	err := converter.Workbook(context.Background(), client, wf, sink, converter.FormatCSV)
	if err != nil {
		t.Fatalf("converter.Workbook failed: %v", err)
	}

	c := capture()
	if c.Method != "GET" || c.Path != "/v4.0/cells/Book1.xlsx" {
		t.Errorf("request = %s %s, want GET /v4.0/cells/Book1.xlsx", c.Method, c.Path)
	}
	if got := c.Query.Get("format"); got != "csv" {
		t.Errorf("format = %q, want csv", got)
	}
	if got := c.Query.Get("folder"); got != "TestData/In" {
		t.Errorf("folder = %q, want TestData/In", got)
	}
	if got := c.Query.Get("storageName"); got != "s3" {
		t.Errorf("storageName = %q, want s3", got)
	}
	if got := string(sink.Bytes()); got != "csv-bytes" {
		t.Errorf("sink = %q, want %q", got, "csv-bytes")
	}
}

func TestWorksheet(t *testing.T) {
	client, capture := testutil.NewServer(t, "pdf-bytes")
	sink := &datasource.BytesSink{}

	wf := &asposecellscloud.WorkbookRef{Name: "Book1.xlsx"}
	err := converter.Worksheet(context.Background(), client, wf, "Sheet1", sink, converter.FormatPDF)
	if err != nil {
		t.Fatalf("converter.Worksheet failed: %v", err)
	}

	c := capture()
	if c.Method != "GET" || c.Path != "/v4.0/cells/Book1.xlsx/worksheets/Sheet1" {
		t.Errorf("request = %s %s, want GET /v4.0/cells/Book1.xlsx/worksheets/Sheet1", c.Method, c.Path)
	}
	if got := c.Query.Get("format"); got != "pdf" {
		t.Errorf("format = %q, want pdf", got)
	}
	if got := string(sink.Bytes()); got != "pdf-bytes" {
		t.Errorf("sink = %q, want %q", got, "pdf-bytes")
	}
}

func TestConvertValidation(t *testing.T) {
	client, _ := testutil.NewServer(t, "")
	ctx := context.Background()

	if err := converter.Convert(ctx, client, nil, &datasource.BytesSink{}, converter.FormatPDF); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("nil source: got %v, want ErrInvalidParam", err)
	}
	if err := converter.Convert(ctx, client, datasource.BytesSource([]byte("x")), nil, converter.FormatPDF); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("nil sink: got %v, want ErrInvalidParam", err)
	}
	if err := converter.Workbook(ctx, client, nil, &datasource.BytesSink{}, converter.FormatPDF); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("nil workbook: got %v, want ErrInvalidParam", err)
	}
	if err := converter.Workbook(ctx, client, &asposecellscloud.WorkbookRef{}, &datasource.BytesSink{}, converter.FormatPDF); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("empty workbook name: got %v, want ErrInvalidParam", err)
	}
	if err := converter.Worksheet(ctx, client, &asposecellscloud.WorkbookRef{Name: "B.xlsx"}, "", &datasource.BytesSink{}, converter.FormatPDF); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("empty worksheet: got %v, want ErrInvalidParam", err)
	}
}
