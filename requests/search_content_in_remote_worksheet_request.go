package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type SearchContentInRemoteWorksheetRequest struct {
    name string
    searchText string
    worksheet string

    folder string
    ignoringCase *bool
    password string
    region string
    storageName string
}

func NewSearchContentInRemoteWorksheetRequest(name string, searchText string, worksheet string, opts ...RequestOption) *SearchContentInRemoteWorksheetRequest {
    req := &SearchContentInRemoteWorksheetRequest{
        name: name,
        searchText: searchText,
        worksheet: worksheet,
    }
    if req.name == "" {
        return nil
    }
    if req.searchText == "" {
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
    if val, ok := cfg.Params["ignoringCase"].(*bool); ok {
        req.ignoringCase = val
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

    return req
}

func (request *SearchContentInRemoteWorksheetRequest) GetMethod() string {
    return "PUT"
}

func (request *SearchContentInRemoteWorksheetRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *SearchContentInRemoteWorksheetRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{worksheet}/search/content"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"worksheet"+"}", fmt.Sprintf("%v", request.worksheet), -1)
    return localVarPath
}

func (request *SearchContentInRemoteWorksheetRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("searchText", fmt.Sprintf("%v", request.searchText))
    if request.ignoringCase != nil {
        localVarQueryParams.Add("ignoringCase", fmt.Sprintf("%v", *request.ignoringCase))
    }
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
    return localVarQueryParams
}

func (request *SearchContentInRemoteWorksheetRequest) GetJSONBody() interface{} {
    return nil
}

func (request *SearchContentInRemoteWorksheetRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *SearchContentInRemoteWorksheetRequest) Description() {
    fmt.Println(strings.Trim("Search text in the worksheet of remoted spreadsheet.", " "))
}