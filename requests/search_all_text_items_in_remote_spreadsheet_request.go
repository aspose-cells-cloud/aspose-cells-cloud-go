package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type SearchAllTextItemsInRemoteSpreadsheetRequest struct {
    folder string
    name string

    password string
    region string
    storageName string
}

func NewSearchAllTextItemsInRemoteSpreadsheetRequest(folder string, name string, opts ...RequestOption) *SearchAllTextItemsInRemoteSpreadsheetRequest {
    req := &SearchAllTextItemsInRemoteSpreadsheetRequest{
        folder: folder,
        name: name,
    }
    if req.folder == "" {
        return nil
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

func (request *SearchAllTextItemsInRemoteSpreadsheetRequest) GetMethod() string {
    return "PUT"
}

func (request *SearchAllTextItemsInRemoteSpreadsheetRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *SearchAllTextItemsInRemoteSpreadsheetRequest) GetPath() string {
    localVarPath := "/cells/{name}/search/content/all-textitems"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    return localVarPath
}

func (request *SearchAllTextItemsInRemoteSpreadsheetRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
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

func (request *SearchAllTextItemsInRemoteSpreadsheetRequest) GetJSONBody() interface{} {
    return nil
}

func (request *SearchAllTextItemsInRemoteSpreadsheetRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *SearchAllTextItemsInRemoteSpreadsheetRequest) Description() {
    fmt.Println(strings.Trim("Get all text items in the remote spreadsheet.", " "))
}