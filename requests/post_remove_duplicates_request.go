package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostRemoveDuplicatesRequest struct {
    removeDuplicatesOptions *models.RemoveDuplicatesOptions

    extraQueryParameters map[string]string
}

func NewPostRemoveDuplicatesRequest(removeDuplicatesOptions *models.RemoveDuplicatesOptions, opts ...RequestOption) *PostRemoveDuplicatesRequest {
    req := &PostRemoveDuplicatesRequest{
        removeDuplicatesOptions: removeDuplicatesOptions,
    }
    if req.removeDuplicatesOptions == nil {
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

func (request *PostRemoveDuplicatesRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *PostRemoveDuplicatesRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
}

func (request *PostRemoveDuplicatesRequest) GetMethod() string {
    return "POST"
}

func (request *PostRemoveDuplicatesRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostRemoveDuplicatesRequest) GetPath() string {
    localVarPath := "/cells/removeduplicates"
    return localVarPath
}

func (request *PostRemoveDuplicatesRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    for k, v := range request.extraQueryParameters {
        localVarQueryParams.Add(k, v)
    }
    return localVarQueryParams
}

func (request *PostRemoveDuplicatesRequest) GetJSONBody() interface{} {
    return &request.removeDuplicatesOptions
}

func (request *PostRemoveDuplicatesRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostRemoveDuplicatesRequest) Description() {
    fmt.Println(strings.Trim("Efficiently remove duplicate substrings from Excel cells. Select a range, specify delimiters, and apply options to eliminate repeated text segments.", " "))
}