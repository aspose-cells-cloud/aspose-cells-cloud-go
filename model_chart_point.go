/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="ChartPoint.go">
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

type ChartPoint struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    Area *Area `json:"Area,omitempty" xml:"Area"`
    Border *Line `json:"Border,omitempty" xml:"Border"`
    DataLabels *DataLabels `json:"DataLabels,omitempty" xml:"DataLabels"`
    Explosion int64 `json:"Explosion,omitempty" xml:"Explosion"`
    Marker *Marker `json:"Marker,omitempty" xml:"Marker"`
    Shadow bool `json:"Shadow,omitempty" xml:"Shadow"`
    XValue string `json:"XValue,omitempty" xml:"XValue"`
    YValue string `json:"YValue,omitempty" xml:"YValue"`
}