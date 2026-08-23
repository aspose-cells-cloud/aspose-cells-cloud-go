package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type GetFileVersionsRequest struct {
    path string

    storageName string
}

func NewGetFileVersionsRequest(path string, opts ...RequestOption) *GetFileVersionsRequest {
    req := &GetFileVersionsRequest{
        path: path,
    }
    if req.path == "" {
        return nil
    }

    cfg := &requestConfig{
        Params: make(map[string]interface{}),
    }
    for _, opt := range opts {
        opt.apply(cfg)
    }

    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *GetFileVersionsRequest) GetMethod() string {
    return "GET"
}

func (request *GetFileVersionsRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *GetFileVersionsRequest) GetPath() string {
    localVarPath := "/cells/storage/version/{path}"
    localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", request.path), -1)
    return localVarPath
}

func (request *GetFileVersionsRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *GetFileVersionsRequest) GetJSONBody() interface{} {
    return nil
}

func (request *GetFileVersionsRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *GetFileVersionsRequest) Description() {
    fmt.Println(strings.Trim("GetFileVersions", " "))
}