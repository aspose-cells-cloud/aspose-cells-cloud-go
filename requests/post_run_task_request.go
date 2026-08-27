package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostRunTaskRequest struct {
    TaskData *models.TaskData

    extraQueryParameters map[string]string
}

func NewPostRunTaskRequest(TaskData *models.TaskData, opts ...RequestOption) *PostRunTaskRequest {
    req := &PostRunTaskRequest{
        TaskData: TaskData,
    }
    if req.TaskData == nil {
        return nil
    }

    cfg := &requestConfig{
        Params: make(map[string]interface{}),
    }
    for _, opt := range opts {
        opt.apply(cfg)
    }

    if len(cfg.extraQueryParams) > 0 {
        if req.extraQueryParameters == nil {
            req.extraQueryParameters = make(map[string]string)
        }
        for k, v := range cfg.extraQueryParams {
            req.extraQueryParameters[k] = v
        }
    }

    return req
}

func (request *PostRunTaskRequest) AddQueryParameter(key, value string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    request.extraQueryParameters[key] = value
}

func (request *PostRunTaskRequest) AddQueryParameters(params map[string]string) {
    if request.extraQueryParameters == nil {
        request.extraQueryParameters = make(map[string]string)
    }
    for k, v := range params {
        request.extraQueryParameters[k] = v
    }
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
    for k, v := range request.extraQueryParameters {
        localVarQueryParams.Add(k, v)
    }
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