/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="ChartOperateParameter.go">
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

type ChartOperateParameter struct {
     
        OperateType string `json:"OperateType,omitempty" xml:"OperateType"`
 
    ChartIndex int64 `json:"ChartIndex,omitempty" xml:"ChartIndex"`
    ChartType string `json:"ChartType,omitempty" xml:"ChartType"`
    UpperLeftRow int64 `json:"UpperLeftRow,omitempty" xml:"UpperLeftRow"`
    UpperLeftColumn int64 `json:"UpperLeftColumn,omitempty" xml:"UpperLeftColumn"`
    LowerRightRow int64 `json:"LowerRightRow,omitempty" xml:"LowerRightRow"`
    LowerRightColumn int64 `json:"LowerRightColumn,omitempty" xml:"LowerRightColumn"`
    Area string `json:"Area,omitempty" xml:"Area"`
    IsVertical bool `json:"IsVertical,omitempty" xml:"IsVertical"`
    CategoryData string `json:"CategoryData,omitempty" xml:"CategoryData"`
    IsAutoGetSerialName bool `json:"IsAutoGetSerialName,omitempty" xml:"IsAutoGetSerialName"`
    Title string `json:"Title,omitempty" xml:"Title"`
}
