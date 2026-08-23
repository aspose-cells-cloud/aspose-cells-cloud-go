package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostCellCharactersRequest struct {
    cellName string
    name string
    sheetName string

    folder string
    options []models.FontSetting
    storageName string
}

func NewPostCellCharactersRequest(cellName string, name string, sheetName string, opts ...RequestOption) *PostCellCharactersRequest {
    req := &PostCellCharactersRequest{
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
    if val, ok := cfg.Params["options"].([]models.FontSetting); ok {
        req.options = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *PostCellCharactersRequest) GetMethod() string {
    return "POST"
}

func (request *PostCellCharactersRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostCellCharactersRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/cells/{cellName}/characters"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"cellName"+"}", fmt.Sprintf("%v", request.cellName), -1)
    return localVarPath
}

func (request *PostCellCharactersRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostCellCharactersRequest) GetJSONBody() interface{} {
    return request.options
}

func (request *PostCellCharactersRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostCellCharactersRequest) Description() {
    fmt.Println(strings.Trim("Set cell characters in the worksheet.", " "))
}