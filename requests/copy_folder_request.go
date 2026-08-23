package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type CopyFolderRequest struct {
    destPath string
    srcPath string

    destStorageName string
    srcStorageName string
}

func NewCopyFolderRequest(destPath string, srcPath string, opts ...RequestOption) *CopyFolderRequest {
    req := &CopyFolderRequest{
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

    return req
}

func (request *CopyFolderRequest) GetMethod() string {
    return "PUT"
}

func (request *CopyFolderRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *CopyFolderRequest) GetPath() string {
    localVarPath := "/cells/storage/folder/copy/{srcPath}"
    localVarPath = strings.Replace(localVarPath, "{"+"srcPath"+"}", fmt.Sprintf("%v", request.srcPath), -1)
    return localVarPath
}

func (request *CopyFolderRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("destPath", fmt.Sprintf("%v", request.destPath))
    if request.srcStorageName != "" {
        localVarQueryParams.Add("srcStorageName", fmt.Sprintf("%v", request.srcStorageName))
    }
    if request.destStorageName != "" {
        localVarQueryParams.Add("destStorageName", fmt.Sprintf("%v", request.destStorageName))
    }
    return localVarQueryParams
}

func (request *CopyFolderRequest) GetJSONBody() interface{} {
    return nil
}

func (request *CopyFolderRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *CopyFolderRequest) Description() {
    fmt.Println(strings.Trim("CopyFolder", " "))
}