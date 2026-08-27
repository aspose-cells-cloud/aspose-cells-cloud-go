package requests

import (
    "fmt"
    "net/url"
    "path/filepath"
    "strings"
)

type ProtectSpreadsheetRequest struct {
    modifyPassword string
    password string
    Spreadsheet string
    SpreadsheetData []byte
    SpreadsheetName string

    outPath string
    outStorageName string
    region string

    extraQueryParameters map[string]string
}

func NewProtectSpreadsheetRequest(modifyPassword string, password string, Spreadsheet string, opts ...RequestOption) *ProtectSpreadsheetRequest {
    req := &ProtectSpreadsheetRequest{
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

func (request *ProtectSpreadsheetRequest) SetSpreadsheetBytes(data []byte, name string) {
    if name == "" {
        name = "Spreadsheet"
    }
    request.SpreadsheetData = data
    request.SpreadsheetName = name
}

func (request *ProtectSpreadsheetRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *ProtectSpreadsheetRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
}

func (request *ProtectSpreadsheetRequest) GetMethod() string {
    return "PUT"
}

func (request *ProtectSpreadsheetRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "multipart/form-data"
    return localVarHeaderParams
}

func (request *ProtectSpreadsheetRequest) GetPath() string {
    localVarPath := "/cells/protection/spreadsheet"
    return localVarPath
}

func (request *ProtectSpreadsheetRequest) GetQueryParameters() url.Values {
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
    for k, v := range request.extraQueryParameters {
        localVarQueryParams.Add(k, v)
    }
    return localVarQueryParams
}

func (request *ProtectSpreadsheetRequest) GetJSONBody() interface{} {
    return nil
}

func (request *ProtectSpreadsheetRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    if request.SpreadsheetData != nil {
        localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
    } else if request.Spreadsheet != "" {
        localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
    }
    return localVarFormParams
}

func (request *ProtectSpreadsheetRequest) Description() {
    fmt.Println(strings.Trim("Applies dual-layer password protection to Excel spreadsheets, supporting both open and modify passwords with encryption.", " "))
}