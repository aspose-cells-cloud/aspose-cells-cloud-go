/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="FormatCondition.go">
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

type FormatCondition struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    Priority int64 `json:"Priority,omitempty" xml:"Priority"`
    Type_ string `json:"Type,omitempty" xml:"Type"`
    StopIfTrue bool `json:"StopIfTrue,omitempty" xml:"StopIfTrue"`
    AboveAverage *AboveAverage `json:"AboveAverage,omitempty" xml:"AboveAverage"`
    ColorScale *ColorScale `json:"ColorScale,omitempty" xml:"ColorScale"`
    DataBar *DataBar `json:"DataBar,omitempty" xml:"DataBar"`
    Formula1 string `json:"Formula1,omitempty" xml:"Formula1"`
    Formula2 string `json:"Formula2,omitempty" xml:"Formula2"`
    IconSet *IconSet `json:"IconSet,omitempty" xml:"IconSet"`
    Operator string `json:"Operator,omitempty" xml:"Operator"`
    Style *Style `json:"Style,omitempty" xml:"Style"`
    Text string `json:"Text,omitempty" xml:"Text"`
    TimePeriod string `json:"TimePeriod,omitempty" xml:"TimePeriod"`
    Top10 *Top10 `json:"Top10,omitempty" xml:"Top10"`
}
