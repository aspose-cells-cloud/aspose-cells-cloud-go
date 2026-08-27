package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostWorksheetSparklineGroupRequest struct {
    name string
    sheetName string
    sparklineGroup *models.SparklineGroup
    sparklineGroupIndex int

    folder string
    storageName string

    extraQueryParameters map[string]string
}

func NewPostWorksheetSparklineGroupRequest(name string, sheetName string, sparklineGroup *models.SparklineGroup, sparklineGroupIndex int, opts ...RequestOption) *PostWorksheetSparklineGroupRequest {
    req := &PostWorksheetSparklineGroupRequest{
        name: name,
        sheetName: sheetName,
        sparklineGroup: sparklineGroup,
        sparklineGroupIndex: sparklineGroupIndex,
    }
    if req.name == "" {
        return nil
    }
    if req.sheetName == "" {
        return nil
    }
    if req.sparklineGroup == nil {
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
    if len(cfg.extraQueryParams) > 0 {
        if req.extraQueryParameters == nil {
            req.extraQueryParameters = make(map[string]string)
        }
        for k, v := range cfg.extraQueryParams {
            req.extraQueryParameters[k] = v
        }
    }

    return req
}

func (request *PostWorksheetSparklineGroupRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *PostWorksheetSparklineGroupRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
}

func (request *PostWorksheetSparklineGroupRequest) GetMethod() string {
    return "POST"
}

func (request *PostWorksheetSparklineGroupRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostWorksheetSparklineGroupRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/sparklineGroups/{sparklineGroupIndex}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sparklineGroupIndex"+"}", fmt.Sprintf("%v", request.sparklineGroupIndex), -1)
    return localVarPath
}

func (request *PostWorksheetSparklineGroupRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    for k, v := range request.extraQueryParameters {
        localVarQueryParams.Add(k, v)
    }
    return localVarQueryParams
}

func (request *PostWorksheetSparklineGroupRequest) GetJSONBody() interface{} {
    return &request.sparklineGroup
}

func (request *PostWorksheetSparklineGroupRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostWorksheetSparklineGroupRequest) Description() {
    fmt.Println(strings.Trim("Update a sparkline group in the worksheet.", " "))
}