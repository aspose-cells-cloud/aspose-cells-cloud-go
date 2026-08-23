package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostWorksheetListColumnsTotalRequest struct {
    listObjectIndex int
    name string
    sheetName string
    tableTotalRequests []models.TableTotalRequest

    folder string
    storageName string
}

func NewPostWorksheetListColumnsTotalRequest(listObjectIndex int, name string, sheetName string, tableTotalRequests []models.TableTotalRequest, opts ...RequestOption) *PostWorksheetListColumnsTotalRequest {
    req := &PostWorksheetListColumnsTotalRequest{
        listObjectIndex: listObjectIndex,
        name: name,
        sheetName: sheetName,
        tableTotalRequests: tableTotalRequests,
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

func (request *PostWorksheetListColumnsTotalRequest) GetMethod() string {
    return "POST"
}

func (request *PostWorksheetListColumnsTotalRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostWorksheetListColumnsTotalRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/listobjects/{listObjectIndex}/listcolumns/total"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"listObjectIndex"+"}", fmt.Sprintf("%v", request.listObjectIndex), -1)
    return localVarPath
}

func (request *PostWorksheetListColumnsTotalRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostWorksheetListColumnsTotalRequest) GetJSONBody() interface{} {
    return request.tableTotalRequests
}

func (request *PostWorksheetListColumnsTotalRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostWorksheetListColumnsTotalRequest) Description() {
    fmt.Println(strings.Trim("Update total of list columns in the table.", " "))
}