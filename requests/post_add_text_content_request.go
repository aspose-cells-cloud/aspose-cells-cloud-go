package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostAddTextContentRequest struct {
    addTextOptions *models.AddTextOptions
}

func NewPostAddTextContentRequest(addTextOptions *models.AddTextOptions) *PostAddTextContentRequest {
    req := &PostAddTextContentRequest{
        addTextOptions: addTextOptions,
    }
    if req.addTextOptions == nil {
        return nil
    }

    return req
}

func (request *PostAddTextContentRequest) GetMethod() string {
    return "POST"
}

func (request *PostAddTextContentRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostAddTextContentRequest) GetPath() string {
    localVarPath := "/cells/addtext"
    return localVarPath
}

func (request *PostAddTextContentRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    return localVarQueryParams
}

func (request *PostAddTextContentRequest) GetJSONBody() interface{} {
    return &request.addTextOptions
}

func (request *PostAddTextContentRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostAddTextContentRequest) Description() {
    fmt.Println(strings.Trim("Adds text content to a specified location within a document. It requires an object that defines the text to be added and the insertion location.", " "))
}