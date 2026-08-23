package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PostAutofitWorksheetRowRequest struct {
    name string
    rowIndex int
    sheetName string

    firstColumn *int
    folder string
    lastColumn *int
    rowCount *int
    storageName string
}

func NewPostAutofitWorksheetRowRequest(name string, rowIndex int, sheetName string, opts ...RequestOption) *PostAutofitWorksheetRowRequest {
    req := &PostAutofitWorksheetRowRequest{
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

    if val, ok := cfg.Params["firstColumn"].(*int); ok {
        req.firstColumn = val
    }
    if val, ok := cfg.Params["folder"].(string); ok {
        req.folder = val
    }
    if val, ok := cfg.Params["lastColumn"].(*int); ok {
        req.lastColumn = val
    }
    if val, ok := cfg.Params["rowCount"].(*int); ok {
        req.rowCount = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *PostAutofitWorksheetRowRequest) GetMethod() string {
    return "POST"
}

func (request *PostAutofitWorksheetRowRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostAutofitWorksheetRowRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/autofitrow"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PostAutofitWorksheetRowRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("rowIndex", fmt.Sprintf("%v", request.rowIndex))
    if request.firstColumn != nil {
        localVarQueryParams.Add("firstColumn", fmt.Sprintf("%v", *request.firstColumn))
    }
    if request.lastColumn != nil {
        localVarQueryParams.Add("lastColumn", fmt.Sprintf("%v", *request.lastColumn))
    }
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    if request.rowCount != nil {
        localVarQueryParams.Add("rowCount", fmt.Sprintf("%v", *request.rowCount))
    }
    return localVarQueryParams
}

func (request *PostAutofitWorksheetRowRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostAutofitWorksheetRowRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostAutofitWorksheetRowRequest) Description() {
    fmt.Println(strings.Trim("Autofit a row in the worksheet.", " "))
}