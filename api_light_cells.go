/*
 *  Copyright (c) 2022 Aspose.Cells Cloud
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
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

/* Create Instance of CellsApiService
@param appSid string Application SID
@param appKey string Application Key
@param basePath string Base service path. Set "" for default
@return *CellsApiService */
func NewLightCellsApiService(appSid string, appKey string, opts ...string) *LightCellsApiService {
	var basePath = ""
	var version = ""
	for i, v := range opts {
		switch i {
		case 0:
			basePath = v
		case 1:
			version = v
		}
	}
	config := NewConfiguration(appSid, appKey, basePath, version)
	client := NewAPIClient(config)
	return client.LightCellsApi
}

type LightCellsApiService service

/*
LightCellsApiService
 * @param file File to upload
 * @param optional nil or *DeleteMetadataOpts - Optional Parameters:
     * @param "Type_" (optional.String) -

@return FilesResult
*/

type DeleteMetadataOpts struct {
	Type_ string
}

// var countryCapitalMap map[string]string /*创建集合 */
//countryCapitalMap = make(map[string]string)
func (a *LightCellsApiService) DeleteMetadata(file map[string]string, localVarOptionals *DeleteMetadataOpts) (FilesResult, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue FilesResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/" + a.client.cfg.Version + "/cells/metadata/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
		localVarQueryParams.Add("type", parameterToString(localVarOptionals.Type_, ""))
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
	// localVarFormParams = file
	for name, path := range file {
		localVarFormParams["@"+name] = []string{path}
	}

	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}
	return localVarReturnValue, localVarHttpResponse, err
}

/*
LightCellsApiService
 * @param file File to upload
 * @param optional nil or *GetMetadataOpts - Optional Parameters:
     * @param "Type_" (optional.String) -

@return []CellsDocumentProperty
*/

type GetMetadataOpts struct {
	Type_ string
}

func (a *LightCellsApiService) GetMetadata(file map[string]string, localVarOptionals *GetMetadataOpts) ([]CellsDocumentProperty, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []CellsDocumentProperty
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/" + a.client.cfg.Version + "/cells/metadata/get"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
		localVarQueryParams.Add("type", parameterToString(localVarOptionals.Type_, ""))
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
	// localVarFormParams = file
	for name, path := range file {
		localVarFormParams["@"+name] = []string{path}
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/*
LightCellsApiService
 * @param file File to upload
 * @param datasource
 * @param optional nil or *PostAssembleOpts - Optional Parameters:
     * @param "Format" (optional.String) -

@return FilesResult
*/

type PostAssembleOpts struct {
	Datasource string
	Format     string
}

func (a *LightCellsApiService) PostAssemble(file map[string]string, localVarOptionals *PostAssembleOpts) (FilesResult, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue FilesResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/" + a.client.cfg.Version + "/cells/assemble"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("datasource", parameterToString(localVarOptionals.Datasource, ""))
	if localVarOptionals != nil {
		localVarQueryParams.Add("format", parameterToString(localVarOptionals.Format, ""))
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
	// localVarFormParams = file
	for name, path := range file {
		localVarFormParams["@"+name] = []string{path}
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/*
LightCellsApiService
 * @param file File to upload
 * @param objecttype

@return FilesResult
*/

type PostClearObjectsOpts struct {
	Objecttype string
}

func (a *LightCellsApiService) PostClearObjects(file map[string]string, localVarOptionals *PostClearObjectsOpts) (FilesResult, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue FilesResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/" + a.client.cfg.Version + "/cells/clearobjects"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("objecttype", parameterToString(localVarOptionals.Objecttype, ""))
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
	// localVarFormParams = file
	for name, path := range file {
		localVarFormParams["@"+name] = []string{path}
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/*
LightCellsApiService
 * @param file File to upload
 * @param objectType
 * @param format

@return FilesResult
*/

type PostExportOpts struct {
	ObjectType              string
	Format                  string
	ExtendedQueryParameters map[string]string
}

func (a *LightCellsApiService) PostExport(file map[string]string, localVarOptionals *PostExportOpts) (FilesResult, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue FilesResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/" + a.client.cfg.Version + "/cells/export"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("objectType", parameterToString(localVarOptionals.ObjectType, ""))
	localVarQueryParams.Add("format", parameterToString(localVarOptionals.Format, ""))

	if localVarOptionals.ExtendedQueryParameters != nil {
		for key, value := range localVarOptionals.ExtendedQueryParameters {
			localVarQueryParams.Add(key, parameterToString(value, ""))
		}
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
	// localVarFormParams = file
	for name, path := range file {
		localVarFormParams["@"+name] = []string{path}
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/*
LightCellsApiService
 * @param file File to upload
 * @param documentProperties Cells document property.

@return FilesResult
*/

type PostImportOpts struct {
	ImportOption interface{}
}

func (a *LightCellsApiService) PostImport(file map[string]string, localVarOptionals *PostImportOpts) (FilesResult, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue FilesResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/" + a.client.cfg.Version + "/cells/import"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"multipart/form-data"}

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

	// localVarFormParams = file
	for name, path := range file {
		localVarFormParams["@"+name] = []string{path}
	}
	// body params
	localVarPostBody = &localVarOptionals.ImportOption
	if localVarPostBody != nil {
		b, _ := json.Marshal(localVarPostBody)
		localVarFormParams["ImportOption"] = []string{string(b)}
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/*
LightCellsApiService
 * @param file File to upload
 * @param optional nil or *PostMergeOpts - Optional Parameters:
     * @param "Format" (optional.String) -
     * @param "MergeToOneSheet" (optional.Bool) -

@return FileInfo
*/

type PostMergeOpts struct {
	Format          string
	MergeToOneSheet bool
}

func (a *LightCellsApiService) PostMerge(file map[string]string, localVarOptionals *PostMergeOpts) (FileInfo, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue FileInfo
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/" + a.client.cfg.Version + "/cells/merge"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil {
		localVarQueryParams.Add("format", parameterToString(localVarOptionals.Format, ""))
	}
	if localVarOptionals != nil {
		localVarQueryParams.Add("mergeToOneSheet", parameterToString(localVarOptionals.MergeToOneSheet, ""))
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
	// localVarFormParams = file
	for name, path := range file {
		localVarFormParams["@"+name] = []string{path}
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/*
LightCellsApiService
 * @param file File to upload
 * @param documentProperties Cells document property.

@return FilesResult
*/

type PostMetadataOpts struct {
	DocumentProperties []CellsDocumentProperty
}

func (a *LightCellsApiService) PostMetadata(file map[string]string, localVarOptionals *PostMetadataOpts) (FilesResult, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue FilesResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/" + a.client.cfg.Version + "/cells/metadata/update"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"multipart/form-data"}

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

	// localVarFormParams = file
	for name, path := range file {
		localVarFormParams["@"+name] = []string{path}
	}
	// body params
	localVarPostBody = &localVarOptionals.DocumentProperties
	if localVarPostBody != nil {
		b, _ := json.Marshal(localVarPostBody)
		localVarFormParams["DocumentProperties"] = []string{string(b)}
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/*
LightCellsApiService
 * @param file File to upload
 * @param password

@return FilesResult
*/

type PostProtectOpts struct {
	Password string
}

func (a *LightCellsApiService) PostProtect(file map[string]string, localVarOptionals *PostProtectOpts) (FilesResult, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue FilesResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/" + a.client.cfg.Version + "/cells/protect"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("password", parameterToString(localVarOptionals.Password, ""))
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
	// localVarFormParams = file
	for name, path := range file {
		localVarFormParams["@"+name] = []string{path}
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/*
LightCellsApiService
 * @param file File to upload
 * @param text
 * @param optional nil or *PostSearchOpts - Optional Parameters:
     * @param "Password" (optional.String) -
     * @param "Sheetname" (optional.String) -

@return []TextItem
*/

type PostSearchOpts struct {
	Text      string
	Password  string
	Sheetname string
}

func (a *LightCellsApiService) PostSearch(file map[string]string, localVarOptionals *PostSearchOpts) ([]TextItem, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []TextItem
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/" + a.client.cfg.Version + "/cells/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("text", parameterToString(localVarOptionals.Text, ""))
	if localVarOptionals != nil {
		localVarQueryParams.Add("password", parameterToString(localVarOptionals.Password, ""))
	}
	if localVarOptionals != nil {
		localVarQueryParams.Add("sheetname", parameterToString(localVarOptionals.Sheetname, ""))
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
	// localVarFormParams = file
	for name, path := range file {
		localVarFormParams["@"+name] = []string{path}
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/*
LightCellsApiService
 * @param file File to upload
 * @param format
 * @param optional nil or *PostSplitOpts - Optional Parameters:
     * @param "Password" (optional.String) -
     * @param "From" (optional.Int32) -
     * @param "To" (optional.Int32) -
@return FilesResult
*/

type PostSplitOpts struct {
	Format   string
	Password string
	From     int32
	To       int32
}

func (a *LightCellsApiService) PostSplit(file map[string]string, localVarOptionals *PostSplitOpts) (FilesResult, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue FilesResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/" + a.client.cfg.Version + "/cells/split"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("format", parameterToString(localVarOptionals.Format, ""))
	if localVarOptionals != nil {
		localVarQueryParams.Add("password", parameterToString(localVarOptionals.Password, ""))
	}
	if localVarOptionals != nil {
		localVarQueryParams.Add("from", parameterToString(localVarOptionals.From, ""))
	}
	if localVarOptionals != nil {
		localVarQueryParams.Add("to", parameterToString(localVarOptionals.To, ""))
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
	// localVarFormParams = file
	for name, path := range file {
		localVarFormParams["@"+name] = []string{path}
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/*
LightCellsApiService
 * @param file File to upload
 * @param password

@return FilesResult
*/

type PostUnlockOpts struct {
	Password string
}

func (a *LightCellsApiService) PostUnlock(file map[string]string, localVarOptionals *PostUnlockOpts) (FilesResult, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue FilesResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/" + a.client.cfg.Version + "/cells/unlock"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("password", parameterToString(localVarOptionals.Password, ""))
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
	// localVarFormParams = file
	for name, path := range file {
		localVarFormParams["@"+name] = []string{path}
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/*
LightCellsApiService
 * @param file File to upload
 * @param text
 * @param color

@return FilesResult
*/

type PostWatermarkOpts struct {
	Text  string
	Color string
}

func (a *LightCellsApiService) PostWatermark(file map[string]string, localVarOptionals *PostWatermarkOpts) (FilesResult, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue FilesResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/" + a.client.cfg.Version + "/cells/watermark"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("text", parameterToString(localVarOptionals.Text, ""))
	localVarQueryParams.Add("color", parameterToString(localVarOptionals.Color, ""))
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
	// localVarFormParams = file
	for name, path := range file {
		localVarFormParams["@"+name] = []string{path}
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/*
LightCellsApiService
 * @param file File to upload
 * @param text
 * @param color

@return FilesResult
*/

type PostCompressOpts struct {
	CompressLevel int64
}

func (a *LightCellsApiService) PostCompress(file map[string]string, localVarOptionals *PostCompressOpts) (FilesResult, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue FilesResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/" + a.client.cfg.Version + "/cells/compress"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("CompressLevel", parameterToString(localVarOptionals.CompressLevel, ""))
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
	// localVarFormParams = file
	for name, path := range file {
		localVarFormParams["@"+name] = []string{path}
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/*
LightCellsApiService
 * @param file File to upload
 * @param text
 * @param color

@return FilesResult
*/

type PostReplaceOpts struct {
	Text      string
	NewText   string
	Password  string
	SheetName string
}

func (a *LightCellsApiService) PostReplace(file map[string]string, localVarOptionals *PostReplaceOpts) (FilesResult, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue FilesResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/" + a.client.cfg.Version + "/cells/replace"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("text", parameterToString(localVarOptionals.Text, ""))
	localVarQueryParams.Add("newtext", parameterToString(localVarOptionals.NewText, ""))
	localVarQueryParams.Add("password", parameterToString(localVarOptionals.Password, ""))
	localVarQueryParams.Add("sheetname", parameterToString(localVarOptionals.SheetName, ""))
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
	// localVarFormParams = file
	for name, path := range file {
		localVarFormParams["@"+name] = []string{path}
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/*
LightCellsApiService
 * @param file File to upload
 * @param rotateType
 * @param color

@return FilesResult
*/

type PostReverseOpts struct {
	RotateType string
	Format     string
}

func (a *LightCellsApiService) PostReverse(file map[string]string, localVarOptionals *PostReverseOpts) (FilesResult, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue FilesResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/" + a.client.cfg.Version + "/cells/reverse"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("rotateType", parameterToString(localVarOptionals.RotateType, ""))
	localVarQueryParams.Add("format", parameterToString(localVarOptionals.Format, ""))
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
	// localVarFormParams = file
	for name, path := range file {
		localVarFormParams["@"+name] = []string{path}
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&localVarReturnValue); err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	return localVarReturnValue, localVarHttpResponse, err
}
