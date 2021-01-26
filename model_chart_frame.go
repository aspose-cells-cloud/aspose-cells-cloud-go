/*
 *  Copyright (c) 2021 Aspose.Cells Cloud
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

type ChartFrame struct {
	Link *Link `json:"link,omitempty" xml:"link"`
	IsInnerMode bool `json:"IsInnerMode,omitempty" xml:"IsInnerMode"`
	ShapeProperties []LinkElement `json:"ShapeProperties,omitempty" xml:"ShapeProperties"`
	AutoScaleFont bool `json:"AutoScaleFont,omitempty" xml:"AutoScaleFont"`
	Area *Area `json:"Area,omitempty" xml:"Area"`
	Height int64 `json:"Height,omitempty" xml:"Height"`
	Width int64 `json:"Width,omitempty" xml:"Width"`
	BackgroundMode string `json:"BackgroundMode,omitempty" xml:"BackgroundMode"`
	IsAutomaticSize bool `json:"IsAutomaticSize,omitempty" xml:"IsAutomaticSize"`
	Y int64 `json:"Y,omitempty" xml:"Y"`
	X int64 `json:"X,omitempty" xml:"X"`
	Shadow bool `json:"Shadow,omitempty" xml:"Shadow"`
	Font *Font `json:"Font,omitempty" xml:"Font"`
	Border *Line `json:"Border,omitempty" xml:"Border"`
}
