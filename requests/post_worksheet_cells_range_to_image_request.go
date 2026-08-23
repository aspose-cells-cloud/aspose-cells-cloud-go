package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostWorksheetCellsRangeToImageRequest struct {
    name string
    rangeConvertRequest *models.RangeConvertRequest
    sheetName string

    folder string
    storageName string
}

func NewPostWorksheetCellsRangeToImageRequest(name string, rangeConvertRequest *models.RangeConvertRequest, sheetName string, opts ...RequestOption) *PostWorksheetCellsRangeToImageRequest {
    req := &PostWorksheetCellsRangeToImageRequest{
        name: name,
        rangeConvertRequest: rangeConvertRequest,
        sheetName: sheetName,
    }
    if req.name == "" {
        return nil
    }
    if req.rangeConvertRequest == nil {
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

func (request *PostWorksheetCellsRangeToImageRequest) GetMethod() string {
    return "POST"
}

func (request *PostWorksheetCellsRangeToImageRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostWorksheetCellsRangeToImageRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/ranges/convertToImage"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PostWorksheetCellsRangeToImageRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostWorksheetCellsRangeToImageRequest) GetJSONBody() interface{} {
    return &request.rangeConvertRequest
}

func (request *PostWorksheetCellsRangeToImageRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostWorksheetCellsRangeToImageRequest) Description() {
    fmt.Println(strings.Trim("PostWorksheetCellsRangeToImage", " "))
}