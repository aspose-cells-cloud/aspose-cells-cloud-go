package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostAnalyzeExcelRequest struct {
    analyzeExcelRequest *models.AnalyzeExcelRequest
}

func NewPostAnalyzeExcelRequest(analyzeExcelRequest *models.AnalyzeExcelRequest) *PostAnalyzeExcelRequest {
    req := &PostAnalyzeExcelRequest{
        analyzeExcelRequest: analyzeExcelRequest,
    }
    if req.analyzeExcelRequest == nil {
        return nil
    }

    return req
}

func (request *PostAnalyzeExcelRequest) GetMethod() string {
    return "POST"
}

func (request *PostAnalyzeExcelRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostAnalyzeExcelRequest) GetPath() string {
    localVarPath := "/cells/analyze"
    return localVarPath
}

func (request *PostAnalyzeExcelRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    return localVarQueryParams
}

func (request *PostAnalyzeExcelRequest) GetJSONBody() interface{} {
    return &request.analyzeExcelRequest
}

func (request *PostAnalyzeExcelRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostAnalyzeExcelRequest) Description() {
    fmt.Println(strings.Trim("Perform business analysis of data in Excel files.", " "))
}