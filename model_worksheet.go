/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="Worksheet.go">
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

type Worksheet struct {
 
    Links []Link `json:"Links,omitempty" xml:"Links"`
    DisplayRightToLeft bool `json:"DisplayRightToLeft,omitempty" xml:"DisplayRightToLeft"`
    DisplayZeros bool `json:"DisplayZeros,omitempty" xml:"DisplayZeros"`
    FirstVisibleColumn int64 `json:"FirstVisibleColumn,omitempty" xml:"FirstVisibleColumn"`
    FirstVisibleRow int64 `json:"FirstVisibleRow,omitempty" xml:"FirstVisibleRow"`
    Name string `json:"Name,omitempty" xml:"Name"`
    Index int64 `json:"Index,omitempty" xml:"Index"`
    IsGridlinesVisible bool `json:"IsGridlinesVisible,omitempty" xml:"IsGridlinesVisible"`
    IsOutlineShown bool `json:"IsOutlineShown,omitempty" xml:"IsOutlineShown"`
    IsPageBreakPreview bool `json:"IsPageBreakPreview,omitempty" xml:"IsPageBreakPreview"`
    IsVisible bool `json:"IsVisible,omitempty" xml:"IsVisible"`
    IsProtected bool `json:"IsProtected,omitempty" xml:"IsProtected"`
    IsRowColumnHeadersVisible bool `json:"IsRowColumnHeadersVisible,omitempty" xml:"IsRowColumnHeadersVisible"`
    IsRulerVisible bool `json:"IsRulerVisible,omitempty" xml:"IsRulerVisible"`
    IsSelected bool `json:"IsSelected,omitempty" xml:"IsSelected"`
    TabColor *Color `json:"TabColor,omitempty" xml:"TabColor"`
    TransitionEntry bool `json:"TransitionEntry,omitempty" xml:"TransitionEntry"`
    TransitionEvaluation bool `json:"TransitionEvaluation,omitempty" xml:"TransitionEvaluation"`
    Type_ string `json:"Type,omitempty" xml:"Type"`
    ViewType string `json:"ViewType,omitempty" xml:"ViewType"`
    VisibilityType string `json:"VisibilityType,omitempty" xml:"VisibilityType"`
    Zoom int64 `json:"Zoom,omitempty" xml:"Zoom"`
    Cells *LinkElement `json:"Cells,omitempty" xml:"Cells"`
    Charts *LinkElement `json:"Charts,omitempty" xml:"Charts"`
    AutoShapes *LinkElement `json:"AutoShapes,omitempty" xml:"AutoShapes"`
    OleObjects *LinkElement `json:"OleObjects,omitempty" xml:"OleObjects"`
    Comments *LinkElement `json:"Comments,omitempty" xml:"Comments"`
    Pictures *LinkElement `json:"Pictures,omitempty" xml:"Pictures"`
    MergedCells *LinkElement `json:"MergedCells,omitempty" xml:"MergedCells"`
    Validations *LinkElement `json:"Validations,omitempty" xml:"Validations"`
    ConditionalFormattings *LinkElement `json:"ConditionalFormattings,omitempty" xml:"ConditionalFormattings"`
    Hyperlinks *LinkElement `json:"Hyperlinks,omitempty" xml:"Hyperlinks"`
}
