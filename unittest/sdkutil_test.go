package unittest_test

import (
	"testing"

	asposecellscloud "asposecellscloud"
	"asposecellscloud/internal/sdkutil"
	"asposecellscloud/requests"
)

func TestWorkbookParams_Nil(t *testing.T) {
	opts := sdkutil.WorkbookParams(nil)
	if len(opts) != 0 {
		t.Errorf("expected empty options for nil ref, got %d", len(opts))
	}
}

func TestWorkbookParams_WithFolder(t *testing.T) {
	ref := &asposecellscloud.WorkbookRef{
		Name:   "test.xlsx",
		Folder: "/docs",
	}
	opts := sdkutil.WorkbookParams(ref)
	if len(opts) != 1 {
		t.Errorf("expected 1 option, got %d", len(opts))
	}
}

func TestWorkbookParams_WithFolderAndStorage(t *testing.T) {
	ref := &asposecellscloud.WorkbookRef{
		Name:        "test.xlsx",
		Folder:      "/docs",
		StorageName: "myStorage",
	}
	opts := sdkutil.WorkbookParams(ref)
	if len(opts) != 2 {
		t.Errorf("expected 2 options, got %d", len(opts))
	}
}

func TestWorkbookParams_EmptyFields(t *testing.T) {
	ref := &asposecellscloud.WorkbookRef{Name: "test.xlsx"}
	opts := sdkutil.WorkbookParams(ref)
	if len(opts) != 0 {
		t.Errorf("expected 0 options for empty folder/storage, got %d", len(opts))
	}
}

func TestApply_NilOptions(t *testing.T) {
	cfg := &sdkutil.Config{}
	sdkutil.Apply(cfg, []sdkutil.ConfigOption{nil, nil})
	if len(cfg.ReqOpts) != 0 {
		t.Error("nil options should not add to ReqOpts")
	}
}

func TestWithCommonParameter(t *testing.T) {
	cfg := &sdkutil.Config{}
	opt := sdkutil.WithCommonParameter("key", "value")
	opt(cfg)
	if len(cfg.ReqOpts) != 1 {
		t.Errorf("expected 1 req option, got %d", len(cfg.ReqOpts))
	}
}

func TestWithQueryParameter(t *testing.T) {
	cfg := &sdkutil.Config{}
	opt := sdkutil.WithQueryParameter("q", "v")
	opt(cfg)
	if len(cfg.ReqOpts) != 1 {
		t.Errorf("expected 1 req option, got %d", len(cfg.ReqOpts))
	}
}

func TestWithRaw(t *testing.T) {
	cfg := &sdkutil.Config{}
	raw := requests.WithCommonParameter("raw", "value")
	opt := sdkutil.WithRaw(raw)
	opt(cfg)
	if len(cfg.ReqOpts) != 1 {
		t.Errorf("expected 1 req option, got %d", len(cfg.ReqOpts))
	}
}
