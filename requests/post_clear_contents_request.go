package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PostClearContentsRequest struct {
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

func NewPostClearContentsRequest(name string, sheetName string, opts ...RequestOption) *PostClearContentsRequest {
    req := &PostClearContentsRequest{
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

func (request *PostClearContentsRequest) GetMethod() string {
    return "POST"
}

func (request *PostClearContentsRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostClearContentsRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/cells/clearcontents"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PostClearContentsRequest) GetQueryParameters() url.Values {
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

func (request *PostClearContentsRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostClearContentsRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostClearContentsRequest) Description() {
    fmt.Println(strings.Trim("Clear cell area contents in the worksheet.", " "))
}