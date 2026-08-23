package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type GetWorksheetsRequest struct {
    name string

    folder string
    storageName string
}

func NewGetWorksheetsRequest(name string, opts ...RequestOption) *GetWorksheetsRequest {
    req := &GetWorksheetsRequest{
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

func (request *GetWorksheetsRequest) GetMethod() string {
    return "GET"
}

func (request *GetWorksheetsRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *GetWorksheetsRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    return localVarPath
}

func (request *GetWorksheetsRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *GetWorksheetsRequest) GetJSONBody() interface{} {
    return nil
}

func (request *GetWorksheetsRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *GetWorksheetsRequest) Description() {
    fmt.Println(strings.Trim("Retrieve the description of worksheets from a workbook.", " "))
}