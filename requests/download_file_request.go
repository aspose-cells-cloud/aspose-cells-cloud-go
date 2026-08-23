package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type DownloadFileRequest struct {
    path string

    storageName string
    versionId string
}

func NewDownloadFileRequest(path string, opts ...RequestOption) *DownloadFileRequest {
    req := &DownloadFileRequest{
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

func (request *DownloadFileRequest) GetMethod() string {
    return "GET"
}

func (request *DownloadFileRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *DownloadFileRequest) GetPath() string {
    localVarPath := "/cells/storage/file/{path}"
    localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", request.path), -1)
    return localVarPath
}

func (request *DownloadFileRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    if request.versionId != "" {
        localVarQueryParams.Add("versionId", fmt.Sprintf("%v", request.versionId))
    }
    return localVarQueryParams
}

func (request *DownloadFileRequest) GetJSONBody() interface{} {
    return nil
}

func (request *DownloadFileRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *DownloadFileRequest) Description() {
    fmt.Println(strings.Trim("DownloadFile", " "))
}