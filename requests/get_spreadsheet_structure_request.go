package requests

import (
    "fmt"
    "net/url"
    "path/filepath"
    "strings"
)

type GetSpreadsheetStructureRequest struct {
    Spreadsheet string
    SpreadsheetData []byte
    SpreadsheetName string

    password string
    region string
}

func NewGetSpreadsheetStructureRequest(Spreadsheet string, opts ...RequestOption) *GetSpreadsheetStructureRequest {
    req := &GetSpreadsheetStructureRequest{
        Spreadsheet: Spreadsheet,
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
    if val, ok := cfg.Params["region"].(string); ok {
        req.region = val
    }

    return req
}

func (request *GetSpreadsheetStructureRequest) SetSpreadsheetBytes(data []byte, name string) {
    if name == "" {
        name = "Spreadsheet"
    }
    request.SpreadsheetData = data
    request.SpreadsheetName = name
}

func (request *GetSpreadsheetStructureRequest) GetMethod() string {
    return "PUT"
}

func (request *GetSpreadsheetStructureRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "multipart/form-data"
    return localVarHeaderParams
}

func (request *GetSpreadsheetStructureRequest) GetPath() string {
    localVarPath := "/cells/spreadsheet/structure"
    return localVarPath
}

func (request *GetSpreadsheetStructureRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.region != "" {
        localVarQueryParams.Add("region", fmt.Sprintf("%v", request.region))
    }
    if request.password != "" {
        localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
    }
    return localVarQueryParams
}

func (request *GetSpreadsheetStructureRequest) GetJSONBody() interface{} {
    return nil
}

func (request *GetSpreadsheetStructureRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    if request.SpreadsheetData != nil {
        localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
    } else if request.Spreadsheet != "" {
        localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
    }
    return localVarFormParams
}

func (request *GetSpreadsheetStructureRequest) Description() {
    fmt.Println(strings.Trim("Structurally convert the core metadata, worksheets, tables, pivot tables, charts, shapes, and other information of an Excel workbook into a JObject type JSON object, for scenarios such as data export, API responses, and log recording.", " "))
}