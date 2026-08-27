package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PutWorksheetIconFilterRequest struct {
    fieldIndex int
    iconId int
    iconSetType string
    name string
    _range string
    sheetName string

    folder string
    matchBlanks *bool
    refresh *bool
    storageName string

    extraQueryParameters map[string]string
}

func NewPutWorksheetIconFilterRequest(fieldIndex int, iconId int, iconSetType string, name string, _range string, sheetName string, opts ...RequestOption) *PutWorksheetIconFilterRequest {
    req := &PutWorksheetIconFilterRequest{
        fieldIndex: fieldIndex,
        iconId: iconId,
        iconSetType: iconSetType,
        name: name,
        _range: _range,
        sheetName: sheetName,
    }
    if req.iconSetType == "" {
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

func (request *PutWorksheetIconFilterRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *PutWorksheetIconFilterRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
}

func (request *PutWorksheetIconFilterRequest) GetMethod() string {
    return "PUT"
}

func (request *PutWorksheetIconFilterRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PutWorksheetIconFilterRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/autoFilter/iconFilter"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PutWorksheetIconFilterRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("range", fmt.Sprintf("%v", request._range))
    localVarQueryParams.Add("fieldIndex", fmt.Sprintf("%v", request.fieldIndex))
    localVarQueryParams.Add("iconSetType", fmt.Sprintf("%v", request.iconSetType))
    localVarQueryParams.Add("iconId", fmt.Sprintf("%v", request.iconId))
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
    for k, v := range request.extraQueryParameters {
        localVarQueryParams.Add(k, v)
    }
    return localVarQueryParams
}

func (request *PutWorksheetIconFilterRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PutWorksheetIconFilterRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PutWorksheetIconFilterRequest) Description() {
    fmt.Println(strings.Trim("Add an icon filter in the worksheet.", " "))
}