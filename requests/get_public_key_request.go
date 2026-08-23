package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type GetPublicKeyRequest struct {
}

func NewGetPublicKeyRequest() *GetPublicKeyRequest {
    req := &GetPublicKeyRequest{
    }

    return req
}

func (request *GetPublicKeyRequest) GetMethod() string {
    return "GET"
}

func (request *GetPublicKeyRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *GetPublicKeyRequest) GetPath() string {
    localVarPath := "/cells/publickey"
    return localVarPath
}

func (request *GetPublicKeyRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    return localVarQueryParams
}

func (request *GetPublicKeyRequest) GetJSONBody() interface{} {
    return nil
}

func (request *GetPublicKeyRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *GetPublicKeyRequest) Description() {
    fmt.Println(strings.Trim("Get an asymmetric public key.", " "))
}