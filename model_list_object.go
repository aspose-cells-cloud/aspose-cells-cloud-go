/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="ListObject.go">
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

type ListObject struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    AutoFilter *AutoFilter `json:"AutoFilter,omitempty" xml:"AutoFilter"`
    DisplayName string `json:"DisplayName,omitempty" xml:"DisplayName"`
    StartColumn int64 `json:"StartColumn,omitempty" xml:"StartColumn"`
    StartRow int64 `json:"StartRow,omitempty" xml:"StartRow"`
    EndColumn int64 `json:"EndColumn,omitempty" xml:"EndColumn"`
    EndRow int64 `json:"EndRow,omitempty" xml:"EndRow"`
    ListColumns []ListColumn `json:"ListColumns,omitempty" xml:"ListColumns"`
    ShowHeaderRow bool `json:"ShowHeaderRow,omitempty" xml:"ShowHeaderRow"`
    ShowTableStyleColumnStripes bool `json:"ShowTableStyleColumnStripes,omitempty" xml:"ShowTableStyleColumnStripes"`
    ShowTableStyleFirstColumn bool `json:"ShowTableStyleFirstColumn,omitempty" xml:"ShowTableStyleFirstColumn"`
    ShowTableStyleLastColumn bool `json:"ShowTableStyleLastColumn,omitempty" xml:"ShowTableStyleLastColumn"`
    ShowTableStyleRowStripes bool `json:"ShowTableStyleRowStripes,omitempty" xml:"ShowTableStyleRowStripes"`
    ShowTotals bool `json:"ShowTotals,omitempty" xml:"ShowTotals"`
    TableStyleName string `json:"TableStyleName,omitempty" xml:"TableStyleName"`
    TableStyleType string `json:"TableStyleType,omitempty" xml:"TableStyleType"`
}
