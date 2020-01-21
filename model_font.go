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

type Font struct {
	IsItalic bool `json:"IsItalic,omitempty" xml:"IsItalic"`
	Name string `json:"Name,omitempty" xml:"Name"`
	DoubleSize float64 `json:"DoubleSize,omitempty" xml:"DoubleSize"`
	Color *Color `json:"Color,omitempty" xml:"Color"`
	IsBold bool `json:"IsBold,omitempty" xml:"IsBold"`
	IsSubscript bool `json:"IsSubscript,omitempty" xml:"IsSubscript"`
	IsSuperscript bool `json:"IsSuperscript,omitempty" xml:"IsSuperscript"`
	IsStrikeout bool `json:"IsStrikeout,omitempty" xml:"IsStrikeout"`
	Underline string `json:"Underline,omitempty" xml:"Underline"`
	Size int32 `json:"Size,omitempty" xml:"Size"`
}
