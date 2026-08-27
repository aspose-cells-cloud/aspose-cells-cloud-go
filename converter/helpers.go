package converter

import (
	"context"
	"fmt"

	"asposecellscloud"
	"asposecellscloud/datasource"
	"asposecellscloud/requests"
)

// runOne executes a single request and verifies its HTTP status code.
func runOne(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient, req asposecellscloud.RequestOption) (*asposecellscloud.RichResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("%w: required parameters missing", asposecellscloud.ErrInvalidParam)
	}
	resps, err := client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	if len(resps) == 0 {
		return nil, fmt.Errorf("%w: empty response", asposecellscloud.ErrRequestFailed)
	}
	if err := asposecellscloud.CheckResponseStatus(resps[0]); err != nil {
		return nil, err
	}
	return resps[0], nil
}

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
