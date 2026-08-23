package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type GetWorksheetMergedCellRequest struct {
    mergedCellIndex int
    name string
    sheetName string

    folder string
    storageName string
}

func NewGetWorksheetMergedCellRequest(mergedCellIndex int, name string, sheetName string, opts ...RequestOption) *GetWorksheetMergedCellRequest {
    req := &GetWorksheetMergedCellRequest{
        mergedCellIndex: mergedCellIndex,
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

func (request *GetWorksheetMergedCellRequest) GetMethod() string {
    return "GET"
}

func (request *GetWorksheetMergedCellRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *GetWorksheetMergedCellRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/mergedCells/{mergedCellIndex}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"mergedCellIndex"+"}", fmt.Sprintf("%v", request.mergedCellIndex), -1)
    return localVarPath
}

func (request *GetWorksheetMergedCellRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *GetWorksheetMergedCellRequest) GetJSONBody() interface{} {
    return nil
}

func (request *GetWorksheetMergedCellRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *GetWorksheetMergedCellRequest) Description() {
    fmt.Println(strings.Trim("Retrieve description of a merged cell by its index in the worksheet.", " "))
}