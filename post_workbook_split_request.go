/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="post_workbook_split_request.go">
*   Copyright (c) 2026 Aspose.Cells Cloud
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
	"fmt"
	"net/url"
	"net/http"
	"strings"
)

type PostWorkbookSplitRequest struct {
    Name string `json:"name,omitempty" xml:"name"`
	Format string `json:"format,omitempty" xml:"format"`
	OutFolder string `json:"out_folder,omitempty" xml:"out_folder"`
	From int64 `json:"from,omitempty" xml:"from"`
	To int64 `json:"to,omitempty" xml:"to"`
	HorizontalResolution int64 `json:"horizontal_resolution,omitempty" xml:"horizontal_resolution"`
	VerticalResolution int64 `json:"vertical_resolution,omitempty" xml:"vertical_resolution"`
	SplitNameRule string `json:"split_name_rule,omitempty" xml:"split_name_rule"`
	Folder string `json:"folder,omitempty" xml:"folder"`
	StorageName string `json:"storage_name,omitempty" xml:"storage_name"`
	OutStorageName string `json:"out_storage_name,omitempty" xml:"out_storage_name"`
	

	ExtendQueryParameterMap	map[string]string `json:"ExtendQueryParameterMap,omitempty" xml:"ExtendQueryParameterMap"`	
}

func (data *PostWorkbookSplitRequest) CreateRequestData( client *APIClient) (localVarRequest *http.Request, err error) {
	var (
		localVarHttpMethod  = strings.ToUpper("POST")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
	)

	// create path and map variables
	localVarPath := client.cfg.BasePath + "/v3.0/cells/{name}/split"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", data.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

    // query params : format
    if data.Format != "" {
        localVarQueryParams.Add("format", parameterToString(data.Format, ""))
    }
    // query params : outFolder
    if data.OutFolder != "" {
        localVarQueryParams.Add("outFolder", parameterToString(data.OutFolder, ""))
    }
    // query params : from
    if &data.From != nil {
        localVarQueryParams.Add("from", parameterToString(data.From, ""))
    }
    // query params : to
    if &data.To != nil {
        localVarQueryParams.Add("to", parameterToString(data.To, ""))
    }
    // query params : horizontalResolution
    if &data.HorizontalResolution != nil {
        localVarQueryParams.Add("horizontalResolution", parameterToString(data.HorizontalResolution, ""))
    }
    // query params : verticalResolution
    if &data.VerticalResolution != nil {
        localVarQueryParams.Add("verticalResolution", parameterToString(data.VerticalResolution, ""))
    }
    // query params : splitNameRule
    if data.SplitNameRule != "" {
        localVarQueryParams.Add("splitNameRule", parameterToString(data.SplitNameRule, ""))
    }
    // query params : folder
    if data.Folder != "" {
        localVarQueryParams.Add("folder", parameterToString(data.Folder, ""))
    }
    // query params : storageName
    if data.StorageName != "" {
        localVarQueryParams.Add("storageName", parameterToString(data.StorageName, ""))
    }
    // query params : outStorageName
    if data.OutStorageName != "" {
        localVarQueryParams.Add("outStorageName", parameterToString(data.OutStorageName, ""))
    }
	if data.ExtendQueryParameterMap != nil {
		for key, value := range data.ExtendQueryParameterMap {
			localVarQueryParams.Add(key, parameterToString(value, ""))
		}
	}
	localVarHttpContentTypes := []string{"application/json"} 

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
	r, err := client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	return r,err
}