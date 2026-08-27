package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostBatchConvertRequest struct {
    batchConvertRequest *models.BatchConvertRequest

    extraQueryParameters map[string]string
}

func NewPostBatchConvertRequest(batchConvertRequest *models.BatchConvertRequest, opts ...RequestOption) *PostBatchConvertRequest {
    req := &PostBatchConvertRequest{
        batchConvertRequest: batchConvertRequest,
    }
    if req.batchConvertRequest == nil {
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

func (request *PostBatchConvertRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *PostBatchConvertRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
}

func (request *PostBatchConvertRequest) GetMethod() string {
    return "POST"
}

func (request *PostBatchConvertRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostBatchConvertRequest) GetPath() string {
    localVarPath := "/cells/batch/convert"
    return localVarPath
}

func (request *PostBatchConvertRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    for k, v := range request.extraQueryParameters {
        localVarQueryParams.Add(k, v)
    }
    return localVarQueryParams
}

func (request *PostBatchConvertRequest) GetJSONBody() interface{} {
    return &request.batchConvertRequest
}

func (request *PostBatchConvertRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostBatchConvertRequest) Description() {
    fmt.Println(strings.Trim("Batch converting files that meet specific matching conditions.", " "))
}