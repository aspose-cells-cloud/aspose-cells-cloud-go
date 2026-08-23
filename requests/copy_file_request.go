package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type CopyFileRequest struct {
    destPath string
    srcPath string

    destStorageName string
    srcStorageName string
    versionId string
}

func NewCopyFileRequest(destPath string, srcPath string, opts ...RequestOption) *CopyFileRequest {
    req := &CopyFileRequest{
        destPath: destPath,
        srcPath: srcPath,
    }
    if req.destPath == "" {
        return nil
    }
    if req.srcPath == "" {
        return nil
    }

    cfg := &requestConfig{
        Params: make(map[string]interface{}),
    }
    for _, opt := range opts {
        opt.apply(cfg)
    }

    if val, ok := cfg.Params["destStorageName"].(string); ok {
        req.destStorageName = val
    }
    if val, ok := cfg.Params["srcStorageName"].(string); ok {
        req.srcStorageName = val
    }
    if val, ok := cfg.Params["versionId"].(string); ok {
        req.versionId = val
    }

    return req
}

func (request *CopyFileRequest) GetMethod() string {
    return "PUT"
}

func (request *CopyFileRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *CopyFileRequest) GetPath() string {
    localVarPath := "/cells/storage/file/copy/{srcPath}"
    localVarPath = strings.Replace(localVarPath, "{"+"srcPath"+"}", fmt.Sprintf("%v", request.srcPath), -1)
    return localVarPath
}

func (request *CopyFileRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("destPath", fmt.Sprintf("%v", request.destPath))
    if request.srcStorageName != "" {
        localVarQueryParams.Add("srcStorageName", fmt.Sprintf("%v", request.srcStorageName))
    }
    if request.destStorageName != "" {
        localVarQueryParams.Add("destStorageName", fmt.Sprintf("%v", request.destStorageName))
    }
    if request.versionId != "" {
        localVarQueryParams.Add("versionId", fmt.Sprintf("%v", request.versionId))
    }
    return localVarQueryParams
}

func (request *CopyFileRequest) GetJSONBody() interface{} {
    return nil
}

func (request *CopyFileRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *CopyFileRequest) Description() {
    fmt.Println(strings.Trim("CopyFile", " "))
}