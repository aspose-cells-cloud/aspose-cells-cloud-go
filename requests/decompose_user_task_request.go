package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type DecomposeUserTaskRequest struct {
    TaskDescription string

    password string
    region string
}

func NewDecomposeUserTaskRequest(TaskDescription string, opts ...RequestOption) *DecomposeUserTaskRequest {
    req := &DecomposeUserTaskRequest{
        TaskDescription: TaskDescription,
    }
    if req.TaskDescription == "" {
        return nil
    }

    cfg := &requestConfig{
        Params: make(map[string]interface{}),
    }
    for _, opt := range opts {
        opt.apply(cfg)
    }

    if val, ok := cfg.Params["password"].(string); ok {
        req.password = val
    }
    if val, ok := cfg.Params["region"].(string); ok {
        req.region = val
    }

    return req
}

func (request *DecomposeUserTaskRequest) GetMethod() string {
    return "PUT"
}

func (request *DecomposeUserTaskRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *DecomposeUserTaskRequest) GetPath() string {
    localVarPath := "/cells/ai/task/decompose"
    return localVarPath
}

func (request *DecomposeUserTaskRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.region != "" {
        localVarQueryParams.Add("region", fmt.Sprintf("%v", request.region))
    }
    if request.password != "" {
        localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
    }
    return localVarQueryParams
}

func (request *DecomposeUserTaskRequest) GetJSONBody() interface{} {
    return &request.TaskDescription
}

func (request *DecomposeUserTaskRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *DecomposeUserTaskRequest) Description() {
    fmt.Println(strings.Trim("AI task decomposition: Convert user objectives to sequential action plans with formatted file export.", " "))
}