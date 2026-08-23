package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type ExportSpreadsheetAsFormatRequest struct {
    format string
    name string

    AutoColumnsFit *bool
    AutoRowsFit *bool
    folder string
    fontsLocation string
    outPath string
    outStorageName string
    password string
    region string
    storageName string
}

func NewExportSpreadsheetAsFormatRequest(format string, name string, opts ...RequestOption) *ExportSpreadsheetAsFormatRequest {
    req := &ExportSpreadsheetAsFormatRequest{
        format: format,
        name: name,
    }
    if req.format == "" {
        return nil
    }
    if req.name == "" {
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
    if val, ok := cfg.Params["folder"].(string); ok {
        req.folder = val
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
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *ExportSpreadsheetAsFormatRequest) GetMethod() string {
    return "GET"
}

func (request *ExportSpreadsheetAsFormatRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *ExportSpreadsheetAsFormatRequest) GetPath() string {
    localVarPath := "/cells/{name}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    return localVarPath
}

func (request *ExportSpreadsheetAsFormatRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("format", fmt.Sprintf("%v", request.format))
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
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

func (request *ExportSpreadsheetAsFormatRequest) GetJSONBody() interface{} {
    return nil
}

func (request *ExportSpreadsheetAsFormatRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *ExportSpreadsheetAsFormatRequest) Description() {
    fmt.Println(strings.Trim("Converts a spreadsheet in cloud storage to the specified format.", " "))
}