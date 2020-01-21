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

type ModelRange struct {
	// Gets the count of columns in the range.
	ColumnCount int32 `json:"ColumnCount" xml:"ColumnCount"`
	// Sets or gets the height of rows in this range
	RowHeight float64 `json:"RowHeight" xml:"RowHeight"`
	// Gets or sets the name of the range.
	Name string `json:"Name,omitempty" xml:"Name"`
	// Gets the index of the first column of the range.
	FirstColumn int32 `json:"FirstColumn" xml:"FirstColumn"`
	// Sets or gets the column width of this range
	ColumnWidth float64 `json:"ColumnWidth" xml:"ColumnWidth"`
	// Gets the range's refers to.
	RefersTo string `json:"RefersTo,omitempty" xml:"RefersTo"`
	// Gets the count of rows in the range.
	RowCount int32 `json:"RowCount" xml:"RowCount"`
	// Gets the index of the first row of the range.
	FirstRow int32 `json:"FirstRow" xml:"FirstRow"`
	// Gets the Aspose.Cells.Range.Worksheetobject which contains this range.
	Worksheet string `json:"Worksheet,omitempty" xml:"Worksheet"`
}
