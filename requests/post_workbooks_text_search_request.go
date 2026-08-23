package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PostWorkbooksTextSearchRequest struct {
    name string
    text string

    folder string
    storageName string
}

func NewPostWorkbooksTextSearchRequest(name string, text string, opts ...RequestOption) *PostWorkbooksTextSearchRequest {
    req := &PostWorkbooksTextSearchRequest{
        name: name,
        text: text,
    }
    if req.name == "" {
        return nil
    }
    if req.text == "" {
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

func (request *PostWorkbooksTextSearchRequest) GetMethod() string {
    return "POST"
}

func (request *PostWorkbooksTextSearchRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostWorkbooksTextSearchRequest) GetPath() string {
    localVarPath := "/cells/{name}/findText"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    return localVarPath
}

func (request *PostWorkbooksTextSearchRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("text", fmt.Sprintf("%v", request.text))
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostWorkbooksTextSearchRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostWorkbooksTextSearchRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostWorkbooksTextSearchRequest) Description() {
    fmt.Println(strings.Trim("Search for text in the workbook.", " "))
}