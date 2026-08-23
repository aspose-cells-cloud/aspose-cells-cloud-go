package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type DeleteHorizontalPageBreaksRequest struct {
    name string
    sheetName string

    folder string
    row *int
    storageName string
}

func NewDeleteHorizontalPageBreaksRequest(name string, sheetName string, opts ...RequestOption) *DeleteHorizontalPageBreaksRequest {
    req := &DeleteHorizontalPageBreaksRequest{
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
    if val, ok := cfg.Params["row"].(*int); ok {
        req.row = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *DeleteHorizontalPageBreaksRequest) GetMethod() string {
    return "DELETE"
}

func (request *DeleteHorizontalPageBreaksRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *DeleteHorizontalPageBreaksRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/horizontalpagebreaks"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *DeleteHorizontalPageBreaksRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.row != nil {
        localVarQueryParams.Add("row", fmt.Sprintf("%v", *request.row))
    }
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *DeleteHorizontalPageBreaksRequest) GetJSONBody() interface{} {
    return nil
}

func (request *DeleteHorizontalPageBreaksRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *DeleteHorizontalPageBreaksRequest) Description() {
    fmt.Println(strings.Trim("Delete horizontal page breaks in the worksheet.", " "))
}