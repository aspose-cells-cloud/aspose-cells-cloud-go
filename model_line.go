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

type Line struct {
	Style string `json:"Style,omitempty" xml:"Style"`
	IsAuto bool `json:"IsAuto,omitempty" xml:"IsAuto"`
	BeginArrowLength string `json:"BeginArrowLength,omitempty" xml:"BeginArrowLength"`
	Weight string `json:"Weight,omitempty" xml:"Weight"`
	GradientFill *GradientFill `json:"GradientFill,omitempty" xml:"GradientFill"`
	Color *Color `json:"Color,omitempty" xml:"Color"`
	CompoundType string `json:"CompoundType,omitempty" xml:"CompoundType"`
	WeightPt float64 `json:"WeightPt,omitempty" xml:"WeightPt"`
	IsVisible bool `json:"IsVisible,omitempty" xml:"IsVisible"`
	JoinType string `json:"JoinType,omitempty" xml:"JoinType"`
	EndArrowLength string `json:"EndArrowLength,omitempty" xml:"EndArrowLength"`
	IsAutomaticColor bool `json:"IsAutomaticColor,omitempty" xml:"IsAutomaticColor"`
	DashType string `json:"DashType,omitempty" xml:"DashType"`
	BeginType string `json:"BeginType,omitempty" xml:"BeginType"`
	CapType string `json:"CapType,omitempty" xml:"CapType"`
	EndType string `json:"EndType,omitempty" xml:"EndType"`
	BeginArrowWidth string `json:"BeginArrowWidth,omitempty" xml:"BeginArrowWidth"`
	EndArrowWidth string `json:"EndArrowWidth,omitempty" xml:"EndArrowWidth"`
	Transparency float64 `json:"Transparency,omitempty" xml:"Transparency"`
}
