package asposecellscloud

import (
	"errors"
	"fmt"
)

var (
	ErrRequestFailed = errors.New("request failed")
	ErrInvalidParam  = errors.New("invalid parameter")
)

// SDKError 统一的 SDK 错误类型
type SDKError struct {
	Code    int
	Message string
	Err     error
}

func (e *SDKError) Error() string {
	return fmt.Sprintf("sdk error [code=%d]: %s: %v", e.Code, e.Message, e.Err)
}

func (e *SDKError) Unwrap() error {
	return e.Err
}
