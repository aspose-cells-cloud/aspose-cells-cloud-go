package requests

import (
    "fmt"
    "net/url"
    "path/filepath"
    "strings"
)

type PostImportRequest struct {
    File string
    FileData []byte
    FileName string

    checkExcelRestriction *bool
    outFormat string
    password string
    region string
}

func NewPostImportRequest(File string, opts ...RequestOption) *PostImportRequest {
    req := &PostImportRequest{
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
    if val, ok := cfg.Params["outFormat"].(string); ok {
        req.outFormat = val
    }
    if val, ok := cfg.Params["password"].(string); ok {
        req.password = val
    }
    if val, ok := cfg.Params["region"].(string); ok {
        req.region = val
    }

    return req
}

func (request *PostImportRequest) SetFileBytes(data []byte, name string) {
    if name == "" {
        name = "File"
    }
    request.FileData = data
    request.FileName = name
}

func (request *PostImportRequest) GetMethod() string {
    return "POST"
}

func (request *PostImportRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "multipart/form-data"
    return localVarHeaderParams
}

func (request *PostImportRequest) GetPath() string {
    localVarPath := "/cells/import"
    return localVarPath
}

func (request *PostImportRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.outFormat != "" {
        localVarQueryParams.Add("outFormat", fmt.Sprintf("%v", request.outFormat))
    }
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

func (request *PostImportRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostImportRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    if request.FileData != nil {
        localVarFormParams[request.FileName] = request.FileData
    } else if request.File != "" {
        localVarFormParams["@"+filepath.Base(request.File)] = request.File
    }
    return localVarFormParams
}

func (request *PostImportRequest) Description() {
    fmt.Println(strings.Trim("Import data into an Excel file and generate output files in various formats.", " "))
}