package unittest_test

import (
	"encoding/json"
	"strings"
	"testing"

	asposecellscloud "asposecellscloud"
)

// --- Pointer helpers ---

func TestInt32Ptr(t *testing.T) {
	p := asposecellscloud.Int32Ptr(42)
	if p == nil || *p != 42 {
		t.Errorf("expected pointer to 42, got %v", p)
	}
}

func TestInt64Ptr(t *testing.T) {
	p := asposecellscloud.Int64Ptr(100)
	if p == nil || *p != 100 {
		t.Errorf("expected pointer to 100, got %v", p)
	}
}

func TestFloat64Ptr(t *testing.T) {
	p := asposecellscloud.Float64Ptr(3.14)
	if p == nil || *p != 3.14 {
		t.Errorf("expected pointer to 3.14, got %v", p)
	}
}

func TestBoolPtr(t *testing.T) {
	p := asposecellscloud.BoolPtr(true)
	if p == nil || *p != true {
		t.Errorf("expected pointer to true, got %v", p)
	}
	p2 := asposecellscloud.BoolPtr(false)
	if p2 == nil || *p2 != false {
		t.Errorf("expected pointer to false, got %v", p2)
	}
}

// --- WorkbookRef ---

func TestNewWorkbookRef(t *testing.T) {
	ref := asposecellscloud.NewWorkbookRef("test.xlsx")
	if ref == nil {
		t.Fatal("NewWorkbookRef returned nil")
	}
	if ref.Name != "test.xlsx" {
		t.Errorf("expected name 'test.xlsx', got %s", ref.Name)
	}
	if ref.Folder != "" {
		t.Errorf("expected empty folder, got %s", ref.Folder)
	}
	if ref.StorageName != "" {
		t.Errorf("expected empty storage name, got %s", ref.StorageName)
	}
}

func TestWorkbookRef_Fields(t *testing.T) {
	ref := &asposecellscloud.WorkbookRef{
		Name:        "book.xlsx",
		Folder:      "/docs",
		StorageName: "myStorage",
	}
	if ref.Name != "book.xlsx" {
		t.Errorf("expected name 'book.xlsx', got %s", ref.Name)
	}
	if ref.Folder != "/docs" {
		t.Errorf("expected folder '/docs', got %s", ref.Folder)
	}
	if ref.StorageName != "myStorage" {
		t.Errorf("expected storage 'myStorage', got %s", ref.StorageName)
	}
}

// --- RichResponse ---

func TestRichResponse_ToString(t *testing.T) {
	resp := &asposecellscloud.RichResponse{
		StatusCode: 200,
		Body:       []byte("hello world"),
	}
	if resp.ToString() != "hello world" {
		t.Errorf("expected 'hello world', got %s", resp.ToString())
	}
}

func TestRichResponse_ToBytes(t *testing.T) {
	data := []byte{1, 2, 3}
	resp := &asposecellscloud.RichResponse{StatusCode: 200, Body: data}
	b := resp.ToBytes()
	if len(b) != 3 || b[0] != 1 {
		t.Errorf("unexpected bytes: %v", b)
	}
}

func TestRichResponse_GetJSON(t *testing.T) {
	type payload struct {
		Name string `json:"name"`
	}
	body, _ := json.Marshal(payload{Name: "test"})
	resp := &asposecellscloud.RichResponse{StatusCode: 200, Body: body}

	var out payload
	if err := resp.GetJSON(&out); err != nil {
		t.Fatalf("GetJSON failed: %v", err)
	}
	if out.Name != "test" {
		t.Errorf("expected name 'test', got %s", out.Name)
	}
}

func TestRichResponse_GetJSON_NilTarget(t *testing.T) {
	resp := &asposecellscloud.RichResponse{StatusCode: 200, Body: []byte(`{}`)}
	if err := resp.GetJSON(nil); err != nil {
		t.Errorf("GetJSON with nil target should return nil, got: %v", err)
	}
}

func TestRichResponse_GetJSON_EmptyBody(t *testing.T) {
	resp := &asposecellscloud.RichResponse{StatusCode: 200, Body: nil}
	var out struct{}
	if err := resp.GetJSON(&out); err != nil {
		t.Errorf("GetJSON with empty body should return nil, got: %v", err)
	}
}

// --- TempFileName ---

func TestNewTempFileName_DefaultExt(t *testing.T) {
	name := asposecellscloud.NewTempFileName("")
	if name == "" {
		t.Error("NewTempFileName should return non-empty string")
	}
	if !strings.HasSuffix(name, "xlsx") {
		t.Errorf("expected .xlsx extension, got %s", name)
	}
}

func TestNewTempFileName_CustomExt(t *testing.T) {
	name := asposecellscloud.NewTempFileName("csv")
	if name == "" {
		t.Error("NewTempFileName should return non-empty string")
	}
	if !strings.HasSuffix(name, "csv") {
		t.Errorf("expected .csv extension, got %s", name)
	}
}

func TestNewTempFileName_Unique(t *testing.T) {
	n1 := asposecellscloud.NewTempFileName("xlsx")
	n2 := asposecellscloud.NewTempFileName("xlsx")
	if n1 == n2 {
		t.Error("two consecutive calls should produce different names")
	}
}
