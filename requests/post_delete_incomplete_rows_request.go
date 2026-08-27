package requests

import (
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostDeleteIncompleteRowsRequest struct {
    deleteIncompleteRowsRequest *models.DeleteIncompleteRowsRequest

    extraQueryParameters map[string]string
}

func NewPostDeleteIncompleteRowsRequest(deleteIncompleteRowsRequest *models.DeleteIncompleteRowsRequest, opts ...Option) *PostDeleteIncompleteRowsRequest {
    req := &PostDeleteIncompleteRowsRequest{
        deleteIncompleteRowsRequest: deleteIncompleteRowsRequest,
    }
    if req.deleteIncompleteRowsRequest == nil {
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

func (request *PostDeleteIncompleteRowsRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *PostDeleteIncompleteRowsRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
}

func (request *PostDeleteIncompleteRowsRequest) GetMethod() string {
    return "POST"
}

func (request *PostDeleteIncompleteRowsRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostDeleteIncompleteRowsRequest) GetPath() string {
    localVarPath := "/cells/deleteincompleterows"
    return localVarPath
}

func (request *PostDeleteIncompleteRowsRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    for k, v := range request.extraQueryParameters {
        localVarQueryParams.Add(k, v)
    }
    return localVarQueryParams
}

func (request *PostDeleteIncompleteRowsRequest) GetJSONBody() interface{} {
    return &request.deleteIncompleteRowsRequest
}

func (request *PostDeleteIncompleteRowsRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostDeleteIncompleteRowsRequest) Description() string {
    return strings.Trim("Deleting incomplete rows of spreadsheet files is mainly used to eliminate incomplete rows in tables and ranges.", " ")
}