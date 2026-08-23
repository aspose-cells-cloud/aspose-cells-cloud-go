package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PostRenameWorksheetRequest struct {
    name string
    newname string
    sheetName string

    folder string
    storageName string
}

func NewPostRenameWorksheetRequest(name string, newname string, sheetName string, opts ...RequestOption) *PostRenameWorksheetRequest {
    req := &PostRenameWorksheetRequest{
        name: name,
        newname: newname,
        sheetName: sheetName,
    }
    if req.name == "" {
        return nil
    }
    if req.newname == "" {
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

func (request *PostRenameWorksheetRequest) GetMethod() string {
    return "POST"
}

func (request *PostRenameWorksheetRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostRenameWorksheetRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/rename"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PostRenameWorksheetRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("newname", fmt.Sprintf("%v", request.newname))
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostRenameWorksheetRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostRenameWorksheetRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostRenameWorksheetRequest) Description() {
    fmt.Println(strings.Trim("Rename worksheet in the workbook.", " "))
}