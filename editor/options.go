package editor

import (
	"asposecellscloud/requests"
)

// Option configures a high-level editor operation.
type Option func(*config)

type config struct {
	reqOpts []requests.RequestOption
}

func apply(cfg *config, opts []Option) {
	for _, o := range opts {
		if o != nil {
			o(cfg)
		}
	}
}

// WithCommonParameter passes an optional parameter to the underlying generated
// request. The name must match the operation parameter name in the spec.
func WithCommonParameter(name string, value interface{}) Option {
	return func(c *config) {
		c.reqOpts = append(c.reqOpts, requests.WithCommonParameter(name, value))
	}
}

// WithQueryParameter appends an implicit query parameter to the request URL.
func WithQueryParameter(name, value string) Option {
	return func(c *config) {
		c.reqOpts = append(c.reqOpts, requests.WithQueryParameter(name, value))
	}
}

// WithQueryParameters appends multiple implicit query parameters to the request
// URL.
func WithQueryParameters(params map[string]string) Option {
	return func(c *config) {
		c.reqOpts = append(c.reqOpts, requests.WithQueryParameters(params))
	}
}

// WithRaw passes any requests-level Option straight through to the generated
// request constructor.
func WithRaw(opts ...requests.RequestOption) Option {
	return func(c *config) {
		c.reqOpts = append(c.reqOpts, opts...)
	}
}

// --- Worksheet / spreadsheet parameters ---

// WithPosition sets the insert position of the new worksheet (AddWorksheet).
func WithPosition(p int) Option {
	return WithCommonParameter("position", intPtr(p))
}

// WithSheetType sets the sheet type for AddWorksheet (e.g. "Worksheet").
func WithSheetType(t string) Option {
	return WithCommonParameter("sheetType", t)
}

// WithTemplate initializes a new workbook from a cloud template file path
// (CreateSpreadsheet).
func WithTemplate(path string) Option {
	return WithCommonParameter("template", path)
}

// WithFormat sets the output format of a created workbook (CreateSpreadsheet).
func WithFormat(f string) Option {
	return WithCommonParameter("format", f)
}

// WithOutPath writes the edited workbook to this cloud path instead of
// returning it in the response body.
func WithOutPath(p string) Option {
	return WithCommonParameter("outPath", p)
}

// WithOutStorageName sets the cloud storage name for the output workbook.
func WithOutStorageName(s string) Option {
	return WithCommonParameter("outStorageName", s)
}

// WithPassword sets the workbook open password.
func WithPassword(p string) Option {
	return WithCommonParameter("password", p)
}

// WithRegion sets the region.
func WithRegion(r string) Option {
	return WithCommonParameter("region", r)
}

func intPtr(v int) *int { return &v }
