package requests

import (
    "fmt"
    "net/url"
    "path/filepath"
    "strings"
)

type PostLockRequest struct {
    File string
    FileData []byte
    FileName string
    password string
}

func NewPostLockRequest(File string, password string) *PostLockRequest {
    req := &PostLockRequest{
        File: File,
        password: password,
    }
    if req.password == "" {
        return nil
    }

    return req
}

func (request *PostLockRequest) SetFileBytes(data []byte, name string) {
    if name == "" {
        name = "File"
    }
    request.FileData = data
    request.FileName = name
}

func (request *PostLockRequest) GetMethod() string {
    return "POST"
}

func (request *PostLockRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "multipart/form-data"
    return localVarHeaderParams
}

func (request *PostLockRequest) GetPath() string {
    localVarPath := "/cells/lock"
    return localVarPath
}

func (request *PostLockRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
    return localVarQueryParams
}

func (request *PostLockRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostLockRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    if request.FileData != nil {
        localVarFormParams[request.FileName] = request.FileData
    } else if request.File != "" {
        localVarFormParams["@"+filepath.Base(request.File)] = request.File
    }
    return localVarFormParams
}

func (request *PostLockRequest) Description() {
    fmt.Println(strings.Trim("Lock Excel files.", " "))
}