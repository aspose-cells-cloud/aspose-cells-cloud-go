/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="Series.go">
*   Copyright (c) 2026 Aspose.Cells Cloud
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

type Series struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    Area *Area `json:"Area,omitempty" xml:"Area"`
    Bar3DShapeType string `json:"Bar3DShapeType,omitempty" xml:"Bar3DShapeType"`
    Border *Line `json:"Border,omitempty" xml:"Border"`
    BubbleScale int64 `json:"BubbleScale,omitempty" xml:"BubbleScale"`
    BubbleSizes string `json:"BubbleSizes,omitempty" xml:"BubbleSizes"`
    CountOfDataValues int64 `json:"CountOfDataValues,omitempty" xml:"CountOfDataValues"`
    DataLabels *DataLabels `json:"DataLabels,omitempty" xml:"DataLabels"`
    DisplayName string `json:"DisplayName,omitempty" xml:"DisplayName"`
    DoughnutHoleSize int64 `json:"DoughnutHoleSize,omitempty" xml:"DoughnutHoleSize"`
    DownBars *DropBars `json:"DownBars,omitempty" xml:"DownBars"`
    DropLines *Line `json:"DropLines,omitempty" xml:"DropLines"`
    Explosion int64 `json:"Explosion,omitempty" xml:"Explosion"`
    FirstSliceAngle int64 `json:"FirstSliceAngle,omitempty" xml:"FirstSliceAngle"`
    GapWidth int64 `json:"GapWidth,omitempty" xml:"GapWidth"`
    Has3DEffect bool `json:"Has3DEffect,omitempty" xml:"Has3DEffect"`
    HasDropLines bool `json:"HasDropLines,omitempty" xml:"HasDropLines"`
    HasHiLoLines bool `json:"HasHiLoLines,omitempty" xml:"HasHiLoLines"`
    HasLeaderLines bool `json:"HasLeaderLines,omitempty" xml:"HasLeaderLines"`
    HasRadarAxisLabels bool `json:"HasRadarAxisLabels,omitempty" xml:"HasRadarAxisLabels"`
    HasSeriesLines bool `json:"HasSeriesLines,omitempty" xml:"HasSeriesLines"`
    HasUpDownBars bool `json:"HasUpDownBars,omitempty" xml:"HasUpDownBars"`
    HiLoLines *Line `json:"HiLoLines,omitempty" xml:"HiLoLines"`
    IsAutoSplit bool `json:"IsAutoSplit,omitempty" xml:"IsAutoSplit"`
    IsColorVaried bool `json:"IsColorVaried,omitempty" xml:"IsColorVaried"`
    LeaderLines *Line `json:"LeaderLines,omitempty" xml:"LeaderLines"`
    LegendEntry *LegendEntry `json:"LegendEntry,omitempty" xml:"LegendEntry"`
    Marker *Marker `json:"Marker,omitempty" xml:"Marker"`
    Name string `json:"Name,omitempty" xml:"Name"`
    Overlap int64 `json:"Overlap,omitempty" xml:"Overlap"`
    PlotOnSecondAxis bool `json:"PlotOnSecondAxis,omitempty" xml:"PlotOnSecondAxis"`
    Points *LinkElement `json:"Points,omitempty" xml:"Points"`
    SecondPlotSize int64 `json:"SecondPlotSize,omitempty" xml:"SecondPlotSize"`
    SeriesLines *Line `json:"SeriesLines,omitempty" xml:"SeriesLines"`
    Shadow bool `json:"Shadow,omitempty" xml:"Shadow"`
    ShowNegativeBubbles bool `json:"ShowNegativeBubbles,omitempty" xml:"ShowNegativeBubbles"`
    SizeRepresents string `json:"SizeRepresents,omitempty" xml:"SizeRepresents"`
    Smooth bool `json:"Smooth,omitempty" xml:"Smooth"`
    SplitType string `json:"SplitType,omitempty" xml:"SplitType"`
    SplitValue float64 `json:"SplitValue,omitempty" xml:"SplitValue"`
    TrendLines *Trendlines `json:"TrendLines,omitempty" xml:"TrendLines"`
    Type_ string `json:"Type,omitempty" xml:"Type"`
    UpBars *DropBars `json:"UpBars,omitempty" xml:"UpBars"`
    Values string `json:"Values,omitempty" xml:"Values"`
    XErrorBar *ErrorBar `json:"XErrorBar,omitempty" xml:"XErrorBar"`
    XValues string `json:"XValues,omitempty" xml:"XValues"`
    YErrorBar *ErrorBar `json:"YErrorBar,omitempty" xml:"YErrorBar"`
}
