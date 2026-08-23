package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostWorksheetCellsRangeStyleRequest struct {
    name string
    rangeOperate *models.RangeSetStyleRequest
    sheetName string

    folder string
    storageName string
}

func NewPostWorksheetCellsRangeStyleRequest(name string, rangeOperate *models.RangeSetStyleRequest, sheetName string, opts ...RequestOption) *PostWorksheetCellsRangeStyleRequest {
    req := &PostWorksheetCellsRangeStyleRequest{
        name: name,
        rangeOperate: rangeOperate,
        sheetName: sheetName,
    }
    if req.name == "" {
        return nil
    }
    if req.rangeOperate == nil {
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

func (request *PostWorksheetCellsRangeStyleRequest) GetMethod() string {
    return "POST"
}

func (request *PostWorksheetCellsRangeStyleRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostWorksheetCellsRangeStyleRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/ranges/style"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PostWorksheetCellsRangeStyleRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostWorksheetCellsRangeStyleRequest) GetJSONBody() interface{} {
    return &request.rangeOperate
}

func (request *PostWorksheetCellsRangeStyleRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostWorksheetCellsRangeStyleRequest) Description() {
    fmt.Println(strings.Trim("Set the style for the specified range.", " "))
}