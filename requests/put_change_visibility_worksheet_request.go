package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PutChangeVisibilityWorksheetRequest struct {
    isVisible bool
    name string
    sheetName string

    folder string
    storageName string
}

func NewPutChangeVisibilityWorksheetRequest(isVisible bool, name string, sheetName string, opts ...RequestOption) *PutChangeVisibilityWorksheetRequest {
    req := &PutChangeVisibilityWorksheetRequest{
        isVisible: isVisible,
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

func (request *PutChangeVisibilityWorksheetRequest) GetMethod() string {
    return "PUT"
}

func (request *PutChangeVisibilityWorksheetRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PutChangeVisibilityWorksheetRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/visible"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PutChangeVisibilityWorksheetRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("isVisible", fmt.Sprintf("%v", request.isVisible))
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PutChangeVisibilityWorksheetRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PutChangeVisibilityWorksheetRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PutChangeVisibilityWorksheetRequest) Description() {
    fmt.Println(strings.Trim("Change worksheet visibility in the workbook.", " "))
}