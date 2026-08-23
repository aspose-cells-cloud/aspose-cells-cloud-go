package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostWorksheetCellsRangeOutlineBorderRequest struct {
    name string
    rangeOperate *models.RangeSetOutlineBorderRequest
    sheetName string

    folder string
    storageName string
}

func NewPostWorksheetCellsRangeOutlineBorderRequest(name string, rangeOperate *models.RangeSetOutlineBorderRequest, sheetName string, opts ...RequestOption) *PostWorksheetCellsRangeOutlineBorderRequest {
    req := &PostWorksheetCellsRangeOutlineBorderRequest{
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

func (request *PostWorksheetCellsRangeOutlineBorderRequest) GetMethod() string {
    return "POST"
}

func (request *PostWorksheetCellsRangeOutlineBorderRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostWorksheetCellsRangeOutlineBorderRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/ranges/outlineBorder"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PostWorksheetCellsRangeOutlineBorderRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostWorksheetCellsRangeOutlineBorderRequest) GetJSONBody() interface{} {
    return &request.rangeOperate
}

func (request *PostWorksheetCellsRangeOutlineBorderRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostWorksheetCellsRangeOutlineBorderRequest) Description() {
    fmt.Println(strings.Trim("Apply an outline border around a range of cells.", " "))
}