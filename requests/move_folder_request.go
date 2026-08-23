package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type MoveFolderRequest struct {
    destPath string
    srcPath string

    destStorageName string
    srcStorageName string
}

func NewMoveFolderRequest(destPath string, srcPath string, opts ...RequestOption) *MoveFolderRequest {
    req := &MoveFolderRequest{
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

func (request *MoveFolderRequest) GetMethod() string {
    return "PUT"
}

func (request *MoveFolderRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *MoveFolderRequest) GetPath() string {
    localVarPath := "/cells/storage/folder/move/{srcPath}"
    localVarPath = strings.Replace(localVarPath, "{"+"srcPath"+"}", fmt.Sprintf("%v", request.srcPath), -1)
    return localVarPath
}

func (request *MoveFolderRequest) GetQueryParameters() url.Values {
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

func (request *MoveFolderRequest) GetJSONBody() interface{} {
    return nil
}

func (request *MoveFolderRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *MoveFolderRequest) Description() {
    fmt.Println(strings.Trim("MoveFolder", " "))
}