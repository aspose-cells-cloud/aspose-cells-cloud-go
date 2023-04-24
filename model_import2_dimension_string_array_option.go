/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="Import2DimensionStringArrayOption.go">
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

type Import2DimensionStringArrayOption struct {
     
        DestinationWorksheet string `json:"DestinationWorksheet,omitempty" xml:"DestinationWorksheet"`
        IsInsert bool `json:"IsInsert,omitempty" xml:"IsInsert"`
        ImportDataType string `json:"ImportDataType,omitempty" xml:"ImportDataType"`
        Source *FileSource `json:"Source,omitempty" xml:"Source"`
 
    FirstRow int64 `json:"FirstRow,omitempty" xml:"FirstRow"`
    FirstColumn int64 `json:"FirstColumn,omitempty" xml:"FirstColumn"`
    Data []string `json:"Data,omitempty" xml:"Data"`
}
