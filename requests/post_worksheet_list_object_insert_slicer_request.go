package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PostWorksheetListObjectInsertSlicerRequest struct {
    columnIndex int
    destCellName string
    listObjectIndex int
    name string
    sheetName string

    folder string
    storageName string
}

func NewPostWorksheetListObjectInsertSlicerRequest(columnIndex int, destCellName string, listObjectIndex int, name string, sheetName string, opts ...RequestOption) *PostWorksheetListObjectInsertSlicerRequest {
    req := &PostWorksheetListObjectInsertSlicerRequest{
        columnIndex: columnIndex,
        destCellName: destCellName,
        listObjectIndex: listObjectIndex,
        name: name,
        sheetName: sheetName,
    }
    if req.destCellName == "" {
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

func (request *PostWorksheetListObjectInsertSlicerRequest) GetMethod() string {
    return "POST"
}

func (request *PostWorksheetListObjectInsertSlicerRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostWorksheetListObjectInsertSlicerRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/listobjects/{listObjectIndex}/InsertSlicer"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"listObjectIndex"+"}", fmt.Sprintf("%v", request.listObjectIndex), -1)
    return localVarPath
}

func (request *PostWorksheetListObjectInsertSlicerRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("columnIndex", fmt.Sprintf("%v", request.columnIndex))
    localVarQueryParams.Add("destCellName", fmt.Sprintf("%v", request.destCellName))
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostWorksheetListObjectInsertSlicerRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostWorksheetListObjectInsertSlicerRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostWorksheetListObjectInsertSlicerRequest) Description() {
    fmt.Println(strings.Trim("Insert slicer for list object.", " "))
}