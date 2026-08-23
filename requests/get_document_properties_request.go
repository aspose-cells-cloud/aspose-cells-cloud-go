package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type GetDocumentPropertiesRequest struct {
    name string

    folder string
    storageName string
    _type string
}

func NewGetDocumentPropertiesRequest(name string, opts ...RequestOption) *GetDocumentPropertiesRequest {
    req := &GetDocumentPropertiesRequest{
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
    if val, ok := cfg.Params["type"].(string); ok {
        req._type = val
    }

    return req
}

func (request *GetDocumentPropertiesRequest) GetMethod() string {
    return "GET"
}

func (request *GetDocumentPropertiesRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *GetDocumentPropertiesRequest) GetPath() string {
    localVarPath := "/cells/{name}/documentproperties"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    return localVarPath
}

func (request *GetDocumentPropertiesRequest) GetQueryParameters() url.Values {
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

func (request *GetDocumentPropertiesRequest) GetJSONBody() interface{} {
    return nil
}

func (request *GetDocumentPropertiesRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *GetDocumentPropertiesRequest) Description() {
    fmt.Println(strings.Trim("Retrieve descriptions of Excel file properties.", " "))
}