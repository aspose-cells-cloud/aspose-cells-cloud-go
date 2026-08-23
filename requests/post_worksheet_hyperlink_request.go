package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostWorksheetHyperlinkRequest struct {
    hyperlink *models.Hyperlink
    hyperlinkIndex int
    name string
    sheetName string

    folder string
    storageName string
}

func NewPostWorksheetHyperlinkRequest(hyperlink *models.Hyperlink, hyperlinkIndex int, name string, sheetName string, opts ...RequestOption) *PostWorksheetHyperlinkRequest {
    req := &PostWorksheetHyperlinkRequest{
        hyperlink: hyperlink,
        hyperlinkIndex: hyperlinkIndex,
        name: name,
        sheetName: sheetName,
    }
    if req.hyperlink == nil {
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

func (request *PostWorksheetHyperlinkRequest) GetMethod() string {
    return "POST"
}

func (request *PostWorksheetHyperlinkRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostWorksheetHyperlinkRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/hyperlinks/{hyperlinkIndex}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"hyperlinkIndex"+"}", fmt.Sprintf("%v", request.hyperlinkIndex), -1)
    return localVarPath
}

func (request *PostWorksheetHyperlinkRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostWorksheetHyperlinkRequest) GetJSONBody() interface{} {
    return &request.hyperlink
}

func (request *PostWorksheetHyperlinkRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostWorksheetHyperlinkRequest) Description() {
    fmt.Println(strings.Trim("Update hyperlink by index in the worksheet.", " "))
}