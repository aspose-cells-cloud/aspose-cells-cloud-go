package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostDataCleansingRequest struct {
    dataCleansingRequest *models.DataCleansingRequest
}

func NewPostDataCleansingRequest(dataCleansingRequest *models.DataCleansingRequest) *PostDataCleansingRequest {
    req := &PostDataCleansingRequest{
        dataCleansingRequest: dataCleansingRequest,
    }
    if req.dataCleansingRequest == nil {
        return nil
    }

    return req
}

func (request *PostDataCleansingRequest) GetMethod() string {
    return "POST"
}

func (request *PostDataCleansingRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostDataCleansingRequest) GetPath() string {
    localVarPath := "/cells/datacleansing"
    return localVarPath
}

func (request *PostDataCleansingRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    return localVarQueryParams
}

func (request *PostDataCleansingRequest) GetJSONBody() interface{} {
    return &request.dataCleansingRequest
}

func (request *PostDataCleansingRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostDataCleansingRequest) Description() {
    fmt.Println(strings.Trim("Data cleansing of spreadsheet files is a data management process used to identify, correct, and remove errors, incompleteness, duplicates, or inaccuracies in tables and ranges.", " "))
}