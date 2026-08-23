package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type DeleteWorksheetSparklineGroupRequest struct {
    name string
    sheetName string
    sparklineIndex int

    folder string
    storageName string
}

func NewDeleteWorksheetSparklineGroupRequest(name string, sheetName string, sparklineIndex int, opts ...RequestOption) *DeleteWorksheetSparklineGroupRequest {
    req := &DeleteWorksheetSparklineGroupRequest{
        name: name,
        sheetName: sheetName,
        sparklineIndex: sparklineIndex,
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

func (request *DeleteWorksheetSparklineGroupRequest) GetMethod() string {
    return "DELETE"
}

func (request *DeleteWorksheetSparklineGroupRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *DeleteWorksheetSparklineGroupRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/sparklineGroups/{sparklineIndex}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sparklineIndex"+"}", fmt.Sprintf("%v", request.sparklineIndex), -1)
    return localVarPath
}

func (request *DeleteWorksheetSparklineGroupRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *DeleteWorksheetSparklineGroupRequest) GetJSONBody() interface{} {
    return nil
}

func (request *DeleteWorksheetSparklineGroupRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *DeleteWorksheetSparklineGroupRequest) Description() {
    fmt.Println(strings.Trim("Delete a sparkline group in the worksheet.", " "))
}