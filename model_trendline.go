/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="Trendline.go">
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

type Trendline struct {
     
        BeginArrowLength string `json:"BeginArrowLength,omitempty" xml:"BeginArrowLength"`
        BeginArrowWidth string `json:"BeginArrowWidth,omitempty" xml:"BeginArrowWidth"`
        BeginType string `json:"BeginType,omitempty" xml:"BeginType"`
        CapType string `json:"CapType,omitempty" xml:"CapType"`
        Color *Color `json:"Color,omitempty" xml:"Color"`
        CompoundType string `json:"CompoundType,omitempty" xml:"CompoundType"`
        DashType string `json:"DashType,omitempty" xml:"DashType"`
        EndArrowLength string `json:"EndArrowLength,omitempty" xml:"EndArrowLength"`
        EndArrowWidth string `json:"EndArrowWidth,omitempty" xml:"EndArrowWidth"`
        EndType string `json:"EndType,omitempty" xml:"EndType"`
        GradientFill *GradientFill `json:"GradientFill,omitempty" xml:"GradientFill"`
        IsAuto bool `json:"IsAuto,omitempty" xml:"IsAuto"`
        IsAutomaticColor bool `json:"IsAutomaticColor,omitempty" xml:"IsAutomaticColor"`
        IsVisible bool `json:"IsVisible,omitempty" xml:"IsVisible"`
        JoinType string `json:"JoinType,omitempty" xml:"JoinType"`
        Style string `json:"Style,omitempty" xml:"Style"`
        Transparency float64 `json:"Transparency,omitempty" xml:"Transparency"`
        Weight string `json:"Weight,omitempty" xml:"Weight"`
        WeightPt float64 `json:"WeightPt,omitempty" xml:"WeightPt"`
 
    Link *Link `json:"link,omitempty" xml:"link"`
    Backward float64 `json:"Backward,omitempty" xml:"Backward"`
    DataLabels *DataLabels `json:"DataLabels,omitempty" xml:"DataLabels"`
    DisplayEquation bool `json:"DisplayEquation,omitempty" xml:"DisplayEquation"`
    DisplayRSquared bool `json:"DisplayRSquared,omitempty" xml:"DisplayRSquared"`
    Forward float64 `json:"Forward,omitempty" xml:"Forward"`
    Intercept float64 `json:"Intercept,omitempty" xml:"Intercept"`
    IsNameAuto bool `json:"IsNameAuto,omitempty" xml:"IsNameAuto"`
    LegendEntry *LegendEntry `json:"LegendEntry,omitempty" xml:"LegendEntry"`
    Name string `json:"Name,omitempty" xml:"Name"`
    Order int64 `json:"Order,omitempty" xml:"Order"`
    Period int64 `json:"Period,omitempty" xml:"Period"`
    Type_ string `json:"Type,omitempty" xml:"Type"`
}
