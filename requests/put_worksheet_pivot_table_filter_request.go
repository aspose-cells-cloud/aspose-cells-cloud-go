package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PutWorksheetPivotTableFilterRequest struct {
    filter *models.PivotFilter
    name string
    pivotTableIndex int
    sheetName string

    folder string
    needReCalculate *bool
    storageName string
}

func NewPutWorksheetPivotTableFilterRequest(filter *models.PivotFilter, name string, pivotTableIndex int, sheetName string, opts ...RequestOption) *PutWorksheetPivotTableFilterRequest {
    req := &PutWorksheetPivotTableFilterRequest{
        filter: filter,
        name: name,
        pivotTableIndex: pivotTableIndex,
        sheetName: sheetName,
    }
    if req.filter == nil {
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
    if val, ok := cfg.Params["needReCalculate"].(*bool); ok {
        req.needReCalculate = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *PutWorksheetPivotTableFilterRequest) GetMethod() string {
    return "PUT"
}

func (request *PutWorksheetPivotTableFilterRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PutWorksheetPivotTableFilterRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/PivotFilters"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"pivotTableIndex"+"}", fmt.Sprintf("%v", request.pivotTableIndex), -1)
    return localVarPath
}

func (request *PutWorksheetPivotTableFilterRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.needReCalculate != nil {
        localVarQueryParams.Add("needReCalculate", fmt.Sprintf("%v", *request.needReCalculate))
    }
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PutWorksheetPivotTableFilterRequest) GetJSONBody() interface{} {
    return &request.filter
}

func (request *PutWorksheetPivotTableFilterRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PutWorksheetPivotTableFilterRequest) Description() {
    fmt.Println(strings.Trim("Add a pivot filter to the PivotTable.", " "))
}