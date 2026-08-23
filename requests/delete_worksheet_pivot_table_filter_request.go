package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type DeleteWorksheetPivotTableFilterRequest struct {
    fieldIndex int
    name string
    pivotTableIndex int
    sheetName string

    folder string
    needReCalculate *bool
    storageName string
}

func NewDeleteWorksheetPivotTableFilterRequest(fieldIndex int, name string, pivotTableIndex int, sheetName string, opts ...RequestOption) *DeleteWorksheetPivotTableFilterRequest {
    req := &DeleteWorksheetPivotTableFilterRequest{
        fieldIndex: fieldIndex,
        name: name,
        pivotTableIndex: pivotTableIndex,
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
    if val, ok := cfg.Params["needReCalculate"].(*bool); ok {
        req.needReCalculate = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *DeleteWorksheetPivotTableFilterRequest) GetMethod() string {
    return "DELETE"
}

func (request *DeleteWorksheetPivotTableFilterRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *DeleteWorksheetPivotTableFilterRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/PivotFilters/{fieldIndex}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"pivotTableIndex"+"}", fmt.Sprintf("%v", request.pivotTableIndex), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"fieldIndex"+"}", fmt.Sprintf("%v", request.fieldIndex), -1)
    return localVarPath
}

func (request *DeleteWorksheetPivotTableFilterRequest) GetQueryParameters() url.Values {
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

func (request *DeleteWorksheetPivotTableFilterRequest) GetJSONBody() interface{} {
    return nil
}

func (request *DeleteWorksheetPivotTableFilterRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *DeleteWorksheetPivotTableFilterRequest) Description() {
    fmt.Println(strings.Trim("Delete a pivot filter in the PivotTable.", " "))
}