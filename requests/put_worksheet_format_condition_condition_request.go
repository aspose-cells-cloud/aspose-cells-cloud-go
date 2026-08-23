package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PutWorksheetFormatConditionConditionRequest struct {
    formula1 string
    formula2 string
    index int
    name string
    operatorType string
    sheetName string
    _type string

    folder string
    storageName string
}

func NewPutWorksheetFormatConditionConditionRequest(formula1 string, formula2 string, index int, name string, operatorType string, sheetName string, _type string, opts ...RequestOption) *PutWorksheetFormatConditionConditionRequest {
    req := &PutWorksheetFormatConditionConditionRequest{
        formula1: formula1,
        formula2: formula2,
        index: index,
        name: name,
        operatorType: operatorType,
        sheetName: sheetName,
        _type: _type,
    }
    if req.formula1 == "" {
        return nil
    }
    if req.formula2 == "" {
        return nil
    }
    if req.name == "" {
        return nil
    }
    if req.operatorType == "" {
        return nil
    }
    if req.sheetName == "" {
        return nil
    }
    if req._type == "" {
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

func (request *PutWorksheetFormatConditionConditionRequest) GetMethod() string {
    return "PUT"
}

func (request *PutWorksheetFormatConditionConditionRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PutWorksheetFormatConditionConditionRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/conditionalFormattings/{index}/condition"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", request.index), -1)
    return localVarPath
}

func (request *PutWorksheetFormatConditionConditionRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("type", fmt.Sprintf("%v", request._type))
    localVarQueryParams.Add("operatorType", fmt.Sprintf("%v", request.operatorType))
    localVarQueryParams.Add("formula1", fmt.Sprintf("%v", request.formula1))
    localVarQueryParams.Add("formula2", fmt.Sprintf("%v", request.formula2))
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PutWorksheetFormatConditionConditionRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PutWorksheetFormatConditionConditionRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PutWorksheetFormatConditionConditionRequest) Description() {
    fmt.Println(strings.Trim("Add a condition for the format condition in the worksheet.", " "))
}