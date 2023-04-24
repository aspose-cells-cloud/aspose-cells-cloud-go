/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="SparklineGroup.go">
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

type SparklineGroup struct {
 
    DisplayHidden bool `json:"DisplayHidden,omitempty" xml:"DisplayHidden"`
    FirstPointColor *CellsColor `json:"FirstPointColor,omitempty" xml:"FirstPointColor"`
    HighPointColor *CellsColor `json:"HighPointColor,omitempty" xml:"HighPointColor"`
    HorizontalAxisColor *CellsColor `json:"HorizontalAxisColor,omitempty" xml:"HorizontalAxisColor"`
    HorizontalAxisDateRange string `json:"HorizontalAxisDateRange,omitempty" xml:"HorizontalAxisDateRange"`
    LastPointColor *CellsColor `json:"LastPointColor,omitempty" xml:"LastPointColor"`
    LineWeight float64 `json:"LineWeight,omitempty" xml:"LineWeight"`
    LowPointColor *CellsColor `json:"LowPointColor,omitempty" xml:"LowPointColor"`
    MarkersColor *CellsColor `json:"MarkersColor,omitempty" xml:"MarkersColor"`
    NegativePointsColor *CellsColor `json:"NegativePointsColor,omitempty" xml:"NegativePointsColor"`
    PlotEmptyCellsType string `json:"PlotEmptyCellsType,omitempty" xml:"PlotEmptyCellsType"`
    PlotRightToLeft bool `json:"PlotRightToLeft,omitempty" xml:"PlotRightToLeft"`
    PresetStyle string `json:"PresetStyle,omitempty" xml:"PresetStyle"`
    SeriesColor *CellsColor `json:"SeriesColor,omitempty" xml:"SeriesColor"`
    ShowFirstPoint bool `json:"ShowFirstPoint,omitempty" xml:"ShowFirstPoint"`
    ShowHighPoint bool `json:"ShowHighPoint,omitempty" xml:"ShowHighPoint"`
    ShowHorizontalAxis bool `json:"ShowHorizontalAxis,omitempty" xml:"ShowHorizontalAxis"`
    ShowLastPoint bool `json:"ShowLastPoint,omitempty" xml:"ShowLastPoint"`
    ShowLowPoint bool `json:"ShowLowPoint,omitempty" xml:"ShowLowPoint"`
    ShowMarkers bool `json:"ShowMarkers,omitempty" xml:"ShowMarkers"`
    ShowNegativePoints bool `json:"ShowNegativePoints,omitempty" xml:"ShowNegativePoints"`
    SparklineCollection []Sparkline `json:"SparklineCollection,omitempty" xml:"SparklineCollection"`
    Type_ string `json:"Type,omitempty" xml:"Type"`
    VerticalAxisMaxValue float64 `json:"VerticalAxisMaxValue,omitempty" xml:"VerticalAxisMaxValue"`
    VerticalAxisMaxValueType string `json:"VerticalAxisMaxValueType,omitempty" xml:"VerticalAxisMaxValueType"`
    VerticalAxisMinValue float64 `json:"VerticalAxisMinValue,omitempty" xml:"VerticalAxisMinValue"`
    VerticalAxisMinValueType string `json:"VerticalAxisMinValueType,omitempty" xml:"VerticalAxisMinValueType"`
}
