package asposecellscloud

import (
	"encoding/json"
	"fmt"

	"asposecellscloud/models"
)

// CheckResponseStatus verifies that an HTTP response is successful (2xx).
// client.Do only reports transport/retry errors and does not inspect the HTTP
// status code, so the high-level functional APIs must call this after every
// request. It returns a *SDKError wrapping the status code and, when
// available, the "Status" field from the API error body.
func CheckResponseStatus(resp *RichResponse) error {
	if resp == nil {
		return fmt.Errorf("%w: nil response", ErrRequestFailed)
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}
	msg := fmt.Sprintf("HTTP %d: %s", resp.StatusCode, string(resp.Body))
	var cc models.CellsCloudResponse
	_ = json.Unmarshal(resp.Body, &cc)
	if cc.Status != "" {
		msg = fmt.Sprintf("HTTP %d: %s", resp.StatusCode, cc.Status)
	}
	return &SDKError{Code: resp.StatusCode, Message: "request failed", Err: fmt.Errorf("%w: %s", ErrRequestFailed, msg)}
}
