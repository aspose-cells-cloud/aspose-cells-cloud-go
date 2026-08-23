package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PostHideWorksheetRowsRequest struct {
    name string
    sheetName string
    startrow int
    totalRows int

    folder string
    storageName string
}

func NewPostHideWorksheetRowsRequest(name string, sheetName string, startrow int, totalRows int, opts ...RequestOption) *PostHideWorksheetRowsRequest {
    req := &PostHideWorksheetRowsRequest{
        name: name,
        sheetName: sheetName,
        startrow: startrow,
        totalRows: totalRows,
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

func (request *PostHideWorksheetRowsRequest) GetMethod() string {
    return "POST"
}

func (request *PostHideWorksheetRowsRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostHideWorksheetRowsRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/cells/rows/hide"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PostHideWorksheetRowsRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("startrow", fmt.Sprintf("%v", request.startrow))
    localVarQueryParams.Add("totalRows", fmt.Sprintf("%v", request.totalRows))
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostHideWorksheetRowsRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostHideWorksheetRowsRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostHideWorksheetRowsRequest) Description() {
    fmt.Println(strings.Trim("Hide rows in worksheet.", " "))
}