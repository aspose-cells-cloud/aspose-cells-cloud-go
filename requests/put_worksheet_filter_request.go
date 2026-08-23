package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PutWorksheetFilterRequest struct {
    criteria string
    fieldIndex int
    name string
    _range string
    sheetName string

    folder string
    matchBlanks *bool
    refresh *bool
    storageName string
}

func NewPutWorksheetFilterRequest(criteria string, fieldIndex int, name string, _range string, sheetName string, opts ...RequestOption) *PutWorksheetFilterRequest {
    req := &PutWorksheetFilterRequest{
        criteria: criteria,
        fieldIndex: fieldIndex,
        name: name,
        _range: _range,
        sheetName: sheetName,
    }
    if req.criteria == "" {
        return nil
    }
    if req.name == "" {
        return nil
    }
    if req._range == "" {
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
    if val, ok := cfg.Params["matchBlanks"].(*bool); ok {
        req.matchBlanks = val
    }
    if val, ok := cfg.Params["refresh"].(*bool); ok {
        req.refresh = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *PutWorksheetFilterRequest) GetMethod() string {
    return "PUT"
}

func (request *PutWorksheetFilterRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PutWorksheetFilterRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/autoFilter/filter"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PutWorksheetFilterRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("range", fmt.Sprintf("%v", request._range))
    localVarQueryParams.Add("fieldIndex", fmt.Sprintf("%v", request.fieldIndex))
    localVarQueryParams.Add("criteria", fmt.Sprintf("%v", request.criteria))
    if request.matchBlanks != nil {
        localVarQueryParams.Add("matchBlanks", fmt.Sprintf("%v", *request.matchBlanks))
    }
    if request.refresh != nil {
        localVarQueryParams.Add("refresh", fmt.Sprintf("%v", *request.refresh))
    }
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PutWorksheetFilterRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PutWorksheetFilterRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PutWorksheetFilterRequest) Description() {
    fmt.Println(strings.Trim("Add a filter for a column in the worksheet.", " "))
}