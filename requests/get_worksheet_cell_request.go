package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type GetWorksheetCellRequest struct {
    cellOrMethodName string
    name string
    sheetName string

    folder string
    storageName string
}

func NewGetWorksheetCellRequest(cellOrMethodName string, name string, sheetName string, opts ...RequestOption) *GetWorksheetCellRequest {
    req := &GetWorksheetCellRequest{
        cellOrMethodName: cellOrMethodName,
        name: name,
        sheetName: sheetName,
    }
    if req.cellOrMethodName == "" {
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

func (request *GetWorksheetCellRequest) GetMethod() string {
    return "GET"
}

func (request *GetWorksheetCellRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *GetWorksheetCellRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/cells/{cellOrMethodName}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"cellOrMethodName"+"}", fmt.Sprintf("%v", request.cellOrMethodName), -1)
    return localVarPath
}

func (request *GetWorksheetCellRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *GetWorksheetCellRequest) GetJSONBody() interface{} {
    return nil
}

func (request *GetWorksheetCellRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *GetWorksheetCellRequest) Description() {
    fmt.Println(strings.Trim("Retrieve cell data using either cell reference or method name in the worksheet.", " "))
}