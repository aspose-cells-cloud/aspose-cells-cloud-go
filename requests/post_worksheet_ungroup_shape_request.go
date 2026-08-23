package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PostWorksheetUngroupShapeRequest struct {
    name string
    shapeindex int
    sheetName string

    folder string
    storageName string
}

func NewPostWorksheetUngroupShapeRequest(name string, shapeindex int, sheetName string, opts ...RequestOption) *PostWorksheetUngroupShapeRequest {
    req := &PostWorksheetUngroupShapeRequest{
        name: name,
        shapeindex: shapeindex,
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

func (request *PostWorksheetUngroupShapeRequest) GetMethod() string {
    return "POST"
}

func (request *PostWorksheetUngroupShapeRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostWorksheetUngroupShapeRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/shapes/{shapeindex}/ungroup"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"shapeindex"+"}", fmt.Sprintf("%v", request.shapeindex), -1)
    return localVarPath
}

func (request *PostWorksheetUngroupShapeRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostWorksheetUngroupShapeRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostWorksheetUngroupShapeRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostWorksheetUngroupShapeRequest) Description() {
    fmt.Println(strings.Trim("Ungroup shapes in the worksheet.", " "))
}