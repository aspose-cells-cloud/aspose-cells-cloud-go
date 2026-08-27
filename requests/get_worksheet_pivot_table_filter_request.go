package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type GetWorksheetPivotTableFilterRequest struct {
    filterIndex int
    name string
    pivotTableIndex int
    sheetName string

    folder string
    storageName string

    extraQueryParameters map[string]string
}

func NewGetWorksheetPivotTableFilterRequest(filterIndex int, name string, pivotTableIndex int, sheetName string, opts ...RequestOption) *GetWorksheetPivotTableFilterRequest {
    req := &GetWorksheetPivotTableFilterRequest{
        filterIndex: filterIndex,
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
    if len(cfg.extraQueryParams) > 0 {
        if req.extraQueryParameters == nil {
            req.extraQueryParameters = make(map[string]string)
        }
        for k, v := range cfg.extraQueryParams {
            req.extraQueryParameters[k] = v
        }
    }

    return req
}

func (request *GetWorksheetPivotTableFilterRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *GetWorksheetPivotTableFilterRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
}

func (request *GetWorksheetPivotTableFilterRequest) GetMethod() string {
    return "GET"
}

func (request *GetWorksheetPivotTableFilterRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *GetWorksheetPivotTableFilterRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/PivotFilters/{filterIndex}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"pivotTableIndex"+"}", fmt.Sprintf("%v", request.pivotTableIndex), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"filterIndex"+"}", fmt.Sprintf("%v", request.filterIndex), -1)
    return localVarPath
}

func (request *GetWorksheetPivotTableFilterRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    for k, v := range request.extraQueryParameters {
        localVarQueryParams.Add(k, v)
    }
    return localVarQueryParams
}

func (request *GetWorksheetPivotTableFilterRequest) GetJSONBody() interface{} {
    return nil
}

func (request *GetWorksheetPivotTableFilterRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *GetWorksheetPivotTableFilterRequest) Description() {
    fmt.Println(strings.Trim("Retrieve PivotTable filters in the worksheet.", " "))
}