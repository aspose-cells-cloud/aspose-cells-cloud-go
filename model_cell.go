/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="Cell.go">
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

type Cell struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    Name string `json:"Name,omitempty" xml:"Name"`
    Row int64 `json:"Row,omitempty" xml:"Row"`
    Column int64 `json:"Column,omitempty" xml:"Column"`
    Value string `json:"Value,omitempty" xml:"Value"`
    Type_ string `json:"Type,omitempty" xml:"Type"`
    Formula string `json:"Formula,omitempty" xml:"Formula"`
    IsFormula bool `json:"IsFormula,omitempty" xml:"IsFormula"`
    IsMerged bool `json:"IsMerged,omitempty" xml:"IsMerged"`
    IsArrayHeader bool `json:"IsArrayHeader,omitempty" xml:"IsArrayHeader"`
    IsInArray bool `json:"IsInArray,omitempty" xml:"IsInArray"`
    IsErrorValue bool `json:"IsErrorValue,omitempty" xml:"IsErrorValue"`
    IsInTable bool `json:"IsInTable,omitempty" xml:"IsInTable"`
    IsStyleSet bool `json:"IsStyleSet,omitempty" xml:"IsStyleSet"`
    HtmlString string `json:"HtmlString,omitempty" xml:"HtmlString"`
    Style *LinkElement `json:"Style,omitempty" xml:"Style"`
    Worksheet string `json:"Worksheet,omitempty" xml:"Worksheet"`
}
