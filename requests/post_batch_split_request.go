package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostBatchSplitRequest struct {
    batchSplitRequest *models.BatchSplitRequest
}

func NewPostBatchSplitRequest(batchSplitRequest *models.BatchSplitRequest) *PostBatchSplitRequest {
    req := &PostBatchSplitRequest{
        batchSplitRequest: batchSplitRequest,
    }
    if req.batchSplitRequest == nil {
        return nil
    }

    return req
}

func (request *PostBatchSplitRequest) GetMethod() string {
    return "POST"
}

func (request *PostBatchSplitRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostBatchSplitRequest) GetPath() string {
    localVarPath := "/cells/batch/split"
    return localVarPath
}

func (request *PostBatchSplitRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    return localVarQueryParams
}

func (request *PostBatchSplitRequest) GetJSONBody() interface{} {
    return &request.batchSplitRequest
}

func (request *PostBatchSplitRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostBatchSplitRequest) Description() {
    fmt.Println(strings.Trim("Batch splitting files that meet specific matching conditions.", " "))
}