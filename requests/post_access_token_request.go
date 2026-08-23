package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PostAccessTokenRequest struct {
}

func NewPostAccessTokenRequest() *PostAccessTokenRequest {
    req := &PostAccessTokenRequest{
    }

    return req
}

func (request *PostAccessTokenRequest) GetMethod() string {
    return "POST"
}

func (request *PostAccessTokenRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostAccessTokenRequest) GetPath() string {
    localVarPath := "/cells/connect/token"
    return localVarPath
}

func (request *PostAccessTokenRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    return localVarQueryParams
}

func (request *PostAccessTokenRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostAccessTokenRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostAccessTokenRequest) Description() {
    fmt.Println(strings.Trim("Get Access Token Result: The Cells Cloud Get Token API acts as a proxy service,", " "))
}