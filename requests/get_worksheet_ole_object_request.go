package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type GetWorksheetOleObjectRequest struct {
    name string
    objectNumber int
    sheetName string

    folder string
    format string
    storageName string
}

func NewGetWorksheetOleObjectRequest(name string, objectNumber int, sheetName string, opts ...RequestOption) *GetWorksheetOleObjectRequest {
    req := &GetWorksheetOleObjectRequest{
        name: name,
        objectNumber: objectNumber,
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

func (request *GetWorksheetOleObjectRequest) GetMethod() string {
    return "GET"
}

func (request *GetWorksheetOleObjectRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *GetWorksheetOleObjectRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/oleobjects/{objectNumber}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"objectNumber"+"}", fmt.Sprintf("%v", request.objectNumber), -1)
    return localVarPath
}

func (request *GetWorksheetOleObjectRequest) GetQueryParameters() url.Values {
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

func (request *GetWorksheetOleObjectRequest) GetJSONBody() interface{} {
    return nil
}

func (request *GetWorksheetOleObjectRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *GetWorksheetOleObjectRequest) Description() {
    fmt.Println(strings.Trim("Retrieve the OLE object in a specified format in the worksheet.", " "))
}