package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PutPivotTableFieldRequest struct {
    name string
    pivotFieldType string
    pivotTableFieldRequest *models.PivotTableFieldRequest
    pivotTableIndex int
    sheetName string

    folder string
    needReCalculate *bool
    storageName string
}

func NewPutPivotTableFieldRequest(name string, pivotFieldType string, pivotTableFieldRequest *models.PivotTableFieldRequest, pivotTableIndex int, sheetName string, opts ...RequestOption) *PutPivotTableFieldRequest {
    req := &PutPivotTableFieldRequest{
        name: name,
        pivotFieldType: pivotFieldType,
        pivotTableFieldRequest: pivotTableFieldRequest,
        pivotTableIndex: pivotTableIndex,
        sheetName: sheetName,
    }
    if req.name == "" {
        return nil
    }
    if req.pivotFieldType == "" {
        return nil
    }
    if req.pivotTableFieldRequest == nil {
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

    return req
}

func (request *PutPivotTableFieldRequest) GetMethod() string {
    return "PUT"
}

func (request *PutPivotTableFieldRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PutPivotTableFieldRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/PivotField"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"pivotTableIndex"+"}", fmt.Sprintf("%v", request.pivotTableIndex), -1)
    return localVarPath
}

func (request *PutPivotTableFieldRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("pivotFieldType", fmt.Sprintf("%v", request.pivotFieldType))
    if request.needReCalculate != nil {
        localVarQueryParams.Add("needReCalculate", fmt.Sprintf("%v", *request.needReCalculate))
    }
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PutPivotTableFieldRequest) GetJSONBody() interface{} {
    return &request.pivotTableFieldRequest
}

func (request *PutPivotTableFieldRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PutPivotTableFieldRequest) Description() {
    fmt.Println(strings.Trim("Add a pivot field in the PivotTable.", " "))
}