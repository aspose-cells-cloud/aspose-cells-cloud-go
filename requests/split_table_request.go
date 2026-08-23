package requests

import (
    "fmt"
    "net/url"
    "path/filepath"
    "strings"
)

type SplitTableRequest struct {
    saveSplitColumn bool
    splitColumnName string
    splitRowNumber int
    Spreadsheet string
    SpreadsheetData []byte
    SpreadsheetName string
    tableName string
    toMultipleFiles bool
    toNewWorkbook bool
    worksheet string

    fontsLocation string
    outPath string
    outStorageName string
    password string
    region string
}

func NewSplitTableRequest(saveSplitColumn bool, splitColumnName string, splitRowNumber int, Spreadsheet string, tableName string, toMultipleFiles bool, toNewWorkbook bool, worksheet string, opts ...RequestOption) *SplitTableRequest {
    req := &SplitTableRequest{
        saveSplitColumn: saveSplitColumn,
        splitColumnName: splitColumnName,
        splitRowNumber: splitRowNumber,
        Spreadsheet: Spreadsheet,
        tableName: tableName,
        toMultipleFiles: toMultipleFiles,
        toNewWorkbook: toNewWorkbook,
        worksheet: worksheet,
    }
    if req.splitColumnName == "" {
        return nil
    }
    if req.tableName == "" {
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

    if val, ok := cfg.Params["fontsLocation"].(string); ok {
        req.fontsLocation = val
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

func (request *SplitTableRequest) SetSpreadsheetBytes(data []byte, name string) {
    if name == "" {
        name = "Spreadsheet"
    }
    request.SpreadsheetData = data
    request.SpreadsheetName = name
}

func (request *SplitTableRequest) GetMethod() string {
    return "PUT"
}

func (request *SplitTableRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "multipart/form-data"
    return localVarHeaderParams
}

func (request *SplitTableRequest) GetPath() string {
    localVarPath := "/cells/split/table"
    return localVarPath
}

func (request *SplitTableRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("worksheet", fmt.Sprintf("%v", request.worksheet))
    localVarQueryParams.Add("tableName", fmt.Sprintf("%v", request.tableName))
    localVarQueryParams.Add("splitColumnName", fmt.Sprintf("%v", request.splitColumnName))
    localVarQueryParams.Add("saveSplitColumn", fmt.Sprintf("%v", request.saveSplitColumn))
    localVarQueryParams.Add("splitRowNumber", fmt.Sprintf("%v", request.splitRowNumber))
    localVarQueryParams.Add("toNewWorkbook", fmt.Sprintf("%v", request.toNewWorkbook))
    localVarQueryParams.Add("toMultipleFiles", fmt.Sprintf("%v", request.toMultipleFiles))
    if request.outPath != "" {
        localVarQueryParams.Add("outPath", fmt.Sprintf("%v", request.outPath))
    }
    if request.outStorageName != "" {
        localVarQueryParams.Add("outStorageName", fmt.Sprintf("%v", request.outStorageName))
    }
    if request.fontsLocation != "" {
        localVarQueryParams.Add("fontsLocation", fmt.Sprintf("%v", request.fontsLocation))
    }
    if request.region != "" {
        localVarQueryParams.Add("region", fmt.Sprintf("%v", request.region))
    }
    if request.password != "" {
        localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
    }
    return localVarQueryParams
}

func (request *SplitTableRequest) GetJSONBody() interface{} {
    return nil
}

func (request *SplitTableRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    if request.SpreadsheetData != nil {
        localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
    } else if request.Spreadsheet != "" {
        localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
    }
    return localVarFormParams
}

func (request *SplitTableRequest) Description() {
    fmt.Println(strings.Trim("Split an Excel worksheet tale into multiple sheets by column value.", " "))
}