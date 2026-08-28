package unittest_test

import (
	"context"
	"errors"
	"testing"

	"asposecellscloud"
	"asposecellscloud/datasource"
	"asposecellscloud/editor"
	"asposecellscloud/internal/testutil"
)

func TestAddWorksheet(t *testing.T) {
	client, capture := testutil.NewServer(t, "updated")
	sink := &datasource.BytesSink{}

	err := editor.AddWorksheet(context.Background(), client, datasource.BytesSource([]byte("xlsx")), sink, "Sheet2", editor.WithPosition(3))
	if err != nil {
		t.Fatalf("editor.AddWorksheet failed: %v", err)
	}

	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/spreadsheet/add/worksheet" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/spreadsheet/add/worksheet", c.Method, c.Path)
	}
	if got := c.Query.Get("sheetName"); got != "Sheet2" {
		t.Errorf("sheetName = %q, want Sheet2", got)
	}
	if got := c.Query.Get("position"); got != "3" {
		t.Errorf("position = %q, want 3", got)
	}
	if got := string(c.Files["Spreadsheet"]); got != "xlsx" {
		t.Errorf("Spreadsheet part = %q, want xlsx", got)
	}
	if got := string(sink.Bytes()); got != "updated" {
		t.Errorf("sink = %q, want updated", got)
	}
}

func TestDeleteWorksheet(t *testing.T) {
	client, capture := testutil.NewServer(t, "updated")
	sink := &datasource.BytesSink{}

	if err := editor.DeleteWorksheet(context.Background(), client, datasource.BytesSource([]byte("x")), sink, "Sheet1"); err != nil {
		t.Fatalf("editor.DeleteWorksheet failed: %v", err)
	}
	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/spreadsheet/delete/worksheet" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/spreadsheet/delete/worksheet", c.Method, c.Path)
	}
	if got := c.Query.Get("sheetName"); got != "Sheet1" {
		t.Errorf("sheetName = %q, want Sheet1", got)
	}
}

func TestRenameWorksheet(t *testing.T) {
	client, capture := testutil.NewServer(t, "updated")
	sink := &datasource.BytesSink{}

	if err := editor.RenameWorksheet(context.Background(), client, datasource.BytesSource([]byte("x")), sink, "Old", "New"); err != nil {
		t.Fatalf("editor.RenameWorksheet failed: %v", err)
	}
	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/spreadsheet/rename/worksheet" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/spreadsheet/rename/worksheet", c.Method, c.Path)
	}
	if got := c.Query.Get("sourceName"); got != "Old" {
		t.Errorf("sourceName = %q, want Old", got)
	}
	if got := c.Query.Get("targetName"); got != "New" {
		t.Errorf("targetName = %q, want New", got)
	}
}

func TestMoveWorksheet(t *testing.T) {
	client, capture := testutil.NewServer(t, "updated")
	sink := &datasource.BytesSink{}

	if err := editor.MoveWorksheet(context.Background(), client, datasource.BytesSource([]byte("x")), sink, "Sheet1", 2); err != nil {
		t.Fatalf("editor.MoveWorksheet failed: %v", err)
	}
	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/spreadsheet/move/worksheet" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/spreadsheet/move/worksheet", c.Method, c.Path)
	}
	if got := c.Query.Get("worksheet"); got != "Sheet1" {
		t.Errorf("worksheet = %q, want Sheet1", got)
	}
	if got := c.Query.Get("position"); got != "2" {
		t.Errorf("position = %q, want 2", got)
	}
}

func TestListWorksheets(t *testing.T) {
	client, capture := testutil.NewServer(t, `[{"WorksheetName":"Sheet1","SheetType":"Worksheet"},{"WorksheetName":"Data","SheetType":"Worksheet"}]`)

	ws, err := editor.ListWorksheets(context.Background(), client, datasource.BytesSource([]byte("x")))
	if err != nil {
		t.Fatalf("editor.ListWorksheets failed: %v", err)
	}
	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/spreadsheet/worksheets" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/spreadsheet/worksheets", c.Method, c.Path)
	}
	if len(ws) != 2 {
		t.Fatalf("got %d worksheets, want 2", len(ws))
	}
	if ws[0].Name != "Sheet1" || ws[0].Type != "Worksheet" {
		t.Errorf("ws[0] = %+v, want {Sheet1 Worksheet}", ws[0])
	}
	if ws[1].Name != "Data" {
		t.Errorf("ws[1].Name = %q, want Data", ws[1].Name)
	}
}

func TestCreateSpreadsheet(t *testing.T) {
	client, capture := testutil.NewServer(t, "new-xlsx")
	sink := &datasource.BytesSink{}

	if err := editor.CreateSpreadsheet(context.Background(), client, sink); err != nil {
		t.Fatalf("editor.CreateSpreadsheet failed: %v", err)
	}
	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/spreadsheet/create" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/spreadsheet/create", c.Method, c.Path)
	}
	if len(c.Files) != 0 {
		t.Errorf("expected no file parts, got %v", c.Files)
	}
	if got := string(sink.Bytes()); got != "new-xlsx" {
		t.Errorf("sink = %q, want new-xlsx", got)
	}
}

func TestEditorValidation(t *testing.T) {
	client, _ := testutil.NewServer(t, "")
	ctx := context.Background()
	src := datasource.BytesSource([]byte("x"))
	sink := &datasource.BytesSink{}

	if err := editor.AddWorksheet(ctx, client, nil, sink, "S"); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("editor.AddWorksheet nil source: got %v, want ErrInvalidParam", err)
	}
	if err := editor.AddWorksheet(ctx, client, src, nil, "S"); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("editor.AddWorksheet nil sink: got %v, want ErrInvalidParam", err)
	}
	if err := editor.AddWorksheet(ctx, client, src, sink, ""); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("editor.AddWorksheet empty sheet name: got %v, want ErrInvalidParam", err)
	}
	if err := editor.DeleteWorksheet(ctx, client, src, sink, ""); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("editor.DeleteWorksheet empty sheet name: got %v, want ErrInvalidParam", err)
	}
	if err := editor.RenameWorksheet(ctx, client, src, sink, "", "New"); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("editor.RenameWorksheet empty old name: got %v, want ErrInvalidParam", err)
	}
	if err := editor.MoveWorksheet(ctx, client, src, sink, "", 0); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("editor.MoveWorksheet empty worksheet: got %v, want ErrInvalidParam", err)
	}
	if _, err := editor.ListWorksheets(ctx, client, nil); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("editor.ListWorksheets nil source: got %v, want ErrInvalidParam", err)
	}
	if err := editor.CreateSpreadsheet(ctx, client, nil); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("editor.CreateSpreadsheet nil sink: got %v, want ErrInvalidParam", err)
	}
}
