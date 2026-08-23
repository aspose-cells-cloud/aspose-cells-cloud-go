package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PostWorksheetAutoFilterRefreshRequest struct {
    name string
    sheetName string

    folder string
    storageName string
}

func NewPostWorksheetAutoFilterRefreshRequest(name string, sheetName string, opts ...RequestOption) *PostWorksheetAutoFilterRefreshRequest {
    req := &PostWorksheetAutoFilterRefreshRequest{
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

    return req
}

func (request *PostWorksheetAutoFilterRefreshRequest) GetMethod() string {
    return "POST"
}

func (request *PostWorksheetAutoFilterRefreshRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostWorksheetAutoFilterRefreshRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/autoFilter/refresh"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PostWorksheetAutoFilterRefreshRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostWorksheetAutoFilterRefreshRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostWorksheetAutoFilterRefreshRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostWorksheetAutoFilterRefreshRequest) Description() {
    fmt.Println(strings.Trim("Refresh auto filters in the worksheet.", " "))
}