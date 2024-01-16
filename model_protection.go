/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="Protection.go">
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

type Protection struct {
 
    AllowDeletingColumn bool `json:"AllowDeletingColumn,omitempty" xml:"AllowDeletingColumn"`
    AllowDeletingRow bool `json:"AllowDeletingRow,omitempty" xml:"AllowDeletingRow"`
    AllowFiltering bool `json:"AllowFiltering,omitempty" xml:"AllowFiltering"`
    AllowFormattingCell bool `json:"AllowFormattingCell,omitempty" xml:"AllowFormattingCell"`
    AllowFormattingColumn bool `json:"AllowFormattingColumn,omitempty" xml:"AllowFormattingColumn"`
    AllowFormattingRow bool `json:"AllowFormattingRow,omitempty" xml:"AllowFormattingRow"`
    AllowInsertingColumn bool `json:"AllowInsertingColumn,omitempty" xml:"AllowInsertingColumn"`
    AllowInsertingHyperlink bool `json:"AllowInsertingHyperlink,omitempty" xml:"AllowInsertingHyperlink"`
    AllowInsertingRow bool `json:"AllowInsertingRow,omitempty" xml:"AllowInsertingRow"`
    AllowSorting bool `json:"AllowSorting,omitempty" xml:"AllowSorting"`
    AllowUsingPivotTable bool `json:"AllowUsingPivotTable,omitempty" xml:"AllowUsingPivotTable"`
    AllowEditingContent bool `json:"AllowEditingContent,omitempty" xml:"AllowEditingContent"`
    AllowEditingObject bool `json:"AllowEditingObject,omitempty" xml:"AllowEditingObject"`
    AllowEditingScenario bool `json:"AllowEditingScenario,omitempty" xml:"AllowEditingScenario"`
    Password string `json:"Password,omitempty" xml:"Password"`
    AllowSelectingLockedCell bool `json:"AllowSelectingLockedCell,omitempty" xml:"AllowSelectingLockedCell"`
    AllowSelectingUnlockedCell bool `json:"AllowSelectingUnlockedCell,omitempty" xml:"AllowSelectingUnlockedCell"`
}
