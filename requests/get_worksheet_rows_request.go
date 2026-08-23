package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type GetWorksheetRowsRequest struct {
    name string
    sheetName string

    count *int
    folder string
    offset *int
    storageName string
}

func NewGetWorksheetRowsRequest(name string, sheetName string, opts ...RequestOption) *GetWorksheetRowsRequest {
    req := &GetWorksheetRowsRequest{
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

    if val, ok := cfg.Params["count"].(*int); ok {
        req.count = val
    }
    if val, ok := cfg.Params["folder"].(string); ok {
        req.folder = val
    }
    if val, ok := cfg.Params["offset"].(*int); ok {
        req.offset = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *GetWorksheetRowsRequest) GetMethod() string {
    return "GET"
}

func (request *GetWorksheetRowsRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *GetWorksheetRowsRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/cells/rows/"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *GetWorksheetRowsRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.offset != nil {
        localVarQueryParams.Add("offset", fmt.Sprintf("%v", *request.offset))
    }
    if request.count != nil {
        localVarQueryParams.Add("count", fmt.Sprintf("%v", *request.count))
    }
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *GetWorksheetRowsRequest) GetJSONBody() interface{} {
    return nil
}

func (request *GetWorksheetRowsRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *GetWorksheetRowsRequest) Description() {
    fmt.Println(strings.Trim("Retrieve descriptions of rows in the worksheet.", " "))
}