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

type FormatCondition struct {
	Link *Link `json:"link,omitempty" xml:"link"`
	AboveAverage *AboveAverage `json:"AboveAverage,omitempty" xml:"AboveAverage"`
	Formula2 string `json:"Formula2,omitempty" xml:"Formula2"`
	Style *Style `json:"Style,omitempty" xml:"Style"`
	Formula1 string `json:"Formula1,omitempty" xml:"Formula1"`
	ColorScale *ColorScale `json:"ColorScale,omitempty" xml:"ColorScale"`
	DataBar *DataBar `json:"DataBar,omitempty" xml:"DataBar"`
	Text string `json:"Text,omitempty" xml:"Text"`
	// True, no rules with lower priority may be applied over this rule, when this     rule evaluates to true.  Only applies for Excel 2007;
	StopIfTrue bool `json:"StopIfTrue,omitempty" xml:"StopIfTrue"`
	// The priority of this conditional formatting rule. This value is used to determine     which format should be evaluated and rendered. Lower numeric values are higher     priority than higher numeric values, where '1' is the highest priority.
	Priority int64 `json:"Priority,omitempty" xml:"Priority"`
	Top10 *Top10 `json:"Top10,omitempty" xml:"Top10"`
	Operator string `json:"Operator,omitempty" xml:"Operator"`
	IconSet *IconSet `json:"IconSet,omitempty" xml:"IconSet"`
	// Gets and sets whether the conditional format Type.             
	Type_ string `json:"Type,omitempty" xml:"Type"`
	TimePeriod string `json:"TimePeriod,omitempty" xml:"TimePeriod"`
}
