package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PostUngroupWorksheetColumnsRequest struct {
    firstIndex int
    lastIndex int
    name string
    sheetName string

    folder string
    storageName string
}

func NewPostUngroupWorksheetColumnsRequest(firstIndex int, lastIndex int, name string, sheetName string, opts ...RequestOption) *PostUngroupWorksheetColumnsRequest {
    req := &PostUngroupWorksheetColumnsRequest{
        firstIndex: firstIndex,
        lastIndex: lastIndex,
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

func (request *PostUngroupWorksheetColumnsRequest) GetMethod() string {
    return "POST"
}

func (request *PostUngroupWorksheetColumnsRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostUngroupWorksheetColumnsRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/cells/columns/ungroup"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PostUngroupWorksheetColumnsRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("firstIndex", fmt.Sprintf("%v", request.firstIndex))
    localVarQueryParams.Add("lastIndex", fmt.Sprintf("%v", request.lastIndex))
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostUngroupWorksheetColumnsRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostUngroupWorksheetColumnsRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostUngroupWorksheetColumnsRequest) Description() {
    fmt.Println(strings.Trim("Ungroup worksheet columns.", " "))
}