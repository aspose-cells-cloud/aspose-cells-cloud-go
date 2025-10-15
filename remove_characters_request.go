/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="remove_characters_request.go">
*   Copyright (c) 2025 Aspose.Cells Cloud
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

type RemoveCharactersRequest struct {
	TheFirstNCharacters int64 `json:"the_first_n_characters,omitempty" xml:"the_first_n_characters"`
	TheLastNCharacters int64 `json:"the_last_n_characters,omitempty" xml:"the_last_n_characters"`
	AllCharactersBeforeText string `json:"all_characters_before_text,omitempty" xml:"all_characters_before_text"`
	AllCharactersAfterText string `json:"all_characters_after_text,omitempty" xml:"all_characters_after_text"`
	RemoveTextMethod string `json:"remove_text_method,omitempty" xml:"remove_text_method"`
	CharacterSets string `json:"character_sets,omitempty" xml:"character_sets"`
	RemoveCustomValue string `json:"remove_custom_value,omitempty" xml:"remove_custom_value"`
	Worksheet string `json:"worksheet,omitempty" xml:"worksheet"`
	Range_ string `json:"range,omitempty" xml:"range"`
	OutPath string `json:"out_path,omitempty" xml:"out_path"`
	OutStorageName string `json:"out_storage_name,omitempty" xml:"out_storage_name"`
	Region string `json:"region,omitempty" xml:"region"`
	Password string `json:"password,omitempty" xml:"password"`
	
	Spreadsheet string  `json:"spreadsheet,omitempty" xml:"spreadsheet"`
	 

	ExtendQueryParameterMap	map[string]string `json:"ExtendQueryParameterMap,omitempty" xml:"ExtendQueryParameterMap"`	
}

func (data *RemoveCharactersRequest) CreateRequestData( client *APIClient) (localVarRequest *http.Request, err error) {
	var (
		localVarHttpMethod  = strings.ToUpper("PUT")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
	)

	// create path and map variables
	localVarPath := client.cfg.BasePath + "/v4.0/cells/content/remove/characters"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

    // query params : theFirstNCharacters
    if &data.TheFirstNCharacters != nil {
        localVarQueryParams.Add("theFirstNCharacters", parameterToString(data.TheFirstNCharacters, ""))
    }
    // query params : theLastNCharacters
    if &data.TheLastNCharacters != nil {
        localVarQueryParams.Add("theLastNCharacters", parameterToString(data.TheLastNCharacters, ""))
    }
    // query params : allCharactersBeforeText
    if data.AllCharactersBeforeText != "" {
        localVarQueryParams.Add("allCharactersBeforeText", parameterToString(data.AllCharactersBeforeText, ""))
    }
    // query params : allCharactersAfterText
    if data.AllCharactersAfterText != "" {
        localVarQueryParams.Add("allCharactersAfterText", parameterToString(data.AllCharactersAfterText, ""))
    }
    // query params : removeTextMethod
    if data.RemoveTextMethod != "" {
        localVarQueryParams.Add("removeTextMethod", parameterToString(data.RemoveTextMethod, ""))
    }
    // query params : characterSets
    if data.CharacterSets != "" {
        localVarQueryParams.Add("characterSets", parameterToString(data.CharacterSets, ""))
    }
    // query params : removeCustomValue
    if data.RemoveCustomValue != "" {
        localVarQueryParams.Add("removeCustomValue", parameterToString(data.RemoveCustomValue, ""))
    }
    // query params : worksheet
    if data.Worksheet != "" {
        localVarQueryParams.Add("worksheet", parameterToString(data.Worksheet, ""))
    }
    // query params : range
    if data.Range_ != "" {
        localVarQueryParams.Add("range", parameterToString(data.Range_, ""))
    }
    // query params : outPath
    if data.OutPath != "" {
        localVarQueryParams.Add("outPath", parameterToString(data.OutPath, ""))
    }
    // query params : outStorageName
    if data.OutStorageName != "" {
        localVarQueryParams.Add("outStorageName", parameterToString(data.OutStorageName, ""))
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
			
		if strings.TrimSpace(data.Spreadsheet) != "" { localVarFormParams["@"+ filepath.Base(data.Spreadsheet)] = []string {data.Spreadsheet}} 
	r, err := client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	return r,err
}