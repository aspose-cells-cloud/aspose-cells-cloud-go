package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PostClearFormatsRequest struct {
    name string
    sheetName string

    endColumn *int
    endRow *int
    folder string
    _range string
    startColumn *int
    startRow *int
    storageName string
}

func NewPostClearFormatsRequest(name string, sheetName string, opts ...RequestOption) *PostClearFormatsRequest {
    req := &PostClearFormatsRequest{
        name: name,
        sheetName: sheetName,
    }
    if req.name == "" {
        return nil
    }
    if req.sheetName == "" {
        return nil
    }

    cfg := &requestConfig{
        Params: make(map[string]interface{}),
    }
    for _, opt := range opts {
        opt.apply(cfg)
    }

    if val, ok := cfg.Params["endColumn"].(*int); ok {
        req.endColumn = val
    }
    if val, ok := cfg.Params["endRow"].(*int); ok {
        req.endRow = val
    }
    if val, ok := cfg.Params["folder"].(string); ok {
        req.folder = val
    }
    if val, ok := cfg.Params["range"].(string); ok {
        req._range = val
    }
    if val, ok := cfg.Params["startColumn"].(*int); ok {
        req.startColumn = val
    }
    if val, ok := cfg.Params["startRow"].(*int); ok {
        req.startRow = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *PostClearFormatsRequest) GetMethod() string {
    return "POST"
}

func (request *PostClearFormatsRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostClearFormatsRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/cells/clearformats"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PostClearFormatsRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request._range != "" {
        localVarQueryParams.Add("range", fmt.Sprintf("%v", request._range))
    }
    if request.startRow != nil {
        localVarQueryParams.Add("startRow", fmt.Sprintf("%v", *request.startRow))
    }
    if request.startColumn != nil {
        localVarQueryParams.Add("startColumn", fmt.Sprintf("%v", *request.startColumn))
    }
    if request.endRow != nil {
        localVarQueryParams.Add("endRow", fmt.Sprintf("%v", *request.endRow))
    }
    if request.endColumn != nil {
        localVarQueryParams.Add("endColumn", fmt.Sprintf("%v", *request.endColumn))
    }
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostClearFormatsRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostClearFormatsRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostClearFormatsRequest) Description() {
    fmt.Println(strings.Trim("Clear cell formats in the worksheet.", " "))
}