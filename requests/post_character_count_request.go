package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostCharacterCountRequest struct {
    characterCountOptions *models.CharacterCountOptions

    extraQueryParameters map[string]string
}

func NewPostCharacterCountRequest(characterCountOptions *models.CharacterCountOptions, opts ...RequestOption) *PostCharacterCountRequest {
    req := &PostCharacterCountRequest{
        characterCountOptions: characterCountOptions,
    }
    if req.characterCountOptions == nil {
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

func (request *PostCharacterCountRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *PostCharacterCountRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
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
    for k, v := range request.extraQueryParameters {
        localVarQueryParams.Add(k, v)
    }
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