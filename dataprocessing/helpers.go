package dataprocessing

import (
	"asposecellscloud"
	"asposecellscloud/datasource"
	"asposecellscloud/requests"
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

// workbookParams expands a WorkbookRef into requests-level options.
func workbookParams(wf *asposecellscloud.WorkbookRef) []requests.RequestOption {
	var out []requests.RequestOption
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
