package requests

import (
    "fmt"
    "net/url"
    "path/filepath"
    "strings"
)

type UnpivotRangeRequest struct {
    cellArea string
    Spreadsheet string
    SpreadsheetData []byte
    SpreadsheetName string
    worksheet string

    outPath string
    outStorageName string
    password string
    region string
    skipEmptyValue *bool
}

func NewUnpivotRangeRequest(cellArea string, Spreadsheet string, worksheet string, opts ...RequestOption) *UnpivotRangeRequest {
    req := &UnpivotRangeRequest{
        cellArea: cellArea,
        Spreadsheet: Spreadsheet,
        worksheet: worksheet,
    }
    if req.cellArea == "" {
        return nil
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

    return req
}

func (request *UnpivotRangeRequest) SetSpreadsheetBytes(data []byte, name string) {
    if name == "" {
        name = "Spreadsheet"
    }
    request.SpreadsheetData = data
    request.SpreadsheetName = name
}

func (request *UnpivotRangeRequest) GetMethod() string {
    return "PUT"
}

func (request *UnpivotRangeRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "multipart/form-data"
    return localVarHeaderParams
}

func (request *UnpivotRangeRequest) GetPath() string {
    localVarPath := "/cells/unpivot/range"
    return localVarPath
}

func (request *UnpivotRangeRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("worksheet", fmt.Sprintf("%v", request.worksheet))
    localVarQueryParams.Add("cellArea", fmt.Sprintf("%v", request.cellArea))
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
    return localVarQueryParams
}

func (request *UnpivotRangeRequest) GetJSONBody() interface{} {
    return nil
}

func (request *UnpivotRangeRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    if request.SpreadsheetData != nil {
        localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
    } else if request.Spreadsheet != "" {
        localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
    }
    return localVarFormParams
}

func (request *UnpivotRangeRequest) Description() {
    fmt.Println(strings.Trim("Switch rows and columns in the spreadsheet.", " "))
}