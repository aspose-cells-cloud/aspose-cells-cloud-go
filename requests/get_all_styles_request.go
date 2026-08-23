package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type GetAllStylesRequest struct {
    name string

    folder string
    storageName string
}

func NewGetAllStylesRequest(name string, opts ...RequestOption) *GetAllStylesRequest {
    req := &GetAllStylesRequest{
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

func (request *GetAllStylesRequest) GetMethod() string {
    return "GET"
}

func (request *GetAllStylesRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *GetAllStylesRequest) GetPath() string {
    localVarPath := "/cells/{name}/allstyles"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    return localVarPath
}

func (request *GetAllStylesRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *GetAllStylesRequest) GetJSONBody() interface{} {
    return nil
}

func (request *GetAllStylesRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *GetAllStylesRequest) Description() {
    fmt.Println(strings.Trim("Get all style in the workbook.", " "))
}