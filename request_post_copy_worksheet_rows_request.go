/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="post_copy_worksheet_rows_request.go">
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

type PostCopyWorksheetRowsRequest struct {
    Name string `json:"name,omitempty" xml:"name"`
    SheetName string `json:"sheet_name,omitempty" xml:"sheet_name"`
	SourceRowIndex int64 `json:"source_row_index,omitempty" xml:"source_row_index"`
	DestinationRowIndex int64 `json:"destination_row_index,omitempty" xml:"destination_row_index"`
	RowNumber int64 `json:"row_number,omitempty" xml:"row_number"`
	Worksheet string `json:"worksheet,omitempty" xml:"worksheet"`
	Folder string `json:"folder,omitempty" xml:"folder"`
	StorageName string `json:"storage_name,omitempty" xml:"storage_name"`
	
}

func (data *PostCopyWorksheetRowsRequest) CreateRequestData( client *APIClient) (localVarRequest *http.Request, err error) {
	var (
		localVarHttpMethod  = strings.ToUpper("POST")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
	)

	// create path and map variables
	localVarPath := client.cfg.BasePath + "/" + client.cfg.Version + "/cells/{name}/worksheets/{sheetName}/cells/rows/copy"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", data.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", data.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

    // query params : sourceRowIndex
    if &data.SourceRowIndex != nil {
        localVarQueryParams.Add("sourceRowIndex", parameterToString(data.SourceRowIndex, ""))
    }

    // query params : destinationRowIndex
    if &data.DestinationRowIndex != nil {
        localVarQueryParams.Add("destinationRowIndex", parameterToString(data.DestinationRowIndex, ""))
    }

    // query params : rowNumber
    if &data.RowNumber != nil {
        localVarQueryParams.Add("rowNumber", parameterToString(data.RowNumber, ""))
    }

    // query params : worksheet
    if data.Worksheet != "" {
        localVarQueryParams.Add("worksheet", parameterToString(data.Worksheet, ""))
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