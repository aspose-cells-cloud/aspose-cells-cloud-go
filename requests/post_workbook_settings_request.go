package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostWorkbookSettingsRequest struct {
    name string
    settings *models.WorkbookSettings

    folder string
    storageName string
}

func NewPostWorkbookSettingsRequest(name string, settings *models.WorkbookSettings, opts ...RequestOption) *PostWorkbookSettingsRequest {
    req := &PostWorkbookSettingsRequest{
        name: name,
        settings: settings,
    }
    if req.name == "" {
        return nil
    }
    if req.settings == nil {
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

func (request *PostWorkbookSettingsRequest) GetMethod() string {
    return "POST"
}

func (request *PostWorkbookSettingsRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostWorkbookSettingsRequest) GetPath() string {
    localVarPath := "/cells/{name}/settings"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    return localVarPath
}

func (request *PostWorkbookSettingsRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostWorkbookSettingsRequest) GetJSONBody() interface{} {
    return &request.settings
}

func (request *PostWorkbookSettingsRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostWorkbookSettingsRequest) Description() {
    fmt.Println(strings.Trim("Update setting in the workbook.", " "))
}