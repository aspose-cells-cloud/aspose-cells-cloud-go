package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type DeleteFileRequest struct {
    path string

    storageName string
    versionId string
}

func NewDeleteFileRequest(path string, opts ...RequestOption) *DeleteFileRequest {
    req := &DeleteFileRequest{
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
    if val, ok := cfg.Params["versionId"].(string); ok {
        req.versionId = val
    }

    return req
}

func (request *DeleteFileRequest) GetMethod() string {
    return "DELETE"
}

func (request *DeleteFileRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *DeleteFileRequest) GetPath() string {
    localVarPath := "/cells/storage/file/{path}"
    localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", request.path), -1)
    return localVarPath
}

func (request *DeleteFileRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    if request.versionId != "" {
        localVarQueryParams.Add("versionId", fmt.Sprintf("%v", request.versionId))
    }
    return localVarQueryParams
}

func (request *DeleteFileRequest) GetJSONBody() interface{} {
    return nil
}

func (request *DeleteFileRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *DeleteFileRequest) Description() {
    fmt.Println(strings.Trim("DeleteFile", " "))
}