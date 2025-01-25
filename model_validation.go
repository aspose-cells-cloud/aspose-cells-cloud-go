/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="Validation.go">
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

type Validation struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    AlertStyle string `json:"AlertStyle,omitempty" xml:"AlertStyle"`
    AreaList []CellArea `json:"AreaList,omitempty" xml:"AreaList"`
    ErrorMessage string `json:"ErrorMessage,omitempty" xml:"ErrorMessage"`
    ErrorTitle string `json:"ErrorTitle,omitempty" xml:"ErrorTitle"`
    Formula1 string `json:"Formula1,omitempty" xml:"Formula1"`
    Formula2 string `json:"Formula2,omitempty" xml:"Formula2"`
    IgnoreBlank bool `json:"IgnoreBlank,omitempty" xml:"IgnoreBlank"`
    InCellDropDown bool `json:"InCellDropDown,omitempty" xml:"InCellDropDown"`
    InputMessage string `json:"InputMessage,omitempty" xml:"InputMessage"`
    InputTitle string `json:"InputTitle,omitempty" xml:"InputTitle"`
    Operator string `json:"Operator,omitempty" xml:"Operator"`
    ShowError bool `json:"ShowError,omitempty" xml:"ShowError"`
    ShowInput bool `json:"ShowInput,omitempty" xml:"ShowInput"`
    Type_ string `json:"Type,omitempty" xml:"Type"`
    Value1 string `json:"Value1,omitempty" xml:"Value1"`
    Value2 string `json:"Value2,omitempty" xml:"Value2"`
}
