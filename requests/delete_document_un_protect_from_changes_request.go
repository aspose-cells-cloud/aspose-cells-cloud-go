package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type DeleteDocumentUnProtectFromChangesRequest struct {
    name string

    folder string
    storageName string
}

func NewDeleteDocumentUnProtectFromChangesRequest(name string, opts ...RequestOption) *DeleteDocumentUnProtectFromChangesRequest {
    req := &DeleteDocumentUnProtectFromChangesRequest{
        name: name,
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

func (request *DeleteDocumentUnProtectFromChangesRequest) GetMethod() string {
    return "DELETE"
}

func (request *DeleteDocumentUnProtectFromChangesRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *DeleteDocumentUnProtectFromChangesRequest) GetPath() string {
    localVarPath := "/cells/{name}/writeProtection"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    return localVarPath
}

func (request *DeleteDocumentUnProtectFromChangesRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *DeleteDocumentUnProtectFromChangesRequest) GetJSONBody() interface{} {
    return nil
}

func (request *DeleteDocumentUnProtectFromChangesRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *DeleteDocumentUnProtectFromChangesRequest) Description() {
    fmt.Println(strings.Trim("Excel file cancel write protection.", " "))
}