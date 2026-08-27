package dataprocessing

import (
	"context"
	"fmt"

	"asposecellscloud"
	"asposecellscloud/datasource"
	"asposecellscloud/requests"
)

// MergeSpreadsheet merges the worksheets of a local spreadsheet into a single
// output workbook in the target format and writes the result bytes to sink.
// Set the target format with WithOutFormat (required); optionally merge
// everything into one sheet with WithMergeInOneSheet.
//
// Uses the v4.0 endpoint PUT /v4.0/cells/merge/spreadsheet?outFormat=...
func MergeSpreadsheet(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	src datasource.DataSource, sink datasource.DataSink, opts ...Option) error {

	if src == nil || sink == nil {
		return fmt.Errorf("%w: source and sink are required", asposecellscloud.ErrInvalidParam)
	}
	cfg := &config{}
	apply(cfg, opts)

	data, err := src.ByteData()
	if err != nil {
		return err
	}

	req := requests.NewMergeSpreadsheetsRequest("", cfg.reqOpts...)
	if req == nil {
		return fmt.Errorf("%w: failed to build merge request", asposecellscloud.ErrInvalidParam)
	}
	req.SetSpreadsheetBytes(data, "Spreadsheet")

	resp, err := asposecellscloud.DoChecked(ctx, client, req)
	if err != nil {
		return err
	}
	return writeToSink(sink, resp.Body)
}

// MergeRemoteSpreadsheet merges a cloud spreadsheet (mergedSpreadsheet) into
// another cloud workbook (wf) and writes the resulting workbook to sink. Set
// the target format with WithOutFormat. The merged file is not written back to
// cloud storage unless WithOutPath is set.
//
// Uses the v4.0 endpoint PUT /v4.0/cells/{name}/merge/spreadsheet.
func MergeRemoteSpreadsheet(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	wf *asposecellscloud.WorkbookRef, mergedSpreadsheet string, sink datasource.DataSink, opts ...Option) error {

	if wf == nil || wf.Name == "" || mergedSpreadsheet == "" || sink == nil {
		return fmt.Errorf("%w: workbook names and sink are required", asposecellscloud.ErrInvalidParam)
	}
	cfg := &config{}
	apply(cfg, opts)

	reqOpts := append(workbookParams(wf), cfg.reqOpts...)
	req := requests.NewMergeRemoteSpreadsheetRequest(mergedSpreadsheet, wf.Name, reqOpts...)
	if req == nil {
		return fmt.Errorf("%w: failed to build merge request", asposecellscloud.ErrInvalidParam)
	}

	resp, err := asposecellscloud.DoChecked(ctx, client, req)
	if err != nil {
		return err
	}
	return writeToSink(sink, resp.Body)
}
