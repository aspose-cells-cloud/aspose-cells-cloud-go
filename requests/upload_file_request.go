package requests

import (
    "fmt"
    "net/url"
    "path/filepath"
    "strings"
)

type UploadFileRequest struct {
    path string
    UploadFiles string
    UploadFilesData []byte
    UploadFilesName string

    storageName string

    extraQueryParameters map[string]string
}

func NewUploadFileRequest(path string, UploadFiles string, opts ...RequestOption) *UploadFileRequest {
    req := &UploadFileRequest{
        path: path,
        UploadFiles: UploadFiles,
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
    if len(cfg.extraQueryParams) > 0 {
        if req.extraQueryParameters == nil {
            req.extraQueryParameters = make(map[string]string)
        }
        for k, v := range cfg.extraQueryParams {
            req.extraQueryParameters[k] = v
        }
    }

    return req
}

func (request *UploadFileRequest) SetUploadFilesBytes(data []byte, name string) {
    if name == "" {
        name = "UploadFiles"
    }
    request.UploadFilesData = data
    request.UploadFilesName = name
}

func (request *UploadFileRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *UploadFileRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
}

func (request *UploadFileRequest) GetMethod() string {
    return "PUT"
}

func (request *UploadFileRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "multipart/form-data"
    return localVarHeaderParams
}

func (request *UploadFileRequest) GetPath() string {
    localVarPath := "/cells/storage/file/{path}"
    localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", request.path), -1)
    return localVarPath
}

func (request *UploadFileRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    for k, v := range request.extraQueryParameters {
        localVarQueryParams.Add(k, v)
    }
    return localVarQueryParams
}

func (request *UploadFileRequest) GetJSONBody() interface{} {
    return nil
}

func (request *UploadFileRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    if request.UploadFilesData != nil {
        localVarFormParams[request.UploadFilesName] = request.UploadFilesData
    } else if request.UploadFiles != "" {
        localVarFormParams["@"+filepath.Base(request.UploadFiles)] = request.UploadFiles
    }
    return localVarFormParams
}

func (request *UploadFileRequest) Description() {
    fmt.Println(strings.Trim("UploadFile", " "))
}