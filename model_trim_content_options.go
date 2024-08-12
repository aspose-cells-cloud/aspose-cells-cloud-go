/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="TrimContentOptions.go">
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

type TrimContentOptions struct {
 
    DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
    FileInfo *FileInfo `json:"FileInfo,omitempty" xml:"FileInfo"`
    TrimContent string `json:"TrimContent,omitempty" xml:"TrimContent"`
    TrimLeading bool `json:"TrimLeading,omitempty" xml:"TrimLeading"`
    TrimTrailing bool `json:"TrimTrailing,omitempty" xml:"TrimTrailing"`
    TrimSpaceBetweenWordTo1 bool `json:"TrimSpaceBetweenWordTo1,omitempty" xml:"TrimSpaceBetweenWordTo1"`
    TrimNonBreakingSpaces bool `json:"TrimNonBreakingSpaces,omitempty" xml:"TrimNonBreakingSpaces"`
    RemoveExtraLineBreaks bool `json:"RemoveExtraLineBreaks,omitempty" xml:"RemoveExtraLineBreaks"`
    RemoveAllLineBreaks bool `json:"RemoveAllLineBreaks,omitempty" xml:"RemoveAllLineBreaks"`
    ScopeOptions *ScopeOptions `json:"ScopeOptions,omitempty" xml:"ScopeOptions"`
}
