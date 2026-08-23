package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type MoveFileRequest struct {
    destPath string
    srcPath string

    destStorageName string
    srcStorageName string
    versionId string
}

func NewMoveFileRequest(destPath string, srcPath string, opts ...RequestOption) *MoveFileRequest {
    req := &MoveFileRequest{
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

func (request *MoveFileRequest) GetMethod() string {
    return "PUT"
}

func (request *MoveFileRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *MoveFileRequest) GetPath() string {
    localVarPath := "/cells/storage/file/move/{srcPath}"
    localVarPath = strings.Replace(localVarPath, "{"+"srcPath"+"}", fmt.Sprintf("%v", request.srcPath), -1)
    return localVarPath
}

func (request *MoveFileRequest) GetQueryParameters() url.Values {
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

func (request *MoveFileRequest) GetJSONBody() interface{} {
    return nil
}

func (request *MoveFileRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *MoveFileRequest) Description() {
    fmt.Println(strings.Trim("MoveFile", " "))
}