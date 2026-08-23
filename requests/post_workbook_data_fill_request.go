package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostWorkbookDataFillRequest struct {
    dataFill *models.DataFill
    name string

    checkExcelRestriction *bool
    folder string
    password string
    region string
    storageName string
}

func NewPostWorkbookDataFillRequest(dataFill *models.DataFill, name string, opts ...RequestOption) *PostWorkbookDataFillRequest {
    req := &PostWorkbookDataFillRequest{
        dataFill: dataFill,
        name: name,
    }
    if req.dataFill == nil {
        return nil
    }
    if req.name == "" {
        return nil
    }

    cfg := &requestConfig{
        Params: make(map[string]interface{}),
    }
    for _, opt := range opts {
        opt.apply(cfg)
    }

    if val, ok := cfg.Params["checkExcelRestriction"].(*bool); ok {
        req.checkExcelRestriction = val
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

func (request *PostWorkbookDataFillRequest) GetMethod() string {
    return "POST"
}

func (request *PostWorkbookDataFillRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostWorkbookDataFillRequest) GetPath() string {
    localVarPath := "/cells/{name}/datafill"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    return localVarPath
}

func (request *PostWorkbookDataFillRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    if request.password != "" {
        localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
    }
    if request.region != "" {
        localVarQueryParams.Add("region", fmt.Sprintf("%v", request.region))
    }
    if request.checkExcelRestriction != nil {
        localVarQueryParams.Add("checkExcelRestriction", fmt.Sprintf("%v", *request.checkExcelRestriction))
    }
    return localVarQueryParams
}

func (request *PostWorkbookDataFillRequest) GetJSONBody() interface{} {
    return &request.dataFill
}

func (request *PostWorkbookDataFillRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostWorkbookDataFillRequest) Description() {
    fmt.Println(strings.Trim("Data filling for spreadsheet files is primarily used to fill empty data in tables and ranges.", " "))
}