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

// Describe the DataBar conditional formatting rule. This conditional formatting    rule displays a gradated data bar in the range of cells.
type DataBar struct {
	// Gets or sets the direction the databar is displayed.
	Direction string `json:"Direction,omitempty" xml:"Direction"`
	// Get or set this DataBar's max value object.  Cannot set null or CFValueObject    with type FormatConditionValueType.Min to it.             
	MaxCfvo *ConditionalFormattingValue `json:"MaxCfvo,omitempty" xml:"MaxCfvo"`
	// Get or set this DataBar's Color.             
	Color *Color `json:"Color,omitempty" xml:"Color"`
	// Represents the min length of data bar .             
	MinLength int64 `json:"MinLength,omitempty" xml:"MinLength"`
	// Gets or sets how a data bar is filled with color.
	BarFillType string `json:"BarFillType,omitempty" xml:"BarFillType"`
	// Get or set this DataBar's min value object.  Cannot set null or CFValueObject   with type FormatConditionValueType.Max to it.             
	MinCfvo *ConditionalFormattingValue `json:"MinCfvo,omitempty" xml:"MinCfvo"`
	// Gets or sets the position of the axis of the data bars specified by a conditional    formatting rule.
	AxisPosition string `json:"AxisPosition,omitempty" xml:"AxisPosition"`
	// Gets the NegativeBarFormat object associated with a data bar conditional     formatting rule.
	NegativeBarFormat *NegativeBarFormat `json:"NegativeBarFormat,omitempty" xml:"NegativeBarFormat"`
	// Gets an object that specifies the border of a data bar.
	BarBorder *DataBarBorder `json:"BarBorder,omitempty" xml:"BarBorder"`
	// Gets the color of the axis for cells with conditional formatting as data bars.
	AxisColor *Color `json:"AxisColor,omitempty" xml:"AxisColor"`
	// Represents the max length of data bar .
	MaxLength int64 `json:"MaxLength,omitempty" xml:"MaxLength"`
	// Get or set the flag indicating whether to show the values of the cells on   which this data bar is applied.  Default value is true.             
	ShowValue bool `json:"ShowValue,omitempty" xml:"ShowValue"`
}
