package requests

import (
    "fmt"
    "net/url"
    "path/filepath"
    "strings"
)

type SearchSpreadsheetAllTextItemsRequest struct {
    Spreadsheet string
    SpreadsheetData []byte
    SpreadsheetName string

    password string
    region string
}

func NewSearchSpreadsheetAllTextItemsRequest(Spreadsheet string, opts ...RequestOption) *SearchSpreadsheetAllTextItemsRequest {
    req := &SearchSpreadsheetAllTextItemsRequest{
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

func (request *SearchSpreadsheetAllTextItemsRequest) SetSpreadsheetBytes(data []byte, name string) {
    if name == "" {
        name = "Spreadsheet"
    }
    request.SpreadsheetData = data
    request.SpreadsheetName = name
}

func (request *SearchSpreadsheetAllTextItemsRequest) GetMethod() string {
    return "PUT"
}

func (request *SearchSpreadsheetAllTextItemsRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "multipart/form-data"
    return localVarHeaderParams
}

func (request *SearchSpreadsheetAllTextItemsRequest) GetPath() string {
    localVarPath := "/cells/search/content/all-textitems"
    return localVarPath
}

func (request *SearchSpreadsheetAllTextItemsRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.region != "" {
        localVarQueryParams.Add("region", fmt.Sprintf("%v", request.region))
    }
    if request.password != "" {
        localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
    }
    return localVarQueryParams
}

func (request *SearchSpreadsheetAllTextItemsRequest) GetJSONBody() interface{} {
    return nil
}

func (request *SearchSpreadsheetAllTextItemsRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    if request.SpreadsheetData != nil {
        localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
    } else if request.Spreadsheet != "" {
        localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
    }
    return localVarFormParams
}

func (request *SearchSpreadsheetAllTextItemsRequest) Description() {
    fmt.Println(strings.Trim("Get all text items in the remote spreadsheet.", " "))
}