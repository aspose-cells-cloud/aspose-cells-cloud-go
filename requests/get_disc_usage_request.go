package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type GetDiscUsageRequest struct {

    storageName string
}

func NewGetDiscUsageRequest(opts ...RequestOption) *GetDiscUsageRequest {
    req := &GetDiscUsageRequest{
    }
    cfg := &requestConfig{
        Params: make(map[string]interface{}),
    }
    for _, opt := range opts {
        opt.apply(cfg)
    }

    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *GetDiscUsageRequest) GetMethod() string {
    return "GET"
}

func (request *GetDiscUsageRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *GetDiscUsageRequest) GetPath() string {
    localVarPath := "/cells/storage/disc"
    return localVarPath
}

func (request *GetDiscUsageRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *GetDiscUsageRequest) GetJSONBody() interface{} {
    return nil
}

func (request *GetDiscUsageRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *GetDiscUsageRequest) Description() {
    fmt.Println(strings.Trim("GetDiscUsage", " "))
}