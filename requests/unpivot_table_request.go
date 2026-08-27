package requests

import (
    "fmt"
    "net/url"
    "path/filepath"
    "strings"
)

type UnpivotTableRequest struct {
    index int
    Spreadsheet string
    SpreadsheetData []byte
    SpreadsheetName string
    worksheet string

    outPath string
    outStorageName string
    password string
    region string
    skipEmptyValue *bool

    extraQueryParameters map[string]string
}

func NewUnpivotTableRequest(index int, Spreadsheet string, worksheet string, opts ...RequestOption) *UnpivotTableRequest {
    req := &UnpivotTableRequest{
        index: index,
        Spreadsheet: Spreadsheet,
        worksheet: worksheet,
    }
    if req.worksheet == "" {
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
    if val, ok := cfg.Params["skipEmptyValue"].(*bool); ok {
        req.skipEmptyValue = val
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

func (request *UnpivotTableRequest) SetSpreadsheetBytes(data []byte, name string) {
    if name == "" {
        name = "Spreadsheet"
    }
    request.SpreadsheetData = data
    request.SpreadsheetName = name
}

func (request *UnpivotTableRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *UnpivotTableRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
}

func (request *UnpivotTableRequest) GetMethod() string {
    return "PUT"
}

func (request *UnpivotTableRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "multipart/form-data"
    return localVarHeaderParams
}

func (request *UnpivotTableRequest) GetPath() string {
    localVarPath := "/cells/unpivot/table"
    return localVarPath
}

func (request *UnpivotTableRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("worksheet", fmt.Sprintf("%v", request.worksheet))
    localVarQueryParams.Add("index", fmt.Sprintf("%v", request.index))
    if request.skipEmptyValue != nil {
        localVarQueryParams.Add("skipEmptyValue", fmt.Sprintf("%v", *request.skipEmptyValue))
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

func (request *UnpivotTableRequest) GetJSONBody() interface{} {
    return nil
}

func (request *UnpivotTableRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    if request.SpreadsheetData != nil {
        localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
    } else if request.Spreadsheet != "" {
        localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
    }
    return localVarFormParams
}

func (request *UnpivotTableRequest) Description() {
    fmt.Println(strings.Trim("Switch rows and columns in the spreadsheet.", " "))
}