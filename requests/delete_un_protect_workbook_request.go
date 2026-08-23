package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type DeleteUnProtectWorkbookRequest struct {
    name string
    password string

    folder string
    storageName string
}

func NewDeleteUnProtectWorkbookRequest(name string, password string, opts ...RequestOption) *DeleteUnProtectWorkbookRequest {
    req := &DeleteUnProtectWorkbookRequest{
        name: name,
        password: password,
    }
    if req.name == "" {
        return nil
    }
    if req.password == "" {
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

func (request *DeleteUnProtectWorkbookRequest) GetMethod() string {
    return "DELETE"
}

func (request *DeleteUnProtectWorkbookRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *DeleteUnProtectWorkbookRequest) GetPath() string {
    localVarPath := "/cells/{name}/protection"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    return localVarPath
}

func (request *DeleteUnProtectWorkbookRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *DeleteUnProtectWorkbookRequest) GetJSONBody() interface{} {
    return nil
}

func (request *DeleteUnProtectWorkbookRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *DeleteUnProtectWorkbookRequest) Description() {
    fmt.Println(strings.Trim("Excel unprotection.", " "))
}