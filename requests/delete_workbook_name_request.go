package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type DeleteWorkbookNameRequest struct {
    name string
    nameName string

    folder string
    storageName string
}

func NewDeleteWorkbookNameRequest(name string, nameName string, opts ...RequestOption) *DeleteWorkbookNameRequest {
    req := &DeleteWorkbookNameRequest{
        name: name,
        nameName: nameName,
    }
    if req.name == "" {
        return nil
    }
    if req.nameName == "" {
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

func (request *DeleteWorkbookNameRequest) GetMethod() string {
    return "DELETE"
}

func (request *DeleteWorkbookNameRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *DeleteWorkbookNameRequest) GetPath() string {
    localVarPath := "/cells/{name}/names/{nameName}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nameName"+"}", fmt.Sprintf("%v", request.nameName), -1)
    return localVarPath
}

func (request *DeleteWorkbookNameRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *DeleteWorkbookNameRequest) GetJSONBody() interface{} {
    return nil
}

func (request *DeleteWorkbookNameRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *DeleteWorkbookNameRequest) Description() {
    fmt.Println(strings.Trim("Delete a named range in the workbook.", " "))
}