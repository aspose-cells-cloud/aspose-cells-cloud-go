package requests

import (
    "fmt"
    "net/url"
    "strings"
)

type ConvertTextInRemoteSpreadsheetRequest struct {
    convertTextType string
    name string
    _range string
    worksheet string

    folder string
    password string
    region string
    sourceCharacters string
    storageName string
    targetCharacters string
}

func NewConvertTextInRemoteSpreadsheetRequest(convertTextType string, name string, _range string, worksheet string, opts ...RequestOption) *ConvertTextInRemoteSpreadsheetRequest {
    req := &ConvertTextInRemoteSpreadsheetRequest{
        convertTextType: convertTextType,
        name: name,
        _range: _range,
        worksheet: worksheet,
    }
    if req.convertTextType == "" {
        return nil
    }
    if req.name == "" {
        return nil
    }
    if req._range == "" {
        return nil
    }
    if req.worksheet == "" {
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
    if val, ok := cfg.Params["password"].(string); ok {
        req.password = val
    }
    if val, ok := cfg.Params["region"].(string); ok {
        req.region = val
    }
    if val, ok := cfg.Params["sourceCharacters"].(string); ok {
        req.sourceCharacters = val
    }
    if val, ok := cfg.Params["storageName"].(string); ok {
        req.storageName = val
    }
    if val, ok := cfg.Params["targetCharacters"].(string); ok {
        req.targetCharacters = val
    }

    return req
}

func (request *ConvertTextInRemoteSpreadsheetRequest) GetMethod() string {
    return "PUT"
}

func (request *ConvertTextInRemoteSpreadsheetRequest) GetHeaderParameters() map[string]string {
    localVarHeaderParams := make(map[string]string)
    localVarHeaderParams["Content-Type"] = "application/json"
    return localVarHeaderParams
}

func (request *ConvertTextInRemoteSpreadsheetRequest) GetPath() string {
    localVarPath := "/cells/{name}/worksheets/{worksheet}/range/{range}/content/convert/text"
    localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", request.name), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"worksheet"+"}", fmt.Sprintf("%v", request.worksheet), -1)
    localVarPath = strings.Replace(localVarPath, "{"+"range"+"}", fmt.Sprintf("%v", request._range), -1)
    return localVarPath
}

func (request *ConvertTextInRemoteSpreadsheetRequest) GetQueryParameters() url.Values {
    localVarQueryParams := url.Values{}
    localVarQueryParams.Add("convertTextType", fmt.Sprintf("%v", request.convertTextType))
    if request.sourceCharacters != "" {
        localVarQueryParams.Add("sourceCharacters", fmt.Sprintf("%v", request.sourceCharacters))
    }
    if request.targetCharacters != "" {
        localVarQueryParams.Add("targetCharacters", fmt.Sprintf("%v", request.targetCharacters))
    }
    if request.folder != "" {
        localVarQueryParams.Add("folder", fmt.Sprintf("%v", request.folder))
    }
    if request.storageName != "" {
        localVarQueryParams.Add("storageName", fmt.Sprintf("%v", request.storageName))
    }
    if request.region != "" {
        localVarQueryParams.Add("region", fmt.Sprintf("%v", request.region))
    }
    if request.password != "" {
        localVarQueryParams.Add("password", fmt.Sprintf("%v", request.password))
    }
    return localVarQueryParams
}

func (request *ConvertTextInRemoteSpreadsheetRequest) GetJSONBody() interface{} {
    return nil
}

func (request *ConvertTextInRemoteSpreadsheetRequest) GetMultipartForm() map[string]interface{} {
    localVarFormParams := make(map[string]interface{})
    return localVarFormParams
}

func (request *ConvertTextInRemoteSpreadsheetRequest) Description() {
    fmt.Println(strings.Trim("Indicates converting the numbers stored as text into the correct number format, replacing unwanted characters and line breaks with the desired characters, and converting accented characters to their equivalent characters without accents.", " "))
}