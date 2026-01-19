/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="import_data_into_spreadsheet_request.go">
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
	"path/filepath"
	"net/url"
	"net/http"
	"strings"
)

type ImportDataIntoSpreadsheetRequest struct {
	Worksheet string `json:"worksheet,omitempty" xml:"worksheet"`
	Startcell string `json:"startcell,omitempty" xml:"startcell"`
	Insert bool `json:"insert,omitempty" xml:"insert"`
	ConvertNumericData bool `json:"convert_numeric_data,omitempty" xml:"convert_numeric_data"`
	Splitter string `json:"splitter,omitempty" xml:"splitter"`
	OutPath string `json:"out_path,omitempty" xml:"out_path"`
	OutStorageName string `json:"out_storage_name,omitempty" xml:"out_storage_name"`
	FontsLocation string `json:"fonts_location,omitempty" xml:"fonts_location"`
	Region string `json:"region,omitempty" xml:"region"`
	Password string `json:"password,omitempty" xml:"password"`
	
	Datafile string  `json:"datafile,omitempty" xml:"datafile"`
	Spreadsheet string  `json:"spreadsheet,omitempty" xml:"spreadsheet"`
	 

	ExtendQueryParameterMap	map[string]string `json:"ExtendQueryParameterMap,omitempty" xml:"ExtendQueryParameterMap"`	
}

func (data *ImportDataIntoSpreadsheetRequest) CreateRequestData( client *APIClient) (localVarRequest *http.Request, err error) {
	var (
		localVarHttpMethod  = strings.ToUpper("PUT")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
	)

	// create path and map variables
	localVarPath := client.cfg.BasePath + "/v4.0/cells/import/data"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

    // query params : worksheet
    if data.Worksheet != "" {
        localVarQueryParams.Add("worksheet", parameterToString(data.Worksheet, ""))
    }
    // query params : startcell
    if data.Startcell != "" {
        localVarQueryParams.Add("startcell", parameterToString(data.Startcell, ""))
    }
    // query params : insert
    if data.Insert {
        localVarQueryParams.Add("insert", parameterToString(data.Insert, ""))
    }
    // query params : convertNumericData
    if data.ConvertNumericData {
        localVarQueryParams.Add("convertNumericData", parameterToString(data.ConvertNumericData, ""))
    }
    // query params : splitter
    if data.Splitter != "" {
        localVarQueryParams.Add("splitter", parameterToString(data.Splitter, ""))
    }
    // query params : outPath
    if data.OutPath != "" {
        localVarQueryParams.Add("outPath", parameterToString(data.OutPath, ""))
    }
    // query params : outStorageName
    if data.OutStorageName != "" {
        localVarQueryParams.Add("outStorageName", parameterToString(data.OutStorageName, ""))
    }
    // query params : fontsLocation
    if data.FontsLocation != "" {
        localVarQueryParams.Add("fontsLocation", parameterToString(data.FontsLocation, ""))
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
			
		if strings.TrimSpace(data.Datafile) != "" { localVarFormParams["@"+ filepath.Base(data.Datafile)] = []string {data.Datafile}} 
		if strings.TrimSpace(data.Spreadsheet) != "" { localVarFormParams["@"+ filepath.Base(data.Spreadsheet)] = []string {data.Spreadsheet}} 
	r, err := client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	return r,err
}