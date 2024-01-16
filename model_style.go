/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="Style.go">
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

type Style struct {
 
    Font *Font `json:"Font,omitempty" xml:"Font"`
    Name string `json:"Name,omitempty" xml:"Name"`
    CultureCustom string `json:"CultureCustom,omitempty" xml:"CultureCustom"`
    Custom string `json:"Custom,omitempty" xml:"Custom"`
    BackgroundColor *Color `json:"BackgroundColor,omitempty" xml:"BackgroundColor"`
    ForegroundColor *Color `json:"ForegroundColor,omitempty" xml:"ForegroundColor"`
    IsFormulaHidden bool `json:"IsFormulaHidden,omitempty" xml:"IsFormulaHidden"`
    IsDateTime bool `json:"IsDateTime,omitempty" xml:"IsDateTime"`
    IsTextWrapped bool `json:"IsTextWrapped,omitempty" xml:"IsTextWrapped"`
    IsGradient bool `json:"IsGradient,omitempty" xml:"IsGradient"`
    IsLocked bool `json:"IsLocked,omitempty" xml:"IsLocked"`
    IsPercent bool `json:"IsPercent,omitempty" xml:"IsPercent"`
    ShrinkToFit bool `json:"ShrinkToFit,omitempty" xml:"ShrinkToFit"`
    IndentLevel int64 `json:"IndentLevel,omitempty" xml:"IndentLevel"`
    Number int64 `json:"Number,omitempty" xml:"Number"`
    RotationAngle int64 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
    Pattern string `json:"Pattern,omitempty" xml:"Pattern"`
    TextDirection string `json:"TextDirection,omitempty" xml:"TextDirection"`
    VerticalAlignment string `json:"VerticalAlignment,omitempty" xml:"VerticalAlignment"`
    HorizontalAlignment string `json:"HorizontalAlignment,omitempty" xml:"HorizontalAlignment"`
    BorderCollection []Border `json:"BorderCollection,omitempty" xml:"BorderCollection"`
    BackgroundThemeColor *ThemeColor `json:"BackgroundThemeColor,omitempty" xml:"BackgroundThemeColor"`
    ForegroundThemeColor *ThemeColor `json:"ForegroundThemeColor,omitempty" xml:"ForegroundThemeColor"`
}
