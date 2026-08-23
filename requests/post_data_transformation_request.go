package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostDataTransformationRequest struct {
    dataTransformationRequest *models.DataTransformationRequest
}

func NewPostDataTransformationRequest(dataTransformationRequest *models.DataTransformationRequest) *PostDataTransformationRequest {
    req := &PostDataTransformationRequest{
        dataTransformationRequest: dataTransformationRequest,
    }
    if req.dataTransformationRequest == nil {
        return nil
    }

    return req
}

func (request *PostDataTransformationRequest) GetMethod() string {
    return "POST"
}

func (request *PostDataTransformationRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostDataTransformationRequest) GetPath() string {
    localVarPath := "/cells/datatransformation"
    return localVarPath
}

func (request *PostDataTransformationRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    return localVarQueryParams
}

func (request *PostDataTransformationRequest) GetJSONBody() interface{} {
    return &request.dataTransformationRequest
}

func (request *PostDataTransformationRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostDataTransformationRequest) Description() {
    fmt.Println(strings.Trim("Transform spreadsheet data is mainly used to pivot columns, unpivot columns.", " "))
}