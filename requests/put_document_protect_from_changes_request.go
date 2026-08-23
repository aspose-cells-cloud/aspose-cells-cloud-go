package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PutDocumentProtectFromChangesRequest struct {
    name string
    password *models.PasswordRequest

    folder string
    storageName string
}

func NewPutDocumentProtectFromChangesRequest(name string, password *models.PasswordRequest, opts ...RequestOption) *PutDocumentProtectFromChangesRequest {
    req := &PutDocumentProtectFromChangesRequest{
        name: name,
        password: password,
    }
    if req.name == "" {
        return nil
    }
    if req.password == nil {
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

func (request *PutDocumentProtectFromChangesRequest) GetMethod() string {
    return "PUT"
}

func (request *PutDocumentProtectFromChangesRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PutDocumentProtectFromChangesRequest) GetPath() string {
    localVarPath := "/cells/{name}/writeProtection"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    return localVarPath
}

func (request *PutDocumentProtectFromChangesRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PutDocumentProtectFromChangesRequest) GetJSONBody() interface{} {
    return &request.password
}

func (request *PutDocumentProtectFromChangesRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PutDocumentProtectFromChangesRequest) Description() {
    fmt.Println(strings.Trim("Excel file write protection.", " "))
}