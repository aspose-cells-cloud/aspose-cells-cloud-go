package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PostSetWorksheetColumnWidthRequest struct {
    columnIndex int
    name string
    sheetName string
    width float64

    count *int
    folder string
    storageName string
}

func NewPostSetWorksheetColumnWidthRequest(columnIndex int, name string, sheetName string, width float64, opts ...RequestOption) *PostSetWorksheetColumnWidthRequest {
    req := &PostSetWorksheetColumnWidthRequest{
        columnIndex: columnIndex,
        name: name,
        sheetName: sheetName,
        width: width,
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

    if val, ok := cfg.Params["count"].(*int); ok {
        req.count = val
    }
    if val, ok := cfg.Params["folder"].(string); ok {
        req.folder = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *PostSetWorksheetColumnWidthRequest) GetMethod() string {
    return "POST"
}

func (request *PostSetWorksheetColumnWidthRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostSetWorksheetColumnWidthRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/cells/columns/{columnIndex}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"columnIndex"+"}", fmt.Sprintf("%v", request.columnIndex), -1)
    return localVarPath
}

func (request *PostSetWorksheetColumnWidthRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("width", fmt.Sprintf("%v", request.width))
    if request.count != nil {
        localVarQueryParams.Add("count", fmt.Sprintf("%v", *request.count))
    }
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostSetWorksheetColumnWidthRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostSetWorksheetColumnWidthRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostSetWorksheetColumnWidthRequest) Description() {
    fmt.Println(strings.Trim("Set worksheet column width.", " "))
}