package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type DeleteWorkbookBackgroundRequest struct {
    name string

    folder string
    storageName string
}

func NewDeleteWorkbookBackgroundRequest(name string, opts ...RequestOption) *DeleteWorkbookBackgroundRequest {
    req := &DeleteWorkbookBackgroundRequest{
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
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *DeleteWorkbookBackgroundRequest) GetMethod() string {
    return "DELETE"
}

func (request *DeleteWorkbookBackgroundRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *DeleteWorkbookBackgroundRequest) GetPath() string {
    localVarPath := "/cells/{name}/background"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    return localVarPath
}

func (request *DeleteWorkbookBackgroundRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *DeleteWorkbookBackgroundRequest) GetJSONBody() interface{} {
    return nil
}

func (request *DeleteWorkbookBackgroundRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *DeleteWorkbookBackgroundRequest) Description() {
    fmt.Println(strings.Trim("Delete background in the workbook.", " "))
}