package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PostPivotTableFieldHideItemRequest struct {
    fieldIndex int
    isHide bool
    itemIndex int
    name string
    pivotFieldType string
    pivotTableIndex int
    sheetName string

    folder string
    needReCalculate *bool
    storageName string

    extraQueryParameters map[string]string
}

func NewPostPivotTableFieldHideItemRequest(fieldIndex int, isHide bool, itemIndex int, name string, pivotFieldType string, pivotTableIndex int, sheetName string, opts ...RequestOption) *PostPivotTableFieldHideItemRequest {
    req := &PostPivotTableFieldHideItemRequest{
        fieldIndex: fieldIndex,
        isHide: isHide,
        itemIndex: itemIndex,
        name: name,
        pivotFieldType: pivotFieldType,
        pivotTableIndex: pivotTableIndex,
        sheetName: sheetName,
    }
    if req.name == "" {
        return nil
    }
    if req.pivotFieldType == "" {
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
    if val, ok := cfg.Params["needReCalculate"].(*bool); ok {
        req.needReCalculate = val
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

func (request *PostPivotTableFieldHideItemRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *PostPivotTableFieldHideItemRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
}

func (request *PostPivotTableFieldHideItemRequest) GetMethod() string {
    return "POST"
}

func (request *PostPivotTableFieldHideItemRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostPivotTableFieldHideItemRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/PivotField/Hide"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"pivotTableIndex"+"}", fmt.Sprintf("%v", request.pivotTableIndex), -1)
    return localVarPath
}

func (request *PostPivotTableFieldHideItemRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("pivotFieldType", fmt.Sprintf("%v", request.pivotFieldType))
    localVarQueryParams.Add("fieldIndex", fmt.Sprintf("%v", request.fieldIndex))
    localVarQueryParams.Add("itemIndex", fmt.Sprintf("%v", request.itemIndex))
    localVarQueryParams.Add("isHide", fmt.Sprintf("%v", request.isHide))
    if request.needReCalculate != nil {
        localVarQueryParams.Add("needReCalculate", fmt.Sprintf("%v", *request.needReCalculate))
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

func (request *PostPivotTableFieldHideItemRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostPivotTableFieldHideItemRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostPivotTableFieldHideItemRequest) Description() {
    fmt.Println(strings.Trim("Hide a pivot field item in the PivotTable.", " "))
}