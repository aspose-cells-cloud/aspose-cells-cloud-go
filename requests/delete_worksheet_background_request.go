package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type DeleteWorksheetBackgroundRequest struct {
    name string
    sheetName string

    folder string
    storageName string
}

func NewDeleteWorksheetBackgroundRequest(name string, sheetName string, opts ...RequestOption) *DeleteWorksheetBackgroundRequest {
    req := &DeleteWorksheetBackgroundRequest{
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

func (request *DeleteWorksheetBackgroundRequest) GetMethod() string {
    return "DELETE"
}

func (request *DeleteWorksheetBackgroundRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *DeleteWorksheetBackgroundRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/background"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *DeleteWorksheetBackgroundRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *DeleteWorksheetBackgroundRequest) GetJSONBody() interface{} {
    return nil
}

func (request *DeleteWorksheetBackgroundRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *DeleteWorksheetBackgroundRequest) Description() {
    fmt.Println(strings.Trim("Delete background image in the worksheet.", " "))
}