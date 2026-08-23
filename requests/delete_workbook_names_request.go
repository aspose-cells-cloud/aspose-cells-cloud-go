package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type DeleteWorkbookNamesRequest struct {
    name string

    folder string
    storageName string
}

func NewDeleteWorkbookNamesRequest(name string, opts ...RequestOption) *DeleteWorkbookNamesRequest {
    req := &DeleteWorkbookNamesRequest{
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

func (request *DeleteWorkbookNamesRequest) GetMethod() string {
    return "DELETE"
}

func (request *DeleteWorkbookNamesRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *DeleteWorkbookNamesRequest) GetPath() string {
    localVarPath := "/cells/{name}/names"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    return localVarPath
}

func (request *DeleteWorkbookNamesRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *DeleteWorkbookNamesRequest) GetJSONBody() interface{} {
    return nil
}

func (request *DeleteWorkbookNamesRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *DeleteWorkbookNamesRequest) Description() {
    fmt.Println(strings.Trim("Delete all named ranges in the workbook.", " "))
}