package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostUpdateWorksheetRangeStyleRequest struct {
    name string
    _range string
    sheetName string
    style *models.Style

    folder string
    storageName string
}

func NewPostUpdateWorksheetRangeStyleRequest(name string, _range string, sheetName string, style *models.Style, opts ...RequestOption) *PostUpdateWorksheetRangeStyleRequest {
    req := &PostUpdateWorksheetRangeStyleRequest{
        name: name,
        _range: _range,
        sheetName: sheetName,
        style: style,
    }
    if req.name == "" {
        return nil
    }
    if req._range == "" {
        return nil
    }
    if req.sheetName == "" {
        return nil
    }
    if req.style == nil {
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

func (request *PostUpdateWorksheetRangeStyleRequest) GetMethod() string {
    return "POST"
}

func (request *PostUpdateWorksheetRangeStyleRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostUpdateWorksheetRangeStyleRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/cells/style"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PostUpdateWorksheetRangeStyleRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("range", fmt.Sprintf("%v", request._range))
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostUpdateWorksheetRangeStyleRequest) GetJSONBody() interface{} {
    return &request.style
}

func (request *PostUpdateWorksheetRangeStyleRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostUpdateWorksheetRangeStyleRequest) Description() {
    fmt.Println(strings.Trim("Update cell range styles in the worksheet.", " "))
}