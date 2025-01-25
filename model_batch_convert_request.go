/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="BatchConvertRequest.go">
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

type BatchConvertRequest struct {
 
    SourceFolder string `json:"SourceFolder,omitempty" xml:"SourceFolder"`
    SourceStorage string `json:"SourceStorage,omitempty" xml:"SourceStorage"`
    MatchCondition *MatchConditionRequest `json:"MatchCondition,omitempty" xml:"MatchCondition"`
    Format string `json:"Format,omitempty" xml:"Format"`
    OutFolder string `json:"OutFolder,omitempty" xml:"OutFolder"`
    OutStorage string `json:"OutStorage,omitempty" xml:"OutStorage"`
    Region string `json:"Region,omitempty" xml:"Region"`
    PageWideFitOnPerSheet bool `json:"PageWideFitOnPerSheet,omitempty" xml:"PageWideFitOnPerSheet"`
    PageTallFitOnPerSheet bool `json:"PageTallFitOnPerSheet,omitempty" xml:"PageTallFitOnPerSheet"`
    SaveOptions *SaveOptions `json:"SaveOptions,omitempty" xml:"SaveOptions"`
}
