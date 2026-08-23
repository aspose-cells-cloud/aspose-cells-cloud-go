package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostUpdateWorksheetPropertyRequest struct {
    name string
    sheet *models.Worksheet
    sheetName string

    folder string
    storageName string
}

func NewPostUpdateWorksheetPropertyRequest(name string, sheet *models.Worksheet, sheetName string, opts ...RequestOption) *PostUpdateWorksheetPropertyRequest {
    req := &PostUpdateWorksheetPropertyRequest{
        name: name,
        sheet: sheet,
        sheetName: sheetName,
    }
    if req.name == "" {
        return nil
    }
    if req.sheet == nil {
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

func (request *PostUpdateWorksheetPropertyRequest) GetMethod() string {
    return "POST"
}

func (request *PostUpdateWorksheetPropertyRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostUpdateWorksheetPropertyRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PostUpdateWorksheetPropertyRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostUpdateWorksheetPropertyRequest) GetJSONBody() interface{} {
    return &request.sheet
}

func (request *PostUpdateWorksheetPropertyRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostUpdateWorksheetPropertyRequest) Description() {
    fmt.Println(strings.Trim("Update worksheet properties in the workbook.", " "))
}