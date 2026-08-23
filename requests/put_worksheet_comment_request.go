package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PutWorksheetCommentRequest struct {
    cellName string
    comment *models.Comment
    name string
    sheetName string

    folder string
    storageName string
}

func NewPutWorksheetCommentRequest(cellName string, comment *models.Comment, name string, sheetName string, opts ...RequestOption) *PutWorksheetCommentRequest {
    req := &PutWorksheetCommentRequest{
        cellName: cellName,
        comment: comment,
        name: name,
        sheetName: sheetName,
    }
    if req.cellName == "" {
        return nil
    }
    if req.comment == nil {
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

    return req
}

func (request *PutWorksheetCommentRequest) GetMethod() string {
    return "PUT"
}

func (request *PutWorksheetCommentRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PutWorksheetCommentRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/comments/{cellName}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"cellName"+"}", fmt.Sprintf("%v", request.cellName), -1)
    return localVarPath
}

func (request *PutWorksheetCommentRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PutWorksheetCommentRequest) GetJSONBody() interface{} {
    return &request.comment
}

func (request *PutWorksheetCommentRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PutWorksheetCommentRequest) Description() {
    fmt.Println(strings.Trim("Add cell comment in the worksheet.", " "))
}