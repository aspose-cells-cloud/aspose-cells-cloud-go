/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="PageSetup.go">
*   Copyright (c) 2023 Aspose.Cells Cloud
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

type PageSetup struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    BlackAndWhite bool `json:"BlackAndWhite,omitempty" xml:"BlackAndWhite"`
    BottomMargin float64 `json:"BottomMargin,omitempty" xml:"BottomMargin"`
    CenterHorizontally bool `json:"CenterHorizontally,omitempty" xml:"CenterHorizontally"`
    CenterVertically bool `json:"CenterVertically,omitempty" xml:"CenterVertically"`
    FirstPageNumber int64 `json:"FirstPageNumber,omitempty" xml:"FirstPageNumber"`
    FitToPagesTall int64 `json:"FitToPagesTall,omitempty" xml:"FitToPagesTall"`
    FitToPagesWide int64 `json:"FitToPagesWide,omitempty" xml:"FitToPagesWide"`
    FooterMargin float64 `json:"FooterMargin,omitempty" xml:"FooterMargin"`
    HeaderMargin float64 `json:"HeaderMargin,omitempty" xml:"HeaderMargin"`
    IsAutoFirstPageNumber bool `json:"IsAutoFirstPageNumber,omitempty" xml:"IsAutoFirstPageNumber"`
    IsHFAlignMargins bool `json:"IsHFAlignMargins,omitempty" xml:"IsHFAlignMargins"`
    IsHFDiffFirst bool `json:"IsHFDiffFirst,omitempty" xml:"IsHFDiffFirst"`
    IsHFDiffOddEven bool `json:"IsHFDiffOddEven,omitempty" xml:"IsHFDiffOddEven"`
    IsHFScaleWithDoc bool `json:"IsHFScaleWithDoc,omitempty" xml:"IsHFScaleWithDoc"`
    IsPercentScale bool `json:"IsPercentScale,omitempty" xml:"IsPercentScale"`
    LeftMargin float64 `json:"LeftMargin,omitempty" xml:"LeftMargin"`
    Order string `json:"Order,omitempty" xml:"Order"`
    Orientation string `json:"Orientation,omitempty" xml:"Orientation"`
    PaperSize string `json:"PaperSize,omitempty" xml:"PaperSize"`
    PrintArea string `json:"PrintArea,omitempty" xml:"PrintArea"`
    PrintComments string `json:"PrintComments,omitempty" xml:"PrintComments"`
    PrintCopies int64 `json:"PrintCopies,omitempty" xml:"PrintCopies"`
    PrintDraft bool `json:"PrintDraft,omitempty" xml:"PrintDraft"`
    PrintErrors string `json:"PrintErrors,omitempty" xml:"PrintErrors"`
    PrintGridlines bool `json:"PrintGridlines,omitempty" xml:"PrintGridlines"`
    PrintHeadings bool `json:"PrintHeadings,omitempty" xml:"PrintHeadings"`
    PrintQuality int64 `json:"PrintQuality,omitempty" xml:"PrintQuality"`
    PrintTitleColumns string `json:"PrintTitleColumns,omitempty" xml:"PrintTitleColumns"`
    PrintTitleRows string `json:"PrintTitleRows,omitempty" xml:"PrintTitleRows"`
    RightMargin float64 `json:"RightMargin,omitempty" xml:"RightMargin"`
    TopMargin float64 `json:"TopMargin,omitempty" xml:"TopMargin"`
    Zoom int64 `json:"Zoom,omitempty" xml:"Zoom"`
    Header []PageSection `json:"Header,omitempty" xml:"Header"`
    Footer []PageSection `json:"Footer,omitempty" xml:"Footer"`
}
