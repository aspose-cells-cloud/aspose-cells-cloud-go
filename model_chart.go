/*
 *  Copyright (c) 2020 Aspose.Cells Cloud
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

type Chart struct {
	Link *Link `json:"link,omitempty" xml:"link"`
	// Gets or sets the angle of the first pie-chart or doughnut-chart slice, in degrees (clockwise from vertical). Applies only to pie, 3-D pie, and doughnut charts, 0 to 360.
	FirstSliceAngle int64 `json:"FirstSliceAngle,omitempty" xml:"FirstSliceAngle"`
	// Returns a Floor object that represents the walls of a 3-D chart.             
	Floor *LinkElement `json:"Floor,omitempty" xml:"Floor"`
	// Gets and sets how to plot the empty cells.
	PlotEmptyCellsType string `json:"PlotEmptyCellsType,omitempty" xml:"PlotEmptyCellsType"`
	// True if Microsoft Excel scales a 3-D chart so that it's closer in size to the equivalent 2-D chart. The RightAngleAxes property must be True.
	AutoScaling bool `json:"AutoScaling,omitempty" xml:"AutoScaling"`
	// Gets and sets the builtin style.
	Style int64 `json:"Style,omitempty" xml:"Style"`
	// Gets the chart's series axis.
	SeriesAxis *LinkElement `json:"SeriesAxis,omitempty" xml:"SeriesAxis"`
	// Gets the chart's Y axis.
	ValueAxis *LinkElement `json:"ValueAxis,omitempty" xml:"ValueAxis"`
	// Gets or sets a value indicating whether the chart displays a data table.
	ShowDataTable bool `json:"ShowDataTable,omitempty" xml:"ShowDataTable"`
	// Indicates whether the chart is a 3d chart.
	Is3D bool `json:"Is3D,omitempty" xml:"Is3D"`
	// Gets the chart area in the worksheet
	ChartArea *LinkElement `json:"ChartArea,omitempty" xml:"ChartArea"`
	// Represents the elevation of the 3-D chart view, in degrees.
	Elevation int64 `json:"Elevation,omitempty" xml:"Elevation"`
	SideWall *LinkElement `json:"SideWall,omitempty" xml:"SideWall"`
	// Gets or sets a chart's type.
	Type_ string `json:"Type,omitempty" xml:"Type"`
	// Gets the chart's title.
	Title *LinkElement `json:"Title,omitempty" xml:"Title"`
	// Returns a Walls object that represents the walls of a 3-D chart.
	Walls *LinkElement `json:"Walls,omitempty" xml:"Walls"`
	BackWall *LinkElement `json:"BackWall,omitempty" xml:"BackWall"`
	// Represents the chart data table.
	ChartDataTable *LinkElement `json:"ChartDataTable,omitempty" xml:"ChartDataTable"`
	// Returns or sets the height of a 3-D chart as a percentage of the chart width (between 5 and 500 percent).
	HeightPercent int64 `json:"HeightPercent,omitempty" xml:"HeightPercent"`
	// Returns or sets the space between bar or column clusters, as a percentage of the bar or column width. The value of this property must be between 0 and 500.             
	GapWidth int64 `json:"GapWidth,omitempty" xml:"GapWidth"`
	// Gets the chart legend.
	Legend *LinkElement `json:"Legend,omitempty" xml:"Legend"`
	// Represents the chartShape;
	ChartObject *LinkElement `json:"ChartObject,omitempty" xml:"ChartObject"`
	// Gets or sets a value indicating whether the chart displays a data table.
	IsRectangularCornered bool `json:"IsRectangularCornered,omitempty" xml:"IsRectangularCornered"`
	// Gets the chart's second X axis.
	SecondCategoryAxis *LinkElement `json:"SecondCategoryAxis,omitempty" xml:"SecondCategoryAxis"`
	// Gets the chart's second Y axis.
	SecondValueAxis *LinkElement `json:"SecondValueAxis,omitempty" xml:"SecondValueAxis"`
	// Represents the way the chart is attached to the cells below it.
	Placement string `json:"Placement,omitempty" xml:"Placement"`
	// Gets and sets the name of the chart.
	Name string `json:"Name,omitempty" xml:"Name"`
	// True if Microsoft Excel resizes the chart to match the size of the chart sheet window.
	SizeWithWindow bool `json:"SizeWithWindow,omitempty" xml:"SizeWithWindow"`
	// True if the chart axes are at right angles.Applies only for 3-D charts(except Column3D and 3-D Pie Charts).
	RightAngleAxes bool `json:"RightAngleAxes,omitempty" xml:"RightAngleAxes"`
	// Indicates whether only plot visible cells.
	PlotVisibleCells bool `json:"PlotVisibleCells,omitempty" xml:"PlotVisibleCells"`
	// Gets or sets a value indicating whether the chart legend will be displayed. Default is true.
	ShowLegend bool `json:"ShowLegend,omitempty" xml:"ShowLegend"`
	// The source is the data of the pivotTable.If PivotSource is not empty ,the chart is PivotChart.
	PivotSource string `json:"PivotSource,omitempty" xml:"PivotSource"`
	// Represents the depth of a 3-D chart as a percentage of the chart width (between 20 and 2000 percent).
	DepthPercent int64 `json:"DepthPercent,omitempty" xml:"DepthPercent"`
	// Gets and sets the printed chart size.
	PrintSize string `json:"PrintSize,omitempty" xml:"PrintSize"`
	// Gets or sets the distance between the data series in a 3-D chart, as a percentage of the marker width.The value of this property must be between 0 and 500.
	GapDepth int64 `json:"GapDepth,omitempty" xml:"GapDepth"`
	// Returns all drawing shapes in this chart.
	Shapes *LinkElement `json:"Shapes,omitempty" xml:"Shapes"`
	// True if gridlines are drawn two-dimensionally on a 3-D chart.
	WallsAndGridlines2D bool `json:"WallsAndGridlines2D,omitempty" xml:"WallsAndGridlines2D"`
	// Gets a SeriesCollection collection representing the data series in the chart.
	NSeries *LinkElement `json:"NSeries,omitempty" xml:"NSeries"`
	// Represents the rotation of the 3-D chart view (the rotation of the plot area around the z-axis, in degrees).
	RotationAngle int64 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
	// Gets the chart's plot area which includes axis tick lables.
	PlotArea *LinkElement `json:"PlotArea,omitempty" xml:"PlotArea"`
	// Gets the chart's X axis. The property is read only
	CategoryAxis *LinkElement `json:"CategoryAxis,omitempty" xml:"CategoryAxis"`
	// Returns or sets the perspective for the 3-D chart view. Must be between 0 and 100.This property is ignored if the RightAngleAxes property is True.
	Perspective int32 `json:"Perspective,omitempty" xml:"Perspective"`
	// Indicates whether hide the pivot chart field buttons only when the chart is PivotChart
	HidePivotFieldButtons bool `json:"HidePivotFieldButtons,omitempty" xml:"HidePivotFieldButtons"`
	// Represents the page setup description in this chart.
	PageSetup *LinkElement `json:"PageSetup,omitempty" xml:"PageSetup"`
}
