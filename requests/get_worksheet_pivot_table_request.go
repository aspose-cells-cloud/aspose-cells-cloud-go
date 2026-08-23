package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type GetWorksheetPivotTableRequest struct {
    name string
    pivottableIndex int
    sheetName string

    folder string
    storageName string
}

func NewGetWorksheetPivotTableRequest(name string, pivottableIndex int, sheetName string, opts ...RequestOption) *GetWorksheetPivotTableRequest {
    req := &GetWorksheetPivotTableRequest{
        name: name,
        pivottableIndex: pivottableIndex,
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

func (request *GetWorksheetPivotTableRequest) GetMethod() string {
    return "GET"
}

func (request *GetWorksheetPivotTableRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *GetWorksheetPivotTableRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/pivottables/{pivottableIndex}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"pivottableIndex"+"}", fmt.Sprintf("%v", request.pivottableIndex), -1)
    return localVarPath
}

func (request *GetWorksheetPivotTableRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *GetWorksheetPivotTableRequest) GetJSONBody() interface{} {
    return nil
}

func (request *GetWorksheetPivotTableRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *GetWorksheetPivotTableRequest) Description() {
    fmt.Println(strings.Trim("Retrieve PivotTable information by index in the worksheet.", " "))
}