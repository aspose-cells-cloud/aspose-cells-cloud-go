package requests

import (
    "fmt"
    "net/url"
    "path/filepath"
    "strings"
)

type UnprotectSpreadsheetRequest struct {
    modifyPassword string
    password string
    Spreadsheet string
    SpreadsheetData []byte
    SpreadsheetName string

    outPath string
    outStorageName string
    region string
}

func NewUnprotectSpreadsheetRequest(modifyPassword string, password string, Spreadsheet string, opts ...RequestOption) *UnprotectSpreadsheetRequest {
    req := &UnprotectSpreadsheetRequest{
        modifyPassword: modifyPassword,
        password: password,
        Spreadsheet: Spreadsheet,
    }
    if req.modifyPassword == "" {
        return nil
    }
    if req.password == "" {
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
    if val, ok := cfg.Params["region"].(string); ok {
        req.region = val
    }

    return req
}

func (request *UnprotectSpreadsheetRequest) SetSpreadsheetBytes(data []byte, name string) {
    if name == "" {
        name = "Spreadsheet"
    }
    request.SpreadsheetData = data
    request.SpreadsheetName = name
}

func (request *UnprotectSpreadsheetRequest) GetMethod() string {
    return "PUT"
}

func (request *UnprotectSpreadsheetRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "multipart/form-data"
    return localVarHeaderParams
}

func (request *UnprotectSpreadsheetRequest) GetPath() string {
    localVarPath := "/cells/unprotection/spreadsheet"
    return localVarPath
}

func (request *UnprotectSpreadsheetRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
    localVarQueryParams.Add("modifyPassword", fmt.Sprintf("%v", request.modifyPassword))
    if request.outPath != "" {
        localVarQueryParams.Add("outPath", fmt.Sprintf("%v", request.outPath))
    }
    if request.outStorageName != "" {
        localVarQueryParams.Add("outStorageName", fmt.Sprintf("%v", request.outStorageName))
    }
    if request.region != "" {
        localVarQueryParams.Add("region", fmt.Sprintf("%v", request.region))
    }
    return localVarQueryParams
}

func (request *UnprotectSpreadsheetRequest) GetJSONBody() interface{} {
    return nil
}

func (request *UnprotectSpreadsheetRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    if request.SpreadsheetData != nil {
        localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
    } else if request.Spreadsheet != "" {
        localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
    }
    return localVarFormParams
}

func (request *UnprotectSpreadsheetRequest) Description() {
    fmt.Println(strings.Trim("Removes dual-layer password protection from Excel spreadsheets, supporting both open and modify passwords with encryption.", " "))
}