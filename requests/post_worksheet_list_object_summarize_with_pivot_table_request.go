package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostWorksheetListObjectSummarizeWithPivotTableRequest struct {
    createPivotTableRequest *models.CreatePivotTableRequest
    destsheetName string
    listObjectIndex int
    name string
    sheetName string

    folder string
    storageName string
}

func NewPostWorksheetListObjectSummarizeWithPivotTableRequest(createPivotTableRequest *models.CreatePivotTableRequest, destsheetName string, listObjectIndex int, name string, sheetName string, opts ...RequestOption) *PostWorksheetListObjectSummarizeWithPivotTableRequest {
    req := &PostWorksheetListObjectSummarizeWithPivotTableRequest{
        createPivotTableRequest: createPivotTableRequest,
        destsheetName: destsheetName,
        listObjectIndex: listObjectIndex,
        name: name,
        sheetName: sheetName,
    }
    if req.createPivotTableRequest == nil {
        return nil
    }
    if req.destsheetName == "" {
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

func (request *PostWorksheetListObjectSummarizeWithPivotTableRequest) GetMethod() string {
    return "POST"
}

func (request *PostWorksheetListObjectSummarizeWithPivotTableRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostWorksheetListObjectSummarizeWithPivotTableRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/listobjects/{listObjectIndex}/SummarizeWithPivotTable"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"listObjectIndex"+"}", fmt.Sprintf("%v", request.listObjectIndex), -1)
    return localVarPath
}

func (request *PostWorksheetListObjectSummarizeWithPivotTableRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("destsheetName", fmt.Sprintf("%v", request.destsheetName))
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostWorksheetListObjectSummarizeWithPivotTableRequest) GetJSONBody() interface{} {
    return &request.createPivotTableRequest
}

func (request *PostWorksheetListObjectSummarizeWithPivotTableRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostWorksheetListObjectSummarizeWithPivotTableRequest) Description() {
    fmt.Println(strings.Trim("Create a pivot table with a list object in the worksheet.", " "))
}