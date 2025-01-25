/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="LoadOptions.go">
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

type LoadOptions struct {
 
    ConvertNumericData string `json:"ConvertNumericData,omitempty" xml:"ConvertNumericData"`
    InterruptMonitor string `json:"InterruptMonitor,omitempty" xml:"InterruptMonitor"`
    LanguageCode string `json:"LanguageCode,omitempty" xml:"LanguageCode"`
    LoadDataOptions string `json:"LoadDataOptions,omitempty" xml:"LoadDataOptions"`
    LoadFormat string `json:"LoadFormat,omitempty" xml:"LoadFormat"`
    OnlyLoadDocumentProperties string `json:"OnlyLoadDocumentProperties,omitempty" xml:"OnlyLoadDocumentProperties"`
    ParsingFormulaOnOpen string `json:"ParsingFormulaOnOpen,omitempty" xml:"ParsingFormulaOnOpen"`
    Password string `json:"Password,omitempty" xml:"Password"`
    Region string `json:"Region,omitempty" xml:"Region"`
    StandardFont string `json:"StandardFont,omitempty" xml:"StandardFont"`
    StandardFontSize float64 `json:"StandardFontSize,omitempty" xml:"StandardFontSize"`
}
