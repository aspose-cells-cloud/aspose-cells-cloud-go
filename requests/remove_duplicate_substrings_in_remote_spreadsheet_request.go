package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type RemoveDuplicateSubstringsInRemoteSpreadsheetRequest struct {
    delimiters string
    name string
    _range string
    worksheet string

    caseSensitive *bool
    folder string
    password string
    region string
    storageName string
    treatConsecutiveDelimitersAsOne *bool
}

func NewRemoveDuplicateSubstringsInRemoteSpreadsheetRequest(delimiters string, name string, _range string, worksheet string, opts ...RequestOption) *RemoveDuplicateSubstringsInRemoteSpreadsheetRequest {
    req := &RemoveDuplicateSubstringsInRemoteSpreadsheetRequest{
        delimiters: delimiters,
        name: name,
        _range: _range,
        worksheet: worksheet,
    }
    if req.delimiters == "" {
        return nil
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
    if val, ok := cfg.Params["folder"].(string); ok {
        req.folder = val
    }
    if val, ok := cfg.Params["password"].(string); ok {
        req.password = val
    }
    if val, ok := cfg.Params["region"].(string); ok {
        req.region = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }
    if val, ok := cfg.Params["treatConsecutiveDelimitersAsOne"].(*bool); ok {
        req.treatConsecutiveDelimitersAsOne = val
    }

    return req
}

func (request *RemoveDuplicateSubstringsInRemoteSpreadsheetRequest) GetMethod() string {
    return "PUT"
}

func (request *RemoveDuplicateSubstringsInRemoteSpreadsheetRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *RemoveDuplicateSubstringsInRemoteSpreadsheetRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{worksheet}/range/{range}/content/remove/duplicate-substrings"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"worksheet"+"}", fmt.Sprintf("%v", request.worksheet), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"range"+"}", fmt.Sprintf("%v", request._range), -1)
    return localVarPath
}

func (request *RemoveDuplicateSubstringsInRemoteSpreadsheetRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("delimiters", fmt.Sprintf("%v", request.delimiters))
    if request.treatConsecutiveDelimitersAsOne != nil {
        localVarQueryParams.Add("treatConsecutiveDelimitersAsOne", fmt.Sprintf("%v", *request.treatConsecutiveDelimitersAsOne))
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
    return localVarQueryParams
}

func (request *RemoveDuplicateSubstringsInRemoteSpreadsheetRequest) GetJSONBody() interface{} {
    return nil
}

func (request *RemoveDuplicateSubstringsInRemoteSpreadsheetRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *RemoveDuplicateSubstringsInRemoteSpreadsheetRequest) Description() {
    fmt.Println(strings.Trim("Finds and removes repeated substrings inside every cell of the chosen range, using user-defined or preset delimiters, while preserving formulas, formatting and data-validation.", " "))
}