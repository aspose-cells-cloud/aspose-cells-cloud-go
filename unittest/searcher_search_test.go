package unittest_test

import (
	"context"
	"errors"
	"testing"

	"asposecellscloud"
	"asposecellscloud/datasource"
	"asposecellscloud/internal/testutil"
	"asposecellscloud/searcher"
)

const searchBody = `{"TextItems":[{"Filename":"a.xlsx","Worksheet":"Sheet1","Position":"Cell:E35","Content":"margin"}]}`

func TestSearch(t *testing.T) {
	client, capture := testutil.NewServer(t, searchBody)

	items, err := searcher.Search(context.Background(), client, datasource.BytesSource([]byte("x")), "Sheet1", "margin")
	if err != nil {
		t.Fatalf("searcher.Search failed: %v", err)
	}
	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/search/content" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/search/content", c.Method, c.Path)
	}
	if got := c.Query.Get("searchText"); got != "margin" {
		t.Errorf("searchText = %q, want margin", got)
	}
	if got := c.Query.Get("worksheet"); got != "Sheet1" {
		t.Errorf("worksheet = %q, want Sheet1", got)
	}
	if got := string(c.Files["Spreadsheet"]); got != "x" {
		t.Errorf("Spreadsheet part = %q, want x", got)
	}
	if len(items) != 1 {
		t.Fatalf("got %d items, want 1", len(items))
	}
	if items[0].Text != "margin" || items[0].Worksheet != "Sheet1" || items[0].Position != "Cell:E35" {
		t.Errorf("item = %+v, want {margin Sheet1 Cell:E35}", items[0])
	}
}

func TestReplace(t *testing.T) {
	client, capture := testutil.NewServer(t, "replaced")
	sink := &datasource.BytesSink{}

	err := searcher.Replace(context.Background(), client, datasource.BytesSource([]byte("x")), sink, "old", "new")
	if err != nil {
		t.Fatalf("searcher.Replace failed: %v", err)
	}
	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/replace/content" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/replace/content", c.Method, c.Path)
	}
	if got := c.Query.Get("searchText"); got != "old" {
		t.Errorf("searchText = %q, want old", got)
	}
	if got := c.Query.Get("replaceText"); got != "new" {
		t.Errorf("replaceText = %q, want new", got)
	}
	if got := string(sink.Bytes()); got != "replaced" {
		t.Errorf("sink = %q, want replaced", got)
	}
}

func TestSearchWorksheet(t *testing.T) {
	client, capture := testutil.NewServer(t, searchBody)
	wf := &asposecellscloud.WorkbookRef{Name: "Book1.xlsx", Folder: "TestData/In", StorageName: "s3"}

	items, err := searcher.SearchWorksheet(context.Background(), client, wf, "Sheet1", "margin")
	if err != nil {
		t.Fatalf("searcher.SearchWorksheet failed: %v", err)
	}
	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/Book1.xlsx/worksheets/Sheet1/search/content" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/Book1.xlsx/worksheets/Sheet1/search/content", c.Method, c.Path)
	}
	if got := c.Query.Get("folder"); got != "TestData/In" {
		t.Errorf("folder = %q, want TestData/In", got)
	}
	if got := c.Query.Get("storageName"); got != "s3" {
		t.Errorf("storageName = %q, want s3", got)
	}
	if len(items) != 1 || items[0].Text != "margin" {
		t.Errorf("items = %+v, want one margin item", items)
	}
}

func TestSearchRange(t *testing.T) {
	client, capture := testutil.NewServer(t, searchBody)
	wf := &asposecellscloud.WorkbookRef{Name: "Book1.xlsx"}

	items, err := searcher.SearchRange(context.Background(), client, wf, "Sheet1", "E35:F40", "margin")
	if err != nil {
		t.Fatalf("searcher.SearchRange failed: %v", err)
	}
	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/Book1.xlsx/worksheets/Sheet1/ranges/E35:F40/search/content" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/Book1.xlsx/worksheets/Sheet1/ranges/E35:F40/search/content", c.Method, c.Path)
	}
	if len(items) != 1 {
		t.Errorf("items = %+v, want one item", items)
	}
}

func TestReplaceWorkbook(t *testing.T) {
	client, capture := testutil.NewServer(t, "ok")
	wf := &asposecellscloud.WorkbookRef{Name: "Book1.xlsx", StorageName: "s3"}

	if err := searcher.ReplaceWorkbook(context.Background(), client, wf, "old", "new"); err != nil {
		t.Fatalf("searcher.ReplaceWorkbook failed: %v", err)
	}
	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/Book1.xlsx/replace/content" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/Book1.xlsx/replace/content", c.Method, c.Path)
	}
	if got := c.Query.Get("searchText"); got != "old" {
		t.Errorf("searchText = %q, want old", got)
	}
	if got := c.Query.Get("replaceText"); got != "new" {
		t.Errorf("replaceText = %q, want new", got)
	}
	if got := c.Query.Get("storageName"); got != "s3" {
		t.Errorf("storageName = %q, want s3", got)
	}
}

func TestReplaceWorksheet(t *testing.T) {
	client, capture := testutil.NewServer(t, "ok")
	wf := &asposecellscloud.WorkbookRef{Name: "Book1.xlsx"}

	if err := searcher.ReplaceWorksheet(context.Background(), client, wf, "Sheet1", "old", "new"); err != nil {
		t.Fatalf("searcher.ReplaceWorksheet failed: %v", err)
	}
	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/Book1.xlsx/worksheets/Sheet1/replace/content" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/Book1.xlsx/worksheets/Sheet1/replace/content", c.Method, c.Path)
	}
}

func TestReplaceRange(t *testing.T) {
	client, capture := testutil.NewServer(t, "ok")
	wf := &asposecellscloud.WorkbookRef{Name: "Book1.xlsx"}

	if err := searcher.ReplaceRange(context.Background(), client, wf, "Sheet1", "E35:F40", "old", "new"); err != nil {
		t.Fatalf("searcher.ReplaceRange failed: %v", err)
	}
	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/Book1.xlsx/worksheets/Sheet1/ranges/E35:F40/replace/content" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/Book1.xlsx/worksheets/Sheet1/ranges/E35:F40/replace/content", c.Method, c.Path)
	}
}

func TestSearcherValidation(t *testing.T) {
	client, _ := testutil.NewServer(t, "")
	ctx := context.Background()
	src := datasource.BytesSource([]byte("x"))
	sink := &datasource.BytesSink{}
	wf := &asposecellscloud.WorkbookRef{Name: "B.xlsx"}

	if _, err := searcher.Search(ctx, client, nil, "S", "t"); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("searcher.Search nil source: got %v, want ErrInvalidParam", err)
	}
	if _, err := searcher.Search(ctx, client, src, "", "t"); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("searcher.Search empty worksheet: got %v, want ErrInvalidParam", err)
	}
	if _, err := searcher.Search(ctx, client, src, "S", ""); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("searcher.Search empty text: got %v, want ErrInvalidParam", err)
	}
	if err := searcher.Replace(ctx, client, nil, sink, "a", "b"); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("searcher.Replace nil source: got %v, want ErrInvalidParam", err)
	}
	if err := searcher.Replace(ctx, client, src, nil, "a", "b"); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("searcher.Replace nil sink: got %v, want ErrInvalidParam", err)
	}
	if _, err := searcher.SearchWorksheet(ctx, client, nil, "S", "t"); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("searcher.SearchWorksheet nil wf: got %v, want ErrInvalidParam", err)
	}
	if _, err := searcher.SearchRange(ctx, client, wf, "S", "", "t"); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("searcher.SearchRange empty cellArea: got %v, want ErrInvalidParam", err)
	}
	if err := searcher.ReplaceWorkbook(ctx, client, wf, "", "new"); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("searcher.ReplaceWorkbook empty old value: got %v, want ErrInvalidParam", err)
	}
	if err := searcher.ReplaceWorksheet(ctx, client, wf, "", "a", "b"); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("searcher.ReplaceWorksheet empty worksheet: got %v, want ErrInvalidParam", err)
	}
	if err := searcher.ReplaceRange(ctx, client, wf, "S", "", "a", "b"); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("searcher.ReplaceRange empty cellArea: got %v, want ErrInvalidParam", err)
	}
}
