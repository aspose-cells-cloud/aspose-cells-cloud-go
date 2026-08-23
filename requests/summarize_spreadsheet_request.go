package requests

import (
    "fmt"
    "net/url"
    "path/filepath"
    "strings"
)

type SummarizeSpreadsheetRequest struct {
    Spreadsheet string
    SpreadsheetData []byte
    SpreadsheetName string

    password string
    region string
}

func NewSummarizeSpreadsheetRequest(Spreadsheet string, opts ...RequestOption) *SummarizeSpreadsheetRequest {
    req := &SummarizeSpreadsheetRequest{
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

func (request *SummarizeSpreadsheetRequest) SetSpreadsheetBytes(data []byte, name string) {
    if name == "" {
        name = "Spreadsheet"
    }
    request.SpreadsheetData = data
    request.SpreadsheetName = name
}

func (request *SummarizeSpreadsheetRequest) GetMethod() string {
    return "PUT"
}

func (request *SummarizeSpreadsheetRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "multipart/form-data"
    return localVarHeaderParams
}

func (request *SummarizeSpreadsheetRequest) GetPath() string {
    localVarPath := "/cells/ai/summarize/spreadsheet"
    return localVarPath
}

func (request *SummarizeSpreadsheetRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.region != "" {
        localVarQueryParams.Add("region", fmt.Sprintf("%v", request.region))
    }
    if request.password != "" {
        localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
    }
    return localVarQueryParams
}

func (request *SummarizeSpreadsheetRequest) GetJSONBody() interface{} {
    return nil
}

func (request *SummarizeSpreadsheetRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    if request.SpreadsheetData != nil {
        localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
    } else if request.Spreadsheet != "" {
        localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
    }
    return localVarFormParams
}

func (request *SummarizeSpreadsheetRequest) Description() {
    fmt.Println(strings.Trim("Summarizes spreadsheet content using AI and returns the summary as a downloadable text file.", " "))
}