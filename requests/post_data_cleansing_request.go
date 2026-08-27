package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostDataCleansingRequest struct {
    dataCleansingRequest *models.DataCleansingRequest

    extraQueryParameters map[string]string
}

func NewPostDataCleansingRequest(dataCleansingRequest *models.DataCleansingRequest, opts ...RequestOption) *PostDataCleansingRequest {
    req := &PostDataCleansingRequest{
        dataCleansingRequest: dataCleansingRequest,
    }
    if req.dataCleansingRequest == nil {
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

func (request *PostDataCleansingRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *PostDataCleansingRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
}

func (request *PostDataCleansingRequest) GetMethod() string {
    return "POST"
}

func (request *PostDataCleansingRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostDataCleansingRequest) GetPath() string {
    localVarPath := "/cells/datacleansing"
    return localVarPath
}

func (request *PostDataCleansingRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    for k, v := range request.extraQueryParameters {
        localVarQueryParams.Add(k, v)
    }
    return localVarQueryParams
}

func (request *PostDataCleansingRequest) GetJSONBody() interface{} {
    return &request.dataCleansingRequest
}

func (request *PostDataCleansingRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostDataCleansingRequest) Description() {
    fmt.Println(strings.Trim("Data cleansing of spreadsheet files is a data management process used to identify, correct, and remove errors, incompleteness, duplicates, or inaccuracies in tables and ranges.", " "))
}