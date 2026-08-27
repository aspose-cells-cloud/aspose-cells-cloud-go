package datacleansing

import (
	"asposecellscloud/requests"
)

// Option configures a high-level data cleansing operation.
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

// --- Cleansing scoping ---

// WithWorksheet scopes duplicate removal to one worksheet (RemoveDuplicates).
func WithWorksheet(name string) Option {
	return WithCommonParameter("worksheet", name)
}

// WithRange scopes duplicate removal to one cell range, A1 style e.g. "E35:F40"
// (RemoveDuplicates).
func WithRange(cellArea string) Option {
	return WithCommonParameter("range", cellArea)
}

// WithTable scopes duplicate removal to a named table (RemoveDuplicates).
func WithTable(name string) Option {
	return WithCommonParameter("table", name)
}

// --- Output / credentials ---

// WithOutPath writes the cleaned workbook to this cloud path instead of
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
