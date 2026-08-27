package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PostCopyCellIntoCellRequest struct {
    destCellName string
    name string
    sheetName string
    worksheet string

    cellname string
    column *int
    folder string
    row *int
    storageName string

    extraQueryParameters map[string]string
}

func NewPostCopyCellIntoCellRequest(destCellName string, name string, sheetName string, worksheet string, opts ...RequestOption) *PostCopyCellIntoCellRequest {
    req := &PostCopyCellIntoCellRequest{
        destCellName: destCellName,
        name: name,
        sheetName: sheetName,
        worksheet: worksheet,
    }
    if req.destCellName == "" {
        return nil
    }
    if req.name == "" {
        return nil
    }
    if req.sheetName == "" {
        return nil
    }
    if req.worksheet == "" {
        return nil
    }

    cfg := &requestConfig{
        Params: make(map[string]interface{}),
    }
    for _, opt := range opts {
        opt.apply(cfg)
    }

    if val, ok := cfg.Params["cellname"].(string); ok {
        req.cellname = val
    }
    if val, ok := cfg.Params["column"].(*int); ok {
        req.column = val
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

func (request *PostCopyCellIntoCellRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *PostCopyCellIntoCellRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
}

func (request *PostCopyCellIntoCellRequest) GetMethod() string {
    return "POST"
}

func (request *PostCopyCellIntoCellRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostCopyCellIntoCellRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/cells/{destCellName}/copy"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"destCellName"+"}", fmt.Sprintf("%v", request.destCellName), -1)
    return localVarPath
}

func (request *PostCopyCellIntoCellRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("worksheet", fmt.Sprintf("%v", request.worksheet))
    if request.cellname != "" {
        localVarQueryParams.Add("cellname", fmt.Sprintf("%v", request.cellname))
    }
    if request.row != nil {
        localVarQueryParams.Add("row", fmt.Sprintf("%v", *request.row))
    }
    if request.column != nil {
        localVarQueryParams.Add("column", fmt.Sprintf("%v", *request.column))
    }
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

func (request *PostCopyCellIntoCellRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostCopyCellIntoCellRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostCopyCellIntoCellRequest) Description() {
    fmt.Println(strings.Trim("Copy data from a source cell to a destination cell in the worksheet.", " "))
}