package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PostPivotTableFieldMoveToRequest struct {
    fieldIndex int
    from string
    name string
    pivotTableIndex int
    sheetName string
    to string

    folder string
    storageName string
}

func NewPostPivotTableFieldMoveToRequest(fieldIndex int, from string, name string, pivotTableIndex int, sheetName string, to string, opts ...RequestOption) *PostPivotTableFieldMoveToRequest {
    req := &PostPivotTableFieldMoveToRequest{
        fieldIndex: fieldIndex,
        from: from,
        name: name,
        pivotTableIndex: pivotTableIndex,
        sheetName: sheetName,
        to: to,
    }
    if req.from == "" {
        return nil
    }
    if req.name == "" {
        return nil
    }
    if req.sheetName == "" {
        return nil
    }
    if req.to == "" {
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

func (request *PostPivotTableFieldMoveToRequest) GetMethod() string {
    return "POST"
}

func (request *PostPivotTableFieldMoveToRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostPivotTableFieldMoveToRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/PivotField/Move"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"pivotTableIndex"+"}", fmt.Sprintf("%v", request.pivotTableIndex), -1)
    return localVarPath
}

func (request *PostPivotTableFieldMoveToRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("fieldIndex", fmt.Sprintf("%v", request.fieldIndex))
    localVarQueryParams.Add("from", fmt.Sprintf("%v", request.from))
    localVarQueryParams.Add("to", fmt.Sprintf("%v", request.to))
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostPivotTableFieldMoveToRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostPivotTableFieldMoveToRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostPivotTableFieldMoveToRequest) Description() {
    fmt.Println(strings.Trim("Move a pivot field in the PivotTable.", " "))
}