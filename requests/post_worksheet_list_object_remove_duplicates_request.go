package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PostWorksheetListObjectRemoveDuplicatesRequest struct {
    listObjectIndex int
    name string
    sheetName string

    folder string
    storageName string
}

func NewPostWorksheetListObjectRemoveDuplicatesRequest(listObjectIndex int, name string, sheetName string, opts ...RequestOption) *PostWorksheetListObjectRemoveDuplicatesRequest {
    req := &PostWorksheetListObjectRemoveDuplicatesRequest{
        listObjectIndex: listObjectIndex,
        name: name,
        sheetName: sheetName,
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

func (request *PostWorksheetListObjectRemoveDuplicatesRequest) GetMethod() string {
    return "POST"
}

func (request *PostWorksheetListObjectRemoveDuplicatesRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostWorksheetListObjectRemoveDuplicatesRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/listobjects/{listObjectIndex}/RemoveDuplicates"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"listObjectIndex"+"}", fmt.Sprintf("%v", request.listObjectIndex), -1)
    return localVarPath
}

func (request *PostWorksheetListObjectRemoveDuplicatesRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostWorksheetListObjectRemoveDuplicatesRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostWorksheetListObjectRemoveDuplicatesRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostWorksheetListObjectRemoveDuplicatesRequest) Description() {
    fmt.Println(strings.Trim("Remove duplicates in list object.", " "))
}