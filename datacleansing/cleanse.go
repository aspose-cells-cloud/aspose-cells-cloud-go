// Package datacleansing provides helpers that clean a spreadsheet by removing
// blank rows, blank columns, blank worksheets and duplicate rows.
//
// Cleansing is exposed as local multipart operations: the spreadsheet is
// uploaded, the server removes the targeted content and returns the cleaned
// workbook as file bytes, which are written to the sink.
package datacleansing

import (
	"context"
	"fmt"

	"asposecellscloud"
	"asposecellscloud/datasource"
	"asposecellscloud/requests"
)

// RemoveBlankRows removes blank rows from a spreadsheet and writes the cleaned
// workbook to sink.
//
// Uses the v4.0 endpoint PUT /v4.0/cells/remove/blank-rows.
func RemoveBlankRows(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	src datasource.DataSource, sink datasource.DataSink, opts ...Option) error {
	return clean(ctx, client, src, sink,
		func(opts ...requests.RequestOption) asposecellscloud.RequestOption {
			return requests.NewRemoveSpreadsheetBlankRowsRequest("", opts...)
		}, opts...)
}

// RemoveBlankColumns removes blank columns from a spreadsheet and writes the
// cleaned workbook to sink.
//
// Uses the v4.0 endpoint PUT /v4.0/cells/remove/blank-columns.
func RemoveBlankColumns(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	src datasource.DataSource, sink datasource.DataSink, opts ...Option) error {
	return clean(ctx, client, src, sink,
		func(opts ...requests.RequestOption) asposecellscloud.RequestOption {
			return requests.NewRemoveSpreadsheetBlankColumnsRequest("", opts...)
		}, opts...)
}

// RemoveBlankWorksheets removes blank worksheets from a spreadsheet and writes
// the cleaned workbook to sink.
//
// Uses the v4.0 endpoint PUT /v4.0/cells/remove/blank-worksheets.
func RemoveBlankWorksheets(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	src datasource.DataSource, sink datasource.DataSink, opts ...Option) error {
	return clean(ctx, client, src, sink,
		func(opts ...requests.RequestOption) asposecellscloud.RequestOption {
			return requests.NewRemoveSpreadsheetBlankWorksheetsRequest("", opts...)
		}, opts...)
}

// RemoveDuplicates removes duplicate rows within a worksheet, cell range or
// table of a spreadsheet and writes the cleaned workbook to sink. Scope the
// removal with WithWorksheet / WithRange / WithTable.
//
// Uses the v4.0 endpoint PUT /v4.0/cells/remove/duplicates.
func RemoveDuplicates(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	src datasource.DataSource, sink datasource.DataSink, opts ...Option) error {
	return clean(ctx, client, src, sink,
		func(opts ...requests.RequestOption) asposecellscloud.RequestOption {
			return requests.NewRemoveDuplicatesRequest("", opts...)
		}, opts...)
}

// clean uploads a spreadsheet, runs the cleansing request and writes the
// cleaned workbook to sink.
func clean(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	src datasource.DataSource, sink datasource.DataSink,
	builder func(...requests.RequestOption) asposecellscloud.RequestOption, opts ...Option) error {

	if src == nil || sink == nil {
		return fmt.Errorf("%w: source and sink are required", asposecellscloud.ErrInvalidParam)
	}
	cfg := &config{}
	apply(cfg, opts)

	data, err := src.ByteData()
	if err != nil {
		return err
	}

	req := builder(cfg.reqOpts...)
	if req == nil {
		return fmt.Errorf("%w: failed to build cleansing request", asposecellscloud.ErrInvalidParam)
	}

	// Attach the spreadsheet to the multipart "Spreadsheet" field.
	switch r := req.(type) {
	case *requests.RemoveSpreadsheetBlankRowsRequest:
		r.SetSpreadsheetBytes(data, "Spreadsheet")
	case *requests.RemoveSpreadsheetBlankColumnsRequest:
		r.SetSpreadsheetBytes(data, "Spreadsheet")
	case *requests.RemoveSpreadsheetBlankWorksheetsRequest:
		r.SetSpreadsheetBytes(data, "Spreadsheet")
	case *requests.RemoveDuplicatesRequest:
		r.SetSpreadsheetBytes(data, "Spreadsheet")
	}

	resp, err := asposecellscloud.DoChecked(ctx, client, req)
	if err != nil {
		return err
	}
	return writeToSink(sink, resp.Body)
}
