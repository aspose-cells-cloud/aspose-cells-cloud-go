package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostConvertWorkbookRequest struct {
    convertWorkbookOptions *models.ConvertWorkbookOptions

    FontsLocation string

    extraQueryParameters map[string]string
}

func NewPostConvertWorkbookRequest(convertWorkbookOptions *models.ConvertWorkbookOptions, opts ...RequestOption) *PostConvertWorkbookRequest {
    req := &PostConvertWorkbookRequest{
        convertWorkbookOptions: convertWorkbookOptions,
    }
    if req.convertWorkbookOptions == nil {
        return nil
    }

    cfg := &requestConfig{
        Params: make(map[string]interface{}),
    }
    for _, opt := range opts {
        opt.apply(cfg)
    }

    if val, ok := cfg.Params["FontsLocation"].(string); ok {
        req.FontsLocation = val
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

func (request *PostConvertWorkbookRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *PostConvertWorkbookRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
}

func (request *PostConvertWorkbookRequest) GetMethod() string {
    return "POST"
}

func (request *PostConvertWorkbookRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostConvertWorkbookRequest) GetPath() string {
    localVarPath := "/cells/convertWorkbook"
    return localVarPath
}

func (request *PostConvertWorkbookRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.FontsLocation != "" {
        localVarQueryParams.Add("FontsLocation", fmt.Sprintf("%v", request.FontsLocation))
    }
    for k, v := range request.extraQueryParameters {
        localVarQueryParams.Add(k, v)
    }
    return localVarQueryParams
}

func (request *PostConvertWorkbookRequest) GetJSONBody() interface{} {
    return &request.convertWorkbookOptions
}

func (request *PostConvertWorkbookRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostConvertWorkbookRequest) Description() {
    fmt.Println(strings.Trim("PostConvertWorkbook", " "))
}