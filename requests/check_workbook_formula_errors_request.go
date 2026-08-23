package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type CheckWorkbookFormulaErrorsRequest struct {
    formulaErrorOptions *models.CheckFormulaErrorOptions
}

func NewCheckWorkbookFormulaErrorsRequest(formulaErrorOptions *models.CheckFormulaErrorOptions) *CheckWorkbookFormulaErrorsRequest {
    req := &CheckWorkbookFormulaErrorsRequest{
        formulaErrorOptions: formulaErrorOptions,
    }
    if req.formulaErrorOptions == nil {
        return nil
    }

    return req
}

func (request *CheckWorkbookFormulaErrorsRequest) GetMethod() string {
    return "POST"
}

func (request *CheckWorkbookFormulaErrorsRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *CheckWorkbookFormulaErrorsRequest) GetPath() string {
    localVarPath := "/cells/checkformulaerrors"
    return localVarPath
}

func (request *CheckWorkbookFormulaErrorsRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    return localVarQueryParams
}

func (request *CheckWorkbookFormulaErrorsRequest) GetJSONBody() interface{} {
    return &request.formulaErrorOptions
}

func (request *CheckWorkbookFormulaErrorsRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *CheckWorkbookFormulaErrorsRequest) Description() {
    fmt.Println(strings.Trim("CheckWorkbookFormulaErrors", " "))
}