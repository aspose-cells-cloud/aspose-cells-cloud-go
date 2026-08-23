package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type CheckCloudServiceHealthRequest struct {
}

func NewCheckCloudServiceHealthRequest() *CheckCloudServiceHealthRequest {
    req := &CheckCloudServiceHealthRequest{
    }

    return req
}

func (request *CheckCloudServiceHealthRequest) GetMethod() string {
    return "GET"
}

func (request *CheckCloudServiceHealthRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *CheckCloudServiceHealthRequest) GetPath() string {
    localVarPath := "/cells/status/check"
    return localVarPath
}

func (request *CheckCloudServiceHealthRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    return localVarQueryParams
}

func (request *CheckCloudServiceHealthRequest) GetJSONBody() interface{} {
    return nil
}

func (request *CheckCloudServiceHealthRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *CheckCloudServiceHealthRequest) Description() {
    fmt.Println(strings.Trim("Check the Health Status of Aspose.Cells Cloud Service.", " "))
}