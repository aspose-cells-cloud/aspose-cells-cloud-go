package requests

import (
    "fmt"
    "net/url"
    "path/filepath"
    "strings"

    "asposecellscloud/models"
)

type PostMetadataRequest struct {
    cellsDocuments []models.CellsDocumentProperty
    File string
    FileData []byte
    FileName string

    checkExcelRestriction *bool
    outFormat string
    password string
    region string
}

func NewPostMetadataRequest(cellsDocuments []models.CellsDocumentProperty, File string, opts ...RequestOption) *PostMetadataRequest {
    req := &PostMetadataRequest{
        cellsDocuments: cellsDocuments,
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

func (request *PostMetadataRequest) SetFileBytes(data []byte, name string) {
    if name == "" {
        name = "File"
    }
    request.FileData = data
    request.FileName = name
}

func (request *PostMetadataRequest) GetMethod() string {
    return "POST"
}

func (request *PostMetadataRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "multipart/form-data"
    return localVarHeaderParams
}

func (request *PostMetadataRequest) GetPath() string {
    localVarPath := "/cells/metadata/update"
    return localVarPath
}

func (request *PostMetadataRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.password != "" {
        localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
    }
    if request.checkExcelRestriction != nil {
        localVarQueryParams.Add("checkExcelRestriction", fmt.Sprintf("%v", *request.checkExcelRestriction))
    }
    if request.outFormat != "" {
        localVarQueryParams.Add("outFormat", fmt.Sprintf("%v", request.outFormat))
    }
    if request.region != "" {
        localVarQueryParams.Add("region", fmt.Sprintf("%v", request.region))
    }
    return localVarQueryParams
}

func (request *PostMetadataRequest) GetJSONBody() interface{} {
    return request.cellsDocuments
}

func (request *PostMetadataRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    if request.FileData != nil {
        localVarFormParams[request.FileName] = request.FileData
    } else if request.File != "" {
        localVarFormParams["@"+filepath.Base(request.File)] = request.File
    }
    return localVarFormParams
}

func (request *PostMetadataRequest) Description() {
    fmt.Println(strings.Trim("Update document properties in Excel file, and save them is various formats.", " "))
}