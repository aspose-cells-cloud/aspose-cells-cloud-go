/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="ConvertWorksheetTaskParameter.go">
*   Copyright (c) 2024 Aspose.Cells Cloud
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

type ConvertWorksheetTaskParameter struct {
     
 
    DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
    Workbook *FileSource `json:"Workbook,omitempty" xml:"Workbook"`
    Sheet string `json:"Sheet,omitempty" xml:"Sheet"`
    TargetDataSource *DataSource `json:"TargetDataSource,omitempty" xml:"TargetDataSource"`
    Target *FileSource `json:"Target,omitempty" xml:"Target"`
    Format string `json:"Format,omitempty" xml:"Format"`
    Area string `json:"Area,omitempty" xml:"Area"`
    PageIndex int64 `json:"PageIndex,omitempty" xml:"PageIndex"`
    VerticalResolution int64 `json:"VerticalResolution,omitempty" xml:"VerticalResolution"`
    HorizontalResolution int64 `json:"HorizontalResolution,omitempty" xml:"HorizontalResolution"`
}
