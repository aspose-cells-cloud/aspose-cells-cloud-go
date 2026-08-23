package requests

import (
    "fmt"
    "net/url"
    "path/filepath"
    "strings"

    "asposecellscloud/models"
)

type PostProtectRequest struct {
    File string
    FileData []byte
    FileName string
    protectWorkbookRequest *models.ProtectWorkbookRequest

    password string
}

func NewPostProtectRequest(File string, protectWorkbookRequest *models.ProtectWorkbookRequest, opts ...RequestOption) *PostProtectRequest {
    req := &PostProtectRequest{
        File: File,
        protectWorkbookRequest: protectWorkbookRequest,
    }
    if req.protectWorkbookRequest == nil {
        return nil
    }

    cfg := &requestConfig{
        Params: make(map[string]interface{}),
    }
    for _, opt := range opts {
        opt.apply(cfg)
    }

    if val, ok := cfg.Params["password"].(string); ok {
        req.password = val
    }

    return req
}

func (request *PostProtectRequest) SetFileBytes(data []byte, name string) {
    if name == "" {
        name = "File"
    }
    request.FileData = data
    request.FileName = name
}

func (request *PostProtectRequest) GetMethod() string {
    return "POST"
}

func (request *PostProtectRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "multipart/form-data"
    return localVarHeaderParams
}

func (request *PostProtectRequest) GetPath() string {
    localVarPath := "/cells/protect"
    return localVarPath
}

func (request *PostProtectRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.password != "" {
        localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
    }
    return localVarQueryParams
}

func (request *PostProtectRequest) GetJSONBody() interface{} {
    return &request.protectWorkbookRequest
}

func (request *PostProtectRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    if request.FileData != nil {
        localVarFormParams[request.FileName] = request.FileData
    } else if request.File != "" {
        localVarFormParams["@"+filepath.Base(request.File)] = request.File
    }
    return localVarFormParams
}

func (request *PostProtectRequest) Description() {
    fmt.Println(strings.Trim("Excel files encryption.", " "))
}