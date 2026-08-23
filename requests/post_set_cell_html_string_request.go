package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PostSetCellHtmlStringRequest struct {
    cellName string
    name string
    sheetName string

    folder string
    storageName string
}

func NewPostSetCellHtmlStringRequest(cellName string, name string, sheetName string, opts ...RequestOption) *PostSetCellHtmlStringRequest {
    req := &PostSetCellHtmlStringRequest{
        cellName: cellName,
        name: name,
        sheetName: sheetName,
    }
    if req.cellName == "" {
        return nil
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

func (request *PostSetCellHtmlStringRequest) GetMethod() string {
    return "POST"
}

func (request *PostSetCellHtmlStringRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostSetCellHtmlStringRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/cells/{cellName}/htmlstring"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"cellName"+"}", fmt.Sprintf("%v", request.cellName), -1)
    return localVarPath
}

func (request *PostSetCellHtmlStringRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostSetCellHtmlStringRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostSetCellHtmlStringRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostSetCellHtmlStringRequest) Description() {
    fmt.Println(strings.Trim("Set the HTML string containing data and specific formats in this cell.", " "))
}