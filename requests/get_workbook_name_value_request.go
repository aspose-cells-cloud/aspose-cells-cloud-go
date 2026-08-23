package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type GetWorkbookNameValueRequest struct {
    name string
    nameName string

    folder string
    storageName string
}

func NewGetWorkbookNameValueRequest(name string, nameName string, opts ...RequestOption) *GetWorkbookNameValueRequest {
    req := &GetWorkbookNameValueRequest{
        name: name,
        nameName: nameName,
    }
    if req.name == "" {
        return nil
    }
    if req.nameName == "" {
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

func (request *GetWorkbookNameValueRequest) GetMethod() string {
    return "GET"
}

func (request *GetWorkbookNameValueRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *GetWorkbookNameValueRequest) GetPath() string {
    localVarPath := "/cells/{name}/names/{nameName}/value"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nameName"+"}", fmt.Sprintf("%v", request.nameName), -1)
    return localVarPath
}

func (request *GetWorkbookNameValueRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *GetWorkbookNameValueRequest) GetJSONBody() interface{} {
    return nil
}

func (request *GetWorkbookNameValueRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *GetWorkbookNameValueRequest) Description() {
    fmt.Println(strings.Trim("Retrieve the value of a named range in the workbook.", " "))
}