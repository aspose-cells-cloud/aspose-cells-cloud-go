// Package sdkutil provides shared helpers for the high-level feature packages
// (converter, reporting, editor, searcher, datacleansing, dataprocessing).
// It is internal to avoid polluting the public API surface.
package sdkutil

import (
	"context"
	"fmt"

	"asposecellscloud"
	"asposecellscloud/datasource"
	"asposecellscloud/requests"
)

// DoChecked executes a single request and verifies its HTTP status code,
// returning the first response on success.
func DoChecked(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient, req asposecellscloud.RequestOption) (*asposecellscloud.RichResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("%w: required parameters missing", asposecellscloud.ErrInvalidParam)
	}
	resps, err := client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	if len(resps) == 0 {
		return nil, fmt.Errorf("%w: empty response", asposecellscloud.ErrRequestFailed)
	}
	if err := asposecellscloud.CheckResponseStatus(resps[0]); err != nil {
		return nil, err
	}
	return resps[0], nil
}

// WriteToSink writes data to a DataSink.
func WriteToSink(sink datasource.DataSink, data []byte) error {
	w, err := sink.Write()
	if err != nil {
		return err
	}
	defer w.Close()
	_, err = w.Write(data)
	return err
}

// WorkbookParams expands a WorkbookRef into request-level options.
func WorkbookParams(wf *asposecellscloud.WorkbookRef) []requests.Option {
	var out []requests.Option
	if wf == nil {
		return out
	}
	if wf.Folder != "" {
		out = append(out, requests.WithCommonParameter("folder", wf.Folder))
	}
	if wf.StorageName != "" {
		out = append(out, requests.WithCommonParameter("storageName", wf.StorageName))
	}
	return out
}

// Config holds the common option state shared by all high-level packages.
type Config struct {
	ReqOpts []requests.Option
}

// ConfigOption configures a Config.
type ConfigOption func(*Config)

// Apply applies a slice of ConfigOptions to cfg.
func Apply(cfg *Config, opts []ConfigOption) {
	for _, o := range opts {
		if o != nil {
			o(cfg)
		}
	}
}

// WithCommonParameter passes an optional parameter to the underlying generated request.
func WithCommonParameter(name string, value interface{}) ConfigOption {
	return func(c *Config) {
		c.ReqOpts = append(c.ReqOpts, requests.WithCommonParameter(name, value))
	}
}

// WithQueryParameter appends a custom query parameter to the request URL.
func WithQueryParameter(name, value string) ConfigOption {
	return func(c *Config) {
		c.ReqOpts = append(c.ReqOpts, requests.WithQueryParameter(name, value))
	}
}

// WithQueryParameters appends multiple custom query parameters.
func WithQueryParameters(params map[string]string) ConfigOption {
	return func(c *Config) {
		c.ReqOpts = append(c.ReqOpts, requests.WithQueryParameters(params))
	}
}

// WithRaw passes any requests-level Option straight through to the generated
// request constructor.
func WithRaw(opts ...requests.Option) ConfigOption {
	return func(c *Config) {
		c.ReqOpts = append(c.ReqOpts, opts...)
	}
}

// SpreadsheetSetter is implemented by request types that accept a spreadsheet
// file via SetSpreadsheetBytes. Using this interface avoids brittle type
// switches over concrete request types.
type SpreadsheetSetter interface {
	SetSpreadsheetBytes(data []byte, name string)
}

// DatafileSetter is implemented by request types that accept a data file
// via SetDatafileBytes (e.g. import endpoints).
type DatafileSetter interface {
	SetDatafileBytes(data []byte, name string)
}
