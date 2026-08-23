package requests

import (
    "fmt"
    "net/url"
    "path/filepath"
    "strings"
)

type PostWatermarkRequest struct {
    color string
    File string
    FileData []byte
    FileName string
    text string

    checkExcelRestriction *bool
    outFormat string
    password string
    region string
}

func NewPostWatermarkRequest(color string, File string, text string, opts ...RequestOption) *PostWatermarkRequest {
    req := &PostWatermarkRequest{
        color: color,
        File: File,
        text: text,
    }
    if req.color == "" {
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

func (request *PostWatermarkRequest) SetFileBytes(data []byte, name string) {
    if name == "" {
        name = "File"
    }
    request.FileData = data
    request.FileName = name
}

func (request *PostWatermarkRequest) GetMethod() string {
    return "POST"
}

func (request *PostWatermarkRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "multipart/form-data"
    return localVarHeaderParams
}

func (request *PostWatermarkRequest) GetPath() string {
    localVarPath := "/cells/watermark"
    return localVarPath
}

func (request *PostWatermarkRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("text", fmt.Sprintf("%v", request.text))
    localVarQueryParams.Add("color", fmt.Sprintf("%v", request.color))
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

func (request *PostWatermarkRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostWatermarkRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    if request.FileData != nil {
        localVarFormParams[request.FileName] = request.FileData
    } else if request.File != "" {
        localVarFormParams["@"+filepath.Base(request.File)] = request.File
    }
    return localVarFormParams
}

func (request *PostWatermarkRequest) Description() {
    fmt.Println(strings.Trim("Add Text Watermark to Excel files and generate output files in various formats.", " "))
}