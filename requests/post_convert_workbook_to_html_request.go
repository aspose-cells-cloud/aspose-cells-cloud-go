package requests

import (
    "fmt"
    "net/url"
    "path/filepath"
    "strings"
)

type PostConvertWorkbookToHtmlRequest struct {
    File string
    FileData []byte
    FileName string

    checkExcelRestriction *bool
    FontsLocation string
    password string
    region string
}

func NewPostConvertWorkbookToHtmlRequest(File string, opts ...RequestOption) *PostConvertWorkbookToHtmlRequest {
    req := &PostConvertWorkbookToHtmlRequest{
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
    if val, ok := cfg.Params["FontsLocation"].(string); ok {
        req.FontsLocation = val
    }
    if val, ok := cfg.Params["password"].(string); ok {
        req.password = val
    }
    if val, ok := cfg.Params["region"].(string); ok {
        req.region = val
    }

    return req
}

func (request *PostConvertWorkbookToHtmlRequest) SetFileBytes(data []byte, name string) {
    if name == "" {
        name = "File"
    }
    request.FileData = data
    request.FileName = name
}

func (request *PostConvertWorkbookToHtmlRequest) GetMethod() string {
    return "POST"
}

func (request *PostConvertWorkbookToHtmlRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "multipart/form-data"
    return localVarHeaderParams
}

func (request *PostConvertWorkbookToHtmlRequest) GetPath() string {
    localVarPath := "/cells/convert/html"
    return localVarPath
}

func (request *PostConvertWorkbookToHtmlRequest) GetQueryParameters() url.Values {
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
    if request.FontsLocation != "" {
        localVarQueryParams.Add("FontsLocation", fmt.Sprintf("%v", request.FontsLocation))
    }
    return localVarQueryParams
}

func (request *PostConvertWorkbookToHtmlRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostConvertWorkbookToHtmlRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    if request.FileData != nil {
        localVarFormParams[request.FileName] = request.FileData
    } else if request.File != "" {
        localVarFormParams["@"+filepath.Base(request.File)] = request.File
    }
    return localVarFormParams
}

func (request *PostConvertWorkbookToHtmlRequest) Description() {
    fmt.Println(strings.Trim("Convert Excel file to HTML files.", " "))
}