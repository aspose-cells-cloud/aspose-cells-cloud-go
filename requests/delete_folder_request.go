package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type DeleteFolderRequest struct {
    path string

    recursive *bool
    storageName string
}

func NewDeleteFolderRequest(path string, opts ...RequestOption) *DeleteFolderRequest {
    req := &DeleteFolderRequest{
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

    if val, ok := cfg.Params["recursive"].(*bool); ok {
        req.recursive = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *DeleteFolderRequest) GetMethod() string {
    return "DELETE"
}

func (request *DeleteFolderRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *DeleteFolderRequest) GetPath() string {
    localVarPath := "/cells/storage/folder/{path}"
    localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", request.path), -1)
    return localVarPath
}

func (request *DeleteFolderRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    if request.recursive != nil {
        localVarQueryParams.Add("recursive", fmt.Sprintf("%v", *request.recursive))
    }
    return localVarQueryParams
}

func (request *DeleteFolderRequest) GetJSONBody() interface{} {
    return nil
}

func (request *DeleteFolderRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *DeleteFolderRequest) Description() {
    fmt.Println(strings.Trim("DeleteFolder", " "))
}