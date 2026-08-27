package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type SearchBrokenLinksInRemoteRangeRequest struct {
    cellArea string
    name string
    worksheet string

    folder string
    password string
    region string
    storageName string

    extraQueryParameters map[string]string
}

func NewSearchBrokenLinksInRemoteRangeRequest(cellArea string, name string, worksheet string, opts ...RequestOption) *SearchBrokenLinksInRemoteRangeRequest {
    req := &SearchBrokenLinksInRemoteRangeRequest{
        cellArea: cellArea,
        name: name,
        worksheet: worksheet,
    }
    if req.cellArea == "" {
        return nil
    }
    if req.name == "" {
        return nil
    }
    if req.worksheet == "" {
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
    if val, ok := cfg.Params["password"].(string); ok {
        req.password = val
    }
    if val, ok := cfg.Params["region"].(string); ok {
        req.region = val
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

func (request *SearchBrokenLinksInRemoteRangeRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *SearchBrokenLinksInRemoteRangeRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
}

func (request *SearchBrokenLinksInRemoteRangeRequest) GetMethod() string {
    return "PUT"
}

func (request *SearchBrokenLinksInRemoteRangeRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *SearchBrokenLinksInRemoteRangeRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{worksheet}/ranges/{cellArea}/search/broken-links"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"worksheet"+"}", fmt.Sprintf("%v", request.worksheet), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"cellArea"+"}", fmt.Sprintf("%v", request.cellArea), -1)
    return localVarPath
}

func (request *SearchBrokenLinksInRemoteRangeRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    if request.region != "" {
        localVarQueryParams.Add("region", fmt.Sprintf("%v", request.region))
    }
    if request.password != "" {
        localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
    }
    for k, v := range request.extraQueryParameters {
        localVarQueryParams.Add(k, v)
    }
    return localVarQueryParams
}

func (request *SearchBrokenLinksInRemoteRangeRequest) GetJSONBody() interface{} {
    return nil
}

func (request *SearchBrokenLinksInRemoteRangeRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *SearchBrokenLinksInRemoteRangeRequest) Description() {
    fmt.Println(strings.Trim("Search broken links in the range of remoted spreadsheet.", " "))
}