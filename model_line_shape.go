/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="LineShape.go">
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

type LineShape struct {
     
        Name string `json:"Name,omitempty" xml:"Name"`
        MsoDrawingType string `json:"MsoDrawingType,omitempty" xml:"MsoDrawingType"`
        AutoShapeType string `json:"AutoShapeType,omitempty" xml:"AutoShapeType"`
        Placement string `json:"Placement,omitempty" xml:"Placement"`
        UpperLeftRow int64 `json:"UpperLeftRow,omitempty" xml:"UpperLeftRow"`
        Top int64 `json:"Top,omitempty" xml:"Top"`
        UpperLeftColumn int64 `json:"UpperLeftColumn,omitempty" xml:"UpperLeftColumn"`
        Left int64 `json:"Left,omitempty" xml:"Left"`
        LowerRightRow int64 `json:"LowerRightRow,omitempty" xml:"LowerRightRow"`
        Bottom int64 `json:"Bottom,omitempty" xml:"Bottom"`
        LowerRightColumn int64 `json:"LowerRightColumn,omitempty" xml:"LowerRightColumn"`
        Right int64 `json:"Right,omitempty" xml:"Right"`
        Width int64 `json:"Width,omitempty" xml:"Width"`
        Height int64 `json:"Height,omitempty" xml:"Height"`
        X int64 `json:"X,omitempty" xml:"X"`
        Y int64 `json:"Y,omitempty" xml:"Y"`
        RotationAngle float64 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
        HtmlText string `json:"HtmlText,omitempty" xml:"HtmlText"`
        Text string `json:"Text,omitempty" xml:"Text"`
        AlternativeText string `json:"AlternativeText,omitempty" xml:"AlternativeText"`
        TextHorizontalAlignment string `json:"TextHorizontalAlignment,omitempty" xml:"TextHorizontalAlignment"`
        TextHorizontalOverflow string `json:"TextHorizontalOverflow,omitempty" xml:"TextHorizontalOverflow"`
        TextOrientationType string `json:"TextOrientationType,omitempty" xml:"TextOrientationType"`
        TextVerticalAlignment string `json:"TextVerticalAlignment,omitempty" xml:"TextVerticalAlignment"`
        TextVerticalOverflow string `json:"TextVerticalOverflow,omitempty" xml:"TextVerticalOverflow"`
        IsGroup bool `json:"IsGroup,omitempty" xml:"IsGroup"`
        IsHidden bool `json:"IsHidden,omitempty" xml:"IsHidden"`
        IsLockAspectRatio bool `json:"IsLockAspectRatio,omitempty" xml:"IsLockAspectRatio"`
        IsLocked bool `json:"IsLocked,omitempty" xml:"IsLocked"`
        IsPrintable bool `json:"IsPrintable,omitempty" xml:"IsPrintable"`
        IsTextWrapped bool `json:"IsTextWrapped,omitempty" xml:"IsTextWrapped"`
        IsWordArt bool `json:"IsWordArt,omitempty" xml:"IsWordArt"`
        LinkedCell string `json:"LinkedCell,omitempty" xml:"LinkedCell"`
        ZOrderPosition int64 `json:"ZOrderPosition,omitempty" xml:"ZOrderPosition"`
        Font *Font `json:"Font,omitempty" xml:"Font"`
        Hyperlink string `json:"Hyperlink,omitempty" xml:"Hyperlink"`
 
    BeginArrowheadLength string `json:"BeginArrowheadLength,omitempty" xml:"BeginArrowheadLength"`
    BeginArrowheadStyle string `json:"BeginArrowheadStyle,omitempty" xml:"BeginArrowheadStyle"`
    BeginArrowheadWidth string `json:"BeginArrowheadWidth,omitempty" xml:"BeginArrowheadWidth"`
    EndArrowheadLength string `json:"EndArrowheadLength,omitempty" xml:"EndArrowheadLength"`
    EndArrowheadStyle string `json:"EndArrowheadStyle,omitempty" xml:"EndArrowheadStyle"`
    EndArrowheadWidth string `json:"EndArrowheadWidth,omitempty" xml:"EndArrowheadWidth"`
}
