package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostWorksheetCellsRangeMoveToRequest struct {
    destColumn int
    destRow int
    name string
    _range *models.Range
    sheetName string

    folder string
    storageName string
}

func NewPostWorksheetCellsRangeMoveToRequest(destColumn int, destRow int, name string, _range *models.Range, sheetName string, opts ...RequestOption) *PostWorksheetCellsRangeMoveToRequest {
    req := &PostWorksheetCellsRangeMoveToRequest{
        destColumn: destColumn,
        destRow: destRow,
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

func (request *PostWorksheetCellsRangeMoveToRequest) GetMethod() string {
    return "POST"
}

func (request *PostWorksheetCellsRangeMoveToRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostWorksheetCellsRangeMoveToRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/ranges/moveto"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PostWorksheetCellsRangeMoveToRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("destRow", fmt.Sprintf("%v", request.destRow))
    localVarQueryParams.Add("destColumn", fmt.Sprintf("%v", request.destColumn))
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostWorksheetCellsRangeMoveToRequest) GetJSONBody() interface{} {
    return &request._range
}

func (request *PostWorksheetCellsRangeMoveToRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostWorksheetCellsRangeMoveToRequest) Description() {
    fmt.Println(strings.Trim("Move the current range to the destination range.", " "))
}