package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostDeleteIncompleteRowsRequest struct {
    deleteIncompleteRowsRequest *models.DeleteIncompleteRowsRequest
}

func NewPostDeleteIncompleteRowsRequest(deleteIncompleteRowsRequest *models.DeleteIncompleteRowsRequest) *PostDeleteIncompleteRowsRequest {
    req := &PostDeleteIncompleteRowsRequest{
        deleteIncompleteRowsRequest: deleteIncompleteRowsRequest,
    }
    if req.deleteIncompleteRowsRequest == nil {
        return nil
    }

    return req
}

func (request *PostDeleteIncompleteRowsRequest) GetMethod() string {
    return "POST"
}

func (request *PostDeleteIncompleteRowsRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostDeleteIncompleteRowsRequest) GetPath() string {
    localVarPath := "/cells/deleteincompleterows"
    return localVarPath
}

func (request *PostDeleteIncompleteRowsRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    return localVarQueryParams
}

func (request *PostDeleteIncompleteRowsRequest) GetJSONBody() interface{} {
    return &request.deleteIncompleteRowsRequest
}

func (request *PostDeleteIncompleteRowsRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostDeleteIncompleteRowsRequest) Description() {
    fmt.Println(strings.Trim("Deleting incomplete rows of spreadsheet files is mainly used to eliminate incomplete rows in tables and ranges.", " "))
}