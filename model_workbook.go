/*
 *  Copyright (c) 2021 Aspose.Cells Cloud
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

type Workbook struct {
	Links []Link `json:"Links,omitempty" xml:"Links"`
	Settings *LinkElement `json:"Settings,omitempty" xml:"Settings"`
	IsProtected string `json:"IsProtected,omitempty" xml:"IsProtected"`
	IsWriteProtected string `json:"IsWriteProtected,omitempty" xml:"IsWriteProtected"`
	FileName string `json:"FileName,omitempty" xml:"FileName"`
	DocumentProperties *LinkElement `json:"DocumentProperties,omitempty" xml:"DocumentProperties"`
	Names *LinkElement `json:"Names,omitempty" xml:"Names"`
	DefaultStyle *LinkElement `json:"DefaultStyle,omitempty" xml:"DefaultStyle"`
	IsEncryption string `json:"IsEncryption,omitempty" xml:"IsEncryption"`
	Worksheets *LinkElement `json:"Worksheets,omitempty" xml:"Worksheets"`
	Password string `json:"Password,omitempty" xml:"Password"`
}
