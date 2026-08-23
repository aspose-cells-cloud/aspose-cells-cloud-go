package requests

import (
    "fmt"
    "net/url"
    "path/filepath"
    "strings"
)

type PutWorksheetBackgroundRequest struct {
    name string
    sheetName string

    File string
    FileData []byte
    FileName string
    folder string
    imageAdaptOption string
    picPath string
    storageName string
}

func NewPutWorksheetBackgroundRequest(name string, sheetName string, opts ...RequestOption) *PutWorksheetBackgroundRequest {
    req := &PutWorksheetBackgroundRequest{
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

    if val, ok := cfg.Params["File"].(string); ok {
        req.File = val
    }
    if val, ok := cfg.Params["folder"].(string); ok {
        req.folder = val
    }
    if val, ok := cfg.Params["imageAdaptOption"].(string); ok {
        req.imageAdaptOption = val
    }
    if val, ok := cfg.Params["picPath"].(string); ok {
        req.picPath = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *PutWorksheetBackgroundRequest) SetFileBytes(data []byte, name string) {
    if name == "" {
        name = "File"
    }
    request.FileData = data
    request.FileName = name
}

func (request *PutWorksheetBackgroundRequest) GetMethod() string {
    return "PUT"
}

func (request *PutWorksheetBackgroundRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "multipart/form-data"
    return localVarHeaderParams
}

func (request *PutWorksheetBackgroundRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/background"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PutWorksheetBackgroundRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.picPath != "" {
        localVarQueryParams.Add("picPath", fmt.Sprintf("%v", request.picPath))
    }
    if request.imageAdaptOption != "" {
        localVarQueryParams.Add("imageAdaptOption", fmt.Sprintf("%v", request.imageAdaptOption))
    }
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PutWorksheetBackgroundRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PutWorksheetBackgroundRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    if request.FileData != nil {
        localVarFormParams[request.FileName] = request.FileData
    } else if request.File != "" {
        localVarFormParams["@"+filepath.Base(request.File)] = request.File
    }
    return localVarFormParams
}

func (request *PutWorksheetBackgroundRequest) Description() {
    fmt.Println(strings.Trim("Set background image in the worksheet.", " "))
}