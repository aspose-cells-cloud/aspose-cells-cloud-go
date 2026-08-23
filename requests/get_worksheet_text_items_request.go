package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type GetWorksheetTextItemsRequest struct {
    name string
    sheetName string

    folder string
    storageName string
}

func NewGetWorksheetTextItemsRequest(name string, sheetName string, opts ...RequestOption) *GetWorksheetTextItemsRequest {
    req := &GetWorksheetTextItemsRequest{
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

func (request *GetWorksheetTextItemsRequest) GetMethod() string {
    return "GET"
}

func (request *GetWorksheetTextItemsRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *GetWorksheetTextItemsRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/textItems"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *GetWorksheetTextItemsRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *GetWorksheetTextItemsRequest) GetJSONBody() interface{} {
    return nil
}

func (request *GetWorksheetTextItemsRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *GetWorksheetTextItemsRequest) Description() {
    fmt.Println(strings.Trim("Retrieve text items in the worksheet.", " "))
}