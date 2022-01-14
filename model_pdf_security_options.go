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

type PdfSecurityOptions struct {
	AnnotationsPermission bool `json:"AnnotationsPermission,omitempty" xml:"AnnotationsPermission"`
	AssembleDocumentPermission bool `json:"AssembleDocumentPermission,omitempty" xml:"AssembleDocumentPermission"`
	// Make the workbook empty after saving the file.
	ExtractContentPermission bool `json:"ExtractContentPermission,omitempty" xml:"ExtractContentPermission"`
	// The cached file folder is used to store some large data.
	ExtractContentPermissionObsolete bool `json:"ExtractContentPermissionObsolete,omitempty" xml:"ExtractContentPermissionObsolete"`
	// Indicates whether validate merged areas before saving the file. The default value is false.             
	FillFormsPermission bool `json:"FillFormsPermission,omitempty" xml:"FillFormsPermission"`
	FullQualityPrintPermission bool `json:"FullQualityPrintPermission,omitempty" xml:"FullQualityPrintPermission"`
	// If true and the directory does not exist, the directory will be automatically created before saving the file.             
	ModifyDocumentPermission bool `json:"ModifyDocumentPermission,omitempty" xml:"ModifyDocumentPermission"`
	OwnerPassword string `json:"OwnerPassword,omitempty" xml:"OwnerPassword"`
	PrintPermission bool `json:"PrintPermission,omitempty" xml:"PrintPermission"`
	UserPassword string `json:"UserPassword,omitempty" xml:"UserPassword"`
}
