package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostRemoveDuplicatesRequest struct {
    removeDuplicatesOptions *models.RemoveDuplicatesOptions
}

func NewPostRemoveDuplicatesRequest(removeDuplicatesOptions *models.RemoveDuplicatesOptions) *PostRemoveDuplicatesRequest {
    req := &PostRemoveDuplicatesRequest{
        removeDuplicatesOptions: removeDuplicatesOptions,
    }
    if req.removeDuplicatesOptions == nil {
        return nil
    }

    return req
}

func (request *PostRemoveDuplicatesRequest) GetMethod() string {
    return "POST"
}

func (request *PostRemoveDuplicatesRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostRemoveDuplicatesRequest) GetPath() string {
    localVarPath := "/cells/removeduplicates"
    return localVarPath
}

func (request *PostRemoveDuplicatesRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    return localVarQueryParams
}

func (request *PostRemoveDuplicatesRequest) GetJSONBody() interface{} {
    return &request.removeDuplicatesOptions
}

func (request *PostRemoveDuplicatesRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostRemoveDuplicatesRequest) Description() {
    fmt.Println(strings.Trim("Efficiently remove duplicate substrings from Excel cells. Select a range, specify delimiters, and apply options to eliminate repeated text segments.", " "))
}