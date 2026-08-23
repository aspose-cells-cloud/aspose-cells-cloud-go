package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PostWorkbookSplitRequest struct {
    name string

    folder string
    format string
    from *int
    horizontalResolution *int
    outFolder string
    outStorageName string
    splitNameRule string
    storageName string
    to *int
    verticalResolution *int
}

func NewPostWorkbookSplitRequest(name string, opts ...RequestOption) *PostWorkbookSplitRequest {
    req := &PostWorkbookSplitRequest{
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
    if val, ok := cfg.Params["format"].(string); ok {
        req.format = val
    }
    if val, ok := cfg.Params["from"].(*int); ok {
        req.from = val
    }
    if val, ok := cfg.Params["horizontalResolution"].(*int); ok {
        req.horizontalResolution = val
    }
    if val, ok := cfg.Params["outFolder"].(string); ok {
        req.outFolder = val
    }
    if val, ok := cfg.Params["outStorageName"].(string); ok {
        req.outStorageName = val
    }
    if val, ok := cfg.Params["splitNameRule"].(string); ok {
        req.splitNameRule = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }
    if val, ok := cfg.Params["to"].(*int); ok {
        req.to = val
    }
    if val, ok := cfg.Params["verticalResolution"].(*int); ok {
        req.verticalResolution = val
    }

    return req
}

func (request *PostWorkbookSplitRequest) GetMethod() string {
    return "POST"
}

func (request *PostWorkbookSplitRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostWorkbookSplitRequest) GetPath() string {
    localVarPath := "/cells/{name}/split"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    return localVarPath
}

func (request *PostWorkbookSplitRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.format != "" {
        localVarQueryParams.Add("format", fmt.Sprintf("%v", request.format))
    }
    if request.outFolder != "" {
        localVarQueryParams.Add("outFolder", fmt.Sprintf("%v", request.outFolder))
    }
    if request.from != nil {
        localVarQueryParams.Add("from", fmt.Sprintf("%v", *request.from))
    }
    if request.to != nil {
        localVarQueryParams.Add("to", fmt.Sprintf("%v", *request.to))
    }
    if request.horizontalResolution != nil {
        localVarQueryParams.Add("horizontalResolution", fmt.Sprintf("%v", *request.horizontalResolution))
    }
    if request.verticalResolution != nil {
        localVarQueryParams.Add("verticalResolution", fmt.Sprintf("%v", *request.verticalResolution))
    }
    if request.splitNameRule != "" {
        localVarQueryParams.Add("splitNameRule", fmt.Sprintf("%v", request.splitNameRule))
    }
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    if request.outStorageName != "" {
        localVarQueryParams.Add("outStorageName", fmt.Sprintf("%v", request.outStorageName))
    }
    return localVarQueryParams
}

func (request *PostWorkbookSplitRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostWorkbookSplitRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostWorkbookSplitRequest) Description() {
    fmt.Println(strings.Trim("Split the workbook with a specific format.", " "))
}