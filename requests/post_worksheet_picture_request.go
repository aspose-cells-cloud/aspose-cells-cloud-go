package requests

import (
    "fmt"
    "net/url"
    "strings"

    "asposecellscloud/models"
)

type PostWorksheetPictureRequest struct {
    name string
    picture *models.Picture
    pictureIndex int
    sheetName string

    folder string
    storageName string
}

func NewPostWorksheetPictureRequest(name string, picture *models.Picture, pictureIndex int, sheetName string, opts ...RequestOption) *PostWorksheetPictureRequest {
    req := &PostWorksheetPictureRequest{
        name: name,
        picture: picture,
        pictureIndex: pictureIndex,
        sheetName: sheetName,
    }
    if req.name == "" {
        return nil
    }
    if req.picture == nil {
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

func (request *PostWorksheetPictureRequest) GetMethod() string {
    return "POST"
}

func (request *PostWorksheetPictureRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *PostWorksheetPictureRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{sheetName}/pictures/{pictureIndex}"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", request.sheetName), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"pictureIndex"+"}", fmt.Sprintf("%v", request.pictureIndex), -1)
    return localVarPath
}

func (request *PostWorksheetPictureRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    return localVarQueryParams
}

func (request *PostWorksheetPictureRequest) GetJSONBody() interface{} {
    return &request.picture
}

func (request *PostWorksheetPictureRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *PostWorksheetPictureRequest) Description() {
    fmt.Println(strings.Trim("Update a picture by index in the worksheet.", " "))
}