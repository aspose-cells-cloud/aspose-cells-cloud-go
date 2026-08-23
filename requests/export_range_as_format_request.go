package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type ExportRangeAsFormatRequest struct {
    format string
    name string
    _range string
    worksheet string

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

func NewExportRangeAsFormatRequest(format string, name string, _range string, worksheet string, opts ...RequestOption) *ExportRangeAsFormatRequest {
    req := &ExportRangeAsFormatRequest{
        format: format,
        name: name,
        _range: _range,
        worksheet: worksheet,
    }
    if req.format == "" {
        return nil
    }
    if req.name == "" {
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

func (request *ExportRangeAsFormatRequest) GetMethod() string {
    return "GET"
}

func (request *ExportRangeAsFormatRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *ExportRangeAsFormatRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{worksheet}/ranges/{range}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"worksheet"+"}", fmt.Sprintf("%v", request.worksheet), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"range"+"}", fmt.Sprintf("%v", request._range), -1)
    return localVarPath
}

func (request *ExportRangeAsFormatRequest) GetQueryParameters() url.Values {
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

func (request *ExportRangeAsFormatRequest) GetJSONBody() interface{} {
    return nil
}

func (request *ExportRangeAsFormatRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *ExportRangeAsFormatRequest) Description() {
    fmt.Println(strings.Trim("Converts the range of spreadsheet in cloud storage to the specified format.", " "))
}