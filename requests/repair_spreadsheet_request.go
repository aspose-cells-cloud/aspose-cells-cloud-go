package requests

import (
    "fmt"
    "net/url"
    "path/filepath"
    "strings"
)

type RepairSpreadsheetRequest struct {
    Spreadsheet string
    SpreadsheetData []byte
    SpreadsheetName string

    outPath string
    outStorageName string
    password string
    region string
}

func NewRepairSpreadsheetRequest(Spreadsheet string, opts ...RequestOption) *RepairSpreadsheetRequest {
    req := &RepairSpreadsheetRequest{
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
    if val, ok := cfg.Params["region"].(string); ok {
        req.region = val
    }

    return req
}

func (request *RepairSpreadsheetRequest) SetSpreadsheetBytes(data []byte, name string) {
    if name == "" {
        name = "Spreadsheet"
    }
    request.SpreadsheetData = data
    request.SpreadsheetName = name
}

func (request *RepairSpreadsheetRequest) GetMethod() string {
    return "PUT"
}

func (request *RepairSpreadsheetRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "multipart/form-data"
    return localVarHeaderParams
}

func (request *RepairSpreadsheetRequest) GetPath() string {
    localVarPath := "/cells/spreadsheet/repair"
    return localVarPath
}

func (request *RepairSpreadsheetRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
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
    return localVarQueryParams
}

func (request *RepairSpreadsheetRequest) GetJSONBody() interface{} {
    return nil
}

func (request *RepairSpreadsheetRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    if request.SpreadsheetData != nil {
        localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
    } else if request.Spreadsheet != "" {
        localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
    }
    return localVarFormParams
}

func (request *RepairSpreadsheetRequest) Description() {
    fmt.Println(strings.Trim("The Web API endpoint allows users to repair a spreadsheet.", " "))
}