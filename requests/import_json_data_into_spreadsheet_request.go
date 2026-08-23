package requests

import (
    "fmt"
    "net/url"
    "path/filepath"
    "strings"
)

type ImportJSONDataIntoSpreadsheetRequest struct {
    datafile string
    datafileData []byte
    datafileName string
    Spreadsheet string
    SpreadsheetData []byte
    SpreadsheetName string
    startcell string
    worksheet string

    fontsLocation string
    insert *bool
    outPath string
    outStorageName string
    password string
    region string
}

func NewImportJSONDataIntoSpreadsheetRequest(datafile string, Spreadsheet string, startcell string, worksheet string, opts ...RequestOption) *ImportJSONDataIntoSpreadsheetRequest {
    req := &ImportJSONDataIntoSpreadsheetRequest{
        datafile: datafile,
        Spreadsheet: Spreadsheet,
        startcell: startcell,
        worksheet: worksheet,
    }
    if req.startcell == "" {
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
    if val, ok := cfg.Params["insert"].(*bool); ok {
        req.insert = val
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

func (request *ImportJSONDataIntoSpreadsheetRequest) SetDatafileBytes(data []byte, name string) {
    if name == "" {
        name = "datafile"
    }
    request.datafileData = data
    request.datafileName = name
}

func (request *ImportJSONDataIntoSpreadsheetRequest) SetSpreadsheetBytes(data []byte, name string) {
    if name == "" {
        name = "Spreadsheet"
    }
    request.SpreadsheetData = data
    request.SpreadsheetName = name
}

func (request *ImportJSONDataIntoSpreadsheetRequest) GetMethod() string {
    return "PUT"
}

func (request *ImportJSONDataIntoSpreadsheetRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "multipart/form-data"
    return localVarHeaderParams
}

func (request *ImportJSONDataIntoSpreadsheetRequest) GetPath() string {
    localVarPath := "/cells/import/data/json"
    return localVarPath
}

func (request *ImportJSONDataIntoSpreadsheetRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("worksheet", fmt.Sprintf("%v", request.worksheet))
    localVarQueryParams.Add("startcell", fmt.Sprintf("%v", request.startcell))
    if request.insert != nil {
        localVarQueryParams.Add("insert", fmt.Sprintf("%v", *request.insert))
    }
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

func (request *ImportJSONDataIntoSpreadsheetRequest) GetJSONBody() interface{} {
    return nil
}

func (request *ImportJSONDataIntoSpreadsheetRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    if request.datafileData != nil {
        localVarFormParams[request.datafileName] = request.datafileData
    } else if request.datafile != "" {
        localVarFormParams["@"+filepath.Base(request.datafile)] = request.datafile
    }
    if request.SpreadsheetData != nil {
        localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
    } else if request.Spreadsheet != "" {
        localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
    }
    return localVarFormParams
}

func (request *ImportJSONDataIntoSpreadsheetRequest) Description() {
    fmt.Println(strings.Trim("Import JSON data file into the local spreadsheet.", " "))
}