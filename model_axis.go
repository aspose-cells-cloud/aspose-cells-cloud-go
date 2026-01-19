/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="Axis.go">
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

type Axis struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    Area *Area `json:"Area,omitempty" xml:"Area"`
    AxisBetweenCategories bool `json:"AxisBetweenCategories,omitempty" xml:"AxisBetweenCategories"`
    AxisLine *Line `json:"AxisLine,omitempty" xml:"AxisLine"`
    BaseUnitScale string `json:"BaseUnitScale,omitempty" xml:"BaseUnitScale"`
    CategoryType string `json:"CategoryType,omitempty" xml:"CategoryType"`
    CrossAt float64 `json:"CrossAt,omitempty" xml:"CrossAt"`
    CrossType string `json:"CrossType,omitempty" xml:"CrossType"`
    DisplayUnit string `json:"DisplayUnit,omitempty" xml:"DisplayUnit"`
    DisplayUnitLabel *DisplayUnitLabel `json:"DisplayUnitLabel,omitempty" xml:"DisplayUnitLabel"`
    HasMultiLevelLabels bool `json:"HasMultiLevelLabels,omitempty" xml:"HasMultiLevelLabels"`
    IsAutomaticMajorUnit bool `json:"IsAutomaticMajorUnit,omitempty" xml:"IsAutomaticMajorUnit"`
    IsAutomaticMaxValue bool `json:"IsAutomaticMaxValue,omitempty" xml:"IsAutomaticMaxValue"`
    IsAutomaticMinorUnit bool `json:"IsAutomaticMinorUnit,omitempty" xml:"IsAutomaticMinorUnit"`
    IsAutomaticMinValue bool `json:"IsAutomaticMinValue,omitempty" xml:"IsAutomaticMinValue"`
    IsDisplayUnitLabelShown bool `json:"IsDisplayUnitLabelShown,omitempty" xml:"IsDisplayUnitLabelShown"`
    IsLogarithmic bool `json:"IsLogarithmic,omitempty" xml:"IsLogarithmic"`
    IsPlotOrderReversed bool `json:"IsPlotOrderReversed,omitempty" xml:"IsPlotOrderReversed"`
    IsVisible bool `json:"IsVisible,omitempty" xml:"IsVisible"`
    LogBase float64 `json:"LogBase,omitempty" xml:"LogBase"`
    MajorGridLines *Line `json:"MajorGridLines,omitempty" xml:"MajorGridLines"`
    MajorTickMark string `json:"MajorTickMark,omitempty" xml:"MajorTickMark"`
    MajorUnit float64 `json:"MajorUnit,omitempty" xml:"MajorUnit"`
    MajorUnitScale string `json:"MajorUnitScale,omitempty" xml:"MajorUnitScale"`
    MaxValue float64 `json:"MaxValue,omitempty" xml:"MaxValue"`
    MinorGridLines *Line `json:"MinorGridLines,omitempty" xml:"MinorGridLines"`
    MinorTickMark string `json:"MinorTickMark,omitempty" xml:"MinorTickMark"`
    MinorUnit float64 `json:"MinorUnit,omitempty" xml:"MinorUnit"`
    MinorUnitScale string `json:"MinorUnitScale,omitempty" xml:"MinorUnitScale"`
    MinValue float64 `json:"MinValue,omitempty" xml:"MinValue"`
    TickLabelPosition string `json:"TickLabelPosition,omitempty" xml:"TickLabelPosition"`
    TickLabels *TickLabels `json:"TickLabels,omitempty" xml:"TickLabels"`
    TickLabelSpacing int64 `json:"TickLabelSpacing,omitempty" xml:"TickLabelSpacing"`
    TickMarkSpacing int64 `json:"TickMarkSpacing,omitempty" xml:"TickMarkSpacing"`
    Title *Title `json:"Title,omitempty" xml:"Title"`
}
