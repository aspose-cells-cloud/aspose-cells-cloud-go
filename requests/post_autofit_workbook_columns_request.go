package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PostAutofitWorkbookColumnsRequest struct {
    name string

    endColumn *int
    folder string
    startColumn *int
    storageName string
}

func NewPostAutofitWorkbookColumnsRequest(name string, opts ...RequestOption) *PostAutofitWorkbookColumnsRequest {
    req := &PostAutofitWorkbookColumnsRequest{
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

    if val, ok := cfg.Params["endColumn"].(*int); ok {
        req.endColumn = val
    }
    if val, ok := cfg.Params["folder"].(string); ok {
        req.folder = val
    }
    if val, ok := cfg.Params["startColumn"].(*int); ok {
        req.startColumn = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *PostAutofitWorkbookColumnsRequest) GetMethod() string {
    return "POST"
}

func (request *PostAutofitWorkbookColumnsRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostAutofitWorkbookColumnsRequest) GetPath() string {
    localVarPath := "/cells/{name}/autofitcolumns"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    return localVarPath
}

func (request *PostAutofitWorkbookColumnsRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.startColumn != nil {
        localVarQueryParams.Add("startColumn", fmt.Sprintf("%v", *request.startColumn))
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

func (request *PostAutofitWorkbookColumnsRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostAutofitWorkbookColumnsRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostAutofitWorkbookColumnsRequest) Description() {
    fmt.Println(strings.Trim("Autofit columns in the workbook.", " "))
}