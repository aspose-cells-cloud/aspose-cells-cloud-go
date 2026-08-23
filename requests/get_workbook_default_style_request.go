package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type GetWorkbookDefaultStyleRequest struct {
    name string

    folder string
    storageName string
}

func NewGetWorkbookDefaultStyleRequest(name string, opts ...RequestOption) *GetWorkbookDefaultStyleRequest {
    req := &GetWorkbookDefaultStyleRequest{
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

func (request *GetWorkbookDefaultStyleRequest) GetMethod() string {
    return "GET"
}

func (request *GetWorkbookDefaultStyleRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *GetWorkbookDefaultStyleRequest) GetPath() string {
    localVarPath := "/cells/{name}/defaultstyle"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    return localVarPath
}

func (request *GetWorkbookDefaultStyleRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *GetWorkbookDefaultStyleRequest) GetJSONBody() interface{} {
    return nil
}

func (request *GetWorkbookDefaultStyleRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *GetWorkbookDefaultStyleRequest) Description() {
    fmt.Println(strings.Trim("Retrieve the description of the default style for the workbook .", " "))
}