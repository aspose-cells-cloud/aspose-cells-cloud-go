/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="put_worksheet_ole_object_request.go">
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

type PutWorksheetOleObjectRequest struct {
    Name string `json:"name,omitempty" xml:"name"`
    SheetName string `json:"sheet_name,omitempty" xml:"sheet_name"`
	UpperLeftRow int64 `json:"upper_left_row,omitempty" xml:"upper_left_row"`
	UpperLeftColumn int64 `json:"upper_left_column,omitempty" xml:"upper_left_column"`
	Height int64 `json:"height,omitempty" xml:"height"`
	Width int64 `json:"width,omitempty" xml:"width"`
	OleFile string `json:"ole_file,omitempty" xml:"ole_file"`
	ImageFile string `json:"image_file,omitempty" xml:"image_file"`
	Folder string `json:"folder,omitempty" xml:"folder"`
	StorageName string `json:"storage_name,omitempty" xml:"storage_name"`
	

	ExtendQueryParameterMap	map[string]string `json:"ExtendQueryParameterMap,omitempty" xml:"ExtendQueryParameterMap"`	
}

func (data *PutWorksheetOleObjectRequest) CreateRequestData( client *APIClient) (localVarRequest *http.Request, err error) {
	var (
		localVarHttpMethod  = strings.ToUpper("PUT")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
	)

	// create path and map variables
	localVarPath := client.cfg.BasePath + "/v3.0/cells/{name}/worksheets/{sheetName}/oleobjects"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", data.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", data.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

    // query params : upperLeftRow
    if &data.UpperLeftRow != nil {
        localVarQueryParams.Add("upperLeftRow", parameterToString(data.UpperLeftRow, ""))
    }
    // query params : upperLeftColumn
    if &data.UpperLeftColumn != nil {
        localVarQueryParams.Add("upperLeftColumn", parameterToString(data.UpperLeftColumn, ""))
    }
    // query params : height
    if &data.Height != nil {
        localVarQueryParams.Add("height", parameterToString(data.Height, ""))
    }
    // query params : width
    if &data.Width != nil {
        localVarQueryParams.Add("width", parameterToString(data.Width, ""))
    }
    // query params : oleFile
    if data.OleFile != "" {
        localVarQueryParams.Add("oleFile", parameterToString(data.OleFile, ""))
    }
    // query params : imageFile
    if data.ImageFile != "" {
        localVarQueryParams.Add("imageFile", parameterToString(data.ImageFile, ""))
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