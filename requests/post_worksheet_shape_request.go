package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostWorksheetShapeRequest struct {
    dto *models.Shape
    name string
    shapeindex int
    sheetName string

    folder string
    storageName string

    extraQueryParameters map[string]string
}

func NewPostWorksheetShapeRequest(dto *models.Shape, name string, shapeindex int, sheetName string, opts ...RequestOption) *PostWorksheetShapeRequest {
    req := &PostWorksheetShapeRequest{
        dto: dto,
        name: name,
        shapeindex: shapeindex,
        sheetName: sheetName,
    }
    if req.dto == nil {
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
    if len(cfg.extraQueryParams) > 0 {
        if req.extraQueryParameters == nil {
            req.extraQueryParameters = make(map[string]string)
        }
        for k, v := range cfg.extraQueryParams {
            req.extraQueryParameters[k] = v
        }
    }

    return req
}

func (request *PostWorksheetShapeRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *PostWorksheetShapeRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
}

func (request *PostWorksheetShapeRequest) GetMethod() string {
    return "POST"
}

func (request *PostWorksheetShapeRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostWorksheetShapeRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/shapes/{shapeindex}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"shapeindex"+"}", fmt.Sprintf("%v", request.shapeindex), -1)
    return localVarPath
}

func (request *PostWorksheetShapeRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    for k, v := range request.extraQueryParameters {
        localVarQueryParams.Add(k, v)
    }
    return localVarQueryParams
}

func (request *PostWorksheetShapeRequest) GetJSONBody() interface{} {
    return &request.dto
}

func (request *PostWorksheetShapeRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostWorksheetShapeRequest) Description() {
    fmt.Println(strings.Trim("Update a shape in the worksheet.", " "))
}