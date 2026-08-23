package requests

import (
    "fmt"
    "net/url"
    "path/filepath"
    "strings"
)

type TranslateTextFileRequest struct {
    Spreadsheet string
    SpreadsheetData []byte
    SpreadsheetName string
    targetLanguage string

    password string
    region string
}

func NewTranslateTextFileRequest(Spreadsheet string, targetLanguage string, opts ...RequestOption) *TranslateTextFileRequest {
    req := &TranslateTextFileRequest{
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

func (request *TranslateTextFileRequest) SetSpreadsheetBytes(data []byte, name string) {
    if name == "" {
        name = "Spreadsheet"
    }
    request.SpreadsheetData = data
    request.SpreadsheetName = name
}

func (request *TranslateTextFileRequest) GetMethod() string {
    return "PUT"
}

func (request *TranslateTextFileRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "multipart/form-data"
    return localVarHeaderParams
}

func (request *TranslateTextFileRequest) GetPath() string {
    localVarPath := "/cells/ai/translate/text-file"
    return localVarPath
}

func (request *TranslateTextFileRequest) GetQueryParameters() url.Values {
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

func (request *TranslateTextFileRequest) GetJSONBody() interface{} {
    return nil
}

func (request *TranslateTextFileRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    if request.SpreadsheetData != nil {
        localVarFormParams[request.SpreadsheetName] = request.SpreadsheetData
    } else if request.Spreadsheet != "" {
        localVarFormParams["@"+filepath.Base(request.Spreadsheet)] = request.Spreadsheet
    }
    return localVarFormParams
}

func (request *TranslateTextFileRequest) Description() {
    fmt.Println(strings.Trim("Translates text file content to the specified target language.", " "))
}