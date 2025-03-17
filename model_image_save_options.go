/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="ImageSaveOptions.go">
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

type ImageSaveOptions struct {
     
        SaveFormat string `json:"SaveFormat,omitempty" xml:"SaveFormat"`
        CachedFileFolder string `json:"CachedFileFolder,omitempty" xml:"CachedFileFolder"`
        ClearData bool `json:"ClearData,omitempty" xml:"ClearData"`
        CreateDirectory bool `json:"CreateDirectory,omitempty" xml:"CreateDirectory"`
        EnableHTTPCompression bool `json:"EnableHTTPCompression,omitempty" xml:"EnableHTTPCompression"`
        RefreshChartCache bool `json:"RefreshChartCache,omitempty" xml:"RefreshChartCache"`
        SortNames bool `json:"SortNames,omitempty" xml:"SortNames"`
        ValidateMergedAreas bool `json:"ValidateMergedAreas,omitempty" xml:"ValidateMergedAreas"`
        MergeAreas bool `json:"MergeAreas,omitempty" xml:"MergeAreas"`
        SortExternalNames bool `json:"SortExternalNames,omitempty" xml:"SortExternalNames"`
        CheckExcelRestriction bool `json:"CheckExcelRestriction,omitempty" xml:"CheckExcelRestriction"`
        UpdateSmartArt bool `json:"UpdateSmartArt,omitempty" xml:"UpdateSmartArt"`
        EncryptDocumentProperties bool `json:"EncryptDocumentProperties,omitempty" xml:"EncryptDocumentProperties"`
 
    ChartImageType string `json:"ChartImageType,omitempty" xml:"ChartImageType"`
    EmbededImageNameInSvg string `json:"EmbededImageNameInSvg,omitempty" xml:"EmbededImageNameInSvg"`
    HorizontalResolution int64 `json:"HorizontalResolution,omitempty" xml:"HorizontalResolution"`
    ImageFormat string `json:"ImageFormat,omitempty" xml:"ImageFormat"`
    IsCellAutoFit bool `json:"IsCellAutoFit,omitempty" xml:"IsCellAutoFit"`
    OnePagePerSheet bool `json:"OnePagePerSheet,omitempty" xml:"OnePagePerSheet"`
    OnlyArea bool `json:"OnlyArea,omitempty" xml:"OnlyArea"`
    PrintingPage string `json:"PrintingPage,omitempty" xml:"PrintingPage"`
    PrintWithStatusDialog bool `json:"PrintWithStatusDialog,omitempty" xml:"PrintWithStatusDialog"`
    Quality int64 `json:"Quality,omitempty" xml:"Quality"`
    TiffCompression string `json:"TiffCompression,omitempty" xml:"TiffCompression"`
    VerticalResolution int64 `json:"VerticalResolution,omitempty" xml:"VerticalResolution"`
}
