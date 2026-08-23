package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type StorageExistsRequest struct {
    storageName string
}

func NewStorageExistsRequest(storageName string) *StorageExistsRequest {
    req := &StorageExistsRequest{
        storageName: storageName,
    }
    if req.storageName == "" {
        return nil
    }

    return req
}

func (request *StorageExistsRequest) GetMethod() string {
    return "GET"
}

func (request *StorageExistsRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *StorageExistsRequest) GetPath() string {
    localVarPath := "/cells/storage/{storageName}/exist"
    localVarPath = strings.Replace(localVarPath, "{"+"storageName"+"}", fmt.Sprintf("%v", request.storageName), -1)
    return localVarPath
}

func (request *StorageExistsRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    return localVarQueryParams
}

func (request *StorageExistsRequest) GetJSONBody() interface{} {
    return nil
}

func (request *StorageExistsRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *StorageExistsRequest) Description() {
    fmt.Println(strings.Trim("StorageExists", " "))
}