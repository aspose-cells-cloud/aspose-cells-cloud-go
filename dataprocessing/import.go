// Package dataprocessing provides data import helpers that merge a data file
// (CSV, JSON or XML) into a template spreadsheet, plus merge/split helpers.
//
// The v4.0 import endpoints (PUT /v4.0/cells/import/data[/csv|/json|/xml]) take
// the data file as the multipart "datafile" field and the target template
// spreadsheet as the multipart "Spreadsheet" field, and return the merged
// workbook as raw file bytes. The data format is selected by the entry point;
// the generic ImportData lets the server infer it from the datafile content.
package dataprocessing

import (
	"context"
	"fmt"

	"asposecellscloud"
	"asposecellscloud/datasource"
	"asposecellscloud/requests"
)

// importReqBuilder builds one of the v4.0 import requests for the given target
// position. The request then receives the datafile and template spreadsheet
// bytes via SetDatafileBytes/SetSpreadsheetBytes.
type importReqBuilder func(startcell, worksheet string, opts ...requests.RequestOption) asposecellscloud.RequestOption

// importLocal uploads datafile into the template spreadsheet at the given
// position and writes the merged workbook to sink.
func importLocal(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	datafile, template datasource.DataSource, sink datasource.DataSink,
	worksheet, startcell string, builder importReqBuilder, opts ...Option) error {

	if datafile == nil || template == nil || sink == nil {
		return fmt.Errorf("%w: datafile, template and sink are required", asposecellscloud.ErrInvalidParam)
	}
	if worksheet == "" || startcell == "" {
		return fmt.Errorf("%w: worksheet and start cell are required", asposecellscloud.ErrInvalidParam)
	}
	cfg := &config{}
	apply(cfg, opts)

	data, err := datafile.ByteData()
	if err != nil {
		return err
	}
	tpl, err := template.ByteData()
	if err != nil {
		return err
	}

	req := builder(startcell, worksheet, cfg.reqOpts...)
	if req == nil {
		return fmt.Errorf("%w: failed to build import request", asposecellscloud.ErrInvalidParam)
	}

	// Set the multipart file fields expected by the v4.0 import endpoints.
	switch r := req.(type) {
	case *requests.ImportCSVDataIntoSpreadsheetRequest:
		r.SetDatafileBytes(data, "datafile")
		r.SetSpreadsheetBytes(tpl, "Spreadsheet")
	case *requests.ImportJSONDataIntoSpreadsheetRequest:
		r.SetDatafileBytes(data, "datafile")
		r.SetSpreadsheetBytes(tpl, "Spreadsheet")
	case *requests.ImportXMLDataIntoSpreadsheetRequest:
		r.SetDatafileBytes(data, "datafile")
		r.SetSpreadsheetBytes(tpl, "Spreadsheet")
	case *requests.ImportDataIntoSpreadsheetRequest:
		r.SetDatafileBytes(data, "datafile")
		r.SetSpreadsheetBytes(tpl, "Spreadsheet")
	}

	resp, err := asposecellscloud.DoChecked(ctx, client, req)
	if err != nil {
		return err
	}
	return writeToSink(sink, resp.Body)
}

// ImportXML imports XML data into a template spreadsheet at the given
// worksheet/startcell position and writes the merged workbook to sink.
func ImportXML(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	datafile, template datasource.DataSource, sink datasource.DataSink,
	worksheet, startcell string, opts ...Option) error {
	return importLocal(ctx, client, datafile, template, sink, worksheet, startcell,
		func(startcell, worksheet string, opts ...requests.RequestOption) asposecellscloud.RequestOption {
			return requests.NewImportXMLDataIntoSpreadsheetRequest("", "", startcell, worksheet, opts...)
		}, opts...)
}

// ImportJSON imports JSON data into a template spreadsheet at the given
// worksheet/startcell position and writes the merged workbook to sink.
func ImportJSON(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	datafile, template datasource.DataSource, sink datasource.DataSink,
	worksheet, startcell string, opts ...Option) error {
	return importLocal(ctx, client, datafile, template, sink, worksheet, startcell,
		func(startcell, worksheet string, opts ...requests.RequestOption) asposecellscloud.RequestOption {
			return requests.NewImportJSONDataIntoSpreadsheetRequest("", "", startcell, worksheet, opts...)
		}, opts...)
}

// ImportCSV imports CSV data into a template spreadsheet at the given
// worksheet/startcell position and writes the merged workbook to sink.
func ImportCSV(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	datafile, template datasource.DataSource, sink datasource.DataSink,
	worksheet, startcell string, opts ...Option) error {
	return importLocal(ctx, client, datafile, template, sink, worksheet, startcell,
		func(startcell, worksheet string, opts ...requests.RequestOption) asposecellscloud.RequestOption {
			return requests.NewImportCSVDataIntoSpreadsheetRequest("", "", startcell, worksheet, opts...)
		}, opts...)
}

// ImportData imports data into a template spreadsheet at the given
// worksheet/startcell position and writes the merged workbook to sink. The
// data format is inferred from the datafile content.
func ImportData(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	datafile, template datasource.DataSource, sink datasource.DataSink,
	worksheet, startcell string, opts ...Option) error {
	return importLocal(ctx, client, datafile, template, sink, worksheet, startcell,
		func(startcell, worksheet string, opts ...requests.RequestOption) asposecellscloud.RequestOption {
			return requests.NewImportDataIntoSpreadsheetRequest("", "", startcell, worksheet, opts...)
		}, opts...)
}
