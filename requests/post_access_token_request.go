package requests

import (
    "net/url"
    "strings"
)

type PostAccessTokenRequest struct {
    extraQueryParameters map[string]string
}

func NewPostAccessTokenRequest(opts ...Option) *PostAccessTokenRequest {
    req := &PostAccessTokenRequest{
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

func (request *PostAccessTokenRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *PostAccessTokenRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
}

func (request *PostAccessTokenRequest) GetMethod() string {
    return "POST"
}

func (request *PostAccessTokenRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostAccessTokenRequest) GetPath() string {
    localVarPath := "/cells/connect/token"
    return localVarPath
}

func (request *PostAccessTokenRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    for k, v := range request.extraQueryParameters {
        localVarQueryParams.Add(k, v)
    }
    return localVarQueryParams
}

func (request *PostAccessTokenRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostAccessTokenRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostAccessTokenRequest) Description() string {
    return strings.Trim("Get Access Token Result: The Cells Cloud Get Token API acts as a proxy service,", " ")
}