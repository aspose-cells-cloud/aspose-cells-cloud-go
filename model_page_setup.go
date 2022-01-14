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

// excel print page setting
type PageSetup struct {
	Link *Link `json:"link,omitempty" xml:"link"`
	// True means that the header/footer of the first page is different with other pages.
	IsHFDiffFirst bool `json:"IsHFDiffFirst,omitempty" xml:"IsHFDiffFirst"`
	// Represents the number of pages wide the worksheet will be scaled to when it's printed.
	FitToPagesWide int64 `json:"FitToPagesWide,omitempty" xml:"FitToPagesWide"`
	// Represents the print quality.
	PrintQuality int64 `json:"PrintQuality,omitempty" xml:"PrintQuality"`
	// Represents if the sheet will be printed without graphics.
	PrintDraft bool `json:"PrintDraft,omitempty" xml:"PrintDraft"`
	// Represents the first page number that will be used when this sheet is printed.
	FirstPageNumber int64 `json:"FirstPageNumber,omitempty" xml:"FirstPageNumber"`
	// Represents the size of the paper.
	PaperSize string `json:"PaperSize,omitempty" xml:"PaperSize"`
	// Represents the way comments are printed with the sheet.
	PrintComments string `json:"PrintComments,omitempty" xml:"PrintComments"`
	// Specifies the type of print error displayed.
	PrintErrors string `json:"PrintErrors,omitempty" xml:"PrintErrors"`
	// Represent if the sheet is printed centered vertically.
	CenterVertically bool `json:"CenterVertically,omitempty" xml:"CenterVertically"`
	// If this property is False, the FitToPagesWide and FitToPagesTall properties control how the worksheet is scaled.
	IsPercentScale bool `json:"IsPercentScale,omitempty" xml:"IsPercentScale"`
	// Represents if elements of the document will be printed in black and white. True/False
	BlackAndWhite bool `json:"BlackAndWhite,omitempty" xml:"BlackAndWhite"`
	// Represents the columns that contain the cells to be repeated on the left side of each page.
	PrintTitleColumns string `json:"PrintTitleColumns,omitempty" xml:"PrintTitleColumns"`
	// Indicates whether header and footer margins are aligned with the page margins.Only applies for Excel 2007.
	IsHFAlignMargins bool `json:"IsHFAlignMargins,omitempty" xml:"IsHFAlignMargins"`
	// Represents the range to be printed.
	PrintArea string `json:"PrintArea,omitempty" xml:"PrintArea"`
	// Represents the distance from the bottom of the page to the footer, in unit of centimeters.
	FooterMargin float64 `json:"FooterMargin,omitempty" xml:"FooterMargin"`
	// Represents the size of the left margin, in unit of centimeters.
	LeftMargin float64 `json:"LeftMargin,omitempty" xml:"LeftMargin"`
	// Represent if the sheet is printed centered horizontally.
	CenterHorizontally bool `json:"CenterHorizontally,omitempty" xml:"CenterHorizontally"`
	// Represents the distance from the top of the page to the header, in unit of centimeters.
	HeaderMargin float64 `json:"HeaderMargin,omitempty" xml:"HeaderMargin"`
	// Represents the size of the top margin, in unit of centimeters.
	TopMargin float64 `json:"TopMargin,omitempty" xml:"TopMargin"`
	// Represents the page footor.
	Footer []PageSection `json:"Footer,omitempty" xml:"Footer"`
	// Represents the number of pages tall the worksheet will be scaled to when it's printed.
	FitToPagesTall int64 `json:"FitToPagesTall,omitempty" xml:"FitToPagesTall"`
	// Indicates whether header and footer are scaled with document scaling.Only applies for Excel 2007. 
	IsHFScaleWithDoc bool `json:"IsHFScaleWithDoc,omitempty" xml:"IsHFScaleWithDoc"`
	// Represents if row and column headings are printed with this page.
	PrintHeadings bool `json:"PrintHeadings,omitempty" xml:"PrintHeadings"`
	// Represents the scaling factor in percent. It should be between 10 and 400.
	Zoom int64 `json:"Zoom,omitempty" xml:"Zoom"`
	// Represents the rows that contain the cells to be repeated at the top of each page.
	PrintTitleRows string `json:"PrintTitleRows,omitempty" xml:"PrintTitleRows"`
	// Represents the order that Microsoft Excel uses to number pages when printing a large worksheet.
	Order string `json:"Order,omitempty" xml:"Order"`
	// Get and sets number of copies to print.
	PrintCopies int64 `json:"PrintCopies,omitempty" xml:"PrintCopies"`
	// Represents page print orientation.
	Orientation string `json:"Orientation,omitempty" xml:"Orientation"`
	// Represents the size of the right margin, in unit of centimeters.
	RightMargin float64 `json:"RightMargin,omitempty" xml:"RightMargin"`
	// Represents if cell gridlines are printed on the page.
	PrintGridlines bool `json:"PrintGridlines,omitempty" xml:"PrintGridlines"`
	// Indicates whether the first the page number is automatically assigned.
	IsAutoFirstPageNumber bool `json:"IsAutoFirstPageNumber,omitempty" xml:"IsAutoFirstPageNumber"`
	// Represents the page header.
	Header []PageSection `json:"Header,omitempty" xml:"Header"`
	// True means that the header/footer of the odd pages is different with odd pages.
	IsHFDiffOddEven bool `json:"IsHFDiffOddEven,omitempty" xml:"IsHFDiffOddEven"`
	// Represents the size of the bottom margin, in unit of centimeters.
	BottomMargin float64 `json:"BottomMargin,omitempty" xml:"BottomMargin"`
}
