package requests

import (
    "fmt"
    "net/url"
    "path/filepath"
    "strings"
)

type RemoveDuplicateSubstringsRequest struct {
    delimiters string
    Spreadsheet string
    SpreadsheetData []byte
    SpreadsheetName string

    caseSensitive *bool
    outPath string
    outStorageName string
    password string
    _range string
    region string
    treatConsecutiveDelimitersAsOne *bool
    worksheet string

    extraQueryParameters map[string]string
}

func NewRemoveDuplicateSubstringsRequest(delimiters string, Spreadsheet string, opts ...RequestOption) *RemoveDuplicateSubstringsRequest {
    req := &RemoveDuplicateSubstringsRequest{
        delimiters: delimiters,
        Spreadsheet: Spreadsheet,
    }
    if req.delimiters == "" {
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
    if val, ok := cfg.Params["outPath"].(string); ok {
        req.outPath = val
    }
    if val, ok := cfg.Params["outStorageName"].(string); ok {
        req.outStorageName = val
    }
    if val, ok := cfg.Params["password"].(string); ok {
        req.password = val
    }
    if val, ok := cfg.Params["range"].(string); ok {
        req._range = val
    }
    if val, ok := cfg.Params["region"].(string); ok {
        req.region = val
    }
    if val, ok := cfg.Params["treatConsecutiveDelimitersAsOne"].(*bool); ok {
        req.treatConsecutiveDelimitersAsOne = val
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

func (request *RemoveDuplicateSubstringsRequest) SetSpreadsheetBytes(data []byte, name string) {
    if name == "" {
        name = "Spreadsheet"
    }
    request.SpreadsheetData = data
    request.SpreadsheetName = name
}

func (request *RemoveDuplicateSubstringsRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *RemoveDuplicateSubstringsRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
}

func (request *RemoveDuplicateSubstringsRequest) GetMethod() string {
    return "PUT"
}

func (request *RemoveDuplicateSubstringsRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "multipart/form-data"
    return localVarHeaderParams
}

func (request *RemoveDuplicateSubstringsRequest) GetPath() string {
    localVarPath := "/cells/content/remove/duplicate-substrings"
    return localVarPath
}

func (request *RemoveDuplicateSubstringsRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("delimiters", fmt.Sprintf("%v", request.delimiters))
    if request.treatConsecutiveDelimitersAsOne != nil {
        localVarQueryParams.Add("treatConsecutiveDelimitersAsOne", fmt.Sprintf("%v", *request.treatConsecutiveDelimitersAsOne))
    }
    if request.caseSensitive != nil {
        localVarQueryParams.Add("caseSensitive", fmt.Sprintf("%v", *request.caseSensitive))
    }
    if request.worksheet != "" {
        localVarQueryParams.Add("worksheet", fmt.Sprintf("%v", request.worksheet))
    }
    if request._range != "" {
        localVarQueryParams.Add("range", fmt.Sprintf("%v", request._range))
    }
    if request.outPath != "" {
        localVarQueryParams.Add("outPath", fmt.Sprintf("%v", request.outPath))
    }
    if request.outStorageName != "" {
        localVarQueryParams.Add("outStorageName", fmt.Sprintf("%v", request.outStorageName))
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

func (request *RemoveDuplicateSubstringsRequest) GetJSONBody() interface{} {
    return nil
}

func (request *RemoveDuplicateSubstringsRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    if request.SpreadsheetData != nil {
        localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
    } else if request.Spreadsheet != "" {
        localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
    }
    return localVarFormParams
}

func (request *RemoveDuplicateSubstringsRequest) Description() {
    fmt.Println(strings.Trim("Finds and removes repeated substrings inside every cell of the chosen range, using user-defined or preset delimiters, while preserving formulas, formatting and data-validation.", " "))
}