package requests

import (
    "fmt"
    "net/url"
    "path/filepath"
    "strings"
)

type SearchSpreadsheetContentRequest struct {
    searchText string
    Spreadsheet string
    SpreadsheetData []byte
    SpreadsheetName string

    cellArea string
    ignoringCase *bool
    password string
    region string
    worksheet string
}

func NewSearchSpreadsheetContentRequest(searchText string, Spreadsheet string, opts ...RequestOption) *SearchSpreadsheetContentRequest {
    req := &SearchSpreadsheetContentRequest{
        searchText: searchText,
        Spreadsheet: Spreadsheet,
    }
    if req.searchText == "" {
        return nil
    }

    cfg := &requestConfig{
        Params: make(map[string]interface{}),
    }
    for _, opt := range opts {
        opt.apply(cfg)
    }

    if val, ok := cfg.Params["cellArea"].(string); ok {
        req.cellArea = val
    }
    if val, ok := cfg.Params["ignoringCase"].(*bool); ok {
        req.ignoringCase = val
    }
    if val, ok := cfg.Params["password"].(string); ok {
        req.password = val
    }
    if val, ok := cfg.Params["region"].(string); ok {
        req.region = val
    }
    if val, ok := cfg.Params["worksheet"].(string); ok {
        req.worksheet = val
    }

    return req
}

func (request *SearchSpreadsheetContentRequest) SetSpreadsheetBytes(data []byte, name string) {
    if name == "" {
        name = "Spreadsheet"
    }
    request.SpreadsheetData = data
    request.SpreadsheetName = name
}

func (request *SearchSpreadsheetContentRequest) GetMethod() string {
    return "PUT"
}

func (request *SearchSpreadsheetContentRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "multipart/form-data"
    return localVarHeaderParams
}

func (request *SearchSpreadsheetContentRequest) GetPath() string {
    localVarPath := "/cells/search/content"
    return localVarPath
}

func (request *SearchSpreadsheetContentRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("searchText", fmt.Sprintf("%v", request.searchText))
    if request.ignoringCase != nil {
        localVarQueryParams.Add("ignoringCase", fmt.Sprintf("%v", *request.ignoringCase))
    }
    if request.worksheet != "" {
        localVarQueryParams.Add("worksheet", fmt.Sprintf("%v", request.worksheet))
    }
    if request.cellArea != "" {
        localVarQueryParams.Add("cellArea", fmt.Sprintf("%v", request.cellArea))
    }
    if request.region != "" {
        localVarQueryParams.Add("region", fmt.Sprintf("%v", request.region))
    }
    if request.password != "" {
        localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
    }
    return localVarQueryParams
}

func (request *SearchSpreadsheetContentRequest) GetJSONBody() interface{} {
    return nil
}

func (request *SearchSpreadsheetContentRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    if request.SpreadsheetData != nil {
        localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
    } else if request.Spreadsheet != "" {
        localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
    }
    return localVarFormParams
}

func (request *SearchSpreadsheetContentRequest) Description() {
    fmt.Println(strings.Trim("Search text in the local spreadsheet.", " "))
}