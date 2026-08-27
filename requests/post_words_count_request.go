package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostWordsCountRequest struct {
    wordsCountOptions *models.WordsCountOptions

    extraQueryParameters map[string]string
}

func NewPostWordsCountRequest(wordsCountOptions *models.WordsCountOptions, opts ...RequestOption) *PostWordsCountRequest {
    req := &PostWordsCountRequest{
        wordsCountOptions: wordsCountOptions,
    }
    if req.wordsCountOptions == nil {
        return nil
    }

    cfg := &requestConfig{
        Params: make(map[string]interface{}),
    }
    for _, opt := range opts {
        opt.apply(cfg)
    }

    if len(cfg.extraQueryParams) > 0 {
        if req.extraQueryParameters == nil {
            req.extraQueryParameters = make(map[string]string)
        }
        for k, v := range cfg.extraQueryParams {
            req.extraQueryParameters[k] = v
        }
    }

    return req
}

func (request *PostWordsCountRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *PostWordsCountRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
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
    for k, v := range request.extraQueryParameters {
        localVarQueryParams.Add(k, v)
    }
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