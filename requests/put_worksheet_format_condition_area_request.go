package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PutWorksheetFormatConditionAreaRequest struct {
    cellArea string
    index int
    name string
    sheetName string

    folder string
    storageName string
}

func NewPutWorksheetFormatConditionAreaRequest(cellArea string, index int, name string, sheetName string, opts ...RequestOption) *PutWorksheetFormatConditionAreaRequest {
    req := &PutWorksheetFormatConditionAreaRequest{
        cellArea: cellArea,
        index: index,
        name: name,
        sheetName: sheetName,
    }
    if req.cellArea == "" {
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

func (request *PutWorksheetFormatConditionAreaRequest) GetMethod() string {
    return "PUT"
}

func (request *PutWorksheetFormatConditionAreaRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PutWorksheetFormatConditionAreaRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/conditionalFormattings/{index}/area"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", request.index), -1)
    return localVarPath
}

func (request *PutWorksheetFormatConditionAreaRequest) GetQueryParameters() url.Values {
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

func (request *PutWorksheetFormatConditionAreaRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PutWorksheetFormatConditionAreaRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PutWorksheetFormatConditionAreaRequest) Description() {
    fmt.Println(strings.Trim("Add a cell area for the format condition in the worksheet.", " "))
}