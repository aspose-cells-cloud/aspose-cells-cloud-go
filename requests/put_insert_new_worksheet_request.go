package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PutInsertNewWorksheetRequest struct {
    index int
    name string
    sheetName string
    sheettype string

    folder string
    newsheetname string
    storageName string
}

func NewPutInsertNewWorksheetRequest(index int, name string, sheetName string, sheettype string, opts ...RequestOption) *PutInsertNewWorksheetRequest {
    req := &PutInsertNewWorksheetRequest{
        index: index,
        name: name,
        sheetName: sheetName,
        sheettype: sheettype,
    }
    if req.name == "" {
        return nil
    }
    if req.sheetName == "" {
        return nil
    }
    if req.sheettype == "" {
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
    if val, ok := cfg.Params["newsheetname"].(string); ok {
        req.newsheetname = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *PutInsertNewWorksheetRequest) GetMethod() string {
    return "PUT"
}

func (request *PutInsertNewWorksheetRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PutInsertNewWorksheetRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/insert"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    return localVarPath
}

func (request *PutInsertNewWorksheetRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("sheetName", fmt.Sprintf("%v", request.sheetName))
    localVarQueryParams.Add("index", fmt.Sprintf("%v", request.index))
    localVarQueryParams.Add("sheettype", fmt.Sprintf("%v", request.sheettype))
    if request.newsheetname != "" {
        localVarQueryParams.Add("newsheetname", fmt.Sprintf("%v", request.newsheetname))
    }
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PutInsertNewWorksheetRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PutInsertNewWorksheetRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PutInsertNewWorksheetRequest) Description() {
    fmt.Println(strings.Trim("Insert a new worksheet in the workbook.", " "))
}