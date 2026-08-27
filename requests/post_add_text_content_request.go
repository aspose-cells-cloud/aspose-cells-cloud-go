package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostAddTextContentRequest struct {
    addTextOptions *models.AddTextOptions

    extraQueryParameters map[string]string
}

func NewPostAddTextContentRequest(addTextOptions *models.AddTextOptions, opts ...RequestOption) *PostAddTextContentRequest {
    req := &PostAddTextContentRequest{
        addTextOptions: addTextOptions,
    }
    if req.addTextOptions == nil {
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

func (request *PostAddTextContentRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *PostAddTextContentRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
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
    for k, v := range request.extraQueryParameters {
        localVarQueryParams.Add(k, v)
    }
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