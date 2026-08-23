package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type GetWorksheetColumnRequest struct {
    columnIndex int
    name string
    sheetName string

    folder string
    storageName string
}

func NewGetWorksheetColumnRequest(columnIndex int, name string, sheetName string, opts ...RequestOption) *GetWorksheetColumnRequest {
    req := &GetWorksheetColumnRequest{
        columnIndex: columnIndex,
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

    if val, ok := cfg.Params["folder"].(string); ok {
        req.folder = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *GetWorksheetColumnRequest) GetMethod() string {
    return "GET"
}

func (request *GetWorksheetColumnRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *GetWorksheetColumnRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/cells/columns/{columnIndex}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"columnIndex"+"}", fmt.Sprintf("%v", request.columnIndex), -1)
    return localVarPath
}

func (request *GetWorksheetColumnRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *GetWorksheetColumnRequest) GetJSONBody() interface{} {
    return nil
}

func (request *GetWorksheetColumnRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *GetWorksheetColumnRequest) Description() {
    fmt.Println(strings.Trim("Retrieve worksheet column data by column index.", " "))
}