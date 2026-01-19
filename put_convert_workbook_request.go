/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="put_convert_workbook_request.go">
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

type PutConvertWorkbookRequest struct {
	Format string `json:"format,omitempty" xml:"format"`
	Password string `json:"password,omitempty" xml:"password"`
	OutPath string `json:"out_path,omitempty" xml:"out_path"`
	StorageName string `json:"storage_name,omitempty" xml:"storage_name"`
	CheckExcelRestriction bool `json:"check_excel_restriction,omitempty" xml:"check_excel_restriction"`
	StreamFormat string `json:"stream_format,omitempty" xml:"stream_format"`
	Region string `json:"region,omitempty" xml:"region"`
	PageWideFitOnPerSheet bool `json:"page_wide_fit_on_per_sheet,omitempty" xml:"page_wide_fit_on_per_sheet"`
	PageTallFitOnPerSheet bool `json:"page_tall_fit_on_per_sheet,omitempty" xml:"page_tall_fit_on_per_sheet"`
	SheetName string `json:"sheet_name,omitempty" xml:"sheet_name"`
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index"`
	OnePagePerSheet bool `json:"one_page_per_sheet,omitempty" xml:"one_page_per_sheet"`
	AutoRowsFit bool `json:"auto_rows_fit,omitempty" xml:"auto_rows_fit"`
	AutoColumnsFit bool `json:"auto_columns_fit,omitempty" xml:"auto_columns_fit"`
	FontsLocation string `json:"fonts_location,omitempty" xml:"fonts_location"`
	
	LocalPath string  `json:"local_path,omitempty" xml:"local_path"`
	 

	ExtendQueryParameterMap	map[string]string `json:"ExtendQueryParameterMap,omitempty" xml:"ExtendQueryParameterMap"`	
}

func (data *PutConvertWorkbookRequest) CreateRequestData( client *APIClient) (localVarRequest *http.Request, err error) {
	var (
		localVarHttpMethod  = strings.ToUpper("PUT")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
	)

	// create path and map variables
	localVarPath := client.cfg.BasePath + "/v3.0/cells/convert"

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
    // query params : outPath
    if data.OutPath != "" {
        localVarQueryParams.Add("outPath", parameterToString(data.OutPath, ""))
    }
    // query params : storageName
    if data.StorageName != "" {
        localVarQueryParams.Add("storageName", parameterToString(data.StorageName, ""))
    }
    // query params : checkExcelRestriction
    if data.CheckExcelRestriction {
        localVarQueryParams.Add("checkExcelRestriction", parameterToString(data.CheckExcelRestriction, ""))
    }
    // query params : streamFormat
    if data.StreamFormat != "" {
        localVarQueryParams.Add("streamFormat", parameterToString(data.StreamFormat, ""))
    }
    // query params : region
    if data.Region != "" {
        localVarQueryParams.Add("region", parameterToString(data.Region, ""))
    }
    // query params : pageWideFitOnPerSheet
    if data.PageWideFitOnPerSheet {
        localVarQueryParams.Add("pageWideFitOnPerSheet", parameterToString(data.PageWideFitOnPerSheet, ""))
    }
    // query params : pageTallFitOnPerSheet
    if data.PageTallFitOnPerSheet {
        localVarQueryParams.Add("pageTallFitOnPerSheet", parameterToString(data.PageTallFitOnPerSheet, ""))
    }
    // query params : sheetName
    if data.SheetName != "" {
        localVarQueryParams.Add("sheetName", parameterToString(data.SheetName, ""))
    }
    // query params : pageIndex
    if &data.PageIndex != nil {
        localVarQueryParams.Add("pageIndex", parameterToString(data.PageIndex, ""))
    }
    // query params : onePagePerSheet
    if data.OnePagePerSheet {
        localVarQueryParams.Add("onePagePerSheet", parameterToString(data.OnePagePerSheet, ""))
    }
    // query params : autoRowsFit
    if data.AutoRowsFit {
        localVarQueryParams.Add("AutoRowsFit", parameterToString(data.AutoRowsFit, ""))
    }
    // query params : autoColumnsFit
    if data.AutoColumnsFit {
        localVarQueryParams.Add("AutoColumnsFit", parameterToString(data.AutoColumnsFit, ""))
    }
    // query params : fontsLocation
    if data.FontsLocation != "" {
        localVarQueryParams.Add("FontsLocation", parameterToString(data.FontsLocation, ""))
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
			
		if strings.TrimSpace(data.LocalPath) != "" { localVarFormParams["@"+ filepath.Base(data.LocalPath)] = []string {data.LocalPath}} 
	r, err := client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	return r,err
}