package asposecellscloud

import (
	"context"
	"fmt"
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
