package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type MergeSpreadsheetsInRemoteFolderRequest struct {
    folder string

    fileMatchExpression string
    fontsLocation string
    mergeInOneSheet *bool
    outFormat string
    outPath string
    outStorageName string
    password string
    region string
    storageName string
}

func NewMergeSpreadsheetsInRemoteFolderRequest(folder string, opts ...RequestOption) *MergeSpreadsheetsInRemoteFolderRequest {
    req := &MergeSpreadsheetsInRemoteFolderRequest{
        folder: folder,
    }
    if req.folder == "" {
        return nil
    }

    cfg := &requestConfig{
        Params: make(map[string]interface{}),
    }
    for _, opt := range opts {
        opt.apply(cfg)
    }

    if val, ok := cfg.Params["fileMatchExpression"].(string); ok {
        req.fileMatchExpression = val
    }
    if val, ok := cfg.Params["fontsLocation"].(string); ok {
        req.fontsLocation = val
    }
    if val, ok := cfg.Params["mergeInOneSheet"].(*bool); ok {
        req.mergeInOneSheet = val
    }
    if val, ok := cfg.Params["outFormat"].(string); ok {
        req.outFormat = val
    }
    if val, ok := cfg.Params["outPath"].(string); ok {
        req.outPath = val
    }
    if val, ok := cfg.Params["outStorageName"].(string); ok {
        req.outStorageName = val
    }
    if val, ok := cfg.Params["password"].(string); ok {
        req.password = val
    }
    if val, ok := cfg.Params["region"].(string); ok {
        req.region = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *MergeSpreadsheetsInRemoteFolderRequest) GetMethod() string {
    return "PUT"
}

func (request *MergeSpreadsheetsInRemoteFolderRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *MergeSpreadsheetsInRemoteFolderRequest) GetPath() string {
    localVarPath := "/cells/merge/remote-spreadsheets"
    return localVarPath
}

func (request *MergeSpreadsheetsInRemoteFolderRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    if request.fileMatchExpression != "" {
        localVarQueryParams.Add("fileMatchExpression", fmt.Sprintf("%v", request.fileMatchExpression))
    }
    if request.outFormat != "" {
        localVarQueryParams.Add("outFormat", fmt.Sprintf("%v", request.outFormat))
    }
    if request.mergeInOneSheet != nil {
        localVarQueryParams.Add("mergeInOneSheet", fmt.Sprintf("%v", *request.mergeInOneSheet))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    if request.outPath != "" {
        localVarQueryParams.Add("outPath", fmt.Sprintf("%v", request.outPath))
    }
    if request.outStorageName != "" {
        localVarQueryParams.Add("outStorageName", fmt.Sprintf("%v", request.outStorageName))
    }
    if request.fontsLocation != "" {
        localVarQueryParams.Add("fontsLocation", fmt.Sprintf("%v", request.fontsLocation))
    }
    if request.region != "" {
        localVarQueryParams.Add("region", fmt.Sprintf("%v", request.region))
    }
    if request.password != "" {
        localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
    }
    return localVarQueryParams
}

func (request *MergeSpreadsheetsInRemoteFolderRequest) GetJSONBody() interface{} {
    return nil
}

func (request *MergeSpreadsheetsInRemoteFolderRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *MergeSpreadsheetsInRemoteFolderRequest) Description() {
    fmt.Println(strings.Trim("Merge spreadsheet files in folder of cloud storage into a specified format file.", " "))
}