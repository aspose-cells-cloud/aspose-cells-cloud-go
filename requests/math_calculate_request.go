package requests

import (
    "fmt"
    "net/url"
    "path/filepath"
    "strings"
)

type MathCalculateRequest struct {
    operation string
    Spreadsheet string
    SpreadsheetData []byte
    SpreadsheetName string
    value string

    password string
    _range string
    region string
    worksheet string
}

func NewMathCalculateRequest(operation string, Spreadsheet string, value string, opts ...RequestOption) *MathCalculateRequest {
    req := &MathCalculateRequest{
        operation: operation,
        Spreadsheet: Spreadsheet,
        value: value,
    }
    if req.operation == "" {
        return nil
    }
    if req.value == "" {
        return nil
    }

    cfg := &requestConfig{
        Params: make(map[string]interface{}),
    }
    for _, opt := range opts {
        opt.apply(cfg)
    }

    if val, ok := cfg.Params["password"].(string); ok {
        req.password = val
    }
    if val, ok := cfg.Params["range"].(string); ok {
        req._range = val
    }
    if val, ok := cfg.Params["region"].(string); ok {
        req.region = val
    }
    if val, ok := cfg.Params["worksheet"].(string); ok {
        req.worksheet = val
    }

    return req
}

func (request *MathCalculateRequest) SetSpreadsheetBytes(data []byte, name string) {
    if name == "" {
        name = "Spreadsheet"
    }
    request.SpreadsheetData = data
    request.SpreadsheetName = name
}

func (request *MathCalculateRequest) GetMethod() string {
    return "PUT"
}

func (request *MathCalculateRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "multipart/form-data"
    return localVarHeaderParams
}

func (request *MathCalculateRequest) GetPath() string {
    localVarPath := "/cells/calculate/math"
    return localVarPath
}

func (request *MathCalculateRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("operation", fmt.Sprintf("%v", request.operation))
    localVarQueryParams.Add("value", fmt.Sprintf("%v", request.value))
    if request.worksheet != "" {
        localVarQueryParams.Add("worksheet", fmt.Sprintf("%v", request.worksheet))
    }
    if request._range != "" {
        localVarQueryParams.Add("range", fmt.Sprintf("%v", request._range))
    }
    if request.region != "" {
        localVarQueryParams.Add("region", fmt.Sprintf("%v", request.region))
    }
    if request.password != "" {
        localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
    }
    return localVarQueryParams
}

func (request *MathCalculateRequest) GetJSONBody() interface{} {
    return nil
}

func (request *MathCalculateRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    if request.SpreadsheetData != nil {
        localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
    } else if request.Spreadsheet != "" {
        localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
    }
    return localVarFormParams
}

func (request *MathCalculateRequest) Description() {
    fmt.Println(strings.Trim("MathCalculate", " "))
}