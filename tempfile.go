package asposecellscloud

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"asposecellscloud/requests"
)

// tempFileSeq is a monotonic counter that guarantees uniqueness of temp file
// names within a process, even when two calls land on the same clock tick.
var tempFileSeq uint64

// NewTempFileName returns a unique cloud-storage file name with the given
// extension (defaults to "xlsx"). The high-level search and import helpers use
// temporary files to round-trip local bytes through cloud storage, because the
// live API exposes several operations only against stored files.
func NewTempFileName(ext string) string {
	if ext == "" {
		ext = "xlsx"
	}
	// The nanosecond timestamp plus a per-process counter guarantees a distinct
	// name on every call; the timestamp keeps names unique across processes.
	return fmt.Sprintf("_cells_sdk_%d_%d.%s", time.Now().UnixNano(), atomic.AddUint64(&tempFileSeq, 1), ext)
}

// UploadBytes uploads in-memory data to a cloud-storage file.
func UploadBytes(ctx context.Context, client *AsposeCellsCloudClient, cloudPath string, data []byte) error {
	if cloudPath == "" {
		return fmt.Errorf("%w: cloud path is required", ErrInvalidParam)
	}
	req := requests.NewUploadFileRequest(cloudPath, "", requests.WithCommonParameter("storageName", ""))
	req.SetUploadFilesBytes(data, "UploadFiles")
	_, err := DoChecked(ctx, client, req)
	return err
}

// DeleteCloudFile removes a file from cloud storage.
func DeleteCloudFile(ctx context.Context, client *AsposeCellsCloudClient, cloudPath string) error {
	if cloudPath == "" {
		return fmt.Errorf("%w: cloud path is required", ErrInvalidParam)
	}
	req := requests.NewDeleteFileRequest(cloudPath, requests.WithCommonParameter("storageName", ""))
	_, err := DoChecked(ctx, client, req)
	return err
}
