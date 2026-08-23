package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type DeleteWorksheetHyperlinkRequest struct {
    hyperlinkIndex int
    name string
    sheetName string

    folder string
    storageName string
}

func NewDeleteWorksheetHyperlinkRequest(hyperlinkIndex int, name string, sheetName string, opts ...RequestOption) *DeleteWorksheetHyperlinkRequest {
    req := &DeleteWorksheetHyperlinkRequest{
        hyperlinkIndex: hyperlinkIndex,
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

func (request *DeleteWorksheetHyperlinkRequest) GetMethod() string {
    return "DELETE"
}

func (request *DeleteWorksheetHyperlinkRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *DeleteWorksheetHyperlinkRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/hyperlinks/{hyperlinkIndex}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"hyperlinkIndex"+"}", fmt.Sprintf("%v", request.hyperlinkIndex), -1)
    return localVarPath
}

func (request *DeleteWorksheetHyperlinkRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *DeleteWorksheetHyperlinkRequest) GetJSONBody() interface{} {
    return nil
}

func (request *DeleteWorksheetHyperlinkRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *DeleteWorksheetHyperlinkRequest) Description() {
    fmt.Println(strings.Trim("Delete hyperlink by index in the worksheet.", " "))
}