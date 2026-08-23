package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostEncryptWorkbookRequest struct {
    encryption *models.WorkbookEncryptionRequest
    name string

    folder string
    storageName string
}

func NewPostEncryptWorkbookRequest(encryption *models.WorkbookEncryptionRequest, name string, opts ...RequestOption) *PostEncryptWorkbookRequest {
    req := &PostEncryptWorkbookRequest{
        encryption: encryption,
        name: name,
    }
    if req.encryption == nil {
        return nil
    }
    if req.name == "" {
        return nil
    }

    cfg := &requestConfig{
        Params: make(map[string]interface{}),
    }
    for _, opt := range opts {
        opt.apply(cfg)
    }

    if val, ok := cfg.Params["folder"].(string); ok {
        req.folder = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *PostEncryptWorkbookRequest) GetMethod() string {
    return "POST"
}

func (request *PostEncryptWorkbookRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostEncryptWorkbookRequest) GetPath() string {
    localVarPath := "/cells/{name}/encryption"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    return localVarPath
}

func (request *PostEncryptWorkbookRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostEncryptWorkbookRequest) GetJSONBody() interface{} {
    return &request.encryption
}

func (request *PostEncryptWorkbookRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostEncryptWorkbookRequest) Description() {
    fmt.Println(strings.Trim("Excel Encryption.", " "))
}