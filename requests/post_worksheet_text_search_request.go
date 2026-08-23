package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PostWorksheetTextSearchRequest struct {
    name string
    sheetName string
    text string

    folder string
    storageName string
}

func NewPostWorksheetTextSearchRequest(name string, sheetName string, text string, opts ...RequestOption) *PostWorksheetTextSearchRequest {
    req := &PostWorksheetTextSearchRequest{
        name: name,
        sheetName: sheetName,
        text: text,
    }
    if req.name == "" {
        return nil
    }
    if req.sheetName == "" {
        return nil
    }
    if req.text == "" {
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

func (request *PostWorksheetTextSearchRequest) GetMethod() string {
    return "POST"
}

func (request *PostWorksheetTextSearchRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostWorksheetTextSearchRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/findText"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PostWorksheetTextSearchRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("text", fmt.Sprintf("%v", request.text))
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostWorksheetTextSearchRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostWorksheetTextSearchRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostWorksheetTextSearchRequest) Description() {
    fmt.Println(strings.Trim("Search for text in the worksheet.", " "))
}