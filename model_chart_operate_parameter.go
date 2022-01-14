/*
 *  Copyright (c) 2022 Aspose.Cells Cloud
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

type ChartOperateParameter struct {
	OperateType string `json:"OperateType,omitempty" xml:"OperateType"`
	Title string `json:"Title,omitempty" xml:"Title"`
	Area string `json:"Area,omitempty" xml:"Area"`
	CategoryData string `json:"CategoryData,omitempty" xml:"CategoryData"`
	UpperLeftRow int64 `json:"UpperLeftRow,omitempty" xml:"UpperLeftRow"`
	LowerRightColumn int64 `json:"LowerRightColumn,omitempty" xml:"LowerRightColumn"`
	LowerRightRow int64 `json:"LowerRightRow,omitempty" xml:"LowerRightRow"`
	IsAutoGetSerialName bool `json:"IsAutoGetSerialName,omitempty" xml:"IsAutoGetSerialName"`
	ChartType string `json:"ChartType,omitempty" xml:"ChartType"`
	IsVertical bool `json:"IsVertical,omitempty" xml:"IsVertical"`
}
