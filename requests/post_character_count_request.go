package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostCharacterCountRequest struct {
    characterCountOptions *models.CharacterCountOptions
}

func NewPostCharacterCountRequest(characterCountOptions *models.CharacterCountOptions) *PostCharacterCountRequest {
    req := &PostCharacterCountRequest{
        characterCountOptions: characterCountOptions,
    }
    if req.characterCountOptions == nil {
        return nil
    }

    return req
}

func (request *PostCharacterCountRequest) GetMethod() string {
    return "POST"
}

func (request *PostCharacterCountRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostCharacterCountRequest) GetPath() string {
    localVarPath := "/cells/charactercount"
    return localVarPath
}

func (request *PostCharacterCountRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    return localVarQueryParams
}

func (request *PostCharacterCountRequest) GetJSONBody() interface{} {
    return &request.characterCountOptions
}

func (request *PostCharacterCountRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostCharacterCountRequest) Description() {
    fmt.Println(strings.Trim("PostCharacterCount", " "))
}