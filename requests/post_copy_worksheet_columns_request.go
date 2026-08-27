package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PostCopyWorksheetColumnsRequest struct {
    columnNumber int
    destinationColumnIndex int
    name string
    sheetName string
    sourceColumnIndex int

    folder string
    storageName string
    worksheet string

    extraQueryParameters map[string]string
}

func NewPostCopyWorksheetColumnsRequest(columnNumber int, destinationColumnIndex int, name string, sheetName string, sourceColumnIndex int, opts ...Option) *PostCopyWorksheetColumnsRequest {
    req := &PostCopyWorksheetColumnsRequest{
        columnNumber: columnNumber,
        destinationColumnIndex: destinationColumnIndex,
        name: name,
        sheetName: sheetName,
        sourceColumnIndex: sourceColumnIndex,
    }
    if req.name == "" {
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
    if val, ok := cfg.Params["worksheet"].(string); ok {
        req.worksheet = val
    }
    if len(cfg.extraQueryParams) > 0 {
        if req.extraQueryParameters == nil {
            req.extraQueryParameters = make(map[string]string)
        }
        for k, v := range cfg.extraQueryParams {
            req.extraQueryParameters[k] = v
        }
    }

    return req
}

func (request *PostCopyWorksheetColumnsRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *PostCopyWorksheetColumnsRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
}

func (request *PostCopyWorksheetColumnsRequest) GetMethod() string {
    return "POST"
}

func (request *PostCopyWorksheetColumnsRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostCopyWorksheetColumnsRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/cells/columns/copy"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PostCopyWorksheetColumnsRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("sourceColumnIndex", fmt.Sprintf("%v", request.sourceColumnIndex))
    localVarQueryParams.Add("destinationColumnIndex", fmt.Sprintf("%v", request.destinationColumnIndex))
    localVarQueryParams.Add("columnNumber", fmt.Sprintf("%v", request.columnNumber))
    if request.worksheet != "" {
        localVarQueryParams.Add("worksheet", fmt.Sprintf("%v", request.worksheet))
    }
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    for k, v := range request.extraQueryParameters {
        localVarQueryParams.Add(k, v)
    }
    return localVarQueryParams
}

func (request *PostCopyWorksheetColumnsRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostCopyWorksheetColumnsRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostCopyWorksheetColumnsRequest) Description() string {
    return strings.Trim("Copy data from source columns to destination columns in the worksheet.", " ")
}