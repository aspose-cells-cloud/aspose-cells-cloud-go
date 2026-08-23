package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type GetWorksheetChartRequest struct {
    chartNumber int
    name string
    sheetName string

    folder string
    format string
    storageName string
}

func NewGetWorksheetChartRequest(chartNumber int, name string, sheetName string, opts ...RequestOption) *GetWorksheetChartRequest {
    req := &GetWorksheetChartRequest{
        chartNumber: chartNumber,
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
    if val, ok := cfg.Params["format"].(string); ok {
        req.format = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *GetWorksheetChartRequest) GetMethod() string {
    return "GET"
}

func (request *GetWorksheetChartRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *GetWorksheetChartRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/charts/{chartNumber}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"chartNumber"+"}", fmt.Sprintf("%v", request.chartNumber), -1)
    return localVarPath
}

func (request *GetWorksheetChartRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.format != "" {
        localVarQueryParams.Add("format", fmt.Sprintf("%v", request.format))
    }
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *GetWorksheetChartRequest) GetJSONBody() interface{} {
    return nil
}

func (request *GetWorksheetChartRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *GetWorksheetChartRequest) Description() {
    fmt.Println(strings.Trim("Retrieve the chart in a specified format.", " "))
}