package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostBatchConvertRequest struct {
    batchConvertRequest *models.BatchConvertRequest
}

func NewPostBatchConvertRequest(batchConvertRequest *models.BatchConvertRequest) *PostBatchConvertRequest {
    req := &PostBatchConvertRequest{
        batchConvertRequest: batchConvertRequest,
    }
    if req.batchConvertRequest == nil {
        return nil
    }

    return req
}

func (request *PostBatchConvertRequest) GetMethod() string {
    return "POST"
}

func (request *PostBatchConvertRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostBatchConvertRequest) GetPath() string {
    localVarPath := "/cells/batch/convert"
    return localVarPath
}

func (request *PostBatchConvertRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    return localVarQueryParams
}

func (request *PostBatchConvertRequest) GetJSONBody() interface{} {
    return &request.batchConvertRequest
}

func (request *PostBatchConvertRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostBatchConvertRequest) Description() {
    fmt.Println(strings.Trim("Batch converting files that meet specific matching conditions.", " "))
}