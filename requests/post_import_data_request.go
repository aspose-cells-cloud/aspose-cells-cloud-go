package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostImportDataRequest struct {
    name string

    folder string
    FontsLocation string
    importOption *models.ImportOption
    region string
    storageName string
}

func NewPostImportDataRequest(name string, opts ...RequestOption) *PostImportDataRequest {
    req := &PostImportDataRequest{
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
    if val, ok := cfg.Params["FontsLocation"].(string); ok {
        req.FontsLocation = val
    }
    if val, ok := cfg.Params["importOption"].(*models.ImportOption); ok {
        req.importOption = val
    }
    if val, ok := cfg.Params["region"].(string); ok {
        req.region = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *PostImportDataRequest) GetMethod() string {
    return "POST"
}

func (request *PostImportDataRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostImportDataRequest) GetPath() string {
    localVarPath := "/cells/{name}/importdata"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    return localVarPath
}

func (request *PostImportDataRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    if request.region != "" {
        localVarQueryParams.Add("region", fmt.Sprintf("%v", request.region))
    }
    if request.FontsLocation != "" {
        localVarQueryParams.Add("FontsLocation", fmt.Sprintf("%v", request.FontsLocation))
    }
    return localVarQueryParams
}

func (request *PostImportDataRequest) GetJSONBody() interface{} {
    return &request.importOption
}

func (request *PostImportDataRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostImportDataRequest) Description() {
    fmt.Println(strings.Trim("Import data into the Excel file.", " "))
}