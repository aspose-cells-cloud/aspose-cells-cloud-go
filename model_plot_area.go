/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="PlotArea.go">
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

type PlotArea struct {
     
        Area *Area `json:"Area,omitempty" xml:"Area"`
        AutoScaleFont bool `json:"AutoScaleFont,omitempty" xml:"AutoScaleFont"`
        BackgroundMode string `json:"BackgroundMode,omitempty" xml:"BackgroundMode"`
        Border *Line `json:"Border,omitempty" xml:"Border"`
        Font *Font `json:"Font,omitempty" xml:"Font"`
        IsAutomaticSize bool `json:"IsAutomaticSize,omitempty" xml:"IsAutomaticSize"`
        IsInnerMode bool `json:"IsInnerMode,omitempty" xml:"IsInnerMode"`
        Shadow bool `json:"Shadow,omitempty" xml:"Shadow"`
        ShapeProperties []LinkElement `json:"ShapeProperties,omitempty" xml:"ShapeProperties"`
        Width int64 `json:"Width,omitempty" xml:"Width"`
        Height int64 `json:"Height,omitempty" xml:"Height"`
        X int64 `json:"X,omitempty" xml:"X"`
        Y int64 `json:"Y,omitempty" xml:"Y"`
 
    InnerHeight int64 `json:"InnerHeight,omitempty" xml:"InnerHeight"`
    InnerWidth int64 `json:"InnerWidth,omitempty" xml:"InnerWidth"`
    InnerX int64 `json:"InnerX,omitempty" xml:"InnerX"`
    InnerY int64 `json:"InnerY,omitempty" xml:"InnerY"`
}
