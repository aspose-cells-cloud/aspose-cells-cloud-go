package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostDataFillRequest struct {
    dataFillRequest *models.DataFillRequest

    extraQueryParameters map[string]string
}

func NewPostDataFillRequest(dataFillRequest *models.DataFillRequest, opts ...RequestOption) *PostDataFillRequest {
    req := &PostDataFillRequest{
        dataFillRequest: dataFillRequest,
    }
    if req.dataFillRequest == nil {
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

func (request *PostDataFillRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *PostDataFillRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
}

func (request *PostDataFillRequest) GetMethod() string {
    return "POST"
}

func (request *PostDataFillRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostDataFillRequest) GetPath() string {
    localVarPath := "/cells/datafill"
    return localVarPath
}

func (request *PostDataFillRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    for k, v := range request.extraQueryParameters {
        localVarQueryParams.Add(k, v)
    }
    return localVarQueryParams
}

func (request *PostDataFillRequest) GetJSONBody() interface{} {
    return &request.dataFillRequest
}

func (request *PostDataFillRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostDataFillRequest) Description() {
    fmt.Println(strings.Trim("Data filling for spreadsheet files is primarily used to fill empty data in tables and ranges.", " "))
}