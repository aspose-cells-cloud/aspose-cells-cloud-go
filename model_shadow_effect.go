/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="ShadowEffect.go">
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

type ShadowEffect struct {
 
    Angle float64 `json:"Angle,omitempty" xml:"Angle"`
    Blur float64 `json:"Blur,omitempty" xml:"Blur"`
    Color *CellsColor `json:"Color,omitempty" xml:"Color"`
    Distance float64 `json:"Distance,omitempty" xml:"Distance"`
    PresetType string `json:"PresetType,omitempty" xml:"PresetType"`
    Size float64 `json:"Size,omitempty" xml:"Size"`
    Transparency float64 `json:"Transparency,omitempty" xml:"Transparency"`
}
