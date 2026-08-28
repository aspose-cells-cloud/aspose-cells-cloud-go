package dataprocessing

import (
	"context"
	"fmt"

	"asposecellscloud"
	"asposecellscloud/datasource"
	"asposecellscloud/internal/sdkutil"
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
	cfg := &sdkutil.Config{}
	sdkutil.Apply(cfg, opts)

	data, err := src.ByteData()
	if err != nil {
		return err
	}

	req := requests.NewMergeSpreadsheetsRequest("", cfg.ReqOpts...)
	req.SetSpreadsheetBytes(data, "Spreadsheet")

	resp, err := asposecellscloud.DoChecked(ctx, client, req)
	if err != nil {
		return err
	}
	return sdkutil.WriteToSink(sink, resp.Body)
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
	cfg := &sdkutil.Config{}
	sdkutil.Apply(cfg, opts)

	reqOpts := append(sdkutil.WorkbookParams(wf), cfg.ReqOpts...)
	req := requests.NewMergeRemoteSpreadsheetRequest(mergedSpreadsheet, wf.Name, reqOpts...)

	resp, err := asposecellscloud.DoChecked(ctx, client, req)
	if err != nil {
		return err
	}
	return sdkutil.WriteToSink(sink, resp.Body)
}
