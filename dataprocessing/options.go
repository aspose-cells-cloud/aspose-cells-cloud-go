package dataprocessing

import (
	"asposecellscloud"
	"asposecellscloud/internal/sdkutil"
	"asposecellscloud/requests"
)

// Option configures a high-level import operation.
type Option = sdkutil.ConfigOption

// WithCommonParameter passes an optional parameter to the underlying generated
// request. The name must match the operation parameter name in the spec.
func WithCommonParameter(name string, value interface{}) Option {
	return sdkutil.WithCommonParameter(name, value)
}

// WithQueryParameter appends an implicit query parameter to the request URL.
func WithQueryParameter(name, value string) Option {
	return sdkutil.WithQueryParameter(name, value)
}

// WithQueryParameters appends multiple implicit query parameters to the request
// URL.
func WithQueryParameters(params map[string]string) Option {
	return sdkutil.WithQueryParameters(params)
}

// WithRaw passes any requests-level Option straight through to the generated
// request constructor.
func WithRaw(opts ...requests.Option) Option {
	return sdkutil.WithRaw(opts...)
}

// WithInsert inserts data instead of overwriting the target cells.
func WithInsert(v bool) Option {
	return WithCommonParameter("insert", asposecellscloud.BoolPtr(v))
}

// WithConvertNumericData converts imported numeric strings to numeric cell
// values.
func WithConvertNumericData(v bool) Option {
	return WithCommonParameter("convertNumericData", asposecellscloud.BoolPtr(v))
}

// WithSplitter sets the column delimiter for CSV imports (default ",").
func WithSplitter(s string) Option {
	return WithCommonParameter("splitter", s)
}

// --- Merge / split parameter group ---

// WithOutFormat sets the target format of the merged or split output
// (MergeSpreadsheet, SplitSpreadsheet, MergeRemoteSpreadsheet,
// SplitRemoteSpreadsheet).
func WithOutFormat(f string) Option {
	return WithCommonParameter("outFormat", f)
}

// WithMergeInOneSheet merges all worksheets into a single sheet (merge).
func WithMergeInOneSheet(v bool) Option {
	return WithCommonParameter("mergeInOneSheet", asposecellscloud.BoolPtr(v))
}

// WithFrom scopes a split to start at the given zero-based worksheet index.
func WithFrom(i int) Option {
	return WithCommonParameter("from", intPtr(i))
}

// WithTo scopes a split to end at the given zero-based worksheet index.
func WithTo(i int) Option {
	return WithCommonParameter("to", intPtr(i))
}

// WithFontsLocation sets the location of fonts used for merge/split output.
func WithFontsLocation(l string) Option {
	return WithCommonParameter("fontsLocation", l)
}

// WithOutPath writes the merged workbook to this cloud path instead of
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
