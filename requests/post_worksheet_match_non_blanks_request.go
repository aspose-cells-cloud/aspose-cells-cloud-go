package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PostWorksheetMatchNonBlanksRequest struct {
    fieldIndex int
    name string
    sheetName string

    folder string
    storageName string
}

func NewPostWorksheetMatchNonBlanksRequest(fieldIndex int, name string, sheetName string, opts ...RequestOption) *PostWorksheetMatchNonBlanksRequest {
    req := &PostWorksheetMatchNonBlanksRequest{
        fieldIndex: fieldIndex,
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

func (request *PostWorksheetMatchNonBlanksRequest) GetMethod() string {
    return "POST"
}

func (request *PostWorksheetMatchNonBlanksRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostWorksheetMatchNonBlanksRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/autoFilter/matchNonBlanks"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PostWorksheetMatchNonBlanksRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("fieldIndex", fmt.Sprintf("%v", request.fieldIndex))
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostWorksheetMatchNonBlanksRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostWorksheetMatchNonBlanksRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostWorksheetMatchNonBlanksRequest) Description() {
    fmt.Println(strings.Trim("Match all not blank cells in the list.", " "))
}