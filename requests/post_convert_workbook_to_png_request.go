package requests

import (
    "fmt"
    "net/url"
    "path/filepath"
    "strings"
)

type PostConvertWorkbookToPNGRequest struct {
    File string
    FileData []byte
    FileName string

    checkExcelRestriction *bool
    password string
    region string
}

func NewPostConvertWorkbookToPNGRequest(File string, opts ...RequestOption) *PostConvertWorkbookToPNGRequest {
    req := &PostConvertWorkbookToPNGRequest{
        File: File,
    }
    cfg := &requestConfig{
        Params: make(map[string]interface{}),
    }
    for _, opt := range opts {
        opt.apply(cfg)
    }

    if val, ok := cfg.Params["checkExcelRestriction"].(*bool); ok {
        req.checkExcelRestriction = val
    }
    if val, ok := cfg.Params["password"].(string); ok {
        req.password = val
    }
    if val, ok := cfg.Params["region"].(string); ok {
        req.region = val
    }

    return req
}

func (request *PostConvertWorkbookToPNGRequest) SetFileBytes(data []byte, name string) {
    if name == "" {
        name = "File"
    }
    request.FileData = data
    request.FileName = name
}

func (request *PostConvertWorkbookToPNGRequest) GetMethod() string {
    return "POST"
}

func (request *PostConvertWorkbookToPNGRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "multipart/form-data"
    return localVarHeaderParams
}

func (request *PostConvertWorkbookToPNGRequest) GetPath() string {
    localVarPath := "/cells/convert/png"
    return localVarPath
}

func (request *PostConvertWorkbookToPNGRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.password != "" {
        localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
    }
    if request.checkExcelRestriction != nil {
        localVarQueryParams.Add("checkExcelRestriction", fmt.Sprintf("%v", *request.checkExcelRestriction))
    }
    if request.region != "" {
        localVarQueryParams.Add("region", fmt.Sprintf("%v", request.region))
    }
    return localVarQueryParams
}

func (request *PostConvertWorkbookToPNGRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostConvertWorkbookToPNGRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    if request.FileData != nil {
        localVarFormParams[request.FileName] = request.FileData
    } else if request.File != "" {
        localVarFormParams["@"+filepath.Base(request.File)] = request.File
    }
    return localVarFormParams
}

func (request *PostConvertWorkbookToPNGRequest) Description() {
    fmt.Println(strings.Trim("Convert Excel file to PNG files.", " "))
}