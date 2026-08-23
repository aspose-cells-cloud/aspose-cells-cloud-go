package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostExtractTextRequest struct {
    extractTextOptions *models.ExtractTextOptions
}

func NewPostExtractTextRequest(extractTextOptions *models.ExtractTextOptions) *PostExtractTextRequest {
    req := &PostExtractTextRequest{
        extractTextOptions: extractTextOptions,
    }
    if req.extractTextOptions == nil {
        return nil
    }

    return req
}

func (request *PostExtractTextRequest) GetMethod() string {
    return "POST"
}

func (request *PostExtractTextRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostExtractTextRequest) GetPath() string {
    localVarPath := "/cells/extracttext"
    return localVarPath
}

func (request *PostExtractTextRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    return localVarQueryParams
}

func (request *PostExtractTextRequest) GetJSONBody() interface{} {
    return &request.extractTextOptions
}

func (request *PostExtractTextRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostExtractTextRequest) Description() {
    fmt.Println(strings.Trim("Effortlessly extract text and numbers from Excel cells with precise options. This API allows extraction of first/last characters, text between delimiters, and numbers from strings, with output as static values or formulas.", " "))
}