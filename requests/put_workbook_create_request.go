package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PutWorkbookCreateRequest struct {
    name string

    checkExcelRestriction *bool
    dataFile string
    folder string
    isWriteOver *bool
    storageName string
    templateFile string
}

func NewPutWorkbookCreateRequest(name string, opts ...RequestOption) *PutWorkbookCreateRequest {
    req := &PutWorkbookCreateRequest{
        name: name,
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
    if val, ok := cfg.Params["dataFile"].(string); ok {
        req.dataFile = val
    }
    if val, ok := cfg.Params["folder"].(string); ok {
        req.folder = val
    }
    if val, ok := cfg.Params["isWriteOver"].(*bool); ok {
        req.isWriteOver = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }
    if val, ok := cfg.Params["templateFile"].(string); ok {
        req.templateFile = val
    }

    return req
}

func (request *PutWorkbookCreateRequest) GetMethod() string {
    return "PUT"
}

func (request *PutWorkbookCreateRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PutWorkbookCreateRequest) GetPath() string {
    localVarPath := "/cells/{name}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    return localVarPath
}

func (request *PutWorkbookCreateRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.templateFile != "" {
        localVarQueryParams.Add("templateFile", fmt.Sprintf("%v", request.templateFile))
    }
    if request.dataFile != "" {
        localVarQueryParams.Add("dataFile", fmt.Sprintf("%v", request.dataFile))
    }
    if request.isWriteOver != nil {
        localVarQueryParams.Add("isWriteOver", fmt.Sprintf("%v", *request.isWriteOver))
    }
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    if request.checkExcelRestriction != nil {
        localVarQueryParams.Add("checkExcelRestriction", fmt.Sprintf("%v", *request.checkExcelRestriction))
    }
    return localVarQueryParams
}

func (request *PutWorkbookCreateRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PutWorkbookCreateRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PutWorkbookCreateRequest) Description() {
    fmt.Println(strings.Trim("Create a new workbook using different methods.", " "))
}