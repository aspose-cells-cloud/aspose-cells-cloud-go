package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type RemoveCharactersInRemoteSpreadsheetRequest struct {
    name string
    _range string
    worksheet string

    caseSensitive *bool
    characterSets string
    folder string
    password string
    region string
    removeCustomValue string
    removeTextMethod string
    storageName string

    extraQueryParameters map[string]string
}

func NewRemoveCharactersInRemoteSpreadsheetRequest(name string, _range string, worksheet string, opts ...Option) *RemoveCharactersInRemoteSpreadsheetRequest {
    req := &RemoveCharactersInRemoteSpreadsheetRequest{
        name: name,
        _range: _range,
        worksheet: worksheet,
    }
    if req.name == "" {
        return nil
    }
    if req._range == "" {
        return nil
    }
    if req.worksheet == "" {
        return nil
    }

    cfg := &requestConfig{
        Params: make(map[string]interface{}),
    }
    for _, opt := range opts {
        opt.apply(cfg)
    }

    if val, ok := cfg.Params["caseSensitive"].(*bool); ok {
        req.caseSensitive = val
    }
    if val, ok := cfg.Params["characterSets"].(string); ok {
        req.characterSets = val
    }
    if val, ok := cfg.Params["folder"].(string); ok {
        req.folder = val
    }
    if val, ok := cfg.Params["password"].(string); ok {
        req.password = val
    }
    if val, ok := cfg.Params["region"].(string); ok {
        req.region = val
    }
    if val, ok := cfg.Params["removeCustomValue"].(string); ok {
        req.removeCustomValue = val
    }
    if val, ok := cfg.Params["removeTextMethod"].(string); ok {
        req.removeTextMethod = val
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

func (request *RemoveCharactersInRemoteSpreadsheetRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *RemoveCharactersInRemoteSpreadsheetRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
}

func (request *RemoveCharactersInRemoteSpreadsheetRequest) GetMethod() string {
    return "PUT"
}

func (request *RemoveCharactersInRemoteSpreadsheetRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *RemoveCharactersInRemoteSpreadsheetRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{worksheet}/range/{range}/content/remove/characters"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"worksheet"+"}", fmt.Sprintf("%v", request.worksheet), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"range"+"}", fmt.Sprintf("%v", request._range), -1)
    return localVarPath
}

func (request *RemoveCharactersInRemoteSpreadsheetRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.removeTextMethod != "" {
        localVarQueryParams.Add("removeTextMethod", fmt.Sprintf("%v", request.removeTextMethod))
    }
    if request.characterSets != "" {
        localVarQueryParams.Add("characterSets", fmt.Sprintf("%v", request.characterSets))
    }
    if request.removeCustomValue != "" {
        localVarQueryParams.Add("removeCustomValue", fmt.Sprintf("%v", request.removeCustomValue))
    }
    if request.caseSensitive != nil {
        localVarQueryParams.Add("caseSensitive", fmt.Sprintf("%v", *request.caseSensitive))
    }
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    if request.region != "" {
        localVarQueryParams.Add("region", fmt.Sprintf("%v", request.region))
    }
    if request.password != "" {
        localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
    }
    for k, v := range request.extraQueryParameters {
        localVarQueryParams.Add(k, v)
    }
    return localVarQueryParams
}

func (request *RemoveCharactersInRemoteSpreadsheetRequest) GetJSONBody() interface{} {
    return nil
}

func (request *RemoveCharactersInRemoteSpreadsheetRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *RemoveCharactersInRemoteSpreadsheetRequest) Description() string {
    return strings.Trim("Deletes user-defined characters, predefined symbol sets, or any substring from every cell in the chosen range while preserving formulas, formatting and data-validation for a remote spreadsheet.", " ")
}