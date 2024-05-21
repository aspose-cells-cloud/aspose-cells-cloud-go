/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="DataLabels.go">
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

type DataLabels struct {
     
        Area *Area `json:"Area,omitempty" xml:"Area"`
        AutoScaleFont bool `json:"AutoScaleFont,omitempty" xml:"AutoScaleFont"`
        BackgroundMode string `json:"BackgroundMode,omitempty" xml:"BackgroundMode"`
        Border *Line `json:"Border,omitempty" xml:"Border"`
        Font *Font `json:"Font,omitempty" xml:"Font"`
        IsAutomaticSize bool `json:"IsAutomaticSize,omitempty" xml:"IsAutomaticSize"`
        IsInnerMode bool `json:"IsInnerMode,omitempty" xml:"IsInnerMode"`
        Shadow bool `json:"Shadow,omitempty" xml:"Shadow"`
        Width int64 `json:"Width,omitempty" xml:"Width"`
        Height int64 `json:"Height,omitempty" xml:"Height"`
        X int64 `json:"X,omitempty" xml:"X"`
        Y int64 `json:"Y,omitempty" xml:"Y"`
 
    IsAutoText bool `json:"IsAutoText,omitempty" xml:"IsAutoText"`
    IsDeleted bool `json:"IsDeleted,omitempty" xml:"IsDeleted"`
    LinkedSource string `json:"LinkedSource,omitempty" xml:"LinkedSource"`
    Number int64 `json:"Number,omitempty" xml:"Number"`
    NumberFormat string `json:"NumberFormat,omitempty" xml:"NumberFormat"`
    NumberFormatLinked bool `json:"NumberFormatLinked,omitempty" xml:"NumberFormatLinked"`
    Position string `json:"Position,omitempty" xml:"Position"`
    RotationAngle int64 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
    Separator string `json:"Separator,omitempty" xml:"Separator"`
    ShowBubbleSize bool `json:"ShowBubbleSize,omitempty" xml:"ShowBubbleSize"`
    ShowCategoryName bool `json:"ShowCategoryName,omitempty" xml:"ShowCategoryName"`
    ShowLegendKey bool `json:"ShowLegendKey,omitempty" xml:"ShowLegendKey"`
    ShowPercentage bool `json:"ShowPercentage,omitempty" xml:"ShowPercentage"`
    ShowSeriesName bool `json:"ShowSeriesName,omitempty" xml:"ShowSeriesName"`
    ShowValue bool `json:"ShowValue,omitempty" xml:"ShowValue"`
    Text string `json:"Text,omitempty" xml:"Text"`
    TextDirection string `json:"TextDirection,omitempty" xml:"TextDirection"`
    TextHorizontalAlignment string `json:"TextHorizontalAlignment,omitempty" xml:"TextHorizontalAlignment"`
    TextVerticalAlignment string `json:"TextVerticalAlignment,omitempty" xml:"TextVerticalAlignment"`
}
