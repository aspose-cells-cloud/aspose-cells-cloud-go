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

// Encapsulates the object that represents a single Workbook cell.
type Cell struct {
	Link *Link `json:"link,omitempty" xml:"link"`
	Style *LinkElement `json:"Style,omitempty" xml:"Style"`
	// Gets and sets the html string which contains data and some formattings in this cell.             
	HtmlString string `json:"HtmlString,omitempty" xml:"HtmlString"`
	// Gets the name of the cell.             
	Name string `json:"Name,omitempty" xml:"Name"`
	// Gets column number (zero based) of the cell.             
	Column int64 `json:"Column" xml:"Column"`
	// Gets the parent worksheet.
	Worksheet string `json:"Worksheet,omitempty" xml:"Worksheet"`
	// Indicates whethe this cell is part of table formula.             
	IsInTable bool `json:"IsInTable" xml:"IsInTable"`
	// Inidicates the cell's formula is and array formula and it is the first cell of the array.
	IsArrayHeader bool `json:"IsArrayHeader" xml:"IsArrayHeader"`
	Value string `json:"Value,omitempty" xml:"Value"`
	// Represents if the specified cell contains formula.             
	IsFormula bool `json:"IsFormula" xml:"IsFormula"`
	// Indicates if the cell's style is set. If return false, it means this cell has a default cell format.             
	IsStyleSet bool `json:"IsStyleSet" xml:"IsStyleSet"`
	// Indicates whether the cell formula is an array formula.
	IsInArray bool `json:"IsInArray" xml:"IsInArray"`
	// Checks if a formula can properly evaluate a result.             
	IsErrorValue bool `json:"IsErrorValue" xml:"IsErrorValue"`
	// Checks if a cell is part of a merged range or not.             
	IsMerged bool `json:"IsMerged" xml:"IsMerged"`
	// Gets or sets a formula of the Aspose.Cells.Cell.
	Formula string `json:"Formula,omitempty" xml:"Formula"`
	// Specifies a cell value type.
	Type_ string `json:"Type,omitempty" xml:"Type"`
	// Gets row number (zero based) of the cell.             
	Row int64 `json:"Row" xml:"Row"`
}
