/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="CellsColor.go">
*   Copyright (c) 2025 Aspose.Cells Cloud
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

type CellsColor struct {
 
    Color *Color `json:"Color,omitempty" xml:"Color"`
    ColorIndex int64 `json:"ColorIndex,omitempty" xml:"ColorIndex"`
    IsShapeColor bool `json:"IsShapeColor,omitempty" xml:"IsShapeColor"`
    Tint float64 `json:"tint,omitempty" xml:"tint"`
    Argb int64 `json:"Argb,omitempty" xml:"Argb"`
    ThemeColor *ThemeColor `json:"ThemeColor,omitempty" xml:"ThemeColor"`
    Type_ string `json:"Type,omitempty" xml:"Type"`
    Transparency float64 `json:"Transparency,omitempty" xml:"Transparency"`
}
