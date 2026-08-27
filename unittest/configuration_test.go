package unittest_test

import (
	"testing"

	asposecellscloud "asposecellscloud"
)

func TestNewConfiguration_DefaultBasePath(t *testing.T) {
	cfg := asposecellscloud.NewConfiguration("id", "secret", "", "v4.0")
	if cfg.BasePath == "" {
		t.Error("BasePath should not be empty")
	}
	if cfg.Version != "v4.0" {
		t.Errorf("expected version v4.0, got %s", cfg.Version)
	}
	if cfg.ClientId != "id" {
		t.Errorf("expected client id 'id', got %s", cfg.ClientId)
	}
}

func TestNewConfiguration_WithBasePath(t *testing.T) {
	cfg := asposecellscloud.NewConfiguration("id", "secret", "https://custom.api.com", "v4.0")
	if cfg.BasePath != "https://custom.api.com" {
		t.Errorf("expected custom base path, got %s", cfg.BasePath)
	}
}

func TestNewConfiguration_StripTrailingSlash(t *testing.T) {
	cfg := asposecellscloud.NewConfiguration("id", "secret", "https://api.example.com/", "v3.0")
	if cfg.BasePath[len(cfg.BasePath)-1] == '/' {
		t.Error("BasePath should not end with /")
	}
}

func TestNewConfiguration_VersionFromBasePath(t *testing.T) {
	// When version param is empty, version is detected from basePath suffix.
	tests := []struct {
		basePath    string
		wantVersion string
	}{
		{"https://api.aspose.cloud/v4.0", "v4.0"},
		{"https://api.aspose.cloud/v3.0", "v3.0"},
		{"https://api.aspose.cloud/v1.1", "v1.1"},
	}
	for _, tt := range tests {
		cfg := asposecellscloud.NewConfiguration("id", "secret", tt.basePath, "")
		if cfg.Version != tt.wantVersion {
			t.Errorf("basePath=%s: expected version %s, got %s", tt.basePath, tt.wantVersion, cfg.Version)
		}
	}
}

func TestNewConfiguration_VersionExplicitOverrides(t *testing.T) {
	cfg := asposecellscloud.NewConfiguration("id", "secret", "https://api.aspose.cloud/v4.0", "v1.1")
	if cfg.Version != "v1.1" {
		t.Errorf("expected explicit version v1.1, got %s", cfg.Version)
	}
}

func TestConfiguration_AddDefaultHeader(t *testing.T) {
	cfg := asposecellscloud.NewConfiguration("id", "secret", "", "v4.0")
	cfg.AddDefaultHeader("X-Custom", "value")
	if cfg.DefaultHeader["X-Custom"] != "value" {
		t.Errorf("expected header X-Custom=value, got %s", cfg.DefaultHeader["X-Custom"])
	}
}
