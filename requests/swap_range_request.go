package requests

import (
    "fmt"
    "net/url"
    "path/filepath"
    "strings"
)

type SwapRangeRequest struct {
    range1 string
    range2 string
    Spreadsheet string
    SpreadsheetData []byte
    SpreadsheetName string
    worksheet1 string
    worksheet2 string

    outPath string
    outStorageName string
    password string
    region string
}

func NewSwapRangeRequest(range1 string, range2 string, Spreadsheet string, worksheet1 string, worksheet2 string, opts ...RequestOption) *SwapRangeRequest {
    req := &SwapRangeRequest{
        range1: range1,
        range2: range2,
        Spreadsheet: Spreadsheet,
        worksheet1: worksheet1,
        worksheet2: worksheet2,
    }
    if req.range1 == "" {
        return nil
    }
    if req.range2 == "" {
        return nil
    }
    if req.worksheet1 == "" {
        return nil
    }
    if req.worksheet2 == "" {
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

    return req
}

func (request *SwapRangeRequest) SetSpreadsheetBytes(data []byte, name string) {
    if name == "" {
        name = "Spreadsheet"
    }
    request.SpreadsheetData = data
    request.SpreadsheetName = name
}

func (request *SwapRangeRequest) GetMethod() string {
    return "PUT"
}

func (request *SwapRangeRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "multipart/form-data"
    return localVarHeaderParams
}

func (request *SwapRangeRequest) GetPath() string {
    localVarPath := "/cells/swap/range"
    return localVarPath
}

func (request *SwapRangeRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("worksheet1", fmt.Sprintf("%v", request.worksheet1))
    localVarQueryParams.Add("range1", fmt.Sprintf("%v", request.range1))
    localVarQueryParams.Add("worksheet2", fmt.Sprintf("%v", request.worksheet2))
    localVarQueryParams.Add("range2", fmt.Sprintf("%v", request.range2))
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

func (request *SwapRangeRequest) GetJSONBody() interface{} {
    return nil
}

func (request *SwapRangeRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    if request.SpreadsheetData != nil {
        localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
    } else if request.Spreadsheet != "" {
        localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
    }
    return localVarFormParams
}

func (request *SwapRangeRequest) Description() {
    fmt.Println(strings.Trim("The Swap Ranges for Excel API provides a powerful tool to move any two columns, rows, ranges, or individual cells within an Excel file. This API allows users to re-arrange their tables quickly and efficiently, ensuring that the original data formatting is preserved and all existing formulas continue to function correctly. By leveraging this API, users can streamline their data manipulation tasks and maintain the integrity of their spreadsheets.", " "))
}