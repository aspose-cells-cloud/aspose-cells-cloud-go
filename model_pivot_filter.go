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

// Represents a PivotFilter in PivotFilter Collection.
type PivotFilter struct {
	// Gets the Evaluation Order of the pivot filter.
	EvaluationOrder int64 `json:"EvaluationOrder,omitempty" xml:"EvaluationOrder"`
	// Gets the name of the pivot filter.
	Name string `json:"Name,omitempty" xml:"Name"`
	// Gets the autofilter type of the pivot filter.
	FilterType string `json:"FilterType,omitempty" xml:"FilterType"`
	// Gets the autofilter of the pivot filter.
	AutoFilter *AutoFilter `json:"AutoFilter,omitempty" xml:"AutoFilter"`
	// Gets the field index of the pivot filter.
	FieldIndex int64 `json:"FieldIndex,omitempty" xml:"FieldIndex"`
	// Gets the measure field index of the pivot filter.             
	MeasureFldIndex int64 `json:"MeasureFldIndex,omitempty" xml:"MeasureFldIndex"`
	// Gets the string value1 of the label pivot filter.             
	Value1 string `json:"Value1,omitempty" xml:"Value1"`
	// Gets the member property field index of the pivot filter.             
	MemberPropertyFieldIndex int64 `json:"MemberPropertyFieldIndex,omitempty" xml:"MemberPropertyFieldIndex"`
	// Gets the string value2 of the label pivot filter.             
	Value2 string `json:"Value2,omitempty" xml:"Value2"`
}
