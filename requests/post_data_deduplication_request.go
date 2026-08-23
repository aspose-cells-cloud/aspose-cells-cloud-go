package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostDataDeduplicationRequest struct {
    dataDeduplicationRequest *models.DataDeduplicationRequest
}

func NewPostDataDeduplicationRequest(dataDeduplicationRequest *models.DataDeduplicationRequest) *PostDataDeduplicationRequest {
    req := &PostDataDeduplicationRequest{
        dataDeduplicationRequest: dataDeduplicationRequest,
    }
    if req.dataDeduplicationRequest == nil {
        return nil
    }

    return req
}

func (request *PostDataDeduplicationRequest) GetMethod() string {
    return "POST"
}

func (request *PostDataDeduplicationRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostDataDeduplicationRequest) GetPath() string {
    localVarPath := "/cells/datadeduplication"
    return localVarPath
}

func (request *PostDataDeduplicationRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    return localVarQueryParams
}

func (request *PostDataDeduplicationRequest) GetJSONBody() interface{} {
    return &request.dataDeduplicationRequest
}

func (request *PostDataDeduplicationRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostDataDeduplicationRequest) Description() {
    fmt.Println(strings.Trim("Data deduplication of spreadsheet files is mainly used to eliminate duplicate data in tables and ranges.", " "))
}