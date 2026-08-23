package requests

import (
    "fmt"
    "net/url"
    "path/filepath"
    "strings"
)

type GetMergedCellsInWorksheetRequest struct {
    Spreadsheet string
    SpreadsheetData []byte
    SpreadsheetName string
    worksheet string

    password string
    region string
}

func NewGetMergedCellsInWorksheetRequest(Spreadsheet string, worksheet string, opts ...RequestOption) *GetMergedCellsInWorksheetRequest {
    req := &GetMergedCellsInWorksheetRequest{
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

    if val, ok := cfg.Params["password"].(string); ok {
        req.password = val
    }
    if val, ok := cfg.Params["region"].(string); ok {
        req.region = val
    }

    return req
}

func (request *GetMergedCellsInWorksheetRequest) SetSpreadsheetBytes(data []byte, name string) {
    if name == "" {
        name = "Spreadsheet"
    }
    request.SpreadsheetData = data
    request.SpreadsheetName = name
}

func (request *GetMergedCellsInWorksheetRequest) GetMethod() string {
    return "PUT"
}

func (request *GetMergedCellsInWorksheetRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "multipart/form-data"
    return localVarHeaderParams
}

func (request *GetMergedCellsInWorksheetRequest) GetPath() string {
    localVarPath := "/cells/spreadsheet/mergedcells"
    return localVarPath
}

func (request *GetMergedCellsInWorksheetRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("worksheet", fmt.Sprintf("%v", request.worksheet))
    if request.region != "" {
        localVarQueryParams.Add("region", fmt.Sprintf("%v", request.region))
    }
    if request.password != "" {
        localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
    }
    return localVarQueryParams
}

func (request *GetMergedCellsInWorksheetRequest) GetJSONBody() interface{} {
    return nil
}

func (request *GetMergedCellsInWorksheetRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    if request.SpreadsheetData != nil {
        localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
    } else if request.Spreadsheet != "" {
        localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
    }
    return localVarFormParams
}

func (request *GetMergedCellsInWorksheetRequest) Description() {
    fmt.Println(strings.Trim("Get all merged cell area form a local spreadsheet worksheet.", " "))
}