package requests

import (
    "fmt"
    "net/url"
    "path/filepath"
    "strings"
)

type AddWorksheetToSpreadsheetRequest struct {
    Spreadsheet string
    SpreadsheetData []byte
    SpreadsheetName string

    outPath string
    outStorageName string
    password string
    position *int
    region string
    sheetName string
    sheetType string

    extraQueryParameters map[string]string
}

func NewAddWorksheetToSpreadsheetRequest(Spreadsheet string, opts ...RequestOption) *AddWorksheetToSpreadsheetRequest {
    req := &AddWorksheetToSpreadsheetRequest{
        Spreadsheet: Spreadsheet,
    }
    cfg := &requestConfig{
        Params: make(map[string]interface{}),
    }
    for _, opt := range opts {
        opt.apply(cfg)
    }

    if val, ok := cfg.Params["outPath"].(string); ok {
        req.outPath = val
    }
    if val, ok := cfg.Params["outStorageName"].(string); ok {
        req.outStorageName = val
    }
    if val, ok := cfg.Params["password"].(string); ok {
        req.password = val
    }
    if val, ok := cfg.Params["position"].(*int); ok {
        req.position = val
    }
    if val, ok := cfg.Params["region"].(string); ok {
        req.region = val
    }
    if val, ok := cfg.Params["sheetName"].(string); ok {
        req.sheetName = val
    }
    if val, ok := cfg.Params["sheetType"].(string); ok {
        req.sheetType = val
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

func (request *AddWorksheetToSpreadsheetRequest) SetSpreadsheetBytes(data []byte, name string) {
    if name == "" {
        name = "Spreadsheet"
    }
    request.SpreadsheetData = data
    request.SpreadsheetName = name
}

func (request *AddWorksheetToSpreadsheetRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *AddWorksheetToSpreadsheetRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
}

func (request *AddWorksheetToSpreadsheetRequest) GetMethod() string {
    return "PUT"
}

func (request *AddWorksheetToSpreadsheetRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "multipart/form-data"
    return localVarHeaderParams
}

func (request *AddWorksheetToSpreadsheetRequest) GetPath() string {
    localVarPath := "/cells/spreadsheet/add/worksheet"
    return localVarPath
}

func (request *AddWorksheetToSpreadsheetRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.sheetType != "" {
        localVarQueryParams.Add("sheetType", fmt.Sprintf("%v", request.sheetType))
    }
    if request.position != nil {
        localVarQueryParams.Add("position", fmt.Sprintf("%v", *request.position))
    }
    if request.sheetName != "" {
        localVarQueryParams.Add("sheetName", fmt.Sprintf("%v", request.sheetName))
    }
    if request.outPath != "" {
        localVarQueryParams.Add("outPath", fmt.Sprintf("%v", request.outPath))
    }
    if request.outStorageName != "" {
        localVarQueryParams.Add("outStorageName", fmt.Sprintf("%v", request.outStorageName))
    }
    if request.region != "" {
        localVarQueryParams.Add("region", fmt.Sprintf("%v", request.region))
    }
    if request.password != "" {
        localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
    }
    for k, v := range request.extraQueryParameters {
        localVarQueryParams.Add(k, v)
    }
    return localVarQueryParams
}

func (request *AddWorksheetToSpreadsheetRequest) GetJSONBody() interface{} {
    return nil
}

func (request *AddWorksheetToSpreadsheetRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    if request.SpreadsheetData != nil {
        localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
    } else if request.Spreadsheet != "" {
        localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
    }
    return localVarFormParams
}

func (request *AddWorksheetToSpreadsheetRequest) Description() {
    fmt.Println(strings.Trim("The Web API enables users to add a new worksheet to a workbook, specifying the worksheet's type, position, and name. This function provides flexibility in managing workbook structure by allowing detailed control over worksheet addition.", " "))
}