package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostRemoveCharactersRequest struct {
    removeCharactersOptions *models.RemoveCharactersOptions
}

func NewPostRemoveCharactersRequest(removeCharactersOptions *models.RemoveCharactersOptions) *PostRemoveCharactersRequest {
    req := &PostRemoveCharactersRequest{
        removeCharactersOptions: removeCharactersOptions,
    }
    if req.removeCharactersOptions == nil {
        return nil
    }

    return req
}

func (request *PostRemoveCharactersRequest) GetMethod() string {
    return "POST"
}

func (request *PostRemoveCharactersRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostRemoveCharactersRequest) GetPath() string {
    localVarPath := "/cells/removecharacters"
    return localVarPath
}

func (request *PostRemoveCharactersRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    return localVarQueryParams
}

func (request *PostRemoveCharactersRequest) GetJSONBody() interface{} {
    return &request.removeCharactersOptions
}

func (request *PostRemoveCharactersRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostRemoveCharactersRequest) Description() {
    fmt.Println(strings.Trim("A comprehensive set of tools for cleaning text content within selected cells. It allows users to remove specific characters, character sets, and substrings, ensuring that the text is standardized and free from unwanted symbols or sequences.", " "))
}