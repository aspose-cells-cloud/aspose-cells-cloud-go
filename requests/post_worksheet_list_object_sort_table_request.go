package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostWorksheetListObjectSortTableRequest struct {
    dataSorter *models.DataSorter
    listObjectIndex int
    name string
    sheetName string

    folder string
    storageName string
}

func NewPostWorksheetListObjectSortTableRequest(dataSorter *models.DataSorter, listObjectIndex int, name string, sheetName string, opts ...RequestOption) *PostWorksheetListObjectSortTableRequest {
    req := &PostWorksheetListObjectSortTableRequest{
        dataSorter: dataSorter,
        listObjectIndex: listObjectIndex,
        name: name,
        sheetName: sheetName,
    }
    if req.dataSorter == nil {
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

func (request *PostWorksheetListObjectSortTableRequest) GetMethod() string {
    return "POST"
}

func (request *PostWorksheetListObjectSortTableRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostWorksheetListObjectSortTableRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/listobjects/{listObjectIndex}/sort"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"listObjectIndex"+"}", fmt.Sprintf("%v", request.listObjectIndex), -1)
    return localVarPath
}

func (request *PostWorksheetListObjectSortTableRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostWorksheetListObjectSortTableRequest) GetJSONBody() interface{} {
    return &request.dataSorter
}

func (request *PostWorksheetListObjectSortTableRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostWorksheetListObjectSortTableRequest) Description() {
    fmt.Println(strings.Trim("Sort list object in the worksheet.", " "))
}