package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostDataFillRequest struct {
    dataFillRequest *models.DataFillRequest
}

func NewPostDataFillRequest(dataFillRequest *models.DataFillRequest) *PostDataFillRequest {
    req := &PostDataFillRequest{
        dataFillRequest: dataFillRequest,
    }
    if req.dataFillRequest == nil {
        return nil
    }

    return req
}

func (request *PostDataFillRequest) GetMethod() string {
    return "POST"
}

func (request *PostDataFillRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostDataFillRequest) GetPath() string {
    localVarPath := "/cells/datafill"
    return localVarPath
}

func (request *PostDataFillRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    return localVarQueryParams
}

func (request *PostDataFillRequest) GetJSONBody() interface{} {
    return &request.dataFillRequest
}

func (request *PostDataFillRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostDataFillRequest) Description() {
    fmt.Println(strings.Trim("Data filling for spreadsheet files is primarily used to fill empty data in tables and ranges.", " "))
}