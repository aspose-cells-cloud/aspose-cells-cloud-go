package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostBatchProtectRequest struct {
    batchProtectRequest *models.BatchProtectRequest
}

func NewPostBatchProtectRequest(batchProtectRequest *models.BatchProtectRequest) *PostBatchProtectRequest {
    req := &PostBatchProtectRequest{
        batchProtectRequest: batchProtectRequest,
    }
    if req.batchProtectRequest == nil {
        return nil
    }

    return req
}

func (request *PostBatchProtectRequest) GetMethod() string {
    return "POST"
}

func (request *PostBatchProtectRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostBatchProtectRequest) GetPath() string {
    localVarPath := "/cells/batch/protect"
    return localVarPath
}

func (request *PostBatchProtectRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    return localVarQueryParams
}

func (request *PostBatchProtectRequest) GetJSONBody() interface{} {
    return &request.batchProtectRequest
}

func (request *PostBatchProtectRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostBatchProtectRequest) Description() {
    fmt.Println(strings.Trim("Batch protecting files that meet specific matching conditions.", " "))
}