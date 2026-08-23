package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PostUnhideWorksheetColumnsRequest struct {
    name string
    sheetName string
    startColumn int
    totalColumns int

    folder string
    storageName string
    width *float64
}

func NewPostUnhideWorksheetColumnsRequest(name string, sheetName string, startColumn int, totalColumns int, opts ...RequestOption) *PostUnhideWorksheetColumnsRequest {
    req := &PostUnhideWorksheetColumnsRequest{
        name: name,
        sheetName: sheetName,
        startColumn: startColumn,
        totalColumns: totalColumns,
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
    if val, ok := cfg.Params["width"].(*float64); ok {
        req.width = val
    }

    return req
}

func (request *PostUnhideWorksheetColumnsRequest) GetMethod() string {
    return "POST"
}

func (request *PostUnhideWorksheetColumnsRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostUnhideWorksheetColumnsRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/cells/columns/unhide"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PostUnhideWorksheetColumnsRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("startColumn", fmt.Sprintf("%v", request.startColumn))
    localVarQueryParams.Add("totalColumns", fmt.Sprintf("%v", request.totalColumns))
    if request.width != nil {
        localVarQueryParams.Add("width", fmt.Sprintf("%v", *request.width))
    }
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostUnhideWorksheetColumnsRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostUnhideWorksheetColumnsRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostUnhideWorksheetColumnsRequest) Description() {
    fmt.Println(strings.Trim("Unhide worksheet columns in the worksheet.", " "))
}