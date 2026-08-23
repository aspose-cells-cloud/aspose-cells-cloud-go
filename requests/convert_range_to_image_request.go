package requests

import (
    "fmt"
    "net/url"
    "path/filepath"
    "strings"
)

type ConvertRangeToImageRequest struct {
    format string
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
    printHeadings *bool
    region string
}

func NewConvertRangeToImageRequest(format string, _range string, Spreadsheet string, worksheet string, opts ...RequestOption) *ConvertRangeToImageRequest {
    req := &ConvertRangeToImageRequest{
        format: format,
        _range: _range,
        Spreadsheet: Spreadsheet,
        worksheet: worksheet,
    }
    if req.format == "" {
        return nil
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
    if val, ok := cfg.Params["printHeadings"].(*bool); ok {
        req.printHeadings = val
    }
    if val, ok := cfg.Params["region"].(string); ok {
        req.region = val
    }

    return req
}

func (request *ConvertRangeToImageRequest) SetSpreadsheetBytes(data []byte, name string) {
    if name == "" {
        name = "Spreadsheet"
    }
    request.SpreadsheetData = data
    request.SpreadsheetName = name
}

func (request *ConvertRangeToImageRequest) GetMethod() string {
    return "PUT"
}

func (request *ConvertRangeToImageRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "multipart/form-data"
    return localVarHeaderParams
}

func (request *ConvertRangeToImageRequest) GetPath() string {
    localVarPath := "/cells/convert/range/image"
    return localVarPath
}

func (request *ConvertRangeToImageRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("worksheet", fmt.Sprintf("%v", request.worksheet))
    localVarQueryParams.Add("range", fmt.Sprintf("%v", request._range))
    localVarQueryParams.Add("format", fmt.Sprintf("%v", request.format))
    if request.printHeadings != nil {
        localVarQueryParams.Add("printHeadings", fmt.Sprintf("%v", *request.printHeadings))
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

func (request *ConvertRangeToImageRequest) GetJSONBody() interface{} {
    return nil
}

func (request *ConvertRangeToImageRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    if request.SpreadsheetData != nil {
        localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
    } else if request.Spreadsheet != "" {
        localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
    }
    return localVarFormParams
}

func (request *ConvertRangeToImageRequest) Description() {
    fmt.Println(strings.Trim("Converts a range of spreadsheet on a local drive to the image file.", " "))
}