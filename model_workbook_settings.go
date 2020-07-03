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

type WorkbookSettings struct {
	NumberGroupSeparator string `json:"NumberGroupSeparator,omitempty" xml:"NumberGroupSeparator"`
	HidePivotFieldList bool `json:"HidePivotFieldList,omitempty" xml:"HidePivotFieldList"`
	// Represents whether the generated spreadsheet will be opened Minimized.             
	IsMinimized bool `json:"IsMinimized,omitempty" xml:"IsMinimized"`
	// Specifies the version of the calculation engine used to calculate values in the workbook.             
	CalculationId string `json:"CalculationId,omitempty" xml:"CalculationId"`
	// Indicates whether re-calculate all formulas on opening file.             
	ReCalculateOnOpen bool `json:"ReCalculateOnOpen,omitempty" xml:"ReCalculateOnOpen"`
	// Whether check restriction of excel file when user modify cells related objects.  For example, excel does not allow inputting string value longer than 32K.  When you input a value longer than 32K such as by Cell.PutValue(string), if this property is true, you will get an Exception.  If this property is false, we will accept your input string value as the cell's value so that later you can output the complete string value for other file formats such as CSV.  However, if you have set such kind of value that is invalid for excel file format, you should not save the workbook as excel file format later. Otherwise there may be unexpected error for the generated excel file.             
	CheckExcelRestriction bool `json:"CheckExcelRestriction,omitempty" xml:"CheckExcelRestriction"`
	// Gets or sets a value indicating whether the generated spreadsheet will contain a horizontal scroll bar.                           Remarks: The default value is true.              
	IsHScrollBarVisible bool `json:"IsHScrollBarVisible,omitempty" xml:"IsHScrollBarVisible"`
	// The height of the window, in unit of point.             
	WindowHeight float64 `json:"WindowHeight,omitempty" xml:"WindowHeight"`
	// The distance from the left edge of the client area to the left edge of the window, in unit of point.             
	WindowLeft float64 `json:"WindowLeft,omitempty" xml:"WindowLeft"`
	// Specifies the stack size for calculating cells recursively.  The large value for this size will give better performance when there are lots of cells need to be calculated recursively.  On the other hand, larger value will raise the stakes of StackOverflowException.  If use gets StackOverflowException when calculating formulas, this value should be decreased.             
	CalcStackSize int64 `json:"CalcStackSize,omitempty" xml:"CalcStackSize"`
	// Gets or sets a value that indicates whether the Workbook is shared.                           Remarks: The default value is false.              
	Shared bool `json:"Shared,omitempty" xml:"Shared"`
	RemovePersonalInformation bool `json:"RemovePersonalInformation,omitempty" xml:"RemovePersonalInformation"`
	// Gets or sets the user interface language of the Workbook version based on CountryCode that has saved the file.             
	LanguageCode string `json:"LanguageCode,omitempty" xml:"LanguageCode"`
	EnableMacros bool `json:"EnableMacros,omitempty" xml:"EnableMacros"`
	IsDefaultEncrypted bool `json:"IsDefaultEncrypted,omitempty" xml:"IsDefaultEncrypted"`
	// Indicates whether to recalculate before saving the document.             
	RecalculateBeforeSave bool `json:"RecalculateBeforeSave,omitempty" xml:"RecalculateBeforeSave"`
	// Indicates whether parsing the formula when reading the file.                           Remarks: Only applies for Excel Xlsx,Xltx, Xltm,Xlsm file because the formulas in the files are stored with a string formula.              
	ParsingFormulaOnOpen bool `json:"ParsingFormulaOnOpen,omitempty" xml:"ParsingFormulaOnOpen"`
	// The distance from the top edge of the client area to the top edge of the window, in unit of point.             
	WindowTop float64 `json:"WindowTop,omitempty" xml:"WindowTop"`
	// Gets or sets the system regional settings based on CountryCode at the time the file was saved.                           Remarks: If you do not want to use the region saved in the file, please reset it after reading the file.              
	Region string `json:"Region,omitempty" xml:"Region"`
	MemorySetting string `json:"MemorySetting,omitempty" xml:"MemorySetting"`
	// Indicates whether update adjacent cells' border.                           Remarks: The default value is true.  For example: the bottom border of the cell A1 is update, the top border of the cell A2 should be changed too.              
	UpdateAdjacentCellsBorder bool `json:"UpdateAdjacentCellsBorder,omitempty" xml:"UpdateAdjacentCellsBorder"`
	CrashSave bool `json:"CrashSave,omitempty" xml:"CrashSave"`
	// Get or sets a value whether the Workbook tabs are displayed.                           Remarks: The default value is true.              
	ShowTabs bool `json:"ShowTabs,omitempty" xml:"ShowTabs"`
	// True if calculations in this workbook will be done using only the precision of the numbers as they're displayed             
	PrecisionAsDisplayed bool `json:"PrecisionAsDisplayed,omitempty" xml:"PrecisionAsDisplayed"`
	// It specifies whether to calculate formulas manually, automatically or automatically except for multiple table operations.             
	CalcMode string `json:"CalcMode,omitempty" xml:"CalcMode"`
	AutoCompressPictures bool `json:"AutoCompressPictures,omitempty" xml:"AutoCompressPictures"`
	// Gets or sets a value which represents if the workbook uses the 1904 date system.             
	Date1904 bool `json:"Date1904,omitempty" xml:"Date1904"`
	NumberDecimalSeparator string `json:"NumberDecimalSeparator,omitempty" xml:"NumberDecimalSeparator"`
	// Indicates if Aspose.Cells will use iteration to resolve circular references.             
	Iteration bool `json:"Iteration,omitempty" xml:"Iteration"`
	// Indicates whether check comptiliblity when saving workbook.                           Remarks:  The default value is true.              
	CheckComptiliblity bool `json:"CheckComptiliblity,omitempty" xml:"CheckComptiliblity"`
	AutoRecover bool `json:"AutoRecover,omitempty" xml:"AutoRecover"`
	// Returns or sets the maximum number of change that Microsoft Excel can use to resolve a circular reference.             
	MaxChange float64 `json:"MaxChange,omitempty" xml:"MaxChange"`
	DataExtractLoad bool `json:"DataExtractLoad,omitempty" xml:"DataExtractLoad"`
	// Gets or sets the first visible worksheet tab.             
	FirstVisibleTab int64 `json:"FirstVisibleTab,omitempty" xml:"FirstVisibleTab"`
	// Indicates whether this workbook is hidden.             
	IsHidden bool `json:"IsHidden,omitempty" xml:"IsHidden"`
	// Indicates if the Read Only Recommended option is selected.             
	RecommendReadOnly bool `json:"RecommendReadOnly,omitempty" xml:"RecommendReadOnly"`
	// Indicates whether and how to show objects in the workbook.             
	DisplayDrawingObjects string `json:"DisplayDrawingObjects,omitempty" xml:"DisplayDrawingObjects"`
	// Specifies the incremental public release of the application.             
	BuildVersion string `json:"BuildVersion,omitempty" xml:"BuildVersion"`
	// Gets or sets a value indicating whether the generated spreadsheet will contain a vertical scroll bar.                           Remarks: The default value is true.              
	IsVScrollBarVisible bool `json:"IsVScrollBarVisible,omitempty" xml:"IsVScrollBarVisible"`
	// The width of the window, in unit of point.             
	WindowWidth float64 `json:"WindowWidth,omitempty" xml:"WindowWidth"`
	// Indicates whether create calculated formulas chain.             
	CreateCalcChain bool `json:"CreateCalcChain,omitempty" xml:"CreateCalcChain"`
	// Returns or sets the maximum number of iterations that Aspose.Cells can use to resolve a circular reference.             
	MaxIteration int64 `json:"MaxIteration,omitempty" xml:"MaxIteration"`
	RepairLoad bool `json:"RepairLoad,omitempty" xml:"RepairLoad"`
	UpdateLinksType string `json:"UpdateLinksType,omitempty" xml:"UpdateLinksType"`
	// Width of worksheet tab bar (in 1/1000 of window width).             
	SheetTabBarWidth int64 `json:"SheetTabBarWidth,omitempty" xml:"SheetTabBarWidth"`
}
