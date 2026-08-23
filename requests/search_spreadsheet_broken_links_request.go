package requests

import (
    "fmt"
    "net/url"
    "path/filepath"
    "strings"
)

type SearchSpreadsheetBrokenLinksRequest struct {
    Spreadsheet string
    SpreadsheetData []byte
    SpreadsheetName string

    cellArea string
    password string
    region string
    worksheet string
}

func NewSearchSpreadsheetBrokenLinksRequest(Spreadsheet string, opts ...RequestOption) *SearchSpreadsheetBrokenLinksRequest {
    req := &SearchSpreadsheetBrokenLinksRequest{
        Spreadsheet: Spreadsheet,
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

func (request *SearchSpreadsheetBrokenLinksRequest) SetSpreadsheetBytes(data []byte, name string) {
    if name == "" {
        name = "Spreadsheet"
    }
    request.SpreadsheetData = data
    request.SpreadsheetName = name
}

func (request *SearchSpreadsheetBrokenLinksRequest) GetMethod() string {
    return "PUT"
}

func (request *SearchSpreadsheetBrokenLinksRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "multipart/form-data"
    return localVarHeaderParams
}

func (request *SearchSpreadsheetBrokenLinksRequest) GetPath() string {
    localVarPath := "/cells/search/broken-links"
    return localVarPath
}

func (request *SearchSpreadsheetBrokenLinksRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
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

func (request *SearchSpreadsheetBrokenLinksRequest) GetJSONBody() interface{} {
    return nil
}

func (request *SearchSpreadsheetBrokenLinksRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    if request.SpreadsheetData != nil {
        localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
    } else if request.Spreadsheet != "" {
        localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
    }
    return localVarFormParams
}

func (request *SearchSpreadsheetBrokenLinksRequest) Description() {
    fmt.Println(strings.Trim("Search broken links in the local spreadsheet.", " "))
}