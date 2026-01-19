/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="put_worksheet_filter_top10_request.go">
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

type PutWorksheetFilterTop10Request struct {
    Name string `json:"name,omitempty" xml:"name"`
    SheetName string `json:"sheet_name,omitempty" xml:"sheet_name"`
	Range_ string `json:"range,omitempty" xml:"range"`
	FieldIndex int64 `json:"field_index,omitempty" xml:"field_index"`
	IsTop bool `json:"is_top,omitempty" xml:"is_top"`
	IsPercent bool `json:"is_percent,omitempty" xml:"is_percent"`
	ItemCount int64 `json:"item_count,omitempty" xml:"item_count"`
	MatchBlanks bool `json:"match_blanks,omitempty" xml:"match_blanks"`
	Refresh bool `json:"refresh,omitempty" xml:"refresh"`
	Folder string `json:"folder,omitempty" xml:"folder"`
	StorageName string `json:"storage_name,omitempty" xml:"storage_name"`
	

	ExtendQueryParameterMap	map[string]string `json:"ExtendQueryParameterMap,omitempty" xml:"ExtendQueryParameterMap"`	
}

func (data *PutWorksheetFilterTop10Request) CreateRequestData( client *APIClient) (localVarRequest *http.Request, err error) {
	var (
		localVarHttpMethod  = strings.ToUpper("PUT")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
	)

	// create path and map variables
	localVarPath := client.cfg.BasePath + "/v3.0/cells/{name}/worksheets/{sheetName}/autoFilter/filterTop10"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", data.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", data.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

    // query params : range
    if data.Range_ != "" {
        localVarQueryParams.Add("range", parameterToString(data.Range_, ""))
    }
    // query params : fieldIndex
    if &data.FieldIndex != nil {
        localVarQueryParams.Add("fieldIndex", parameterToString(data.FieldIndex, ""))
    }
    // query params : isTop
    if data.IsTop {
        localVarQueryParams.Add("isTop", parameterToString(data.IsTop, ""))
    }
    // query params : isPercent
    if data.IsPercent {
        localVarQueryParams.Add("isPercent", parameterToString(data.IsPercent, ""))
    }
    // query params : itemCount
    if &data.ItemCount != nil {
        localVarQueryParams.Add("itemCount", parameterToString(data.ItemCount, ""))
    }
    // query params : matchBlanks
    if data.MatchBlanks {
        localVarQueryParams.Add("matchBlanks", parameterToString(data.MatchBlanks, ""))
    }
    // query params : refresh
    if data.Refresh {
        localVarQueryParams.Add("refresh", parameterToString(data.Refresh, ""))
    }
    // query params : folder
    if data.Folder != "" {
        localVarQueryParams.Add("folder", parameterToString(data.Folder, ""))
    }
    // query params : storageName
    if data.StorageName != "" {
        localVarQueryParams.Add("storageName", parameterToString(data.StorageName, ""))
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