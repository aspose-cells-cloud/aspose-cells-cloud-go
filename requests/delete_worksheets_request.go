package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type DeleteWorksheetsRequest struct {
    name string

    folder string
    matchCondition *models.MatchConditionRequest
    storageName string

    extraQueryParameters map[string]string
}

func NewDeleteWorksheetsRequest(name string, opts ...Option) *DeleteWorksheetsRequest {
    req := &DeleteWorksheetsRequest{
        name: name,
    }
    if req.name == "" {
        return nil
    }

    cfg := &requestConfig{
        Params: make(map[string]interface{}),
    }
    for _, opt := range opts {
        opt.apply(cfg)
    }

    if val, ok := cfg.Params["folder"].(string); ok {
        req.folder = val
    }
    if val, ok := cfg.Params["matchCondition"].(*models.MatchConditionRequest); ok {
        req.matchCondition = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
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

func (request *DeleteWorksheetsRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *DeleteWorksheetsRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
}

func (request *DeleteWorksheetsRequest) GetMethod() string {
    return "DELETE"
}

func (request *DeleteWorksheetsRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *DeleteWorksheetsRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    return localVarPath
}

func (request *DeleteWorksheetsRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    for k, v := range request.extraQueryParameters {
        localVarQueryParams.Add(k, v)
    }
    return localVarQueryParams
}

func (request *DeleteWorksheetsRequest) GetJSONBody() interface{} {
    return &request.matchCondition
}

func (request *DeleteWorksheetsRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *DeleteWorksheetsRequest) Description() string {
    return strings.Trim("Delete matched worksheets in the workbook.", " ")
}