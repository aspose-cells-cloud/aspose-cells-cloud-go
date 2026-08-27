package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type GetVerticalPageBreaksRequest struct {
    name string
    sheetName string

    folder string
    storageName string

    extraQueryParameters map[string]string
}

func NewGetVerticalPageBreaksRequest(name string, sheetName string, opts ...RequestOption) *GetVerticalPageBreaksRequest {
    req := &GetVerticalPageBreaksRequest{
        name: name,
        sheetName: sheetName,
    }
    if req.name == "" {
        return nil
    }
    if req.sheetName == "" {
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

func (request *GetVerticalPageBreaksRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *GetVerticalPageBreaksRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
}

func (request *GetVerticalPageBreaksRequest) GetMethod() string {
    return "GET"
}

func (request *GetVerticalPageBreaksRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *GetVerticalPageBreaksRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/verticalpagebreaks"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *GetVerticalPageBreaksRequest) GetQueryParameters() url.Values {
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

func (request *GetVerticalPageBreaksRequest) GetJSONBody() interface{} {
    return nil
}

func (request *GetVerticalPageBreaksRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *GetVerticalPageBreaksRequest) Description() {
    fmt.Println(strings.Trim("Retrieve descriptions of vertical page breaks in the worksheet.", " "))
}