package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PostWorkbookTextReplaceRequest struct {
    name string
    newValue string
    oldValue string

    folder string
    storageName string
}

func NewPostWorkbookTextReplaceRequest(name string, newValue string, oldValue string, opts ...RequestOption) *PostWorkbookTextReplaceRequest {
    req := &PostWorkbookTextReplaceRequest{
        name: name,
        newValue: newValue,
        oldValue: oldValue,
    }
    if req.name == "" {
        return nil
    }
    if req.newValue == "" {
        return nil
    }
    if req.oldValue == "" {
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
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *PostWorkbookTextReplaceRequest) GetMethod() string {
    return "POST"
}

func (request *PostWorkbookTextReplaceRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostWorkbookTextReplaceRequest) GetPath() string {
    localVarPath := "/cells/{name}/replaceText"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    return localVarPath
}

func (request *PostWorkbookTextReplaceRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("oldValue", fmt.Sprintf("%v", request.oldValue))
    localVarQueryParams.Add("newValue", fmt.Sprintf("%v", request.newValue))
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostWorkbookTextReplaceRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostWorkbookTextReplaceRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostWorkbookTextReplaceRequest) Description() {
    fmt.Println(strings.Trim("Replace text in the workbook.", " "))
}