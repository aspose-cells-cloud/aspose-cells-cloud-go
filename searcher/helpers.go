package searcher

import (
	"asposecellscloud"
	"asposecellscloud/datasource"
	"asposecellscloud/internal/sdkutil"
	"asposecellscloud/requests"
)

// writeToSink writes data to a DataSink.
func writeToSink(sink datasource.DataSink, data []byte) error {
	return sdkutil.WriteToSink(sink, data)
}

// workbookParams expands a WorkbookRef into requests-level options.
func workbookParams(wf *asposecellscloud.WorkbookRef) []requests.Option {
	return sdkutil.WorkbookParams(wf)
}
