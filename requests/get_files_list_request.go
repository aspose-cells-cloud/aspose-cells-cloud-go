package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type GetFilesListRequest struct {

    path string
    storageName string
}

func NewGetFilesListRequest(opts ...RequestOption) *GetFilesListRequest {
    req := &GetFilesListRequest{
    }
    cfg := &requestConfig{
        Params: make(map[string]interface{}),
    }
    for _, opt := range opts {
        opt.apply(cfg)
    }

    if val, ok := cfg.Params["path"].(string); ok {
        req.path = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *GetFilesListRequest) GetMethod() string {
    return "GET"
}

func (request *GetFilesListRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *GetFilesListRequest) GetPath() string {
    localVarPath := "/cells/storage/folder/{path}"
    localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", request.path), -1)
    return localVarPath
}

func (request *GetFilesListRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *GetFilesListRequest) GetJSONBody() interface{} {
    return nil
}

func (request *GetFilesListRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *GetFilesListRequest) Description() {
    fmt.Println(strings.Trim("GetFilesList", " "))
}