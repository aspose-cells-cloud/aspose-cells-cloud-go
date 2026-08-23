package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PostWorkbooksMergeRequest struct {
    mergeWith string
    name string

    folder string
    mergedStorageName string
    storageName string
}

func NewPostWorkbooksMergeRequest(mergeWith string, name string, opts ...RequestOption) *PostWorkbooksMergeRequest {
    req := &PostWorkbooksMergeRequest{
        mergeWith: mergeWith,
        name: name,
    }
    if req.mergeWith == "" {
        return nil
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
    if val, ok := cfg.Params["mergedStorageName"].(string); ok {
        req.mergedStorageName = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *PostWorkbooksMergeRequest) GetMethod() string {
    return "POST"
}

func (request *PostWorkbooksMergeRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostWorkbooksMergeRequest) GetPath() string {
    localVarPath := "/cells/{name}/merge"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    return localVarPath
}

func (request *PostWorkbooksMergeRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("mergeWith", fmt.Sprintf("%v", request.mergeWith))
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    if request.mergedStorageName != "" {
        localVarQueryParams.Add("mergedStorageName", fmt.Sprintf("%v", request.mergedStorageName))
    }
    return localVarQueryParams
}

func (request *PostWorkbooksMergeRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostWorkbooksMergeRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostWorkbooksMergeRequest) Description() {
    fmt.Println(strings.Trim("Merge a workbook into the existing workbook.", " "))
}