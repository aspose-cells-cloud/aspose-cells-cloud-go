/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="put_worksheet_custom_filter_request.go">
*   Copyright (c) 2023 Aspose.Cells Cloud
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

type PutWorksheetCustomFilterRequest struct {
    Name string `json:"name,omitempty" xml:"name"`
    SheetName string `json:"sheet_name,omitempty" xml:"sheet_name"`
	Range_ string `json:"range,omitempty" xml:"range"`
	FieldIndex int64 `json:"field_index,omitempty" xml:"field_index"`
	OperatorType1 string `json:"operator_type1,omitempty" xml:"operator_type1"`
	Criteria1 string `json:"criteria1,omitempty" xml:"criteria1"`
	IsAnd bool `json:"is_and,omitempty" xml:"is_and"`
	OperatorType2 string `json:"operator_type2,omitempty" xml:"operator_type2"`
	Criteria2 string `json:"criteria2,omitempty" xml:"criteria2"`
	MatchBlanks bool `json:"match_blanks,omitempty" xml:"match_blanks"`
	Refresh bool `json:"refresh,omitempty" xml:"refresh"`
	Folder string `json:"folder,omitempty" xml:"folder"`
	StorageName string `json:"storage_name,omitempty" xml:"storage_name"`
	
}

func (data *PutWorksheetCustomFilterRequest) CreateRequestData( client *APIClient) (localVarRequest *http.Request, err error) {
	var (
		localVarHttpMethod  = strings.ToUpper("PUT")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
	)

	// create path and map variables
	localVarPath := client.cfg.BasePath + "/" + client.cfg.Version + "/cells/{name}/worksheets/{sheetName}/autoFilter/custom"
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

    // query params : operatorType1
    if data.OperatorType1 != "" {
        localVarQueryParams.Add("operatorType1", parameterToString(data.OperatorType1, ""))
    }

    // query params : criteria1
    if data.Criteria1 != "" {
        localVarQueryParams.Add("criteria1", parameterToString(data.Criteria1, ""))
    }

    // query params : isAnd
    if &data.IsAnd != nil {
        localVarQueryParams.Add("isAnd", parameterToString(data.IsAnd, ""))
    }

    // query params : operatorType2
    if data.OperatorType2 != "" {
        localVarQueryParams.Add("operatorType2", parameterToString(data.OperatorType2, ""))
    }

    // query params : criteria2
    if data.Criteria2 != "" {
        localVarQueryParams.Add("criteria2", parameterToString(data.Criteria2, ""))
    }

    // query params : matchBlanks
    if &data.MatchBlanks != nil {
        localVarQueryParams.Add("matchBlanks", parameterToString(data.MatchBlanks, ""))
    }

    // query params : refresh
    if &data.Refresh != nil {
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
