package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type DeleteWorksheetConditionalFormattingsRequest struct {
    name string
    sheetName string

    folder string
    storageName string
}

func NewDeleteWorksheetConditionalFormattingsRequest(name string, sheetName string, opts ...RequestOption) *DeleteWorksheetConditionalFormattingsRequest {
    req := &DeleteWorksheetConditionalFormattingsRequest{
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

func (request *DeleteWorksheetConditionalFormattingsRequest) GetMethod() string {
    return "DELETE"
}

func (request *DeleteWorksheetConditionalFormattingsRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *DeleteWorksheetConditionalFormattingsRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/conditionalFormattings"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *DeleteWorksheetConditionalFormattingsRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *DeleteWorksheetConditionalFormattingsRequest) GetJSONBody() interface{} {
    return nil
}

func (request *DeleteWorksheetConditionalFormattingsRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *DeleteWorksheetConditionalFormattingsRequest) Description() {
    fmt.Println(strings.Trim("Clear all conditional formattings in the worksheet.", " "))
}