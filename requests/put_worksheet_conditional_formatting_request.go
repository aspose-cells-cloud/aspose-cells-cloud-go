package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PutWorksheetConditionalFormattingRequest struct {
    cellArea string
    formatcondition *models.FormatCondition
    name string
    sheetName string

    folder string
    storageName string
}

func NewPutWorksheetConditionalFormattingRequest(cellArea string, formatcondition *models.FormatCondition, name string, sheetName string, opts ...RequestOption) *PutWorksheetConditionalFormattingRequest {
    req := &PutWorksheetConditionalFormattingRequest{
        cellArea: cellArea,
        formatcondition: formatcondition,
        name: name,
        sheetName: sheetName,
    }
    if req.cellArea == "" {
        return nil
    }
    if req.formatcondition == nil {
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

func (request *PutWorksheetConditionalFormattingRequest) GetMethod() string {
    return "PUT"
}

func (request *PutWorksheetConditionalFormattingRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PutWorksheetConditionalFormattingRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/conditionalFormattings"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PutWorksheetConditionalFormattingRequest) GetQueryParameters() url.Values {
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

func (request *PutWorksheetConditionalFormattingRequest) GetJSONBody() interface{} {
    return &request.formatcondition
}

func (request *PutWorksheetConditionalFormattingRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PutWorksheetConditionalFormattingRequest) Description() {
    fmt.Println(strings.Trim("Add conditional formatting in the worksheet.", " "))
}