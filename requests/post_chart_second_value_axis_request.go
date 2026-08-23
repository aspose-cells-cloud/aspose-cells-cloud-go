package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostChartSecondValueAxisRequest struct {
    axis *models.Axis
    chartIndex int
    name string
    sheetName string

    folder string
    storageName string
}

func NewPostChartSecondValueAxisRequest(axis *models.Axis, chartIndex int, name string, sheetName string, opts ...RequestOption) *PostChartSecondValueAxisRequest {
    req := &PostChartSecondValueAxisRequest{
        axis: axis,
        chartIndex: chartIndex,
        name: name,
        sheetName: sheetName,
    }
    if req.axis == nil {
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

func (request *PostChartSecondValueAxisRequest) GetMethod() string {
    return "POST"
}

func (request *PostChartSecondValueAxisRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostChartSecondValueAxisRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/charts/{chartIndex}/secondvalueaxis"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"chartIndex"+"}", fmt.Sprintf("%v", request.chartIndex), -1)
    return localVarPath
}

func (request *PostChartSecondValueAxisRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostChartSecondValueAxisRequest) GetJSONBody() interface{} {
    return &request.axis
}

func (request *PostChartSecondValueAxisRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostChartSecondValueAxisRequest) Description() {
    fmt.Println(strings.Trim("Update chart sencond value axis in the chart.", " "))
}