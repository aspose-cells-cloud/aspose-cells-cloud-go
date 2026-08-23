package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type GetCellsCloudServicesHealthCheckRequest struct {
}

func NewGetCellsCloudServicesHealthCheckRequest() *GetCellsCloudServicesHealthCheckRequest {
    req := &GetCellsCloudServicesHealthCheckRequest{
    }

    return req
}

func (request *GetCellsCloudServicesHealthCheckRequest) GetMethod() string {
    return "GET"
}

func (request *GetCellsCloudServicesHealthCheckRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *GetCellsCloudServicesHealthCheckRequest) GetPath() string {
    localVarPath := "/cells"
    return localVarPath
}

func (request *GetCellsCloudServicesHealthCheckRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    return localVarQueryParams
}

func (request *GetCellsCloudServicesHealthCheckRequest) GetJSONBody() interface{} {
    return nil
}

func (request *GetCellsCloudServicesHealthCheckRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *GetCellsCloudServicesHealthCheckRequest) Description() {
    fmt.Println(strings.Trim("Retrieve cell descriptions in a specified format.", " "))
}