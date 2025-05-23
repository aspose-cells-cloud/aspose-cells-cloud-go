/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="DocxSaveOptions.go">
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

type DocxSaveOptions struct {
     
        DefaultFont string `json:"DefaultFont,omitempty" xml:"DefaultFont"`
        CheckWorkbookDefaultFont bool `json:"CheckWorkbookDefaultFont,omitempty" xml:"CheckWorkbookDefaultFont"`
        CheckFontCompatibility bool `json:"CheckFontCompatibility,omitempty" xml:"CheckFontCompatibility"`
        IsFontSubstitutionCharGranularity bool `json:"IsFontSubstitutionCharGranularity,omitempty" xml:"IsFontSubstitutionCharGranularity"`
        OnePagePerSheet bool `json:"OnePagePerSheet,omitempty" xml:"OnePagePerSheet"`
        AllColumnsInOnePagePerSheet bool `json:"AllColumnsInOnePagePerSheet,omitempty" xml:"AllColumnsInOnePagePerSheet"`
        IgnoreError bool `json:"IgnoreError,omitempty" xml:"IgnoreError"`
        OutputBlankPageWhenNothingToPrint bool `json:"OutputBlankPageWhenNothingToPrint,omitempty" xml:"OutputBlankPageWhenNothingToPrint"`
        PageIndex int64 `json:"PageIndex,omitempty" xml:"PageIndex"`
        PageCount int64 `json:"PageCount,omitempty" xml:"PageCount"`
        PrintingPageType string `json:"PrintingPageType,omitempty" xml:"PrintingPageType"`
        GridlineType string `json:"GridlineType,omitempty" xml:"GridlineType"`
        TextCrossType string `json:"TextCrossType,omitempty" xml:"TextCrossType"`
        DefaultEditLanguage string `json:"DefaultEditLanguage,omitempty" xml:"DefaultEditLanguage"`
        EmfRenderSetting string `json:"EmfRenderSetting,omitempty" xml:"EmfRenderSetting"`
 
}
