/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="put_worksheet_add_chart_request.go">
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

type PutWorksheetAddChartRequest struct {
    Name string `json:"name,omitempty" xml:"name"`
    SheetName string `json:"sheet_name,omitempty" xml:"sheet_name"`
	ChartType string `json:"chart_type,omitempty" xml:"chart_type"`
	UpperLeftRow int64 `json:"upper_left_row,omitempty" xml:"upper_left_row"`
	UpperLeftColumn int64 `json:"upper_left_column,omitempty" xml:"upper_left_column"`
	LowerRightRow int64 `json:"lower_right_row,omitempty" xml:"lower_right_row"`
	LowerRightColumn int64 `json:"lower_right_column,omitempty" xml:"lower_right_column"`
	Area string `json:"area,omitempty" xml:"area"`
	IsVertical bool `json:"is_vertical,omitempty" xml:"is_vertical"`
	CategoryData string `json:"category_data,omitempty" xml:"category_data"`
	IsAutoGetSerialName bool `json:"is_auto_get_serial_name,omitempty" xml:"is_auto_get_serial_name"`
	Title string `json:"title,omitempty" xml:"title"`
	Folder string `json:"folder,omitempty" xml:"folder"`
	DataLabels bool `json:"data_labels,omitempty" xml:"data_labels"`
	DataLabelsPosition string `json:"data_labels_position,omitempty" xml:"data_labels_position"`
	PivotTableSheet string `json:"pivot_table_sheet,omitempty" xml:"pivot_table_sheet"`
	PivotTableName string `json:"pivot_table_name,omitempty" xml:"pivot_table_name"`
	StorageName string `json:"storage_name,omitempty" xml:"storage_name"`
	
}

func (data *PutWorksheetAddChartRequest) CreateRequestData( client *APIClient) (localVarRequest *http.Request, err error) {
	var (
		localVarHttpMethod  = strings.ToUpper("PUT")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
	)

	// create path and map variables
	localVarPath := client.cfg.BasePath + "/" + client.cfg.Version + "/cells/{name}/worksheets/{sheetName}/charts"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", data.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", data.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

    // query params : chartType
    if data.ChartType != "" {
        localVarQueryParams.Add("chartType", parameterToString(data.ChartType, ""))
    }

    // query params : upperLeftRow
    if &data.UpperLeftRow != nil {
        localVarQueryParams.Add("upperLeftRow", parameterToString(data.UpperLeftRow, ""))
    }

    // query params : upperLeftColumn
    if &data.UpperLeftColumn != nil {
        localVarQueryParams.Add("upperLeftColumn", parameterToString(data.UpperLeftColumn, ""))
    }

    // query params : lowerRightRow
    if &data.LowerRightRow != nil {
        localVarQueryParams.Add("lowerRightRow", parameterToString(data.LowerRightRow, ""))
    }

    // query params : lowerRightColumn
    if &data.LowerRightColumn != nil {
        localVarQueryParams.Add("lowerRightColumn", parameterToString(data.LowerRightColumn, ""))
    }

    // query params : area
    if data.Area != "" {
        localVarQueryParams.Add("area", parameterToString(data.Area, ""))
    }

    // query params : isVertical
    if &data.IsVertical != nil {
        localVarQueryParams.Add("isVertical", parameterToString(data.IsVertical, ""))
    }

    // query params : categoryData
    if data.CategoryData != "" {
        localVarQueryParams.Add("categoryData", parameterToString(data.CategoryData, ""))
    }

    // query params : isAutoGetSerialName
    if &data.IsAutoGetSerialName != nil {
        localVarQueryParams.Add("isAutoGetSerialName", parameterToString(data.IsAutoGetSerialName, ""))
    }

    // query params : title
    if data.Title != "" {
        localVarQueryParams.Add("title", parameterToString(data.Title, ""))
    }

    // query params : folder
    if data.Folder != "" {
        localVarQueryParams.Add("folder", parameterToString(data.Folder, ""))
    }

    // query params : dataLabels
    if &data.DataLabels != nil {
        localVarQueryParams.Add("dataLabels", parameterToString(data.DataLabels, ""))
    }

    // query params : dataLabelsPosition
    if data.DataLabelsPosition != "" {
        localVarQueryParams.Add("dataLabelsPosition", parameterToString(data.DataLabelsPosition, ""))
    }

    // query params : pivotTableSheet
    if data.PivotTableSheet != "" {
        localVarQueryParams.Add("pivotTableSheet", parameterToString(data.PivotTableSheet, ""))
    }

    // query params : pivotTableName
    if data.PivotTableName != "" {
        localVarQueryParams.Add("pivotTableName", parameterToString(data.PivotTableName, ""))
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
