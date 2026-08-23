package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PostHideWorksheetColumnsRequest struct {
    name string
    sheetName string
    startColumn int
    totalColumns int

    folder string
    storageName string
}

func NewPostHideWorksheetColumnsRequest(name string, sheetName string, startColumn int, totalColumns int, opts ...RequestOption) *PostHideWorksheetColumnsRequest {
    req := &PostHideWorksheetColumnsRequest{
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

    return req
}

func (request *PostHideWorksheetColumnsRequest) GetMethod() string {
    return "POST"
}

func (request *PostHideWorksheetColumnsRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostHideWorksheetColumnsRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/cells/columns/hide"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PostHideWorksheetColumnsRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("startColumn", fmt.Sprintf("%v", request.startColumn))
    localVarQueryParams.Add("totalColumns", fmt.Sprintf("%v", request.totalColumns))
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostHideWorksheetColumnsRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostHideWorksheetColumnsRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostHideWorksheetColumnsRequest) Description() {
    fmt.Println(strings.Trim("Hide worksheet columns in the worksheet.", " "))
}