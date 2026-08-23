package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostWorksheetRangeSortRequest struct {
    cellArea string
    dataSorter *models.DataSorter
    name string
    sheetName string

    folder string
    storageName string
}

func NewPostWorksheetRangeSortRequest(cellArea string, dataSorter *models.DataSorter, name string, sheetName string, opts ...RequestOption) *PostWorksheetRangeSortRequest {
    req := &PostWorksheetRangeSortRequest{
        cellArea: cellArea,
        dataSorter: dataSorter,
        name: name,
        sheetName: sheetName,
    }
    if req.cellArea == "" {
        return nil
    }
    if req.dataSorter == nil {
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

func (request *PostWorksheetRangeSortRequest) GetMethod() string {
    return "POST"
}

func (request *PostWorksheetRangeSortRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostWorksheetRangeSortRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/sort"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PostWorksheetRangeSortRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("cellArea", fmt.Sprintf("%v", request.cellArea))
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostWorksheetRangeSortRequest) GetJSONBody() interface{} {
    return &request.dataSorter
}

func (request *PostWorksheetRangeSortRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostWorksheetRangeSortRequest) Description() {
    fmt.Println(strings.Trim("Sort a range in the worksheet.", " "))
}