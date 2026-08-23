package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type GetWorksheetRowRequest struct {
    name string
    rowIndex int
    sheetName string

    folder string
    storageName string
}

func NewGetWorksheetRowRequest(name string, rowIndex int, sheetName string, opts ...RequestOption) *GetWorksheetRowRequest {
    req := &GetWorksheetRowRequest{
        name: name,
        rowIndex: rowIndex,
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

func (request *GetWorksheetRowRequest) GetMethod() string {
    return "GET"
}

func (request *GetWorksheetRowRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *GetWorksheetRowRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/cells/rows/{rowIndex}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"rowIndex"+"}", fmt.Sprintf("%v", request.rowIndex), -1)
    return localVarPath
}

func (request *GetWorksheetRowRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *GetWorksheetRowRequest) GetJSONBody() interface{} {
    return nil
}

func (request *GetWorksheetRowRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *GetWorksheetRowRequest) Description() {
    fmt.Println(strings.Trim("Retrieve row data by the row's index in the worksheet.", " "))
}