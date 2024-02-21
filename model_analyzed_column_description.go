/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="AnalyzedColumnDescription.go">
*   Copyright (c) 2024 Aspose.Cells Cloud
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

type AnalyzedColumnDescription struct {
 
    Index int64 `json:"Index,omitempty" xml:"Index"`
    ColumnIndex int64 `json:"ColumnIndex,omitempty" xml:"ColumnIndex"`
    Title string `json:"Title,omitempty" xml:"Title"`
    RepetitionRate float64 `json:"RepetitionRate,omitempty" xml:"RepetitionRate"`
    ColumnDataDataType string `json:"ColumnDataDataType,omitempty" xml:"ColumnDataDataType"`
    NumberCategoryType string `json:"NumberCategoryType,omitempty" xml:"NumberCategoryType"`
    TextCategoryType string `json:"TextCategoryType,omitempty" xml:"TextCategoryType"`
    StyleNumber int64 `json:"StyleNumber,omitempty" xml:"StyleNumber"`
    ColumnDataExceptionDescription string `json:"columnDataExceptionDescription,omitempty" xml:"columnDataExceptionDescription"`
}
