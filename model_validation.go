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

type Validation struct {
	Link *Link `json:"link,omitempty" xml:"link"`
	// Represents the value or expression associated with the second part of the    data validation.             
	Formula2 string `json:"Formula2,omitempty" xml:"Formula2"`
	// Represents the value or expression associated with the data validation.
	Formula1 string `json:"Formula1,omitempty" xml:"Formula1"`
	// Indicates whether the data validation error message will be displayed whenever    the user enters invalid data.
	ShowError bool `json:"ShowError,omitempty" xml:"ShowError"`
	// Represents the data validation error message.
	ErrorMessage string `json:"ErrorMessage,omitempty" xml:"ErrorMessage"`
	// Indicates whether data validation displays a drop-down list that contains    acceptable values.
	InCellDropDown bool `json:"InCellDropDown,omitempty" xml:"InCellDropDown"`
	// Indicates whether the data validation input message will be displayed whenever    the user selects a cell in the data validation range.
	ShowInput bool `json:"ShowInput,omitempty" xml:"ShowInput"`
	// Represents the validation alert style.Information,Stop,Warning             
	AlertStyle string `json:"AlertStyle,omitempty" xml:"AlertStyle"`
	// Represents the title of the data-validation input dialog box.
	InputTitle string `json:"InputTitle,omitempty" xml:"InputTitle"`
	// Indicates whether blank values are permitted by the range data validation.
	IgnoreBlank bool `json:"IgnoreBlank,omitempty" xml:"IgnoreBlank"`
	// Represents the first value associated with the data validation.             
	Value2 string `json:"Value2,omitempty" xml:"Value2"`
	// Represents the first value associated with the data validation.
	Value1 string `json:"Value1,omitempty" xml:"Value1"`
	// Represents the operator for the data validation. Between,Equal,GreaterThan,GreaterOrEqual,LessThan,LessOrEqual,None,NotBetween,NotEqual
	Operator string `json:"Operator,omitempty" xml:"Operator"`
	// Represents the title of the data-validation error dialog box.
	ErrorTitle string `json:"ErrorTitle,omitempty" xml:"ErrorTitle"`
	// Represents the data validation type. AnyValue ,WholeNumber,Decimal,List,Date,Time,TextLength,Custom             
	Type_ string `json:"Type,omitempty" xml:"Type"`
	// Represents the data validation input message.
	InputMessage string `json:"InputMessage,omitempty" xml:"InputMessage"`
	// Represents a collection of Aspose.Cells.CellArea which contains the data     validation settings.
	AreaList []CellArea `json:"AreaList,omitempty" xml:"AreaList"`
}
