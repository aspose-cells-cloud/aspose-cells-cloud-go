package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PostUpdateWorksheetZoomRequest struct {
    name string
    sheetName string
    value int

    folder string
    storageName string
}

func NewPostUpdateWorksheetZoomRequest(name string, sheetName string, value int, opts ...RequestOption) *PostUpdateWorksheetZoomRequest {
    req := &PostUpdateWorksheetZoomRequest{
        name: name,
        sheetName: sheetName,
        value: value,
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

func (request *PostUpdateWorksheetZoomRequest) GetMethod() string {
    return "POST"
}

func (request *PostUpdateWorksheetZoomRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostUpdateWorksheetZoomRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/zoom"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PostUpdateWorksheetZoomRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("value", fmt.Sprintf("%v", request.value))
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostUpdateWorksheetZoomRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostUpdateWorksheetZoomRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostUpdateWorksheetZoomRequest) Description() {
    fmt.Println(strings.Trim("Update the scaling percentage in the worksheet. It should be between 10 and 400.", " "))
}