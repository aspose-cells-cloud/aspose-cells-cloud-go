package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type CreateFolderRequest struct {
    path string

    storageName string
}

func NewCreateFolderRequest(path string, opts ...RequestOption) *CreateFolderRequest {
    req := &CreateFolderRequest{
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

func (request *CreateFolderRequest) GetMethod() string {
    return "PUT"
}

func (request *CreateFolderRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *CreateFolderRequest) GetPath() string {
    localVarPath := "/cells/storage/folder/{path}"
    localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", request.path), -1)
    return localVarPath
}

func (request *CreateFolderRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *CreateFolderRequest) GetJSONBody() interface{} {
    return nil
}

func (request *CreateFolderRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *CreateFolderRequest) Description() {
    fmt.Println(strings.Trim("CreateFolder", " "))
}