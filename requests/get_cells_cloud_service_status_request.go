package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type GetCellsCloudServiceStatusRequest struct {
}

func NewGetCellsCloudServiceStatusRequest() *GetCellsCloudServiceStatusRequest {
    req := &GetCellsCloudServiceStatusRequest{
    }

    return req
}

func (request *GetCellsCloudServiceStatusRequest) GetMethod() string {
    return "GET"
}

func (request *GetCellsCloudServiceStatusRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *GetCellsCloudServiceStatusRequest) GetPath() string {
    localVarPath := "/cells/status/check"
    return localVarPath
}

func (request *GetCellsCloudServiceStatusRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    return localVarQueryParams
}

func (request *GetCellsCloudServiceStatusRequest) GetJSONBody() interface{} {
    return nil
}

func (request *GetCellsCloudServiceStatusRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *GetCellsCloudServiceStatusRequest) Description() {
    fmt.Println(strings.Trim("Aspose.Cells Cloud service health status check.", " "))
}