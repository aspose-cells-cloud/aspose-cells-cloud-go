package dataprocessing

import (
	"context"
	"fmt"

	"asposecellscloud"
	"asposecellscloud/datasource"
	"asposecellscloud/requests"
)

// SplitSpreadsheet splits a local spreadsheet by worksheets into output files
// in the target format and writes the resulting file(s) to sink. Set the target
// format with WithOutFormat (required); optionally scope the split to a range
// of worksheets with WithFrom/WithTo.
//
// Uses the v4.0 endpoint PUT /v4.0/cells/split/spreadsheet?outFormat=...
func SplitSpreadsheet(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
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

	req := requests.NewSplitSpreadsheetRequest("", cfg.reqOpts...)
	if req == nil {
		return fmt.Errorf("%w: failed to build split request", asposecellscloud.ErrInvalidParam)
	}
	req.SetSpreadsheetBytes(data, "Spreadsheet")

	resp, err := asposecellscloud.DoChecked(ctx, client, req)
	if err != nil {
		return err
	}
	return writeToSink(sink, resp.Body)
}

// SplitRemoteSpreadsheet splits a cloud spreadsheet by worksheets in the target
// format. The split files are written to cloud storage at outPath (required),
// not returned in the response body; the raw response is returned so the caller
// can inspect the resulting file list. Set the target format with WithOutFormat.
//
// Uses the v4.0 endpoint PUT /v4.0/cells/{name}/split/spreadsheet.
func SplitRemoteSpreadsheet(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	wf *asposecellscloud.WorkbookRef, outPath string, opts ...Option) (*asposecellscloud.RichResponse, error) {

	if wf == nil || wf.Name == "" || outPath == "" {
		return nil, fmt.Errorf("%w: workbook name and outPath are required", asposecellscloud.ErrInvalidParam)
	}
	cfg := &config{}
	apply(cfg, opts)
	cfg.reqOpts = append(cfg.reqOpts, requests.WithCommonParameter("outPath", outPath))

	reqOpts := append(workbookParams(wf), cfg.reqOpts...)
	req := requests.NewSplitRemoteSpreadsheetRequest(wf.Name, reqOpts...)
	if req == nil {
		return nil, fmt.Errorf("%w: failed to build split request", asposecellscloud.ErrInvalidParam)
	}

	return asposecellscloud.DoChecked(ctx, client, req)
}
