package requests

import (
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostRemoveCharactersRequest struct {
    removeCharactersOptions *models.RemoveCharactersOptions

    extraQueryParameters map[string]string
}

func NewPostRemoveCharactersRequest(removeCharactersOptions *models.RemoveCharactersOptions, opts ...Option) *PostRemoveCharactersRequest {
    req := &PostRemoveCharactersRequest{
        removeCharactersOptions: removeCharactersOptions,
    }
    if req.removeCharactersOptions == nil {
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

func (request *PostRemoveCharactersRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *PostRemoveCharactersRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
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
    for k, v := range request.extraQueryParameters {
        localVarQueryParams.Add(k, v)
    }
    return localVarQueryParams
}

func (request *PostRemoveCharactersRequest) GetJSONBody() interface{} {
    return &request.removeCharactersOptions
}

func (request *PostRemoveCharactersRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostRemoveCharactersRequest) Description() string {
    return strings.Trim("A comprehensive set of tools for cleaning text content within selected cells. It allows users to remove specific characters, character sets, and substrings, ensuring that the text is standardized and free from unwanted symbols or sequences.", " ")
}