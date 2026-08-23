package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type PostDigitalSignatureRequest struct {
    digitalsignaturefile string
    name string
    password string

    folder string
    storageName string
}

func NewPostDigitalSignatureRequest(digitalsignaturefile string, name string, password string, opts ...RequestOption) *PostDigitalSignatureRequest {
    req := &PostDigitalSignatureRequest{
        digitalsignaturefile: digitalsignaturefile,
        name: name,
        password: password,
    }
    if req.digitalsignaturefile == "" {
        return nil
    }
    if req.name == "" {
        return nil
    }
    if req.password == "" {
        return nil
    }

    cfg := &requestConfig{
        Params: make(map[string]interface{}),
    }
    for _, opt := range opts {
        opt.apply(cfg)
    }

    if val, ok := cfg.Params["folder"].(string); ok {
        req.folder = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }

    return req
}

func (request *PostDigitalSignatureRequest) GetMethod() string {
    return "POST"
}

func (request *PostDigitalSignatureRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostDigitalSignatureRequest) GetPath() string {
    localVarPath := "/cells/{name}/digitalsignature"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    return localVarPath
}

func (request *PostDigitalSignatureRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("digitalsignaturefile", fmt.Sprintf("%v", request.digitalsignaturefile))
    localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostDigitalSignatureRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostDigitalSignatureRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostDigitalSignatureRequest) Description() {
    fmt.Println(strings.Trim("Excel file digital signature.", " "))
}