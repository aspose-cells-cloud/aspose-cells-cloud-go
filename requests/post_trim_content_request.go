package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostTrimContentRequest struct {
    trimContentOptions *models.TrimContentOptions
}

func NewPostTrimContentRequest(trimContentOptions *models.TrimContentOptions) *PostTrimContentRequest {
    req := &PostTrimContentRequest{
        trimContentOptions: trimContentOptions,
    }
    if req.trimContentOptions == nil {
        return nil
    }

    return req
}

func (request *PostTrimContentRequest) GetMethod() string {
    return "POST"
}

func (request *PostTrimContentRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostTrimContentRequest) GetPath() string {
    localVarPath := "/cells/trimcontent"
    return localVarPath
}

func (request *PostTrimContentRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    return localVarQueryParams
}

func (request *PostTrimContentRequest) GetJSONBody() interface{} {
    return &request.trimContentOptions
}

func (request *PostTrimContentRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostTrimContentRequest) Description() {
    fmt.Println(strings.Trim("The PostTrimContent API is designed to process and trim content within a specified range in a spreadsheet. This API allows users to remove extra spaces, line breaks, or other unnecessary characters from the content of selected cells. It is particularly useful for cleaning up data entries and ensuring consistency in spreadsheet formatting", " "))
}