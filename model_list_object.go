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

type ListObject struct {
	Link *Link `json:"link,omitempty" xml:"link"`
	// Gets and sets whether this ListObject show total row.
	ShowTotals bool `json:"ShowTotals,omitempty" xml:"ShowTotals"`
	// Gets and the built-in table style.
	TableStyleType string `json:"TableStyleType,omitempty" xml:"TableStyleType"`
	// Gets and sets the display name.Gets the data range of the ListObject.
	DisplayName string `json:"DisplayName,omitempty" xml:"DisplayName"`
	// Gets and sets whether this ListObject show header row.             
	ShowHeaderRow bool `json:"ShowHeaderRow,omitempty" xml:"ShowHeaderRow"`
	// Gets the start column of the range.
	StartColumn int32 `json:"StartColumn,omitempty" xml:"StartColumn"`
	// Indicates whether the last column in the table should have the style applied.
	ShowTableStyleLastColumn bool `json:"ShowTableStyleLastColumn,omitempty" xml:"ShowTableStyleLastColumn"`
	// Indicates whether column stripe formatting is applied.
	ShowTableStyleColumnStripes bool `json:"ShowTableStyleColumnStripes,omitempty" xml:"ShowTableStyleColumnStripes"`
	// Inidicates whether the first column in the table should have the style applied.
	ShowTableStyleFirstColumn bool `json:"ShowTableStyleFirstColumn,omitempty" xml:"ShowTableStyleFirstColumn"`
	// Gets the start row of the range.
	StartRow int32 `json:"StartRow,omitempty" xml:"StartRow"`
	// Gets auto filter.             
	AutoFilter *AutoFilter `json:"AutoFilter,omitempty" xml:"AutoFilter"`
	// Indicates whether row stripe formatting is applied.
	ShowTableStyleRowStripes bool `json:"ShowTableStyleRowStripes,omitempty" xml:"ShowTableStyleRowStripes"`
	// Gets the end column of the range.
	EndColumn int32 `json:"EndColumn,omitempty" xml:"EndColumn"`
	// Gets and sets the table style name.
	TableStyleName string `json:"TableStyleName,omitempty" xml:"TableStyleName"`
	// Gets ListColumns of the ListObject.
	ListColumns []ListColumn `json:"ListColumns,omitempty" xml:"ListColumns"`
	// Gets the end row of the range.
	EndRow int32 `json:"EndRow,omitempty" xml:"EndRow"`
}
