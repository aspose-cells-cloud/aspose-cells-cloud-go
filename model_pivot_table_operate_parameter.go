/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="PivotTableOperateParameter.go">
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

type PivotTableOperateParameter struct {
     
        OperateType string `json:"OperateType,omitempty" xml:"OperateType"`
 
    SourceData string `json:"SourceData,omitempty" xml:"SourceData"`
    DestCellName string `json:"DestCellName,omitempty" xml:"DestCellName"`
    TableName string `json:"TableName,omitempty" xml:"TableName"`
    UseSameSource bool `json:"UseSameSource,omitempty" xml:"UseSameSource"`
    PivotTableIndex int64 `json:"PivotTableIndex,omitempty" xml:"PivotTableIndex"`
    PivotFieldRows []int64 `json:"PivotFieldRows,omitempty" xml:"PivotFieldRows"`
    PivotFieldColumns []int64 `json:"PivotFieldColumns,omitempty" xml:"PivotFieldColumns"`
    PivotFieldData []int64 `json:"PivotFieldData,omitempty" xml:"PivotFieldData"`
}
