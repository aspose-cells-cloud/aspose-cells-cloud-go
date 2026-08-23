package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type DeleteWorksheetFilterRequest struct {
    fieldIndex int
    name string
    sheetName string

    criteria string
    folder string
    storageName string
}

func NewDeleteWorksheetFilterRequest(fieldIndex int, name string, sheetName string, opts ...RequestOption) *DeleteWorksheetFilterRequest {
    req := &DeleteWorksheetFilterRequest{
        fieldIndex: fieldIndex,
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

    if val, ok := cfg.Params["criteria"].(string); ok {
        req.criteria = val
    }
    if val, ok := cfg.Params["folder"].(string); ok {
        req.folder = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *DeleteWorksheetFilterRequest) GetMethod() string {
    return "DELETE"
}

func (request *DeleteWorksheetFilterRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *DeleteWorksheetFilterRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/autoFilter/filter"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *DeleteWorksheetFilterRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("fieldIndex", fmt.Sprintf("%v", request.fieldIndex))
    if request.criteria != "" {
        localVarQueryParams.Add("criteria", fmt.Sprintf("%v", request.criteria))
    }
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *DeleteWorksheetFilterRequest) GetJSONBody() interface{} {
    return nil
}

func (request *DeleteWorksheetFilterRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *DeleteWorksheetFilterRequest) Description() {
    fmt.Println(strings.Trim("Delete a filter for a column in the worksheet.", " "))
}