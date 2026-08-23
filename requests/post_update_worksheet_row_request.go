package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PostUpdateWorksheetRowRequest struct {
    name string
    rowIndex int
    sheetName string

    count *int
    folder string
    height *float64
    storageName string
}

func NewPostUpdateWorksheetRowRequest(name string, rowIndex int, sheetName string, opts ...RequestOption) *PostUpdateWorksheetRowRequest {
    req := &PostUpdateWorksheetRowRequest{
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

    if val, ok := cfg.Params["count"].(*int); ok {
        req.count = val
    }
    if val, ok := cfg.Params["folder"].(string); ok {
        req.folder = val
    }
    if val, ok := cfg.Params["height"].(*float64); ok {
        req.height = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *PostUpdateWorksheetRowRequest) GetMethod() string {
    return "POST"
}

func (request *PostUpdateWorksheetRowRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostUpdateWorksheetRowRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/cells/rows/{rowIndex}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"rowIndex"+"}", fmt.Sprintf("%v", request.rowIndex), -1)
    return localVarPath
}

func (request *PostUpdateWorksheetRowRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.height != nil {
        localVarQueryParams.Add("height", fmt.Sprintf("%v", *request.height))
    }
    if request.count != nil {
        localVarQueryParams.Add("count", fmt.Sprintf("%v", *request.count))
    }
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostUpdateWorksheetRowRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostUpdateWorksheetRowRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostUpdateWorksheetRowRequest) Description() {
    fmt.Println(strings.Trim("Update height of rows in the worksheet.", " "))
}