package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PostWorksheetPivotTableMoveRequest struct {
    name string
    pivotTableIndex int
    sheetName string

    column *int
    destCellName string
    folder string
    row *int
    storageName string
}

func NewPostWorksheetPivotTableMoveRequest(name string, pivotTableIndex int, sheetName string, opts ...RequestOption) *PostWorksheetPivotTableMoveRequest {
    req := &PostWorksheetPivotTableMoveRequest{
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

    if val, ok := cfg.Params["column"].(*int); ok {
        req.column = val
    }
    if val, ok := cfg.Params["destCellName"].(string); ok {
        req.destCellName = val
    }
    if val, ok := cfg.Params["folder"].(string); ok {
        req.folder = val
    }
    if val, ok := cfg.Params["row"].(*int); ok {
        req.row = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *PostWorksheetPivotTableMoveRequest) GetMethod() string {
    return "POST"
}

func (request *PostWorksheetPivotTableMoveRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostWorksheetPivotTableMoveRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/Move"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"pivotTableIndex"+"}", fmt.Sprintf("%v", request.pivotTableIndex), -1)
    return localVarPath
}

func (request *PostWorksheetPivotTableMoveRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.row != nil {
        localVarQueryParams.Add("row", fmt.Sprintf("%v", *request.row))
    }
    if request.column != nil {
        localVarQueryParams.Add("column", fmt.Sprintf("%v", *request.column))
    }
    if request.destCellName != "" {
        localVarQueryParams.Add("destCellName", fmt.Sprintf("%v", request.destCellName))
    }
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostWorksheetPivotTableMoveRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostWorksheetPivotTableMoveRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostWorksheetPivotTableMoveRequest) Description() {
    fmt.Println(strings.Trim("Move PivotTable in the worksheet.", " "))
}