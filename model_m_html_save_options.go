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

type MHtmlSaveOptions struct {
	EnableHTTPCompression bool `json:"EnableHTTPCompression,omitempty" xml:"EnableHTTPCompression"`
	SaveFormat string `json:"SaveFormat,omitempty" xml:"SaveFormat"`
	// Make the workbook empty after saving the file.
	ClearData bool `json:"ClearData,omitempty" xml:"ClearData"`
	// The cached file folder is used to store some large data.
	CachedFileFolder string `json:"CachedFileFolder,omitempty" xml:"CachedFileFolder"`
	// Indicates whether validate merged areas before saving the file. The default value is false.             
	ValidateMergedAreas bool `json:"ValidateMergedAreas,omitempty" xml:"ValidateMergedAreas"`
	RefreshChartCache bool `json:"RefreshChartCache,omitempty" xml:"RefreshChartCache"`
	// If true and the directory does not exist, the directory will be automatically created before saving the file.             
	CreateDirectory bool `json:"CreateDirectory,omitempty" xml:"CreateDirectory"`
	SortNames bool `json:"SortNames,omitempty" xml:"SortNames"`
	AttachedFilesDirectory string `json:"AttachedFilesDirectory,omitempty" xml:"AttachedFilesDirectory"`
	AttachedFilesUrlPrefix string `json:"AttachedFilesUrlPrefix,omitempty" xml:"AttachedFilesUrlPrefix"`
	Encoding string `json:"Encoding,omitempty" xml:"Encoding"`
	ExportActiveWorksheetOnly bool `json:"ExportActiveWorksheetOnly,omitempty" xml:"ExportActiveWorksheetOnly"`
	ExportChartImageFormat string `json:"ExportChartImageFormat,omitempty" xml:"ExportChartImageFormat"`
	ExportImagesAsBase64 bool `json:"ExportImagesAsBase64,omitempty" xml:"ExportImagesAsBase64"`
	HiddenColDisplayType string `json:"HiddenColDisplayType,omitempty" xml:"HiddenColDisplayType"`
	HiddenRowDisplayType string `json:"HiddenRowDisplayType,omitempty" xml:"HiddenRowDisplayType"`
	HtmlCrossStringType string `json:"HtmlCrossStringType,omitempty" xml:"HtmlCrossStringType"`
	IsExpImageToTempDir bool `json:"IsExpImageToTempDir,omitempty" xml:"IsExpImageToTempDir"`
	PageTitle string `json:"PageTitle,omitempty" xml:"PageTitle"`
	ParseHtmlTagInCell bool `json:"ParseHtmlTagInCell,omitempty" xml:"ParseHtmlTagInCell"`
}
