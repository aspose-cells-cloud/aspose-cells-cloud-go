/*
 *  Copyright (c) 2021 Aspose.Cells Cloud
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

type ImportPictureOption struct {
	Source *FileSource `json:"Source,omitempty" xml:"Source"`
	ImportDataType string `json:"ImportDataType,omitempty" xml:"ImportDataType"`
	DestinationWorksheet string `json:"DestinationWorksheet,omitempty" xml:"DestinationWorksheet"`
	IsInsert bool `json:"IsInsert,omitempty" xml:"IsInsert"`
	// Upper Left Row.
	UpperLeftRow int32 `json:"UpperLeftRow,omitempty" xml:"UpperLeftRow"`
	// Upper Left Column.
	UpperLeftColumn int32 `json:"UpperLeftColumn,omitempty" xml:"UpperLeftColumn"`
	// Lower Right Row.
	LowerRightRow int32 `json:"LowerRightRow,omitempty" xml:"LowerRightRow"`
	// Lower Right Column.
	LowerRightColumn int32 `json:"LowerRightColumn,omitempty" xml:"LowerRightColumn"`
	// Filename.
	Filename string `json:"Filename,omitempty" xml:"Filename"`
	// data : base64  string.
	Data string `json:"Data,omitempty" xml:"Data"`
}
