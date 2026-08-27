package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type CheckWrokbookExternalReferenceRequest struct {
    checkExternalReferenceOptions *models.CheckExternalReferenceOptions

    extraQueryParameters map[string]string
}

func NewCheckWrokbookExternalReferenceRequest(checkExternalReferenceOptions *models.CheckExternalReferenceOptions, opts ...RequestOption) *CheckWrokbookExternalReferenceRequest {
    req := &CheckWrokbookExternalReferenceRequest{
        checkExternalReferenceOptions: checkExternalReferenceOptions,
    }
    if req.checkExternalReferenceOptions == nil {
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

func (request *CheckWrokbookExternalReferenceRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *CheckWrokbookExternalReferenceRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
}

func (request *CheckWrokbookExternalReferenceRequest) GetMethod() string {
    return "POST"
}

func (request *CheckWrokbookExternalReferenceRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *CheckWrokbookExternalReferenceRequest) GetPath() string {
    localVarPath := "/cells/checkexternalreference"
    return localVarPath
}

func (request *CheckWrokbookExternalReferenceRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    for k, v := range request.extraQueryParameters {
        localVarQueryParams.Add(k, v)
    }
    return localVarQueryParams
}

func (request *CheckWrokbookExternalReferenceRequest) GetJSONBody() interface{} {
    return &request.checkExternalReferenceOptions
}

func (request *CheckWrokbookExternalReferenceRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *CheckWrokbookExternalReferenceRequest) Description() {
    fmt.Println(strings.Trim("Export Excel internal elements or the workbook itself to various format files.", " "))
}