package requests

import (
    "fmt"
    "net/url"
    "path/filepath"
    "strings"
)

type RemoveCharactersRequest struct {
    Spreadsheet string
    SpreadsheetData []byte
    SpreadsheetName string

    caseSensitive *bool
    characterSets string
    outPath string
    outStorageName string
    password string
    _range string
    region string
    removeCustomValue string
    removeTextMethod string
    worksheet string
}

func NewRemoveCharactersRequest(Spreadsheet string, opts ...RequestOption) *RemoveCharactersRequest {
    req := &RemoveCharactersRequest{
        Spreadsheet: Spreadsheet,
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
    if val, ok := cfg.Params["removeCustomValue"].(string); ok {
        req.removeCustomValue = val
    }
    if val, ok := cfg.Params["removeTextMethod"].(string); ok {
        req.removeTextMethod = val
    }
    if val, ok := cfg.Params["worksheet"].(string); ok {
        req.worksheet = val
    }

    return req
}

func (request *RemoveCharactersRequest) SetSpreadsheetBytes(data []byte, name string) {
    if name == "" {
        name = "Spreadsheet"
    }
    request.SpreadsheetData = data
    request.SpreadsheetName = name
}

func (request *RemoveCharactersRequest) GetMethod() string {
    return "PUT"
}

func (request *RemoveCharactersRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "multipart/form-data"
    return localVarHeaderParams
}

func (request *RemoveCharactersRequest) GetPath() string {
    localVarPath := "/cells/content/remove/characters"
    return localVarPath
}

func (request *RemoveCharactersRequest) GetQueryParameters() url.Values {
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
    return localVarQueryParams
}

func (request *RemoveCharactersRequest) GetJSONBody() interface{} {
    return nil
}

func (request *RemoveCharactersRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    if request.SpreadsheetData != nil {
        localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
    } else if request.Spreadsheet != "" {
        localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
    }
    return localVarFormParams
}

func (request *RemoveCharactersRequest) Description() {
    fmt.Println(strings.Trim("Deletes user-defined characters, predefined symbol sets, or any substring from every cell in the chosen range while preserving formulas, formatting and data-validation.", " "))
}