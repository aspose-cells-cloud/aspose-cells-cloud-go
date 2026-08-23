package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PostFooterRequest struct {
    isFirstPage bool
    name string
    script string
    section int
    sheetName string

    folder string
    storageName string
}

func NewPostFooterRequest(isFirstPage bool, name string, script string, section int, sheetName string, opts ...RequestOption) *PostFooterRequest {
    req := &PostFooterRequest{
        isFirstPage: isFirstPage,
        name: name,
        script: script,
        section: section,
        sheetName: sheetName,
    }
    if req.name == "" {
        return nil
    }
    if req.script == "" {
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

    if val, ok := cfg.Params["folder"].(string); ok {
        req.folder = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *PostFooterRequest) GetMethod() string {
    return "POST"
}

func (request *PostFooterRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostFooterRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/pagesetup/footer"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PostFooterRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("section", fmt.Sprintf("%v", request.section))
    localVarQueryParams.Add("script", fmt.Sprintf("%v", request.script))
    localVarQueryParams.Add("isFirstPage", fmt.Sprintf("%v", request.isFirstPage))
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostFooterRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostFooterRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostFooterRequest) Description() {
    fmt.Println(strings.Trim("Update page footer in the worksheet.", " "))
}