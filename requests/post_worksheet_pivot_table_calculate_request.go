package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PostWorksheetPivotTableCalculateRequest struct {
    name string
    pivotTableIndex int
    sheetName string

    folder string
    storageName string
}

func NewPostWorksheetPivotTableCalculateRequest(name string, pivotTableIndex int, sheetName string, opts ...RequestOption) *PostWorksheetPivotTableCalculateRequest {
    req := &PostWorksheetPivotTableCalculateRequest{
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
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *PostWorksheetPivotTableCalculateRequest) GetMethod() string {
    return "POST"
}

func (request *PostWorksheetPivotTableCalculateRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostWorksheetPivotTableCalculateRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/Calculate"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"pivotTableIndex"+"}", fmt.Sprintf("%v", request.pivotTableIndex), -1)
    return localVarPath
}

func (request *PostWorksheetPivotTableCalculateRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostWorksheetPivotTableCalculateRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostWorksheetPivotTableCalculateRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostWorksheetPivotTableCalculateRequest) Description() {
    fmt.Println(strings.Trim("Calculate pivottable's data to cells.", " "))
}