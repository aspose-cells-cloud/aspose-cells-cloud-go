package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostPageSetupRequest struct {
    name string
    pageSetup *models.PageSetup
    sheetName string

    folder string
    storageName string
}

func NewPostPageSetupRequest(name string, pageSetup *models.PageSetup, sheetName string, opts ...RequestOption) *PostPageSetupRequest {
    req := &PostPageSetupRequest{
        name: name,
        pageSetup: pageSetup,
        sheetName: sheetName,
    }
    if req.name == "" {
        return nil
    }
    if req.pageSetup == nil {
        return nil
    }
    if req.sheetName == "" {
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

func (request *PostPageSetupRequest) GetMethod() string {
    return "POST"
}

func (request *PostPageSetupRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostPageSetupRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/pagesetup"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    return localVarPath
}

func (request *PostPageSetupRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostPageSetupRequest) GetJSONBody() interface{} {
    return &request.pageSetup
}

func (request *PostPageSetupRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostPageSetupRequest) Description() {
    fmt.Println(strings.Trim("Update page setup in the worksheet.", " "))
}