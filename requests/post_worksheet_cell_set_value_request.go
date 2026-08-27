package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PostWorksheetCellSetValueRequest struct {
    cellName string
    name string
    sheetName string

    folder string
    formula string
    storageName string
    _type string
    value string

    extraQueryParameters map[string]string
}

func NewPostWorksheetCellSetValueRequest(cellName string, name string, sheetName string, opts ...Option) *PostWorksheetCellSetValueRequest {
    req := &PostWorksheetCellSetValueRequest{
        cellName: cellName,
        name: name,
        sheetName: sheetName,
    }
    if req.cellName == "" {
        return nil
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
    if val, ok := cfg.Params["formula"].(string); ok {
        req.formula = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }
    if val, ok := cfg.Params["type"].(string); ok {
        req._type = val
    }
    if val, ok := cfg.Params["value"].(string); ok {
        req.value = val
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

func (request *PostWorksheetCellSetValueRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *PostWorksheetCellSetValueRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
}

func (request *PostWorksheetCellSetValueRequest) GetMethod() string {
    return "POST"
}

func (request *PostWorksheetCellSetValueRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostWorksheetCellSetValueRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/cells/{cellName}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"cellName"+"}", fmt.Sprintf("%v", request.cellName), -1)
    return localVarPath
}

func (request *PostWorksheetCellSetValueRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.value != "" {
        localVarQueryParams.Add("value", fmt.Sprintf("%v", request.value))
    }
    if request._type != "" {
        localVarQueryParams.Add("type", fmt.Sprintf("%v", request._type))
    }
    if request.formula != "" {
        localVarQueryParams.Add("formula", fmt.Sprintf("%v", request.formula))
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

func (request *PostWorksheetCellSetValueRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostWorksheetCellSetValueRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostWorksheetCellSetValueRequest) Description() string {
    return strings.Trim("Set cell value using cell name in the worksheet.", " ")
}