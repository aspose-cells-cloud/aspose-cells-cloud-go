package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostWordsCountRequest struct {
    wordsCountOptions *models.WordsCountOptions
}

func NewPostWordsCountRequest(wordsCountOptions *models.WordsCountOptions) *PostWordsCountRequest {
    req := &PostWordsCountRequest{
        wordsCountOptions: wordsCountOptions,
    }
    if req.wordsCountOptions == nil {
        return nil
    }

    return req
}

func (request *PostWordsCountRequest) GetMethod() string {
    return "POST"
}

func (request *PostWordsCountRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostWordsCountRequest) GetPath() string {
    localVarPath := "/cells/wordscount"
    return localVarPath
}

func (request *PostWordsCountRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    return localVarQueryParams
}

func (request *PostWordsCountRequest) GetJSONBody() interface{} {
    return &request.wordsCountOptions
}

func (request *PostWordsCountRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostWordsCountRequest) Description() {
    fmt.Println(strings.Trim("PostWordsCount", " "))
}