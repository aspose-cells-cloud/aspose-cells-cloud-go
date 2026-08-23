package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostRowStyleRequest struct {
    name string
    rowIndex int
    sheetName string
    style *models.Style

    folder string
    storageName string
}

func NewPostRowStyleRequest(name string, rowIndex int, sheetName string, style *models.Style, opts ...RequestOption) *PostRowStyleRequest {
    req := &PostRowStyleRequest{
        name: name,
        rowIndex: rowIndex,
        sheetName: sheetName,
        style: style,
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

func (request *PostRowStyleRequest) GetMethod() string {
    return "POST"
}

func (request *PostRowStyleRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostRowStyleRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/cells/rows/{rowIndex}/style"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"rowIndex"+"}", fmt.Sprintf("%v", request.rowIndex), -1)
    return localVarPath
}

func (request *PostRowStyleRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostRowStyleRequest) GetJSONBody() interface{} {
    return &request.style
}

func (request *PostRowStyleRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostRowStyleRequest) Description() {
    fmt.Println(strings.Trim("Apply formats to an entire row in the worksheet.", " "))
}