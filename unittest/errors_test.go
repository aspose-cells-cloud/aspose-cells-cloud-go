package unittest_test

import (
	"errors"
	"strings"
	"testing"

	asposecellscloud "asposecellscloud"
)

func TestSDKError_Error(t *testing.T) {
	err := &asposecellscloud.SDKError{Code: 404, Message: "not found", Err: errors.New("resource missing")}
	s := err.Error()
	if !strings.Contains(s, "404") {
		t.Errorf("expected error string to contain '404', got: %s", s)
	}
}

func TestSDKError_Unwrap(t *testing.T) {
	inner := errors.New("inner error")
	err := &asposecellscloud.SDKError{Code: 500, Message: "server error", Err: inner}
	if !errors.Is(err, inner) {
		t.Error("Unwrap should return the inner error")
	}
}

func TestSDKError_WrapsWithFmtErrorf(t *testing.T) {
	sdkErr := &asposecellscloud.SDKError{Code: -1, Message: "failed", Err: asposecellscloud.ErrRequestFailed}
	if !errors.Is(sdkErr, asposecellscloud.ErrRequestFailed) {
		t.Error("should be able to match wrapped sentinel error")
	}
}
