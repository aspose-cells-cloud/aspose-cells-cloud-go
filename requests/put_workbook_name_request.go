package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PutWorkbookNameRequest struct {
    name string
    newName *models.Name

    folder string
    storageName string
}

func NewPutWorkbookNameRequest(name string, newName *models.Name, opts ...RequestOption) *PutWorkbookNameRequest {
    req := &PutWorkbookNameRequest{
        name: name,
        newName: newName,
    }
    if req.name == "" {
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

func (request *PutWorkbookNameRequest) GetMethod() string {
    return "PUT"
}

func (request *PutWorkbookNameRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PutWorkbookNameRequest) GetPath() string {
    localVarPath := "/cells/{name}/names"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    return localVarPath
}

func (request *PutWorkbookNameRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PutWorkbookNameRequest) GetJSONBody() interface{} {
    return &request.newName
}

func (request *PutWorkbookNameRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PutWorkbookNameRequest) Description() {
    fmt.Println(strings.Trim("Define a new name in the workbook.", " "))
}