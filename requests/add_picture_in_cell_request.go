package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type AddPictureInCellRequest struct {
    cellName string
    name string
    picturePath string
    sheetName string

    folder string
    storageName string

    extraQueryParameters map[string]string
}

func NewAddPictureInCellRequest(cellName string, name string, picturePath string, sheetName string, opts ...Option) *AddPictureInCellRequest {
    req := &AddPictureInCellRequest{
        cellName: cellName,
        name: name,
        picturePath: picturePath,
        sheetName: sheetName,
    }
    if req.cellName == "" {
        return nil
    }
    if req.name == "" {
        return nil
    }
    if req.picturePath == "" {
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

func (request *AddPictureInCellRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *AddPictureInCellRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
}

func (request *AddPictureInCellRequest) GetMethod() string {
    return "POST"
}

func (request *AddPictureInCellRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *AddPictureInCellRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/pictures/addPictureInCell"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *AddPictureInCellRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("cellName", fmt.Sprintf("%v", request.cellName))
    localVarQueryParams.Add("picturePath", fmt.Sprintf("%v", request.picturePath))
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

func (request *AddPictureInCellRequest) GetJSONBody() interface{} {
    return nil
}

func (request *AddPictureInCellRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *AddPictureInCellRequest) Description() string {
    return strings.Trim("add new picture in the cells.", " ")
}