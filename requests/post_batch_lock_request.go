package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostBatchLockRequest struct {
    batchLockRequest *models.BatchLockRequest
}

func NewPostBatchLockRequest(batchLockRequest *models.BatchLockRequest) *PostBatchLockRequest {
    req := &PostBatchLockRequest{
        batchLockRequest: batchLockRequest,
    }
    if req.batchLockRequest == nil {
        return nil
    }

    return req
}

func (request *PostBatchLockRequest) GetMethod() string {
    return "POST"
}

func (request *PostBatchLockRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostBatchLockRequest) GetPath() string {
    localVarPath := "/cells/batch/lock"
    return localVarPath
}

func (request *PostBatchLockRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    return localVarQueryParams
}

func (request *PostBatchLockRequest) GetJSONBody() interface{} {
    return &request.batchLockRequest
}

func (request *PostBatchLockRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostBatchLockRequest) Description() {
    fmt.Println(strings.Trim("Batch locking files that meet specific matching conditions.", " "))
}