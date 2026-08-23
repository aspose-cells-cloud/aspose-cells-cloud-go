package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostChartCategoryAxisRequest struct {
    axis *models.Axis
    chartIndex int
    name string
    sheetName string

    folder string
    storageName string
}

func NewPostChartCategoryAxisRequest(axis *models.Axis, chartIndex int, name string, sheetName string, opts ...RequestOption) *PostChartCategoryAxisRequest {
    req := &PostChartCategoryAxisRequest{
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

func (request *PostChartCategoryAxisRequest) GetMethod() string {
    return "POST"
}

func (request *PostChartCategoryAxisRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostChartCategoryAxisRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/charts/{chartIndex}/categoryaxis"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"chartIndex"+"}", fmt.Sprintf("%v", request.chartIndex), -1)
    return localVarPath
}

func (request *PostChartCategoryAxisRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostChartCategoryAxisRequest) GetJSONBody() interface{} {
    return &request.axis
}

func (request *PostChartCategoryAxisRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostChartCategoryAxisRequest) Description() {
    fmt.Println(strings.Trim("Update chart category axis in the chart.", " "))
}