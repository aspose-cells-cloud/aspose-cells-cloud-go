package datacleansing

import (
	"asposecellscloud/datasource"
	"asposecellscloud/internal/sdkutil"
)

// writeToSink writes data to a DataSink.
func writeToSink(sink datasource.DataSink, data []byte) error {
	return sdkutil.WriteToSink(sink, data)
}
