package datacleansing

import (
	"asposecellscloud/datasource"
)

// writeToSink writes data to a DataSink.
func writeToSink(sink datasource.DataSink, data []byte) error {
	w, err := sink.Write()
	if err != nil {
		return err
	}
	defer w.Close()
	_, err = w.Write(data)
	return err
}
