package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type GetMergedCellsInRemotedWorksheetRequest struct {
    name string
    worksheet string

    folder string
    password string
    region string
    storageName string
}

func NewGetMergedCellsInRemotedWorksheetRequest(name string, worksheet string, opts ...RequestOption) *GetMergedCellsInRemotedWorksheetRequest {
    req := &GetMergedCellsInRemotedWorksheetRequest{
        name: name,
        worksheet: worksheet,
    }
    if req.name == "" {
        return nil
    }
    if req.worksheet == "" {
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
    if val, ok := cfg.Params["password"].(string); ok {
        req.password = val
    }
    if val, ok := cfg.Params["region"].(string); ok {
        req.region = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *GetMergedCellsInRemotedWorksheetRequest) GetMethod() string {
    return "GET"
}

func (request *GetMergedCellsInRemotedWorksheetRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *GetMergedCellsInRemotedWorksheetRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{worksheet}/mergedcells"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"worksheet"+"}", fmt.Sprintf("%v", request.worksheet), -1)
    return localVarPath
}

func (request *GetMergedCellsInRemotedWorksheetRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    if request.region != "" {
        localVarQueryParams.Add("region", fmt.Sprintf("%v", request.region))
    }
    if request.password != "" {
        localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
    }
    return localVarQueryParams
}

func (request *GetMergedCellsInRemotedWorksheetRequest) GetJSONBody() interface{} {
    return nil
}

func (request *GetMergedCellsInRemotedWorksheetRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *GetMergedCellsInRemotedWorksheetRequest) Description() {
    fmt.Println(strings.Trim("Get all merged cell area form a remote spreadsheet worksheet.", " "))
}