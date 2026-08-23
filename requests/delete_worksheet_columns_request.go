package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type DeleteWorksheetColumnsRequest struct {
    columnIndex int
    columns int
    name string
    sheetName string
    updateReference bool

    folder string
    storageName string
}

func NewDeleteWorksheetColumnsRequest(columnIndex int, columns int, name string, sheetName string, updateReference bool, opts ...RequestOption) *DeleteWorksheetColumnsRequest {
    req := &DeleteWorksheetColumnsRequest{
        columnIndex: columnIndex,
        columns: columns,
        name: name,
        sheetName: sheetName,
        updateReference: updateReference,
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

func (request *DeleteWorksheetColumnsRequest) GetMethod() string {
    return "DELETE"
}

func (request *DeleteWorksheetColumnsRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *DeleteWorksheetColumnsRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/cells/columns/{columnIndex}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"columnIndex"+"}", fmt.Sprintf("%v", request.columnIndex), -1)
    return localVarPath
}

func (request *DeleteWorksheetColumnsRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("columns", fmt.Sprintf("%v", request.columns))
    localVarQueryParams.Add("updateReference", fmt.Sprintf("%v", request.updateReference))
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *DeleteWorksheetColumnsRequest) GetJSONBody() interface{} {
    return nil
}

func (request *DeleteWorksheetColumnsRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *DeleteWorksheetColumnsRequest) Description() {
    fmt.Println(strings.Trim("Delete worksheet columns in the worksheet.", " "))
}