package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type DeleteWorksheetValidationRequest struct {
    name string
    sheetName string
    validationIndex int

    folder string
    storageName string
}

func NewDeleteWorksheetValidationRequest(name string, sheetName string, validationIndex int, opts ...RequestOption) *DeleteWorksheetValidationRequest {
    req := &DeleteWorksheetValidationRequest{
        name: name,
        sheetName: sheetName,
        validationIndex: validationIndex,
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

func (request *DeleteWorksheetValidationRequest) GetMethod() string {
    return "DELETE"
}

func (request *DeleteWorksheetValidationRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *DeleteWorksheetValidationRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/validations/{validationIndex}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"validationIndex"+"}", fmt.Sprintf("%v", request.validationIndex), -1)
    return localVarPath
}

func (request *DeleteWorksheetValidationRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *DeleteWorksheetValidationRequest) GetJSONBody() interface{} {
    return nil
}

func (request *DeleteWorksheetValidationRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *DeleteWorksheetValidationRequest) Description() {
    fmt.Println(strings.Trim("Delete a validation by index in worksheet.", " "))
}