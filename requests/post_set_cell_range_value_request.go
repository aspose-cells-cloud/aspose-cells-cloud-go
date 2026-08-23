package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PostSetCellRangeValueRequest struct {
    cellarea string
    name string
    sheetName string
    _type string
    value string

    folder string
    storageName string
}

func NewPostSetCellRangeValueRequest(cellarea string, name string, sheetName string, _type string, value string, opts ...RequestOption) *PostSetCellRangeValueRequest {
    req := &PostSetCellRangeValueRequest{
        cellarea: cellarea,
        name: name,
        sheetName: sheetName,
        _type: _type,
        value: value,
    }
    if req.cellarea == "" {
        return nil
    }
    if req.name == "" {
        return nil
    }
    if req.sheetName == "" {
        return nil
    }
    if req._type == "" {
        return nil
    }
    if req.value == "" {
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

func (request *PostSetCellRangeValueRequest) GetMethod() string {
    return "POST"
}

func (request *PostSetCellRangeValueRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostSetCellRangeValueRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/cells"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PostSetCellRangeValueRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("cellarea", fmt.Sprintf("%v", request.cellarea))
    localVarQueryParams.Add("value", fmt.Sprintf("%v", request.value))
    localVarQueryParams.Add("type", fmt.Sprintf("%v", request._type))
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostSetCellRangeValueRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostSetCellRangeValueRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostSetCellRangeValueRequest) Description() {
    fmt.Println(strings.Trim("Set the value of the range in the worksheet.", " "))
}