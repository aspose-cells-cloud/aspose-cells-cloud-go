package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PostWorksheetCalculateFormulaRequest struct {
    formula string
    name string
    sheetName string

    folder string
    storageName string

    extraQueryParameters map[string]string
}

func NewPostWorksheetCalculateFormulaRequest(formula string, name string, sheetName string, opts ...RequestOption) *PostWorksheetCalculateFormulaRequest {
    req := &PostWorksheetCalculateFormulaRequest{
        formula: formula,
        name: name,
        sheetName: sheetName,
    }
    if req.formula == "" {
        return nil
    }
    if req.name == "" {
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

func (request *PostWorksheetCalculateFormulaRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *PostWorksheetCalculateFormulaRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
}

func (request *PostWorksheetCalculateFormulaRequest) GetMethod() string {
    return "POST"
}

func (request *PostWorksheetCalculateFormulaRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostWorksheetCalculateFormulaRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/calculateformula"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PostWorksheetCalculateFormulaRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("formula", fmt.Sprintf("%v", request.formula))
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

func (request *PostWorksheetCalculateFormulaRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostWorksheetCalculateFormulaRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostWorksheetCalculateFormulaRequest) Description() {
    fmt.Println(strings.Trim("Calculate formula in the worksheet.", " "))
}