package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostWorksheetCommentRequest struct {
    cellName string
    comment *models.Comment
    name string
    sheetName string

    folder string
    storageName string
}

func NewPostWorksheetCommentRequest(cellName string, comment *models.Comment, name string, sheetName string, opts ...RequestOption) *PostWorksheetCommentRequest {
    req := &PostWorksheetCommentRequest{
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

func (request *PostWorksheetCommentRequest) GetMethod() string {
    return "POST"
}

func (request *PostWorksheetCommentRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostWorksheetCommentRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/comments/{cellName}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"cellName"+"}", fmt.Sprintf("%v", request.cellName), -1)
    return localVarPath
}

func (request *PostWorksheetCommentRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostWorksheetCommentRequest) GetJSONBody() interface{} {
    return &request.comment
}

func (request *PostWorksheetCommentRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostWorksheetCommentRequest) Description() {
    fmt.Println(strings.Trim("Update cell comment in the worksheet.", " "))
}