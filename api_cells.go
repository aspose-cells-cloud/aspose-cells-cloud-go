/*
 *  Copyright (c) 2020 Aspose.Cells Cloud
 *  Permission is hereby granted, free of charge, to any person obtaining a copy
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
 *
 */

package asposecellscloud

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"os"
)


/* Create Instance of CellsApiService
 @param appSid string Application SID
 @param appKey string Application Key
 @param basePath string Base service path. Set "" for default
 @return *PdfApiService */
 func NewCellsApiService(appSid string, appKey string, basePath string) *CellsApiService {
	config := NewConfiguration(appSid, appKey, basePath)
	client := NewAPIClient(config)
	return client.CellsApi
}

type CellsApiService service

/* 
CellsApiService Removes a date filter.             
 * @param name
 * @param sheetName
 * @param fieldIndex
 * @param dateTimeGroupingType
 * @param optional nil or *CellsAutoFilterDeleteWorksheetDateFilterOpts - Optional Parameters:
     * @param "Year" (optional.Int32) - 
     * @param "Month" (optional.Int32) - 
     * @param "Day" (optional.Int32) - 
     * @param "Hour" (optional.Int32) - 
     * @param "Minute" (optional.Int32) - 
     * @param "Second" (optional.Int32) - 
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsAutoFilterDeleteWorksheetDateFilterOpts struct { 
	Name string
	SheetName string
	FieldIndex int32
	DateTimeGroupingType string
	Year int32
	Month int32
	Day int32
	Hour int32
	Minute int32
	Second int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsAutoFilterDeleteWorksheetDateFilter(    localVarOptionals *CellsAutoFilterDeleteWorksheetDateFilterOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/autoFilter/dateFilter"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("fieldIndex", parameterToString(localVarOptionals.FieldIndex, ""))
	localVarQueryParams.Add("dateTimeGroupingType", parameterToString(localVarOptionals.DateTimeGroupingType, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("year", parameterToString(localVarOptionals.Year.Value(), ""))
		localVarQueryParams.Add("year", parameterToString(localVarOptionals.Year, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("month", parameterToString(localVarOptionals.Month.Value(), ""))
		localVarQueryParams.Add("month", parameterToString(localVarOptionals.Month, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("day", parameterToString(localVarOptionals.Day.Value(), ""))
		localVarQueryParams.Add("day", parameterToString(localVarOptionals.Day, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("hour", parameterToString(localVarOptionals.Hour.Value(), ""))
		localVarQueryParams.Add("hour", parameterToString(localVarOptionals.Hour, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("minute", parameterToString(localVarOptionals.Minute.Value(), ""))
		localVarQueryParams.Add("minute", parameterToString(localVarOptionals.Minute, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("second", parameterToString(localVarOptionals.Second.Value(), ""))
		localVarQueryParams.Add("second", parameterToString(localVarOptionals.Second, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Delete a filter for a filter column.             
 * @param name
 * @param sheetName
 * @param fieldIndex
 * @param optional nil or *CellsAutoFilterDeleteWorksheetFilterOpts - Optional Parameters:
     * @param "Criteria" (optional.String) - 
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsAutoFilterDeleteWorksheetFilterOpts struct { 
	Name string
	SheetName string
	FieldIndex int32
	Criteria string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsAutoFilterDeleteWorksheetFilter(    localVarOptionals *CellsAutoFilterDeleteWorksheetFilterOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/autoFilter/filter"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("fieldIndex", parameterToString(localVarOptionals.FieldIndex, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("criteria", parameterToString(localVarOptionals.Criteria.Value(), ""))
		localVarQueryParams.Add("criteria", parameterToString(localVarOptionals.Criteria, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get Auto filter Description
 * @param name
 * @param sheetName
 * @param optional nil or *CellsAutoFilterGetWorksheetAutoFilterOpts - Optional Parameters:
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return AutoFilterResponse
*/


type CellsAutoFilterGetWorksheetAutoFilterOpts struct { 
	Name string
	SheetName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsAutoFilterGetWorksheetAutoFilter(    localVarOptionals *CellsAutoFilterGetWorksheetAutoFilterOpts) (AutoFilterResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue AutoFilterResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/autoFilter"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService
 * @param name
 * @param sheetName
 * @param optional nil or *CellsAutoFilterPostWorksheetAutoFilterRefreshOpts - Optional Parameters:
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsAutoFilterPostWorksheetAutoFilterRefreshOpts struct { 
	Name string
	SheetName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsAutoFilterPostWorksheetAutoFilterRefresh(    localVarOptionals *CellsAutoFilterPostWorksheetAutoFilterRefreshOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/autoFilter/refresh"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Match all blank cell in the list.
 * @param name
 * @param sheetName
 * @param fieldIndex
 * @param optional nil or *CellsAutoFilterPostWorksheetMatchBlanksOpts - Optional Parameters:
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsAutoFilterPostWorksheetMatchBlanksOpts struct { 
	Name string
	SheetName string
	FieldIndex int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsAutoFilterPostWorksheetMatchBlanks(    localVarOptionals *CellsAutoFilterPostWorksheetMatchBlanksOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/autoFilter/matchBlanks"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("fieldIndex", parameterToString(localVarOptionals.FieldIndex, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Match all not blank cell in the list.             
 * @param name
 * @param sheetName
 * @param fieldIndex
 * @param optional nil or *CellsAutoFilterPostWorksheetMatchNonBlanksOpts - Optional Parameters:
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsAutoFilterPostWorksheetMatchNonBlanksOpts struct { 
	Name string
	SheetName string
	FieldIndex int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsAutoFilterPostWorksheetMatchNonBlanks(    localVarOptionals *CellsAutoFilterPostWorksheetMatchNonBlanksOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/autoFilter/matchNonBlanks"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("fieldIndex", parameterToString(localVarOptionals.FieldIndex, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService
 * @param name
 * @param sheetName
 * @param range_
 * @param fieldIndex
 * @param optional nil or *CellsAutoFilterPutWorksheetColorFilterOpts - Optional Parameters:
     * @param "ColorFilter" (optional.Interface of ColorFilterRequest) - 
     * @param "MatchBlanks" (optional.Bool) - 
     * @param "Refresh" (optional.Bool) - 
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsAutoFilterPutWorksheetColorFilterOpts struct { 
	Name string
	SheetName string
	Range_ string
	FieldIndex int32
	ColorFilter *ColorFilterRequest
	MatchBlanks bool
	Refresh bool
	Folder string
	Storage string
}


func (a *CellsApiService) CellsAutoFilterPutWorksheetColorFilter(    localVarOptionals *CellsAutoFilterPutWorksheetColorFilterOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/autoFilter/colorFilter"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("range", parameterToString(localVarOptionals.Range_, ""))
	localVarQueryParams.Add("fieldIndex", parameterToString(localVarOptionals.FieldIndex, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("matchBlanks", parameterToString(localVarOptionals.MatchBlanks.Value(), ""))
		localVarQueryParams.Add("matchBlanks", parameterToString(localVarOptionals.MatchBlanks, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("refresh", parameterToString(localVarOptionals.Refresh.Value(), ""))
		localVarQueryParams.Add("refresh", parameterToString(localVarOptionals.Refresh, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.ColorFilter.IsSet() {
//		
//		localVarOptionalColorFilter, localVarOptionalColorFilterok := localVarOptionals.ColorFilter.Value().(ColorFilterRequest)
//		if !localVarOptionalColorFilterok {
//				return localVarReturnValue, nil, reportError("colorFilter should be ColorFilterRequest")
//		}
//		localVarPostBody = &localVarOptionalColorFilter
//	}
	if localVarOptionals != nil &&  &localVarOptionals.ColorFilter != nil {
		
		localVarPostBody = &localVarOptionals.ColorFilter
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Filters a list with a custom criteria.             
 * @param name
 * @param sheetName
 * @param range_
 * @param fieldIndex
 * @param operatorType1
 * @param criteria1
 * @param optional nil or *CellsAutoFilterPutWorksheetCustomFilterOpts - Optional Parameters:
     * @param "IsAnd" (optional.Bool) - 
     * @param "OperatorType2" (optional.String) - 
     * @param "Criteria2" (optional.String) - 
     * @param "MatchBlanks" (optional.Bool) - 
     * @param "Refresh" (optional.Bool) - 
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsAutoFilterPutWorksheetCustomFilterOpts struct { 
	Name string
	SheetName string
	Range_ string
	FieldIndex int32
	OperatorType1 string
	Criteria1 string
	IsAnd bool
	OperatorType2 string
	Criteria2 string
	MatchBlanks bool
	Refresh bool
	Folder string
	Storage string
}


func (a *CellsApiService) CellsAutoFilterPutWorksheetCustomFilter(    localVarOptionals *CellsAutoFilterPutWorksheetCustomFilterOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/autoFilter/custom"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("range", parameterToString(localVarOptionals.Range_, ""))
	localVarQueryParams.Add("fieldIndex", parameterToString(localVarOptionals.FieldIndex, ""))
	localVarQueryParams.Add("operatorType1", parameterToString(localVarOptionals.OperatorType1, ""))
	localVarQueryParams.Add("criteria1", parameterToString(localVarOptionals.Criteria1, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("isAnd", parameterToString(localVarOptionals.IsAnd.Value(), ""))
		localVarQueryParams.Add("isAnd", parameterToString(localVarOptionals.IsAnd, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("operatorType2", parameterToString(localVarOptionals.OperatorType2.Value(), ""))
		localVarQueryParams.Add("operatorType2", parameterToString(localVarOptionals.OperatorType2, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("criteria2", parameterToString(localVarOptionals.Criteria2.Value(), ""))
		localVarQueryParams.Add("criteria2", parameterToString(localVarOptionals.Criteria2, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("matchBlanks", parameterToString(localVarOptionals.MatchBlanks.Value(), ""))
		localVarQueryParams.Add("matchBlanks", parameterToString(localVarOptionals.MatchBlanks, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("refresh", parameterToString(localVarOptionals.Refresh.Value(), ""))
		localVarQueryParams.Add("refresh", parameterToString(localVarOptionals.Refresh, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService add date filter in worksheet 
 * @param name
 * @param sheetName
 * @param range_
 * @param fieldIndex
 * @param dateTimeGroupingType
 * @param optional nil or *CellsAutoFilterPutWorksheetDateFilterOpts - Optional Parameters:
     * @param "Year" (optional.Int32) - 
     * @param "Month" (optional.Int32) - 
     * @param "Day" (optional.Int32) - 
     * @param "Hour" (optional.Int32) - 
     * @param "Minute" (optional.Int32) - 
     * @param "Second" (optional.Int32) - 
     * @param "MatchBlanks" (optional.Bool) - 
     * @param "Refresh" (optional.Bool) - 
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsAutoFilterPutWorksheetDateFilterOpts struct { 
	Name string
	SheetName string
	Range_ string
	FieldIndex int32
	DateTimeGroupingType string
	Year int32
	Month int32
	Day int32
	Hour int32
	Minute int32
	Second int32
	MatchBlanks bool
	Refresh bool
	Folder string
	Storage string
}


func (a *CellsApiService) CellsAutoFilterPutWorksheetDateFilter(    localVarOptionals *CellsAutoFilterPutWorksheetDateFilterOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/autoFilter/dateFilter"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("range", parameterToString(localVarOptionals.Range_, ""))
	localVarQueryParams.Add("fieldIndex", parameterToString(localVarOptionals.FieldIndex, ""))
	localVarQueryParams.Add("dateTimeGroupingType", parameterToString(localVarOptionals.DateTimeGroupingType, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("year", parameterToString(localVarOptionals.Year.Value(), ""))
		localVarQueryParams.Add("year", parameterToString(localVarOptionals.Year, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("month", parameterToString(localVarOptionals.Month.Value(), ""))
		localVarQueryParams.Add("month", parameterToString(localVarOptionals.Month, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("day", parameterToString(localVarOptionals.Day.Value(), ""))
		localVarQueryParams.Add("day", parameterToString(localVarOptionals.Day, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("hour", parameterToString(localVarOptionals.Hour.Value(), ""))
		localVarQueryParams.Add("hour", parameterToString(localVarOptionals.Hour, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("minute", parameterToString(localVarOptionals.Minute.Value(), ""))
		localVarQueryParams.Add("minute", parameterToString(localVarOptionals.Minute, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("second", parameterToString(localVarOptionals.Second.Value(), ""))
		localVarQueryParams.Add("second", parameterToString(localVarOptionals.Second, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("matchBlanks", parameterToString(localVarOptionals.MatchBlanks.Value(), ""))
		localVarQueryParams.Add("matchBlanks", parameterToString(localVarOptionals.MatchBlanks, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("refresh", parameterToString(localVarOptionals.Refresh.Value(), ""))
		localVarQueryParams.Add("refresh", parameterToString(localVarOptionals.Refresh, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService
 * @param name
 * @param sheetName
 * @param range_
 * @param fieldIndex
 * @param dynamicFilterType
 * @param optional nil or *CellsAutoFilterPutWorksheetDynamicFilterOpts - Optional Parameters:
     * @param "MatchBlanks" (optional.Bool) - 
     * @param "Refresh" (optional.Bool) - 
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsAutoFilterPutWorksheetDynamicFilterOpts struct { 
	Name string
	SheetName string
	Range_ string
	FieldIndex int32
	DynamicFilterType string
	MatchBlanks bool
	Refresh bool
	Folder string
	Storage string
}


func (a *CellsApiService) CellsAutoFilterPutWorksheetDynamicFilter(    localVarOptionals *CellsAutoFilterPutWorksheetDynamicFilterOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/autoFilter/dynamicFilter"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("range", parameterToString(localVarOptionals.Range_, ""))
	localVarQueryParams.Add("fieldIndex", parameterToString(localVarOptionals.FieldIndex, ""))
	localVarQueryParams.Add("dynamicFilterType", parameterToString(localVarOptionals.DynamicFilterType, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("matchBlanks", parameterToString(localVarOptionals.MatchBlanks.Value(), ""))
		localVarQueryParams.Add("matchBlanks", parameterToString(localVarOptionals.MatchBlanks, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("refresh", parameterToString(localVarOptionals.Refresh.Value(), ""))
		localVarQueryParams.Add("refresh", parameterToString(localVarOptionals.Refresh, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Adds a filter for a filter column.             
 * @param name
 * @param sheetName
 * @param range_
 * @param fieldIndex
 * @param criteria
 * @param optional nil or *CellsAutoFilterPutWorksheetFilterOpts - Optional Parameters:
     * @param "MatchBlanks" (optional.Bool) - 
     * @param "Refresh" (optional.Bool) - 
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsAutoFilterPutWorksheetFilterOpts struct { 
	Name string
	SheetName string
	Range_ string
	FieldIndex int32
	Criteria string
	MatchBlanks bool
	Refresh bool
	Folder string
	Storage string
}


func (a *CellsApiService) CellsAutoFilterPutWorksheetFilter(    localVarOptionals *CellsAutoFilterPutWorksheetFilterOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/autoFilter/filter"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("range", parameterToString(localVarOptionals.Range_, ""))
	localVarQueryParams.Add("fieldIndex", parameterToString(localVarOptionals.FieldIndex, ""))
	localVarQueryParams.Add("criteria", parameterToString(localVarOptionals.Criteria, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("matchBlanks", parameterToString(localVarOptionals.MatchBlanks.Value(), ""))
		localVarQueryParams.Add("matchBlanks", parameterToString(localVarOptionals.MatchBlanks, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("refresh", parameterToString(localVarOptionals.Refresh.Value(), ""))
		localVarQueryParams.Add("refresh", parameterToString(localVarOptionals.Refresh, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Filter the top 10 item in the list
 * @param name
 * @param sheetName
 * @param range_
 * @param fieldIndex
 * @param isTop
 * @param isPercent
 * @param itemCount
 * @param optional nil or *CellsAutoFilterPutWorksheetFilterTop10Opts - Optional Parameters:
     * @param "MatchBlanks" (optional.Bool) - 
     * @param "Refresh" (optional.Bool) - 
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsAutoFilterPutWorksheetFilterTop10Opts struct { 
	Name string
	SheetName string
	Range_ string
	FieldIndex int32
	IsTop bool
	IsPercent bool
	ItemCount int32
	MatchBlanks bool
	Refresh bool
	Folder string
	Storage string
}


func (a *CellsApiService) CellsAutoFilterPutWorksheetFilterTop10(    localVarOptionals *CellsAutoFilterPutWorksheetFilterTop10Opts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/autoFilter/filterTop10"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("range", parameterToString(localVarOptionals.Range_, ""))
	localVarQueryParams.Add("fieldIndex", parameterToString(localVarOptionals.FieldIndex, ""))
	localVarQueryParams.Add("isTop", parameterToString(localVarOptionals.IsTop, ""))
	localVarQueryParams.Add("isPercent", parameterToString(localVarOptionals.IsPercent, ""))
	localVarQueryParams.Add("itemCount", parameterToString(localVarOptionals.ItemCount, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("matchBlanks", parameterToString(localVarOptionals.MatchBlanks.Value(), ""))
		localVarQueryParams.Add("matchBlanks", parameterToString(localVarOptionals.MatchBlanks, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("refresh", parameterToString(localVarOptionals.Refresh.Value(), ""))
		localVarQueryParams.Add("refresh", parameterToString(localVarOptionals.Refresh, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Adds an icon filter.
 * @param name
 * @param sheetName
 * @param range_
 * @param fieldIndex
 * @param iconSetType
 * @param iconId
 * @param optional nil or *CellsAutoFilterPutWorksheetIconFilterOpts - Optional Parameters:
     * @param "MatchBlanks" (optional.Bool) - 
     * @param "Refresh" (optional.Bool) - 
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsAutoFilterPutWorksheetIconFilterOpts struct { 
	Name string
	SheetName string
	Range_ string
	FieldIndex int32
	IconSetType string
	IconId int32
	MatchBlanks bool
	Refresh bool
	Folder string
	Storage string
}


func (a *CellsApiService) CellsAutoFilterPutWorksheetIconFilter(    localVarOptionals *CellsAutoFilterPutWorksheetIconFilterOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/autoFilter/iconFilter"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("range", parameterToString(localVarOptionals.Range_, ""))
	localVarQueryParams.Add("fieldIndex", parameterToString(localVarOptionals.FieldIndex, ""))
	localVarQueryParams.Add("iconSetType", parameterToString(localVarOptionals.IconSetType, ""))
	localVarQueryParams.Add("iconId", parameterToString(localVarOptionals.IconId, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("matchBlanks", parameterToString(localVarOptionals.MatchBlanks.Value(), ""))
		localVarQueryParams.Add("matchBlanks", parameterToString(localVarOptionals.MatchBlanks, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("refresh", parameterToString(localVarOptionals.Refresh.Value(), ""))
		localVarQueryParams.Add("refresh", parameterToString(localVarOptionals.Refresh, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get autoshape info.
 * @param name Document name.
 * @param sheetName Worksheet name.
 * @param autoshapeNumber The autoshape number.
 * @param optional nil or *CellsAutoshapesGetWorksheetAutoshapeOpts - Optional Parameters:
     * @param "Format" (optional.String) -  Exported format.
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return *os.File
*/


type CellsAutoshapesGetWorksheetAutoshapeOpts struct { 
	Name string
	SheetName string
	AutoshapeNumber int32
	Format string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsAutoshapesGetWorksheetAutoshape(    localVarOptionals *CellsAutoshapesGetWorksheetAutoshapeOpts) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/autoshapes/{autoshapeNumber}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"autoshapeNumber"+"}", fmt.Sprintf("%v", localVarOptionals.AutoshapeNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("format", parameterToString(localVarOptionals.Format.Value(), ""))
		localVarQueryParams.Add("format", parameterToString(localVarOptionals.Format, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get worksheet autoshapes info.
 * @param name Document name.
 * @param sheetName The worksheet name.
 * @param optional nil or *CellsAutoshapesGetWorksheetAutoshapesOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return AutoShapesResponse
*/


type CellsAutoshapesGetWorksheetAutoshapesOpts struct { 
	Name string
	SheetName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsAutoshapesGetWorksheetAutoshapes(    localVarOptionals *CellsAutoshapesGetWorksheetAutoshapesOpts) (AutoShapesResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue AutoShapesResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/autoshapes"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get chart area info.
 * @param name Workbook name.
 * @param sheetName Worksheet name.
 * @param chartIndex The chart index.
 * @param optional nil or *CellsChartAreaGetChartAreaOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return ChartAreaResponse
*/


type CellsChartAreaGetChartAreaOpts struct { 
	Name string
	SheetName string
	ChartIndex int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsChartAreaGetChartArea(    localVarOptionals *CellsChartAreaGetChartAreaOpts) (ChartAreaResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ChartAreaResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/charts/{chartIndex}/chartArea"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"chartIndex"+"}", fmt.Sprintf("%v", localVarOptionals.ChartIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get chart area border info.
 * @param name Workbook name.
 * @param sheetName Worksheet name.
 * @param chartIndex The chart index.
 * @param optional nil or *CellsChartAreaGetChartAreaBorderOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return LineResponse
*/


type CellsChartAreaGetChartAreaBorderOpts struct { 
	Name string
	SheetName string
	ChartIndex int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsChartAreaGetChartAreaBorder(    localVarOptionals *CellsChartAreaGetChartAreaBorderOpts) (LineResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue LineResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/charts/{chartIndex}/chartArea/border"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"chartIndex"+"}", fmt.Sprintf("%v", localVarOptionals.ChartIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get chart area fill format info.
 * @param name Workbook name.
 * @param sheetName Worksheet name.
 * @param chartIndex The chart index.
 * @param optional nil or *CellsChartAreaGetChartAreaFillFormatOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return FillFormatResponse
*/


type CellsChartAreaGetChartAreaFillFormatOpts struct { 
	Name string
	SheetName string
	ChartIndex int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsChartAreaGetChartAreaFillFormat(    localVarOptionals *CellsChartAreaGetChartAreaFillFormatOpts) (FillFormatResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue FillFormatResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/charts/{chartIndex}/chartArea/fillFormat"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"chartIndex"+"}", fmt.Sprintf("%v", localVarOptionals.ChartIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Hide legend in chart
 * @param name Workbook name.
 * @param sheetName Worksheet name.
 * @param chartIndex The chart index.
 * @param optional nil or *CellsChartsDeleteWorksheetChartLegendOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsChartsDeleteWorksheetChartLegendOpts struct { 
	Name string
	SheetName string
	ChartIndex int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsChartsDeleteWorksheetChartLegend(    localVarOptionals *CellsChartsDeleteWorksheetChartLegendOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/charts/{chartIndex}/legend"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"chartIndex"+"}", fmt.Sprintf("%v", localVarOptionals.ChartIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Hide title in chart
 * @param name Workbook name.
 * @param sheetName Worksheet name.
 * @param chartIndex The chart index.
 * @param optional nil or *CellsChartsDeleteWorksheetChartTitleOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsChartsDeleteWorksheetChartTitleOpts struct { 
	Name string
	SheetName string
	ChartIndex int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsChartsDeleteWorksheetChartTitle(    localVarOptionals *CellsChartsDeleteWorksheetChartTitleOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/charts/{chartIndex}/title"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"chartIndex"+"}", fmt.Sprintf("%v", localVarOptionals.ChartIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Clear the charts.
 * @param name Workbook name.
 * @param sheetName The worksheet name.
 * @param optional nil or *CellsChartsDeleteWorksheetClearChartsOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsChartsDeleteWorksheetClearChartsOpts struct { 
	Name string
	SheetName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsChartsDeleteWorksheetClearCharts(    localVarOptionals *CellsChartsDeleteWorksheetClearChartsOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/charts"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Delete worksheet chart by index.
 * @param name Workbook name.
 * @param sheetName Worksheet name.
 * @param chartIndex The chart index.
 * @param optional nil or *CellsChartsDeleteWorksheetDeleteChartOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return ChartsResponse
*/


type CellsChartsDeleteWorksheetDeleteChartOpts struct { 
	Name string
	SheetName string
	ChartIndex int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsChartsDeleteWorksheetDeleteChart(    localVarOptionals *CellsChartsDeleteWorksheetDeleteChartOpts) (ChartsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ChartsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/charts/{chartIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"chartIndex"+"}", fmt.Sprintf("%v", localVarOptionals.ChartIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get chart info.
 * @param name Document name.
 * @param sheetName Worksheet name.
 * @param chartNumber The chart number.
 * @param optional nil or *CellsChartsGetWorksheetChartOpts - Optional Parameters:
     * @param "Format" (optional.String) -  The exported file format.
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return *os.File
*/


type CellsChartsGetWorksheetChartOpts struct { 
	Name string
	SheetName string
	ChartNumber int32
	Format string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsChartsGetWorksheetChart(    localVarOptionals *CellsChartsGetWorksheetChartOpts) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/charts/{chartNumber}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"chartNumber"+"}", fmt.Sprintf("%v", localVarOptionals.ChartNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("format", parameterToString(localVarOptionals.Format.Value(), ""))
		localVarQueryParams.Add("format", parameterToString(localVarOptionals.Format, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get chart legend
 * @param name Workbook name.
 * @param sheetName Worksheet name.
 * @param chartIndex The chart index.
 * @param optional nil or *CellsChartsGetWorksheetChartLegendOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return LegendResponse
*/


type CellsChartsGetWorksheetChartLegendOpts struct { 
	Name string
	SheetName string
	ChartIndex int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsChartsGetWorksheetChartLegend(    localVarOptionals *CellsChartsGetWorksheetChartLegendOpts) (LegendResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue LegendResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/charts/{chartIndex}/legend"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"chartIndex"+"}", fmt.Sprintf("%v", localVarOptionals.ChartIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get chart title
 * @param name Workbook name.
 * @param sheetName Worksheet name.
 * @param chartIndex The chart index.
 * @param optional nil or *CellsChartsGetWorksheetChartTitleOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return TitleResponse
*/


type CellsChartsGetWorksheetChartTitleOpts struct { 
	Name string
	SheetName string
	ChartIndex int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsChartsGetWorksheetChartTitle(    localVarOptionals *CellsChartsGetWorksheetChartTitleOpts) (TitleResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue TitleResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/charts/{chartIndex}/title"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"chartIndex"+"}", fmt.Sprintf("%v", localVarOptionals.ChartIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get worksheet charts info.
 * @param name Document name.
 * @param sheetName The worksheet name.
 * @param optional nil or *CellsChartsGetWorksheetChartsOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return ChartsResponse
*/


type CellsChartsGetWorksheetChartsOpts struct { 
	Name string
	SheetName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsChartsGetWorksheetCharts(    localVarOptionals *CellsChartsGetWorksheetChartsOpts) (ChartsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ChartsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/charts"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Update chart propreties
 * @param name
 * @param sheetName
 * @param chartIndex
 * @param optional nil or *CellsChartsPostWorksheetChartOpts - Optional Parameters:
     * @param "Chart" (optional.Interface of Chart) - 
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsChartsPostWorksheetChartOpts struct { 
	Name string
	SheetName string
	ChartIndex int32
	Chart *Chart
	Folder string
	Storage string
}


func (a *CellsApiService) CellsChartsPostWorksheetChart(    localVarOptionals *CellsChartsPostWorksheetChartOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/charts/{chartIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"chartIndex"+"}", fmt.Sprintf("%v", localVarOptionals.ChartIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Chart.IsSet() {
//		
//		localVarOptionalChart, localVarOptionalChartok := localVarOptionals.Chart.Value().(Chart)
//		if !localVarOptionalChartok {
//				return localVarReturnValue, nil, reportError("chart should be Chart")
//		}
//		localVarPostBody = &localVarOptionalChart
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Chart != nil {
		
		localVarPostBody = &localVarOptionals.Chart
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Update chart legend
 * @param name Workbook name.
 * @param sheetName Worksheet name.
 * @param chartIndex The chart index.
 * @param optional nil or *CellsChartsPostWorksheetChartLegendOpts - Optional Parameters:
     * @param "Legend" (optional.Interface of Legend) - 
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return LegendResponse
*/


type CellsChartsPostWorksheetChartLegendOpts struct { 
	Name string
	SheetName string
	ChartIndex int32
	Legend *Legend
	Folder string
	Storage string
}


func (a *CellsApiService) CellsChartsPostWorksheetChartLegend(    localVarOptionals *CellsChartsPostWorksheetChartLegendOpts) (LegendResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue LegendResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/charts/{chartIndex}/legend"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"chartIndex"+"}", fmt.Sprintf("%v", localVarOptionals.ChartIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Legend.IsSet() {
//		
//		localVarOptionalLegend, localVarOptionalLegendok := localVarOptionals.Legend.Value().(Legend)
//		if !localVarOptionalLegendok {
//				return localVarReturnValue, nil, reportError("legend should be Legend")
//		}
//		localVarPostBody = &localVarOptionalLegend
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Legend != nil {
		
		localVarPostBody = &localVarOptionals.Legend
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Update chart title
 * @param name Workbook name.
 * @param sheetName Worksheet name.
 * @param chartIndex The chart index.
 * @param optional nil or *CellsChartsPostWorksheetChartTitleOpts - Optional Parameters:
     * @param "Title" (optional.Interface of Title) -  Chart title
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return TitleResponse
*/


type CellsChartsPostWorksheetChartTitleOpts struct { 
	Name string
	SheetName string
	ChartIndex int32
	Title *Title
	Folder string
	Storage string
}


func (a *CellsApiService) CellsChartsPostWorksheetChartTitle(    localVarOptionals *CellsChartsPostWorksheetChartTitleOpts) (TitleResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue TitleResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/charts/{chartIndex}/title"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"chartIndex"+"}", fmt.Sprintf("%v", localVarOptionals.ChartIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Title.IsSet() {
//		
//		localVarOptionalTitle, localVarOptionalTitleok := localVarOptionals.Title.Value().(Title)
//		if !localVarOptionalTitleok {
//				return localVarReturnValue, nil, reportError("title should be Title")
//		}
//		localVarPostBody = &localVarOptionalTitle
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Title != nil {
		
		localVarPostBody = &localVarOptionals.Title
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Add new chart to worksheet.
 * @param name Workbook name.
 * @param sheetName The worksheet name.
 * @param chartType Chart type, please refer property Type in chart resource.
 * @param optional nil or *CellsChartsPutWorksheetAddChartOpts - Optional Parameters:
     * @param "UpperLeftRow" (optional.Int32) -  New chart upper left row.
     * @param "UpperLeftColumn" (optional.Int32) -  New chart upperleft column.
     * @param "LowerRightRow" (optional.Int32) -  New chart lower right row.
     * @param "LowerRightColumn" (optional.Int32) -  New chart lower right column.
     * @param "Area" (optional.String) -  Specifies values from which to plot the data series. 
     * @param "IsVertical" (optional.Bool) -  Specifies whether to plot the series from a range of cell values by row or by column. 
     * @param "CategoryData" (optional.String) -  Gets or sets the range of category Axis values. It can be a range of cells (such as, \&quot;d1:e10\&quot;). 
     * @param "IsAutoGetSerialName" (optional.Bool) -  Specifies whether auto update serial name. 
     * @param "Title" (optional.String) -  Specifies chart title name.
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return ChartsResponse
*/


type CellsChartsPutWorksheetAddChartOpts struct { 
	Name string
	SheetName string
	ChartType string
	UpperLeftRow int32
	UpperLeftColumn int32
	LowerRightRow int32
	LowerRightColumn int32
	Area string
	IsVertical bool
	CategoryData string
	IsAutoGetSerialName bool
	Title string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsChartsPutWorksheetAddChart(    localVarOptionals *CellsChartsPutWorksheetAddChartOpts) (ChartsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ChartsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/charts"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("chartType", parameterToString(localVarOptionals.ChartType, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("upperLeftRow", parameterToString(localVarOptionals.UpperLeftRow.Value(), ""))
		localVarQueryParams.Add("upperLeftRow", parameterToString(localVarOptionals.UpperLeftRow, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("upperLeftColumn", parameterToString(localVarOptionals.UpperLeftColumn.Value(), ""))
		localVarQueryParams.Add("upperLeftColumn", parameterToString(localVarOptionals.UpperLeftColumn, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("lowerRightRow", parameterToString(localVarOptionals.LowerRightRow.Value(), ""))
		localVarQueryParams.Add("lowerRightRow", parameterToString(localVarOptionals.LowerRightRow, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("lowerRightColumn", parameterToString(localVarOptionals.LowerRightColumn.Value(), ""))
		localVarQueryParams.Add("lowerRightColumn", parameterToString(localVarOptionals.LowerRightColumn, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("area", parameterToString(localVarOptionals.Area.Value(), ""))
		localVarQueryParams.Add("area", parameterToString(localVarOptionals.Area, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("isVertical", parameterToString(localVarOptionals.IsVertical.Value(), ""))
		localVarQueryParams.Add("isVertical", parameterToString(localVarOptionals.IsVertical, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("categoryData", parameterToString(localVarOptionals.CategoryData.Value(), ""))
		localVarQueryParams.Add("categoryData", parameterToString(localVarOptionals.CategoryData, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("isAutoGetSerialName", parameterToString(localVarOptionals.IsAutoGetSerialName.Value(), ""))
		localVarQueryParams.Add("isAutoGetSerialName", parameterToString(localVarOptionals.IsAutoGetSerialName, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("title", parameterToString(localVarOptionals.Title.Value(), ""))
		localVarQueryParams.Add("title", parameterToString(localVarOptionals.Title, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Show legend in chart
 * @param name Workbook name.
 * @param sheetName Worksheet name.
 * @param chartIndex The chart index.
 * @param optional nil or *CellsChartsPutWorksheetChartLegendOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsChartsPutWorksheetChartLegendOpts struct { 
	Name string
	SheetName string
	ChartIndex int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsChartsPutWorksheetChartLegend(    localVarOptionals *CellsChartsPutWorksheetChartLegendOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/charts/{chartIndex}/legend"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"chartIndex"+"}", fmt.Sprintf("%v", localVarOptionals.ChartIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Add chart title / Set chart title visible
 * @param name Workbook name.
 * @param sheetName Worksheet name.
 * @param chartIndex The chart index.
 * @param optional nil or *CellsChartsPutWorksheetChartTitleOpts - Optional Parameters:
     * @param "Title" (optional.Interface of Title) -  Chart title.
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return TitleResponse
*/


type CellsChartsPutWorksheetChartTitleOpts struct { 
	Name string
	SheetName string
	ChartIndex int32
	Title *Title
	Folder string
	Storage string
}


func (a *CellsApiService) CellsChartsPutWorksheetChartTitle(    localVarOptionals *CellsChartsPutWorksheetChartTitleOpts) (TitleResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue TitleResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/charts/{chartIndex}/title"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"chartIndex"+"}", fmt.Sprintf("%v", localVarOptionals.ChartIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Title.IsSet() {
//		
//		localVarOptionalTitle, localVarOptionalTitleok := localVarOptionals.Title.Value().(Title)
//		if !localVarOptionalTitleok {
//				return localVarReturnValue, nil, reportError("title should be Title")
//		}
//		localVarPostBody = &localVarOptionalTitle
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Title != nil {
		
		localVarPostBody = &localVarOptionals.Title
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Remove conditional formatting
 * @param name
 * @param sheetName
 * @param index
 * @param optional nil or *CellsConditionalFormattingsDeleteWorksheetConditionalFormattingOpts - Optional Parameters:
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsConditionalFormattingsDeleteWorksheetConditionalFormattingOpts struct { 
	Name string
	SheetName string
	Index int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsConditionalFormattingsDeleteWorksheetConditionalFormatting(    localVarOptionals *CellsConditionalFormattingsDeleteWorksheetConditionalFormattingOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/conditionalFormattings/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", localVarOptionals.Index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Remove cell area from conditional formatting.
 * @param name
 * @param sheetName
 * @param startRow
 * @param startColumn
 * @param totalRows
 * @param totalColumns
 * @param optional nil or *CellsConditionalFormattingsDeleteWorksheetConditionalFormattingAreaOpts - Optional Parameters:
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsConditionalFormattingsDeleteWorksheetConditionalFormattingAreaOpts struct { 
	Name string
	SheetName string
	StartRow int32
	StartColumn int32
	TotalRows int32
	TotalColumns int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsConditionalFormattingsDeleteWorksheetConditionalFormattingArea(    localVarOptionals *CellsConditionalFormattingsDeleteWorksheetConditionalFormattingAreaOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/conditionalFormattings/area"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("startRow", parameterToString(localVarOptionals.StartRow, ""))
	localVarQueryParams.Add("startColumn", parameterToString(localVarOptionals.StartColumn, ""))
	localVarQueryParams.Add("totalRows", parameterToString(localVarOptionals.TotalRows, ""))
	localVarQueryParams.Add("totalColumns", parameterToString(localVarOptionals.TotalColumns, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Clear all condition formattings
 * @param name
 * @param sheetName
 * @param optional nil or *CellsConditionalFormattingsDeleteWorksheetConditionalFormattingsOpts - Optional Parameters:
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsConditionalFormattingsDeleteWorksheetConditionalFormattingsOpts struct { 
	Name string
	SheetName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsConditionalFormattingsDeleteWorksheetConditionalFormattings(    localVarOptionals *CellsConditionalFormattingsDeleteWorksheetConditionalFormattingsOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/conditionalFormattings"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get conditional formatting
 * @param name
 * @param sheetName
 * @param index
 * @param optional nil or *CellsConditionalFormattingsGetWorksheetConditionalFormattingOpts - Optional Parameters:
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return ConditionalFormattingResponse
*/


type CellsConditionalFormattingsGetWorksheetConditionalFormattingOpts struct { 
	Name string
	SheetName string
	Index int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsConditionalFormattingsGetWorksheetConditionalFormatting(    localVarOptionals *CellsConditionalFormattingsGetWorksheetConditionalFormattingOpts) (ConditionalFormattingResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ConditionalFormattingResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/conditionalFormattings/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", localVarOptionals.Index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get conditional formattings 
 * @param name
 * @param sheetName
 * @param optional nil or *CellsConditionalFormattingsGetWorksheetConditionalFormattingsOpts - Optional Parameters:
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return ConditionalFormattingsResponse
*/


type CellsConditionalFormattingsGetWorksheetConditionalFormattingsOpts struct { 
	Name string
	SheetName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsConditionalFormattingsGetWorksheetConditionalFormattings(    localVarOptionals *CellsConditionalFormattingsGetWorksheetConditionalFormattingsOpts) (ConditionalFormattingsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ConditionalFormattingsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/conditionalFormattings"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Add a condition formatting.
 * @param name
 * @param sheetName
 * @param cellArea
 * @param optional nil or *CellsConditionalFormattingsPutWorksheetConditionalFormattingOpts - Optional Parameters:
     * @param "FormatCondition" (optional.Interface of FormatCondition) - 
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsConditionalFormattingsPutWorksheetConditionalFormattingOpts struct { 
	Name string
	SheetName string
	CellArea string
	FormatCondition *FormatCondition
	Folder string
	Storage string
}


func (a *CellsApiService) CellsConditionalFormattingsPutWorksheetConditionalFormatting(    localVarOptionals *CellsConditionalFormattingsPutWorksheetConditionalFormattingOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/conditionalFormattings"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("cellArea", parameterToString(localVarOptionals.CellArea, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.FormatCondition.IsSet() {
//		
//		localVarOptionalFormatCondition, localVarOptionalFormatConditionok := localVarOptionals.FormatCondition.Value().(FormatCondition)
//		if !localVarOptionalFormatConditionok {
//				return localVarReturnValue, nil, reportError("formatCondition should be FormatCondition")
//		}
//		localVarPostBody = &localVarOptionalFormatCondition
//	}
	if localVarOptionals != nil &&  &localVarOptionals.FormatCondition != nil {
		
		localVarPostBody = &localVarOptionals.FormatCondition
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Add a format condition.
 * @param name
 * @param sheetName
 * @param index
 * @param cellArea
 * @param type_
 * @param operatorType
 * @param formula1
 * @param formula2
 * @param optional nil or *CellsConditionalFormattingsPutWorksheetFormatConditionOpts - Optional Parameters:
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsConditionalFormattingsPutWorksheetFormatConditionOpts struct { 
	Name string
	SheetName string
	Index int32
	CellArea string
	Type_ string
	OperatorType string
	Formula1 string
	Formula2 string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsConditionalFormattingsPutWorksheetFormatCondition(    localVarOptionals *CellsConditionalFormattingsPutWorksheetFormatConditionOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/conditionalFormattings/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", localVarOptionals.Index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("cellArea", parameterToString(localVarOptionals.CellArea, ""))
	localVarQueryParams.Add("type", parameterToString(localVarOptionals.Type_, ""))
	localVarQueryParams.Add("operatorType", parameterToString(localVarOptionals.OperatorType, ""))
	localVarQueryParams.Add("formula1", parameterToString(localVarOptionals.Formula1, ""))
	localVarQueryParams.Add("formula2", parameterToString(localVarOptionals.Formula2, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService add a cell area for format condition             
 * @param name
 * @param sheetName
 * @param index
 * @param cellArea
 * @param optional nil or *CellsConditionalFormattingsPutWorksheetFormatConditionAreaOpts - Optional Parameters:
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsConditionalFormattingsPutWorksheetFormatConditionAreaOpts struct { 
	Name string
	SheetName string
	Index int32
	CellArea string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsConditionalFormattingsPutWorksheetFormatConditionArea(    localVarOptionals *CellsConditionalFormattingsPutWorksheetFormatConditionAreaOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/conditionalFormattings/{index}/area"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", localVarOptionals.Index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("cellArea", parameterToString(localVarOptionals.CellArea, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Add a condition for format condition.
 * @param name
 * @param sheetName
 * @param index
 * @param type_
 * @param operatorType
 * @param formula1
 * @param formula2
 * @param optional nil or *CellsConditionalFormattingsPutWorksheetFormatConditionConditionOpts - Optional Parameters:
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsConditionalFormattingsPutWorksheetFormatConditionConditionOpts struct { 
	Name string
	SheetName string
	Index int32
	Type_ string
	OperatorType string
	Formula1 string
	Formula2 string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsConditionalFormattingsPutWorksheetFormatConditionCondition(    localVarOptionals *CellsConditionalFormattingsPutWorksheetFormatConditionConditionOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/conditionalFormattings/{index}/condition"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", localVarOptionals.Index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("type", parameterToString(localVarOptionals.Type_, ""))
	localVarQueryParams.Add("operatorType", parameterToString(localVarOptionals.OperatorType, ""))
	localVarQueryParams.Add("formula1", parameterToString(localVarOptionals.Formula1, ""))
	localVarQueryParams.Add("formula2", parameterToString(localVarOptionals.Formula2, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Delete worksheet columns.
 * @param name The workbook name.
 * @param sheetName The worksheet name.
 * @param columnIndex The column index.
 * @param columns The columns.
 * @param updateReference The update reference.
 * @param optional nil or *CellsDeleteWorksheetColumnsOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return ColumnsResponse
*/


type CellsDeleteWorksheetColumnsOpts struct { 
	Name string
	SheetName string
	ColumnIndex int32
	Columns int32
	UpdateReference bool
	Folder string
	Storage string
}


func (a *CellsApiService) CellsDeleteWorksheetColumns(    localVarOptionals *CellsDeleteWorksheetColumnsOpts) (ColumnsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ColumnsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/columns/{columnIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"columnIndex"+"}", fmt.Sprintf("%v", localVarOptionals.ColumnIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("columns", parameterToString(localVarOptionals.Columns, ""))
	localVarQueryParams.Add("updateReference", parameterToString(localVarOptionals.UpdateReference, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Delete worksheet row.
 * @param name The workbook name.
 * @param sheetName The worksheet bame.
 * @param rowIndex The row index.
 * @param optional nil or *CellsDeleteWorksheetRowOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsDeleteWorksheetRowOpts struct { 
	Name string
	SheetName string
	RowIndex int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsDeleteWorksheetRow(    localVarOptionals *CellsDeleteWorksheetRowOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/rows/{rowIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rowIndex"+"}", fmt.Sprintf("%v", localVarOptionals.RowIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Delete several worksheet rows.
 * @param name The workbook name.
 * @param sheetName The worksheet bame.
 * @param startrow The begin row index to be operated.
 * @param optional nil or *CellsDeleteWorksheetRowsOpts - Optional Parameters:
     * @param "TotalRows" (optional.Int32) -  Number of rows to be operated.
     * @param "UpdateReference" (optional.Bool) -  Indicates if update references in other worksheets.
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsDeleteWorksheetRowsOpts struct { 
	Name string
	SheetName string
	Startrow int32
	TotalRows int32
	UpdateReference bool
	Folder string
	Storage string
}


func (a *CellsApiService) CellsDeleteWorksheetRows(    localVarOptionals *CellsDeleteWorksheetRowsOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/rows"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("startrow", parameterToString(localVarOptionals.Startrow, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("totalRows", parameterToString(localVarOptionals.TotalRows.Value(), ""))
		localVarQueryParams.Add("totalRows", parameterToString(localVarOptionals.TotalRows, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("updateReference", parameterToString(localVarOptionals.UpdateReference.Value(), ""))
		localVarQueryParams.Add("updateReference", parameterToString(localVarOptionals.UpdateReference, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Read cell data by cell&#39;s name.
 * @param name Document name.
 * @param sheetName Worksheet name.
 * @param cellName The cell&#39;s  name.
 * @param optional nil or *CellsGetCellHtmlStringOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return interface{}
*/


type CellsGetCellHtmlStringOpts struct { 
	Name string
	SheetName string
	CellName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsGetCellHtmlString(    localVarOptionals *CellsGetCellHtmlStringOpts) (interface{}, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/{cellName}/htmlstring"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cellName"+"}", fmt.Sprintf("%v", localVarOptionals.CellName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Read cell data by cell&#39;s name.
 * @param name Document name.
 * @param sheetName Worksheet name.
 * @param cellOrMethodName The cell&#39;s or method name. (Method name like firstcell, endcell etc.)
 * @param optional nil or *CellsGetWorksheetCellOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return interface{}
*/


type CellsGetWorksheetCellOpts struct { 
	Name string
	SheetName string
	CellOrMethodName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsGetWorksheetCell(    localVarOptionals *CellsGetWorksheetCellOpts) (interface{}, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/{cellOrMethodName}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cellOrMethodName"+"}", fmt.Sprintf("%v", localVarOptionals.CellOrMethodName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Read cell&#39;s style info.
 * @param name Document name.
 * @param sheetName Worksheet name.
 * @param cellName Cell&#39;s name.
 * @param optional nil or *CellsGetWorksheetCellStyleOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return StyleResponse
*/


type CellsGetWorksheetCellStyleOpts struct { 
	Name string
	SheetName string
	CellName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsGetWorksheetCellStyle(    localVarOptionals *CellsGetWorksheetCellStyleOpts) (StyleResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue StyleResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/{cellName}/style"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cellName"+"}", fmt.Sprintf("%v", localVarOptionals.CellName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get cells info.
 * @param name Document name.
 * @param sheetName Worksheet name.
 * @param optional nil or *CellsGetWorksheetCellsOpts - Optional Parameters:
     * @param "Offest" (optional.Int32) -  Begginig offset.
     * @param "Count" (optional.Int32) -  Maximum amount of cells in the response.
     * @param "Folder" (optional.String) -  Document&#39;s folder name.
     * @param "Storage" (optional.String) -  storage name.

@return CellsResponse
*/


type CellsGetWorksheetCellsOpts struct { 
	Name string
	SheetName string
	Offest int32
	Count int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsGetWorksheetCells(    localVarOptionals *CellsGetWorksheetCellsOpts) (CellsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("offest", parameterToString(localVarOptionals.Offest.Value(), ""))
		localVarQueryParams.Add("offest", parameterToString(localVarOptionals.Offest, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("count", parameterToString(localVarOptionals.Count.Value(), ""))
		localVarQueryParams.Add("count", parameterToString(localVarOptionals.Count, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Read worksheet column data by column&#39;s index.
 * @param name The workbook name.
 * @param sheetName The worksheet name.
 * @param columnIndex The column index.
 * @param optional nil or *CellsGetWorksheetColumnOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return ColumnResponse
*/


type CellsGetWorksheetColumnOpts struct { 
	Name string
	SheetName string
	ColumnIndex int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsGetWorksheetColumn(    localVarOptionals *CellsGetWorksheetColumnOpts) (ColumnResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ColumnResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/columns/{columnIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"columnIndex"+"}", fmt.Sprintf("%v", localVarOptionals.ColumnIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Read worksheet columns info.
 * @param name The workbook name.
 * @param sheetName The worksheet name.
 * @param optional nil or *CellsGetWorksheetColumnsOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The workdook folder.
     * @param "Storage" (optional.String) -  storage name.

@return ColumnsResponse
*/


type CellsGetWorksheetColumnsOpts struct { 
	Name string
	SheetName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsGetWorksheetColumns(    localVarOptionals *CellsGetWorksheetColumnsOpts) (ColumnsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ColumnsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/columns"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Read worksheet row data by row&#39;s index.
 * @param name The workbook name.
 * @param sheetName The worksheet name.
 * @param rowIndex The row index.
 * @param optional nil or *CellsGetWorksheetRowOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return RowResponse
*/


type CellsGetWorksheetRowOpts struct { 
	Name string
	SheetName string
	RowIndex int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsGetWorksheetRow(    localVarOptionals *CellsGetWorksheetRowOpts) (RowResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue RowResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/rows/{rowIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rowIndex"+"}", fmt.Sprintf("%v", localVarOptionals.RowIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Read worksheet rows info.
 * @param name The workbook name.
 * @param sheetName The worksheet name.
 * @param optional nil or *CellsGetWorksheetRowsOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The workdook folder.
     * @param "Storage" (optional.String) -  storage name.

@return RowsResponse
*/


type CellsGetWorksheetRowsOpts struct { 
	Name string
	SheetName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsGetWorksheetRows(    localVarOptionals *CellsGetWorksheetRowsOpts) (RowsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue RowsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/rows"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Delete worksheet hyperlink by index.
 * @param name Document name.
 * @param sheetName Worksheet name.
 * @param hyperlinkIndex The hyperlink&#39;s index.
 * @param optional nil or *CellsHypelinksDeleteWorksheetHyperlinkOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsHypelinksDeleteWorksheetHyperlinkOpts struct { 
	Name string
	SheetName string
	HyperlinkIndex int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsHypelinksDeleteWorksheetHyperlink(    localVarOptionals *CellsHypelinksDeleteWorksheetHyperlinkOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/hyperlinks/{hyperlinkIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"hyperlinkIndex"+"}", fmt.Sprintf("%v", localVarOptionals.HyperlinkIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Delete all hyperlinks in worksheet.
 * @param name Document name.
 * @param sheetName Worksheet name.
 * @param optional nil or *CellsHypelinksDeleteWorksheetHyperlinksOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsHypelinksDeleteWorksheetHyperlinksOpts struct { 
	Name string
	SheetName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsHypelinksDeleteWorksheetHyperlinks(    localVarOptionals *CellsHypelinksDeleteWorksheetHyperlinksOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/hyperlinks"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get worksheet hyperlink by index.
 * @param name Document name.
 * @param sheetName Worksheet name.
 * @param hyperlinkIndex The hyperlink&#39;s index.
 * @param optional nil or *CellsHypelinksGetWorksheetHyperlinkOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return HyperlinkResponse
*/


type CellsHypelinksGetWorksheetHyperlinkOpts struct { 
	Name string
	SheetName string
	HyperlinkIndex int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsHypelinksGetWorksheetHyperlink(    localVarOptionals *CellsHypelinksGetWorksheetHyperlinkOpts) (HyperlinkResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue HyperlinkResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/hyperlinks/{hyperlinkIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"hyperlinkIndex"+"}", fmt.Sprintf("%v", localVarOptionals.HyperlinkIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get worksheet hyperlinks.
 * @param name Document name.
 * @param sheetName The worksheet name.
 * @param optional nil or *CellsHypelinksGetWorksheetHyperlinksOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return HyperlinksResponse
*/


type CellsHypelinksGetWorksheetHyperlinksOpts struct { 
	Name string
	SheetName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsHypelinksGetWorksheetHyperlinks(    localVarOptionals *CellsHypelinksGetWorksheetHyperlinksOpts) (HyperlinksResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue HyperlinksResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/hyperlinks"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Update worksheet hyperlink by index.
 * @param name Document name.
 * @param sheetName Worksheet name.
 * @param hyperlinkIndex The hyperlink&#39;s index.
 * @param optional nil or *CellsHypelinksPostWorksheetHyperlinkOpts - Optional Parameters:
     * @param "Hyperlink" (optional.Interface of Hyperlink) -  Hyperlink object
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return HyperlinkResponse
*/


type CellsHypelinksPostWorksheetHyperlinkOpts struct { 
	Name string
	SheetName string
	HyperlinkIndex int32
	Hyperlink *Hyperlink
	Folder string
	Storage string
}


func (a *CellsApiService) CellsHypelinksPostWorksheetHyperlink(    localVarOptionals *CellsHypelinksPostWorksheetHyperlinkOpts) (HyperlinkResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue HyperlinkResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/hyperlinks/{hyperlinkIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"hyperlinkIndex"+"}", fmt.Sprintf("%v", localVarOptionals.HyperlinkIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Hyperlink.IsSet() {
//		
//		localVarOptionalHyperlink, localVarOptionalHyperlinkok := localVarOptionals.Hyperlink.Value().(Hyperlink)
//		if !localVarOptionalHyperlinkok {
//				return localVarReturnValue, nil, reportError("hyperlink should be Hyperlink")
//		}
//		localVarPostBody = &localVarOptionalHyperlink
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Hyperlink != nil {
		
		localVarPostBody = &localVarOptionals.Hyperlink
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Add worksheet hyperlink.
 * @param name Document name.
 * @param sheetName Worksheet name.
 * @param firstRow
 * @param firstColumn
 * @param totalRows
 * @param totalColumns
 * @param address
 * @param optional nil or *CellsHypelinksPutWorksheetHyperlinkOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return HyperlinkResponse
*/


type CellsHypelinksPutWorksheetHyperlinkOpts struct { 
	Name string
	SheetName string
	FirstRow int32
	FirstColumn int32
	TotalRows int32
	TotalColumns int32
	Address string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsHypelinksPutWorksheetHyperlink(    localVarOptionals *CellsHypelinksPutWorksheetHyperlinkOpts) (HyperlinkResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue HyperlinkResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/hyperlinks"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("firstRow", parameterToString(localVarOptionals.FirstRow, ""))
	localVarQueryParams.Add("firstColumn", parameterToString(localVarOptionals.FirstColumn, ""))
	localVarQueryParams.Add("totalRows", parameterToString(localVarOptionals.TotalRows, ""))
	localVarQueryParams.Add("totalColumns", parameterToString(localVarOptionals.TotalColumns, ""))
	localVarQueryParams.Add("address", parameterToString(localVarOptionals.Address, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Delete worksheet list object by index
 * @param name Document name.
 * @param sheetName The worksheet name.
 * @param listObjectIndex List object index
 * @param optional nil or *CellsListObjectsDeleteWorksheetListObjectOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsListObjectsDeleteWorksheetListObjectOpts struct { 
	Name string
	SheetName string
	ListObjectIndex int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsListObjectsDeleteWorksheetListObject(    localVarOptionals *CellsListObjectsDeleteWorksheetListObjectOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/listobjects/{listObjectIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"listObjectIndex"+"}", fmt.Sprintf("%v", localVarOptionals.ListObjectIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Delete worksheet list objects
 * @param name Document name.
 * @param sheetName The worksheet name.
 * @param optional nil or *CellsListObjectsDeleteWorksheetListObjectsOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsListObjectsDeleteWorksheetListObjectsOpts struct { 
	Name string
	SheetName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsListObjectsDeleteWorksheetListObjects(    localVarOptionals *CellsListObjectsDeleteWorksheetListObjectsOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/listobjects"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get worksheet list object info by index.
 * @param name Document name.
 * @param sheetName The worksheet name.
 * @param listobjectindex list object index.
 * @param optional nil or *CellsListObjectsGetWorksheetListObjectOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return ListObjectResponse
*/


type CellsListObjectsGetWorksheetListObjectOpts struct { 
	Name string
	SheetName string
	Listobjectindex int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsListObjectsGetWorksheetListObject(    localVarOptionals *CellsListObjectsGetWorksheetListObjectOpts) (ListObjectResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ListObjectResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/listobjects/{listobjectindex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"listobjectindex"+"}", fmt.Sprintf("%v", localVarOptionals.Listobjectindex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get worksheet listobjects info.
 * @param name Document name.
 * @param sheetName The worksheet name.
 * @param optional nil or *CellsListObjectsGetWorksheetListObjectsOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return ListObjectsResponse
*/


type CellsListObjectsGetWorksheetListObjectsOpts struct { 
	Name string
	SheetName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsListObjectsGetWorksheetListObjects(    localVarOptionals *CellsListObjectsGetWorksheetListObjectsOpts) (ListObjectsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ListObjectsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/listobjects"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Update  list object 
 * @param name Document name.
 * @param sheetName The worksheet name.
 * @param listObjectIndex list Object index
 * @param optional nil or *CellsListObjectsPostWorksheetListObjectOpts - Optional Parameters:
     * @param "ListObject" (optional.Interface of ListObject) -  listObject dto in request body.
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsListObjectsPostWorksheetListObjectOpts struct { 
	Name string
	SheetName string
	ListObjectIndex int32
	ListObject *ListObject
	Folder string
	Storage string
}


func (a *CellsApiService) CellsListObjectsPostWorksheetListObject(    localVarOptionals *CellsListObjectsPostWorksheetListObjectOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/listobjects/{listObjectIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"listObjectIndex"+"}", fmt.Sprintf("%v", localVarOptionals.ListObjectIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.ListObject.IsSet() {
//		
//		localVarOptionalListObject, localVarOptionalListObjectok := localVarOptionals.ListObject.Value().(ListObject)
//		if !localVarOptionalListObjectok {
//				return localVarReturnValue, nil, reportError("listObject should be ListObject")
//		}
//		localVarPostBody = &localVarOptionalListObject
//	}
	if localVarOptionals != nil &&  &localVarOptionals.ListObject != nil {
		
		localVarPostBody = &localVarOptionals.ListObject
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService
 * @param name
 * @param sheetName
 * @param listObjectIndex
 * @param optional nil or *CellsListObjectsPostWorksheetListObjectConvertToRangeOpts - Optional Parameters:
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsListObjectsPostWorksheetListObjectConvertToRangeOpts struct { 
	Name string
	SheetName string
	ListObjectIndex int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsListObjectsPostWorksheetListObjectConvertToRange(    localVarOptionals *CellsListObjectsPostWorksheetListObjectConvertToRangeOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/listobjects/{listObjectIndex}/ConvertToRange"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"listObjectIndex"+"}", fmt.Sprintf("%v", localVarOptionals.ListObjectIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService
 * @param name
 * @param sheetName
 * @param listObjectIndex
 * @param optional nil or *CellsListObjectsPostWorksheetListObjectSortTableOpts - Optional Parameters:
     * @param "DataSorter" (optional.Interface of DataSorter) - 
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsListObjectsPostWorksheetListObjectSortTableOpts struct { 
	Name string
	SheetName string
	ListObjectIndex int32
	DataSorter *DataSorter
	Folder string
	Storage string
}


func (a *CellsApiService) CellsListObjectsPostWorksheetListObjectSortTable(    localVarOptionals *CellsListObjectsPostWorksheetListObjectSortTableOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/listobjects/{listObjectIndex}/sort"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"listObjectIndex"+"}", fmt.Sprintf("%v", localVarOptionals.ListObjectIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.DataSorter.IsSet() {
//		
//		localVarOptionalDataSorter, localVarOptionalDataSorterok := localVarOptionals.DataSorter.Value().(DataSorter)
//		if !localVarOptionalDataSorterok {
//				return localVarReturnValue, nil, reportError("dataSorter should be DataSorter")
//		}
//		localVarPostBody = &localVarOptionalDataSorter
//	}
	if localVarOptionals != nil &&  &localVarOptionals.DataSorter != nil {
		
		localVarPostBody = &localVarOptionals.DataSorter
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService
 * @param name
 * @param sheetName
 * @param listObjectIndex
 * @param destsheetName
 * @param optional nil or *CellsListObjectsPostWorksheetListObjectSummarizeWithPivotTableOpts - Optional Parameters:
     * @param "Request" (optional.Interface of CreatePivotTableRequest) - 
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsListObjectsPostWorksheetListObjectSummarizeWithPivotTableOpts struct { 
	Name string
	SheetName string
	ListObjectIndex int32
	DestsheetName string
	Request *CreatePivotTableRequest
	Folder string
	Storage string
}


func (a *CellsApiService) CellsListObjectsPostWorksheetListObjectSummarizeWithPivotTable(    localVarOptionals *CellsListObjectsPostWorksheetListObjectSummarizeWithPivotTableOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/listobjects/{listObjectIndex}/SummarizeWithPivotTable"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"listObjectIndex"+"}", fmt.Sprintf("%v", localVarOptionals.ListObjectIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("destsheetName", parameterToString(localVarOptionals.DestsheetName, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Request.IsSet() {
//		
//		localVarOptionalRequest, localVarOptionalRequestok := localVarOptionals.Request.Value().(CreatePivotTableRequest)
//		if !localVarOptionalRequestok {
//				return localVarReturnValue, nil, reportError("request should be CreatePivotTableRequest")
//		}
//		localVarPostBody = &localVarOptionalRequest
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Request != nil {
		
		localVarPostBody = &localVarOptionals.Request
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Add a list object into worksheet.
 * @param name Document name.
 * @param sheetName The worksheet name.
 * @param startRow The start row of the list range.
 * @param startColumn The start row of the list range.
 * @param endRow The start row of the list range.
 * @param endColumn The start row of the list range.
 * @param optional nil or *CellsListObjectsPutWorksheetListObjectOpts - Optional Parameters:
     * @param "HasHeaders" (optional.Bool) -  Whether the range has headers.
     * @param "ListObject" (optional.Interface of ListObject) -  List Object
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return ListObjectResponse
*/


type CellsListObjectsPutWorksheetListObjectOpts struct { 
	Name string
	SheetName string
	StartRow int32
	StartColumn int32
	EndRow int32
	EndColumn int32
	HasHeaders bool
	ListObject *ListObject
	Folder string
	Storage string
}


func (a *CellsApiService) CellsListObjectsPutWorksheetListObject(    localVarOptionals *CellsListObjectsPutWorksheetListObjectOpts) (ListObjectResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ListObjectResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/listobjects"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("startRow", parameterToString(localVarOptionals.StartRow, ""))
	localVarQueryParams.Add("startColumn", parameterToString(localVarOptionals.StartColumn, ""))
	localVarQueryParams.Add("endRow", parameterToString(localVarOptionals.EndRow, ""))
	localVarQueryParams.Add("endColumn", parameterToString(localVarOptionals.EndColumn, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("hasHeaders", parameterToString(localVarOptionals.HasHeaders.Value(), ""))
		localVarQueryParams.Add("hasHeaders", parameterToString(localVarOptionals.HasHeaders, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.ListObject.IsSet() {
//		
//		localVarOptionalListObject, localVarOptionalListObjectok := localVarOptionals.ListObject.Value().(ListObject)
//		if !localVarOptionalListObjectok {
//				return localVarReturnValue, nil, reportError("listObject should be ListObject")
//		}
//		localVarPostBody = &localVarOptionalListObject
//	}
	if localVarOptionals != nil &&  &localVarOptionals.ListObject != nil {
		
		localVarPostBody = &localVarOptionals.ListObject
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Delete OLE object.
 * @param name The workbook name.
 * @param sheetName The worsheet name.
 * @param oleObjectIndex Ole object index
 * @param optional nil or *CellsOleObjectsDeleteWorksheetOleObjectOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsOleObjectsDeleteWorksheetOleObjectOpts struct { 
	Name string
	SheetName string
	OleObjectIndex int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsOleObjectsDeleteWorksheetOleObject(    localVarOptionals *CellsOleObjectsDeleteWorksheetOleObjectOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/oleobjects/{oleObjectIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"oleObjectIndex"+"}", fmt.Sprintf("%v", localVarOptionals.OleObjectIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Delete all OLE objects.
 * @param name The workbook name.
 * @param sheetName The worsheet name.
 * @param optional nil or *CellsOleObjectsDeleteWorksheetOleObjectsOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsOleObjectsDeleteWorksheetOleObjectsOpts struct { 
	Name string
	SheetName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsOleObjectsDeleteWorksheetOleObjects(    localVarOptionals *CellsOleObjectsDeleteWorksheetOleObjectsOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/oleobjects"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get OLE object info.
 * @param name Document name.
 * @param sheetName Worksheet name.
 * @param objectNumber The object number.
 * @param optional nil or *CellsOleObjectsGetWorksheetOleObjectOpts - Optional Parameters:
     * @param "Format" (optional.String) -  The exported object format.
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return *os.File
*/


type CellsOleObjectsGetWorksheetOleObjectOpts struct { 
	Name string
	SheetName string
	ObjectNumber int32
	Format string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsOleObjectsGetWorksheetOleObject(    localVarOptionals *CellsOleObjectsGetWorksheetOleObjectOpts) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/oleobjects/{objectNumber}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"objectNumber"+"}", fmt.Sprintf("%v", localVarOptionals.ObjectNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("format", parameterToString(localVarOptionals.Format.Value(), ""))
		localVarQueryParams.Add("format", parameterToString(localVarOptionals.Format, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get worksheet OLE objects info.
 * @param name Document name.
 * @param sheetName The worksheet name.
 * @param optional nil or *CellsOleObjectsGetWorksheetOleObjectsOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return OleObjectsResponse
*/


type CellsOleObjectsGetWorksheetOleObjectsOpts struct { 
	Name string
	SheetName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsOleObjectsGetWorksheetOleObjects(    localVarOptionals *CellsOleObjectsGetWorksheetOleObjectsOpts) (OleObjectsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue OleObjectsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/oleobjects"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Update OLE object.
 * @param name The workbook name.
 * @param sheetName The worsheet name.
 * @param oleObjectIndex Ole object index
 * @param optional nil or *CellsOleObjectsPostUpdateWorksheetOleObjectOpts - Optional Parameters:
     * @param "Ole" (optional.Interface of OleObject) -  Ole Object
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsOleObjectsPostUpdateWorksheetOleObjectOpts struct { 
	Name string
	SheetName string
	OleObjectIndex int32
	Ole *OleObject
	Folder string
	Storage string
}


func (a *CellsApiService) CellsOleObjectsPostUpdateWorksheetOleObject(    localVarOptionals *CellsOleObjectsPostUpdateWorksheetOleObjectOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/oleobjects/{oleObjectIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"oleObjectIndex"+"}", fmt.Sprintf("%v", localVarOptionals.OleObjectIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Ole.IsSet() {
//		
//		localVarOptionalOle, localVarOptionalOleok := localVarOptionals.Ole.Value().(OleObject)
//		if !localVarOptionalOleok {
//				return localVarReturnValue, nil, reportError("ole should be OleObject")
//		}
//		localVarPostBody = &localVarOptionalOle
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Ole != nil {
		
		localVarPostBody = &localVarOptionals.Ole
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Add OLE object
 * @param name The workbook name.
 * @param sheetName The worsheet name.
 * @param optional nil or *CellsOleObjectsPutWorksheetOleObjectOpts - Optional Parameters:
     * @param "OleObject" (optional.Interface of OleObject) -  Ole Object
     * @param "UpperLeftRow" (optional.Int32) -  Upper left row index
     * @param "UpperLeftColumn" (optional.Int32) -  Upper left column index
     * @param "Height" (optional.Int32) -  Height of oleObject, in unit of pixel
     * @param "Width" (optional.Int32) -  Width of oleObject, in unit of pixel
     * @param "OleFile" (optional.String) -  OLE filename
     * @param "ImageFile" (optional.String) -  Image filename
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return OleObjectResponse
*/


type CellsOleObjectsPutWorksheetOleObjectOpts struct { 
	Name string
	SheetName string
	OleObject *OleObject
	UpperLeftRow int32
	UpperLeftColumn int32
	Height int32
	Width int32
	OleFile string
	ImageFile string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsOleObjectsPutWorksheetOleObject(    localVarOptionals *CellsOleObjectsPutWorksheetOleObjectOpts) (OleObjectResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue OleObjectResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/oleobjects"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("upperLeftRow", parameterToString(localVarOptionals.UpperLeftRow.Value(), ""))
		localVarQueryParams.Add("upperLeftRow", parameterToString(localVarOptionals.UpperLeftRow, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("upperLeftColumn", parameterToString(localVarOptionals.UpperLeftColumn.Value(), ""))
		localVarQueryParams.Add("upperLeftColumn", parameterToString(localVarOptionals.UpperLeftColumn, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("height", parameterToString(localVarOptionals.Height.Value(), ""))
		localVarQueryParams.Add("height", parameterToString(localVarOptionals.Height, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("width", parameterToString(localVarOptionals.Width.Value(), ""))
		localVarQueryParams.Add("width", parameterToString(localVarOptionals.Width, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("oleFile", parameterToString(localVarOptionals.OleFile.Value(), ""))
		localVarQueryParams.Add("oleFile", parameterToString(localVarOptionals.OleFile, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("imageFile", parameterToString(localVarOptionals.ImageFile.Value(), ""))
		localVarQueryParams.Add("imageFile", parameterToString(localVarOptionals.ImageFile, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.OleObject.IsSet() {
//		
//		localVarOptionalOleObject, localVarOptionalOleObjectok := localVarOptionals.OleObject.Value().(OleObject)
//		if !localVarOptionalOleObjectok {
//				return localVarReturnValue, nil, reportError("oleObject should be OleObject")
//		}
//		localVarPostBody = &localVarOptionalOleObject
//	}
	if localVarOptionals != nil &&  &localVarOptionals.OleObject != nil {
		
		localVarPostBody = &localVarOptionals.OleObject
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService
 * @param name
 * @param sheetName
 * @param index
 * @param optional nil or *CellsPageBreaksDeleteHorizontalPageBreakOpts - Optional Parameters:
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPageBreaksDeleteHorizontalPageBreakOpts struct { 
	Name string
	SheetName string
	Index int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPageBreaksDeleteHorizontalPageBreak(    localVarOptionals *CellsPageBreaksDeleteHorizontalPageBreakOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/horizontalpagebreaks/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", localVarOptionals.Index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService
 * @param name
 * @param sheetName
 * @param optional nil or *CellsPageBreaksDeleteHorizontalPageBreaksOpts - Optional Parameters:
     * @param "Row" (optional.Int32) - 
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPageBreaksDeleteHorizontalPageBreaksOpts struct { 
	Name string
	SheetName string
	Row int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPageBreaksDeleteHorizontalPageBreaks(    localVarOptionals *CellsPageBreaksDeleteHorizontalPageBreaksOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/horizontalpagebreaks"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("row", parameterToString(localVarOptionals.Row.Value(), ""))
		localVarQueryParams.Add("row", parameterToString(localVarOptionals.Row, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService
 * @param name
 * @param sheetName
 * @param index
 * @param optional nil or *CellsPageBreaksDeleteVerticalPageBreakOpts - Optional Parameters:
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPageBreaksDeleteVerticalPageBreakOpts struct { 
	Name string
	SheetName string
	Index int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPageBreaksDeleteVerticalPageBreak(    localVarOptionals *CellsPageBreaksDeleteVerticalPageBreakOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/verticalpagebreaks/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", localVarOptionals.Index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService
 * @param name
 * @param sheetName
 * @param optional nil or *CellsPageBreaksDeleteVerticalPageBreaksOpts - Optional Parameters:
     * @param "Column" (optional.Int32) - 
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPageBreaksDeleteVerticalPageBreaksOpts struct { 
	Name string
	SheetName string
	Column int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPageBreaksDeleteVerticalPageBreaks(    localVarOptionals *CellsPageBreaksDeleteVerticalPageBreaksOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/verticalpagebreaks"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("column", parameterToString(localVarOptionals.Column.Value(), ""))
		localVarQueryParams.Add("column", parameterToString(localVarOptionals.Column, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService
 * @param name
 * @param sheetName
 * @param index
 * @param optional nil or *CellsPageBreaksGetHorizontalPageBreakOpts - Optional Parameters:
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return HorizontalPageBreakResponse
*/


type CellsPageBreaksGetHorizontalPageBreakOpts struct { 
	Name string
	SheetName string
	Index int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPageBreaksGetHorizontalPageBreak(    localVarOptionals *CellsPageBreaksGetHorizontalPageBreakOpts) (HorizontalPageBreakResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue HorizontalPageBreakResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/horizontalpagebreaks/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", localVarOptionals.Index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService
 * @param name
 * @param sheetName
 * @param optional nil or *CellsPageBreaksGetHorizontalPageBreaksOpts - Optional Parameters:
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return HorizontalPageBreaksResponse
*/


type CellsPageBreaksGetHorizontalPageBreaksOpts struct { 
	Name string
	SheetName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPageBreaksGetHorizontalPageBreaks(    localVarOptionals *CellsPageBreaksGetHorizontalPageBreaksOpts) (HorizontalPageBreaksResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue HorizontalPageBreaksResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/horizontalpagebreaks"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService
 * @param name
 * @param sheetName
 * @param index
 * @param optional nil or *CellsPageBreaksGetVerticalPageBreakOpts - Optional Parameters:
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return VerticalPageBreakResponse
*/


type CellsPageBreaksGetVerticalPageBreakOpts struct { 
	Name string
	SheetName string
	Index int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPageBreaksGetVerticalPageBreak(    localVarOptionals *CellsPageBreaksGetVerticalPageBreakOpts) (VerticalPageBreakResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue VerticalPageBreakResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/verticalpagebreaks/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", fmt.Sprintf("%v", localVarOptionals.Index), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService
 * @param name
 * @param sheetName
 * @param optional nil or *CellsPageBreaksGetVerticalPageBreaksOpts - Optional Parameters:
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return VerticalPageBreaksResponse
*/


type CellsPageBreaksGetVerticalPageBreaksOpts struct { 
	Name string
	SheetName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPageBreaksGetVerticalPageBreaks(    localVarOptionals *CellsPageBreaksGetVerticalPageBreaksOpts) (VerticalPageBreaksResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue VerticalPageBreaksResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/verticalpagebreaks"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService
 * @param name
 * @param sheetName
 * @param optional nil or *CellsPageBreaksPutHorizontalPageBreakOpts - Optional Parameters:
     * @param "Cellname" (optional.String) - 
     * @param "Row" (optional.Int32) - 
     * @param "Column" (optional.Int32) - 
     * @param "StartColumn" (optional.Int32) - 
     * @param "EndColumn" (optional.Int32) - 
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPageBreaksPutHorizontalPageBreakOpts struct { 
	Name string
	SheetName string
	Cellname string
	Row int32
	Column int32
	StartColumn int32
	EndColumn int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPageBreaksPutHorizontalPageBreak(    localVarOptionals *CellsPageBreaksPutHorizontalPageBreakOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/horizontalpagebreaks"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("cellname", parameterToString(localVarOptionals.Cellname.Value(), ""))
		localVarQueryParams.Add("cellname", parameterToString(localVarOptionals.Cellname, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("row", parameterToString(localVarOptionals.Row.Value(), ""))
		localVarQueryParams.Add("row", parameterToString(localVarOptionals.Row, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("column", parameterToString(localVarOptionals.Column.Value(), ""))
		localVarQueryParams.Add("column", parameterToString(localVarOptionals.Column, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("startColumn", parameterToString(localVarOptionals.StartColumn.Value(), ""))
		localVarQueryParams.Add("startColumn", parameterToString(localVarOptionals.StartColumn, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("endColumn", parameterToString(localVarOptionals.EndColumn.Value(), ""))
		localVarQueryParams.Add("endColumn", parameterToString(localVarOptionals.EndColumn, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService
 * @param name
 * @param sheetName
 * @param optional nil or *CellsPageBreaksPutVerticalPageBreakOpts - Optional Parameters:
     * @param "Cellname" (optional.String) - 
     * @param "Column" (optional.Int32) - 
     * @param "Row" (optional.Int32) - 
     * @param "StartRow" (optional.Int32) - 
     * @param "EndRow" (optional.Int32) - 
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPageBreaksPutVerticalPageBreakOpts struct { 
	Name string
	SheetName string
	Cellname string
	Column int32
	Row int32
	StartRow int32
	EndRow int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPageBreaksPutVerticalPageBreak(    localVarOptionals *CellsPageBreaksPutVerticalPageBreakOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/verticalpagebreaks"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("cellname", parameterToString(localVarOptionals.Cellname.Value(), ""))
		localVarQueryParams.Add("cellname", parameterToString(localVarOptionals.Cellname, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("column", parameterToString(localVarOptionals.Column.Value(), ""))
		localVarQueryParams.Add("column", parameterToString(localVarOptionals.Column, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("row", parameterToString(localVarOptionals.Row.Value(), ""))
		localVarQueryParams.Add("row", parameterToString(localVarOptionals.Row, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("startRow", parameterToString(localVarOptionals.StartRow.Value(), ""))
		localVarQueryParams.Add("startRow", parameterToString(localVarOptionals.StartRow, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("endRow", parameterToString(localVarOptionals.EndRow.Value(), ""))
		localVarQueryParams.Add("endRow", parameterToString(localVarOptionals.EndRow, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService clear header footer
 * @param name
 * @param sheetName
 * @param optional nil or *CellsPageSetupDeleteHeaderFooterOpts - Optional Parameters:
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPageSetupDeleteHeaderFooterOpts struct { 
	Name string
	SheetName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPageSetupDeleteHeaderFooter(    localVarOptionals *CellsPageSetupDeleteHeaderFooterOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/pagesetup/clearheaderfooter"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService get page footer information
 * @param name
 * @param sheetName
 * @param optional nil or *CellsPageSetupGetFooterOpts - Optional Parameters:
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return PageSectionsResponse
*/


type CellsPageSetupGetFooterOpts struct { 
	Name string
	SheetName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPageSetupGetFooter(    localVarOptionals *CellsPageSetupGetFooterOpts) (PageSectionsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue PageSectionsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/pagesetup/footer"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService get page header information
 * @param name
 * @param sheetName
 * @param optional nil or *CellsPageSetupGetHeaderOpts - Optional Parameters:
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return PageSectionsResponse
*/


type CellsPageSetupGetHeaderOpts struct { 
	Name string
	SheetName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPageSetupGetHeader(    localVarOptionals *CellsPageSetupGetHeaderOpts) (PageSectionsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue PageSectionsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/pagesetup/header"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get Page Setup information.             
 * @param name
 * @param sheetName
 * @param optional nil or *CellsPageSetupGetPageSetupOpts - Optional Parameters:
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return PageSetupResponse
*/


type CellsPageSetupGetPageSetupOpts struct { 
	Name string
	SheetName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPageSetupGetPageSetup(    localVarOptionals *CellsPageSetupGetPageSetupOpts) (PageSetupResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue PageSetupResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/pagesetup"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService update  page footer information 
 * @param name
 * @param sheetName
 * @param section
 * @param script
 * @param isFirstPage
 * @param optional nil or *CellsPageSetupPostFooterOpts - Optional Parameters:
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPageSetupPostFooterOpts struct { 
	Name string
	SheetName string
	Section int32
	Script string
	IsFirstPage bool
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPageSetupPostFooter(    localVarOptionals *CellsPageSetupPostFooterOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/pagesetup/footer"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("section", parameterToString(localVarOptionals.Section, ""))
	localVarQueryParams.Add("script", parameterToString(localVarOptionals.Script, ""))
	localVarQueryParams.Add("isFirstPage", parameterToString(localVarOptionals.IsFirstPage, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService update  page header information 
 * @param name
 * @param sheetName
 * @param section
 * @param script
 * @param isFirstPage
 * @param optional nil or *CellsPageSetupPostHeaderOpts - Optional Parameters:
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPageSetupPostHeaderOpts struct { 
	Name string
	SheetName string
	Section int32
	Script string
	IsFirstPage bool
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPageSetupPostHeader(    localVarOptionals *CellsPageSetupPostHeaderOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/pagesetup/header"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("section", parameterToString(localVarOptionals.Section, ""))
	localVarQueryParams.Add("script", parameterToString(localVarOptionals.Script, ""))
	localVarQueryParams.Add("isFirstPage", parameterToString(localVarOptionals.IsFirstPage, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Update Page Setup information.
 * @param name
 * @param sheetName
 * @param optional nil or *CellsPageSetupPostPageSetupOpts - Optional Parameters:
     * @param "PageSetup" (optional.Interface of PageSetup) - 
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPageSetupPostPageSetupOpts struct { 
	Name string
	SheetName string
	PageSetup *PageSetup
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPageSetupPostPageSetup(    localVarOptionals *CellsPageSetupPostPageSetupOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/pagesetup"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.PageSetup.IsSet() {
//		
//		localVarOptionalPageSetup, localVarOptionalPageSetupok := localVarOptionals.PageSetup.Value().(PageSetup)
//		if !localVarOptionalPageSetupok {
//				return localVarReturnValue, nil, reportError("pageSetup should be PageSetup")
//		}
//		localVarPostBody = &localVarOptionalPageSetup
//	}
	if localVarOptionals != nil &&  &localVarOptionals.PageSetup != nil {
		
		localVarPostBody = &localVarOptionals.PageSetup
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Delete a picture object in worksheet
 * @param name The workbook name.
 * @param sheetName The worsheet name.
 * @param pictureIndex Picture index
 * @param optional nil or *CellsPicturesDeleteWorksheetPictureOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPicturesDeleteWorksheetPictureOpts struct { 
	Name string
	SheetName string
	PictureIndex int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPicturesDeleteWorksheetPicture(    localVarOptionals *CellsPicturesDeleteWorksheetPictureOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/pictures/{pictureIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pictureIndex"+"}", fmt.Sprintf("%v", localVarOptionals.PictureIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Delete all pictures in worksheet.
 * @param name Document name.
 * @param sheetName Worksheet name.
 * @param optional nil or *CellsPicturesDeleteWorksheetPicturesOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPicturesDeleteWorksheetPicturesOpts struct { 
	Name string
	SheetName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPicturesDeleteWorksheetPictures(    localVarOptionals *CellsPicturesDeleteWorksheetPicturesOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/pictures"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService GRead worksheet picture by number.
 * @param name Document name.
 * @param sheetName Worksheet name.
 * @param pictureIndex The picture index.
 * @param optional nil or *CellsPicturesGetWorksheetPictureOpts - Optional Parameters:
     * @param "Format" (optional.String) -  The exported object format.
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return *os.File
*/


type CellsPicturesGetWorksheetPictureOpts struct { 
	Name string
	SheetName string
	PictureIndex int32
	Format string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPicturesGetWorksheetPicture(    localVarOptionals *CellsPicturesGetWorksheetPictureOpts) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/pictures/{pictureIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pictureIndex"+"}", fmt.Sprintf("%v", localVarOptionals.PictureIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("format", parameterToString(localVarOptionals.Format.Value(), ""))
		localVarQueryParams.Add("format", parameterToString(localVarOptionals.Format, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Read worksheet pictures.
 * @param name Document name.
 * @param sheetName The worksheet name.
 * @param optional nil or *CellsPicturesGetWorksheetPicturesOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return PicturesResponse
*/


type CellsPicturesGetWorksheetPicturesOpts struct { 
	Name string
	SheetName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPicturesGetWorksheetPictures(    localVarOptionals *CellsPicturesGetWorksheetPicturesOpts) (PicturesResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue PicturesResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/pictures"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Update worksheet picture by index.
 * @param name Document name.
 * @param sheetName Worksheet name.
 * @param pictureIndex The picture&#39;s index.
 * @param optional nil or *CellsPicturesPostWorksheetPictureOpts - Optional Parameters:
     * @param "Picture" (optional.Interface of Picture) -  Picture object
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return PictureResponse
*/


type CellsPicturesPostWorksheetPictureOpts struct { 
	Name string
	SheetName string
	PictureIndex int32
	Picture *Picture
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPicturesPostWorksheetPicture(    localVarOptionals *CellsPicturesPostWorksheetPictureOpts) (PictureResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue PictureResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/pictures/{pictureIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pictureIndex"+"}", fmt.Sprintf("%v", localVarOptionals.PictureIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Picture.IsSet() {
//		
//		localVarOptionalPicture, localVarOptionalPictureok := localVarOptionals.Picture.Value().(Picture)
//		if !localVarOptionalPictureok {
//				return localVarReturnValue, nil, reportError("picture should be Picture")
//		}
//		localVarPostBody = &localVarOptionalPicture
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Picture != nil {
		
		localVarPostBody = &localVarOptionals.Picture
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Add a new worksheet picture.
 * @param name The workbook name.
 * @param sheetName The worsheet name.
 * @param optional nil or *CellsPicturesPutWorksheetAddPictureOpts - Optional Parameters:
     * @param "Picture" (optional.Interface of Picture) -  Pictute object
     * @param "UpperLeftRow" (optional.Int32) -  The image upper left row.
     * @param "UpperLeftColumn" (optional.Int32) -  The image upper left column.
     * @param "LowerRightRow" (optional.Int32) -  The image low right row.
     * @param "LowerRightColumn" (optional.Int32) -  The image low right column.
     * @param "PicturePath" (optional.String) -  The picture path, if not provided the picture data is inspected in the request body.
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return PicturesResponse
*/


type CellsPicturesPutWorksheetAddPictureOpts struct { 
	Name string
	SheetName string
	Picture *Picture
	UpperLeftRow int32
	UpperLeftColumn int32
	LowerRightRow int32
	LowerRightColumn int32
	PicturePath string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPicturesPutWorksheetAddPicture(    localVarOptionals *CellsPicturesPutWorksheetAddPictureOpts) (PicturesResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue PicturesResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/pictures"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("upperLeftRow", parameterToString(localVarOptionals.UpperLeftRow.Value(), ""))
		localVarQueryParams.Add("upperLeftRow", parameterToString(localVarOptionals.UpperLeftRow, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("upperLeftColumn", parameterToString(localVarOptionals.UpperLeftColumn.Value(), ""))
		localVarQueryParams.Add("upperLeftColumn", parameterToString(localVarOptionals.UpperLeftColumn, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("lowerRightRow", parameterToString(localVarOptionals.LowerRightRow.Value(), ""))
		localVarQueryParams.Add("lowerRightRow", parameterToString(localVarOptionals.LowerRightRow, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("lowerRightColumn", parameterToString(localVarOptionals.LowerRightColumn.Value(), ""))
		localVarQueryParams.Add("lowerRightColumn", parameterToString(localVarOptionals.LowerRightColumn, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("picturePath", parameterToString(localVarOptionals.PicturePath.Value(), ""))
		localVarQueryParams.Add("picturePath", parameterToString(localVarOptionals.PicturePath, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Picture.IsSet() {
//		
//		localVarOptionalPicture, localVarOptionalPictureok := localVarOptionals.Picture.Value().(Picture)
//		if !localVarOptionalPictureok {
//				return localVarReturnValue, nil, reportError("picture should be Picture")
//		}
//		localVarPostBody = &localVarOptionalPicture
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Picture != nil {
		
		localVarPostBody = &localVarOptionals.Picture
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Delete pivot field into into pivot table
 * @param name Document name.
 * @param sheetName The worksheet name.
 * @param pivotTableIndex Pivot table index
 * @param pivotFieldType The fields area type.
 * @param optional nil or *CellsPivotTablesDeletePivotTableFieldOpts - Optional Parameters:
     * @param "Request" (optional.Interface of PivotTableFieldRequest) -  Dto that conrains field indexes
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPivotTablesDeletePivotTableFieldOpts struct { 
	Name string
	SheetName string
	PivotTableIndex int32
	PivotFieldType string
	Request *PivotTableFieldRequest
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPivotTablesDeletePivotTableField(    localVarOptionals *CellsPivotTablesDeletePivotTableFieldOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/PivotField"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pivotTableIndex"+"}", fmt.Sprintf("%v", localVarOptionals.PivotTableIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("pivotFieldType", parameterToString(localVarOptionals.PivotFieldType, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Request.IsSet() {
//		
//		localVarOptionalRequest, localVarOptionalRequestok := localVarOptionals.Request.Value().(PivotTableFieldRequest)
//		if !localVarOptionalRequestok {
//				return localVarReturnValue, nil, reportError("request should be PivotTableFieldRequest")
//		}
//		localVarPostBody = &localVarOptionalRequest
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Request != nil {
		
		localVarPostBody = &localVarOptionals.Request
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Delete worksheet pivot table by index
 * @param name Document name.
 * @param sheetName The worksheet name.
 * @param pivotTableIndex Pivot table index
 * @param optional nil or *CellsPivotTablesDeleteWorksheetPivotTableOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPivotTablesDeleteWorksheetPivotTableOpts struct { 
	Name string
	SheetName string
	PivotTableIndex int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPivotTablesDeleteWorksheetPivotTable(    localVarOptionals *CellsPivotTablesDeleteWorksheetPivotTableOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pivotTableIndex"+"}", fmt.Sprintf("%v", localVarOptionals.PivotTableIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService delete  pivot filter for piovt table             
 * @param name
 * @param sheetName
 * @param pivotTableIndex
 * @param fieldIndex
 * @param optional nil or *CellsPivotTablesDeleteWorksheetPivotTableFilterOpts - Optional Parameters:
     * @param "NeedReCalculate" (optional.Bool) - 
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPivotTablesDeleteWorksheetPivotTableFilterOpts struct { 
	Name string
	SheetName string
	PivotTableIndex int32
	FieldIndex int32
	NeedReCalculate bool
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPivotTablesDeleteWorksheetPivotTableFilter(    localVarOptionals *CellsPivotTablesDeleteWorksheetPivotTableFilterOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/PivotFilters/{fieldIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pivotTableIndex"+"}", fmt.Sprintf("%v", localVarOptionals.PivotTableIndex), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"fieldIndex"+"}", fmt.Sprintf("%v", localVarOptionals.FieldIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("needReCalculate", parameterToString(localVarOptionals.NeedReCalculate.Value(), ""))
		localVarQueryParams.Add("needReCalculate", parameterToString(localVarOptionals.NeedReCalculate, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService delete all pivot filters for piovt table
 * @param name
 * @param sheetName
 * @param pivotTableIndex
 * @param optional nil or *CellsPivotTablesDeleteWorksheetPivotTableFiltersOpts - Optional Parameters:
     * @param "NeedReCalculate" (optional.Bool) - 
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPivotTablesDeleteWorksheetPivotTableFiltersOpts struct { 
	Name string
	SheetName string
	PivotTableIndex int32
	NeedReCalculate bool
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPivotTablesDeleteWorksheetPivotTableFilters(    localVarOptionals *CellsPivotTablesDeleteWorksheetPivotTableFiltersOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/PivotFilters"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pivotTableIndex"+"}", fmt.Sprintf("%v", localVarOptionals.PivotTableIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("needReCalculate", parameterToString(localVarOptionals.NeedReCalculate.Value(), ""))
		localVarQueryParams.Add("needReCalculate", parameterToString(localVarOptionals.NeedReCalculate, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Delete worksheet pivot tables
 * @param name Document name.
 * @param sheetName The worksheet name.
 * @param optional nil or *CellsPivotTablesDeleteWorksheetPivotTablesOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPivotTablesDeleteWorksheetPivotTablesOpts struct { 
	Name string
	SheetName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPivotTablesDeleteWorksheetPivotTables(    localVarOptionals *CellsPivotTablesDeleteWorksheetPivotTablesOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/pivottables"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get pivot field into into pivot table
 * @param name Document name.
 * @param sheetName The worksheet name.
 * @param pivotTableIndex Pivot table index
 * @param pivotFieldIndex The field index in the base fields.
 * @param pivotFieldType The fields area type.
 * @param optional nil or *CellsPivotTablesGetPivotTableFieldOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return PivotFieldResponse
*/


type CellsPivotTablesGetPivotTableFieldOpts struct { 
	Name string
	SheetName string
	PivotTableIndex int32
	PivotFieldIndex int32
	PivotFieldType string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPivotTablesGetPivotTableField(    localVarOptionals *CellsPivotTablesGetPivotTableFieldOpts) (PivotFieldResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue PivotFieldResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/PivotField"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pivotTableIndex"+"}", fmt.Sprintf("%v", localVarOptionals.PivotTableIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("pivotFieldIndex", parameterToString(localVarOptionals.PivotFieldIndex, ""))
	localVarQueryParams.Add("pivotFieldType", parameterToString(localVarOptionals.PivotFieldType, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get worksheet pivottable info by index.
 * @param name Document name.
 * @param sheetName The worksheet name.
 * @param pivottableIndex
 * @param optional nil or *CellsPivotTablesGetWorksheetPivotTableOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return PivotTableResponse
*/


type CellsPivotTablesGetWorksheetPivotTableOpts struct { 
	Name string
	SheetName string
	PivottableIndex int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPivotTablesGetWorksheetPivotTable(    localVarOptionals *CellsPivotTablesGetWorksheetPivotTableOpts) (PivotTableResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue PivotTableResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/pivottables/{pivottableIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pivottableIndex"+"}", fmt.Sprintf("%v", localVarOptionals.PivottableIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService
 * @param name
 * @param sheetName
 * @param pivotTableIndex
 * @param filterIndex
 * @param optional nil or *CellsPivotTablesGetWorksheetPivotTableFilterOpts - Optional Parameters:
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return PivotFilterResponse
*/


type CellsPivotTablesGetWorksheetPivotTableFilterOpts struct { 
	Name string
	SheetName string
	PivotTableIndex int32
	FilterIndex int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPivotTablesGetWorksheetPivotTableFilter(    localVarOptionals *CellsPivotTablesGetWorksheetPivotTableFilterOpts) (PivotFilterResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue PivotFilterResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/PivotFilters/{filterIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pivotTableIndex"+"}", fmt.Sprintf("%v", localVarOptionals.PivotTableIndex), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"filterIndex"+"}", fmt.Sprintf("%v", localVarOptionals.FilterIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService
 * @param name
 * @param sheetName
 * @param pivotTableIndex
 * @param optional nil or *CellsPivotTablesGetWorksheetPivotTableFiltersOpts - Optional Parameters:
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return PivotFiltersResponse
*/


type CellsPivotTablesGetWorksheetPivotTableFiltersOpts struct { 
	Name string
	SheetName string
	PivotTableIndex int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPivotTablesGetWorksheetPivotTableFilters(    localVarOptionals *CellsPivotTablesGetWorksheetPivotTableFiltersOpts) (PivotFiltersResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue PivotFiltersResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/PivotFilters"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pivotTableIndex"+"}", fmt.Sprintf("%v", localVarOptionals.PivotTableIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get worksheet pivottables info.
 * @param name Document name.
 * @param sheetName The worksheet name.
 * @param optional nil or *CellsPivotTablesGetWorksheetPivotTablesOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return PivotTablesResponse
*/


type CellsPivotTablesGetWorksheetPivotTablesOpts struct { 
	Name string
	SheetName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPivotTablesGetWorksheetPivotTables(    localVarOptionals *CellsPivotTablesGetWorksheetPivotTablesOpts) (PivotTablesResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue PivotTablesResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/pivottables"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Update cell style for pivot table
 * @param name Document name.
 * @param sheetName The worksheet name.
 * @param pivotTableIndex Pivot table index
 * @param column
 * @param row
 * @param optional nil or *CellsPivotTablesPostPivotTableCellStyleOpts - Optional Parameters:
     * @param "Style" (optional.Interface of Style) -  Style dto in request body.
     * @param "NeedReCalculate" (optional.Bool) - 
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPivotTablesPostPivotTableCellStyleOpts struct { 
	Name string
	SheetName string
	PivotTableIndex int32
	Column int32
	Row int32
	Style *Style
	NeedReCalculate bool
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPivotTablesPostPivotTableCellStyle(    localVarOptionals *CellsPivotTablesPostPivotTableCellStyleOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/Format"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pivotTableIndex"+"}", fmt.Sprintf("%v", localVarOptionals.PivotTableIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("column", parameterToString(localVarOptionals.Column, ""))
	localVarQueryParams.Add("row", parameterToString(localVarOptionals.Row, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("needReCalculate", parameterToString(localVarOptionals.NeedReCalculate.Value(), ""))
		localVarQueryParams.Add("needReCalculate", parameterToString(localVarOptionals.NeedReCalculate, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Style.IsSet() {
//		
//		localVarOptionalStyle, localVarOptionalStyleok := localVarOptionals.Style.Value().(Style)
//		if !localVarOptionalStyleok {
//				return localVarReturnValue, nil, reportError("style should be Style")
//		}
//		localVarPostBody = &localVarOptionalStyle
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Style != nil {
		
		localVarPostBody = &localVarOptionals.Style
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService
 * @param name
 * @param sheetName
 * @param pivotTableIndex
 * @param pivotFieldType
 * @param fieldIndex
 * @param itemIndex
 * @param isHide
 * @param optional nil or *CellsPivotTablesPostPivotTableFieldHideItemOpts - Optional Parameters:
     * @param "NeedReCalculate" (optional.Bool) - 
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPivotTablesPostPivotTableFieldHideItemOpts struct { 
	Name string
	SheetName string
	PivotTableIndex int32
	PivotFieldType string
	FieldIndex int32
	ItemIndex int32
	IsHide bool
	NeedReCalculate bool
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPivotTablesPostPivotTableFieldHideItem(    localVarOptionals *CellsPivotTablesPostPivotTableFieldHideItemOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/PivotField/Hide"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pivotTableIndex"+"}", fmt.Sprintf("%v", localVarOptionals.PivotTableIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("pivotFieldType", parameterToString(localVarOptionals.PivotFieldType, ""))
	localVarQueryParams.Add("fieldIndex", parameterToString(localVarOptionals.FieldIndex, ""))
	localVarQueryParams.Add("itemIndex", parameterToString(localVarOptionals.ItemIndex, ""))
	localVarQueryParams.Add("isHide", parameterToString(localVarOptionals.IsHide, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("needReCalculate", parameterToString(localVarOptionals.NeedReCalculate.Value(), ""))
		localVarQueryParams.Add("needReCalculate", parameterToString(localVarOptionals.NeedReCalculate, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService
 * @param name
 * @param sheetName
 * @param pivotTableIndex
 * @param fieldIndex
 * @param from
 * @param to
 * @param optional nil or *CellsPivotTablesPostPivotTableFieldMoveToOpts - Optional Parameters:
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPivotTablesPostPivotTableFieldMoveToOpts struct { 
	Name string
	SheetName string
	PivotTableIndex int32
	FieldIndex int32
	From string
	To string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPivotTablesPostPivotTableFieldMoveTo(    localVarOptionals *CellsPivotTablesPostPivotTableFieldMoveToOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/PivotField/Move"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pivotTableIndex"+"}", fmt.Sprintf("%v", localVarOptionals.PivotTableIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("fieldIndex", parameterToString(localVarOptionals.FieldIndex, ""))
	localVarQueryParams.Add("from", parameterToString(localVarOptionals.From, ""))
	localVarQueryParams.Add("to", parameterToString(localVarOptionals.To, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Update style for pivot table
 * @param name Document name.
 * @param sheetName The worksheet name.
 * @param pivotTableIndex Pivot table index
 * @param optional nil or *CellsPivotTablesPostPivotTableStyleOpts - Optional Parameters:
     * @param "Style" (optional.Interface of Style) -  Style dto in request body.
     * @param "NeedReCalculate" (optional.Bool) - 
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPivotTablesPostPivotTableStyleOpts struct { 
	Name string
	SheetName string
	PivotTableIndex int32
	Style *Style
	NeedReCalculate bool
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPivotTablesPostPivotTableStyle(    localVarOptionals *CellsPivotTablesPostPivotTableStyleOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/FormatAll"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pivotTableIndex"+"}", fmt.Sprintf("%v", localVarOptionals.PivotTableIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("needReCalculate", parameterToString(localVarOptionals.NeedReCalculate.Value(), ""))
		localVarQueryParams.Add("needReCalculate", parameterToString(localVarOptionals.NeedReCalculate, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Style.IsSet() {
//		
//		localVarOptionalStyle, localVarOptionalStyleok := localVarOptionals.Style.Value().(Style)
//		if !localVarOptionalStyleok {
//				return localVarReturnValue, nil, reportError("style should be Style")
//		}
//		localVarPostBody = &localVarOptionalStyle
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Style != nil {
		
		localVarPostBody = &localVarOptionals.Style
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Calculates pivottable&#39;s data to cells.
 * @param name Document name.
 * @param sheetName The worksheet name.
 * @param pivotTableIndex Pivot table index
 * @param optional nil or *CellsPivotTablesPostWorksheetPivotTableCalculateOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPivotTablesPostWorksheetPivotTableCalculateOpts struct { 
	Name string
	SheetName string
	PivotTableIndex int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPivotTablesPostWorksheetPivotTableCalculate(    localVarOptionals *CellsPivotTablesPostWorksheetPivotTableCalculateOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/Calculate"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pivotTableIndex"+"}", fmt.Sprintf("%v", localVarOptionals.PivotTableIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService
 * @param name
 * @param sheetName
 * @param pivotTableIndex
 * @param optional nil or *CellsPivotTablesPostWorksheetPivotTableMoveOpts - Optional Parameters:
     * @param "Row" (optional.Int32) - 
     * @param "Column" (optional.Int32) - 
     * @param "DestCellName" (optional.String) - 
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPivotTablesPostWorksheetPivotTableMoveOpts struct { 
	Name string
	SheetName string
	PivotTableIndex int32
	Row int32
	Column int32
	DestCellName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPivotTablesPostWorksheetPivotTableMove(    localVarOptionals *CellsPivotTablesPostWorksheetPivotTableMoveOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/Move"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pivotTableIndex"+"}", fmt.Sprintf("%v", localVarOptionals.PivotTableIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("row", parameterToString(localVarOptionals.Row.Value(), ""))
		localVarQueryParams.Add("row", parameterToString(localVarOptionals.Row, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("column", parameterToString(localVarOptionals.Column.Value(), ""))
		localVarQueryParams.Add("column", parameterToString(localVarOptionals.Column, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("destCellName", parameterToString(localVarOptionals.DestCellName.Value(), ""))
		localVarQueryParams.Add("destCellName", parameterToString(localVarOptionals.DestCellName, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Add pivot field into into pivot table
 * @param name Document name.
 * @param sheetName The worksheet name.
 * @param pivotTableIndex Pivot table index
 * @param pivotFieldType The fields area type.
 * @param optional nil or *CellsPivotTablesPutPivotTableFieldOpts - Optional Parameters:
     * @param "Request" (optional.Interface of PivotTableFieldRequest) -  Dto that conrains field indexes
     * @param "NeedReCalculate" (optional.Bool) - 
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPivotTablesPutPivotTableFieldOpts struct { 
	Name string
	SheetName string
	PivotTableIndex int32
	PivotFieldType string
	Request *PivotTableFieldRequest
	NeedReCalculate bool
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPivotTablesPutPivotTableField(    localVarOptionals *CellsPivotTablesPutPivotTableFieldOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/PivotField"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pivotTableIndex"+"}", fmt.Sprintf("%v", localVarOptionals.PivotTableIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("pivotFieldType", parameterToString(localVarOptionals.PivotFieldType, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("needReCalculate", parameterToString(localVarOptionals.NeedReCalculate.Value(), ""))
		localVarQueryParams.Add("needReCalculate", parameterToString(localVarOptionals.NeedReCalculate, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Request.IsSet() {
//		
//		localVarOptionalRequest, localVarOptionalRequestok := localVarOptionals.Request.Value().(PivotTableFieldRequest)
//		if !localVarOptionalRequestok {
//				return localVarReturnValue, nil, reportError("request should be PivotTableFieldRequest")
//		}
//		localVarPostBody = &localVarOptionalRequest
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Request != nil {
		
		localVarPostBody = &localVarOptionals.Request
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Add a pivot table into worksheet.
 * @param name Document name.
 * @param sheetName The worksheet name.
 * @param optional nil or *CellsPivotTablesPutWorksheetPivotTableOpts - Optional Parameters:
     * @param "Request" (optional.Interface of CreatePivotTableRequest) -  CreatePivotTableRequest dto in request body.
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.
     * @param "SourceData" (optional.String) -  The data for the new PivotTable cache.
     * @param "DestCellName" (optional.String) -  The cell in the upper-left corner of the PivotTable report&#39;s destination range.
     * @param "TableName" (optional.String) -  The name of the new PivotTable report.
     * @param "UseSameSource" (optional.Bool) -  Indicates whether using same data source when another existing pivot table has used this data source. If the property is true, it will save memory.

@return PivotTableResponse
*/


type CellsPivotTablesPutWorksheetPivotTableOpts struct { 
	Name string
	SheetName string
	Request *CreatePivotTableRequest
	Folder string
	Storage string
	SourceData string
	DestCellName string
	TableName string
	UseSameSource bool
}


func (a *CellsApiService) CellsPivotTablesPutWorksheetPivotTable(    localVarOptionals *CellsPivotTablesPutWorksheetPivotTableOpts) (PivotTableResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue PivotTableResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/pivottables"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("sourceData", parameterToString(localVarOptionals.SourceData.Value(), ""))
		localVarQueryParams.Add("sourceData", parameterToString(localVarOptionals.SourceData, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("destCellName", parameterToString(localVarOptionals.DestCellName.Value(), ""))
		localVarQueryParams.Add("destCellName", parameterToString(localVarOptionals.DestCellName, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("tableName", parameterToString(localVarOptionals.TableName.Value(), ""))
		localVarQueryParams.Add("tableName", parameterToString(localVarOptionals.TableName, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("useSameSource", parameterToString(localVarOptionals.UseSameSource.Value(), ""))
		localVarQueryParams.Add("useSameSource", parameterToString(localVarOptionals.UseSameSource, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Request.IsSet() {
//		
//		localVarOptionalRequest, localVarOptionalRequestok := localVarOptionals.Request.Value().(CreatePivotTableRequest)
//		if !localVarOptionalRequestok {
//				return localVarReturnValue, nil, reportError("request should be CreatePivotTableRequest")
//		}
//		localVarPostBody = &localVarOptionalRequest
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Request != nil {
		
		localVarPostBody = &localVarOptionals.Request
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Add pivot filter for piovt table index
 * @param name
 * @param sheetName
 * @param pivotTableIndex
 * @param optional nil or *CellsPivotTablesPutWorksheetPivotTableFilterOpts - Optional Parameters:
     * @param "Filter" (optional.Interface of PivotFilter) - 
     * @param "NeedReCalculate" (optional.Bool) - 
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPivotTablesPutWorksheetPivotTableFilterOpts struct { 
	Name string
	SheetName string
	PivotTableIndex int32
	Filter *PivotFilter
	NeedReCalculate bool
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPivotTablesPutWorksheetPivotTableFilter(    localVarOptionals *CellsPivotTablesPutWorksheetPivotTableFilterOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/PivotFilters"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pivotTableIndex"+"}", fmt.Sprintf("%v", localVarOptionals.PivotTableIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("needReCalculate", parameterToString(localVarOptionals.NeedReCalculate.Value(), ""))
		localVarQueryParams.Add("needReCalculate", parameterToString(localVarOptionals.NeedReCalculate, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Filter.IsSet() {
//		
//		localVarOptionalFilter, localVarOptionalFilterok := localVarOptionals.Filter.Value().(PivotFilter)
//		if !localVarOptionalFilterok {
//				return localVarReturnValue, nil, reportError("filter should be PivotFilter")
//		}
//		localVarPostBody = &localVarOptionalFilter
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Filter != nil {
		
		localVarPostBody = &localVarOptionals.Filter
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Cell calculate formula
 * @param name
 * @param sheetName
 * @param cellName
 * @param optional nil or *CellsPostCellCalculateOpts - Optional Parameters:
     * @param "Options" (optional.Interface of CalculationOptions) - 
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPostCellCalculateOpts struct { 
	Name string
	SheetName string
	CellName string
	Options *CalculationOptions
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPostCellCalculate(    localVarOptionals *CellsPostCellCalculateOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/{cellName}/calculate"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cellName"+"}", fmt.Sprintf("%v", localVarOptionals.CellName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Options.IsSet() {
//		
//		localVarOptionalOptions, localVarOptionalOptionsok := localVarOptionals.Options.Value().(CalculationOptions)
//		if !localVarOptionalOptionsok {
//				return localVarReturnValue, nil, reportError("options should be CalculationOptions")
//		}
//		localVarPostBody = &localVarOptionalOptions
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Options != nil {
		
		localVarPostBody = &localVarOptionals.Options
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Set cell characters 
 * @param name
 * @param sheetName
 * @param cellName
 * @param optional nil or *CellsPostCellCharactersOpts - Optional Parameters:
     * @param "Options" (optional.Interface of []FontSetting) - 
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPostCellCharactersOpts struct { 
	Name string
	SheetName string
	CellName string
	Options *[]FontSetting
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPostCellCharacters(    localVarOptionals *CellsPostCellCharactersOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/{cellName}/characters"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cellName"+"}", fmt.Sprintf("%v", localVarOptionals.CellName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Options.IsSet() {
//		
//		localVarOptionalOptions, localVarOptionalOptionsok := localVarOptionals.Options.Value().([]FontSetting)
//		if !localVarOptionalOptionsok {
//				return localVarReturnValue, nil, reportError("options should be []FontSetting")
//		}
//		localVarPostBody = &localVarOptionalOptions
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Options != nil {
		
		localVarPostBody = &localVarOptionals.Options
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Clear cells contents.
 * @param name Workbook name.
 * @param sheetName Worksheet name.
 * @param optional nil or *CellsPostClearContentsOpts - Optional Parameters:
     * @param "Range_" (optional.String) -  The range.
     * @param "StartRow" (optional.Int32) -  The start row.
     * @param "StartColumn" (optional.Int32) -  The start column.
     * @param "EndRow" (optional.Int32) -  The end row.
     * @param "EndColumn" (optional.Int32) -  The end column.
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPostClearContentsOpts struct { 
	Name string
	SheetName string
	Range_ string
	StartRow int32
	StartColumn int32
	EndRow int32
	EndColumn int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPostClearContents(    localVarOptionals *CellsPostClearContentsOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/clearcontents"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("range", parameterToString(localVarOptionals.Range_.Value(), ""))
		localVarQueryParams.Add("range", parameterToString(localVarOptionals.Range_, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("startRow", parameterToString(localVarOptionals.StartRow.Value(), ""))
		localVarQueryParams.Add("startRow", parameterToString(localVarOptionals.StartRow, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("startColumn", parameterToString(localVarOptionals.StartColumn.Value(), ""))
		localVarQueryParams.Add("startColumn", parameterToString(localVarOptionals.StartColumn, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("endRow", parameterToString(localVarOptionals.EndRow.Value(), ""))
		localVarQueryParams.Add("endRow", parameterToString(localVarOptionals.EndRow, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("endColumn", parameterToString(localVarOptionals.EndColumn.Value(), ""))
		localVarQueryParams.Add("endColumn", parameterToString(localVarOptionals.EndColumn, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Clear cells contents.
 * @param name Workbook name.
 * @param sheetName Worksheet name.
 * @param optional nil or *CellsPostClearFormatsOpts - Optional Parameters:
     * @param "Range_" (optional.String) -  The range.
     * @param "StartRow" (optional.Int32) -  The start row.
     * @param "StartColumn" (optional.Int32) -  The start column.
     * @param "EndRow" (optional.Int32) -  The end row.
     * @param "EndColumn" (optional.Int32) -  The end column.
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPostClearFormatsOpts struct { 
	Name string
	SheetName string
	Range_ string
	StartRow int32
	StartColumn int32
	EndRow int32
	EndColumn int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPostClearFormats(    localVarOptionals *CellsPostClearFormatsOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/clearformats"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("range", parameterToString(localVarOptionals.Range_.Value(), ""))
		localVarQueryParams.Add("range", parameterToString(localVarOptionals.Range_, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("startRow", parameterToString(localVarOptionals.StartRow.Value(), ""))
		localVarQueryParams.Add("startRow", parameterToString(localVarOptionals.StartRow, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("startColumn", parameterToString(localVarOptionals.StartColumn.Value(), ""))
		localVarQueryParams.Add("startColumn", parameterToString(localVarOptionals.StartColumn, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("endRow", parameterToString(localVarOptionals.EndRow.Value(), ""))
		localVarQueryParams.Add("endRow", parameterToString(localVarOptionals.EndRow, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("endColumn", parameterToString(localVarOptionals.EndColumn.Value(), ""))
		localVarQueryParams.Add("endColumn", parameterToString(localVarOptionals.EndColumn, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Set column style
 * @param name The workbook name.
 * @param sheetName The worksheet name.
 * @param columnIndex The column index.
 * @param optional nil or *CellsPostColumnStyleOpts - Optional Parameters:
     * @param "Style" (optional.Interface of Style) -  Style dto
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPostColumnStyleOpts struct { 
	Name string
	SheetName string
	ColumnIndex int32
	Style *Style
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPostColumnStyle(    localVarOptionals *CellsPostColumnStyleOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/columns/{columnIndex}/style"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"columnIndex"+"}", fmt.Sprintf("%v", localVarOptionals.ColumnIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Style.IsSet() {
//		
//		localVarOptionalStyle, localVarOptionalStyleok := localVarOptionals.Style.Value().(Style)
//		if !localVarOptionalStyleok {
//				return localVarReturnValue, nil, reportError("style should be Style")
//		}
//		localVarPostBody = &localVarOptionalStyle
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Style != nil {
		
		localVarPostBody = &localVarOptionals.Style
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Copy cell into cell
 * @param name Workbook name.
 * @param destCellName Destination cell name
 * @param sheetName Destination worksheet name.
 * @param worksheet Source worksheet name.
 * @param optional nil or *CellsPostCopyCellIntoCellOpts - Optional Parameters:
     * @param "Cellname" (optional.String) -  Source cell name
     * @param "Row" (optional.Int32) -  Source row
     * @param "Column" (optional.Int32) -  Source column
     * @param "Folder" (optional.String) -  Folder name
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPostCopyCellIntoCellOpts struct { 
	Name string
	DestCellName string
	SheetName string
	Worksheet string
	Cellname string
	Row int32
	Column int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPostCopyCellIntoCell(    localVarOptionals *CellsPostCopyCellIntoCellOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/{destCellName}/copy"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"destCellName"+"}", fmt.Sprintf("%v", localVarOptionals.DestCellName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("worksheet", parameterToString(localVarOptionals.Worksheet, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("cellname", parameterToString(localVarOptionals.Cellname.Value(), ""))
		localVarQueryParams.Add("cellname", parameterToString(localVarOptionals.Cellname, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("row", parameterToString(localVarOptionals.Row.Value(), ""))
		localVarQueryParams.Add("row", parameterToString(localVarOptionals.Row, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("column", parameterToString(localVarOptionals.Column.Value(), ""))
		localVarQueryParams.Add("column", parameterToString(localVarOptionals.Column, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Copy worksheet columns.
 * @param name The workbook name.
 * @param sheetName The worksheet name.
 * @param sourceColumnIndex Source column index
 * @param destinationColumnIndex Destination column index
 * @param columnNumber The copied column number
 * @param optional nil or *CellsPostCopyWorksheetColumnsOpts - Optional Parameters:
     * @param "Worksheet" (optional.String) -  The Worksheet
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPostCopyWorksheetColumnsOpts struct { 
	Name string
	SheetName string
	SourceColumnIndex int32
	DestinationColumnIndex int32
	ColumnNumber int32
	Worksheet string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPostCopyWorksheetColumns(    localVarOptionals *CellsPostCopyWorksheetColumnsOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/columns/copy"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("sourceColumnIndex", parameterToString(localVarOptionals.SourceColumnIndex, ""))
	localVarQueryParams.Add("destinationColumnIndex", parameterToString(localVarOptionals.DestinationColumnIndex, ""))
	localVarQueryParams.Add("columnNumber", parameterToString(localVarOptionals.ColumnNumber, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("worksheet", parameterToString(localVarOptionals.Worksheet.Value(), ""))
		localVarQueryParams.Add("worksheet", parameterToString(localVarOptionals.Worksheet, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Copy worksheet rows.
 * @param name The workbook name.
 * @param sheetName The worksheet name.
 * @param sourceRowIndex Source row index
 * @param destinationRowIndex Destination row index
 * @param rowNumber The copied row number
 * @param optional nil or *CellsPostCopyWorksheetRowsOpts - Optional Parameters:
     * @param "Worksheet" (optional.String) -  worksheet
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPostCopyWorksheetRowsOpts struct { 
	Name string
	SheetName string
	SourceRowIndex int32
	DestinationRowIndex int32
	RowNumber int32
	Worksheet string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPostCopyWorksheetRows(    localVarOptionals *CellsPostCopyWorksheetRowsOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/rows/copy"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("sourceRowIndex", parameterToString(localVarOptionals.SourceRowIndex, ""))
	localVarQueryParams.Add("destinationRowIndex", parameterToString(localVarOptionals.DestinationRowIndex, ""))
	localVarQueryParams.Add("rowNumber", parameterToString(localVarOptionals.RowNumber, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("worksheet", parameterToString(localVarOptionals.Worksheet.Value(), ""))
		localVarQueryParams.Add("worksheet", parameterToString(localVarOptionals.Worksheet, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Group worksheet columns.
 * @param name The workbook name.
 * @param sheetName The worksheet name.
 * @param firstIndex The first column index to be operated.
 * @param lastIndex The last column index to be operated.
 * @param optional nil or *CellsPostGroupWorksheetColumnsOpts - Optional Parameters:
     * @param "Hide" (optional.Bool) -  columns visible state
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPostGroupWorksheetColumnsOpts struct { 
	Name string
	SheetName string
	FirstIndex int32
	LastIndex int32
	Hide bool
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPostGroupWorksheetColumns(    localVarOptionals *CellsPostGroupWorksheetColumnsOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/columns/group"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("firstIndex", parameterToString(localVarOptionals.FirstIndex, ""))
	localVarQueryParams.Add("lastIndex", parameterToString(localVarOptionals.LastIndex, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("hide", parameterToString(localVarOptionals.Hide.Value(), ""))
		localVarQueryParams.Add("hide", parameterToString(localVarOptionals.Hide, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Group worksheet rows.
 * @param name The workbook name.
 * @param sheetName The worksheet name.
 * @param firstIndex The first row index to be operated.
 * @param lastIndex The last row index to be operated.
 * @param optional nil or *CellsPostGroupWorksheetRowsOpts - Optional Parameters:
     * @param "Hide" (optional.Bool) -  rows visible state
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPostGroupWorksheetRowsOpts struct { 
	Name string
	SheetName string
	FirstIndex int32
	LastIndex int32
	Hide bool
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPostGroupWorksheetRows(    localVarOptionals *CellsPostGroupWorksheetRowsOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/rows/group"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("firstIndex", parameterToString(localVarOptionals.FirstIndex, ""))
	localVarQueryParams.Add("lastIndex", parameterToString(localVarOptionals.LastIndex, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("hide", parameterToString(localVarOptionals.Hide.Value(), ""))
		localVarQueryParams.Add("hide", parameterToString(localVarOptionals.Hide, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Hide worksheet columns.
 * @param name The workbook name.
 * @param sheetName The worksheet name.
 * @param startColumn The begin column index to be operated.
 * @param totalColumns Number of columns to be operated.
 * @param optional nil or *CellsPostHideWorksheetColumnsOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPostHideWorksheetColumnsOpts struct { 
	Name string
	SheetName string
	StartColumn int32
	TotalColumns int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPostHideWorksheetColumns(    localVarOptionals *CellsPostHideWorksheetColumnsOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/columns/hide"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("startColumn", parameterToString(localVarOptionals.StartColumn, ""))
	localVarQueryParams.Add("totalColumns", parameterToString(localVarOptionals.TotalColumns, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Hide worksheet rows.
 * @param name The workbook name.
 * @param sheetName The worksheet name.
 * @param startrow The begin row index to be operated.
 * @param totalRows Number of rows to be operated.
 * @param optional nil or *CellsPostHideWorksheetRowsOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPostHideWorksheetRowsOpts struct { 
	Name string
	SheetName string
	Startrow int32
	TotalRows int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPostHideWorksheetRows(    localVarOptionals *CellsPostHideWorksheetRowsOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/rows/hide"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("startrow", parameterToString(localVarOptionals.Startrow, ""))
	localVarQueryParams.Add("totalRows", parameterToString(localVarOptionals.TotalRows, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Set row style.
 * @param name The workbook name.
 * @param sheetName The worksheet name.
 * @param rowIndex The row index.
 * @param optional nil or *CellsPostRowStyleOpts - Optional Parameters:
     * @param "Style" (optional.Interface of Style) -  Style dto
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPostRowStyleOpts struct { 
	Name string
	SheetName string
	RowIndex int32
	Style *Style
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPostRowStyle(    localVarOptionals *CellsPostRowStyleOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/rows/{rowIndex}/style"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rowIndex"+"}", fmt.Sprintf("%v", localVarOptionals.RowIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Style.IsSet() {
//		
//		localVarOptionalStyle, localVarOptionalStyleok := localVarOptionals.Style.Value().(Style)
//		if !localVarOptionalStyleok {
//				return localVarReturnValue, nil, reportError("style should be Style")
//		}
//		localVarPostBody = &localVarOptionalStyle
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Style != nil {
		
		localVarPostBody = &localVarOptionals.Style
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Set htmlstring value into cell
 * @param name Workbook name.
 * @param sheetName Worksheet name.
 * @param cellName The cell name.
 * @param htmlString
 * @param optional nil or *CellsPostSetCellHtmlStringOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellResponse
*/


type CellsPostSetCellHtmlStringOpts struct { 
	Name string
	SheetName string
	CellName string
	HtmlString string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPostSetCellHtmlString(    localVarOptionals *CellsPostSetCellHtmlStringOpts) (CellResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/{cellName}/htmlstring"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cellName"+"}", fmt.Sprintf("%v", localVarOptionals.CellName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
	localVarPostBody = &localVarOptionals.HtmlString
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Set cell range value 
 * @param name Workbook name.
 * @param sheetName Worksheet name.
 * @param cellarea Cell area (like \&quot;A1:C2\&quot;)
 * @param value Range value
 * @param type_ Value data type (like \&quot;int\&quot;)
 * @param optional nil or *CellsPostSetCellRangeValueOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Folder name
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPostSetCellRangeValueOpts struct { 
	Name string
	SheetName string
	Cellarea string
	Value string
	Type_ string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPostSetCellRangeValue(    localVarOptionals *CellsPostSetCellRangeValueOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("cellarea", parameterToString(localVarOptionals.Cellarea, ""))
	localVarQueryParams.Add("value", parameterToString(localVarOptionals.Value, ""))
	localVarQueryParams.Add("type", parameterToString(localVarOptionals.Type_, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Set worksheet column width.
 * @param name The workbook name.
 * @param sheetName The worksheet name.
 * @param columnIndex The column index.
 * @param width The width.
 * @param optional nil or *CellsPostSetWorksheetColumnWidthOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return ColumnResponse
*/


type CellsPostSetWorksheetColumnWidthOpts struct { 
	Name string
	SheetName string
	ColumnIndex int32
	Width float64
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPostSetWorksheetColumnWidth(    localVarOptionals *CellsPostSetWorksheetColumnWidthOpts) (ColumnResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ColumnResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/columns/{columnIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"columnIndex"+"}", fmt.Sprintf("%v", localVarOptionals.ColumnIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("width", parameterToString(localVarOptionals.Width, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Ungroup worksheet columns.
 * @param name The workbook name.
 * @param sheetName The worksheet name.
 * @param firstIndex The first column index to be operated.
 * @param lastIndex The last column index to be operated.
 * @param optional nil or *CellsPostUngroupWorksheetColumnsOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPostUngroupWorksheetColumnsOpts struct { 
	Name string
	SheetName string
	FirstIndex int32
	LastIndex int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPostUngroupWorksheetColumns(    localVarOptionals *CellsPostUngroupWorksheetColumnsOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/columns/ungroup"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("firstIndex", parameterToString(localVarOptionals.FirstIndex, ""))
	localVarQueryParams.Add("lastIndex", parameterToString(localVarOptionals.LastIndex, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Ungroup worksheet rows.
 * @param name The workbook name.
 * @param sheetName The worksheet name.
 * @param firstIndex The first row index to be operated.
 * @param lastIndex The last row index to be operated.
 * @param optional nil or *CellsPostUngroupWorksheetRowsOpts - Optional Parameters:
     * @param "IsAll" (optional.Bool) -  Is all row to be operated
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPostUngroupWorksheetRowsOpts struct { 
	Name string
	SheetName string
	FirstIndex int32
	LastIndex int32
	IsAll bool
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPostUngroupWorksheetRows(    localVarOptionals *CellsPostUngroupWorksheetRowsOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/rows/ungroup"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("firstIndex", parameterToString(localVarOptionals.FirstIndex, ""))
	localVarQueryParams.Add("lastIndex", parameterToString(localVarOptionals.LastIndex, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("isAll", parameterToString(localVarOptionals.IsAll.Value(), ""))
		localVarQueryParams.Add("isAll", parameterToString(localVarOptionals.IsAll, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Unhide worksheet columns.
 * @param name The workbook name.
 * @param sheetName The worksheet name.
 * @param startcolumn The begin column index to be operated.
 * @param totalColumns Number of columns to be operated.
 * @param optional nil or *CellsPostUnhideWorksheetColumnsOpts - Optional Parameters:
     * @param "Width" (optional.Float64) -  The new column width.
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPostUnhideWorksheetColumnsOpts struct { 
	Name string
	SheetName string
	Startcolumn int32
	TotalColumns int32
	Width float64
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPostUnhideWorksheetColumns(    localVarOptionals *CellsPostUnhideWorksheetColumnsOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/columns/unhide"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("startcolumn", parameterToString(localVarOptionals.Startcolumn, ""))
	localVarQueryParams.Add("totalColumns", parameterToString(localVarOptionals.TotalColumns, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("width", parameterToString(localVarOptionals.Width.Value(), ""))
		localVarQueryParams.Add("width", parameterToString(localVarOptionals.Width, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Unhide worksheet rows.
 * @param name The workbook name.
 * @param sheetName The worksheet name.
 * @param startrow The begin row index to be operated.
 * @param totalRows Number of rows to be operated.
 * @param optional nil or *CellsPostUnhideWorksheetRowsOpts - Optional Parameters:
     * @param "Height" (optional.Float64) -  The new row height.
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPostUnhideWorksheetRowsOpts struct { 
	Name string
	SheetName string
	Startrow int32
	TotalRows int32
	Height float64
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPostUnhideWorksheetRows(    localVarOptionals *CellsPostUnhideWorksheetRowsOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/rows/unhide"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("startrow", parameterToString(localVarOptionals.Startrow, ""))
	localVarQueryParams.Add("totalRows", parameterToString(localVarOptionals.TotalRows, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("height", parameterToString(localVarOptionals.Height.Value(), ""))
		localVarQueryParams.Add("height", parameterToString(localVarOptionals.Height, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Update cell&#39;s style.
 * @param name Workbook name.
 * @param sheetName Worksheet name.
 * @param cellName The cell name.
 * @param optional nil or *CellsPostUpdateWorksheetCellStyleOpts - Optional Parameters:
     * @param "Style" (optional.Interface of Style) -  with update style settings.
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return StyleResponse
*/


type CellsPostUpdateWorksheetCellStyleOpts struct { 
	Name string
	SheetName string
	CellName string
	Style *Style
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPostUpdateWorksheetCellStyle(    localVarOptionals *CellsPostUpdateWorksheetCellStyleOpts) (StyleResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue StyleResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/{cellName}/style"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cellName"+"}", fmt.Sprintf("%v", localVarOptionals.CellName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Style.IsSet() {
//		
//		localVarOptionalStyle, localVarOptionalStyleok := localVarOptionals.Style.Value().(Style)
//		if !localVarOptionalStyleok {
//				return localVarReturnValue, nil, reportError("style should be Style")
//		}
//		localVarPostBody = &localVarOptionalStyle
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Style != nil {
		
		localVarPostBody = &localVarOptionals.Style
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Update cell&#39;s range style.
 * @param name Workbook name.
 * @param sheetName Worksheet name.
 * @param range_ The range.
 * @param optional nil or *CellsPostUpdateWorksheetRangeStyleOpts - Optional Parameters:
     * @param "Style" (optional.Interface of Style) -  with update style settings.
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPostUpdateWorksheetRangeStyleOpts struct { 
	Name string
	SheetName string
	Range_ string
	Style *Style
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPostUpdateWorksheetRangeStyle(    localVarOptionals *CellsPostUpdateWorksheetRangeStyleOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/style"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("range", parameterToString(localVarOptionals.Range_, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Style.IsSet() {
//		
//		localVarOptionalStyle, localVarOptionalStyleok := localVarOptionals.Style.Value().(Style)
//		if !localVarOptionalStyleok {
//				return localVarReturnValue, nil, reportError("style should be Style")
//		}
//		localVarPostBody = &localVarOptionalStyle
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Style != nil {
		
		localVarPostBody = &localVarOptionals.Style
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Update worksheet row.
 * @param name The workbook name.
 * @param sheetName The worksheet name.
 * @param rowIndex The row index.
 * @param optional nil or *CellsPostUpdateWorksheetRowOpts - Optional Parameters:
     * @param "Height" (optional.Float64) -  The new row height.
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return RowResponse
*/


type CellsPostUpdateWorksheetRowOpts struct { 
	Name string
	SheetName string
	RowIndex int32
	Height float64
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPostUpdateWorksheetRow(    localVarOptionals *CellsPostUpdateWorksheetRowOpts) (RowResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue RowResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/rows/{rowIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rowIndex"+"}", fmt.Sprintf("%v", localVarOptionals.RowIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("height", parameterToString(localVarOptionals.Height.Value(), ""))
		localVarQueryParams.Add("height", parameterToString(localVarOptionals.Height, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Set cell value.
 * @param name The document name.
 * @param sheetName The worksheet name.
 * @param cellName The cell name.
 * @param optional nil or *CellsPostWorksheetCellSetValueOpts - Optional Parameters:
     * @param "Value" (optional.String) -  The cell value.
     * @param "Type_" (optional.String) -  The value type.
     * @param "Formula" (optional.String) -  Formula for cell
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellResponse
*/


type CellsPostWorksheetCellSetValueOpts struct { 
	Name string
	SheetName string
	CellName string
	Value string
	Type_ string
	Formula string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPostWorksheetCellSetValue(    localVarOptionals *CellsPostWorksheetCellSetValueOpts) (CellResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/{cellName}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cellName"+"}", fmt.Sprintf("%v", localVarOptionals.CellName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("value", parameterToString(localVarOptionals.Value.Value(), ""))
		localVarQueryParams.Add("value", parameterToString(localVarOptionals.Value, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("type", parameterToString(localVarOptionals.Type_.Value(), ""))
		localVarQueryParams.Add("type", parameterToString(localVarOptionals.Type_, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("formula", parameterToString(localVarOptionals.Formula.Value(), ""))
		localVarQueryParams.Add("formula", parameterToString(localVarOptionals.Formula, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Merge cells.
 * @param name The workbook name.
 * @param sheetName The worksheet name.
 * @param startRow The start row.
 * @param startColumn The start column.
 * @param totalRows The total rows
 * @param totalColumns The total columns.
 * @param optional nil or *CellsPostWorksheetMergeOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPostWorksheetMergeOpts struct { 
	Name string
	SheetName string
	StartRow int32
	StartColumn int32
	TotalRows int32
	TotalColumns int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPostWorksheetMerge(    localVarOptionals *CellsPostWorksheetMergeOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/merge"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("startRow", parameterToString(localVarOptionals.StartRow, ""))
	localVarQueryParams.Add("startColumn", parameterToString(localVarOptionals.StartColumn, ""))
	localVarQueryParams.Add("totalRows", parameterToString(localVarOptionals.TotalRows, ""))
	localVarQueryParams.Add("totalColumns", parameterToString(localVarOptionals.TotalColumns, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Unmerge cells.
 * @param name The workbook name.
 * @param sheetName The worksheet name.
 * @param startRow The start row.
 * @param startColumn The start column.
 * @param totalRows The total rows
 * @param totalColumns The total columns.
 * @param optional nil or *CellsPostWorksheetUnmergeOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPostWorksheetUnmergeOpts struct { 
	Name string
	SheetName string
	StartRow int32
	StartColumn int32
	TotalRows int32
	TotalColumns int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPostWorksheetUnmerge(    localVarOptionals *CellsPostWorksheetUnmergeOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/unmerge"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("startRow", parameterToString(localVarOptionals.StartRow, ""))
	localVarQueryParams.Add("startColumn", parameterToString(localVarOptionals.StartColumn, ""))
	localVarQueryParams.Add("totalRows", parameterToString(localVarOptionals.TotalRows, ""))
	localVarQueryParams.Add("totalColumns", parameterToString(localVarOptionals.TotalColumns, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Delete all custom document properties and clean built-in ones.
 * @param name The document name.
 * @param optional nil or *CellsPropertiesDeleteDocumentPropertiesOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsDocumentPropertiesResponse
*/


type CellsPropertiesDeleteDocumentPropertiesOpts struct { 
	Name string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPropertiesDeleteDocumentProperties(    localVarOptionals *CellsPropertiesDeleteDocumentPropertiesOpts) (CellsDocumentPropertiesResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsDocumentPropertiesResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/documentproperties"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Delete document property.
 * @param name The document name.
 * @param propertyName The property name.
 * @param optional nil or *CellsPropertiesDeleteDocumentPropertyOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsDocumentPropertiesResponse
*/


type CellsPropertiesDeleteDocumentPropertyOpts struct { 
	Name string
	PropertyName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPropertiesDeleteDocumentProperty(    localVarOptionals *CellsPropertiesDeleteDocumentPropertyOpts) (CellsDocumentPropertiesResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsDocumentPropertiesResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/documentproperties/{propertyName}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"propertyName"+"}", fmt.Sprintf("%v", localVarOptionals.PropertyName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Read document properties.
 * @param name The document name.
 * @param optional nil or *CellsPropertiesGetDocumentPropertiesOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsDocumentPropertiesResponse
*/


type CellsPropertiesGetDocumentPropertiesOpts struct { 
	Name string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPropertiesGetDocumentProperties(    localVarOptionals *CellsPropertiesGetDocumentPropertiesOpts) (CellsDocumentPropertiesResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsDocumentPropertiesResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/documentproperties"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Read document property by name.
 * @param name The document name.
 * @param propertyName The property name.
 * @param optional nil or *CellsPropertiesGetDocumentPropertyOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsDocumentPropertyResponse
*/


type CellsPropertiesGetDocumentPropertyOpts struct { 
	Name string
	PropertyName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPropertiesGetDocumentProperty(    localVarOptionals *CellsPropertiesGetDocumentPropertyOpts) (CellsDocumentPropertyResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsDocumentPropertyResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/documentproperties/{propertyName}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"propertyName"+"}", fmt.Sprintf("%v", localVarOptionals.PropertyName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Set/create document property.
 * @param name The document name.
 * @param propertyName The property name.
 * @param optional nil or *CellsPropertiesPutDocumentPropertyOpts - Optional Parameters:
     * @param "Property" (optional.Interface of CellsDocumentProperty) -  with new property value.
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsDocumentPropertyResponse
*/


type CellsPropertiesPutDocumentPropertyOpts struct { 
	Name string
	PropertyName string
	Property *CellsDocumentProperty
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPropertiesPutDocumentProperty(    localVarOptionals *CellsPropertiesPutDocumentPropertyOpts) (CellsDocumentPropertyResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsDocumentPropertyResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/documentproperties/{propertyName}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"propertyName"+"}", fmt.Sprintf("%v", localVarOptionals.PropertyName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Property.IsSet() {
//		
//		localVarOptionalProperty, localVarOptionalPropertyok := localVarOptionals.Property.Value().(CellsDocumentProperty)
//		if !localVarOptionalPropertyok {
//				return localVarReturnValue, nil, reportError("property should be CellsDocumentProperty")
//		}
//		localVarPostBody = &localVarOptionalProperty
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Property != nil {
		
		localVarPostBody = &localVarOptionals.Property
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Insert worksheet columns.
 * @param name The workbook name.
 * @param sheetName The worksheet name.
 * @param columnIndex The column index.
 * @param columns The columns.
 * @param optional nil or *CellsPutInsertWorksheetColumnsOpts - Optional Parameters:
     * @param "UpdateReference" (optional.Bool) -  The update reference.
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return ColumnsResponse
*/


type CellsPutInsertWorksheetColumnsOpts struct { 
	Name string
	SheetName string
	ColumnIndex int32
	Columns int32
	UpdateReference bool
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPutInsertWorksheetColumns(    localVarOptionals *CellsPutInsertWorksheetColumnsOpts) (ColumnsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ColumnsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/columns/{columnIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"columnIndex"+"}", fmt.Sprintf("%v", localVarOptionals.ColumnIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("columns", parameterToString(localVarOptionals.Columns, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("updateReference", parameterToString(localVarOptionals.UpdateReference.Value(), ""))
		localVarQueryParams.Add("updateReference", parameterToString(localVarOptionals.UpdateReference, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Insert new worksheet row.
 * @param name The workbook name.
 * @param sheetName The worksheet name.
 * @param rowIndex The new row index.
 * @param optional nil or *CellsPutInsertWorksheetRowOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return RowResponse
*/


type CellsPutInsertWorksheetRowOpts struct { 
	Name string
	SheetName string
	RowIndex int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPutInsertWorksheetRow(    localVarOptionals *CellsPutInsertWorksheetRowOpts) (RowResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue RowResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/rows/{rowIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rowIndex"+"}", fmt.Sprintf("%v", localVarOptionals.RowIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Insert several new worksheet rows.
 * @param name The workbook name.
 * @param sheetName The worksheet name.
 * @param startrow The begin row index to be operated.
 * @param optional nil or *CellsPutInsertWorksheetRowsOpts - Optional Parameters:
     * @param "TotalRows" (optional.Int32) -  Number of rows to be operated.
     * @param "UpdateReference" (optional.Bool) -  Indicates if update references in other worksheets.
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsPutInsertWorksheetRowsOpts struct { 
	Name string
	SheetName string
	Startrow int32
	TotalRows int32
	UpdateReference bool
	Folder string
	Storage string
}


func (a *CellsApiService) CellsPutInsertWorksheetRows(    localVarOptionals *CellsPutInsertWorksheetRowsOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/cells/rows"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("startrow", parameterToString(localVarOptionals.Startrow, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("totalRows", parameterToString(localVarOptionals.TotalRows.Value(), ""))
		localVarQueryParams.Add("totalRows", parameterToString(localVarOptionals.TotalRows, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("updateReference", parameterToString(localVarOptionals.UpdateReference.Value(), ""))
		localVarQueryParams.Add("updateReference", parameterToString(localVarOptionals.UpdateReference, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get cells list in a range by range name or row column indexes  
 * @param name workbook name
 * @param sheetName worksheet name
 * @param optional nil or *CellsRangesGetWorksheetCellsRangeValueOpts - Optional Parameters:
     * @param "Namerange" (optional.String) -  range name, for example: &#39;A1:B2&#39; or &#39;range_name1&#39;
     * @param "FirstRow" (optional.Int32) -  the first row of the range
     * @param "FirstColumn" (optional.Int32) -  the first column of the range
     * @param "RowCount" (optional.Int32) -  the count of rows in the range
     * @param "ColumnCount" (optional.Int32) -  the count of columns in the range
     * @param "Folder" (optional.String) -  Workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return RangeValueResponse
*/


type CellsRangesGetWorksheetCellsRangeValueOpts struct { 
	Name string
	SheetName string
	Namerange string
	FirstRow int32
	FirstColumn int32
	RowCount int32
	ColumnCount int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsRangesGetWorksheetCellsRangeValue(    localVarOptionals *CellsRangesGetWorksheetCellsRangeValueOpts) (RangeValueResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue RangeValueResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/ranges/value"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("namerange", parameterToString(localVarOptionals.Namerange.Value(), ""))
		localVarQueryParams.Add("namerange", parameterToString(localVarOptionals.Namerange, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("firstRow", parameterToString(localVarOptionals.FirstRow.Value(), ""))
		localVarQueryParams.Add("firstRow", parameterToString(localVarOptionals.FirstRow, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("firstColumn", parameterToString(localVarOptionals.FirstColumn.Value(), ""))
		localVarQueryParams.Add("firstColumn", parameterToString(localVarOptionals.FirstColumn, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("rowCount", parameterToString(localVarOptionals.RowCount.Value(), ""))
		localVarQueryParams.Add("rowCount", parameterToString(localVarOptionals.RowCount, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("columnCount", parameterToString(localVarOptionals.ColumnCount.Value(), ""))
		localVarQueryParams.Add("columnCount", parameterToString(localVarOptionals.ColumnCount, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Set column width of range
 * @param name
 * @param sheetName
 * @param value
 * @param optional nil or *CellsRangesPostWorksheetCellsRangeColumnWidthOpts - Optional Parameters:
     * @param "Range_" (optional.Interface of ModelRange) - 
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsRangesPostWorksheetCellsRangeColumnWidthOpts struct { 
	Name string
	SheetName string
	Value float64
	Range_ *ModelRange
	Folder string
	Storage string
}


func (a *CellsApiService) CellsRangesPostWorksheetCellsRangeColumnWidth(    localVarOptionals *CellsRangesPostWorksheetCellsRangeColumnWidthOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/ranges/columnWidth"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("value", parameterToString(localVarOptionals.Value, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Range_.IsSet() {
//		
//		localVarOptionalRange_, localVarOptionalRange_ok := localVarOptionals.Range_.Value().(ModelRange)
//		if !localVarOptionalRange_ok {
//				return localVarReturnValue, nil, reportError("range_ should be ModelRange")
//		}
//		localVarPostBody = &localVarOptionalRange_
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Range_ != nil {
		
		localVarPostBody = &localVarOptionals.Range_
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Combines a range of cells into a single cell.              
 * @param name workbook name
 * @param sheetName worksheet name
 * @param optional nil or *CellsRangesPostWorksheetCellsRangeMergeOpts - Optional Parameters:
     * @param "Range_" (optional.Interface of ModelRange) -  range in worksheet 
     * @param "Folder" (optional.String) -  Workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsRangesPostWorksheetCellsRangeMergeOpts struct { 
	Name string
	SheetName string
	Range_ *ModelRange
	Folder string
	Storage string
}


func (a *CellsApiService) CellsRangesPostWorksheetCellsRangeMerge(    localVarOptionals *CellsRangesPostWorksheetCellsRangeMergeOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/ranges/merge"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Range_.IsSet() {
//		
//		localVarOptionalRange_, localVarOptionalRange_ok := localVarOptionals.Range_.Value().(ModelRange)
//		if !localVarOptionalRange_ok {
//				return localVarReturnValue, nil, reportError("range_ should be ModelRange")
//		}
//		localVarPostBody = &localVarOptionalRange_
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Range_ != nil {
		
		localVarPostBody = &localVarOptionals.Range_
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Move the current range to the dest range.             
 * @param name workbook name
 * @param sheetName worksheet name
 * @param destRow The start row of the dest range.
 * @param destColumn The start column of the dest range.
 * @param optional nil or *CellsRangesPostWorksheetCellsRangeMoveToOpts - Optional Parameters:
     * @param "Range_" (optional.Interface of ModelRange) -  range in worksheet 
     * @param "Folder" (optional.String) -  Workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsRangesPostWorksheetCellsRangeMoveToOpts struct { 
	Name string
	SheetName string
	DestRow int32
	DestColumn int32
	Range_ *ModelRange
	Folder string
	Storage string
}


func (a *CellsApiService) CellsRangesPostWorksheetCellsRangeMoveTo(    localVarOptionals *CellsRangesPostWorksheetCellsRangeMoveToOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/ranges/moveto"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("destRow", parameterToString(localVarOptionals.DestRow, ""))
	localVarQueryParams.Add("destColumn", parameterToString(localVarOptionals.DestColumn, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Range_.IsSet() {
//		
//		localVarOptionalRange_, localVarOptionalRange_ok := localVarOptionals.Range_.Value().(ModelRange)
//		if !localVarOptionalRange_ok {
//				return localVarReturnValue, nil, reportError("range_ should be ModelRange")
//		}
//		localVarPostBody = &localVarOptionalRange_
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Range_ != nil {
		
		localVarPostBody = &localVarOptionals.Range_
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Sets outline border around a range of cells.
 * @param name workbook name
 * @param sheetName worksheet name
 * @param optional nil or *CellsRangesPostWorksheetCellsRangeOutlineBorderOpts - Optional Parameters:
     * @param "RangeOperate" (optional.Interface of RangeSetOutlineBorderRequest) -  Range Set OutlineBorder Request 
     * @param "Folder" (optional.String) -  Workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsRangesPostWorksheetCellsRangeOutlineBorderOpts struct { 
	Name string
	SheetName string
	RangeOperate *RangeSetOutlineBorderRequest
	Folder string
	Storage string
}


func (a *CellsApiService) CellsRangesPostWorksheetCellsRangeOutlineBorder(    localVarOptionals *CellsRangesPostWorksheetCellsRangeOutlineBorderOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/ranges/outlineBorder"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.RangeOperate.IsSet() {
//		
//		localVarOptionalRangeOperate, localVarOptionalRangeOperateok := localVarOptionals.RangeOperate.Value().(RangeSetOutlineBorderRequest)
//		if !localVarOptionalRangeOperateok {
//				return localVarReturnValue, nil, reportError("rangeOperate should be RangeSetOutlineBorderRequest")
//		}
//		localVarPostBody = &localVarOptionalRangeOperate
//	}
	if localVarOptionals != nil &&  &localVarOptionals.RangeOperate != nil {
		
		localVarPostBody = &localVarOptionals.RangeOperate
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService set row height of range
 * @param name
 * @param sheetName
 * @param value
 * @param optional nil or *CellsRangesPostWorksheetCellsRangeRowHeightOpts - Optional Parameters:
     * @param "Range_" (optional.Interface of ModelRange) - 
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsRangesPostWorksheetCellsRangeRowHeightOpts struct { 
	Name string
	SheetName string
	Value float64
	Range_ *ModelRange
	Folder string
	Storage string
}


func (a *CellsApiService) CellsRangesPostWorksheetCellsRangeRowHeight(    localVarOptionals *CellsRangesPostWorksheetCellsRangeRowHeightOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/ranges/rowHeight"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("value", parameterToString(localVarOptionals.Value, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Range_.IsSet() {
//		
//		localVarOptionalRange_, localVarOptionalRange_ok := localVarOptionals.Range_.Value().(ModelRange)
//		if !localVarOptionalRange_ok {
//				return localVarReturnValue, nil, reportError("range_ should be ModelRange")
//		}
//		localVarPostBody = &localVarOptionalRange_
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Range_ != nil {
		
		localVarPostBody = &localVarOptionals.Range_
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Sets the style of the range.             
 * @param name workbook name
 * @param sheetName worksheet name
 * @param optional nil or *CellsRangesPostWorksheetCellsRangeStyleOpts - Optional Parameters:
     * @param "RangeOperate" (optional.Interface of RangeSetStyleRequest) -  Range Set Style Request 
     * @param "Folder" (optional.String) -  Workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsRangesPostWorksheetCellsRangeStyleOpts struct { 
	Name string
	SheetName string
	RangeOperate *RangeSetStyleRequest
	Folder string
	Storage string
}


func (a *CellsApiService) CellsRangesPostWorksheetCellsRangeStyle(    localVarOptionals *CellsRangesPostWorksheetCellsRangeStyleOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/ranges/style"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.RangeOperate.IsSet() {
//		
//		localVarOptionalRangeOperate, localVarOptionalRangeOperateok := localVarOptionals.RangeOperate.Value().(RangeSetStyleRequest)
//		if !localVarOptionalRangeOperateok {
//				return localVarReturnValue, nil, reportError("rangeOperate should be RangeSetStyleRequest")
//		}
//		localVarPostBody = &localVarOptionalRangeOperate
//	}
	if localVarOptionals != nil &&  &localVarOptionals.RangeOperate != nil {
		
		localVarPostBody = &localVarOptionals.RangeOperate
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Unmerges merged cells of this range.             
 * @param name workbook name
 * @param sheetName worksheet name
 * @param optional nil or *CellsRangesPostWorksheetCellsRangeUnmergeOpts - Optional Parameters:
     * @param "Range_" (optional.Interface of ModelRange) -  range in worksheet 
     * @param "Folder" (optional.String) -  Workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsRangesPostWorksheetCellsRangeUnmergeOpts struct { 
	Name string
	SheetName string
	Range_ *ModelRange
	Folder string
	Storage string
}


func (a *CellsApiService) CellsRangesPostWorksheetCellsRangeUnmerge(    localVarOptionals *CellsRangesPostWorksheetCellsRangeUnmergeOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/ranges/unmerge"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Range_.IsSet() {
//		
//		localVarOptionalRange_, localVarOptionalRange_ok := localVarOptionals.Range_.Value().(ModelRange)
//		if !localVarOptionalRange_ok {
//				return localVarReturnValue, nil, reportError("range_ should be ModelRange")
//		}
//		localVarPostBody = &localVarOptionalRange_
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Range_ != nil {
		
		localVarPostBody = &localVarOptionals.Range_
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Puts a value into the range, if appropriate the value will be converted to other data type and cell&#39;s number format will be reset.             
 * @param name workbook name
 * @param sheetName worksheet name
 * @param value Input value
 * @param optional nil or *CellsRangesPostWorksheetCellsRangeValueOpts - Optional Parameters:
     * @param "Range_" (optional.Interface of ModelRange) -  range in worksheet 
     * @param "IsConverted" (optional.Bool) -  True: converted to other data type if appropriate.
     * @param "SetStyle" (optional.Bool) -  True: set the number format to cell&#39;s style when converting to other data type
     * @param "Folder" (optional.String) -  Workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsRangesPostWorksheetCellsRangeValueOpts struct { 
	Name string
	SheetName string
	Value string
	Range_ *ModelRange
	IsConverted bool
	SetStyle bool
	Folder string
	Storage string
}


func (a *CellsApiService) CellsRangesPostWorksheetCellsRangeValue(    localVarOptionals *CellsRangesPostWorksheetCellsRangeValueOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/ranges/value"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("Value", parameterToString(localVarOptionals.Value, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("isConverted", parameterToString(localVarOptionals.IsConverted.Value(), ""))
		localVarQueryParams.Add("isConverted", parameterToString(localVarOptionals.IsConverted, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("setStyle", parameterToString(localVarOptionals.SetStyle.Value(), ""))
		localVarQueryParams.Add("setStyle", parameterToString(localVarOptionals.SetStyle, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Range_.IsSet() {
//		
//		localVarOptionalRange_, localVarOptionalRange_ok := localVarOptionals.Range_.Value().(ModelRange)
//		if !localVarOptionalRange_ok {
//				return localVarReturnValue, nil, reportError("range_ should be ModelRange")
//		}
//		localVarPostBody = &localVarOptionalRange_
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Range_ != nil {
		
		localVarPostBody = &localVarOptionals.Range_
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService copy range in the worksheet
 * @param name workbook name
 * @param sheetName worksheet name
 * @param optional nil or *CellsRangesPostWorksheetCellsRangesOpts - Optional Parameters:
     * @param "RangeOperate" (optional.Interface of RangeCopyRequest) -  copydata,copystyle,copyto,copyvalue
     * @param "Folder" (optional.String) -  Workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsRangesPostWorksheetCellsRangesOpts struct { 
	Name string
	SheetName string
	RangeOperate *RangeCopyRequest
	Folder string
	Storage string
}


func (a *CellsApiService) CellsRangesPostWorksheetCellsRanges(    localVarOptionals *CellsRangesPostWorksheetCellsRangesOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/ranges"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.RangeOperate.IsSet() {
//		
//		localVarOptionalRangeOperate, localVarOptionalRangeOperateok := localVarOptionals.RangeOperate.Value().(RangeCopyRequest)
//		if !localVarOptionalRangeOperateok {
//				return localVarReturnValue, nil, reportError("rangeOperate should be RangeCopyRequest")
//		}
//		localVarPostBody = &localVarOptionalRangeOperate
//	}
	if localVarOptionals != nil &&  &localVarOptionals.RangeOperate != nil {
		
		localVarPostBody = &localVarOptionals.RangeOperate
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Convert document and save result to storage.
 * @param name The document name.
 * @param optional nil or *CellsSaveAsPostDocumentSaveAsOpts - Optional Parameters:
     * @param "SaveOptions" (optional.Interface of SaveOptions) -  Save options.
     * @param "Newfilename" (optional.String) -  The new file name.
     * @param "IsAutoFitRows" (optional.Bool) -  Autofit rows.
     * @param "IsAutoFitColumns" (optional.Bool) -  Autofit columns.
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return SaveResponse
*/


type CellsSaveAsPostDocumentSaveAsOpts struct { 
	Name string
	SaveOptions *SaveOptions
	Newfilename string
	IsAutoFitRows bool
	IsAutoFitColumns bool
	Folder string
	Storage string
}


func (a *CellsApiService) CellsSaveAsPostDocumentSaveAs(    localVarOptionals *CellsSaveAsPostDocumentSaveAsOpts) (SaveResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue SaveResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/SaveAs"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("newfilename", parameterToString(localVarOptionals.Newfilename.Value(), ""))
		localVarQueryParams.Add("newfilename", parameterToString(localVarOptionals.Newfilename, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("isAutoFitRows", parameterToString(localVarOptionals.IsAutoFitRows.Value(), ""))
		localVarQueryParams.Add("isAutoFitRows", parameterToString(localVarOptionals.IsAutoFitRows, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("isAutoFitColumns", parameterToString(localVarOptionals.IsAutoFitColumns.Value(), ""))
		localVarQueryParams.Add("isAutoFitColumns", parameterToString(localVarOptionals.IsAutoFitColumns, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.SaveOptions.IsSet() {
//		
//		localVarOptionalSaveOptions, localVarOptionalSaveOptionsok := localVarOptionals.SaveOptions.Value().(SaveOptions)
//		if !localVarOptionalSaveOptionsok {
//				return localVarReturnValue, nil, reportError("saveOptions should be SaveOptions")
//		}
//		localVarPostBody = &localVarOptionalSaveOptions
//	}
	if localVarOptionals != nil &&  &localVarOptionals.SaveOptions != nil {
		
		localVarPostBody = &localVarOptionals.SaveOptions
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Delete a shape in worksheet
 * @param name document name.
 * @param sheetName worksheet name.
 * @param shapeindex shape index in worksheet shapes.
 * @param optional nil or *CellsShapesDeleteWorksheetShapeOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsShapesDeleteWorksheetShapeOpts struct { 
	Name string
	SheetName string
	Shapeindex int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsShapesDeleteWorksheetShape(    localVarOptionals *CellsShapesDeleteWorksheetShapeOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/shapes/{shapeindex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"shapeindex"+"}", fmt.Sprintf("%v", localVarOptionals.Shapeindex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService delete all shapes in worksheet
 * @param name document name.
 * @param sheetName worksheet name.
 * @param optional nil or *CellsShapesDeleteWorksheetShapesOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsShapesDeleteWorksheetShapesOpts struct { 
	Name string
	SheetName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsShapesDeleteWorksheetShapes(    localVarOptionals *CellsShapesDeleteWorksheetShapesOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/shapes"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get worksheet shape
 * @param name document name.
 * @param sheetName worksheet name.
 * @param shapeindex shape index in worksheet shapes.
 * @param optional nil or *CellsShapesGetWorksheetShapeOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return ShapeResponse
*/


type CellsShapesGetWorksheetShapeOpts struct { 
	Name string
	SheetName string
	Shapeindex int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsShapesGetWorksheetShape(    localVarOptionals *CellsShapesGetWorksheetShapeOpts) (ShapeResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ShapeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/shapes/{shapeindex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"shapeindex"+"}", fmt.Sprintf("%v", localVarOptionals.Shapeindex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get worksheet shapes 
 * @param name document name.
 * @param sheetName worksheet name.
 * @param optional nil or *CellsShapesGetWorksheetShapesOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return ShapesResponse
*/


type CellsShapesGetWorksheetShapesOpts struct { 
	Name string
	SheetName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsShapesGetWorksheetShapes(    localVarOptionals *CellsShapesGetWorksheetShapesOpts) (ShapesResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ShapesResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/shapes"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Update a shape in worksheet
 * @param name document name.
 * @param sheetName worksheet name.
 * @param shapeindex shape index in worksheet shapes.
 * @param optional nil or *CellsShapesPostWorksheetShapeOpts - Optional Parameters:
     * @param "Dto" (optional.Interface of Shape) - 
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsShapesPostWorksheetShapeOpts struct { 
	Name string
	SheetName string
	Shapeindex int32
	Dto *Shape
	Folder string
	Storage string
}


func (a *CellsApiService) CellsShapesPostWorksheetShape(    localVarOptionals *CellsShapesPostWorksheetShapeOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/shapes/{shapeindex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"shapeindex"+"}", fmt.Sprintf("%v", localVarOptionals.Shapeindex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Dto.IsSet() {
//		
//		localVarOptionalDto, localVarOptionalDtook := localVarOptionals.Dto.Value().(Shape)
//		if !localVarOptionalDtook {
//				return localVarReturnValue, nil, reportError("dto should be Shape")
//		}
//		localVarPostBody = &localVarOptionalDto
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Dto != nil {
		
		localVarPostBody = &localVarOptionals.Dto
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Add shape in worksheet
 * @param name document name.
 * @param sheetName worksheet name.
 * @param drawingType shape object type
 * @param upperLeftRow Upper left row index.
 * @param upperLeftColumn Upper left column index.
 * @param top Represents the vertical offset of Spinner from its left row, in unit of pixel.
 * @param left Represents the horizontal offset of Spinner from its left column, in unit of pixel.
 * @param width Represents the height of Spinner, in unit of pixel.
 * @param height Represents the width of Spinner, in unit of pixel.
 * @param optional nil or *CellsShapesPutWorksheetShapeOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return ShapeResponse
*/


type CellsShapesPutWorksheetShapeOpts struct { 
	Name string
	SheetName string
	DrawingType string
	UpperLeftRow int32
	UpperLeftColumn int32
	Top int32
	Left int32
	Width int32
	Height int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsShapesPutWorksheetShape(    localVarOptionals *CellsShapesPutWorksheetShapeOpts) (ShapeResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ShapeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/shapes"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("DrawingType", parameterToString(localVarOptionals.DrawingType, ""))
	localVarQueryParams.Add("upperLeftRow", parameterToString(localVarOptionals.UpperLeftRow, ""))
	localVarQueryParams.Add("upperLeftColumn", parameterToString(localVarOptionals.UpperLeftColumn, ""))
	localVarQueryParams.Add("top", parameterToString(localVarOptionals.Top, ""))
	localVarQueryParams.Add("left", parameterToString(localVarOptionals.Left, ""))
	localVarQueryParams.Add("width", parameterToString(localVarOptionals.Width, ""))
	localVarQueryParams.Add("height", parameterToString(localVarOptionals.Height, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Run tasks  
 * @param taskData

@return interface{}
*/


type CellsTaskPostRunTaskOpts struct { 
	TaskData interface{}
}


func (a *CellsApiService) CellsTaskPostRunTask(    localVarOptionals *CellsTaskPostRunTaskOpts) (interface{}, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/task/runtask"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
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
	// body params
	localVarPostBody = &localVarOptionals.TaskData
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Decrypt document.
 * @param name The document name.
 * @param optional nil or *CellsWorkbookDeleteDecryptDocumentOpts - Optional Parameters:
     * @param "Encryption" (optional.Interface of WorkbookEncryptionRequest) -  Encryption settings, only password can be specified.
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsWorkbookDeleteDecryptDocumentOpts struct { 
	Name string
	Encryption *WorkbookEncryptionRequest
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorkbookDeleteDecryptDocument(    localVarOptionals *CellsWorkbookDeleteDecryptDocumentOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/encryption"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Encryption.IsSet() {
//		
//		localVarOptionalEncryption, localVarOptionalEncryptionok := localVarOptionals.Encryption.Value().(WorkbookEncryptionRequest)
//		if !localVarOptionalEncryptionok {
//				return localVarReturnValue, nil, reportError("encryption should be WorkbookEncryptionRequest")
//		}
//		localVarPostBody = &localVarOptionalEncryption
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Encryption != nil {
		
		localVarPostBody = &localVarOptionals.Encryption
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Unprotect document from changes.
 * @param name The document name.
 * @param optional nil or *CellsWorkbookDeleteDocumentUnprotectFromChangesOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsWorkbookDeleteDocumentUnprotectFromChangesOpts struct { 
	Name string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorkbookDeleteDocumentUnprotectFromChanges(    localVarOptionals *CellsWorkbookDeleteDocumentUnprotectFromChangesOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/writeProtection"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Unprotect document.
 * @param name The document name.
 * @param optional nil or *CellsWorkbookDeleteUnprotectDocumentOpts - Optional Parameters:
     * @param "Protection" (optional.Interface of WorkbookProtectionRequest) -  Protection settings, only password can be specified.
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsWorkbookDeleteUnprotectDocumentOpts struct { 
	Name string
	Protection *WorkbookProtectionRequest
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorkbookDeleteUnprotectDocument(    localVarOptionals *CellsWorkbookDeleteUnprotectDocumentOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/protection"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Protection.IsSet() {
//		
//		localVarOptionalProtection, localVarOptionalProtectionok := localVarOptionals.Protection.Value().(WorkbookProtectionRequest)
//		if !localVarOptionalProtectionok {
//				return localVarReturnValue, nil, reportError("protection should be WorkbookProtectionRequest")
//		}
//		localVarPostBody = &localVarOptionalProtection
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Protection != nil {
		
		localVarPostBody = &localVarOptionals.Protection
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Clean workbook&#39;s names.
 * @param name The workbook name.
 * @param nameName The name.
 * @param optional nil or *CellsWorkbookDeleteWorkbookNameOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsWorkbookDeleteWorkbookNameOpts struct { 
	Name string
	NameName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorkbookDeleteWorkbookName(    localVarOptionals *CellsWorkbookDeleteWorkbookNameOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/names/{nameName}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nameName"+"}", fmt.Sprintf("%v", localVarOptionals.NameName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Clean workbook&#39;s names.
 * @param name The workbook name.
 * @param optional nil or *CellsWorkbookDeleteWorkbookNamesOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsWorkbookDeleteWorkbookNamesOpts struct { 
	Name string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorkbookDeleteWorkbookNames(    localVarOptionals *CellsWorkbookDeleteWorkbookNamesOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/names"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Read workbook info or export.
 * @param name The document name.
 * @param optional nil or *CellsWorkbookGetWorkbookOpts - Optional Parameters:
     * @param "Password" (optional.String) -  The document password.
     * @param "Format" (optional.String) -  The exported file format.
     * @param "IsAutoFit" (optional.Bool) -  Set document rows to be autofit.
     * @param "OnlySaveTable" (optional.Bool) -  Only save table data.
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.
     * @param "OutPath" (optional.String) -  The document output folder.

@return *os.File
*/


type CellsWorkbookGetWorkbookOpts struct { 
	Name string
	Password string
	Format string
	IsAutoFit bool
	OnlySaveTable bool
	Folder string
	Storage string
	OutPath string
}


func (a *CellsApiService) CellsWorkbookGetWorkbook(    localVarOptionals *CellsWorkbookGetWorkbookOpts) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("password", parameterToString(localVarOptionals.Password.Value(), ""))
		localVarQueryParams.Add("password", parameterToString(localVarOptionals.Password, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("format", parameterToString(localVarOptionals.Format.Value(), ""))
		localVarQueryParams.Add("format", parameterToString(localVarOptionals.Format, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("isAutoFit", parameterToString(localVarOptionals.IsAutoFit.Value(), ""))
		localVarQueryParams.Add("isAutoFit", parameterToString(localVarOptionals.IsAutoFit, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("onlySaveTable", parameterToString(localVarOptionals.OnlySaveTable.Value(), ""))
		localVarQueryParams.Add("onlySaveTable", parameterToString(localVarOptionals.OnlySaveTable, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("outPath", parameterToString(localVarOptionals.OutPath.Value(), ""))
		localVarQueryParams.Add("outPath", parameterToString(localVarOptionals.OutPath, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Read workbook default style info.
 * @param name The workbook name.
 * @param optional nil or *CellsWorkbookGetWorkbookDefaultStyleOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return StyleResponse
*/


type CellsWorkbookGetWorkbookDefaultStyleOpts struct { 
	Name string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorkbookGetWorkbookDefaultStyle(    localVarOptionals *CellsWorkbookGetWorkbookDefaultStyleOpts) (StyleResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue StyleResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/defaultstyle"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Read workbook&#39;s name.
 * @param name The workbook name.
 * @param nameName The name.
 * @param optional nil or *CellsWorkbookGetWorkbookNameOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return NameResponse
*/


type CellsWorkbookGetWorkbookNameOpts struct { 
	Name string
	NameName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorkbookGetWorkbookName(    localVarOptionals *CellsWorkbookGetWorkbookNameOpts) (NameResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue NameResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/names/{nameName}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nameName"+"}", fmt.Sprintf("%v", localVarOptionals.NameName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get workbook&#39;s name value.
 * @param name The workbook name.
 * @param nameName The name.
 * @param optional nil or *CellsWorkbookGetWorkbookNameValueOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return RangeValueResponse
*/


type CellsWorkbookGetWorkbookNameValueOpts struct { 
	Name string
	NameName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorkbookGetWorkbookNameValue(    localVarOptionals *CellsWorkbookGetWorkbookNameValueOpts) (RangeValueResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue RangeValueResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/names/{nameName}/value"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nameName"+"}", fmt.Sprintf("%v", localVarOptionals.NameName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Read workbook&#39;s names.
 * @param name The workbook name.
 * @param optional nil or *CellsWorkbookGetWorkbookNamesOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return NamesResponse
*/


type CellsWorkbookGetWorkbookNamesOpts struct { 
	Name string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorkbookGetWorkbookNames(    localVarOptionals *CellsWorkbookGetWorkbookNamesOpts) (NamesResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue NamesResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/names"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get Workbook Settings DTO
 * @param name Document name.
 * @param optional nil or *CellsWorkbookGetWorkbookSettingsOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return WorkbookSettingsResponse
*/


type CellsWorkbookGetWorkbookSettingsOpts struct { 
	Name string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorkbookGetWorkbookSettings(    localVarOptionals *CellsWorkbookGetWorkbookSettingsOpts) (WorkbookSettingsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue WorkbookSettingsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/settings"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Read workbook&#39;s text items.
 * @param name The workbook name.
 * @param optional nil or *CellsWorkbookGetWorkbookTextItemsOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return TextItemsResponse
*/


type CellsWorkbookGetWorkbookTextItemsOpts struct { 
	Name string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorkbookGetWorkbookTextItems(    localVarOptionals *CellsWorkbookGetWorkbookTextItemsOpts) (TextItemsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue TextItemsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/textItems"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Autofit workbook rows.
 * @param name Document name.
 * @param optional nil or *CellsWorkbookPostAutofitWorkbookRowsOpts - Optional Parameters:
     * @param "AutoFitterOptions" (optional.Interface of AutoFitterOptions) -  Auto Fitter Options.
     * @param "StartRow" (optional.Int32) -  Start row.
     * @param "EndRow" (optional.Int32) -  End row.
     * @param "OnlyAuto" (optional.Bool) -  Only auto.
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsWorkbookPostAutofitWorkbookRowsOpts struct { 
	Name string
	AutoFitterOptions *AutoFitterOptions
	StartRow int32
	EndRow int32
	OnlyAuto bool
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorkbookPostAutofitWorkbookRows(    localVarOptionals *CellsWorkbookPostAutofitWorkbookRowsOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/autofitrows"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("startRow", parameterToString(localVarOptionals.StartRow.Value(), ""))
		localVarQueryParams.Add("startRow", parameterToString(localVarOptionals.StartRow, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("endRow", parameterToString(localVarOptionals.EndRow.Value(), ""))
		localVarQueryParams.Add("endRow", parameterToString(localVarOptionals.EndRow, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("onlyAuto", parameterToString(localVarOptionals.OnlyAuto.Value(), ""))
		localVarQueryParams.Add("onlyAuto", parameterToString(localVarOptionals.OnlyAuto, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.AutoFitterOptions.IsSet() {
//		
//		localVarOptionalAutoFitterOptions, localVarOptionalAutoFitterOptionsok := localVarOptionals.AutoFitterOptions.Value().(AutoFitterOptions)
//		if !localVarOptionalAutoFitterOptionsok {
//				return localVarReturnValue, nil, reportError("autoFitterOptions should be AutoFitterOptions")
//		}
//		localVarPostBody = &localVarOptionalAutoFitterOptions
//	}
	if localVarOptionals != nil &&  &localVarOptionals.AutoFitterOptions != nil {
		
		localVarPostBody = &localVarOptionals.AutoFitterOptions
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Encript document.
 * @param name The document name.
 * @param optional nil or *CellsWorkbookPostEncryptDocumentOpts - Optional Parameters:
     * @param "Encryption" (optional.Interface of WorkbookEncryptionRequest) -  Encryption parameters.
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsWorkbookPostEncryptDocumentOpts struct { 
	Name string
	Encryption *WorkbookEncryptionRequest
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorkbookPostEncryptDocument(    localVarOptionals *CellsWorkbookPostEncryptDocumentOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/encryption"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Encryption.IsSet() {
//		
//		localVarOptionalEncryption, localVarOptionalEncryptionok := localVarOptionals.Encryption.Value().(WorkbookEncryptionRequest)
//		if !localVarOptionalEncryptionok {
//				return localVarReturnValue, nil, reportError("encryption should be WorkbookEncryptionRequest")
//		}
//		localVarPostBody = &localVarOptionalEncryption
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Encryption != nil {
		
		localVarPostBody = &localVarOptionals.Encryption
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService
 * @param name
 * @param importData
 * @param optional nil or *CellsWorkbookPostImportDataOpts - Optional Parameters:
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsWorkbookPostImportDataOpts struct { 
	Name string
	ImportData interface{}
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorkbookPostImportData(    localVarOptionals *CellsWorkbookPostImportDataOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/importdata"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
	localVarPostBody = localVarOptionals.ImportData
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Protect document.
 * @param name The document name.
 * @param optional nil or *CellsWorkbookPostProtectDocumentOpts - Optional Parameters:
     * @param "Protection" (optional.Interface of WorkbookProtectionRequest) -  The protection settings.
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsWorkbookPostProtectDocumentOpts struct { 
	Name string
	Protection *WorkbookProtectionRequest
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorkbookPostProtectDocument(    localVarOptionals *CellsWorkbookPostProtectDocumentOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/protection"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Protection.IsSet() {
//		
//		localVarOptionalProtection, localVarOptionalProtectionok := localVarOptionals.Protection.Value().(WorkbookProtectionRequest)
//		if !localVarOptionalProtectionok {
//				return localVarReturnValue, nil, reportError("protection should be WorkbookProtectionRequest")
//		}
//		localVarPostBody = &localVarOptionalProtection
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Protection != nil {
		
		localVarPostBody = &localVarOptionals.Protection
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Calculate all formulas in workbook.
 * @param name Document name.
 * @param optional nil or *CellsWorkbookPostWorkbookCalculateFormulaOpts - Optional Parameters:
     * @param "Options" (optional.Interface of CalculationOptions) -  Calculation Options.
     * @param "IgnoreError" (optional.Bool) -  ignore Error.
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsWorkbookPostWorkbookCalculateFormulaOpts struct { 
	Name string
	Options *CalculationOptions
	IgnoreError bool
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorkbookPostWorkbookCalculateFormula(    localVarOptionals *CellsWorkbookPostWorkbookCalculateFormulaOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/calculateformula"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("ignoreError", parameterToString(localVarOptionals.IgnoreError.Value(), ""))
		localVarQueryParams.Add("ignoreError", parameterToString(localVarOptionals.IgnoreError, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Options.IsSet() {
//		
//		localVarOptionalOptions, localVarOptionalOptionsok := localVarOptionals.Options.Value().(CalculationOptions)
//		if !localVarOptionalOptionsok {
//				return localVarReturnValue, nil, reportError("options should be CalculationOptions")
//		}
//		localVarPostBody = &localVarOptionalOptions
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Options != nil {
		
		localVarPostBody = &localVarOptionals.Options
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Smart marker processing result.
 * @param name The workbook name.
 * @param optional nil or *CellsWorkbookPostWorkbookGetSmartMarkerResultOpts - Optional Parameters:
     * @param "XmlFile" (optional.String) -  The xml file full path, if empty the data is read from request body.
     * @param "Folder" (optional.String) -  The workbook folder full path.
     * @param "Storage" (optional.String) -  storage name.
     * @param "OutPath" (optional.String) -  Path to save result

@return *os.File
*/


type CellsWorkbookPostWorkbookGetSmartMarkerResultOpts struct { 
	Name string
	XmlFile string
	Folder string
	Storage string
	OutPath string
}


func (a *CellsApiService) CellsWorkbookPostWorkbookGetSmartMarkerResult(    localVarOptionals *CellsWorkbookPostWorkbookGetSmartMarkerResultOpts) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/smartmarker"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("xmlFile", parameterToString(localVarOptionals.XmlFile.Value(), ""))
		localVarQueryParams.Add("xmlFile", parameterToString(localVarOptionals.XmlFile, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("outPath", parameterToString(localVarOptionals.OutPath.Value(), ""))
		localVarQueryParams.Add("outPath", parameterToString(localVarOptionals.OutPath, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Update Workbook setting 
 * @param name Document name.
 * @param optional nil or *CellsWorkbookPostWorkbookSettingsOpts - Optional Parameters:
     * @param "Settings" (optional.Interface of WorkbookSettings) -  Workbook Setting DTO
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsWorkbookPostWorkbookSettingsOpts struct { 
	Name string
	Settings *WorkbookSettings
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorkbookPostWorkbookSettings(    localVarOptionals *CellsWorkbookPostWorkbookSettingsOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/settings"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Settings.IsSet() {
//		
//		localVarOptionalSettings, localVarOptionalSettingsok := localVarOptionals.Settings.Value().(WorkbookSettings)
//		if !localVarOptionalSettingsok {
//				return localVarReturnValue, nil, reportError("settings should be WorkbookSettings")
//		}
//		localVarPostBody = &localVarOptionalSettings
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Settings != nil {
		
		localVarPostBody = &localVarOptionals.Settings
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Split workbook.
 * @param name The workbook name.
 * @param optional nil or *CellsWorkbookPostWorkbookSplitOpts - Optional Parameters:
     * @param "Format" (optional.String) -  Split format.
     * @param "From" (optional.Int32) -  Start worksheet index.
     * @param "To" (optional.Int32) -  End worksheet index.
     * @param "HorizontalResolution" (optional.Int32) -  Image horizontal resolution.
     * @param "VerticalResolution" (optional.Int32) -  Image vertical resolution.
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return SplitResultResponse
*/


type CellsWorkbookPostWorkbookSplitOpts struct { 
	Name string
	Format string
	From int32
	To int32
	HorizontalResolution int32
	VerticalResolution int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorkbookPostWorkbookSplit(    localVarOptionals *CellsWorkbookPostWorkbookSplitOpts) (SplitResultResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue SplitResultResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/split"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("format", parameterToString(localVarOptionals.Format.Value(), ""))
		localVarQueryParams.Add("format", parameterToString(localVarOptionals.Format, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("from", parameterToString(localVarOptionals.From.Value(), ""))
		localVarQueryParams.Add("from", parameterToString(localVarOptionals.From, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("to", parameterToString(localVarOptionals.To.Value(), ""))
		localVarQueryParams.Add("to", parameterToString(localVarOptionals.To, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("horizontalResolution", parameterToString(localVarOptionals.HorizontalResolution.Value(), ""))
		localVarQueryParams.Add("horizontalResolution", parameterToString(localVarOptionals.HorizontalResolution, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("verticalResolution", parameterToString(localVarOptionals.VerticalResolution.Value(), ""))
		localVarQueryParams.Add("verticalResolution", parameterToString(localVarOptionals.VerticalResolution, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Merge workbooks.
 * @param name Workbook name.
 * @param mergeWith The workbook to merge with.
 * @param optional nil or *CellsWorkbookPostWorkbooksMergeOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Source workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return WorkbookResponse
*/


type CellsWorkbookPostWorkbooksMergeOpts struct { 
	Name string
	MergeWith string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorkbookPostWorkbooksMerge(    localVarOptionals *CellsWorkbookPostWorkbooksMergeOpts) (WorkbookResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue WorkbookResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/merge"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("mergeWith", parameterToString(localVarOptionals.MergeWith, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Replace text.
 * @param name Document name.
 * @param oldValue The old value.
 * @param newValue The new value.
 * @param optional nil or *CellsWorkbookPostWorkbooksTextReplaceOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return WorkbookReplaceResponse
*/


type CellsWorkbookPostWorkbooksTextReplaceOpts struct { 
	Name string
	OldValue string
	NewValue string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorkbookPostWorkbooksTextReplace(    localVarOptionals *CellsWorkbookPostWorkbooksTextReplaceOpts) (WorkbookReplaceResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue WorkbookReplaceResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/replaceText"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("oldValue", parameterToString(localVarOptionals.OldValue, ""))
	localVarQueryParams.Add("newValue", parameterToString(localVarOptionals.NewValue, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Search text.
 * @param name Document name.
 * @param text Text sample.
 * @param optional nil or *CellsWorkbookPostWorkbooksTextSearchOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return TextItemsResponse
*/


type CellsWorkbookPostWorkbooksTextSearchOpts struct { 
	Name string
	Text string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorkbookPostWorkbooksTextSearch(    localVarOptionals *CellsWorkbookPostWorkbooksTextSearchOpts) (TextItemsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue TextItemsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/findText"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("text", parameterToString(localVarOptionals.Text, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Convert workbook from request content to some format.
 * @param workbook
 * @param optional nil or *CellsWorkbookPutConvertWorkbookOpts - Optional Parameters:
     * @param "Format" (optional.String) -  The format to convert.
     * @param "Password" (optional.String) -  The workbook password.
     * @param "OutPath" (optional.String) -  Path to save result

@return *os.File
*/


type CellsWorkbookPutConvertWorkbookOpts struct { 

	Format string
	Password string
	OutPath string
}


func (a *CellsApiService) CellsWorkbookPutConvertWorkbook(   workbook *os.File ,    localVarOptionals *CellsWorkbookPutConvertWorkbookOpts) ( []byte, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/convert"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("format", parameterToString(localVarOptionals.Format.Value(), ""))
		localVarQueryParams.Add("format", parameterToString(localVarOptionals.Format, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("password", parameterToString(localVarOptionals.Password.Value(), ""))
		localVarQueryParams.Add("password", parameterToString(localVarOptionals.Password, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("outPath", parameterToString(localVarOptionals.OutPath.Value(), ""))
		localVarQueryParams.Add("outPath", parameterToString(localVarOptionals.OutPath, ""))
	}
	// to determine the Content-Type header
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
	var localVarFile (*os.File) =  workbook
	if localVarFile != nil {
		fbs, _ := ioutil.ReadAll(localVarFile)
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body)
    if err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }
	
	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Protect document from changes.
 * @param name Document name.
 * @param optional nil or *CellsWorkbookPutDocumentProtectFromChangesOpts - Optional Parameters:
     * @param "Password" (optional.Interface of PasswordRequest) -  Modification password.
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsWorkbookPutDocumentProtectFromChangesOpts struct { 
	Name string
	Password *PasswordRequest
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorkbookPutDocumentProtectFromChanges(    localVarOptionals *CellsWorkbookPutDocumentProtectFromChangesOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/writeProtection"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Password.IsSet() {
//		
//		localVarOptionalPassword, localVarOptionalPasswordok := localVarOptionals.Password.Value().(PasswordRequest)
//		if !localVarOptionalPasswordok {
//				return localVarReturnValue, nil, reportError("password should be PasswordRequest")
//		}
//		localVarPostBody = &localVarOptionalPassword
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Password != nil {
		
		localVarPostBody = &localVarOptionals.Password
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Create new workbook using deferent methods.
 * @param name The new document name.
 * @param optional nil or *CellsWorkbookPutWorkbookCreateOpts - Optional Parameters:
     * @param "TemplateFile" (optional.String) -  The template file, if the data not provided default workbook is created.
     * @param "DataFile" (optional.String) -  Smart marker data file, if the data not provided the request content is checked for the data.
     * @param "Folder" (optional.String) -  The new document folder.
     * @param "Storage" (optional.String) -  storage name.

@return WorkbookResponse
*/


type CellsWorkbookPutWorkbookCreateOpts struct { 
	Name string
	TemplateFile string
	DataFile string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorkbookPutWorkbookCreate(    localVarOptionals *CellsWorkbookPutWorkbookCreateOpts) (WorkbookResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue WorkbookResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("templateFile", parameterToString(localVarOptionals.TemplateFile.Value(), ""))
		localVarQueryParams.Add("templateFile", parameterToString(localVarOptionals.TemplateFile, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("dataFile", parameterToString(localVarOptionals.DataFile.Value(), ""))
		localVarQueryParams.Add("dataFile", parameterToString(localVarOptionals.DataFile, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Delete worksheet validation by index.
 * @param name Document name.
 * @param sheetName Worksheet name.
 * @param validationIndex The validation index.
 * @param optional nil or *CellsWorksheetValidationsDeleteWorksheetValidationOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return ValidationResponse
*/


type CellsWorksheetValidationsDeleteWorksheetValidationOpts struct { 
	Name string
	SheetName string
	ValidationIndex int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetValidationsDeleteWorksheetValidation(    localVarOptionals *CellsWorksheetValidationsDeleteWorksheetValidationOpts) (ValidationResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ValidationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/validations/{validationIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"validationIndex"+"}", fmt.Sprintf("%v", localVarOptionals.ValidationIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Clear all validation in worksheet.
 * @param name Document name.
 * @param sheetName Worksheet name.
 * @param optional nil or *CellsWorksheetValidationsDeleteWorksheetValidationsOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsWorksheetValidationsDeleteWorksheetValidationsOpts struct { 
	Name string
	SheetName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetValidationsDeleteWorksheetValidations(    localVarOptionals *CellsWorksheetValidationsDeleteWorksheetValidationsOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/validations"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get worksheet validation by index.
 * @param name Document name.
 * @param sheetName Worksheet name.
 * @param validationIndex The validation index.
 * @param optional nil or *CellsWorksheetValidationsGetWorksheetValidationOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return ValidationResponse
*/


type CellsWorksheetValidationsGetWorksheetValidationOpts struct { 
	Name string
	SheetName string
	ValidationIndex int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetValidationsGetWorksheetValidation(    localVarOptionals *CellsWorksheetValidationsGetWorksheetValidationOpts) (ValidationResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ValidationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/validations/{validationIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"validationIndex"+"}", fmt.Sprintf("%v", localVarOptionals.ValidationIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get worksheet validations.
 * @param name Document name.
 * @param sheetName Worksheet name.
 * @param optional nil or *CellsWorksheetValidationsGetWorksheetValidationsOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document folder.
     * @param "Storage" (optional.String) -  storage name.

@return ValidationsResponse
*/


type CellsWorksheetValidationsGetWorksheetValidationsOpts struct { 
	Name string
	SheetName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetValidationsGetWorksheetValidations(    localVarOptionals *CellsWorksheetValidationsGetWorksheetValidationsOpts) (ValidationsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ValidationsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/validations"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Update worksheet validation by index.
 * @param name Document name.
 * @param sheetName Worksheet name.
 * @param validationIndex The validation index.
 * @param optional nil or *CellsWorksheetValidationsPostWorksheetValidationOpts - Optional Parameters:
     * @param "Validation" (optional.Interface of Validation) - 
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return ValidationResponse
*/


type CellsWorksheetValidationsPostWorksheetValidationOpts struct { 
	Name string
	SheetName string
	ValidationIndex int32
	Validation *Validation
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetValidationsPostWorksheetValidation(    localVarOptionals *CellsWorksheetValidationsPostWorksheetValidationOpts) (ValidationResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ValidationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/validations/{validationIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"validationIndex"+"}", fmt.Sprintf("%v", localVarOptionals.ValidationIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Validation.IsSet() {
//		
//		localVarOptionalValidation, localVarOptionalValidationok := localVarOptionals.Validation.Value().(Validation)
//		if !localVarOptionalValidationok {
//				return localVarReturnValue, nil, reportError("validation should be Validation")
//		}
//		localVarPostBody = &localVarOptionalValidation
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Validation != nil {
		
		localVarPostBody = &localVarOptionals.Validation
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Add worksheet validation at index.
 * @param name Document name.
 * @param sheetName Worksheet name.
 * @param optional nil or *CellsWorksheetValidationsPutWorksheetValidationOpts - Optional Parameters:
     * @param "Range_" (optional.String) -  Specified cells area
     * @param "Validation" (optional.Interface of Validation) -  validation
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return ValidationResponse
*/


type CellsWorksheetValidationsPutWorksheetValidationOpts struct { 
	Name string
	SheetName string
	Range_ string
	Validation *Validation
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetValidationsPutWorksheetValidation(    localVarOptionals *CellsWorksheetValidationsPutWorksheetValidationOpts) (ValidationResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ValidationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/validations"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("range", parameterToString(localVarOptionals.Range_.Value(), ""))
		localVarQueryParams.Add("range", parameterToString(localVarOptionals.Range_, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Validation.IsSet() {
//		
//		localVarOptionalValidation, localVarOptionalValidationok := localVarOptionals.Validation.Value().(Validation)
//		if !localVarOptionalValidationok {
//				return localVarReturnValue, nil, reportError("validation should be Validation")
//		}
//		localVarPostBody = &localVarOptionalValidation
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Validation != nil {
		
		localVarPostBody = &localVarOptionals.Validation
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Unprotect worksheet.
 * @param name Document name.
 * @param sheetName The worksheet name.
 * @param optional nil or *CellsWorksheetsDeleteUnprotectWorksheetOpts - Optional Parameters:
     * @param "ProtectParameter" (optional.Interface of ProtectSheetParameter) -  with protection settings. Only password is used here.
     * @param "Folder" (optional.String) -  Document folder.
     * @param "Storage" (optional.String) -  storage name.

@return WorksheetResponse
*/


type CellsWorksheetsDeleteUnprotectWorksheetOpts struct { 
	Name string
	SheetName string
	ProtectParameter *ProtectSheetParameter
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetsDeleteUnprotectWorksheet(    localVarOptionals *CellsWorksheetsDeleteUnprotectWorksheetOpts) (WorksheetResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue WorksheetResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/protection"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.ProtectParameter.IsSet() {
//		
//		localVarOptionalProtectParameter, localVarOptionalProtectParameterok := localVarOptionals.ProtectParameter.Value().(ProtectSheetParameter)
//		if !localVarOptionalProtectParameterok {
//				return localVarReturnValue, nil, reportError("protectParameter should be ProtectSheetParameter")
//		}
//		localVarPostBody = &localVarOptionalProtectParameter
//	}
	if localVarOptionals != nil &&  &localVarOptionals.ProtectParameter != nil {
		
		localVarPostBody = &localVarOptionals.ProtectParameter
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Delete worksheet.
 * @param name Document name.
 * @param sheetName The worksheet name.
 * @param optional nil or *CellsWorksheetsDeleteWorksheetOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return WorksheetsResponse
*/


type CellsWorksheetsDeleteWorksheetOpts struct { 
	Name string
	SheetName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetsDeleteWorksheet(    localVarOptionals *CellsWorksheetsDeleteWorksheetOpts) (WorksheetsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue WorksheetsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Set worksheet background image.
 * @param name
 * @param sheetName
 * @param optional nil or *CellsWorksheetsDeleteWorksheetBackgroundOpts - Optional Parameters:
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsWorksheetsDeleteWorksheetBackgroundOpts struct { 
	Name string
	SheetName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetsDeleteWorksheetBackground(    localVarOptionals *CellsWorksheetsDeleteWorksheetBackgroundOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/background"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Delete worksheet&#39;s cell comment.
 * @param name The document name.
 * @param sheetName The worksheet name.
 * @param cellName The cell name
 * @param optional nil or *CellsWorksheetsDeleteWorksheetCommentOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsWorksheetsDeleteWorksheetCommentOpts struct { 
	Name string
	SheetName string
	CellName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetsDeleteWorksheetComment(    localVarOptionals *CellsWorksheetsDeleteWorksheetCommentOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/comments/{cellName}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cellName"+"}", fmt.Sprintf("%v", localVarOptionals.CellName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Delete all comments for worksheet.
 * @param name
 * @param sheetName
 * @param optional nil or *CellsWorksheetsDeleteWorksheetCommentsOpts - Optional Parameters:
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsWorksheetsDeleteWorksheetCommentsOpts struct { 
	Name string
	SheetName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetsDeleteWorksheetComments(    localVarOptionals *CellsWorksheetsDeleteWorksheetCommentsOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/comments"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Unfreeze panes
 * @param name
 * @param sheetName
 * @param row
 * @param column
 * @param freezedRows
 * @param freezedColumns
 * @param optional nil or *CellsWorksheetsDeleteWorksheetFreezePanesOpts - Optional Parameters:
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsWorksheetsDeleteWorksheetFreezePanesOpts struct { 
	Name string
	SheetName string
	Row int32
	Column int32
	FreezedRows int32
	FreezedColumns int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetsDeleteWorksheetFreezePanes(    localVarOptionals *CellsWorksheetsDeleteWorksheetFreezePanesOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/freezepanes"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("row", parameterToString(localVarOptionals.Row, ""))
	localVarQueryParams.Add("column", parameterToString(localVarOptionals.Column, ""))
	localVarQueryParams.Add("freezedRows", parameterToString(localVarOptionals.FreezedRows, ""))
	localVarQueryParams.Add("freezedColumns", parameterToString(localVarOptionals.FreezedColumns, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Read worksheets ranges info.
 * @param name Document name.
 * @param optional nil or *CellsWorksheetsGetNamedRangesOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document folder.
     * @param "Storage" (optional.String) -  storage name.

@return RangesResponse
*/


type CellsWorksheetsGetNamedRangesOpts struct { 
	Name string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetsGetNamedRanges(    localVarOptionals *CellsWorksheetsGetNamedRangesOpts) (RangesResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue RangesResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/ranges"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Read worksheet info or export.
 * @param name The document name.
 * @param sheetName The worksheet name.
 * @param optional nil or *CellsWorksheetsGetWorksheetOpts - Optional Parameters:
     * @param "Format" (optional.String) -  The exported file format.
     * @param "VerticalResolution" (optional.Int32) -  Image vertical resolution.
     * @param "HorizontalResolution" (optional.Int32) -  Image horizontal resolution.
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return *os.File
*/


type CellsWorksheetsGetWorksheetOpts struct { 
	Name string
	SheetName string
	Format string
	VerticalResolution int32
	HorizontalResolution int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetsGetWorksheet(    localVarOptionals *CellsWorksheetsGetWorksheetOpts) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("format", parameterToString(localVarOptionals.Format.Value(), ""))
		localVarQueryParams.Add("format", parameterToString(localVarOptionals.Format, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("verticalResolution", parameterToString(localVarOptionals.VerticalResolution.Value(), ""))
		localVarQueryParams.Add("verticalResolution", parameterToString(localVarOptionals.VerticalResolution, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("horizontalResolution", parameterToString(localVarOptionals.HorizontalResolution.Value(), ""))
		localVarQueryParams.Add("horizontalResolution", parameterToString(localVarOptionals.HorizontalResolution, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Calculate formula value.
 * @param name Document name.
 * @param sheetName Worksheet name.
 * @param formula The formula.
 * @param optional nil or *CellsWorksheetsGetWorksheetCalculateFormulaOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return SingleValueResponse
*/


type CellsWorksheetsGetWorksheetCalculateFormulaOpts struct { 
	Name string
	SheetName string
	Formula string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetsGetWorksheetCalculateFormula(    localVarOptionals *CellsWorksheetsGetWorksheetCalculateFormulaOpts) (SingleValueResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue SingleValueResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/formulaResult"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("formula", parameterToString(localVarOptionals.Formula, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get worksheet comment by cell name.
 * @param name The document name.
 * @param sheetName The worksheet name.
 * @param cellName The cell name
 * @param optional nil or *CellsWorksheetsGetWorksheetCommentOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return CommentResponse
*/


type CellsWorksheetsGetWorksheetCommentOpts struct { 
	Name string
	SheetName string
	CellName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetsGetWorksheetComment(    localVarOptionals *CellsWorksheetsGetWorksheetCommentOpts) (CommentResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CommentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/comments/{cellName}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cellName"+"}", fmt.Sprintf("%v", localVarOptionals.CellName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get worksheet comments.
 * @param name Workbook name.
 * @param sheetName The worksheet name.
 * @param optional nil or *CellsWorksheetsGetWorksheetCommentsOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return CommentsResponse
*/


type CellsWorksheetsGetWorksheetCommentsOpts struct { 
	Name string
	SheetName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetsGetWorksheetComments(    localVarOptionals *CellsWorksheetsGetWorksheetCommentsOpts) (CommentsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CommentsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/comments"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get worksheet merged cell by its index.
 * @param name Document name.
 * @param sheetName Worksheet name.
 * @param mergedCellIndex Merged cell index.
 * @param optional nil or *CellsWorksheetsGetWorksheetMergedCellOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document folder.
     * @param "Storage" (optional.String) -  storage name.

@return MergedCellResponse
*/


type CellsWorksheetsGetWorksheetMergedCellOpts struct { 
	Name string
	SheetName string
	MergedCellIndex int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetsGetWorksheetMergedCell(    localVarOptionals *CellsWorksheetsGetWorksheetMergedCellOpts) (MergedCellResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue MergedCellResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/mergedCells/{mergedCellIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mergedCellIndex"+"}", fmt.Sprintf("%v", localVarOptionals.MergedCellIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get worksheet merged cells.
 * @param name Document name.
 * @param sheetName The workseet name.
 * @param optional nil or *CellsWorksheetsGetWorksheetMergedCellsOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document folder.
     * @param "Storage" (optional.String) -  storage name.

@return MergedCellsResponse
*/


type CellsWorksheetsGetWorksheetMergedCellsOpts struct { 
	Name string
	SheetName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetsGetWorksheetMergedCells(    localVarOptionals *CellsWorksheetsGetWorksheetMergedCellsOpts) (MergedCellsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue MergedCellsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/mergedCells"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get worksheet text items.
 * @param name Workbook name.
 * @param sheetName The worksheet name.
 * @param optional nil or *CellsWorksheetsGetWorksheetTextItemsOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The workbook&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return TextItemsResponse
*/


type CellsWorksheetsGetWorksheetTextItemsOpts struct { 
	Name string
	SheetName string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetsGetWorksheetTextItems(    localVarOptionals *CellsWorksheetsGetWorksheetTextItemsOpts) (TextItemsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue TextItemsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/textItems"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Read worksheets info.
 * @param name Document name.
 * @param optional nil or *CellsWorksheetsGetWorksheetsOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document folder.
     * @param "Storage" (optional.String) -  storage name.

@return WorksheetsResponse
*/


type CellsWorksheetsGetWorksheetsOpts struct { 
	Name string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetsGetWorksheets(    localVarOptionals *CellsWorksheetsGetWorksheetsOpts) (WorksheetsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue WorksheetsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService
 * @param name
 * @param sheetName
 * @param firstColumn
 * @param lastColumn
 * @param optional nil or *CellsWorksheetsPostAutofitWorksheetColumnsOpts - Optional Parameters:
     * @param "AutoFitterOptions" (optional.Interface of AutoFitterOptions) - 
     * @param "FirstRow" (optional.Int32) - 
     * @param "LastRow" (optional.Int32) - 
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsWorksheetsPostAutofitWorksheetColumnsOpts struct { 
	Name string
	SheetName string
	FirstColumn int32
	LastColumn int32
	AutoFitterOptions *AutoFitterOptions
	FirstRow int32
	LastRow int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetsPostAutofitWorksheetColumns(    localVarOptionals *CellsWorksheetsPostAutofitWorksheetColumnsOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/autofitcolumns"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("firstColumn", parameterToString(localVarOptionals.FirstColumn, ""))
	localVarQueryParams.Add("lastColumn", parameterToString(localVarOptionals.LastColumn, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("firstRow", parameterToString(localVarOptionals.FirstRow.Value(), ""))
		localVarQueryParams.Add("firstRow", parameterToString(localVarOptionals.FirstRow, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("lastRow", parameterToString(localVarOptionals.LastRow.Value(), ""))
		localVarQueryParams.Add("lastRow", parameterToString(localVarOptionals.LastRow, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.AutoFitterOptions.IsSet() {
//		
//		localVarOptionalAutoFitterOptions, localVarOptionalAutoFitterOptionsok := localVarOptionals.AutoFitterOptions.Value().(AutoFitterOptions)
//		if !localVarOptionalAutoFitterOptionsok {
//				return localVarReturnValue, nil, reportError("autoFitterOptions should be AutoFitterOptions")
//		}
//		localVarPostBody = &localVarOptionalAutoFitterOptions
//	}
	if localVarOptionals != nil &&  &localVarOptionals.AutoFitterOptions != nil {
		
		localVarPostBody = &localVarOptionals.AutoFitterOptions
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService
 * @param name
 * @param sheetName
 * @param rowIndex
 * @param firstColumn
 * @param lastColumn
 * @param optional nil or *CellsWorksheetsPostAutofitWorksheetRowOpts - Optional Parameters:
     * @param "AutoFitterOptions" (optional.Interface of AutoFitterOptions) - 
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsWorksheetsPostAutofitWorksheetRowOpts struct { 
	Name string
	SheetName string
	RowIndex int32
	FirstColumn int32
	LastColumn int32
	AutoFitterOptions *AutoFitterOptions
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetsPostAutofitWorksheetRow(    localVarOptionals *CellsWorksheetsPostAutofitWorksheetRowOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/autofitrow"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("rowIndex", parameterToString(localVarOptionals.RowIndex, ""))
	localVarQueryParams.Add("firstColumn", parameterToString(localVarOptionals.FirstColumn, ""))
	localVarQueryParams.Add("lastColumn", parameterToString(localVarOptionals.LastColumn, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.AutoFitterOptions.IsSet() {
//		
//		localVarOptionalAutoFitterOptions, localVarOptionalAutoFitterOptionsok := localVarOptionals.AutoFitterOptions.Value().(AutoFitterOptions)
//		if !localVarOptionalAutoFitterOptionsok {
//				return localVarReturnValue, nil, reportError("autoFitterOptions should be AutoFitterOptions")
//		}
//		localVarPostBody = &localVarOptionalAutoFitterOptions
//	}
	if localVarOptionals != nil &&  &localVarOptionals.AutoFitterOptions != nil {
		
		localVarPostBody = &localVarOptionals.AutoFitterOptions
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Autofit worksheet rows.
 * @param name Document name.
 * @param sheetName The worksheet name.
 * @param optional nil or *CellsWorksheetsPostAutofitWorksheetRowsOpts - Optional Parameters:
     * @param "AutoFitterOptions" (optional.Interface of AutoFitterOptions) -  Auto Fitter Options.
     * @param "StartRow" (optional.Int32) -  Start row.
     * @param "EndRow" (optional.Int32) -  End row.
     * @param "OnlyAuto" (optional.Bool) -  Only auto.
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsWorksheetsPostAutofitWorksheetRowsOpts struct { 
	Name string
	SheetName string
	AutoFitterOptions *AutoFitterOptions
	StartRow int32
	EndRow int32
	OnlyAuto bool
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetsPostAutofitWorksheetRows(    localVarOptionals *CellsWorksheetsPostAutofitWorksheetRowsOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/autofitrows"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("startRow", parameterToString(localVarOptionals.StartRow.Value(), ""))
		localVarQueryParams.Add("startRow", parameterToString(localVarOptionals.StartRow, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("endRow", parameterToString(localVarOptionals.EndRow.Value(), ""))
		localVarQueryParams.Add("endRow", parameterToString(localVarOptionals.EndRow, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("onlyAuto", parameterToString(localVarOptionals.OnlyAuto.Value(), ""))
		localVarQueryParams.Add("onlyAuto", parameterToString(localVarOptionals.OnlyAuto, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.AutoFitterOptions.IsSet() {
//		
//		localVarOptionalAutoFitterOptions, localVarOptionalAutoFitterOptionsok := localVarOptionals.AutoFitterOptions.Value().(AutoFitterOptions)
//		if !localVarOptionalAutoFitterOptionsok {
//				return localVarReturnValue, nil, reportError("autoFitterOptions should be AutoFitterOptions")
//		}
//		localVarPostBody = &localVarOptionalAutoFitterOptions
//	}
	if localVarOptionals != nil &&  &localVarOptionals.AutoFitterOptions != nil {
		
		localVarPostBody = &localVarOptionals.AutoFitterOptions
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService
 * @param name
 * @param sheetName
 * @param sourceSheet
 * @param optional nil or *CellsWorksheetsPostCopyWorksheetOpts - Optional Parameters:
     * @param "Options" (optional.Interface of CopyOptions) - 
     * @param "SourceWorkbook" (optional.String) - 
     * @param "SourceFolder" (optional.String) - 
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsWorksheetsPostCopyWorksheetOpts struct { 
	Name string
	SheetName string
	SourceSheet string
	Options *CopyOptions
	SourceWorkbook string
	SourceFolder string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetsPostCopyWorksheet(    localVarOptionals *CellsWorksheetsPostCopyWorksheetOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/copy"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("sourceSheet", parameterToString(localVarOptionals.SourceSheet, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("sourceWorkbook", parameterToString(localVarOptionals.SourceWorkbook.Value(), ""))
		localVarQueryParams.Add("sourceWorkbook", parameterToString(localVarOptionals.SourceWorkbook, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("sourceFolder", parameterToString(localVarOptionals.SourceFolder.Value(), ""))
		localVarQueryParams.Add("sourceFolder", parameterToString(localVarOptionals.SourceFolder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Options.IsSet() {
//		
//		localVarOptionalOptions, localVarOptionalOptionsok := localVarOptionals.Options.Value().(CopyOptions)
//		if !localVarOptionalOptionsok {
//				return localVarReturnValue, nil, reportError("options should be CopyOptions")
//		}
//		localVarPostBody = &localVarOptionalOptions
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Options != nil {
		
		localVarPostBody = &localVarOptionals.Options
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Move worksheet.
 * @param name Document name.
 * @param sheetName The worksheet name.
 * @param optional nil or *CellsWorksheetsPostMoveWorksheetOpts - Optional Parameters:
     * @param "Moving" (optional.Interface of WorksheetMovingRequest) -  with moving parameters.
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return WorksheetsResponse
*/


type CellsWorksheetsPostMoveWorksheetOpts struct { 
	Name string
	SheetName string
	Moving *WorksheetMovingRequest
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetsPostMoveWorksheet(    localVarOptionals *CellsWorksheetsPostMoveWorksheetOpts) (WorksheetsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue WorksheetsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/position"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Moving.IsSet() {
//		
//		localVarOptionalMoving, localVarOptionalMovingok := localVarOptionals.Moving.Value().(WorksheetMovingRequest)
//		if !localVarOptionalMovingok {
//				return localVarReturnValue, nil, reportError("moving should be WorksheetMovingRequest")
//		}
//		localVarPostBody = &localVarOptionalMoving
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Moving != nil {
		
		localVarPostBody = &localVarOptionals.Moving
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Rename worksheet
 * @param name
 * @param sheetName
 * @param newname
 * @param optional nil or *CellsWorksheetsPostRenameWorksheetOpts - Optional Parameters:
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsWorksheetsPostRenameWorksheetOpts struct { 
	Name string
	SheetName string
	Newname string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetsPostRenameWorksheet(    localVarOptionals *CellsWorksheetsPostRenameWorksheetOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/rename"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("newname", parameterToString(localVarOptionals.Newname, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Update worksheet property
 * @param name
 * @param sheetName
 * @param optional nil or *CellsWorksheetsPostUpdateWorksheetPropertyOpts - Optional Parameters:
     * @param "Sheet" (optional.Interface of Worksheet) - 
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return WorksheetResponse
*/


type CellsWorksheetsPostUpdateWorksheetPropertyOpts struct { 
	Name string
	SheetName string
	Sheet *Worksheet
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetsPostUpdateWorksheetProperty(    localVarOptionals *CellsWorksheetsPostUpdateWorksheetPropertyOpts) (WorksheetResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue WorksheetResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Sheet.IsSet() {
//		
//		localVarOptionalSheet, localVarOptionalSheetok := localVarOptionals.Sheet.Value().(Worksheet)
//		if !localVarOptionalSheetok {
//				return localVarReturnValue, nil, reportError("sheet should be Worksheet")
//		}
//		localVarPostBody = &localVarOptionalSheet
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Sheet != nil {
		
		localVarPostBody = &localVarOptionals.Sheet
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService
 * @param name
 * @param sheetName
 * @param value
 * @param optional nil or *CellsWorksheetsPostUpdateWorksheetZoomOpts - Optional Parameters:
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsWorksheetsPostUpdateWorksheetZoomOpts struct { 
	Name string
	SheetName string
	Value int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetsPostUpdateWorksheetZoom(    localVarOptionals *CellsWorksheetsPostUpdateWorksheetZoomOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/zoom"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("value", parameterToString(localVarOptionals.Value, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Update worksheet&#39;s cell comment.
 * @param name The document name.
 * @param sheetName The worksheet name.
 * @param cellName The cell name
 * @param optional nil or *CellsWorksheetsPostWorksheetCommentOpts - Optional Parameters:
     * @param "Comment" (optional.Interface of Comment) -  Comment object
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsWorksheetsPostWorksheetCommentOpts struct { 
	Name string
	SheetName string
	CellName string
	Comment *Comment
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetsPostWorksheetComment(    localVarOptionals *CellsWorksheetsPostWorksheetCommentOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/comments/{cellName}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cellName"+"}", fmt.Sprintf("%v", localVarOptionals.CellName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Comment.IsSet() {
//		
//		localVarOptionalComment, localVarOptionalCommentok := localVarOptionals.Comment.Value().(Comment)
//		if !localVarOptionalCommentok {
//				return localVarReturnValue, nil, reportError("comment should be Comment")
//		}
//		localVarPostBody = &localVarOptionalComment
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Comment != nil {
		
		localVarPostBody = &localVarOptionals.Comment
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Sort worksheet range.
 * @param name The workbook name.
 * @param sheetName The worksheet name.
 * @param cellArea The range to sort.
 * @param optional nil or *CellsWorksheetsPostWorksheetRangeSortOpts - Optional Parameters:
     * @param "DataSorter" (optional.Interface of DataSorter) -  with sorting settings.
     * @param "Folder" (optional.String) -  The workbook folder.
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsWorksheetsPostWorksheetRangeSortOpts struct { 
	Name string
	SheetName string
	CellArea string
	DataSorter *DataSorter
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetsPostWorksheetRangeSort(    localVarOptionals *CellsWorksheetsPostWorksheetRangeSortOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/sort"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("cellArea", parameterToString(localVarOptionals.CellArea, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.DataSorter.IsSet() {
//		
//		localVarOptionalDataSorter, localVarOptionalDataSorterok := localVarOptionals.DataSorter.Value().(DataSorter)
//		if !localVarOptionalDataSorterok {
//				return localVarReturnValue, nil, reportError("dataSorter should be DataSorter")
//		}
//		localVarPostBody = &localVarOptionalDataSorter
//	}
	if localVarOptionals != nil &&  &localVarOptionals.DataSorter != nil {
		
		localVarPostBody = &localVarOptionals.DataSorter
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Search text.
 * @param name Document name.
 * @param sheetName The worksheet name.
 * @param text Text to search.
 * @param optional nil or *CellsWorksheetsPostWorksheetTextSearchOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return TextItemsResponse
*/


type CellsWorksheetsPostWorksheetTextSearchOpts struct { 
	Name string
	SheetName string
	Text string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetsPostWorksheetTextSearch(    localVarOptionals *CellsWorksheetsPostWorksheetTextSearchOpts) (TextItemsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue TextItemsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/findText"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("text", parameterToString(localVarOptionals.Text, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Replace text.
 * @param name Document name.
 * @param sheetName Worksheet name.
 * @param oldValue The old text to replace.
 * @param newValue The new text to replace by.
 * @param optional nil or *CellsWorksheetsPostWorsheetTextReplaceOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  Document&#39;s folder.
     * @param "Storage" (optional.String) -  storage name.

@return WorksheetReplaceResponse
*/


type CellsWorksheetsPostWorsheetTextReplaceOpts struct { 
	Name string
	SheetName string
	OldValue string
	NewValue string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetsPostWorsheetTextReplace(    localVarOptionals *CellsWorksheetsPostWorsheetTextReplaceOpts) (WorksheetReplaceResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue WorksheetReplaceResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/replaceText"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("oldValue", parameterToString(localVarOptionals.OldValue, ""))
	localVarQueryParams.Add("newValue", parameterToString(localVarOptionals.NewValue, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Add new worksheet.
 * @param name Document name.
 * @param sheetName The new sheet name.
 * @param optional nil or *CellsWorksheetsPutAddNewWorksheetOpts - Optional Parameters:
     * @param "Position" (optional.Int32) -  The new sheet position.
     * @param "Sheettype" (optional.String) -  The new sheet type.
     * @param "Folder" (optional.String) -  Document folder.
     * @param "Storage" (optional.String) -  storage name.

@return WorksheetsResponse
*/


type CellsWorksheetsPutAddNewWorksheetOpts struct { 
	Name string
	SheetName string
	Position int32
	Sheettype string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetsPutAddNewWorksheet(    localVarOptionals *CellsWorksheetsPutAddNewWorksheetOpts) (WorksheetsResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue WorksheetsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("position", parameterToString(localVarOptionals.Position.Value(), ""))
		localVarQueryParams.Add("position", parameterToString(localVarOptionals.Position, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("sheettype", parameterToString(localVarOptionals.Sheettype.Value(), ""))
		localVarQueryParams.Add("sheettype", parameterToString(localVarOptionals.Sheettype, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Change worksheet visibility.
 * @param name Document name.
 * @param sheetName Worksheet name.
 * @param isVisible New worksheet visibility value.
 * @param optional nil or *CellsWorksheetsPutChangeVisibilityWorksheetOpts - Optional Parameters:
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return WorksheetResponse
*/


type CellsWorksheetsPutChangeVisibilityWorksheetOpts struct { 
	Name string
	SheetName string
	IsVisible bool
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetsPutChangeVisibilityWorksheet(    localVarOptionals *CellsWorksheetsPutChangeVisibilityWorksheetOpts) (WorksheetResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue WorksheetResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/visible"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("isVisible", parameterToString(localVarOptionals.IsVisible, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Protect worksheet.
 * @param name Document name.
 * @param sheetName The worksheet name.
 * @param optional nil or *CellsWorksheetsPutProtectWorksheetOpts - Optional Parameters:
     * @param "ProtectParameter" (optional.Interface of ProtectSheetParameter) -  with protection settings.
     * @param "Folder" (optional.String) -  Document folder.
     * @param "Storage" (optional.String) -  storage name.

@return WorksheetResponse
*/


type CellsWorksheetsPutProtectWorksheetOpts struct { 
	Name string
	SheetName string
	ProtectParameter *ProtectSheetParameter
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetsPutProtectWorksheet(    localVarOptionals *CellsWorksheetsPutProtectWorksheetOpts) (WorksheetResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue WorksheetResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/protection"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.ProtectParameter.IsSet() {
//		
//		localVarOptionalProtectParameter, localVarOptionalProtectParameterok := localVarOptionals.ProtectParameter.Value().(ProtectSheetParameter)
//		if !localVarOptionalProtectParameterok {
//				return localVarReturnValue, nil, reportError("protectParameter should be ProtectSheetParameter")
//		}
//		localVarPostBody = &localVarOptionalProtectParameter
//	}
	if localVarOptionals != nil &&  &localVarOptionals.ProtectParameter != nil {
		
		localVarPostBody = &localVarOptionals.ProtectParameter
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Set worksheet background image.
 * @param name
 * @param sheetName
 * @param png
 * @param optional nil or *CellsWorksheetsPutWorksheetBackgroundOpts - Optional Parameters:
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsWorksheetsPutWorksheetBackgroundOpts struct { 
	Name string
	SheetName string
	Png string
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetsPutWorksheetBackground(    localVarOptionals *CellsWorksheetsPutWorksheetBackgroundOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/background"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
	localVarPostBody = &localVarOptionals.Png
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Add worksheet&#39;s cell comment.
 * @param name The document name.
 * @param sheetName The worksheet name.
 * @param cellName The cell name
 * @param optional nil or *CellsWorksheetsPutWorksheetCommentOpts - Optional Parameters:
     * @param "Comment" (optional.Interface of Comment) -  Comment object
     * @param "Folder" (optional.String) -  The document folder.
     * @param "Storage" (optional.String) -  storage name.

@return CommentResponse
*/


type CellsWorksheetsPutWorksheetCommentOpts struct { 
	Name string
	SheetName string
	CellName string
	Comment *Comment
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetsPutWorksheetComment(    localVarOptionals *CellsWorksheetsPutWorksheetCommentOpts) (CommentResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CommentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/comments/{cellName}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cellName"+"}", fmt.Sprintf("%v", localVarOptionals.CellName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	// body params
//	if localVarOptionals != nil && localVarOptionals.Comment.IsSet() {
//		
//		localVarOptionalComment, localVarOptionalCommentok := localVarOptionals.Comment.Value().(Comment)
//		if !localVarOptionalCommentok {
//				return localVarReturnValue, nil, reportError("comment should be Comment")
//		}
//		localVarPostBody = &localVarOptionalComment
//	}
	if localVarOptionals != nil &&  &localVarOptionals.Comment != nil {
		
		localVarPostBody = &localVarOptionals.Comment
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Set freeze panes
 * @param name
 * @param sheetName
 * @param row
 * @param column
 * @param freezedRows
 * @param freezedColumns
 * @param optional nil or *CellsWorksheetsPutWorksheetFreezePanesOpts - Optional Parameters:
     * @param "Folder" (optional.String) - 
     * @param "Storage" (optional.String) -  storage name.

@return CellsCloudResponse
*/


type CellsWorksheetsPutWorksheetFreezePanesOpts struct { 
	Name string
	SheetName string
	Row int32
	Column int32
	FreezedRows int32
	FreezedColumns int32
	Folder string
	Storage string
}


func (a *CellsApiService) CellsWorksheetsPutWorksheetFreezePanes(    localVarOptionals *CellsWorksheetsPutWorksheetFreezePanesOpts) (CellsCloudResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue CellsCloudResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/{name}/worksheets/{sheetName}/freezepanes"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", localVarOptionals.Name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sheetName"+"}", fmt.Sprintf("%v", localVarOptionals.SheetName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("row", parameterToString(localVarOptionals.Row, ""))
	localVarQueryParams.Add("column", parameterToString(localVarOptionals.Column, ""))
	localVarQueryParams.Add("freezedRows", parameterToString(localVarOptionals.FreezedRows, ""))
	localVarQueryParams.Add("freezedColumns", parameterToString(localVarOptionals.FreezedColumns, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder.Value(), ""))
		localVarQueryParams.Add("folder", parameterToString(localVarOptionals.Folder, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage.Value(), ""))
		localVarQueryParams.Add("storage", parameterToString(localVarOptionals.Storage, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Copy file
 * @param srcPath Source file path e.g. &#39;/folder/file.ext&#39;
 * @param destPath Destination file path
 * @param optional nil or *CopyFileOpts - Optional Parameters:
     * @param "SrcStorageName" (optional.String) -  Source storage name
     * @param "DestStorageName" (optional.String) -  Destination storage name
     * @param "VersionId" (optional.String) -  File version ID to copy


*/


type CopyFileOpts struct { 
	SrcPath string
	DestPath string
	SrcStorageName string
	DestStorageName string
	VersionId string
}


func (a *CellsApiService) CopyFile(    localVarOptionals *CopyFileOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/storage/file/copy/{srcPath}"
	localVarPath = strings.Replace(localVarPath, "{"+"srcPath"+"}", fmt.Sprintf("%v", localVarOptionals.SrcPath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("destPath", parameterToString(localVarOptionals.DestPath, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("srcStorageName", parameterToString(localVarOptionals.SrcStorageName.Value(), ""))
		localVarQueryParams.Add("srcStorageName", parameterToString(localVarOptionals.SrcStorageName, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("destStorageName", parameterToString(localVarOptionals.DestStorageName.Value(), ""))
		localVarQueryParams.Add("destStorageName", parameterToString(localVarOptionals.DestStorageName, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("versionId", parameterToString(localVarOptionals.VersionId.Value(), ""))
		localVarQueryParams.Add("versionId", parameterToString(localVarOptionals.VersionId, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarHttpResponse, err
}

/* 
CellsApiService Copy folder
 * @param srcPath Source folder path e.g. &#39;/src&#39;
 * @param destPath Destination folder path e.g. &#39;/dst&#39;
 * @param optional nil or *CopyFolderOpts - Optional Parameters:
     * @param "SrcStorageName" (optional.String) -  Source storage name
     * @param "DestStorageName" (optional.String) -  Destination storage name


*/


type CopyFolderOpts struct { 
	SrcPath string
	DestPath string
	SrcStorageName string
	DestStorageName string
}


func (a *CellsApiService) CopyFolder(    localVarOptionals *CopyFolderOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/storage/folder/copy/{srcPath}"
	localVarPath = strings.Replace(localVarPath, "{"+"srcPath"+"}", fmt.Sprintf("%v", localVarOptionals.SrcPath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("destPath", parameterToString(localVarOptionals.DestPath, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("srcStorageName", parameterToString(localVarOptionals.SrcStorageName.Value(), ""))
		localVarQueryParams.Add("srcStorageName", parameterToString(localVarOptionals.SrcStorageName, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("destStorageName", parameterToString(localVarOptionals.DestStorageName.Value(), ""))
		localVarQueryParams.Add("destStorageName", parameterToString(localVarOptionals.DestStorageName, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarHttpResponse, err
}

/* 
CellsApiService Create the folder
 * @param path Folder path to create e.g. &#39;folder_1/folder_2/&#39;
 * @param optional nil or *CreateFolderOpts - Optional Parameters:
     * @param "StorageName" (optional.String) -  Storage name


*/


type CreateFolderOpts struct { 
	Path string
	StorageName string
}


func (a *CellsApiService) CreateFolder(    localVarOptionals *CreateFolderOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/storage/folder/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", localVarOptionals.Path), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("storageName", parameterToString(localVarOptionals.StorageName.Value(), ""))
		localVarQueryParams.Add("storageName", parameterToString(localVarOptionals.StorageName, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarHttpResponse, err
}

/* 
CellsApiService Delete file
 * @param path File path e.g. &#39;/folder/file.ext&#39;
 * @param optional nil or *DeleteFileOpts - Optional Parameters:
     * @param "StorageName" (optional.String) -  Storage name
     * @param "VersionId" (optional.String) -  File version ID to delete


*/


type DeleteFileOpts struct { 
	Path string
	StorageName string
	VersionId string
}


func (a *CellsApiService) DeleteFile(    localVarOptionals *DeleteFileOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/storage/file/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", localVarOptionals.Path), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("storageName", parameterToString(localVarOptionals.StorageName.Value(), ""))
		localVarQueryParams.Add("storageName", parameterToString(localVarOptionals.StorageName, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("versionId", parameterToString(localVarOptionals.VersionId.Value(), ""))
		localVarQueryParams.Add("versionId", parameterToString(localVarOptionals.VersionId, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarHttpResponse, err
}

/* 
CellsApiService Delete folder
 * @param path Folder path e.g. &#39;/folder&#39;
 * @param optional nil or *DeleteFolderOpts - Optional Parameters:
     * @param "StorageName" (optional.String) -  Storage name
     * @param "Recursive" (optional.Bool) -  Enable to delete folders, subfolders and files


*/


type DeleteFolderOpts struct { 
	Path string
	StorageName string
	Recursive bool
}


func (a *CellsApiService) DeleteFolder(    localVarOptionals *DeleteFolderOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/storage/folder/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", localVarOptionals.Path), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("storageName", parameterToString(localVarOptionals.StorageName.Value(), ""))
		localVarQueryParams.Add("storageName", parameterToString(localVarOptionals.StorageName, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("recursive", parameterToString(localVarOptionals.Recursive.Value(), ""))
		localVarQueryParams.Add("recursive", parameterToString(localVarOptionals.Recursive, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarHttpResponse, err
}

/* 
CellsApiService Download file
 * @param path File path e.g. &#39;/folder/file.ext&#39;
 * @param optional nil or *DownloadFileOpts - Optional Parameters:
     * @param "StorageName" (optional.String) -  Storage name
     * @param "VersionId" (optional.String) -  File version ID to download

@return *os.File
*/


type DownloadFileOpts struct { 
	Path string
	StorageName string
	VersionId string
}


func (a *CellsApiService) DownloadFile(    localVarOptionals *DownloadFileOpts) ([]byte, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/storage/file/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", localVarOptionals.Path), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("storageName", parameterToString(localVarOptionals.StorageName.Value(), ""))
		localVarQueryParams.Add("storageName", parameterToString(localVarOptionals.StorageName, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("versionId", parameterToString(localVarOptionals.VersionId.Value(), ""))
		localVarQueryParams.Add("versionId", parameterToString(localVarOptionals.VersionId, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"multipart/form-data"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	localVarReturnValue, err = ioutil.ReadAll(localVarHttpResponse.Body)
    if err != nil {
        return localVarReturnValue, localVarHttpResponse, err
    }

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get disc usage
 * @param optional nil or *GetDiscUsageOpts - Optional Parameters:
     * @param "StorageName" (optional.String) -  Storage name

@return DiscUsage
*/


type GetDiscUsageOpts struct { 
	StorageName string
}


func (a *CellsApiService) GetDiscUsage(    localVarOptionals *GetDiscUsageOpts) (DiscUsage, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue DiscUsage
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/storage/disc"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("storageName", parameterToString(localVarOptionals.StorageName.Value(), ""))
		localVarQueryParams.Add("storageName", parameterToString(localVarOptionals.StorageName, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get file versions
 * @param path File path e.g. &#39;/file.ext&#39;
 * @param optional nil or *GetFileVersionsOpts - Optional Parameters:
     * @param "StorageName" (optional.String) -  Storage name

@return FileVersions
*/


type GetFileVersionsOpts struct { 
	Path string
	StorageName string
}


func (a *CellsApiService) GetFileVersions(    localVarOptionals *GetFileVersionsOpts) (FileVersions, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue FileVersions
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/storage/version/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", localVarOptionals.Path), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("storageName", parameterToString(localVarOptionals.StorageName.Value(), ""))
		localVarQueryParams.Add("storageName", parameterToString(localVarOptionals.StorageName, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Get all files and folders within a folder
 * @param path Folder path e.g. &#39;/folder&#39;
 * @param optional nil or *GetFilesListOpts - Optional Parameters:
     * @param "StorageName" (optional.String) -  Storage name

@return FilesList
*/


type GetFilesListOpts struct { 
	Path string
	StorageName string
}


func (a *CellsApiService) GetFilesList(    localVarOptionals *GetFilesListOpts) (FilesList, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue FilesList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/storage/folder/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", localVarOptionals.Path), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("storageName", parameterToString(localVarOptionals.StorageName.Value(), ""))
		localVarQueryParams.Add("storageName", parameterToString(localVarOptionals.StorageName, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Move file
 * @param srcPath Source file path e.g. &#39;/src.ext&#39;
 * @param destPath Destination file path e.g. &#39;/dest.ext&#39;
 * @param optional nil or *MoveFileOpts - Optional Parameters:
     * @param "SrcStorageName" (optional.String) -  Source storage name
     * @param "DestStorageName" (optional.String) -  Destination storage name
     * @param "VersionId" (optional.String) -  File version ID to move


*/


type MoveFileOpts struct { 
	SrcPath string
	DestPath string
	SrcStorageName string
	DestStorageName string
	VersionId string
}


func (a *CellsApiService) MoveFile(    localVarOptionals *MoveFileOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/storage/file/move/{srcPath}"
	localVarPath = strings.Replace(localVarPath, "{"+"srcPath"+"}", fmt.Sprintf("%v", localVarOptionals.SrcPath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("destPath", parameterToString(localVarOptionals.DestPath, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("srcStorageName", parameterToString(localVarOptionals.SrcStorageName.Value(), ""))
		localVarQueryParams.Add("srcStorageName", parameterToString(localVarOptionals.SrcStorageName, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("destStorageName", parameterToString(localVarOptionals.DestStorageName.Value(), ""))
		localVarQueryParams.Add("destStorageName", parameterToString(localVarOptionals.DestStorageName, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("versionId", parameterToString(localVarOptionals.VersionId.Value(), ""))
		localVarQueryParams.Add("versionId", parameterToString(localVarOptionals.VersionId, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarHttpResponse, err
}

/* 
CellsApiService Move folder
 * @param srcPath Folder path to move e.g. &#39;/folder&#39;
 * @param destPath Destination folder path to move to e.g &#39;/dst&#39;
 * @param optional nil or *MoveFolderOpts - Optional Parameters:
     * @param "SrcStorageName" (optional.String) -  Source storage name
     * @param "DestStorageName" (optional.String) -  Destination storage name


*/


type MoveFolderOpts struct { 
	SrcPath string
	DestPath string
	SrcStorageName string
	DestStorageName string
}


func (a *CellsApiService) MoveFolder(    localVarOptionals *MoveFolderOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/storage/folder/move/{srcPath}"
	localVarPath = strings.Replace(localVarPath, "{"+"srcPath"+"}", fmt.Sprintf("%v", localVarOptionals.SrcPath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("destPath", parameterToString(localVarOptionals.DestPath, ""))
	if localVarOptionals != nil {
//		localVarQueryParams.Add("srcStorageName", parameterToString(localVarOptionals.SrcStorageName.Value(), ""))
		localVarQueryParams.Add("srcStorageName", parameterToString(localVarOptionals.SrcStorageName, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("destStorageName", parameterToString(localVarOptionals.DestStorageName.Value(), ""))
		localVarQueryParams.Add("destStorageName", parameterToString(localVarOptionals.DestStorageName, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarHttpResponse, err
}

/* 
CellsApiService Get Access token
 * @param grantType Grant Type
 * @param clientId App SID
 * @param clientSecret App Key

@return AccessTokenResponse
*/


type OAuthPostOpts struct { 
	GrantType string
	ClientId string
	ClientSecret string
}


func (a *CellsApiService) OAuthPost(    localVarOptionals *OAuthPostOpts) (AccessTokenResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue AccessTokenResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/connect/token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

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
	localVarFormParams.Add("grant_type", parameterToString(localVarOptionals.GrantType, ""))
	localVarFormParams.Add("client_id", parameterToString(localVarOptionals.ClientId, ""))
	localVarFormParams.Add("client_secret", parameterToString(localVarOptionals.ClientSecret, ""))
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Check if file or folder exists
 * @param path File or folder path e.g. &#39;/file.ext&#39; or &#39;/folder&#39;
 * @param optional nil or *ObjectExistsOpts - Optional Parameters:
     * @param "StorageName" (optional.String) -  Storage name
     * @param "VersionId" (optional.String) -  File version ID

@return ObjectExist
*/


type ObjectExistsOpts struct { 
	Path string
	StorageName string
	VersionId string
}


func (a *CellsApiService) ObjectExists(    localVarOptionals *ObjectExistsOpts) (ObjectExist, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ObjectExist
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/storage/exist/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", localVarOptionals.Path), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("storageName", parameterToString(localVarOptionals.StorageName.Value(), ""))
		localVarQueryParams.Add("storageName", parameterToString(localVarOptionals.StorageName, ""))
	}
	if localVarOptionals != nil {
//		localVarQueryParams.Add("versionId", parameterToString(localVarOptionals.VersionId.Value(), ""))
		localVarQueryParams.Add("versionId", parameterToString(localVarOptionals.VersionId, ""))
	}
	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Check if storage exists
 * @param storageName Storage name

@return StorageExist
*/


type StorageExistsOpts struct { 
	StorageName string
}


func (a *CellsApiService) StorageExists(    localVarOptionals *StorageExistsOpts) (StorageExist, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue StorageExist
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/storage/{storageName}/exist"
	localVarPath = strings.Replace(localVarPath, "{"+"storageName"+"}", fmt.Sprintf("%v", localVarOptionals.StorageName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
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
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/* 
CellsApiService Upload file
 * @param path Path where to upload including filename and extension e.g. /file.ext or /Folder 1/file.ext             If the content is multipart and path does not contains the file name it tries to get them from filename parameter             from Content-Disposition header.             
 * @param file File to upload
 * @param optional nil or *UploadFileOpts - Optional Parameters:
     * @param "StorageName" (optional.String) -  Storage name

@return FilesUploadResult
*/


type UploadFileOpts struct { 
	Path string

	StorageName string
}


func (a *CellsApiService) UploadFile(   file *os.File ,    localVarOptionals *UploadFileOpts) (FilesUploadResult, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue FilesUploadResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cells/storage/file/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", localVarOptionals.Path), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
//		localVarQueryParams.Add("storageName", parameterToString(localVarOptionals.StorageName.Value(), ""))
		localVarQueryParams.Add("storageName", parameterToString(localVarOptionals.StorageName, ""))
	}
	// to determine the Content-Type header
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
	var localVarFile (*os.File) =  file
	if localVarFile != nil {
		fbs, _ := ioutil.ReadAll(localVarFile)
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
	}
	r, err := a.client.prepareRequest( localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarReturnValue, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarReturnValue, localVarHttpResponse, err
}
