package requests

import (
    "fmt"
    "net/url"
    "path/filepath"
    "strings"
)

type AddTextRequest struct {
    position string
    Spreadsheet string
    SpreadsheetData []byte
    SpreadsheetName string
    text string

    outPath string
    outStorageName string
    password string
    _range string
    region string
    selectText string
    skipEmptyCells *bool
    worksheet string

    extraQueryParameters map[string]string
}

func NewAddTextRequest(position string, Spreadsheet string, text string, opts ...Option) *AddTextRequest {
    req := &AddTextRequest{
        position: position,
        Spreadsheet: Spreadsheet,
        text: text,
    }
    if req.position == "" {
        return nil
    }
    if req.text == "" {
        return nil
    }

    cfg := &requestConfig{
        Params: make(map[string]interface{}),
    }
    for _, opt := range opts {
        opt.apply(cfg)
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
    if val, ok := cfg.Params["selectText"].(string); ok {
        req.selectText = val
    }
    if val, ok := cfg.Params["skipEmptyCells"].(*bool); ok {
        req.skipEmptyCells = val
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

func (request *AddTextRequest) SetSpreadsheetBytes(data []byte, name string) {
    if name == "" {
        name = "Spreadsheet"
    }
    request.SpreadsheetData = data
    request.SpreadsheetName = name
}

func (request *AddTextRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *AddTextRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
}

func (request *AddTextRequest) GetMethod() string {
    return "PUT"
}

func (request *AddTextRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "multipart/form-data"
    return localVarHeaderParams
}

func (request *AddTextRequest) GetPath() string {
    localVarPath := "/cells/content/add/text"
    return localVarPath
}

func (request *AddTextRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("text", fmt.Sprintf("%v", request.text))
    localVarQueryParams.Add("position", fmt.Sprintf("%v", request.position))
    if request.selectText != "" {
        localVarQueryParams.Add("selectText", fmt.Sprintf("%v", request.selectText))
    }
    if request.skipEmptyCells != nil {
        localVarQueryParams.Add("skipEmptyCells", fmt.Sprintf("%v", *request.skipEmptyCells))
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

func (request *AddTextRequest) GetJSONBody() interface{} {
    return nil
}

func (request *AddTextRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    if request.SpreadsheetData != nil {
        localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
    } else if request.Spreadsheet != "" {
        localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
    }
    return localVarFormParams
}

func (request *AddTextRequest) Description() string {
    return strings.Trim("Specify appending text to multiple cells at once, allowing you to add prefixes, suffixes, labels, or any specific characters. You can choose the exact position of the text—in the beginning, at the end, or before or after certain characters in the cell.", " ")
}