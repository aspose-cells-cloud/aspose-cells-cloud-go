/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="PivotFilter.go">
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

type PivotFilter struct {
 
    AutoFilter *AutoFilter `json:"AutoFilter,omitempty" xml:"AutoFilter"`
    EvaluationOrder int64 `json:"EvaluationOrder,omitempty" xml:"EvaluationOrder"`
    FieldIndex int64 `json:"FieldIndex,omitempty" xml:"FieldIndex"`
    FilterType string `json:"FilterType,omitempty" xml:"FilterType"`
    ValueFieldIndex int64 `json:"ValueFieldIndex,omitempty" xml:"ValueFieldIndex"`
    MemberPropertyFieldIndex int64 `json:"MemberPropertyFieldIndex,omitempty" xml:"MemberPropertyFieldIndex"`
    Name string `json:"Name,omitempty" xml:"Name"`
    Value1 string `json:"Value1,omitempty" xml:"Value1"`
    Value2 string `json:"Value2,omitempty" xml:"Value2"`
    Top10Filter *Top10Filter `json:"Top10Filter,omitempty" xml:"Top10Filter"`
}
