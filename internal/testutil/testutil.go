// Package testutil provides a shared httptest harness for the high-level
// feature-package tests (converter, editor, searcher, reporting,
// dataprocessing, datacleansing).
package testutil

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"asposecellscloud"
)

// RequestCapture records the API request the SDK sent to the test server.
type RequestCapture struct {
	Method string
	Path   string
	Query  url.Values
	Auth   string
	CType  string
	Fields map[string]string // multipart text fields, keyed by field name
	Files  map[string][]byte // multipart file parts, keyed by part name
}

// NewServer returns a client wired to an httptest server that serves the OAuth
// token endpoint and records one API call, replying with body.
func NewServer(t *testing.T, body string) (*asposecellscloud.AsposeCellsCloudClient, func() RequestCapture) {
	t.Helper()
	c := RequestCapture{Fields: map[string]string{}, Files: map[string][]byte{}}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/cells/connect/token") {
			io.WriteString(w, `{"access_token":"tok123","token_type":"Bearer","expires_in":3600}`)
			return
		}
		c.Method = r.Method
		c.Path = r.URL.Path
		c.Query = r.URL.Query()
		c.Auth = r.Header.Get("Authorization")
		c.CType = r.Header.Get("Content-Type")
		if strings.HasPrefix(c.CType, "multipart/form-data") {
			if err := r.ParseMultipartForm(64 << 20); err != nil {
				t.Errorf("parse multipart form: %v", err)
				return
			}
			for k, v := range r.MultipartForm.Value {
				c.Fields[k] = v[0]
			}
			for k, fhs := range r.MultipartForm.File {
				f, err := fhs[0].Open()
				if err != nil {
					t.Errorf("open file part %q: %v", k, err)
					continue
				}
				b, _ := io.ReadAll(f)
				f.Close()
				c.Files[k] = b
			}
		}
		io.WriteString(w, body)
	}))
	t.Cleanup(ts.Close)
	return asposecellscloud.NewAsposeCellsCloudClient("id", "secret", ts.URL), func() RequestCapture { return c }
}
