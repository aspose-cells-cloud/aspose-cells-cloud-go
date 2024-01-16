/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="Chart.go">
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

type Chart struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    AutoScaling bool `json:"AutoScaling,omitempty" xml:"AutoScaling"`
    BackWall *LinkElement `json:"BackWall,omitempty" xml:"BackWall"`
    CategoryAxis *LinkElement `json:"CategoryAxis,omitempty" xml:"CategoryAxis"`
    ChartArea *LinkElement `json:"ChartArea,omitempty" xml:"ChartArea"`
    ChartDataTable *LinkElement `json:"ChartDataTable,omitempty" xml:"ChartDataTable"`
    ChartObject *LinkElement `json:"ChartObject,omitempty" xml:"ChartObject"`
    DepthPercent int64 `json:"DepthPercent,omitempty" xml:"DepthPercent"`
    Elevation int64 `json:"Elevation,omitempty" xml:"Elevation"`
    FirstSliceAngle int64 `json:"FirstSliceAngle,omitempty" xml:"FirstSliceAngle"`
    Floor *LinkElement `json:"Floor,omitempty" xml:"Floor"`
    GapDepth int64 `json:"GapDepth,omitempty" xml:"GapDepth"`
    GapWidth int64 `json:"GapWidth,omitempty" xml:"GapWidth"`
    HeightPercent int64 `json:"HeightPercent,omitempty" xml:"HeightPercent"`
    HidePivotFieldButtons bool `json:"HidePivotFieldButtons,omitempty" xml:"HidePivotFieldButtons"`
    Is3D bool `json:"Is3D,omitempty" xml:"Is3D"`
    IsRectangularCornered bool `json:"IsRectangularCornered,omitempty" xml:"IsRectangularCornered"`
    Legend *LinkElement `json:"Legend,omitempty" xml:"Legend"`
    Name string `json:"Name,omitempty" xml:"Name"`
    NSeries *LinkElement `json:"NSeries,omitempty" xml:"NSeries"`
    PageSetup *LinkElement `json:"PageSetup,omitempty" xml:"PageSetup"`
    Perspective int64 `json:"Perspective,omitempty" xml:"Perspective"`
    PivotSource string `json:"PivotSource,omitempty" xml:"PivotSource"`
    Placement string `json:"Placement,omitempty" xml:"Placement"`
    PlotArea *LinkElement `json:"PlotArea,omitempty" xml:"PlotArea"`
    PlotEmptyCellsType string `json:"PlotEmptyCellsType,omitempty" xml:"PlotEmptyCellsType"`
    PlotVisibleCells bool `json:"PlotVisibleCells,omitempty" xml:"PlotVisibleCells"`
    PrintSize string `json:"PrintSize,omitempty" xml:"PrintSize"`
    RightAngleAxes bool `json:"RightAngleAxes,omitempty" xml:"RightAngleAxes"`
    RotationAngle int64 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
    SecondCategoryAxis *LinkElement `json:"SecondCategoryAxis,omitempty" xml:"SecondCategoryAxis"`
    SecondValueAxis *LinkElement `json:"SecondValueAxis,omitempty" xml:"SecondValueAxis"`
    SeriesAxis *LinkElement `json:"SeriesAxis,omitempty" xml:"SeriesAxis"`
    Shapes *LinkElement `json:"Shapes,omitempty" xml:"Shapes"`
    ShowDataTable bool `json:"ShowDataTable,omitempty" xml:"ShowDataTable"`
    ShowLegend bool `json:"ShowLegend,omitempty" xml:"ShowLegend"`
    SideWall *LinkElement `json:"SideWall,omitempty" xml:"SideWall"`
    SizeWithWindow bool `json:"SizeWithWindow,omitempty" xml:"SizeWithWindow"`
    Style int64 `json:"Style,omitempty" xml:"Style"`
    Title *LinkElement `json:"Title,omitempty" xml:"Title"`
    Type_ string `json:"Type,omitempty" xml:"Type"`
    ValueAxis *LinkElement `json:"ValueAxis,omitempty" xml:"ValueAxis"`
    Walls *LinkElement `json:"Walls,omitempty" xml:"Walls"`
    WallsAndGridlines2D bool `json:"WallsAndGridlines2D,omitempty" xml:"WallsAndGridlines2D"`
}
