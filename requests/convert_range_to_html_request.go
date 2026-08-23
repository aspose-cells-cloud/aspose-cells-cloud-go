package requests

import (
    "fmt"
    "net/url"
    "path/filepath"
    "strings"
)

type ConvertRangeToHtmlRequest struct {
    _range string
    Spreadsheet string
    SpreadsheetData []byte
    SpreadsheetName string
    worksheet string

    AutoColumnsFit *bool
    AutoRowsFit *bool
    fontsLocation string
    outPath string
    outStorageName string
    password string
    region string
}

func NewConvertRangeToHtmlRequest(_range string, Spreadsheet string, worksheet string, opts ...RequestOption) *ConvertRangeToHtmlRequest {
    req := &ConvertRangeToHtmlRequest{
        _range: _range,
        Spreadsheet: Spreadsheet,
        worksheet: worksheet,
    }
    if req._range == "" {
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

    if val, ok := cfg.Params["AutoColumnsFit"].(*bool); ok {
        req.AutoColumnsFit = val
    }
    if val, ok := cfg.Params["AutoRowsFit"].(*bool); ok {
        req.AutoRowsFit = val
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

func (request *ConvertRangeToHtmlRequest) SetSpreadsheetBytes(data []byte, name string) {
    if name == "" {
        name = "Spreadsheet"
    }
    request.SpreadsheetData = data
    request.SpreadsheetName = name
}

func (request *ConvertRangeToHtmlRequest) GetMethod() string {
    return "PUT"
}

func (request *ConvertRangeToHtmlRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "multipart/form-data"
    return localVarHeaderParams
}

func (request *ConvertRangeToHtmlRequest) GetPath() string {
    localVarPath := "/cells/convert/range/html"
    return localVarPath
}

func (request *ConvertRangeToHtmlRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("worksheet", fmt.Sprintf("%v", request.worksheet))
    localVarQueryParams.Add("range", fmt.Sprintf("%v", request._range))
    if request.outPath != "" {
        localVarQueryParams.Add("outPath", fmt.Sprintf("%v", request.outPath))
    }
    if request.outStorageName != "" {
        localVarQueryParams.Add("outStorageName", fmt.Sprintf("%v", request.outStorageName))
    }
    if request.fontsLocation != "" {
        localVarQueryParams.Add("fontsLocation", fmt.Sprintf("%v", request.fontsLocation))
    }
    if request.AutoRowsFit != nil {
        localVarQueryParams.Add("AutoRowsFit", fmt.Sprintf("%v", *request.AutoRowsFit))
    }
    if request.AutoColumnsFit != nil {
        localVarQueryParams.Add("AutoColumnsFit", fmt.Sprintf("%v", *request.AutoColumnsFit))
    }
    if request.region != "" {
        localVarQueryParams.Add("region", fmt.Sprintf("%v", request.region))
    }
    if request.password != "" {
        localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
    }
    return localVarQueryParams
}

func (request *ConvertRangeToHtmlRequest) GetJSONBody() interface{} {
    return nil
}

func (request *ConvertRangeToHtmlRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    if request.SpreadsheetData != nil {
        localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
    } else if request.Spreadsheet != "" {
        localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
    }
    return localVarFormParams
}

func (request *ConvertRangeToHtmlRequest) Description() {
    fmt.Println(strings.Trim("Converts a range of spreadsheet on a local drive to the html file.", " "))
}