package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostUpdateWordCaseRequest struct {
    wordCaseOptions *models.WordCaseOptions
}

func NewPostUpdateWordCaseRequest(wordCaseOptions *models.WordCaseOptions) *PostUpdateWordCaseRequest {
    req := &PostUpdateWordCaseRequest{
        wordCaseOptions: wordCaseOptions,
    }
    if req.wordCaseOptions == nil {
        return nil
    }

    return req
}

func (request *PostUpdateWordCaseRequest) GetMethod() string {
    return "POST"
}

func (request *PostUpdateWordCaseRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostUpdateWordCaseRequest) GetPath() string {
    localVarPath := "/cells/updatewordcase"
    return localVarPath
}

func (request *PostUpdateWordCaseRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    return localVarQueryParams
}

func (request *PostUpdateWordCaseRequest) GetJSONBody() interface{} {
    return &request.wordCaseOptions
}

func (request *PostUpdateWordCaseRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostUpdateWordCaseRequest) Description() {
    fmt.Println(strings.Trim("Managing inconsistent text case in spreadsheets (Excel, Google Sheets, CSV) can be frustrating, especially with large datasets. The PostUpdateWordCase WEB API solves this by automating text case conversions, ensuring clean and standardized data.", " "))
}