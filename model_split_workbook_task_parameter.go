/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="SplitWorkbookTaskParameter.go">
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

type SplitWorkbookTaskParameter struct {
     
 
    Workbook *FileSource `json:"Workbook,omitempty" xml:"Workbook"`
    DestinationFilePosition *FileSource `json:"DestinationFilePosition,omitempty" xml:"DestinationFilePosition"`
    DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
    TargetDataSource *DataSource `json:"TargetDataSource,omitempty" xml:"TargetDataSource"`
    DestinationFileFormat string `json:"DestinationFileFormat,omitempty" xml:"DestinationFileFormat"`
    SplitNameRule string `json:"SplitNameRule,omitempty" xml:"SplitNameRule"`
    VerticalResolution int64 `json:"VerticalResolution,omitempty" xml:"VerticalResolution"`
    HorizontalResolution int64 `json:"HorizontalResolution,omitempty" xml:"HorizontalResolution"`
}
