package requests

import (
    "fmt"
    "net/url"
    "path/filepath"
    "strings"
)

type PostUnlockRequest struct {
    File string
    FileData []byte
    FileName string
    password string
}

func NewPostUnlockRequest(File string, password string) *PostUnlockRequest {
    req := &PostUnlockRequest{
        File: File,
        password: password,
    }
    if req.password == "" {
        return nil
    }

    return req
}

func (request *PostUnlockRequest) SetFileBytes(data []byte, name string) {
    if name == "" {
        name = "File"
    }
    request.FileData = data
    request.FileName = name
}

func (request *PostUnlockRequest) GetMethod() string {
    return "POST"
}

func (request *PostUnlockRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "multipart/form-data"
    return localVarHeaderParams
}

func (request *PostUnlockRequest) GetPath() string {
    localVarPath := "/cells/unlock"
    return localVarPath
}

func (request *PostUnlockRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
    return localVarQueryParams
}

func (request *PostUnlockRequest) GetJSONBody() interface{} {
    return nil
}

func (request *PostUnlockRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    if request.FileData != nil {
        localVarFormParams[request.FileName] = request.FileData
    } else if request.File != "" {
        localVarFormParams["@"+filepath.Base(request.File)] = request.File
    }
    return localVarFormParams
}

func (request *PostUnlockRequest) Description() {
    fmt.Println(strings.Trim("Unlock Excel files.", " "))
}