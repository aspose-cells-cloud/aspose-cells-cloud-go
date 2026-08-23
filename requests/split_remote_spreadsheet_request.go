package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type SplitRemoteSpreadsheetRequest struct {
    name string

    folder string
    fontsLocation string
    from *int
    outFormat string
    outPath string
    outStorageName string
    password string
    region string
    storageName string
    to *int
}

func NewSplitRemoteSpreadsheetRequest(name string, opts ...RequestOption) *SplitRemoteSpreadsheetRequest {
    req := &SplitRemoteSpreadsheetRequest{
        name: name,
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

    if val, ok := cfg.Params["folder"].(string); ok {
        req.folder = val
    }
    if val, ok := cfg.Params["fontsLocation"].(string); ok {
        req.fontsLocation = val
    }
    if val, ok := cfg.Params["from"].(*int); ok {
        req.from = val
    }
    if val, ok := cfg.Params["outFormat"].(string); ok {
        req.outFormat = val
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
    if val, ok := cfg.Params["to"].(*int); ok {
        req.to = val
    }

    return req
}

func (request *SplitRemoteSpreadsheetRequest) GetMethod() string {
    return "PUT"
}

func (request *SplitRemoteSpreadsheetRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *SplitRemoteSpreadsheetRequest) GetPath() string {
    localVarPath := "/cells/{name}/split/spreadsheet"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    return localVarPath
}

func (request *SplitRemoteSpreadsheetRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.from != nil {
        localVarQueryParams.Add("from", fmt.Sprintf("%v", *request.from))
    }
    if request.to != nil {
        localVarQueryParams.Add("to", fmt.Sprintf("%v", *request.to))
    }
    if request.outFormat != "" {
        localVarQueryParams.Add("outFormat", fmt.Sprintf("%v", request.outFormat))
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
    if request.region != "" {
        localVarQueryParams.Add("region", fmt.Sprintf("%v", request.region))
    }
    if request.password != "" {
        localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
    }
    return localVarQueryParams
}

func (request *SplitRemoteSpreadsheetRequest) GetJSONBody() interface{} {
    return nil
}

func (request *SplitRemoteSpreadsheetRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *SplitRemoteSpreadsheetRequest) Description() {
    fmt.Println(strings.Trim("Split a spreadsheet in cloud storage into the specified format, multi-file.", " "))
}