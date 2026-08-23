package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostCopyWorksheetRequest struct {
    name string
    options *models.CopyOptions
    sheetName string
    sourceSheet string

    folder string
    sourceFolder string
    sourceWorkbook string
    storageName string
}

func NewPostCopyWorksheetRequest(name string, options *models.CopyOptions, sheetName string, sourceSheet string, opts ...RequestOption) *PostCopyWorksheetRequest {
    req := &PostCopyWorksheetRequest{
        name: name,
        options: options,
        sheetName: sheetName,
        sourceSheet: sourceSheet,
    }
    if req.name == "" {
        return nil
    }
    if req.options == nil {
        return nil
    }
    if req.sheetName == "" {
        return nil
    }
    if req.sourceSheet == "" {
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
    if val, ok := cfg.Params["sourceFolder"].(string); ok {
        req.sourceFolder = val
    }
    if val, ok := cfg.Params["sourceWorkbook"].(string); ok {
        req.sourceWorkbook = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *PostCopyWorksheetRequest) GetMethod() string {
    return "POST"
}

func (request *PostCopyWorksheetRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostCopyWorksheetRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/copy"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PostCopyWorksheetRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("sourceSheet", fmt.Sprintf("%v", request.sourceSheet))
    if request.sourceWorkbook != "" {
        localVarQueryParams.Add("sourceWorkbook", fmt.Sprintf("%v", request.sourceWorkbook))
    }
    if request.sourceFolder != "" {
        localVarQueryParams.Add("sourceFolder", fmt.Sprintf("%v", request.sourceFolder))
    }
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostCopyWorksheetRequest) GetJSONBody() interface{} {
    return &request.options
}

func (request *PostCopyWorksheetRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostCopyWorksheetRequest) Description() {
    fmt.Println(strings.Trim("Copy contents and formats from another worksheet.", " "))
}