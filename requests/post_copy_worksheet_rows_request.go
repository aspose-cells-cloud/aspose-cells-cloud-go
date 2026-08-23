package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PostCopyWorksheetRowsRequest struct {
    destinationRowIndex int
    name string
    rowNumber int
    sheetName string
    sourceRowIndex int

    folder string
    storageName string
    worksheet string
}

func NewPostCopyWorksheetRowsRequest(destinationRowIndex int, name string, rowNumber int, sheetName string, sourceRowIndex int, opts ...RequestOption) *PostCopyWorksheetRowsRequest {
    req := &PostCopyWorksheetRowsRequest{
        destinationRowIndex: destinationRowIndex,
        name: name,
        rowNumber: rowNumber,
        sheetName: sheetName,
        sourceRowIndex: sourceRowIndex,
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
    if val, ok := cfg.Params["worksheet"].(string); ok {
        req.worksheet = val
    }

    return req
}

func (request *PostCopyWorksheetRowsRequest) GetMethod() string {
    return "POST"
}

func (request *PostCopyWorksheetRowsRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostCopyWorksheetRowsRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/cells/rows/copy"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PostCopyWorksheetRowsRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("sourceRowIndex", fmt.Sprintf("%v", request.sourceRowIndex))
    localVarQueryParams.Add("destinationRowIndex", fmt.Sprintf("%v", request.destinationRowIndex))
    localVarQueryParams.Add("rowNumber", fmt.Sprintf("%v", request.rowNumber))
    if request.worksheet != "" {
        localVarQueryParams.Add("worksheet", fmt.Sprintf("%v", request.worksheet))
    }
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostCopyWorksheetRowsRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostCopyWorksheetRowsRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostCopyWorksheetRowsRequest) Description() {
    fmt.Println(strings.Trim("Copy data and formats from specific entire rows in the worksheet.", " "))
}