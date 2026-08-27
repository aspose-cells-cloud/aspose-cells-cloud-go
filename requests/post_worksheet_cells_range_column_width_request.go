package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostWorksheetCellsRangeColumnWidthRequest struct {
    name string
    _range *models.Range
    sheetName string
    value float64

    folder string
    storageName string

    extraQueryParameters map[string]string
}

func NewPostWorksheetCellsRangeColumnWidthRequest(name string, _range *models.Range, sheetName string, value float64, opts ...RequestOption) *PostWorksheetCellsRangeColumnWidthRequest {
    req := &PostWorksheetCellsRangeColumnWidthRequest{
        name: name,
        _range: _range,
        sheetName: sheetName,
        value: value,
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

func (request *PostWorksheetCellsRangeColumnWidthRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *PostWorksheetCellsRangeColumnWidthRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
}

func (request *PostWorksheetCellsRangeColumnWidthRequest) GetMethod() string {
    return "POST"
}

func (request *PostWorksheetCellsRangeColumnWidthRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostWorksheetCellsRangeColumnWidthRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/ranges/columnWidth"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PostWorksheetCellsRangeColumnWidthRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("value", fmt.Sprintf("%v", request.value))
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

func (request *PostWorksheetCellsRangeColumnWidthRequest) GetJSONBody() interface{} {
    return &request._range
}

func (request *PostWorksheetCellsRangeColumnWidthRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostWorksheetCellsRangeColumnWidthRequest) Description() {
    fmt.Println(strings.Trim("Set the column width of the specified range.", " "))
}