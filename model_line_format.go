/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="LineFormat.go">
*   Copyright (c) 2023 Aspose.Cells Cloud
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

type LineFormat struct {
     
        Type_ string `json:"Type,omitempty" xml:"Type"`
        SolidFill *SolidFill `json:"SolidFill,omitempty" xml:"SolidFill"`
        PatternFill *PatternFill `json:"PatternFill,omitempty" xml:"PatternFill"`
        TextureFill *TextureFill `json:"TextureFill,omitempty" xml:"TextureFill"`
        GradientFill *GradientFill `json:"GradientFill,omitempty" xml:"GradientFill"`
        ImageData string `json:"ImageData,omitempty" xml:"ImageData"`
 
    BeginArrowheadLength string `json:"BeginArrowheadLength,omitempty" xml:"BeginArrowheadLength"`
    BeginArrowheadStyle string `json:"BeginArrowheadStyle,omitempty" xml:"BeginArrowheadStyle"`
    BeginArrowheadWidth string `json:"BeginArrowheadWidth,omitempty" xml:"BeginArrowheadWidth"`
    CapType string `json:"CapType,omitempty" xml:"CapType"`
    CompoundType string `json:"CompoundType,omitempty" xml:"CompoundType"`
    DashStyle string `json:"DashStyle,omitempty" xml:"DashStyle"`
    EndArrowheadLength string `json:"EndArrowheadLength,omitempty" xml:"EndArrowheadLength"`
    EndArrowheadStyle string `json:"EndArrowheadStyle,omitempty" xml:"EndArrowheadStyle"`
    EndArrowheadWidth string `json:"EndArrowheadWidth,omitempty" xml:"EndArrowheadWidth"`
    JoinType string `json:"JoinType,omitempty" xml:"JoinType"`
    Weight float64 `json:"Weight,omitempty" xml:"Weight"`
}
