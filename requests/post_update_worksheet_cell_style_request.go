package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostUpdateWorksheetCellStyleRequest struct {
    cellName string
    name string
    sheetName string
    style *models.Style

    folder string
    storageName string
}

func NewPostUpdateWorksheetCellStyleRequest(cellName string, name string, sheetName string, style *models.Style, opts ...RequestOption) *PostUpdateWorksheetCellStyleRequest {
    req := &PostUpdateWorksheetCellStyleRequest{
        cellName: cellName,
        name: name,
        sheetName: sheetName,
        style: style,
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

func (request *PostUpdateWorksheetCellStyleRequest) GetMethod() string {
    return "POST"
}

func (request *PostUpdateWorksheetCellStyleRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostUpdateWorksheetCellStyleRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/cells/{cellName}/style"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"cellName"+"}", fmt.Sprintf("%v", request.cellName), -1)
    return localVarPath
}

func (request *PostUpdateWorksheetCellStyleRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostUpdateWorksheetCellStyleRequest) GetJSONBody() interface{} {
    return &request.style
}

func (request *PostUpdateWorksheetCellStyleRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostUpdateWorksheetCellStyleRequest) Description() {
    fmt.Println(strings.Trim("Set cell style using cell name in the worksheet.", " "))
}