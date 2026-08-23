package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostChartSecondCategoryAxisRequest struct {
    axis *models.Axis
    chartIndex int
    name string
    sheetName string

    folder string
    storageName string
}

func NewPostChartSecondCategoryAxisRequest(axis *models.Axis, chartIndex int, name string, sheetName string, opts ...RequestOption) *PostChartSecondCategoryAxisRequest {
    req := &PostChartSecondCategoryAxisRequest{
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

func (request *PostChartSecondCategoryAxisRequest) GetMethod() string {
    return "POST"
}

func (request *PostChartSecondCategoryAxisRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostChartSecondCategoryAxisRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/charts/{chartIndex}/secondcategoryaxis"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"chartIndex"+"}", fmt.Sprintf("%v", request.chartIndex), -1)
    return localVarPath
}

func (request *PostChartSecondCategoryAxisRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostChartSecondCategoryAxisRequest) GetJSONBody() interface{} {
    return &request.axis
}

func (request *PostChartSecondCategoryAxisRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostChartSecondCategoryAxisRequest) Description() {
    fmt.Println(strings.Trim("Update chart sencond category axis in the chart.", " "))
}