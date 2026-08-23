package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostSplitTextRequest struct {
    splitTextOptions *models.SplitTextOptions
}

func NewPostSplitTextRequest(splitTextOptions *models.SplitTextOptions) *PostSplitTextRequest {
    req := &PostSplitTextRequest{
        splitTextOptions: splitTextOptions,
    }
    if req.splitTextOptions == nil {
        return nil
    }

    return req
}

func (request *PostSplitTextRequest) GetMethod() string {
    return "POST"
}

func (request *PostSplitTextRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostSplitTextRequest) GetPath() string {
    localVarPath := "/cells/splittext"
    return localVarPath
}

func (request *PostSplitTextRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    return localVarQueryParams
}

func (request *PostSplitTextRequest) GetJSONBody() interface{} {
    return &request.splitTextOptions
}

func (request *PostSplitTextRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostSplitTextRequest) Description() {
    fmt.Println(strings.Trim("Efficiently divides Excel cell content into columns or rows based on specified delimiters or patterns. Supports Character-based splitting, Custom string splitting, Mask and wildcard splitting for pattern-based division, Line break division, Column or row splitting, Delimiter removal or retention.", " "))
}