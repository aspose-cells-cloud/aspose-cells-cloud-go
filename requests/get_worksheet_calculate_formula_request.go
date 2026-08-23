package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type GetWorksheetCalculateFormulaRequest struct {
    formula string
    name string
    sheetName string

    folder string
    storageName string
}

func NewGetWorksheetCalculateFormulaRequest(formula string, name string, sheetName string, opts ...RequestOption) *GetWorksheetCalculateFormulaRequest {
    req := &GetWorksheetCalculateFormulaRequest{
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

    return req
}

func (request *GetWorksheetCalculateFormulaRequest) GetMethod() string {
    return "GET"
}

func (request *GetWorksheetCalculateFormulaRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *GetWorksheetCalculateFormulaRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/formulaResult"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *GetWorksheetCalculateFormulaRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("formula", fmt.Sprintf("%v", request.formula))
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *GetWorksheetCalculateFormulaRequest) GetJSONBody() interface{} {
    return nil
}

func (request *GetWorksheetCalculateFormulaRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *GetWorksheetCalculateFormulaRequest) Description() {
    fmt.Println(strings.Trim("Calculate formula in the worksheet.", " "))
}