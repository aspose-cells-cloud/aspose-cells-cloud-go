package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type GetWorksheetChartsRequest struct {
    name string
    sheetName string

    folder string
    storageName string
}

func NewGetWorksheetChartsRequest(name string, sheetName string, opts ...RequestOption) *GetWorksheetChartsRequest {
    req := &GetWorksheetChartsRequest{
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

func (request *GetWorksheetChartsRequest) GetMethod() string {
    return "GET"
}

func (request *GetWorksheetChartsRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *GetWorksheetChartsRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/charts"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *GetWorksheetChartsRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *GetWorksheetChartsRequest) GetJSONBody() interface{} {
    return nil
}

func (request *GetWorksheetChartsRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *GetWorksheetChartsRequest) Description() {
    fmt.Println(strings.Trim("Retrieve descriptions of charts in the worksheet.", " "))
}