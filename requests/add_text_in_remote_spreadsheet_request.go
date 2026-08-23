package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type AddTextInRemoteSpreadsheetRequest struct {
    name string
    position string
    _range string
    text string
    worksheet string

    folder string
    password string
    region string
    selectText string
    skipEmptyCells *bool
    storageName string
}

func NewAddTextInRemoteSpreadsheetRequest(name string, position string, _range string, text string, worksheet string, opts ...RequestOption) *AddTextInRemoteSpreadsheetRequest {
    req := &AddTextInRemoteSpreadsheetRequest{
        name: name,
        position: position,
        _range: _range,
        text: text,
        worksheet: worksheet,
    }
    if req.name == "" {
        return nil
    }
    if req.position == "" {
        return nil
    }
    if req._range == "" {
        return nil
    }
    if req.text == "" {
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

    if val, ok := cfg.Params["folder"].(string); ok {
        req.folder = val
    }
    if val, ok := cfg.Params["password"].(string); ok {
        req.password = val
    }
    if val, ok := cfg.Params["region"].(string); ok {
        req.region = val
    }
    if val, ok := cfg.Params["selectText"].(string); ok {
        req.selectText = val
    }
    if val, ok := cfg.Params["skipEmptyCells"].(*bool); ok {
        req.skipEmptyCells = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *AddTextInRemoteSpreadsheetRequest) GetMethod() string {
    return "PUT"
}

func (request *AddTextInRemoteSpreadsheetRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *AddTextInRemoteSpreadsheetRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{worksheet}/range/{range}/content/add/text"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"worksheet"+"}", fmt.Sprintf("%v", request.worksheet), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"range"+"}", fmt.Sprintf("%v", request._range), -1)
    return localVarPath
}

func (request *AddTextInRemoteSpreadsheetRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("text", fmt.Sprintf("%v", request.text))
    localVarQueryParams.Add("position", fmt.Sprintf("%v", request.position))
    if request.selectText != "" {
        localVarQueryParams.Add("selectText", fmt.Sprintf("%v", request.selectText))
    }
    if request.skipEmptyCells != nil {
        localVarQueryParams.Add("skipEmptyCells", fmt.Sprintf("%v", *request.skipEmptyCells))
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

func (request *AddTextInRemoteSpreadsheetRequest) GetJSONBody() interface{} {
    return nil
}

func (request *AddTextInRemoteSpreadsheetRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *AddTextInRemoteSpreadsheetRequest) Description() {
    fmt.Println(strings.Trim("Specify appending text to multiple cells at once, allowing you to add prefixes, suffixes, labels, or any specific characters. You can choose the exact position of the text—in the beginning, at the end, or before or after certain characters in the cell.", " "))
}