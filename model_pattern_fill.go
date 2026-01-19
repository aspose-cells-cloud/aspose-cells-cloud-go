/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="PatternFill.go">
*   Copyright (c) 2026 Aspose.Cells Cloud
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

type PatternFill struct {
 
    Pattern string `json:"Pattern,omitempty" xml:"Pattern"`
    BackgroundCellsColor *CellsColor `json:"BackgroundCellsColor,omitempty" xml:"BackgroundCellsColor"`
    ForegroundCellsColor *CellsColor `json:"ForegroundCellsColor,omitempty" xml:"ForegroundCellsColor"`
    ForegroundColor *Color `json:"ForegroundColor,omitempty" xml:"ForegroundColor"`
    BackgroundColor *Color `json:"BackgroundColor,omitempty" xml:"BackgroundColor"`
    BackTransparency float64 `json:"BackTransparency,omitempty" xml:"BackTransparency"`
    ForeTransparency float64 `json:"ForeTransparency,omitempty" xml:"ForeTransparency"`
}
