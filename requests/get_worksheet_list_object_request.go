package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type GetWorksheetListObjectRequest struct {
    listobjectindex int
    name string
    sheetName string

    folder string
    format string
    storageName string
}

func NewGetWorksheetListObjectRequest(listobjectindex int, name string, sheetName string, opts ...RequestOption) *GetWorksheetListObjectRequest {
    req := &GetWorksheetListObjectRequest{
        listobjectindex: listobjectindex,
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
    if val, ok := cfg.Params["format"].(string); ok {
        req.format = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *GetWorksheetListObjectRequest) GetMethod() string {
    return "GET"
}

func (request *GetWorksheetListObjectRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *GetWorksheetListObjectRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/listobjects/{listobjectindex}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"listobjectindex"+"}", fmt.Sprintf("%v", request.listobjectindex), -1)
    return localVarPath
}

func (request *GetWorksheetListObjectRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.format != "" {
        localVarQueryParams.Add("format", fmt.Sprintf("%v", request.format))
    }
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *GetWorksheetListObjectRequest) GetJSONBody() interface{} {
    return nil
}

func (request *GetWorksheetListObjectRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *GetWorksheetListObjectRequest) Description() {
    fmt.Println(strings.Trim("Retrieve list object description by index in the worksheet.", " "))
}