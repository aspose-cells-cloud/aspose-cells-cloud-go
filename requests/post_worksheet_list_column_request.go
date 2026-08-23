package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostWorksheetListColumnRequest struct {
    columnIndex int
    listColumn *models.ListColumn
    listObjectIndex int
    name string
    sheetName string

    folder string
    storageName string
}

func NewPostWorksheetListColumnRequest(columnIndex int, listColumn *models.ListColumn, listObjectIndex int, name string, sheetName string, opts ...RequestOption) *PostWorksheetListColumnRequest {
    req := &PostWorksheetListColumnRequest{
        columnIndex: columnIndex,
        listColumn: listColumn,
        listObjectIndex: listObjectIndex,
        name: name,
        sheetName: sheetName,
    }
    if req.listColumn == nil {
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

func (request *PostWorksheetListColumnRequest) GetMethod() string {
    return "POST"
}

func (request *PostWorksheetListColumnRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostWorksheetListColumnRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/listobjects/{listObjectIndex}/listcolumns/{columnIndex}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"listObjectIndex"+"}", fmt.Sprintf("%v", request.listObjectIndex), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"columnIndex"+"}", fmt.Sprintf("%v", request.columnIndex), -1)
    return localVarPath
}

func (request *PostWorksheetListColumnRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostWorksheetListColumnRequest) GetJSONBody() interface{} {
    return &request.listColumn
}

func (request *PostWorksheetListColumnRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostWorksheetListColumnRequest) Description() {
    fmt.Println(strings.Trim("Update list column in list object.", " "))
}