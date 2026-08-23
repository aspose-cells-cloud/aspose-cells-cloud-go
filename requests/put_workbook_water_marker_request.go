package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PutWorkbookWaterMarkerRequest struct {
    name string
    textWaterMarkerRequest *models.TextWaterMarkerRequest

    folder string
    storageName string
}

func NewPutWorkbookWaterMarkerRequest(name string, textWaterMarkerRequest *models.TextWaterMarkerRequest, opts ...RequestOption) *PutWorkbookWaterMarkerRequest {
    req := &PutWorkbookWaterMarkerRequest{
        name: name,
        textWaterMarkerRequest: textWaterMarkerRequest,
    }
    if req.name == "" {
        return nil
    }
    if req.textWaterMarkerRequest == nil {
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

func (request *PutWorkbookWaterMarkerRequest) GetMethod() string {
    return "PUT"
}

func (request *PutWorkbookWaterMarkerRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PutWorkbookWaterMarkerRequest) GetPath() string {
    localVarPath := "/cells/{name}/watermarker"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    return localVarPath
}

func (request *PutWorkbookWaterMarkerRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PutWorkbookWaterMarkerRequest) GetJSONBody() interface{} {
    return &request.textWaterMarkerRequest
}

func (request *PutWorkbookWaterMarkerRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PutWorkbookWaterMarkerRequest) Description() {
    fmt.Println(strings.Trim("Set water marker in the workbook.", " "))
}