package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PutWorksheetPivotTableRequest struct {
    name string
    sheetName string

    destCellName string
    folder string
    sourceData string
    storageName string
    tableName string
    useSameSource *bool
}

func NewPutWorksheetPivotTableRequest(name string, sheetName string, opts ...RequestOption) *PutWorksheetPivotTableRequest {
    req := &PutWorksheetPivotTableRequest{
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

    if val, ok := cfg.Params["destCellName"].(string); ok {
        req.destCellName = val
    }
    if val, ok := cfg.Params["folder"].(string); ok {
        req.folder = val
    }
    if val, ok := cfg.Params["sourceData"].(string); ok {
        req.sourceData = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }
    if val, ok := cfg.Params["tableName"].(string); ok {
        req.tableName = val
    }
    if val, ok := cfg.Params["useSameSource"].(*bool); ok {
        req.useSameSource = val
    }

    return req
}

func (request *PutWorksheetPivotTableRequest) GetMethod() string {
    return "PUT"
}

func (request *PutWorksheetPivotTableRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PutWorksheetPivotTableRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/pivottables"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PutWorksheetPivotTableRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.sourceData != "" {
        localVarQueryParams.Add("sourceData", fmt.Sprintf("%v", request.sourceData))
    }
    if request.destCellName != "" {
        localVarQueryParams.Add("destCellName", fmt.Sprintf("%v", request.destCellName))
    }
    if request.tableName != "" {
        localVarQueryParams.Add("tableName", fmt.Sprintf("%v", request.tableName))
    }
    if request.useSameSource != nil {
        localVarQueryParams.Add("useSameSource", fmt.Sprintf("%v", *request.useSameSource))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PutWorksheetPivotTableRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PutWorksheetPivotTableRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PutWorksheetPivotTableRequest) Description() {
    fmt.Println(strings.Trim("Add a PivotTable in the worksheet.", " "))
}