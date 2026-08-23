package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type CheckWrokbookExternalReferenceRequest struct {
    checkExternalReferenceOptions *models.CheckExternalReferenceOptions
}

func NewCheckWrokbookExternalReferenceRequest(checkExternalReferenceOptions *models.CheckExternalReferenceOptions) *CheckWrokbookExternalReferenceRequest {
    req := &CheckWrokbookExternalReferenceRequest{
        checkExternalReferenceOptions: checkExternalReferenceOptions,
    }
    if req.checkExternalReferenceOptions == nil {
        return nil
    }

    return req
}

func (request *CheckWrokbookExternalReferenceRequest) GetMethod() string {
    return "POST"
}

func (request *CheckWrokbookExternalReferenceRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *CheckWrokbookExternalReferenceRequest) GetPath() string {
    localVarPath := "/cells/checkexternalreference"
    return localVarPath
}

func (request *CheckWrokbookExternalReferenceRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    return localVarQueryParams
}

func (request *CheckWrokbookExternalReferenceRequest) GetJSONBody() interface{} {
    return &request.checkExternalReferenceOptions
}

func (request *CheckWrokbookExternalReferenceRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *CheckWrokbookExternalReferenceRequest) Description() {
    fmt.Println(strings.Trim("Export Excel internal elements or the workbook itself to various format files.", " "))
}