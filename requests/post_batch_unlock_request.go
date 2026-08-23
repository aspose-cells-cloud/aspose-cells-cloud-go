package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostBatchUnlockRequest struct {
    batchLockRequest *models.BatchLockRequest
}

func NewPostBatchUnlockRequest(batchLockRequest *models.BatchLockRequest) *PostBatchUnlockRequest {
    req := &PostBatchUnlockRequest{
        batchLockRequest: batchLockRequest,
    }
    if req.batchLockRequest == nil {
        return nil
    }

    return req
}

func (request *PostBatchUnlockRequest) GetMethod() string {
    return "POST"
}

func (request *PostBatchUnlockRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostBatchUnlockRequest) GetPath() string {
    localVarPath := "/cells/batch/unlock"
    return localVarPath
}

func (request *PostBatchUnlockRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    return localVarQueryParams
}

func (request *PostBatchUnlockRequest) GetJSONBody() interface{} {
    return &request.batchLockRequest
}

func (request *PostBatchUnlockRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostBatchUnlockRequest) Description() {
    fmt.Println(strings.Trim("Batch unlocking files that meet specific matching conditions.", " "))
}