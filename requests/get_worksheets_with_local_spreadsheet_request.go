package requests

import (
    "fmt"
    "net/url"
    "path/filepath"
    "strings"
)

type GetWorksheetsWithLocalSpreadsheetRequest struct {
    Spreadsheet string
    SpreadsheetData []byte
    SpreadsheetName string

    password string
    region string
}

func NewGetWorksheetsWithLocalSpreadsheetRequest(Spreadsheet string, opts ...RequestOption) *GetWorksheetsWithLocalSpreadsheetRequest {
    req := &GetWorksheetsWithLocalSpreadsheetRequest{
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

func (request *GetWorksheetsWithLocalSpreadsheetRequest) SetSpreadsheetBytes(data []byte, name string) {
    if name == "" {
        name = "Spreadsheet"
    }
    request.SpreadsheetData = data
    request.SpreadsheetName = name
}

func (request *GetWorksheetsWithLocalSpreadsheetRequest) GetMethod() string {
    return "PUT"
}

func (request *GetWorksheetsWithLocalSpreadsheetRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "multipart/form-data"
    return localVarHeaderParams
}

func (request *GetWorksheetsWithLocalSpreadsheetRequest) GetPath() string {
    localVarPath := "/cells/spreadsheet/worksheets"
    return localVarPath
}

func (request *GetWorksheetsWithLocalSpreadsheetRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.region != "" {
        localVarQueryParams.Add("region", fmt.Sprintf("%v", request.region))
    }
    if request.password != "" {
        localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
    }
    return localVarQueryParams
}

func (request *GetWorksheetsWithLocalSpreadsheetRequest) GetJSONBody() interface{} {
    return nil
}

func (request *GetWorksheetsWithLocalSpreadsheetRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    if request.SpreadsheetData != nil {
        localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
    } else if request.Spreadsheet != "" {
        localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
    }
    return localVarFormParams
}

func (request *GetWorksheetsWithLocalSpreadsheetRequest) Description() {
    fmt.Println(strings.Trim("Fetches a complete list of worksheets from the currently active local spreadsheet.", " "))
}