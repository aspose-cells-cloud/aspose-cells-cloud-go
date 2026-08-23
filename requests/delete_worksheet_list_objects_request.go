package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type DeleteWorksheetListObjectsRequest struct {
    name string
    sheetName string

    folder string
    storageName string
}

func NewDeleteWorksheetListObjectsRequest(name string, sheetName string, opts ...RequestOption) *DeleteWorksheetListObjectsRequest {
    req := &DeleteWorksheetListObjectsRequest{
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

func (request *DeleteWorksheetListObjectsRequest) GetMethod() string {
    return "DELETE"
}

func (request *DeleteWorksheetListObjectsRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *DeleteWorksheetListObjectsRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/listobjects"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *DeleteWorksheetListObjectsRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *DeleteWorksheetListObjectsRequest) GetJSONBody() interface{} {
    return nil
}

func (request *DeleteWorksheetListObjectsRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *DeleteWorksheetListObjectsRequest) Description() {
    fmt.Println(strings.Trim("Delete ListObjects in the worksheet.", " "))
}