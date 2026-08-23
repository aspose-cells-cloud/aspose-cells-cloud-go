package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type DeleteWorksheetCellsRangeRequest struct {
    name string
    _range string
    sheetName string
    shift string

    folder string
    storageName string
}

func NewDeleteWorksheetCellsRangeRequest(name string, _range string, sheetName string, shift string, opts ...RequestOption) *DeleteWorksheetCellsRangeRequest {
    req := &DeleteWorksheetCellsRangeRequest{
        name: name,
        _range: _range,
        sheetName: sheetName,
        shift: shift,
    }
    if req.name == "" {
        return nil
    }
    if req._range == "" {
        return nil
    }
    if req.sheetName == "" {
        return nil
    }
    if req.shift == "" {
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

func (request *DeleteWorksheetCellsRangeRequest) GetMethod() string {
    return "DELETE"
}

func (request *DeleteWorksheetCellsRangeRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *DeleteWorksheetCellsRangeRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/ranges"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *DeleteWorksheetCellsRangeRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("range", fmt.Sprintf("%v", request._range))
    localVarQueryParams.Add("shift", fmt.Sprintf("%v", request.shift))
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *DeleteWorksheetCellsRangeRequest) GetJSONBody() interface{} {
    return nil
}

func (request *DeleteWorksheetCellsRangeRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *DeleteWorksheetCellsRangeRequest) Description() {
    fmt.Println(strings.Trim("Delete a range of cells and shift existing cells based on the specified shift option.", " "))
}