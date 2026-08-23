package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type GetAsposeCellsCloudStatusRequest struct {
}

func NewGetAsposeCellsCloudStatusRequest() *GetAsposeCellsCloudStatusRequest {
    req := &GetAsposeCellsCloudStatusRequest{
    }

    return req
}

func (request *GetAsposeCellsCloudStatusRequest) GetMethod() string {
    return "GET"
}

func (request *GetAsposeCellsCloudStatusRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *GetAsposeCellsCloudStatusRequest) GetPath() string {
    localVarPath := "/cells"
    return localVarPath
}

func (request *GetAsposeCellsCloudStatusRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    return localVarQueryParams
}

func (request *GetAsposeCellsCloudStatusRequest) GetJSONBody() interface{} {
    return nil
}

func (request *GetAsposeCellsCloudStatusRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *GetAsposeCellsCloudStatusRequest) Description() {
    fmt.Println(strings.Trim("Check the Health Status of Aspose.Cells Cloud Service.", " "))
}