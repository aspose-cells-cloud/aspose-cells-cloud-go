package asposecellscloud

import (
	"context"
	"fmt"
	"time"

	"asposecellscloud/requests"
)

// DoChecked executes a single request and verifies its HTTP status code,
// returning the first response on success. It is the shared building block
// used by the high-level feature packages.
func DoChecked(ctx context.Context, client *AsposeCellsCloudClient, req RequestOption) (*RichResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("%w: required parameters missing", ErrInvalidParam)
	}
	resps, err := client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	if len(resps) == 0 {
		return nil, fmt.Errorf("%w: empty response", ErrRequestFailed)
	}
	if err := CheckResponseStatus(resps[0]); err != nil {
		return nil, err
	}
	return resps[0], nil
}

// NewTempFileName returns a unique cloud-storage file name with the given
// extension (defaults to "xlsx"). The high-level search and import helpers use
// temporary files to round-trip local bytes through cloud storage, because the
// live API exposes several operations only against stored files.
func NewTempFileName(ext string) string {
	if ext == "" {
		ext = "xlsx"
	}
	return fmt.Sprintf("_cells_sdk_%d.%s", time.Now().UnixNano(), ext)
}

// UploadBytes uploads in-memory data to a cloud-storage file.
func UploadBytes(ctx context.Context, client *AsposeCellsCloudClient, cloudPath string, data []byte) error {
	req := requests.NewUploadFileRequest(cloudPath, "", requests.WithCommonParameter("storageName", ""))
	if req == nil {
		return fmt.Errorf("%w: cloud path is required", ErrInvalidParam)
	}
	req.SetUploadFilesBytes(data, "UploadFiles")
	_, err := DoChecked(ctx, client, req)
	return err
}

// DeleteCloudFile removes a file from cloud storage.
func DeleteCloudFile(ctx context.Context, client *AsposeCellsCloudClient, cloudPath string) error {
	req := requests.NewDeleteFileRequest(cloudPath, requests.WithCommonParameter("storageName", ""))
	if req == nil {
		return nil
	}
	_, err := DoChecked(ctx, client, req)
	return err
}
