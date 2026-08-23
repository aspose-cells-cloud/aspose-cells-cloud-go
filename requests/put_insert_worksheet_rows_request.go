package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PutInsertWorksheetRowsRequest struct {
    name string
    sheetName string
    startrow int

    folder string
    storageName string
    totalRows *int
    updateReference *bool
}

func NewPutInsertWorksheetRowsRequest(name string, sheetName string, startrow int, opts ...RequestOption) *PutInsertWorksheetRowsRequest {
    req := &PutInsertWorksheetRowsRequest{
        name: name,
        sheetName: sheetName,
        startrow: startrow,
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
    if val, ok := cfg.Params["totalRows"].(*int); ok {
        req.totalRows = val
    }
    if val, ok := cfg.Params["updateReference"].(*bool); ok {
        req.updateReference = val
    }

    return req
}

func (request *PutInsertWorksheetRowsRequest) GetMethod() string {
    return "PUT"
}

func (request *PutInsertWorksheetRowsRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PutInsertWorksheetRowsRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/cells/rows/"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PutInsertWorksheetRowsRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("startrow", fmt.Sprintf("%v", request.startrow))
    if request.totalRows != nil {
        localVarQueryParams.Add("totalRows", fmt.Sprintf("%v", *request.totalRows))
    }
    if request.updateReference != nil {
        localVarQueryParams.Add("updateReference", fmt.Sprintf("%v", *request.updateReference))
    }
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PutInsertWorksheetRowsRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PutInsertWorksheetRowsRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PutInsertWorksheetRowsRequest) Description() {
    fmt.Println(strings.Trim("Insert several new rows in the worksheet.", " "))
}