package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostWorksheetValidationRequest struct {
    name string
    sheetName string
    validation *models.Validation
    validationIndex int

    folder string
    storageName string
}

func NewPostWorksheetValidationRequest(name string, sheetName string, validation *models.Validation, validationIndex int, opts ...RequestOption) *PostWorksheetValidationRequest {
    req := &PostWorksheetValidationRequest{
        name: name,
        sheetName: sheetName,
        validation: validation,
        validationIndex: validationIndex,
    }
    if req.name == "" {
        return nil
    }
    if req.sheetName == "" {
        return nil
    }
    if req.validation == nil {
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

func (request *PostWorksheetValidationRequest) GetMethod() string {
    return "POST"
}

func (request *PostWorksheetValidationRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostWorksheetValidationRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/validations/{validationIndex}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"validationIndex"+"}", fmt.Sprintf("%v", request.validationIndex), -1)
    return localVarPath
}

func (request *PostWorksheetValidationRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostWorksheetValidationRequest) GetJSONBody() interface{} {
    return &request.validation
}

func (request *PostWorksheetValidationRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostWorksheetValidationRequest) Description() {
    fmt.Println(strings.Trim("Update a validation by index in the worksheet.", " "))
}