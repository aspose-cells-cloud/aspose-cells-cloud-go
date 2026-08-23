package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostWorksheetChartLegendRequest struct {
    chartIndex int
    legend *models.Legend
    name string
    sheetName string

    folder string
    storageName string
}

func NewPostWorksheetChartLegendRequest(chartIndex int, legend *models.Legend, name string, sheetName string, opts ...RequestOption) *PostWorksheetChartLegendRequest {
    req := &PostWorksheetChartLegendRequest{
        chartIndex: chartIndex,
        legend: legend,
        name: name,
        sheetName: sheetName,
    }
    if req.legend == nil {
        return nil
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

func (request *PostWorksheetChartLegendRequest) GetMethod() string {
    return "POST"
}

func (request *PostWorksheetChartLegendRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostWorksheetChartLegendRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/charts/{chartIndex}/legend"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"chartIndex"+"}", fmt.Sprintf("%v", request.chartIndex), -1)
    return localVarPath
}

func (request *PostWorksheetChartLegendRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostWorksheetChartLegendRequest) GetJSONBody() interface{} {
    return &request.legend
}

func (request *PostWorksheetChartLegendRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostWorksheetChartLegendRequest) Description() {
    fmt.Println(strings.Trim("Update chart legend in the worksheet.", " "))
}