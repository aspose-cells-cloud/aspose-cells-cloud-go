package searcher

import (
	"asposecellscloud"
	"asposecellscloud/requests"
)

// Option configures a high-level search/replace operation.
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

// --- Search / replace scoping ---

// WithWorksheet scopes the search or replace to one worksheet.
func WithWorksheet(name string) Option {
	return WithCommonParameter("worksheet", name)
}

// WithCellArea scopes the search or replace to one cell range (A1 style, e.g.
// "E35:F40").
func WithCellArea(cellArea string) Option {
	return WithCommonParameter("cellArea", cellArea)
}

// WithIgnoringCase performs a case-insensitive search. Note the live v4.0
// endpoint still matches case-sensitively; the parameter is accepted for
// compatibility.
func WithIgnoringCase(v bool) Option {
	return WithCommonParameter("ignoringCase", asposecellscloud.BoolPtr(v))
}

// WithPassword sets the workbook open password.
func WithPassword(p string) Option {
	return WithCommonParameter("password", p)
}

// WithRegion sets the region.
func WithRegion(r string) Option {
	return WithCommonParameter("region", r)
}

// WithFolder sets the cloud folder that contains the workbook.
func WithFolder(f string) Option {
	return WithCommonParameter("folder", f)
}

// WithStorageName sets the cloud storage name.
func WithStorageName(s string) Option {
	return WithCommonParameter("storageName", s)
}
