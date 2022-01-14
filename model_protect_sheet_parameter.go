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

type ProtectSheetParameter struct {
	AllowSelectingUnlockedCell string `json:"AllowSelectingUnlockedCell,omitempty" xml:"AllowSelectingUnlockedCell"`
	AllowFiltering string `json:"AllowFiltering,omitempty" xml:"AllowFiltering"`
	AllowDeletingColumn string `json:"AllowDeletingColumn,omitempty" xml:"AllowDeletingColumn"`
	AllowSelectingLockedCell string `json:"AllowSelectingLockedCell,omitempty" xml:"AllowSelectingLockedCell"`
	AllowUsingPivotTable string `json:"AllowUsingPivotTable,omitempty" xml:"AllowUsingPivotTable"`
	AllowEditArea []string `json:"AllowEditArea,omitempty" xml:"AllowEditArea"`
	AllowInsertingHyperlink string `json:"AllowInsertingHyperlink,omitempty" xml:"AllowInsertingHyperlink"`
	AllowFormattingCell string `json:"AllowFormattingCell,omitempty" xml:"AllowFormattingCell"`
	AllowFormattingRow string `json:"AllowFormattingRow,omitempty" xml:"AllowFormattingRow"`
	AllowInsertingRow string `json:"AllowInsertingRow,omitempty" xml:"AllowInsertingRow"`
	AllowFormattingColumn string `json:"AllowFormattingColumn,omitempty" xml:"AllowFormattingColumn"`
	AllowSorting string `json:"AllowSorting,omitempty" xml:"AllowSorting"`
	AllowInsertingColumn string `json:"AllowInsertingColumn,omitempty" xml:"AllowInsertingColumn"`
	Password string `json:"Password,omitempty" xml:"Password"`
	AllowDeletingRow string `json:"AllowDeletingRow,omitempty" xml:"AllowDeletingRow"`
	ProtectionType string `json:"ProtectionType,omitempty" xml:"ProtectionType"`
}
