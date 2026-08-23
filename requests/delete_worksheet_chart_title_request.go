package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type DeleteWorksheetChartTitleRequest struct {
    chartIndex int
    name string
    sheetName string

    folder string
    storageName string
}

func NewDeleteWorksheetChartTitleRequest(chartIndex int, name string, sheetName string, opts ...RequestOption) *DeleteWorksheetChartTitleRequest {
    req := &DeleteWorksheetChartTitleRequest{
        chartIndex: chartIndex,
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

func (request *DeleteWorksheetChartTitleRequest) GetMethod() string {
    return "DELETE"
}

func (request *DeleteWorksheetChartTitleRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *DeleteWorksheetChartTitleRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/charts/{chartIndex}/title"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"chartIndex"+"}", fmt.Sprintf("%v", request.chartIndex), -1)
    return localVarPath
}

func (request *DeleteWorksheetChartTitleRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *DeleteWorksheetChartTitleRequest) GetJSONBody() interface{} {
    return nil
}

func (request *DeleteWorksheetChartTitleRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *DeleteWorksheetChartTitleRequest) Description() {
    fmt.Println(strings.Trim("Hide chart title in the worksheet.", " "))
}