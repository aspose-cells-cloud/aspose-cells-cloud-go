/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="WorksheetDataStatistics.go">
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

type WorksheetDataStatistics struct {
 
    Name string `json:"Name,omitempty" xml:"Name"`
    ChartsCount int64 `json:"ChartsCount,omitempty" xml:"ChartsCount"`
    TablesCount int64 `json:"TablesCount,omitempty" xml:"TablesCount"`
    PivotTablesCount int64 `json:"PivotTablesCount,omitempty" xml:"PivotTablesCount"`
    ShapesCount int64 `json:"ShapesCount,omitempty" xml:"ShapesCount"`
    HyperlinksCount int64 `json:"HyperlinksCount,omitempty" xml:"HyperlinksCount"`
    QueryTablesCount int64 `json:"QueryTablesCount,omitempty" xml:"QueryTablesCount"`
    CellsCount int64 `json:"CellsCount,omitempty" xml:"CellsCount"`
    CellsCountInTable int64 `json:"CellsCountInTable,omitempty" xml:"CellsCountInTable"`
    CellsCountIsFormula int64 `json:"CellsCountIsFormula,omitempty" xml:"CellsCountIsFormula"`
}
