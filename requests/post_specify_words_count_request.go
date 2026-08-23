package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostSpecifyWordsCountRequest struct {
    specifyWordsCountOptions *models.SpecifyWordsCountOptions
}

func NewPostSpecifyWordsCountRequest(specifyWordsCountOptions *models.SpecifyWordsCountOptions) *PostSpecifyWordsCountRequest {
    req := &PostSpecifyWordsCountRequest{
        specifyWordsCountOptions: specifyWordsCountOptions,
    }
    if req.specifyWordsCountOptions == nil {
        return nil
    }

    return req
}

func (request *PostSpecifyWordsCountRequest) GetMethod() string {
    return "POST"
}

func (request *PostSpecifyWordsCountRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostSpecifyWordsCountRequest) GetPath() string {
    localVarPath := "/cells/specifywordscount"
    return localVarPath
}

func (request *PostSpecifyWordsCountRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    return localVarQueryParams
}

func (request *PostSpecifyWordsCountRequest) GetJSONBody() interface{} {
    return &request.specifyWordsCountOptions
}

func (request *PostSpecifyWordsCountRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostSpecifyWordsCountRequest) Description() {
    fmt.Println(strings.Trim("PostSpecifyWordsCount", " "))
}