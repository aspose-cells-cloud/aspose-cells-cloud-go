package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostWorkbookNameRequest struct {
    name string
    nameName string
    newName *models.Name

    folder string
    storageName string
}

func NewPostWorkbookNameRequest(name string, nameName string, newName *models.Name, opts ...RequestOption) *PostWorkbookNameRequest {
    req := &PostWorkbookNameRequest{
        name: name,
        nameName: nameName,
        newName: newName,
    }
    if req.name == "" {
        return nil
    }
    if req.nameName == "" {
        return nil
    }
    if req.newName == nil {
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

func (request *PostWorkbookNameRequest) GetMethod() string {
    return "POST"
}

func (request *PostWorkbookNameRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostWorkbookNameRequest) GetPath() string {
    localVarPath := "/cells/{name}/names/{nameName}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"nameName"+"}", fmt.Sprintf("%v", request.nameName), -1)
    return localVarPath
}

func (request *PostWorkbookNameRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostWorkbookNameRequest) GetJSONBody() interface{} {
    return &request.newName
}

func (request *PostWorkbookNameRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostWorkbookNameRequest) Description() {
    fmt.Println(strings.Trim("Update a named range in the workbook.", " "))
}