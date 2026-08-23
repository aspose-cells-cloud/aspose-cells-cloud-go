package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PutWorksheetSparklineGroupRequest struct {
    dataRange string
    isVertical bool
    locationRange string
    name string
    sheetName string
    _type string

    folder string
    storageName string
}

func NewPutWorksheetSparklineGroupRequest(dataRange string, isVertical bool, locationRange string, name string, sheetName string, _type string, opts ...RequestOption) *PutWorksheetSparklineGroupRequest {
    req := &PutWorksheetSparklineGroupRequest{
        dataRange: dataRange,
        isVertical: isVertical,
        locationRange: locationRange,
        name: name,
        sheetName: sheetName,
        _type: _type,
    }
    if req.dataRange == "" {
        return nil
    }
    if req.locationRange == "" {
        return nil
    }
    if req.name == "" {
        return nil
    }
    if req.sheetName == "" {
        return nil
    }
    if req._type == "" {
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

func (request *PutWorksheetSparklineGroupRequest) GetMethod() string {
    return "PUT"
}

func (request *PutWorksheetSparklineGroupRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PutWorksheetSparklineGroupRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/sparklineGroups"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PutWorksheetSparklineGroupRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("type", fmt.Sprintf("%v", request._type))
    localVarQueryParams.Add("dataRange", fmt.Sprintf("%v", request.dataRange))
    localVarQueryParams.Add("isVertical", fmt.Sprintf("%v", request.isVertical))
    localVarQueryParams.Add("locationRange", fmt.Sprintf("%v", request.locationRange))
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PutWorksheetSparklineGroupRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PutWorksheetSparklineGroupRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PutWorksheetSparklineGroupRequest) Description() {
    fmt.Println(strings.Trim("Add a sparkline group in the worksheet.", " "))
}