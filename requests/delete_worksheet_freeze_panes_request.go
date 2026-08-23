package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type DeleteWorksheetFreezePanesRequest struct {
    column int
    freezedColumns int
    freezedRows int
    name string
    row int
    sheetName string

    folder string
    storageName string
}

func NewDeleteWorksheetFreezePanesRequest(column int, freezedColumns int, freezedRows int, name string, row int, sheetName string, opts ...RequestOption) *DeleteWorksheetFreezePanesRequest {
    req := &DeleteWorksheetFreezePanesRequest{
        column: column,
        freezedColumns: freezedColumns,
        freezedRows: freezedRows,
        name: name,
        row: row,
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

func (request *DeleteWorksheetFreezePanesRequest) GetMethod() string {
    return "DELETE"
}

func (request *DeleteWorksheetFreezePanesRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *DeleteWorksheetFreezePanesRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/freezepanes"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *DeleteWorksheetFreezePanesRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("row", fmt.Sprintf("%v", request.row))
    localVarQueryParams.Add("column", fmt.Sprintf("%v", request.column))
    localVarQueryParams.Add("freezedRows", fmt.Sprintf("%v", request.freezedRows))
    localVarQueryParams.Add("freezedColumns", fmt.Sprintf("%v", request.freezedColumns))
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *DeleteWorksheetFreezePanesRequest) GetJSONBody() interface{} {
    return nil
}

func (request *DeleteWorksheetFreezePanesRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *DeleteWorksheetFreezePanesRequest) Description() {
    fmt.Println(strings.Trim("Unfreeze panes in worksheet.", " "))
}