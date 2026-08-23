package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type GetDocumentPropertyRequest struct {
    name string
    propertyName string

    folder string
    storageName string
}

func NewGetDocumentPropertyRequest(name string, propertyName string, opts ...RequestOption) *GetDocumentPropertyRequest {
    req := &GetDocumentPropertyRequest{
        name: name,
        propertyName: propertyName,
    }
    if req.name == "" {
        return nil
    }
    if req.propertyName == "" {
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

func (request *GetDocumentPropertyRequest) GetMethod() string {
    return "GET"
}

func (request *GetDocumentPropertyRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *GetDocumentPropertyRequest) GetPath() string {
    localVarPath := "/cells/{name}/documentproperties/{propertyName}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"propertyName"+"}", fmt.Sprintf("%v", request.propertyName), -1)
    return localVarPath
}

func (request *GetDocumentPropertyRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *GetDocumentPropertyRequest) GetJSONBody() interface{} {
    return nil
}

func (request *GetDocumentPropertyRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *GetDocumentPropertyRequest) Description() {
    fmt.Println(strings.Trim("Get Excel property by name.", " "))
}