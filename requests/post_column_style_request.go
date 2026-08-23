package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostColumnStyleRequest struct {
    columnIndex int
    name string
    sheetName string
    style *models.Style

    folder string
    storageName string
}

func NewPostColumnStyleRequest(columnIndex int, name string, sheetName string, style *models.Style, opts ...RequestOption) *PostColumnStyleRequest {
    req := &PostColumnStyleRequest{
        columnIndex: columnIndex,
        name: name,
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

func (request *PostColumnStyleRequest) GetMethod() string {
    return "POST"
}

func (request *PostColumnStyleRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostColumnStyleRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/cells/columns/{columnIndex}/style"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"columnIndex"+"}", fmt.Sprintf("%v", request.columnIndex), -1)
    return localVarPath
}

func (request *PostColumnStyleRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostColumnStyleRequest) GetJSONBody() interface{} {
    return &request.style
}

func (request *PostColumnStyleRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostColumnStyleRequest) Description() {
    fmt.Println(strings.Trim("Set column style in the worksheet.", " "))
}