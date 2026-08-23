package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostWorksheetCellsRangeUnMergeRequest struct {
    name string
    _range *models.Range
    sheetName string

    folder string
    storageName string
}

func NewPostWorksheetCellsRangeUnMergeRequest(name string, _range *models.Range, sheetName string, opts ...RequestOption) *PostWorksheetCellsRangeUnMergeRequest {
    req := &PostWorksheetCellsRangeUnMergeRequest{
        name: name,
        _range: _range,
        sheetName: sheetName,
    }
    if req.name == "" {
        return nil
    }
    if req._range == nil {
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

func (request *PostWorksheetCellsRangeUnMergeRequest) GetMethod() string {
    return "POST"
}

func (request *PostWorksheetCellsRangeUnMergeRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostWorksheetCellsRangeUnMergeRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/ranges/unmerge"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PostWorksheetCellsRangeUnMergeRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostWorksheetCellsRangeUnMergeRequest) GetJSONBody() interface{} {
    return &request._range
}

func (request *PostWorksheetCellsRangeUnMergeRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostWorksheetCellsRangeUnMergeRequest) Description() {
    fmt.Println(strings.Trim("Unmerge merged cells within this range.", " "))
}