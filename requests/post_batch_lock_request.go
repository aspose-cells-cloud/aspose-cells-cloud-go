package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostBatchLockRequest struct {
    batchLockRequest *models.BatchLockRequest

    extraQueryParameters map[string]string
}

func NewPostBatchLockRequest(batchLockRequest *models.BatchLockRequest, opts ...RequestOption) *PostBatchLockRequest {
    req := &PostBatchLockRequest{
        batchLockRequest: batchLockRequest,
    }
    if req.batchLockRequest == nil {
        return nil
    }

    cfg := &requestConfig{
        Params: make(map[string]interface{}),
    }
    for _, opt := range opts {
        opt.apply(cfg)
    }

    if len(cfg.extraQueryParams) > 0 {
        if req.extraQueryParameters == nil {
            req.extraQueryParameters = make(map[string]string)
        }
        for k, v := range cfg.extraQueryParams {
            req.extraQueryParameters[k] = v
        }
    }

    return req
}

func (request *PostBatchLockRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *PostBatchLockRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
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
    for k, v := range request.extraQueryParameters {
        localVarQueryParams.Add(k, v)
    }
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