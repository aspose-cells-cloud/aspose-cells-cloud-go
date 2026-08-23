package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type DeleteDocumentPropertyRequest struct {
    name string
    propertyName string

    folder string
    storageName string
    _type string
}

func NewDeleteDocumentPropertyRequest(name string, propertyName string, opts ...RequestOption) *DeleteDocumentPropertyRequest {
    req := &DeleteDocumentPropertyRequest{
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
    if val, ok := cfg.Params["type"].(string); ok {
        req._type = val
    }

    return req
}

func (request *DeleteDocumentPropertyRequest) GetMethod() string {
    return "DELETE"
}

func (request *DeleteDocumentPropertyRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *DeleteDocumentPropertyRequest) GetPath() string {
    localVarPath := "/cells/{name}/documentproperties/{propertyName}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"propertyName"+"}", fmt.Sprintf("%v", request.propertyName), -1)
    return localVarPath
}

func (request *DeleteDocumentPropertyRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request._type != "" {
        localVarQueryParams.Add("type", fmt.Sprintf("%v", request._type))
    }
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *DeleteDocumentPropertyRequest) GetJSONBody() interface{} {
    return nil
}

func (request *DeleteDocumentPropertyRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *DeleteDocumentPropertyRequest) Description() {
    fmt.Println(strings.Trim("Delete an Excel property.", " "))
}