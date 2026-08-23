package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type GetWorkbookNamesRequest struct {
    name string

    folder string
    storageName string
}

func NewGetWorkbookNamesRequest(name string, opts ...RequestOption) *GetWorkbookNamesRequest {
    req := &GetWorkbookNamesRequest{
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

func (request *GetWorkbookNamesRequest) GetMethod() string {
    return "GET"
}

func (request *GetWorkbookNamesRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *GetWorkbookNamesRequest) GetPath() string {
    localVarPath := "/cells/{name}/names"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    return localVarPath
}

func (request *GetWorkbookNamesRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *GetWorkbookNamesRequest) GetJSONBody() interface{} {
    return nil
}

func (request *GetWorkbookNamesRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *GetWorkbookNamesRequest) Description() {
    fmt.Println(strings.Trim("Retrieve named ranges in the workbook.", " "))
}