/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="get_workbook_request.go">
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

type GetWorkbookRequest struct {
    Name string `json:"name,omitempty" xml:"name"`
	Format string `json:"format,omitempty" xml:"format"`
	Password string `json:"password,omitempty" xml:"password"`
	IsAutoFit bool `json:"is_auto_fit,omitempty" xml:"is_auto_fit"`
	OnlySaveTable bool `json:"only_save_table,omitempty" xml:"only_save_table"`
	Folder string `json:"folder,omitempty" xml:"folder"`
	OutPath string `json:"out_path,omitempty" xml:"out_path"`
	StorageName string `json:"storage_name,omitempty" xml:"storage_name"`
	OutStorageName string `json:"out_storage_name,omitempty" xml:"out_storage_name"`
	CheckExcelRestriction bool `json:"check_excel_restriction,omitempty" xml:"check_excel_restriction"`
	
}

func (data *GetWorkbookRequest) CreateRequestData( client *APIClient) (localVarRequest *http.Request, err error) {
	var (
		localVarHttpMethod  = strings.ToUpper("GET")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
	)

	// create path and map variables
	localVarPath := client.cfg.BasePath + "/" + client.cfg.Version + "/cells/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", data.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

    // query params : format
    if data.Format != "" {
        localVarQueryParams.Add("format", parameterToString(data.Format, ""))
    }

    // query params : password
    if data.Password != "" {
        localVarQueryParams.Add("password", parameterToString(data.Password, ""))
    }

    // query params : isAutoFit
    if &data.IsAutoFit != nil {
        localVarQueryParams.Add("isAutoFit", parameterToString(data.IsAutoFit, ""))
    }

    // query params : onlySaveTable
    if &data.OnlySaveTable != nil {
        localVarQueryParams.Add("onlySaveTable", parameterToString(data.OnlySaveTable, ""))
    }

    // query params : folder
    if data.Folder != "" {
        localVarQueryParams.Add("folder", parameterToString(data.Folder, ""))
    }

    // query params : outPath
    if data.OutPath != "" {
        localVarQueryParams.Add("outPath", parameterToString(data.OutPath, ""))
    }

    // query params : storageName
    if data.StorageName != "" {
        localVarQueryParams.Add("storageName", parameterToString(data.StorageName, ""))
    }

    // query params : outStorageName
    if data.OutStorageName != "" {
        localVarQueryParams.Add("outStorageName", parameterToString(data.OutStorageName, ""))
    }

    // query params : checkExcelRestriction
    if &data.CheckExcelRestriction != nil {
        localVarQueryParams.Add("checkExcelRestriction", parameterToString(data.CheckExcelRestriction, ""))
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