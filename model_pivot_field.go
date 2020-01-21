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

type PivotField struct {
	PivotItems []PivotItem `json:"PivotItems,omitempty" xml:"PivotItems"`
	DisplayName string `json:"DisplayName,omitempty" xml:"DisplayName"`
	NumberFormat string `json:"NumberFormat,omitempty" xml:"NumberFormat"`
	DragToColumn bool `json:"DragToColumn,omitempty" xml:"DragToColumn"`
	IsAutoShow bool `json:"IsAutoShow,omitempty" xml:"IsAutoShow"`
	IsRepeatItemLabels bool `json:"IsRepeatItemLabels,omitempty" xml:"IsRepeatItemLabels"`
	DragToRow bool `json:"DragToRow,omitempty" xml:"DragToRow"`
	IsAutoSort bool `json:"IsAutoSort,omitempty" xml:"IsAutoSort"`
	InsertBlankRow bool `json:"InsertBlankRow,omitempty" xml:"InsertBlankRow"`
	ShowSubtotalAtTop bool `json:"ShowSubtotalAtTop,omitempty" xml:"ShowSubtotalAtTop"`
	ShowCompact bool `json:"ShowCompact,omitempty" xml:"ShowCompact"`
	Function string `json:"Function,omitempty" xml:"Function"`
	IsMultipleItemSelectionAllowed bool `json:"IsMultipleItemSelectionAllowed,omitempty" xml:"IsMultipleItemSelectionAllowed"`
	DataDisplayFormat string `json:"DataDisplayFormat,omitempty" xml:"DataDisplayFormat"`
	BaseItemPosition string `json:"BaseItemPosition,omitempty" xml:"BaseItemPosition"`
	IsInsertPageBreaksBetweenItems bool `json:"IsInsertPageBreaksBetweenItems,omitempty" xml:"IsInsertPageBreaksBetweenItems"`
	ShowAllItems bool `json:"ShowAllItems,omitempty" xml:"ShowAllItems"`
	BaseItem int32 `json:"BaseItem,omitempty" xml:"BaseItem"`
	ItemCount int32 `json:"ItemCount,omitempty" xml:"ItemCount"`
	Name string `json:"Name,omitempty" xml:"Name"`
	ShowInOutlineForm bool `json:"ShowInOutlineForm,omitempty" xml:"ShowInOutlineForm"`
	Items []string `json:"Items,omitempty" xml:"Items"`
	AutoShowField int32 `json:"AutoShowField,omitempty" xml:"AutoShowField"`
	IsAutoSubtotals bool `json:"IsAutoSubtotals,omitempty" xml:"IsAutoSubtotals"`
	IsIncludeNewItemsInFilter bool `json:"IsIncludeNewItemsInFilter,omitempty" xml:"IsIncludeNewItemsInFilter"`
	CurrentPageItem int32 `json:"CurrentPageItem,omitempty" xml:"CurrentPageItem"`
	Position int32 `json:"Position,omitempty" xml:"Position"`
	IsAscendSort bool `json:"IsAscendSort,omitempty" xml:"IsAscendSort"`
	IsAscendShow bool `json:"IsAscendShow,omitempty" xml:"IsAscendShow"`
	BaseField int32 `json:"BaseField,omitempty" xml:"BaseField"`
	AutoSortField int32 `json:"AutoSortField,omitempty" xml:"AutoSortField"`
	AutoShowCount int32 `json:"AutoShowCount,omitempty" xml:"AutoShowCount"`
	Number int32 `json:"Number,omitempty" xml:"Number"`
	DragToPage bool `json:"DragToPage,omitempty" xml:"DragToPage"`
	DragToData bool `json:"DragToData,omitempty" xml:"DragToData"`
	BaseIndex int32 `json:"BaseIndex,omitempty" xml:"BaseIndex"`
	OriginalItems []string `json:"OriginalItems,omitempty" xml:"OriginalItems"`
	DragToHide bool `json:"DragToHide,omitempty" xml:"DragToHide"`
	IsCalculatedField bool `json:"IsCalculatedField,omitempty" xml:"IsCalculatedField"`
}
