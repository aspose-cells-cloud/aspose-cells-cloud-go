package requests

import (
    "fmt"
    "net/url"
    "path/filepath"
    "strings"
)

type RenameWorksheetInSpreadsheetRequest struct {
    sourceName string
    Spreadsheet string
    SpreadsheetData []byte
    SpreadsheetName string
    targetName string

    outPath string
    outStorageName string
    password string
    region string

    extraQueryParameters map[string]string
}

func NewRenameWorksheetInSpreadsheetRequest(sourceName string, Spreadsheet string, targetName string, opts ...RequestOption) *RenameWorksheetInSpreadsheetRequest {
    req := &RenameWorksheetInSpreadsheetRequest{
        sourceName: sourceName,
        Spreadsheet: Spreadsheet,
        targetName: targetName,
    }
    if req.sourceName == "" {
        return nil
    }
    if req.targetName == "" {
        return nil
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
    if val, ok := cfg.Params["region"].(string); ok {
        req.region = val
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

func (request *RenameWorksheetInSpreadsheetRequest) SetSpreadsheetBytes(data []byte, name string) {
    if name == "" {
        name = "Spreadsheet"
    }
    request.SpreadsheetData = data
    request.SpreadsheetName = name
}

func (request *RenameWorksheetInSpreadsheetRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *RenameWorksheetInSpreadsheetRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
}

func (request *RenameWorksheetInSpreadsheetRequest) GetMethod() string {
    return "PUT"
}

func (request *RenameWorksheetInSpreadsheetRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "multipart/form-data"
    return localVarHeaderParams
}

func (request *RenameWorksheetInSpreadsheetRequest) GetPath() string {
    localVarPath := "/cells/spreadsheet/rename/worksheet"
    return localVarPath
}

func (request *RenameWorksheetInSpreadsheetRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("sourceName", fmt.Sprintf("%v", request.sourceName))
    localVarQueryParams.Add("targetName", fmt.Sprintf("%v", request.targetName))
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

func (request *RenameWorksheetInSpreadsheetRequest) GetJSONBody() interface{} {
    return nil
}

func (request *RenameWorksheetInSpreadsheetRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    if request.SpreadsheetData != nil {
        localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
    } else if request.Spreadsheet != "" {
        localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
    }
    return localVarFormParams
}

func (request *RenameWorksheetInSpreadsheetRequest) Description() {
    fmt.Println(strings.Trim("The Web API endpoint allows users to rename a specified worksheet within a workbook. This function provides a straightforward way to update worksheet names, enhancing workbook organization and readability.", " "))
}