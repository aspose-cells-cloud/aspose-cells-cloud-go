package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostConvertTextRequest struct {
    convertTextOptions *models.ConvertTextOptions
}

func NewPostConvertTextRequest(convertTextOptions *models.ConvertTextOptions) *PostConvertTextRequest {
    req := &PostConvertTextRequest{
        convertTextOptions: convertTextOptions,
    }
    if req.convertTextOptions == nil {
        return nil
    }

    return req
}

func (request *PostConvertTextRequest) GetMethod() string {
    return "POST"
}

func (request *PostConvertTextRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostConvertTextRequest) GetPath() string {
    localVarPath := "/cells/converttext"
    return localVarPath
}

func (request *PostConvertTextRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    return localVarQueryParams
}

func (request *PostConvertTextRequest) GetJSONBody() interface{} {
    return &request.convertTextOptions
}

func (request *PostConvertTextRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostConvertTextRequest) Description() {
    fmt.Println(strings.Trim("Enhance Excel data through essential text conversions: convert text to numbers, replace characters and line breaks, and remove accents.", " "))
}