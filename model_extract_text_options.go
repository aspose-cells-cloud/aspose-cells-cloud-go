/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="ExtractTextOptions.go">
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

type ExtractTextOptions struct {
     
        Name string `json:"Name,omitempty" xml:"Name"`
 
    DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
    FileInfo *FileInfo `json:"FileInfo,omitempty" xml:"FileInfo"`
    Worksheet string `json:"Worksheet,omitempty" xml:"Worksheet"`
    Range_ string `json:"Range,omitempty" xml:"Range"`
    ExtractTextType string `json:"ExtractTextType,omitempty" xml:"ExtractTextType"`
    BeforeText string `json:"BeforeText,omitempty" xml:"BeforeText"`
    AfterText string `json:"AfterText,omitempty" xml:"AfterText"`
    BeforePosition int64 `json:"BeforePosition,omitempty" xml:"BeforePosition"`
    AfterPosition int64 `json:"AfterPosition,omitempty" xml:"AfterPosition"`
    OutPositionRange string `json:"OutPositionRange,omitempty" xml:"OutPositionRange"`
}
