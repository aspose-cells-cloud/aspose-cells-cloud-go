package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PutWorksheetHyperlinkRequest struct {
    address string
    firstColumn int
    firstRow int
    name string
    sheetName string
    totalColumns int
    totalRows int

    folder string
    storageName string
}

func NewPutWorksheetHyperlinkRequest(address string, firstColumn int, firstRow int, name string, sheetName string, totalColumns int, totalRows int, opts ...RequestOption) *PutWorksheetHyperlinkRequest {
    req := &PutWorksheetHyperlinkRequest{
        address: address,
        firstColumn: firstColumn,
        firstRow: firstRow,
        name: name,
        sheetName: sheetName,
        totalColumns: totalColumns,
        totalRows: totalRows,
    }
    if req.address == "" {
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

func (request *PutWorksheetHyperlinkRequest) GetMethod() string {
    return "PUT"
}

func (request *PutWorksheetHyperlinkRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PutWorksheetHyperlinkRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/hyperlinks"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PutWorksheetHyperlinkRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("firstRow", fmt.Sprintf("%v", request.firstRow))
    localVarQueryParams.Add("firstColumn", fmt.Sprintf("%v", request.firstColumn))
    localVarQueryParams.Add("totalRows", fmt.Sprintf("%v", request.totalRows))
    localVarQueryParams.Add("totalColumns", fmt.Sprintf("%v", request.totalColumns))
    localVarQueryParams.Add("address", fmt.Sprintf("%v", request.address))
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PutWorksheetHyperlinkRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PutWorksheetHyperlinkRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PutWorksheetHyperlinkRequest) Description() {
    fmt.Println(strings.Trim("Add hyperlink in the worksheet.", " "))
}