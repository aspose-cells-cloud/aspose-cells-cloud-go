package converter

import (
	"context"

	"asposecellscloud"
	"asposecellscloud/datasource"
	"asposecellscloud/internal/sdkutil"
	"asposecellscloud/requests"
)

// runOne executes a single request and verifies its HTTP status code.
func runOne(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient, req asposecellscloud.RequestOption) (*asposecellscloud.RichResponse, error) {
	return sdkutil.DoChecked(ctx, client, req)
}

// writeToSink writes data to a DataSink.
func writeToSink(sink datasource.DataSink, data []byte) error {
	return sdkutil.WriteToSink(sink, data)
}

// workbookParams expands a WorkbookRef into requests-level options.
func workbookParams(wf *asposecellscloud.WorkbookRef) []requests.Option {
	return sdkutil.WorkbookParams(wf)
}
