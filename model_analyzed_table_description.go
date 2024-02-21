/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="AnalyzedTableDescription.go">
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

type AnalyzedTableDescription struct {
 
    Name string `json:"Name,omitempty" xml:"Name"`
    SheetName string `json:"SheetName,omitempty" xml:"SheetName"`
    Columns []AnalyzedColumnDescription `json:"Columns,omitempty" xml:"Columns"`
    DateColumns []int64 `json:"DateColumns,omitempty" xml:"DateColumns"`
    NumberColumns []int64 `json:"NumberColumns,omitempty" xml:"NumberColumns"`
    TextColumns []int64 `json:"TextColumns,omitempty" xml:"TextColumns"`
    ExceptionColumns []int64 `json:"ExceptionColumns,omitempty" xml:"ExceptionColumns"`
    HasTableHeaderRow bool `json:"HasTableHeaderRow,omitempty" xml:"HasTableHeaderRow"`
    HasTableTotalRow bool `json:"HasTableTotalRow,omitempty" xml:"HasTableTotalRow"`
    StartDataColumnIndex int64 `json:"StartDataColumnIndex,omitempty" xml:"StartDataColumnIndex"`
    EndDataColumnIndex int64 `json:"EndDataColumnIndex,omitempty" xml:"EndDataColumnIndex"`
    StartDataRowIndex int64 `json:"StartDataRowIndex,omitempty" xml:"StartDataRowIndex"`
    EndDataRowIndex int64 `json:"EndDataRowIndex,omitempty" xml:"EndDataRowIndex"`
    Thumbnail string `json:"Thumbnail,omitempty" xml:"Thumbnail"`
    DiscoverCharts []DiscoverChart `json:"DiscoverCharts,omitempty" xml:"DiscoverCharts"`
    DiscoverPivotTables []DiscoverPivotTable `json:"DiscoverPivotTables,omitempty" xml:"DiscoverPivotTables"`
}
