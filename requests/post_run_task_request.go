package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostRunTaskRequest struct {
    TaskData *models.TaskData
}

func NewPostRunTaskRequest(TaskData *models.TaskData) *PostRunTaskRequest {
    req := &PostRunTaskRequest{
        TaskData: TaskData,
    }
    if req.TaskData == nil {
        return nil
    }

    return req
}

func (request *PostRunTaskRequest) GetMethod() string {
    return "POST"
}

func (request *PostRunTaskRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostRunTaskRequest) GetPath() string {
    localVarPath := "/cells/task/runtask"
    return localVarPath
}

func (request *PostRunTaskRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    return localVarQueryParams
}

func (request *PostRunTaskRequest) GetJSONBody() interface{} {
    return &request.TaskData
}

func (request *PostRunTaskRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostRunTaskRequest) Description() {
    fmt.Println(strings.Trim("Run tasks.", " "))
}