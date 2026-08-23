package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type GetNamedRangeValueRequest struct {
    name string
    namerange string

    folder string
    storageName string
}

func NewGetNamedRangeValueRequest(name string, namerange string, opts ...RequestOption) *GetNamedRangeValueRequest {
    req := &GetNamedRangeValueRequest{
        name: name,
        namerange: namerange,
    }
    if req.name == "" {
        return nil
    }
    if req.namerange == "" {
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

func (request *GetNamedRangeValueRequest) GetMethod() string {
    return "GET"
}

func (request *GetNamedRangeValueRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *GetNamedRangeValueRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/ranges/{namerange}/value"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"namerange"+"}", fmt.Sprintf("%v", request.namerange), -1)
    return localVarPath
}

func (request *GetNamedRangeValueRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *GetNamedRangeValueRequest) GetJSONBody() interface{} {
    return nil
}

func (request *GetNamedRangeValueRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *GetNamedRangeValueRequest) Description() {
    fmt.Println(strings.Trim("Retrieve values in range.", " "))
}