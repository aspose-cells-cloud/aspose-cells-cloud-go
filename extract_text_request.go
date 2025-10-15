/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="extract_text_request.go">
*   Copyright (c) 2025 Aspose.Cells Cloud
* </copyright>
* <summary>
*   Permission is hereby granted, free of charge, to any person obtaining a copy
*  of this software and associated documentation files (the "Software"), to deal
*  in the Software without restriction, including without limitation the rights
*  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
*  copies of the Software, and to permit persons to whom the Software is
*  furnished to do so, subject to the following conditions:
* 
*  The above copyright notice and this permission notice shall be included in all
*  copies or substantial portions of the Software.
* 
*  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
*  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
*  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
*  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
*  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
*  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
*  SOFTWARE.
* </summary> 
-------------------------------------------------------------------------------------------------------------------- **/

package asposecellscloud

import (
	"path/filepath"
	"net/url"
	"net/http"
	"strings"
)

type ExtractTextRequest struct {
	ExtractTextType string `json:"extract_text_type,omitempty" xml:"extract_text_type"`
	BeforeText string `json:"before_text,omitempty" xml:"before_text"`
	AfterText string `json:"after_text,omitempty" xml:"after_text"`
	BeforePosition int64 `json:"before_position,omitempty" xml:"before_position"`
	AfterPosition int64 `json:"after_position,omitempty" xml:"after_position"`
	OutPositionRange string `json:"out_position_range,omitempty" xml:"out_position_range"`
	Worksheet string `json:"worksheet,omitempty" xml:"worksheet"`
	Range_ string `json:"range,omitempty" xml:"range"`
	OutPath string `json:"out_path,omitempty" xml:"out_path"`
	OutStorageName string `json:"out_storage_name,omitempty" xml:"out_storage_name"`
	Region string `json:"region,omitempty" xml:"region"`
	Password string `json:"password,omitempty" xml:"password"`
	
	Spreadsheet string  `json:"spreadsheet,omitempty" xml:"spreadsheet"`
	 

	ExtendQueryParameterMap	map[string]string `json:"ExtendQueryParameterMap,omitempty" xml:"ExtendQueryParameterMap"`	
}

func (data *ExtractTextRequest) CreateRequestData( client *APIClient) (localVarRequest *http.Request, err error) {
	var (
		localVarHttpMethod  = strings.ToUpper("PUT")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
	)

	// create path and map variables
	localVarPath := client.cfg.BasePath + "/v4.0/cells/content/extract/text"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

    // query params : extractTextType
    if data.ExtractTextType != "" {
        localVarQueryParams.Add("extractTextType", parameterToString(data.ExtractTextType, ""))
    }
    // query params : beforeText
    if data.BeforeText != "" {
        localVarQueryParams.Add("beforeText", parameterToString(data.BeforeText, ""))
    }
    // query params : afterText
    if data.AfterText != "" {
        localVarQueryParams.Add("afterText", parameterToString(data.AfterText, ""))
    }
    // query params : beforePosition
    if &data.BeforePosition != nil {
        localVarQueryParams.Add("beforePosition", parameterToString(data.BeforePosition, ""))
    }
    // query params : afterPosition
    if &data.AfterPosition != nil {
        localVarQueryParams.Add("afterPosition", parameterToString(data.AfterPosition, ""))
    }
    // query params : outPositionRange
    if data.OutPositionRange != "" {
        localVarQueryParams.Add("outPositionRange", parameterToString(data.OutPositionRange, ""))
    }
    // query params : worksheet
    if data.Worksheet != "" {
        localVarQueryParams.Add("worksheet", parameterToString(data.Worksheet, ""))
    }
    // query params : range
    if data.Range_ != "" {
        localVarQueryParams.Add("range", parameterToString(data.Range_, ""))
    }
    // query params : outPath
    if data.OutPath != "" {
        localVarQueryParams.Add("outPath", parameterToString(data.OutPath, ""))
    }
    // query params : outStorageName
    if data.OutStorageName != "" {
        localVarQueryParams.Add("outStorageName", parameterToString(data.OutStorageName, ""))
    }
    // query params : region
    if data.Region != "" {
        localVarQueryParams.Add("region", parameterToString(data.Region, ""))
    }
    // query params : password
    if data.Password != "" {
        localVarQueryParams.Add("password", parameterToString(data.Password, ""))
    }
	if data.ExtendQueryParameterMap != nil {
		for key, value := range data.ExtendQueryParameterMap {
			localVarQueryParams.Add(key, parameterToString(value, ""))
		}
	}
	localVarHttpContentTypes := []string{"multipart/form-data"} 

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
			
		if strings.TrimSpace(data.Spreadsheet) != "" { localVarFormParams["@"+ filepath.Base(data.Spreadsheet)] = []string {data.Spreadsheet}} 
	r, err := client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	return r,err
}