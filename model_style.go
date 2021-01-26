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

type Style struct {
	Link *Link `json:"link,omitempty" xml:"link"`
	Pattern string `json:"Pattern,omitempty" xml:"Pattern"`
	TextDirection string `json:"TextDirection,omitempty" xml:"TextDirection"`
	Custom string `json:"Custom,omitempty" xml:"Custom"`
	ShrinkToFit bool `json:"ShrinkToFit,omitempty" xml:"ShrinkToFit"`
	IsDateTime bool `json:"IsDateTime,omitempty" xml:"IsDateTime"`
	CultureCustom string `json:"CultureCustom,omitempty" xml:"CultureCustom"`
	RotationAngle int64 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
	IndentLevel int64 `json:"IndentLevel,omitempty" xml:"IndentLevel"`
	IsPercent bool `json:"IsPercent,omitempty" xml:"IsPercent"`
	ForegroundColor *Color `json:"ForegroundColor,omitempty" xml:"ForegroundColor"`
	Name string `json:"Name,omitempty" xml:"Name"`
	ForegroundThemeColor *ThemeColor `json:"ForegroundThemeColor,omitempty" xml:"ForegroundThemeColor"`
	BorderCollection []Border `json:"BorderCollection,omitempty" xml:"BorderCollection"`
	IsLocked bool `json:"IsLocked,omitempty" xml:"IsLocked"`
	VerticalAlignment string `json:"VerticalAlignment,omitempty" xml:"VerticalAlignment"`
	BackgroundColor *Color `json:"BackgroundColor,omitempty" xml:"BackgroundColor"`
	BackgroundThemeColor *ThemeColor `json:"BackgroundThemeColor,omitempty" xml:"BackgroundThemeColor"`
	IsFormulaHidden bool `json:"IsFormulaHidden,omitempty" xml:"IsFormulaHidden"`
	IsGradient bool `json:"IsGradient,omitempty" xml:"IsGradient"`
	Number int64 `json:"Number,omitempty" xml:"Number"`
	HorizontalAlignment string `json:"HorizontalAlignment,omitempty" xml:"HorizontalAlignment"`
	IsTextWrapped bool `json:"IsTextWrapped,omitempty" xml:"IsTextWrapped"`
	Font *Font `json:"Font,omitempty" xml:"Font"`
}
