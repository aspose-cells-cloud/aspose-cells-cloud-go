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

type Worksheet struct {
	// Gets the index of sheet in the worksheets collection.             
	Index int64 `json:"Index" xml:"Index"`
	Pictures *LinkElement `json:"Pictures,omitempty" xml:"Pictures"`
	Charts *LinkElement `json:"Charts,omitempty" xml:"Charts"`
	Comments *LinkElement `json:"Comments,omitempty" xml:"Comments"`
	Hyperlinks *LinkElement `json:"Hyperlinks,omitempty" xml:"Hyperlinks"`
	// Represents if the worksheet is visible.             
	IsVisible bool `json:"IsVisible,omitempty" xml:"IsVisible"`
	// Gets and sets the view type.
	ViewType string `json:"ViewType,omitempty" xml:"ViewType"`
	// Represents worksheet type
	Type_ string `json:"Type,omitempty" xml:"Type"`
	// Gets or sets a value indicating whether the gridelines are visible.Default     is true.
	IsGridlinesVisible bool `json:"IsGridlinesVisible,omitempty" xml:"IsGridlinesVisible"`
	// Gets or sets a value indicating whether the worksheet will display row and column headers.Default is true.             
	IsRowColumnHeadersVisible bool `json:"IsRowColumnHeadersVisible,omitempty" xml:"IsRowColumnHeadersVisible"`
	// Indications the specified worksheet is shown in normal view or page break preview.
	IsPageBreakPreview bool `json:"IsPageBreakPreview,omitempty" xml:"IsPageBreakPreview"`
	// True if zero values are displayed.
	DisplayZeros bool `json:"DisplayZeros,omitempty" xml:"DisplayZeros"`
	// Flag indicating whether the Transition Formula Evaluation (Lotus compatibility) option is enabled.             
	TransitionEvaluation bool `json:"TransitionEvaluation,omitempty" xml:"TransitionEvaluation"`
	// Indicates if the specified worksheet is displayed from right to left instead    of from left to right.  Default is false.             
	DisplayRightToLeft bool `json:"DisplayRightToLeft,omitempty" xml:"DisplayRightToLeft"`
	// Represents first visible column index.
	FirstVisibleColumn int64 `json:"FirstVisibleColumn,omitempty" xml:"FirstVisibleColumn"`
	OleObjects *LinkElement `json:"OleObjects,omitempty" xml:"OleObjects"`
	// Indicates whether show outline.             
	IsOutlineShown bool `json:"IsOutlineShown,omitempty" xml:"IsOutlineShown"`
	// Gets or sets the name of the worksheet.             
	Name string `json:"Name,omitempty" xml:"Name"`
	AutoShapes *LinkElement `json:"AutoShapes,omitempty" xml:"AutoShapes"`
	Cells *LinkElement `json:"Cells,omitempty" xml:"Cells"`
	Validations *LinkElement `json:"Validations,omitempty" xml:"Validations"`
	// Represents the scaling factor in percent. It should be btween 10 and 400.             
	Zoom int64 `json:"Zoom,omitempty" xml:"Zoom"`
	ConditionalFormattings *LinkElement `json:"ConditionalFormattings,omitempty" xml:"ConditionalFormattings"`
	// Indicates whether this worksheet is selected when the workbook is opened.
	IsSelected bool `json:"IsSelected,omitempty" xml:"IsSelected"`
	// Represents worksheet tab color.
	TabColor *Color `json:"TabColor,omitempty" xml:"TabColor"`
	// Represents first visible row index.             
	FirstVisibleRow int64 `json:"FirstVisibleRow,omitempty" xml:"FirstVisibleRow"`
	// Flag indicating whether the Transition Formula Entry (Lotus compatibility) option is enabled.
	TransitionEntry bool `json:"TransitionEntry,omitempty" xml:"TransitionEntry"`
	// Indicates the state for this sheet visibility             
	VisibilityType string `json:"VisibilityType,omitempty" xml:"VisibilityType"`
	// Indicates whether the ruler is visible. Only apply for page break preview.
	IsRulerVisible bool `json:"IsRulerVisible,omitempty" xml:"IsRulerVisible"`
	Links []Link `json:"Links,omitempty" xml:"Links"`
	// Indicates if the worksheet is protected.
	IsProtected bool `json:"IsProtected" xml:"IsProtected"`
	MergedCells *LinkElement `json:"MergedCells,omitempty" xml:"MergedCells"`
}
