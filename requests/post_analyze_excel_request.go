package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostAnalyzeExcelRequest struct {
    analyzeExcelRequest *models.AnalyzeExcelRequest

    extraQueryParameters map[string]string
}

func NewPostAnalyzeExcelRequest(analyzeExcelRequest *models.AnalyzeExcelRequest, opts ...RequestOption) *PostAnalyzeExcelRequest {
    req := &PostAnalyzeExcelRequest{
        analyzeExcelRequest: analyzeExcelRequest,
    }
    if req.analyzeExcelRequest == nil {
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

func (request *PostAnalyzeExcelRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *PostAnalyzeExcelRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
}

func (request *PostAnalyzeExcelRequest) GetMethod() string {
    return "POST"
}

func (request *PostAnalyzeExcelRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostAnalyzeExcelRequest) GetPath() string {
    localVarPath := "/cells/analyze"
    return localVarPath
}

func (request *PostAnalyzeExcelRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    for k, v := range request.extraQueryParameters {
        localVarQueryParams.Add(k, v)
    }
    return localVarQueryParams
}

func (request *PostAnalyzeExcelRequest) GetJSONBody() interface{} {
    return &request.analyzeExcelRequest
}

func (request *PostAnalyzeExcelRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostAnalyzeExcelRequest) Description() {
    fmt.Println(strings.Trim("Perform business analysis of data in Excel files.", " "))
}