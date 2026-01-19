/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="WorkbookSettings.go">
*   Copyright (c) 2026 Aspose.Cells Cloud
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

type WorkbookSettings struct {
 
    AutoCompressPictures bool `json:"AutoCompressPictures,omitempty" xml:"AutoCompressPictures"`
    AutoRecover bool `json:"AutoRecover,omitempty" xml:"AutoRecover"`
    BuildVersion string `json:"BuildVersion,omitempty" xml:"BuildVersion"`
    CalcMode string `json:"CalcMode,omitempty" xml:"CalcMode"`
    CalculationId string `json:"CalculationId,omitempty" xml:"CalculationId"`
    CheckComptiliblity bool `json:"CheckComptiliblity,omitempty" xml:"CheckComptiliblity"`
    CheckExcelRestriction bool `json:"CheckExcelRestriction,omitempty" xml:"CheckExcelRestriction"`
    CrashSave bool `json:"CrashSave,omitempty" xml:"CrashSave"`
    CreateCalcChain bool `json:"CreateCalcChain,omitempty" xml:"CreateCalcChain"`
    DataExtractLoad bool `json:"DataExtractLoad,omitempty" xml:"DataExtractLoad"`
    Date1904 bool `json:"Date1904,omitempty" xml:"Date1904"`
    DisplayDrawingObjects string `json:"DisplayDrawingObjects,omitempty" xml:"DisplayDrawingObjects"`
    EnableMacros bool `json:"EnableMacros,omitempty" xml:"EnableMacros"`
    FirstVisibleTab int64 `json:"FirstVisibleTab,omitempty" xml:"FirstVisibleTab"`
    HidePivotFieldList bool `json:"HidePivotFieldList,omitempty" xml:"HidePivotFieldList"`
    IsDefaultEncrypted bool `json:"IsDefaultEncrypted,omitempty" xml:"IsDefaultEncrypted"`
    IsHidden bool `json:"IsHidden,omitempty" xml:"IsHidden"`
    IsHScrollBarVisible bool `json:"IsHScrollBarVisible,omitempty" xml:"IsHScrollBarVisible"`
    IsMinimized bool `json:"IsMinimized,omitempty" xml:"IsMinimized"`
    IsVScrollBarVisible bool `json:"IsVScrollBarVisible,omitempty" xml:"IsVScrollBarVisible"`
    Iteration bool `json:"Iteration,omitempty" xml:"Iteration"`
    LanguageCode string `json:"LanguageCode,omitempty" xml:"LanguageCode"`
    MaxChange float64 `json:"MaxChange,omitempty" xml:"MaxChange"`
    MaxIteration int64 `json:"MaxIteration,omitempty" xml:"MaxIteration"`
    MemorySetting string `json:"MemorySetting,omitempty" xml:"MemorySetting"`
    NumberDecimalSeparator string `json:"NumberDecimalSeparator,omitempty" xml:"NumberDecimalSeparator"`
    NumberGroupSeparator string `json:"NumberGroupSeparator,omitempty" xml:"NumberGroupSeparator"`
    ParsingFormulaOnOpen bool `json:"ParsingFormulaOnOpen,omitempty" xml:"ParsingFormulaOnOpen"`
    PrecisionAsDisplayed bool `json:"PrecisionAsDisplayed,omitempty" xml:"PrecisionAsDisplayed"`
    RecalculateBeforeSave bool `json:"RecalculateBeforeSave,omitempty" xml:"RecalculateBeforeSave"`
    ReCalculateOnOpen bool `json:"ReCalculateOnOpen,omitempty" xml:"ReCalculateOnOpen"`
    RecommendReadOnly bool `json:"RecommendReadOnly,omitempty" xml:"RecommendReadOnly"`
    Region string `json:"Region,omitempty" xml:"Region"`
    RemovePersonalInformation bool `json:"RemovePersonalInformation,omitempty" xml:"RemovePersonalInformation"`
    RepairLoad bool `json:"RepairLoad,omitempty" xml:"RepairLoad"`
    Shared bool `json:"Shared,omitempty" xml:"Shared"`
    SheetTabBarWidth int64 `json:"SheetTabBarWidth,omitempty" xml:"SheetTabBarWidth"`
    ShowTabs bool `json:"ShowTabs,omitempty" xml:"ShowTabs"`
    UpdateAdjacentCellsBorder bool `json:"UpdateAdjacentCellsBorder,omitempty" xml:"UpdateAdjacentCellsBorder"`
    UpdateLinksType string `json:"UpdateLinksType,omitempty" xml:"UpdateLinksType"`
    WindowHeight float64 `json:"WindowHeight,omitempty" xml:"WindowHeight"`
    WindowLeft float64 `json:"WindowLeft,omitempty" xml:"WindowLeft"`
    WindowTop float64 `json:"WindowTop,omitempty" xml:"WindowTop"`
    WindowWidth float64 `json:"WindowWidth,omitempty" xml:"WindowWidth"`
    Author string `json:"Author,omitempty" xml:"Author"`
    CheckCustomNumberFormat bool `json:"CheckCustomNumberFormat,omitempty" xml:"CheckCustomNumberFormat"`
    ProtectionType string `json:"ProtectionType,omitempty" xml:"ProtectionType"`
    GlobalizationSettings *GlobalizationSettings `json:"GlobalizationSettings,omitempty" xml:"GlobalizationSettings"`
    Password string `json:"Password,omitempty" xml:"Password"`
    WriteProtection *WriteProtection `json:"WriteProtection,omitempty" xml:"WriteProtection"`
    IsEncrypted bool `json:"IsEncrypted,omitempty" xml:"IsEncrypted"`
    IsProtected bool `json:"IsProtected,omitempty" xml:"IsProtected"`
    MaxRow int64 `json:"MaxRow,omitempty" xml:"MaxRow"`
    MaxColumn int64 `json:"MaxColumn,omitempty" xml:"MaxColumn"`
    SignificantDigits int64 `json:"SignificantDigits,omitempty" xml:"SignificantDigits"`
    CheckCompatibility bool `json:"CheckCompatibility,omitempty" xml:"CheckCompatibility"`
    PaperSize string `json:"PaperSize,omitempty" xml:"PaperSize"`
    MaxRowsOfSharedFormula int64 `json:"MaxRowsOfSharedFormula,omitempty" xml:"MaxRowsOfSharedFormula"`
    Compliance string `json:"Compliance,omitempty" xml:"Compliance"`
    QuotePrefixToStyle bool `json:"QuotePrefixToStyle,omitempty" xml:"QuotePrefixToStyle"`
    FormulaSettings *FormulaSettings `json:"FormulaSettings,omitempty" xml:"FormulaSettings"`
    ForceFullCalculate bool `json:"ForceFullCalculate,omitempty" xml:"ForceFullCalculate"`
}
