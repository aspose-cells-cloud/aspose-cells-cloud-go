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

type PivotTable struct {
	Link *Link `json:"link,omitempty" xml:"link"`
	ShowPivotStyleLastColumn bool `json:"ShowPivotStyleLastColumn,omitempty" xml:"ShowPivotStyleLastColumn"`
	RowHeaderCaption string `json:"RowHeaderCaption,omitempty" xml:"RowHeaderCaption"`
	ColumnRange *CellArea `json:"ColumnRange,omitempty" xml:"ColumnRange"`
	RefreshDataOnOpeningFile bool `json:"RefreshDataOnOpeningFile,omitempty" xml:"RefreshDataOnOpeningFile"`
	PageFields []PivotField `json:"PageFields,omitempty" xml:"PageFields"`
	DataFields []PivotField `json:"DataFields,omitempty" xml:"DataFields"`
	DataBodyRange *CellArea `json:"DataBodyRange,omitempty" xml:"DataBodyRange"`
	ShowDrill bool `json:"ShowDrill,omitempty" xml:"ShowDrill"`
	RefreshDataFlag bool `json:"RefreshDataFlag,omitempty" xml:"RefreshDataFlag"`
	ColumnGrand bool `json:"ColumnGrand,omitempty" xml:"ColumnGrand"`
	PivotTableStyleName string `json:"PivotTableStyleName,omitempty" xml:"PivotTableStyleName"`
	PivotFilters []PivotFilter `json:"PivotFilters,omitempty" xml:"PivotFilters"`
	NullString string `json:"NullString,omitempty" xml:"NullString"`
	ItemPrintTitles bool `json:"ItemPrintTitles,omitempty" xml:"ItemPrintTitles"`
	DisplayNullString bool `json:"DisplayNullString,omitempty" xml:"DisplayNullString"`
	EnableFieldList bool `json:"EnableFieldList,omitempty" xml:"EnableFieldList"`
	TableRange2 *CellArea `json:"TableRange2,omitempty" xml:"TableRange2"`
	RowFields []PivotField `json:"RowFields,omitempty" xml:"RowFields"`
	PageFieldOrder string `json:"PageFieldOrder,omitempty" xml:"PageFieldOrder"`
	AutoFormatType string `json:"AutoFormatType,omitempty" xml:"AutoFormatType"`
	EnableDataValueEditing bool `json:"EnableDataValueEditing,omitempty" xml:"EnableDataValueEditing"`
	ShowPivotStyleRowHeader bool `json:"ShowPivotStyleRowHeader,omitempty" xml:"ShowPivotStyleRowHeader"`
	IsGridDropZones bool `json:"IsGridDropZones,omitempty" xml:"IsGridDropZones"`
	EnableWizard bool `json:"EnableWizard,omitempty" xml:"EnableWizard"`
	ShowMemberPropertyTips bool `json:"ShowMemberPropertyTips,omitempty" xml:"ShowMemberPropertyTips"`
	AltTextDescription string `json:"AltTextDescription,omitempty" xml:"AltTextDescription"`
	ShowDataTips bool `json:"ShowDataTips,omitempty" xml:"ShowDataTips"`
	PrintTitles bool `json:"PrintTitles,omitempty" xml:"PrintTitles"`
	TableRange1 *CellArea `json:"TableRange1,omitempty" xml:"TableRange1"`
	ShowEmptyRow bool `json:"ShowEmptyRow,omitempty" xml:"ShowEmptyRow"`
	IsMultipleFieldFilters bool `json:"IsMultipleFieldFilters,omitempty" xml:"IsMultipleFieldFilters"`
	ShowEmptyCol bool `json:"ShowEmptyCol,omitempty" xml:"ShowEmptyCol"`
	ShowRowHeaderCaption bool `json:"ShowRowHeaderCaption,omitempty" xml:"ShowRowHeaderCaption"`
	HasBlankRows bool `json:"HasBlankRows,omitempty" xml:"HasBlankRows"`
	DataSource []string `json:"DataSource,omitempty" xml:"DataSource"`
	Tag string `json:"Tag,omitempty" xml:"Tag"`
	EnableDrilldown bool `json:"EnableDrilldown,omitempty" xml:"EnableDrilldown"`
	Indent int32 `json:"Indent,omitempty" xml:"Indent"`
	Name string `json:"Name,omitempty" xml:"Name"`
	RowGrand bool `json:"RowGrand,omitempty" xml:"RowGrand"`
	GrandTotalName string `json:"GrandTotalName,omitempty" xml:"GrandTotalName"`
	DisplayErrorString bool `json:"DisplayErrorString,omitempty" xml:"DisplayErrorString"`
	RowRange *CellArea `json:"RowRange,omitempty" xml:"RowRange"`
	IsSelected bool `json:"IsSelected,omitempty" xml:"IsSelected"`
	ColumnFields []PivotField `json:"ColumnFields,omitempty" xml:"ColumnFields"`
	ColumnHeaderCaption string `json:"ColumnHeaderCaption,omitempty" xml:"ColumnHeaderCaption"`
	ShowValuesRow bool `json:"ShowValuesRow,omitempty" xml:"ShowValuesRow"`
	EnableFieldDialog bool `json:"EnableFieldDialog,omitempty" xml:"EnableFieldDialog"`
	MissingItemsLimit string `json:"MissingItemsLimit,omitempty" xml:"MissingItemsLimit"`
	ShowPivotStyleRowStripes bool `json:"ShowPivotStyleRowStripes,omitempty" xml:"ShowPivotStyleRowStripes"`
	ManualUpdate bool `json:"ManualUpdate,omitempty" xml:"ManualUpdate"`
	IsAutoFormat bool `json:"IsAutoFormat,omitempty" xml:"IsAutoFormat"`
	DisplayImmediateItems bool `json:"DisplayImmediateItems,omitempty" xml:"DisplayImmediateItems"`
	ErrorString string `json:"ErrorString,omitempty" xml:"ErrorString"`
	CustomListSort bool `json:"CustomListSort,omitempty" xml:"CustomListSort"`
	MergeLabels bool `json:"MergeLabels,omitempty" xml:"MergeLabels"`
	PageFieldWrapCount int32 `json:"PageFieldWrapCount,omitempty" xml:"PageFieldWrapCount"`
	ShowPivotStyleColumnStripes bool `json:"ShowPivotStyleColumnStripes,omitempty" xml:"ShowPivotStyleColumnStripes"`
	FieldListSortAscending bool `json:"FieldListSortAscending,omitempty" xml:"FieldListSortAscending"`
	AltTextTitle string `json:"AltTextTitle,omitempty" xml:"AltTextTitle"`
	PreserveFormatting bool `json:"PreserveFormatting,omitempty" xml:"PreserveFormatting"`
	PivotTableStyleType string `json:"PivotTableStyleType,omitempty" xml:"PivotTableStyleType"`
	DataField *PivotField `json:"DataField,omitempty" xml:"DataField"`
	SaveData bool `json:"SaveData,omitempty" xml:"SaveData"`
	SubtotalHiddenPageItems bool `json:"SubtotalHiddenPageItems,omitempty" xml:"SubtotalHiddenPageItems"`
	PrintDrill bool `json:"PrintDrill,omitempty" xml:"PrintDrill"`
	ShowPivotStyleColumnHeader bool `json:"ShowPivotStyleColumnHeader,omitempty" xml:"ShowPivotStyleColumnHeader"`
	BaseFields []PivotField `json:"BaseFields,omitempty" xml:"BaseFields"`
}
