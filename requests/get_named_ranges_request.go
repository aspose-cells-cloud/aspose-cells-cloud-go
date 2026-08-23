package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type GetNamedRangesRequest struct {
    name string

    folder string
    storageName string
}

func NewGetNamedRangesRequest(name string, opts ...RequestOption) *GetNamedRangesRequest {
    req := &GetNamedRangesRequest{
        name: name,
    }
    if req.name == "" {
        return nil
    }

    cfg := &requestConfig{
        Params: make(map[string]interface{}),
    }
    for _, opt := range opts {
        opt.apply(cfg)
    }

    if val, ok := cfg.Params["folder"].(string); ok {
        req.folder = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *GetNamedRangesRequest) GetMethod() string {
    return "GET"
}

func (request *GetNamedRangesRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *GetNamedRangesRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/ranges"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    return localVarPath
}

func (request *GetNamedRangesRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *GetNamedRangesRequest) GetJSONBody() interface{} {
    return nil
}

func (request *GetNamedRangesRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *GetNamedRangesRequest) Description() {
    fmt.Println(strings.Trim("Retrieve descriptions of ranges in the worksheets.", " "))
}