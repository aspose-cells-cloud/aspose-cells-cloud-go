/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="DataBar.go">
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

type DataBar struct {
 
    AxisColor *Color `json:"AxisColor,omitempty" xml:"AxisColor"`
    AxisPosition string `json:"AxisPosition,omitempty" xml:"AxisPosition"`
    BarBorder *DataBarBorder `json:"BarBorder,omitempty" xml:"BarBorder"`
    BarFillType string `json:"BarFillType,omitempty" xml:"BarFillType"`
    Color *Color `json:"Color,omitempty" xml:"Color"`
    Direction string `json:"Direction,omitempty" xml:"Direction"`
    MaxCfvo *ConditionalFormattingValue `json:"MaxCfvo,omitempty" xml:"MaxCfvo"`
    MaxLength int64 `json:"MaxLength,omitempty" xml:"MaxLength"`
    MinCfvo *ConditionalFormattingValue `json:"MinCfvo,omitempty" xml:"MinCfvo"`
    MinLength int64 `json:"MinLength,omitempty" xml:"MinLength"`
    NegativeBarFormat *NegativeBarFormat `json:"NegativeBarFormat,omitempty" xml:"NegativeBarFormat"`
    ShowValue bool `json:"ShowValue,omitempty" xml:"ShowValue"`
}
