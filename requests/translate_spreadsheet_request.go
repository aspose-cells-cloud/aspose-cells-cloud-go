package requests

import (
    "fmt"
    "net/url"
    "path/filepath"
    "strings"
)

type TranslateSpreadsheetRequest struct {
    Spreadsheet string
    SpreadsheetData []byte
    SpreadsheetName string
    targetLanguage string

    password string
    region string
}

func NewTranslateSpreadsheetRequest(Spreadsheet string, targetLanguage string, opts ...RequestOption) *TranslateSpreadsheetRequest {
    req := &TranslateSpreadsheetRequest{
        Spreadsheet: Spreadsheet,
        targetLanguage: targetLanguage,
    }
    if req.targetLanguage == "" {
        return nil
    }

    cfg := &requestConfig{
        Params: make(map[string]interface{}),
    }
    for _, opt := range opts {
        opt.apply(cfg)
    }

    if val, ok := cfg.Params["password"].(string); ok {
        req.password = val
    }
    if val, ok := cfg.Params["region"].(string); ok {
        req.region = val
    }

    return req
}

func (request *TranslateSpreadsheetRequest) SetSpreadsheetBytes(data []byte, name string) {
    if name == "" {
        name = "Spreadsheet"
    }
    request.SpreadsheetData = data
    request.SpreadsheetName = name
}

func (request *TranslateSpreadsheetRequest) GetMethod() string {
    return "PUT"
}

func (request *TranslateSpreadsheetRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "multipart/form-data"
    return localVarHeaderParams
}

func (request *TranslateSpreadsheetRequest) GetPath() string {
    localVarPath := "/cells/ai/translate/spreadsheet"
    return localVarPath
}

func (request *TranslateSpreadsheetRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("targetLanguage", fmt.Sprintf("%v", request.targetLanguage))
    if request.region != "" {
        localVarQueryParams.Add("region", fmt.Sprintf("%v", request.region))
    }
    if request.password != "" {
        localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
    }
    return localVarQueryParams
}

func (request *TranslateSpreadsheetRequest) GetJSONBody() interface{} {
    return nil
}

func (request *TranslateSpreadsheetRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    if request.SpreadsheetData != nil {
        localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
    } else if request.Spreadsheet != "" {
        localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
    }
    return localVarFormParams
}

func (request *TranslateSpreadsheetRequest) Description() {
    fmt.Println(strings.Trim("Translates the entire spreadsheet to the specified target language.", " "))
}