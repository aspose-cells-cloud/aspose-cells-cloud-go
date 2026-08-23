package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type GetHorizontalPageBreakRequest struct {
    index int
    name string
    sheetName string

    folder string
    storageName string
}

func NewGetHorizontalPageBreakRequest(index int, name string, sheetName string, opts ...RequestOption) *GetHorizontalPageBreakRequest {
    req := &GetHorizontalPageBreakRequest{
        index: index,
        name: name,
        sheetName: sheetName,
    }
    if req.name == "" {
        return nil
    }
    if req.sheetName == "" {
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

func (request *GetHorizontalPageBreakRequest) GetMethod() string {
    return "GET"
}

func (request *GetHorizontalPageBreakRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *GetHorizontalPageBreakRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/horizontalpagebreaks/{index}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", request.index), -1)
    return localVarPath
}

func (request *GetHorizontalPageBreakRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *GetHorizontalPageBreakRequest) GetJSONBody() interface{} {
    return nil
}

func (request *GetHorizontalPageBreakRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *GetHorizontalPageBreakRequest) Description() {
    fmt.Println(strings.Trim("Retrieve a horizontal page break descripton in the worksheet.", " "))
}