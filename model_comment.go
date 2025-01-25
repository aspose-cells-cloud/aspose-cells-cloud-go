/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="Comment.go">
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

type Comment struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    CellName string `json:"CellName,omitempty" xml:"CellName"`
    Author string `json:"Author,omitempty" xml:"Author"`
    HtmlNote string `json:"HtmlNote,omitempty" xml:"HtmlNote"`
    Note string `json:"Note,omitempty" xml:"Note"`
    AutoSize bool `json:"AutoSize,omitempty" xml:"AutoSize"`
    IsVisible bool `json:"IsVisible,omitempty" xml:"IsVisible"`
    Width int64 `json:"Width,omitempty" xml:"Width"`
    Height int64 `json:"Height,omitempty" xml:"Height"`
    TextHorizontalAlignment string `json:"TextHorizontalAlignment,omitempty" xml:"TextHorizontalAlignment"`
    TextOrientationType string `json:"TextOrientationType,omitempty" xml:"TextOrientationType"`
    TextVerticalAlignment string `json:"TextVerticalAlignment,omitempty" xml:"TextVerticalAlignment"`
}
