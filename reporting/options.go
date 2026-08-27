package reporting

import (
	"asposecellscloud/requests"
)

// Option configures a high-level report/analysis operation.
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

// WithRaw passes any requests-level Option straight through to the generated
// request constructor.
func WithRaw(opts ...requests.RequestOption) Option {
	return func(c *config) {
		c.reqOpts = append(c.reqOpts, opts...)
	}
}

// WithWorksheet restricts an analysis/calculation to a worksheet.
func WithWorksheet(w string) Option {
	return WithCommonParameter("worksheet", w)
}

// WithRange restricts an analysis/calculation to a cell range (A1 style).
func WithRange(r string) Option {
	return WithCommonParameter("range", r)
}

// WithOperation sets the aggregation/calculation operation (e.g. count/sum).
func WithOperation(op string) Option {
	return WithCommonParameter("operation", op)
}

// WithColorPosition sets the color position (e.g. fill/font) for
// AggregateByColor.
func WithColorPosition(p string) Option {
	return WithCommonParameter("colorPosition", p)
}

// WithPassword sets the workbook open password.
func WithPassword(p string) Option {
	return WithCommonParameter("password", p)
}

// WithRegion sets the region.
func WithRegion(r string) Option {
	return WithCommonParameter("region", r)
}
