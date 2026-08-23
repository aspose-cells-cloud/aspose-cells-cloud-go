package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PutWorksheetValidationRequest struct {
    name string
    sheetName string

    folder string
    _range string
    storageName string
}

func NewPutWorksheetValidationRequest(name string, sheetName string, opts ...RequestOption) *PutWorksheetValidationRequest {
    req := &PutWorksheetValidationRequest{
        name: name,
        sheetName: sheetName,
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
    if val, ok := cfg.Params["range"].(string); ok {
        req._range = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *PutWorksheetValidationRequest) GetMethod() string {
    return "PUT"
}

func (request *PutWorksheetValidationRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PutWorksheetValidationRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/validations"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PutWorksheetValidationRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request._range != "" {
        localVarQueryParams.Add("range", fmt.Sprintf("%v", request._range))
    }
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PutWorksheetValidationRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PutWorksheetValidationRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PutWorksheetValidationRequest) Description() {
    fmt.Println(strings.Trim("Add a validation at index in the worksheet.", " "))
}