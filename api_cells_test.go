/*
 *  Copyright (c) 2022 Aspose.Cells Cloud
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

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"
)

var isDockerSDK, errcon = strconv.ParseBool(os.Getenv("CellsCloudTestIsDockerTest"))

func TestCellsAutoFilterDeleteWorksheetDateFilter(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsAutoFilterDeleteWorksheetDateFilterOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.DateTimeGroupingType = "Day"
	args.Year = 2019
	args.Month = 12
	args.Day = 1
	args.Hour = 12
	args.Minute = 1
	args.Second = 1
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsAutoFilterDeleteWorksheetDateFilter(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsAutoFilterDeleteWorksheetDateFilter - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsAutoFilterDeleteWorksheetFilter(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsAutoFilterDeleteWorksheetFilterOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Criteria = "test"
	args.FieldIndex = 1
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsAutoFilterDeleteWorksheetFilter(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsAutoFilterDeleteWorksheetFilter - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsAutoFilterGetWorksheetAutoFilter(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsAutoFilterGetWorksheetAutoFilterOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsAutoFilterGetWorksheetAutoFilter(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsAutoFilterGetWorksheetAutoFilter - %d\n", GetBaseTest().GetTestNumber(), response.Code)
		fmt.Println(json.Marshal(*(response.AutoFilter)))
	}
}

func TestCellsAutoFilterPostWorksheetAutoFilterRefresh(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsAutoFilterPostWorksheetAutoFilterRefreshOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsAutoFilterPostWorksheetAutoFilterRefresh(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsAutoFilterPostWorksheetAutoFilterRefresh - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsAutoFilterPostWorksheetMatchBlanks(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsAutoFilterPostWorksheetMatchBlanksOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.FieldIndex = 0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsAutoFilterPostWorksheetMatchBlanks(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsAutoFilterPostWorksheetMatchBlanks - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsAutoFilterPostWorksheetMatchNonBlanks(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsAutoFilterPostWorksheetMatchNonBlanksOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.FieldIndex = 0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsAutoFilterPostWorksheetMatchNonBlanks(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsAutoFilterPostWorksheetMatchNonBlanks - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsAutoFilterPutWorksheetColorFilter(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsAutoFilterPutWorksheetColorFilterOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.FieldIndex = 0
	args.Range_ = GetRange()
	args.Refresh = true
	args.MatchBlanks = false
	args.ColorFilter = new(ColorFilterRequest)
	args.ColorFilter.Pattern = "Solid"

	color := new(Color)
	color.R = 255
	color.A = 255
	color.B = 255
	color.G = 0

	cellsColor := new(CellsColor)
	cellsColor.Color = color
	cellsColor.ThemeColor = new(ThemeColor)
	cellsColor.ThemeColor.ColorType = "Text2"
	cellsColor.ThemeColor.Tint = 1.0
	cellsColor.Type_ = "Automatic"
	args.ColorFilter.BackgroundColor = cellsColor

	color1 := new(Color)
	color1.R = 255
	color1.A = 255
	color1.B = 255
	color1.G = 0

	cellsColor1 := new(CellsColor)
	cellsColor1.Color = color1
	cellsColor1.ThemeColor = new(ThemeColor)
	cellsColor1.ThemeColor.ColorType = "Text2"
	cellsColor1.ThemeColor.Tint = 1.0
	cellsColor1.Type_ = "Automatic"
	args.ColorFilter.ForegroundColor = cellsColor1

	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsAutoFilterPutWorksheetColorFilter(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsAutoFilterPutWorksheetColorFilter - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsAutoFilterPutWorksheetCustomFilter(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsAutoFilterPutWorksheetCustomFilterOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Range_ = GetRange()
	args.FieldIndex = 0
	args.OperatorType1 = "LessOrEqual"
	args.Criteria1 = "test"
	args.OperatorType2 = "LessOrEqual"
	args.Criteria2 = "test"
	args.IsAnd = true
	args.Refresh = true
	args.MatchBlanks = true
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsAutoFilterPutWorksheetCustomFilter(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsAutoFilterPutWorksheetCustomFilter - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsAutoFilterPutWorksheetDateFilter(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsAutoFilterPutWorksheetDateFilterOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.FieldIndex = 0
	args.DateTimeGroupingType = "Day"
	args.Year = 2019
	args.Month = 12
	args.Day = 1
	args.Hour = 12
	args.Minute = 1
	args.Second = 1
	args.MatchBlanks = true

	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsAutoFilterPutWorksheetDateFilter(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsAutoFilterPutWorksheetDateFilter - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsAutoFilterPutWorksheetDynamicFilter(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsAutoFilterPutWorksheetDynamicFilterOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Range_ = GetRange()
	args.FieldIndex = 0
	args.DynamicFilterType = "May"
	args.MatchBlanks = true
	args.Refresh = true
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsAutoFilterPutWorksheetDynamicFilter(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsAutoFilterPutWorksheetDynamicFilter - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsAutoFilterPutWorksheetFilter(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsAutoFilterPutWorksheetFilterOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Criteria = "test"
	args.FieldIndex = 0
	args.MatchBlanks = true
	args.Refresh = true
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsAutoFilterPutWorksheetFilter(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsAutoFilterPutWorksheetFilter - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsAutoFilterPutWorksheetFilterTop10(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsAutoFilterPutWorksheetFilterTop10Opts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.FieldIndex = 0
	args.IsPercent = true
	args.IsTop = true
	args.Range_ = GetRange()
	args.MatchBlanks = true
	args.ItemCount = 1
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsAutoFilterPutWorksheetFilterTop10(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsAutoFilterPutWorksheetFilterTop10 - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsAutoFilterPutWorksheetIconFilter(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsAutoFilterPutWorksheetIconFilterOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.FieldIndex = 0
	args.IconId = 0
	args.IconSetType = "None"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsAutoFilterPutWorksheetIconFilter(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsAutoFilterPutWorksheetIconFilter - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsAutoshapesGetWorksheetAutoshape(t *testing.T) {
	if isDockerSDK {
		return
	}
	name := GetMyDoc()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsAutoshapesGetWorksheetAutoshapeOpts)
	args.Name = GetMyDoc()
	args.SheetName = GetSheet2()
	args.AutoshapeNumber = 4
	args.Format = "png"
	args.Folder = GetBaseTest().remoteFolder

	_, httpResponse, err := GetBaseTest().CellsAPI.CellsAutoshapesGetWorksheetAutoshape(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsAutoshapesGetWorksheetAutoshape - %d\n", GetBaseTest().GetTestNumber(), 0)
	}
}

func TestCellsAutoshapesGetWorksheetAutoshapes(t *testing.T) {
	if isDockerSDK {
		return
	}
	name := GetMyDoc()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsAutoshapesGetWorksheetAutoshapesOpts)
	args.Name = GetMyDoc()
	args.SheetName = GetSheet2()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsAutoshapesGetWorksheetAutoshapes(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsAutoshapesGetWorksheetAutoshapes - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsChartAreaGetChartArea(t *testing.T) {
	name := GetMyDoc()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsChartAreaGetChartAreaOpts)
	args.Name = GetMyDoc()
	args.SheetName = GetSheet3()
	args.ChartIndex = 0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsChartAreaGetChartArea(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsChartAreaGetChartArea - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsChartAreaGetChartAreaBorder(t *testing.T) {
	name := GetMyDoc()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsChartAreaGetChartAreaBorderOpts)
	args.Name = GetMyDoc()
	args.SheetName = GetSheet3()
	args.ChartIndex = 0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsChartAreaGetChartAreaBorder(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsChartAreaGetChartAreaBorder - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsChartAreaGetChartAreaFillFormat(t *testing.T) {
	name := GetMyDoc()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsChartAreaGetChartAreaFillFormatOpts)
	args.Name = GetMyDoc()
	args.SheetName = GetSheet3()
	args.ChartIndex = 0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsChartAreaGetChartAreaFillFormat(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsChartAreaGetChartAreaFillFormat - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsChartsDeleteWorksheetChartLegend(t *testing.T) {
	name := GetMyDoc()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsChartsDeleteWorksheetChartLegendOpts)
	args.Name = GetMyDoc()
	args.SheetName = GetSheet3()
	args.ChartIndex = 0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsChartsDeleteWorksheetChartLegend(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsChartsDeleteWorksheetChartLegend - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsChartsDeleteWorksheetChartTitle(t *testing.T) {
	name := GetMyDoc()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsChartsDeleteWorksheetChartTitleOpts)
	args.Name = GetMyDoc()
	args.SheetName = GetSheet3()
	args.ChartIndex = 0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsChartsDeleteWorksheetChartTitle(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsChartsDeleteWorksheetChartTitle - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsChartsDeleteWorksheetClearCharts(t *testing.T) {
	name := GetMyDoc()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsChartsDeleteWorksheetClearChartsOpts)
	args.Name = GetMyDoc()
	args.SheetName = GetSheet3()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsChartsDeleteWorksheetClearCharts(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsChartsDeleteWorksheetClearCharts - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsChartsDeleteWorksheetDeleteChart(t *testing.T) {
	name := GetMyDoc()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsChartsDeleteWorksheetDeleteChartOpts)
	args.Name = GetMyDoc()
	args.SheetName = GetSheet3()
	args.ChartIndex = 0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsChartsDeleteWorksheetDeleteChart(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsChartsDeleteWorksheetDeleteChart - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsChartsGetWorksheetChart(t *testing.T) {
	name := GetMyDoc()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsChartsGetWorksheetChartOpts)
	args.Name = GetMyDoc()
	args.SheetName = GetSheet3()
	args.ChartNumber = 0
	args.Folder = GetBaseTest().remoteFolder

	_, httpResponse, err := GetBaseTest().CellsAPI.CellsChartsGetWorksheetChart(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsChartsGetWorksheetChart - %d\n", GetBaseTest().GetTestNumber(), 0)
	}
}

func TestCellsChartsGetWorksheetChartLegend(t *testing.T) {
	name := GetMyDoc()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsChartsGetWorksheetChartLegendOpts)
	args.Name = GetMyDoc()
	args.SheetName = GetSheet3()
	args.ChartIndex = 0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsChartsGetWorksheetChartLegend(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsChartsGetWorksheetChartLegend - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsChartsGetWorksheetChartTitle(t *testing.T) {
	name := GetMyDoc()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsChartsGetWorksheetChartTitleOpts)
	args.Name = GetMyDoc()
	args.SheetName = GetSheet3()
	args.ChartIndex = 0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsChartsGetWorksheetChartTitle(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsChartsGetWorksheetChartTitle - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsChartsGetWorksheetCharts(t *testing.T) {
	name := GetMyDoc()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsChartsGetWorksheetChartsOpts)
	args.Name = GetMyDoc()
	args.SheetName = GetSheet3()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsChartsGetWorksheetCharts(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsChartsGetWorksheetCharts - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsChartsPostWorksheetChart(t *testing.T) {
	name := GetMyDoc()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsChartsPostWorksheetChartOpts)
	args.Name = GetMyDoc()
	args.SheetName = GetSheet3()
	args.ChartIndex = 0
	args.Chart = new(Chart)
	args.Chart.Name = "Test1"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsChartsPostWorksheetChart(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsChartsPostWorksheetChart - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsChartsPostWorksheetChartLegend(t *testing.T) {
	name := GetMyDoc()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsChartsPostWorksheetChartLegendOpts)
	args.Name = GetMyDoc()
	args.SheetName = GetSheet3()
	args.ChartIndex = 0
	args.Legend = new(Legend)
	args.Legend.Width = 10
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsChartsPostWorksheetChartLegend(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsChartsPostWorksheetChartLegend - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsChartsPostWorksheetChartTitle(t *testing.T) {
	name := GetMyDoc()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsChartsPostWorksheetChartTitleOpts)
	args.Name = GetMyDoc()
	args.SheetName = GetSheet3()
	args.ChartIndex = 0
	args.Title = new(Title)
	args.Title.IsVisible = true
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsChartsPostWorksheetChartTitle(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsChartsPostWorksheetChartTitle - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsChartsPutWorksheetAddChart(t *testing.T) {
	name := GetMyDoc()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsChartsPutWorksheetAddChartOpts)
	args.Name = GetMyDoc()
	args.SheetName = GetSheet4()
	args.ChartType = "Pie"
	args.UpperLeftColumn = 5
	args.UpperLeftRow = 5
	args.LowerRightColumn = 10
	args.LowerRightRow = 10
	args.Area = "C7:D11"
	args.IsVertical = true
	// args.PivotTableName = nil
	// args.PivotTableSheet = nil
	// args.DataLabels = true
	// args.DataLabelsPosition = nil
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsChartsPutWorksheetAddChart(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsChartsPutWorksheetAddChart - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsChartsPutWorksheetChartLegend(t *testing.T) {
	name := GetMyDoc()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsChartsPutWorksheetChartLegendOpts)
	args.Name = GetMyDoc()
	args.SheetName = GetSheet3()
	args.ChartIndex = 0
	args.Folder = GetBaseTest().remoteFolder
	response, httpResponse, err := GetBaseTest().CellsAPI.CellsChartsPutWorksheetChartLegend(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsChartsPutWorksheetChartLegend - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsChartsPutWorksheetChartTitle(t *testing.T) {
	name := GetMyDoc()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsChartsPutWorksheetChartTitleOpts)
	args.Name = GetMyDoc()
	args.SheetName = GetSheet3()
	args.ChartIndex = 0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsChartsPutWorksheetChartTitle(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsChartsPutWorksheetChartTitle - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsConditionalFormattingsDeleteWorksheetConditionalFormatting(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsConditionalFormattingsDeleteWorksheetConditionalFormattingOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Index = 0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsConditionalFormattingsDeleteWorksheetConditionalFormatting(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsConditionalFormattingsDeleteWorksheetConditionalFormatting - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsConditionalFormattingsDeleteWorksheetConditionalFormattingArea(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsConditionalFormattingsDeleteWorksheetConditionalFormattingAreaOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.StartColumn = 1
	args.StartColumn = 1
	args.TotalColumns = (6)
	args.TotalRows = (4)
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsConditionalFormattingsDeleteWorksheetConditionalFormattingArea(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsConditionalFormattingsDeleteWorksheetConditionalFormattingArea - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsConditionalFormattingsDeleteWorksheetConditionalFormattings(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsConditionalFormattingsDeleteWorksheetConditionalFormattingsOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsConditionalFormattingsDeleteWorksheetConditionalFormattings(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsConditionalFormattingsDeleteWorksheetConditionalFormattings - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsConditionalFormattingsGetWorksheetConditionalFormatting(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsConditionalFormattingsGetWorksheetConditionalFormattingOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Index = 0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsConditionalFormattingsGetWorksheetConditionalFormatting(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsConditionalFormattingsGetWorksheetConditionalFormatting - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsConditionalFormattingsGetWorksheetConditionalFormattings(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsConditionalFormattingsGetWorksheetConditionalFormattingsOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsConditionalFormattingsGetWorksheetConditionalFormattings(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsConditionalFormattingsGetWorksheetConditionalFormattings - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsConditionalFormattingsPutWorksheetConditionalFormatting(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsConditionalFormattingsPutWorksheetConditionalFormattingOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.CellArea = GetRange()
	args.FormatCondition = new(FormatCondition)
	args.FormatCondition.Type_ = "CellValue"
	args.FormatCondition.Operator = "Between"
	args.FormatCondition.Formula1 = "v1"
	args.FormatCondition.Formula2 = "v2"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsConditionalFormattingsPutWorksheetConditionalFormatting(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsConditionalFormattingsPutWorksheetConditionalFormatting - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsConditionalFormattingsPutWorksheetFormatCondition(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsConditionalFormattingsPutWorksheetFormatConditionOpts)
	args.Name = GetBook1()
	args.Index = 0
	args.SheetName = GetSheet1()
	args.CellArea = GetRange()
	args.Formula1 = "v1"
	args.Formula2 = "v2"
	args.Type_ = "CellValue"
	args.OperatorType = "Between"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsConditionalFormattingsPutWorksheetFormatCondition(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsConditionalFormattingsPutWorksheetFormatCondition - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsConditionalFormattingsPutWorksheetFormatConditionArea(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsConditionalFormattingsPutWorksheetFormatConditionAreaOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.CellArea = GetRange()
	args.Index = 0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsConditionalFormattingsPutWorksheetFormatConditionArea(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsConditionalFormattingsPutWorksheetFormatConditionArea - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsConditionalFormattingsPutWorksheetFormatConditionCondition(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsConditionalFormattingsPutWorksheetFormatConditionConditionOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Formula1 = "v1"
	args.Formula2 = "v2"
	args.Type_ = "CellValue"
	args.Index = 0
	args.OperatorType = "Between"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsConditionalFormattingsPutWorksheetFormatConditionCondition(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsConditionalFormattingsPutWorksheetFormatConditionCondition - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsDeleteWorksheetColumns(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsDeleteWorksheetColumnsOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.ColumnIndex = 0
	args.Columns = (3)
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsDeleteWorksheetColumns(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsDeleteWorksheetColumns - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsDeleteWorksheetRow(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsDeleteWorksheetRowOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.RowIndex = (3)
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsDeleteWorksheetRow(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsDeleteWorksheetRow - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsDeleteWorksheetRows(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsDeleteWorksheetRowsOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Startrow = 1
	args.TotalRows = (10)
	args.UpdateReference = false
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsDeleteWorksheetRows(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsDeleteWorksheetRows - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsGetCellHtmlString(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsGetCellHtmlStringOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.CellName = "A1"
	args.Folder = GetBaseTest().remoteFolder

	_, httpResponse, err := GetBaseTest().CellsAPI.CellsGetCellHtmlString(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsGetCellHtmlString - %d\n", GetBaseTest().GetTestNumber(), 1)
	}
}

func TestCellsGetWorksheetCell(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsGetWorksheetCellOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.CellOrMethodName = "A1"
	args.Folder = GetBaseTest().remoteFolder

	_, httpResponse, err := GetBaseTest().CellsAPI.CellsGetWorksheetCell(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsGetWorksheetCell - %d\n", GetBaseTest().GetTestNumber(), 1)
	}
}

func TestCellsGetWorksheetCellStyle(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsGetWorksheetCellStyleOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.CellName = "A1"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsGetWorksheetCellStyle(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsGetWorksheetCellStyle - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsGetWorksheetCells(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsGetWorksheetCellsOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Count = (4)
	args.Offest = (10)
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsGetWorksheetCells(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsGetWorksheetCells - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsGetWorksheetColumn(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsGetWorksheetColumnOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Folder = GetBaseTest().remoteFolder
	args.ColumnIndex = 1

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsGetWorksheetColumn(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsGetWorksheetColumn - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsGetWorksheetColumns(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsGetWorksheetColumnsOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsGetWorksheetColumns(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsGetWorksheetColumns - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsGetWorksheetRow(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsGetWorksheetRowOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.RowIndex = 0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsGetWorksheetRow(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsGetWorksheetRow - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsGetWorksheetRows(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsGetWorksheetRowsOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsGetWorksheetRows(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsGetWorksheetRows - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsHypelinksDeleteWorksheetHyperlink(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsHypelinksDeleteWorksheetHyperlinkOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.HyperlinkIndex = 0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsHypelinksDeleteWorksheetHyperlink(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsHypelinksDeleteWorksheetHyperlink - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsHypelinksDeleteWorksheetHyperlinks(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsHypelinksDeleteWorksheetHyperlinksOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsHypelinksDeleteWorksheetHyperlinks(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsHypelinksDeleteWorksheetHyperlinks - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsHypelinksGetWorksheetHyperlink(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsHypelinksGetWorksheetHyperlinkOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.HyperlinkIndex = 0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsHypelinksGetWorksheetHyperlink(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsHypelinksGetWorksheetHyperlink - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsHypelinksGetWorksheetHyperlinks(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsHypelinksGetWorksheetHyperlinksOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsHypelinksGetWorksheetHyperlinks(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsHypelinksGetWorksheetHyperlinks - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsHypelinksPostWorksheetHyperlink(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsHypelinksPostWorksheetHyperlinkOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.HyperlinkIndex = 0
	args.Hyperlink = new(Hyperlink)
	args.Hyperlink.Address = "https://api.aspose.cloud"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsHypelinksPostWorksheetHyperlink(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsHypelinksPostWorksheetHyperlink - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}
func TestCellsHypelinksPutWorksheetHyperlink(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsHypelinksPutWorksheetHyperlinkOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Address = "https://api.aspose.cloud"
	args.FirstColumn = 1
	args.FirstRow = 1
	args.TotalColumns = 1
	args.TotalRows = 1
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsHypelinksPutWorksheetHyperlink(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsHypelinksPutWorksheetHyperlink - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsListObjectsDeleteWorksheetListObject(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsListObjectsDeleteWorksheetListObjectOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet7()
	args.ListObjectIndex = 0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsListObjectsDeleteWorksheetListObject(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsListObjectsDeleteWorksheetListObject - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsListObjectsDeleteWorksheetListObjects(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsListObjectsDeleteWorksheetListObjectsOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet7()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsListObjectsDeleteWorksheetListObjects(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsListObjectsDeleteWorksheetListObjects - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsListObjectsGetWorksheetListObject(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsListObjectsGetWorksheetListObjectOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet7()
	args.Listobjectindex = 0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsListObjectsGetWorksheetListObject(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsListObjectsGetWorksheetListObject - %d\n", GetBaseTest().GetTestNumber(), response)
	}
}

func TestCellsListObjectsGetWorksheetListObjects(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsListObjectsGetWorksheetListObjectsOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet7()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsListObjectsGetWorksheetListObjects(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsListObjectsGetWorksheetListObjects - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsListObjectsPostWorksheetListObject(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsListObjectsPostWorksheetListObjectOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet7()
	args.ListObjectIndex = 0
	args.ListObject = new(ListObject)
	args.ListObject.StartColumn = 10
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsListObjectsPostWorksheetListObject(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsListObjectsPostWorksheetListObject - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsListObjectsPostWorksheetListObjectConvertToRange(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsListObjectsPostWorksheetListObjectConvertToRangeOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet7()
	args.ListObjectIndex = 0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsListObjectsPostWorksheetListObjectConvertToRange(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsListObjectsPostWorksheetListObjectConvertToRange - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsListObjectsPostWorksheetListObjectSortTable(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsListObjectsPostWorksheetListObjectSortTableOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet7()
	args.ListObjectIndex = 0
	args.DataSorter = new(DataSorter)
	args.DataSorter.CaseSensitive = true
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsListObjectsPostWorksheetListObjectSortTable(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsListObjectsPostWorksheetListObjectSortTable - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsListObjectsPostWorksheetListObjectSummarizeWithPivotTable(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsListObjectsPostWorksheetListObjectSummarizeWithPivotTableOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet7()
	args.ListObjectIndex = 0
	args.DestsheetName = "Sheet2"
	args.Request = new(CreatePivotTableRequest)
	args.Request.DestCellName = "C1"
	args.Request.Name = "testp"
	args.Request.PivotFieldColumns = []int64{2}
	args.Request.PivotFieldData = []int64{1}
	args.Request.PivotFieldRows = []int64{0}
	args.Request.SourceData = "=Sheet2!A1:E8"
	args.Request.UseSameSource = true
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsListObjectsPostWorksheetListObjectSummarizeWithPivotTable(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsListObjectsPostWorksheetListObjectSummarizeWithPivotTable - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsListObjectsPutWorksheetListObject(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsListObjectsPutWorksheetListObjectOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet7()
	args.EndColumn = (12)
	args.EndRow = (10)
	args.HasHeaders = true
	args.StartColumn = 1
	args.StartRow = 1
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsListObjectsPutWorksheetListObject(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsListObjectsPutWorksheetListObject - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsOleObjectsDeleteWorksheetOleObject(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsOleObjectsDeleteWorksheetOleObjectOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet6()
	args.OleObjectIndex = 0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsOleObjectsDeleteWorksheetOleObject(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsOleObjectsDeleteWorksheetOleObject - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsOleObjectsDeleteWorksheetOleObjects(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsOleObjectsDeleteWorksheetOleObjectsOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet6()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsOleObjectsDeleteWorksheetOleObjects(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsOleObjectsDeleteWorksheetOleObjects - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsOleObjectsGetWorksheetOleObject(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsOleObjectsGetWorksheetOleObjectOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet6()
	args.ObjectNumber = 0
	args.Folder = GetBaseTest().remoteFolder

	_, httpResponse, err := GetBaseTest().CellsAPI.CellsOleObjectsGetWorksheetOleObject(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsOleObjectsGetWorksheetOleObject - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
}

func TestCellsOleObjectsGetWorksheetOleObjects(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsOleObjectsGetWorksheetOleObjectsOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet6()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsOleObjectsGetWorksheetOleObjects(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsOleObjectsGetWorksheetOleObjects - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsOleObjectsPostUpdateWorksheetOleObject(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsOleObjectsPostUpdateWorksheetOleObjectOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet6()
	args.OleObjectIndex = 0
	args.Ole = new(OleObject)
	args.Ole.Height = (10)
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsOleObjectsPostUpdateWorksheetOleObject(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsOleObjectsPostUpdateWorksheetOleObject - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsOleObjectsPutWorksheetOleObject(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	name = GetWordJpg()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	name = GetOLEDoc()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	args := new(CellsOleObjectsPutWorksheetOleObjectOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet6()
	args.Height = 100
	args.ImageFile = GetBaseTest().remoteFolder + "/" + GetWordJpg()
	args.OleFile = GetBaseTest().remoteFolder + "/" + GetOLEDoc()
	args.UpperLeftColumn = 1
	args.UpperLeftRow = 1
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsOleObjectsPutWorksheetOleObject(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsOleObjectsPutWorksheetOleObject - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPageBreaksDeleteHorizontalPageBreak(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPageBreaksDeleteHorizontalPageBreakOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Index = 0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPageBreaksDeleteHorizontalPageBreak(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPageBreaksDeleteHorizontalPageBreak - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPageBreaksDeleteHorizontalPageBreaks(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPageBreaksDeleteHorizontalPageBreaksOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPageBreaksDeleteHorizontalPageBreaks(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPageBreaksDeleteHorizontalPageBreaks - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPageBreaksDeleteVerticalPageBreak(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPageBreaksDeleteVerticalPageBreakOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Index = 0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPageBreaksDeleteVerticalPageBreak(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPageBreaksDeleteVerticalPageBreak - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPageBreaksDeleteVerticalPageBreaks(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPageBreaksDeleteVerticalPageBreaksOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPageBreaksDeleteVerticalPageBreaks(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPageBreaksDeleteVerticalPageBreaks - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPageBreaksGetHorizontalPageBreak(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPageBreaksGetHorizontalPageBreakOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Index = 0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPageBreaksGetHorizontalPageBreak(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPageBreaksGetHorizontalPageBreak - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPageBreaksGetHorizontalPageBreaks(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPageBreaksGetHorizontalPageBreaksOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPageBreaksGetHorizontalPageBreaks(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPageBreaksGetHorizontalPageBreaks - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPageBreaksGetVerticalPageBreak(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPageBreaksGetVerticalPageBreakOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Index = 0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPageBreaksGetVerticalPageBreak(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPageBreaksGetVerticalPageBreak - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPageBreaksGetVerticalPageBreaks(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPageBreaksGetVerticalPageBreaksOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPageBreaksGetVerticalPageBreaks(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPageBreaksGetVerticalPageBreaks - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPageBreaksPutHorizontalPageBreak(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPageBreaksPutHorizontalPageBreakOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Cellname = "A1"
	args.Column = 1
	args.Row = 1
	args.EndColumn = (2)
	args.StartColumn = (2)
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPageBreaksPutHorizontalPageBreak(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPageBreaksPutHorizontalPageBreak - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPageBreaksPutVerticalPageBreak(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPageBreaksPutVerticalPageBreakOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Cellname = "A1"
	args.Column = 1
	args.Row = 1
	args.StartRow = (2)
	args.EndRow = (2)
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPageBreaksPutVerticalPageBreak(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPageBreaksPutVerticalPageBreak - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPageSetupDeleteHeaderFooter(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPageSetupDeleteHeaderFooterOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPageSetupDeleteHeaderFooter(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPageSetupDeleteHeaderFooter - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPageSetupGetFooter(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPageSetupGetFooterOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPageSetupGetFooter(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPageSetupGetFooter - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPageSetupGetHeader(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPageSetupGetHeaderOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPageSetupGetHeader(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPageSetupGetHeader - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPageSetupGetPageSetup(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPageSetupGetPageSetupOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPageSetupGetPageSetup(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPageSetupGetPageSetup - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPageSetupPostFooter(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPageSetupPostFooterOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.IsFirstPage = true
	args.Section = 0
	args.Script = "test"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPageSetupPostFooter(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPageSetupPostFooter - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPageSetupPostHeader(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPageSetupPostHeaderOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.IsFirstPage = true
	args.Section = 0
	args.Script = "test"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPageSetupPostHeader(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPageSetupPostHeader - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPageSetupPostPageSetup(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPageSetupPostPageSetupOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.PageSetup = new(PageSetup)
	args.PageSetup.BlackAndWhite = true
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPageSetupPostPageSetup(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPageSetupPostPageSetup - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPicturesDeleteWorksheetPicture(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPicturesDeleteWorksheetPictureOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet6()
	args.PictureIndex = 0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPicturesDeleteWorksheetPicture(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPicturesDeleteWorksheetPicture - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPicturesDeleteWorksheetPictures(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPicturesDeleteWorksheetPicturesOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet6()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPicturesDeleteWorksheetPictures(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPicturesDeleteWorksheetPictures - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPicturesGetWorksheetPicture(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPicturesGetWorksheetPictureOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet6()
	args.Folder = GetBaseTest().remoteFolder

	_, httpResponse, err := GetBaseTest().CellsAPI.CellsPicturesGetWorksheetPicture(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPicturesGetWorksheetPicture - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
}

func TestCellsPicturesGetWorksheetPictures(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPicturesGetWorksheetPicturesOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet6()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPicturesGetWorksheetPictures(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPicturesGetWorksheetPictures - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPicturesPostWorksheetPicture(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPicturesPostWorksheetPictureOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet6()
	args.PictureIndex = 0
	args.Picture = new(Picture)
	args.Picture.Height = 100
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPicturesPostWorksheetPicture(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPicturesPostWorksheetPicture - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPicturesPutWorksheetAddPicture(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	name = GetWaterMarkPng()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	args := new(CellsPicturesPutWorksheetAddPictureOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet6()
	args.LowerRightColumn = (10)
	args.LowerRightRow = (10)
	args.UpperLeftColumn = (10)
	args.UpperLeftRow = (10)
	args.PicturePath = GetBaseTest().remoteFolder + "\\" + GetWaterMarkPng()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPicturesPutWorksheetAddPicture(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPicturesPutWorksheetAddPicture - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPivotTablesDeletePivotTableField(t *testing.T) {
	name := GetPivTestFile()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPivotTablesDeletePivotTableFieldOpts)
	args.Name = GetPivTestFile()
	args.SheetName = GetSheet4()
	args.PivotTableIndex = 0
	args.PivotFieldType = "row"
	args.Request = new(PivotTableFieldRequest)
	args.Request.Data = []int64{1}
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPivotTablesDeletePivotTableField(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPivotTablesDeletePivotTableField - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPivotTablesDeleteWorksheetPivotTable(t *testing.T) {
	name := GetPivTestFile()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPivotTablesDeleteWorksheetPivotTableOpts)
	args.Name = GetPivTestFile()
	args.SheetName = GetSheet4()
	args.PivotTableIndex = 0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPivotTablesDeleteWorksheetPivotTable(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPivotTablesDeleteWorksheetPivotTable - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

// func TestCellsPivotTablesDeleteWorksheetPivotTableFilter(t *testing.T) {
// 	name := GetPivTestFile()
// 	if err := GetBaseTest().UploadFile(name); err != nil {
// 		t.Error(err)
// 	}

// 	args := new(CellsPivotTablesDeleteWorksheetPivotTableFilterOpts)
// 	args.Name = GetPivTestFile()
// 	args.SheetName = GetSheet4()
// 	args.PivotTableIndex = 0
// 	args.NeedReCalculate = true
// 	args.FieldIndex = 0
// 	args.Folder = GetBaseTest().remoteFolder

// 	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPivotTablesDeleteWorksheetPivotTableFilter(args)
// 	if err != nil {
// 		t.Error(err)
// 	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
// 		t.Fail()
// 	} else {
// 		fmt.Printf("%d\tTestCellsPivotTablesDeleteWorksheetPivotTableFilter - %d\n", GetBaseTest().GetTestNumber(), response.Code)
// 	}
// }

func TestCellsPivotTablesDeleteWorksheetPivotTableFilters(t *testing.T) {
	name := GetPivTestFile()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPivotTablesDeleteWorksheetPivotTableFiltersOpts)
	args.Name = GetPivTestFile()
	args.SheetName = GetSheet4()
	args.NeedReCalculate = true
	args.PivotTableIndex = 0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPivotTablesDeleteWorksheetPivotTableFilters(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPivotTablesDeleteWorksheetPivotTableFilters - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPivotTablesDeleteWorksheetPivotTables(t *testing.T) {
	name := GetPivTestFile()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPivotTablesDeleteWorksheetPivotTablesOpts)
	args.Name = GetPivTestFile()
	args.SheetName = GetSheet4()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPivotTablesDeleteWorksheetPivotTables(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPivotTablesDeleteWorksheetPivotTables - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPivotTablesGetPivotTableField(t *testing.T) {
	name := GetPivTestFile()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPivotTablesGetPivotTableFieldOpts)
	args.Name = GetPivTestFile()
	args.SheetName = GetSheet4()
	args.PivotTableIndex = 0
	args.PivotFieldIndex = 0
	args.PivotFieldType = "row"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPivotTablesGetPivotTableField(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPivotTablesGetPivotTableField - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPivotTablesGetWorksheetPivotTable(t *testing.T) {
	name := GetPivTestFile()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPivotTablesGetWorksheetPivotTableOpts)
	args.Name = GetPivTestFile()
	args.SheetName = GetSheet4()
	args.PivottableIndex = 0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPivotTablesGetWorksheetPivotTable(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPivotTablesGetWorksheetPivotTable - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPivotTablesGetWorksheetPivotTableFilters(t *testing.T) {
	name := GetPivTestFile()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPivotTablesGetWorksheetPivotTableFiltersOpts)
	args.Name = GetPivTestFile()
	args.SheetName = GetSheet4()
	args.PivotTableIndex = 0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPivotTablesGetWorksheetPivotTableFilters(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPivotTablesGetWorksheetPivotTableFilters - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPivotTablesGetWorksheetPivotTables(t *testing.T) {
	name := GetPivTestFile()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPivotTablesGetWorksheetPivotTablesOpts)
	args.Name = GetPivTestFile()
	args.SheetName = GetSheet4()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPivotTablesGetWorksheetPivotTables(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPivotTablesGetWorksheetPivotTables - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPivotTablesPostPivotTableCellStyle(t *testing.T) {
	name := GetPivTestFile()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPivotTablesPostPivotTableCellStyleOpts)
	args.Name = GetPivTestFile()
	args.SheetName = GetSheet4()
	args.PivotTableIndex = 0
	args.NeedReCalculate = true
	args.Row = 1
	args.Column = 1
	args.Style = new(Style)
	args.Style.Custom = "##.#"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPivotTablesPostPivotTableCellStyle(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPivotTablesPostPivotTableCellStyle - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPivotTablesPostPivotTableFieldHideItem(t *testing.T) {
	name := GetPivTestFile()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPivotTablesPostPivotTableFieldHideItemOpts)
	args.Name = GetPivTestFile()
	args.SheetName = GetSheet4()
	args.PivotTableIndex = 0
	args.PivotFieldType = "Row"
	args.FieldIndex = 0
	args.IsHide = true
	args.NeedReCalculate = true
	args.ItemIndex = 1
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPivotTablesPostPivotTableFieldHideItem(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPivotTablesPostPivotTableFieldHideItem - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPivotTablesPostPivotTableFieldMoveTo(t *testing.T) {
	name := GetPivTestFile()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPivotTablesPostPivotTableFieldMoveToOpts)
	args.Name = GetPivTestFile()
	args.SheetName = GetSheet4()
	args.PivotTableIndex = 0
	args.FieldIndex = 0
	args.From = "Row"
	args.To = "Column"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPivotTablesPostPivotTableFieldMoveTo(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPivotTablesPostPivotTableFieldMoveTo - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}
func TestCellsPivotTablesPostPivotTableUpdatePivotField(t *testing.T) {
	name := GetPivTestFile()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPivotTablesPostPivotTableUpdatePivotFieldOpts)
	args.Name = GetPivTestFile()
	args.SheetName = GetSheet4()
	args.PivotTableIndex = 0
	args.PivotFieldType = "Data"
	args.PivotField = new(PivotField)
	args.PivotField.Number = 5
	args.NeedReCalculate = true
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPivotTablesPostPivotTableUpdatePivotField(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t TestCellsPivotTablesPostPivotTableUpdatePivotField - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPivotTablesPostPivotTableUpdatePivotFields(t *testing.T) {
	name := GetPivTestFile()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPivotTablesPostPivotTableUpdatePivotFieldsOpts)
	args.Name = GetPivTestFile()
	args.SheetName = GetSheet4()
	args.PivotTableIndex = 0
	args.PivotFieldType = "Data"
	args.PivotField = new(PivotField)
	args.PivotField.Number = 5
	args.NeedReCalculate = true
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPivotTablesPostPivotTableUpdatePivotFields(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t TestCellsPivotTablesPostPivotTableUpdatePivotField - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}
func TestCellsPivotTablesPostPivotTableStyle(t *testing.T) {
	name := GetPivTestFile()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPivotTablesPostPivotTableStyleOpts)
	args.Name = GetPivTestFile()
	args.SheetName = GetSheet4()
	args.PivotTableIndex = 0
	args.Style = new(Style)
	args.Style.Custom = "##.#"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPivotTablesPostPivotTableStyle(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPivotTablesPostPivotTableStyle - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPivotTablesPostWorksheetPivotTableCalculate(t *testing.T) {
	name := GetPivTestFile()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPivotTablesPostWorksheetPivotTableCalculateOpts)
	args.Name = GetPivTestFile()
	args.SheetName = GetSheet4()
	args.PivotTableIndex = 0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPivotTablesPostWorksheetPivotTableCalculate(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPivotTablesPostWorksheetPivotTableCalculate - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPivotTablesPostWorksheetPivotTableMove(t *testing.T) {
	name := GetPivTestFile()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPivotTablesPostWorksheetPivotTableMoveOpts)
	args.Name = GetPivTestFile()
	args.SheetName = GetSheet4()
	args.PivotTableIndex = 0
	args.Column = 1
	args.Row = 1
	args.DestCellName = "C10"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPivotTablesPostWorksheetPivotTableMove(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPivotTablesPostWorksheetPivotTableMove - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPivotTablesPutPivotTableField(t *testing.T) {
	name := GetPivTestFile()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPivotTablesPutPivotTableFieldOpts)
	args.Name = GetPivTestFile()
	args.SheetName = GetSheet4()
	args.PivotTableIndex = 0
	args.PivotFieldType = "Row"
	args.Request = new(PivotTableFieldRequest)
	args.Request.Data = []int64{1}
	args.NeedReCalculate = true
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPivotTablesPutPivotTableField(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPivotTablesPutPivotTableField - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPivotTablesPutWorksheetPivotTable(t *testing.T) {
	name := GetPivTestFile()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPivotTablesPutWorksheetPivotTableOpts)
	args.Name = GetPivTestFile()
	args.SheetName = GetSheet4()
	args.DestCellName = "C1"
	args.SourceData = "Sheet1!C6:E13"
	args.TableName = "TestPivot"
	args.UseSameSource = true

	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPivotTablesPutWorksheetPivotTable(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPivotTablesPutWorksheetPivotTable - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPivotTablesPutWorksheetPivotTableFilter(t *testing.T) {
	name := GetPivTestFile()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPivotTablesPutWorksheetPivotTableFilterOpts)
	args.Name = GetPivTestFile()
	args.SheetName = GetSheet4()
	args.NeedReCalculate = true
	args.PivotTableIndex = 0
	args.Filter = new(PivotFilter)
	args.Filter.FieldIndex = 1
	args.Filter.FilterType = "Count"
	args.Filter.AutoFilter = new(AutoFilter)

	filterColumn := new(FilterColumn)
	filterColumn.FilterType = "Top10"
	filterColumn.FieldIndex = 0
	filterColumn.Top10Filter = new(Top10Filter)
	filterColumn.Top10Filter.Items = 1
	filterColumn.Top10Filter.IsPercent = true
	filterColumn.Top10Filter.IsTop = true
	args.Filter.AutoFilter.FilterColumns = []FilterColumn{}
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPivotTablesPutWorksheetPivotTableFilter(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPivotTablesPutWorksheetPivotTableFilter - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}

	args1 := new(CellsPivotTablesGetWorksheetPivotTableFilterOpts)
	args1.Name = GetPivTestFile()
	args1.SheetName = GetSheet4()
	args1.PivotTableIndex = 0
	args1.FilterIndex = 0
	args1.Folder = GetBaseTest().remoteFolder

	response1, httpResponse1, err1 := GetBaseTest().CellsAPI.CellsPivotTablesGetWorksheetPivotTableFilter(args1)
	if err1 != nil {
		t.Error(err1)
	} else if httpResponse1.StatusCode < 200 || httpResponse1.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPivotTablesGetWorksheetPivotTableFilter - %d\n", GetBaseTest().GetTestNumber(), response1.Code)
	}

	args2 := new(CellsPivotTablesDeleteWorksheetPivotTableFilterOpts)
	args2.Name = GetPivTestFile()
	args2.SheetName = GetSheet4()
	args2.PivotTableIndex = 0
	args2.NeedReCalculate = true
	args2.FieldIndex = 0
	args2.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err2 := GetBaseTest().CellsAPI.CellsPivotTablesDeleteWorksheetPivotTableFilter(args2)
	if err2 != nil {
		t.Error(err2)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPivotTablesDeleteWorksheetPivotTableFilter - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsSparklineGroupGPPD(t *testing.T) {
	name := GetPivTestFile()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	argsGets := new(CellsSparklineGroupsGetWorksheetSparklineGroupsOpts)
	argsGets.Name = GetPivTestFile()
	argsGets.SheetName = GetSheet1()
	argsGets.Folder = GetBaseTest().remoteFolder

	responseGets, httpResponseGets, errGets := GetBaseTest().CellsAPI.CellsSparklineGroupsGetWorksheetSparklineGroups(argsGets)
	if errGets != nil {
		t.Error(errGets)
	} else if httpResponseGets.StatusCode < 200 || httpResponseGets.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t CellsSparklineGroupsGetWorksheetSparklineGroups - %d\n", GetBaseTest().GetTestNumber(), responseGets.Code)
	}

	argsPut := new(CellsSparklineGroupsPutWorksheetSparklineGroupOpts)
	argsPut.Name = GetPivTestFile()
	argsPut.SheetName = GetSheet1()
	argsPut.DataRange = "C6:E13"
	argsPut.IsVertical = false
	argsPut.LocationRange = "G6:G13"
	argsPut.Type_ = "Line"
	argsPut.Folder = GetBaseTest().remoteFolder

	responsePut, httpResponsePut, errPut := GetBaseTest().CellsAPI.CellsSparklineGroupsPutWorksheetSparklineGroup(argsPut)
	if errPut != nil {
		t.Error(errPut)
	} else if httpResponsePut.StatusCode < 200 || httpResponsePut.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t CellsSparklineGroupsPutWorksheetSparklineGroup - %d\n", GetBaseTest().GetTestNumber(), responsePut.Code)
	}

	argsPost := new(CellsSparklineGroupsPostWorksheetSparklineGroupOpts)
	argsPost.Name = GetPivTestFile()
	argsPost.SheetName = GetSheet1()
	argsPost.SparklineGroupIndex = 0
	argsPost.SparklineGroup = new(SparklineGroup)
	argsPost.SparklineGroup.DisplayHidden = true
	argsPost.Folder = GetBaseTest().remoteFolder

	responsePost, httpResponsePost, errPost := GetBaseTest().CellsAPI.CellsSparklineGroupsPostWorksheetSparklineGroup(argsPost)
	if errPost != nil {
		t.Error(errPost)
	} else if httpResponsePost.StatusCode < 200 || httpResponsePost.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t CellsSparklineGroupsPostWorksheetSparklineGroup - %d\n", GetBaseTest().GetTestNumber(), responsePost.Code)
	}

	argsGet := new(CellsSparklineGroupsGetWorksheetSparklineGroupOpts)
	argsGet.Name = GetPivTestFile()
	argsGet.SheetName = GetSheet1()
	argsGet.SparklineGroupIndex = 0
	argsGet.Folder = GetBaseTest().remoteFolder

	responseGet, httpResponseGet, errGet := GetBaseTest().CellsAPI.CellsSparklineGroupsGetWorksheetSparklineGroup(argsGet)
	if errGet != nil {
		t.Error(errGet)
	} else if httpResponseGet.StatusCode < 200 || httpResponseGet.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t CellsSparklineGroupsDeleteWorksheetSparklineGroup - %d\n", GetBaseTest().GetTestNumber(), responseGet.Code)
	}

	argsDelete := new(CellsSparklineGroupsDeleteWorksheetSparklineGroupOpts)
	argsDelete.Name = GetPivTestFile()
	argsDelete.SheetName = GetSheet1()
	argsDelete.SparklineGroupIndex = 0
	argsDelete.Folder = GetBaseTest().remoteFolder

	responseDelete, httpResponseDelete, errDelete := GetBaseTest().CellsAPI.CellsSparklineGroupsDeleteWorksheetSparklineGroup(argsDelete)
	if errDelete != nil {
		t.Error(errDelete)
	} else if httpResponseDelete.StatusCode < 200 || httpResponseDelete.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t CellsSparklineGroupsDeleteWorksheetSparklineGroup - %d\n", GetBaseTest().GetTestNumber(), responseDelete.Code)
	}

	argsDeletes := new(CellsSparklineGroupsDeleteWorksheetSparklineGroupsOpts)
	argsDeletes.Name = GetPivTestFile()
	argsDeletes.SheetName = GetSheet1()
	argsDeletes.Folder = GetBaseTest().remoteFolder

	responseDeletes, httpResponseDeletes, errDeletes := GetBaseTest().CellsAPI.CellsSparklineGroupsDeleteWorksheetSparklineGroups(argsDeletes)
	if errDeletes != nil {
		t.Error(errDeletes)
	} else if httpResponseDeletes.StatusCode < 200 || httpResponseDeletes.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t CellsSparklineGroupsDeleteWorksheetSparklineGroup - %d\n", GetBaseTest().GetTestNumber(), responseDeletes.Code)
	}

}

func TestCellsPostCellCalculate(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPostCellCalculateOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.CellName = "A1"
	args.Options = new(CalculationOptions)
	args.Options.IgnoreError = true
	args.Options.Recursive = true
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPostCellCalculate(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostCellCalculate - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPostCellCharacters(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPostCellCharactersOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet2()
	args.CellName = "G8"

	fontSetting := new(FontSetting)
	fontSetting.Font = new(Font)
	fontSetting.Font.IsBold = true
	fontSetting.Length = 2
	fontSetting.StartIndex = 0

	args.Options = &[]FontSetting{*fontSetting}

	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPostCellCharacters(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostCellCharacters - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPostClearContents(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPostClearContentsOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.StartColumn = 1
	args.StartRow = 1
	args.EndColumn = (10)
	args.EndRow = (10)
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPostClearContents(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostClearContents - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPostClearFormats(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPostClearFormatsOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Range_ = GetRange()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPostClearFormats(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostClearFormats - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPostColumnStyle(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPostColumnStyleOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.ColumnIndex = 1
	args.Style = new(Style)
	args.Style.Custom = "##.#"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPostColumnStyle(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostColumnStyle - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPostCopyCellIntoCell(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPostCopyCellIntoCellOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.DestCellName = "C1"
	args.Worksheet = GetSheet2()
	args.Cellname = "A1"
	args.Row = 1
	args.Column = 1
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPostCopyCellIntoCell(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostCopyCellIntoCell - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPostCopyWorksheetColumns(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPostCopyWorksheetColumnsOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.ColumnNumber = 1
	args.DestinationColumnIndex = (21)
	args.SourceColumnIndex = 1
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPostCopyWorksheetColumns(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostCopyWorksheetColumns - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPostCopyWorksheetRows(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPostCopyWorksheetRowsOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.DestinationRowIndex = 1
	args.SourceRowIndex = (10)
	args.RowNumber = (3)
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPostCopyWorksheetRows(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostCopyWorksheetRows - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPostGroupWorksheetColumns(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPostGroupWorksheetColumnsOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.FirstIndex = 1
	args.LastIndex = (10)
	args.Hide = true
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPostGroupWorksheetColumns(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostGroupWorksheetColumns - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPostGroupWorksheetRows(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPostGroupWorksheetRowsOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.FirstIndex = 1
	args.LastIndex = (10)
	args.Hide = true
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPostGroupWorksheetRows(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostGroupWorksheetRows - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPostHideWorksheetColumns(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPostHideWorksheetColumnsOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.StartColumn = 1
	args.TotalColumns = (10)
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPostHideWorksheetColumns(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostHideWorksheetColumns - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPostHideWorksheetRows(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPostHideWorksheetRowsOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Startrow = 1
	args.TotalRows = (10)
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPostHideWorksheetRows(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostHideWorksheetRows - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPostRowStyle(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPostRowStyleOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.RowIndex = (10)
	args.Style = new(Style)
	args.Style.Custom = "##.#"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPostRowStyle(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostRowStyle - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPostSetCellHtmlString(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPostSetCellHtmlStringOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.CellName = "A1"
	args.HtmlString = "Test"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPostSetCellHtmlString(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostSetCellHtmlString - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPostSetCellRangeValue(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPostSetCellRangeValueOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Cellarea = GetRange()
	args.Type_ = "String"
	args.Value = "123"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPostSetCellRangeValue(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostSetCellRangeValue - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPostSetWorksheetColumnWidth(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPostSetWorksheetColumnWidthOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.ColumnIndex = 0
	args.Width = 64.0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPostSetWorksheetColumnWidth(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostSetWorksheetColumnWidth - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPostUngroupWorksheetColumns(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPostUngroupWorksheetColumnsOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.FirstIndex = 1
	args.LastIndex = (10)
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPostUngroupWorksheetColumns(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostUngroupWorksheetColumns - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPostUngroupWorksheetRows(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPostUngroupWorksheetRowsOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.FirstIndex = 1
	args.IsAll = true
	args.LastIndex = (10)
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPostUngroupWorksheetRows(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostUngroupWorksheetRows - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPostUnhideWorksheetColumns(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPostUnhideWorksheetColumnsOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Startcolumn = 1
	args.TotalColumns = (9)
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPostUnhideWorksheetColumns(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostUnhideWorksheetColumns - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPostUnhideWorksheetRows(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPostUnhideWorksheetRowsOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Startrow = 0
	args.TotalRows = (10)
	args.Height = 19.0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPostUnhideWorksheetRows(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostUnhideWorksheetRows - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPostUpdateWorksheetCellStyle(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPostUpdateWorksheetCellStyleOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.CellName = "A1"
	args.Style = new(Style)
	args.Style.Custom = "##.#"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPostUpdateWorksheetCellStyle(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostUpdateWorksheetCellStyle - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPostUpdateWorksheetRangeStyle(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPostUpdateWorksheetRangeStyleOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Range_ = GetRange()
	args.Style = new(Style)
	args.Style.Custom = "##.#"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPostUpdateWorksheetRangeStyle(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostUpdateWorksheetRangeStyle - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPostUpdateWorksheetRow(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPostUpdateWorksheetRowOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.RowIndex = 1
	args.Height = 64.0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPostUpdateWorksheetRow(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostUpdateWorksheetRow - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPostWorksheetCellSetValue(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPostWorksheetCellSetValueOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.CellName = "A1"
	args.Formula = "Now()"
	args.Type_ = "String"
	args.Value = "1234"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPostWorksheetCellSetValue(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostWorksheetCellSetValue - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPostWorksheetMerge(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPostWorksheetMergeOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.StartColumn = 0
	args.StartRow = 0
	args.TotalColumns = (10)
	args.TotalRows = (10)
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPostWorksheetMerge(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostWorksheetMerge - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPostWorksheetUnmerge(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPostWorksheetUnmergeOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.StartColumn = 1
	args.StartRow = 1
	args.TotalColumns = (10)
	args.TotalRows = (10)
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPostWorksheetUnmerge(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostWorksheetUnmerge - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPropertiesDeleteDocumentProperties(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPropertiesDeleteDocumentPropertiesOpts)
	args.Name = GetBook1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPropertiesDeleteDocumentProperties(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPropertiesDeleteDocumentProperties - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPropertiesDeleteDocumentProperty(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPropertiesDeleteDocumentPropertyOpts)
	args.Name = GetBook1()
	args.PropertyName = "Author"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPropertiesDeleteDocumentProperty(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPropertiesDeleteDocumentProperty - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPropertiesGetDocumentProperties(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPropertiesGetDocumentPropertiesOpts)
	args.Name = GetBook1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPropertiesGetDocumentProperties(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPropertiesGetDocumentProperties - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPropertiesGetDocumentProperty(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPropertiesGetDocumentPropertyOpts)
	args.Name = GetBook1()
	args.PropertyName = "Author"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPropertiesGetDocumentProperty(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPropertiesGetDocumentProperty - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPropertiesPutDocumentProperty(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPropertiesPutDocumentPropertyOpts)
	args.Name = GetBook1()
	args.PropertyName = "author"
	args.Property = new(CellsDocumentProperty)
	args.Property.Name = "author"
	args.Property.Value = "roy"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPropertiesPutDocumentProperty(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPropertiesPutDocumentProperty - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPutInsertWorksheetColumns(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPutInsertWorksheetColumnsOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.ColumnIndex = 1
	args.Columns = (10)
	args.UpdateReference = true
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPutInsertWorksheetColumns(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPutInsertWorksheetColumns - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPutInsertWorksheetRow(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPutInsertWorksheetRowOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.RowIndex = 0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPutInsertWorksheetRow(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPutInsertWorksheetRow - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPutInsertWorksheetRows(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPutInsertWorksheetRowsOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Startrow = 1
	args.TotalRows = (10)
	args.UpdateReference = true
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsPutInsertWorksheetRows(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPutInsertWorksheetRows - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsRangesGetWorksheetCellsRangeValue(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsRangesGetWorksheetCellsRangeValueOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.ColumnCount = (3)
	args.FirstColumn = 1
	args.FirstRow = 1
	args.RowCount = (3)
	args.Namerange = GetRange()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsRangesGetWorksheetCellsRangeValue(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsRangesGetWorksheetCellsRangeValue - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsShapesDeleteWorksheetShape(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsShapesDeleteWorksheetShapeOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsShapesDeleteWorksheetShape(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsShapesDeleteWorksheetShape - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsShapesDeleteWorksheetShapes(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsShapesDeleteWorksheetShapesOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsShapesDeleteWorksheetShapes(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsShapesDeleteWorksheetShapes - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsShapesGetWorksheetShape(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsShapesGetWorksheetShapeOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsShapesGetWorksheetShape(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsShapesGetWorksheetShape - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsShapesGetWorksheetShapes(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsShapesGetWorksheetShapesOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsShapesGetWorksheetShapes(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsShapesGetWorksheetShapes - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsShapesPostWorksheetShape(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsShapesPostWorksheetShapeOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet4()
	args.Shapeindex = 0
	args.Dto = new(Shape)
	args.Dto.Bottom = 10
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsShapesPostWorksheetShape(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsShapesPostWorksheetShape - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsShapesPutWorksheetShape(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsShapesPutWorksheetShapeOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.ShapeDTO = nil
	args.DrawingType = "arc"
	args.UpperLeftColumn = 1
	args.UpperLeftRow = 1
	args.Left = 1
	args.Top = 1
	args.Height = (100)
	args.Width = (100)
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsShapesPutWorksheetShape(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsShapesPutWorksheetShape - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsRangesPostWorksheetCellsRangeColumnWidth(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsRangesPostWorksheetCellsRangeColumnWidthOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Range_ = new(ModelRange)
	args.Range_.Worksheet = GetSheet1()
	args.Range_.FirstColumn = 1
	args.Range_.FirstRow = 1
	args.Range_.ColumnWidth = 100.0
	args.Range_.RowHeight = 100.0
	args.Range_.RowCount = (100)
	args.Range_.ColumnCount = (100)
	args.Value = 123.0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsRangesPostWorksheetCellsRangeColumnWidth(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsRangesPostWorksheetCellsRangeColumnWidth - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsRangesPostWorksheetCellsRangeMerge(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsRangesPostWorksheetCellsRangeMergeOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Range_ = new(ModelRange)
	args.Range_.Worksheet = GetSheet1()
	args.Range_.FirstColumn = 1
	args.Range_.FirstRow = 1
	args.Range_.ColumnWidth = 100.0
	args.Range_.RowHeight = 100.0
	args.Range_.RowCount = (100)
	args.Range_.ColumnCount = (100)
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsRangesPostWorksheetCellsRangeMerge(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsRangesPostWorksheetCellsRangeMerge - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsRangesPostWorksheetCellsRangeMoveTo(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsRangesPostWorksheetCellsRangeMoveToOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.DestColumn = (10)
	args.DestRow = (10)
	args.Range_ = new(ModelRange)
	args.Range_.Worksheet = GetSheet1()
	args.Range_.FirstColumn = 1
	args.Range_.FirstRow = 1
	args.Range_.ColumnWidth = 100.0
	args.Range_.RowHeight = 100.0
	args.Range_.RowCount = (100)
	args.Range_.ColumnCount = (100)

	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsRangesPostWorksheetCellsRangeMoveTo(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsRangesPostWorksheetCellsRangeMoveTo - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsRangesPostWorksheetCellsRangeOutlineBorder(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsRangesPostWorksheetCellsRangeOutlineBorderOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.RangeOperate = new(RangeSetOutlineBorderRequest)
	args.RangeOperate.Range_ = new(ModelRange)
	args.RangeOperate.Range_.Worksheet = GetSheet1()
	args.RangeOperate.Range_.FirstColumn = 1
	args.RangeOperate.Range_.FirstRow = 1
	args.RangeOperate.Range_.ColumnWidth = 100.0
	args.RangeOperate.Range_.RowHeight = 100.0
	args.RangeOperate.Range_.RowCount = (100)
	args.RangeOperate.Range_.ColumnCount = (100)
	args.RangeOperate.BorderColor = new(Color)
	args.RangeOperate.BorderColor.A = 255
	args.RangeOperate.BorderColor.B = 255
	args.RangeOperate.BorderColor.G = 0
	args.RangeOperate.BorderColor.R = 255
	args.RangeOperate.BorderEdge = "LeftBorder"
	args.RangeOperate.BorderStyle = "Dotted"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsRangesPostWorksheetCellsRangeOutlineBorder(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsRangesPostWorksheetCellsRangeOutlineBorder - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsRangesPostWorksheetCellsRangeRowHeight(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsRangesPostWorksheetCellsRangeRowHeightOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Range_ = new(ModelRange)
	args.Range_.Worksheet = GetSheet1()
	args.Range_.FirstColumn = 1
	args.Range_.FirstRow = 1
	args.Range_.ColumnWidth = 100.0
	args.Range_.RowHeight = 100.0
	args.Range_.RowCount = (100)
	args.Range_.ColumnCount = (100)
	args.Value = 124.0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsRangesPostWorksheetCellsRangeRowHeight(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsRangesPostWorksheetCellsRangeRowHeight - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsRangesPostWorksheetCellsRangeStyle(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsRangesPostWorksheetCellsRangeStyleOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.RangeOperate = new(RangeSetStyleRequest)
	args.RangeOperate.Range_ = new(ModelRange)
	args.RangeOperate.Range_.Worksheet = GetSheet1()
	args.RangeOperate.Range_.FirstColumn = 1
	args.RangeOperate.Range_.FirstRow = 1
	args.RangeOperate.Range_.ColumnWidth = 100.0
	args.RangeOperate.Range_.RowHeight = 100.0
	args.RangeOperate.Range_.RowCount = (100)
	args.RangeOperate.Range_.ColumnCount = (100)
	args.RangeOperate.Style = new(Style)
	args.RangeOperate.Style.Custom = "##.#"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsRangesPostWorksheetCellsRangeStyle(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsRangesPostWorksheetCellsRangeStyle - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsRangesPostWorksheetCellsRangeUnmerge(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsRangesPostWorksheetCellsRangeUnmergeOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Range_ = new(ModelRange)
	args.Range_.Worksheet = GetSheet1()
	args.Range_.FirstColumn = 1
	args.Range_.FirstRow = 1
	args.Range_.ColumnWidth = 100.0
	args.Range_.RowHeight = 100.0
	args.Range_.RowCount = (100)
	args.Range_.ColumnCount = (100)
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsRangesPostWorksheetCellsRangeUnmerge(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsRangesPostWorksheetCellsRangeUnmerge - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsRangesPostWorksheetCellsRangeValue(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsRangesPostWorksheetCellsRangeValueOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Value = "12.0"
	args.IsConverted = true
	args.SetStyle = true
	args.Range_ = new(ModelRange)
	args.Range_.Worksheet = GetSheet1()
	args.Range_.FirstColumn = 1
	args.Range_.FirstRow = 1
	args.Range_.ColumnWidth = 100.0
	args.Range_.RowHeight = 100.0
	args.Range_.RowCount = (100)
	args.Range_.ColumnCount = (100)
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsRangesPostWorksheetCellsRangeValue(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsRangesPostWorksheetCellsRangeValue - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsRangesPostWorksheetCellsRanges(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsRangesPostWorksheetCellsRangesOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.RangeOperate = new(RangeCopyRequest)
	args.RangeOperate.Operate = "copydata"
	args.RangeOperate.Source = new(ModelRange)
	args.RangeOperate.Source.Worksheet = GetSheet1()
	args.RangeOperate.Source.FirstColumn = 1
	args.RangeOperate.Source.FirstRow = 1
	args.RangeOperate.Source.ColumnWidth = 100.0
	args.RangeOperate.Source.RowHeight = 100.0
	args.RangeOperate.Source.RowCount = (100)
	args.RangeOperate.Source.ColumnCount = (100)
	args.RangeOperate.Target = new(ModelRange)
	args.RangeOperate.Target.Worksheet = GetSheet1()
	args.RangeOperate.Target.FirstColumn = 1
	args.RangeOperate.Target.FirstRow = 1
	args.RangeOperate.Target.ColumnWidth = 100.0
	args.RangeOperate.Target.RowHeight = 100.0
	args.RangeOperate.Target.RowCount = (100)
	args.RangeOperate.Target.ColumnCount = (100)
	args.RangeOperate.PasteOptions = new(PasteOptions)
	args.RangeOperate.PasteOptions.OnlyVisibleCells = true
	args.RangeOperate.PasteOptions.SkipBlanks = true
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsRangesPostWorksheetCellsRanges(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsRangesPostWorksheetCellsRanges - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsRangesDeleteWorksheetCellsRange(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsRangesDeleteWorksheetCellsRangeOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Range_ = "A1:B4"
	args.Shift = "Up"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsRangesDeleteWorksheetCellsRange(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsRangesDeleteWorksheetCellsRange - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsRangesPutWorksheetCellsRange(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsRangesPutWorksheetCellsRangeOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Range_ = "A1:B4"
	args.Shift = "Down"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsRangesPutWorksheetCellsRange(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsRangesPutWorksheetCellsRange - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsSaveAsPostDocumentSaveAs(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsSaveAsPostDocumentSaveAsOpts)
	args.Name = GetBook1()
	args.SaveOptions = nil
	args.Newfilename = GetBaseTest().remoteFolder + "/newfilego.pdf"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsSaveAsPostDocumentSaveAs(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsSaveAsPostDocumentSaveAs - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorkbookDeleteDecryptDocument(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorkbookDeleteDecryptDocumentOpts)
	args.Name = GetBook1()
	args.Encryption = new(WorkbookEncryptionRequest)
	args.Encryption.EncryptionType = "XOR"
	args.Encryption.KeyLength = 128
	args.Encryption.Password = "123456"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookDeleteDecryptDocument(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorkbookDeleteDecryptDocument - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorkbookDeleteDocumentUnprotectFromChanges(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorkbookDeleteDocumentUnprotectFromChangesOpts)
	args.Name = GetBook1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookDeleteDocumentUnprotectFromChanges(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorkbookDeleteDocumentUnprotectFromChanges - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorkbookDeleteUnprotectDocument(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorkbookDeleteUnprotectDocumentOpts)
	args.Name = GetBook1()
	args.Protection = new(WorkbookProtectionRequest)
	args.Protection.Password = "123456"
	args.Protection.ProtectionType = "All"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookDeleteUnprotectDocument(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorkbookDeleteUnprotectDocument - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorkbookDeleteWorkbookName(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorkbookDeleteWorkbookNameOpts)
	args.Name = GetBook1()
	args.NameName = "Name_2"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookDeleteWorkbookName(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorkbookDeleteWorkbookName - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorkbookDeleteWorkbookNames(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorkbookDeleteWorkbookNamesOpts)
	args.Name = GetBook1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookDeleteWorkbookNames(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorkbookDeleteWorkbookNames - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorkbookGetWorkbook(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorkbookGetWorkbookOpts)
	args.Name = GetBook1()
	args.Format = "PDF"
	args.Folder = GetBaseTest().remoteFolder

	_, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookGetWorkbook(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorkbookGetWorkbook - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
}
func TestCellsWorkbookGetWorkbookToOtherStorage(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorkbookGetWorkbookOpts)
	args.Name = GetBook1()
	args.Format = "PDF"
	args.Folder = GetBaseTest().remoteFolder
	args.OutPath = "Freeing/TestCellsWorkbookGetWorkbookToOtherStorage.pdf"
	args.OutStorageName = "DropBox"
	_, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookGetWorkbook(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorkbookGetWorkbook - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
}

func TestCellsWorkbookGetWorkbookDefaultStyle(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorkbookGetWorkbookDefaultStyleOpts)
	args.Name = GetBook1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookGetWorkbookDefaultStyle(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorkbookGetWorkbookDefaultStyle - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorkbookGetWorkbookName(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorkbookGetWorkbookNameOpts)
	args.Name = GetBook1()
	args.NameName = "Name_2"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookGetWorkbookName(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorkbookGetWorkbookName - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorkbookGetWorkbookNameValue(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorkbookGetWorkbookNameValueOpts)
	args.Name = GetBook1()
	args.NameName = "Name_2"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookGetWorkbookNameValue(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorkbookGetWorkbookNameValue - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorkbookGetWorkbookNames(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorkbookGetWorkbookNamesOpts)
	args.Name = GetBook1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookGetWorkbookNames(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorkbookGetWorkbookNames - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorkbookGetWorkbookSettings(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorkbookGetWorkbookSettingsOpts)
	args.Name = GetBook1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookGetWorkbookSettings(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorkbookGetWorkbookSettings - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorkbookGetWorkbookTextItems(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorkbookGetWorkbookTextItemsOpts)
	args.Name = GetBook1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookGetWorkbookTextItems(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorkbookGetWorkbookTextItems - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorkbookPostAutofitWorkbookRows(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorkbookPostAutofitWorkbookRowsOpts)
	args.Name = GetBook1()
	args.StartRow = 0
	args.EndRow = (10)
	args.OnlyAuto = true
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookPostAutofitWorkbookRows(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorkbookPostAutofitWorkbookRows - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorkbookPostAutofitWorkbookColumns(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorkbookPostAutofitWorkbookColumnsOpts)
	args.Name = GetBook1()
	args.StartColumn = 0
	args.EndColumn = (10)
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookPostAutofitWorkbookColumns(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorkbookPostAutofitWorkbookColumns - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorkbookPostEncryptDocument(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorkbookPostEncryptDocumentOpts)
	args.Name = GetBook1()
	args.Encryption = new(WorkbookEncryptionRequest)
	args.Encryption.EncryptionType = "XOR"
	args.Encryption.KeyLength = 128
	args.Encryption.Password = "123456"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookPostEncryptDocument(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorkbookPostEncryptDocument - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorkbookPostProtectDocument(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorkbookPostProtectDocumentOpts)
	args.Name = GetBook1()
	args.Protection = new(WorkbookProtectionRequest)
	args.Protection.Password = "123456"
	args.Protection.ProtectionType = "All"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookPostProtectDocument(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorkbookPostProtectDocument - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorkbookPostWorkbookCalculateFormula(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorkbookPostWorkbookCalculateFormulaOpts)
	args.Name = GetBook1()
	args.IgnoreError = true
	args.Options = new(CalculationOptions)
	args.Options.Recursive = true
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookPostWorkbookCalculateFormula(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorkbookPostWorkbookCalculateFormula - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorkbookPostWorkbookGetSmartMarkerResult(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	name = GetReportDataXml()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	args := new(CellsWorkbookPostWorkbookGetSmartMarkerResultOpts)
	args.Name = GetBook1()
	args.OutPath = "GoTest/NewSmartMark.xlsx"
	args.XmlFile = "GoTest/" + GetReportDataXml()
	args.Folder = GetBaseTest().remoteFolder

	_, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookPostWorkbookGetSmartMarkerResult(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorkbookPostWorkbookGetSmartMarkerResult - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
}
func TestCellsWorkbookPostWorkbookGetSmartMarkerResultToOtherStorage(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	name = GetReportDataXml()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	args := new(CellsWorkbookPostWorkbookGetSmartMarkerResultOpts)
	args.Name = GetBook1()
	args.OutPath = "GoTest/NewSmartMark.xlsx"
	args.XmlFile = "GoTest/" + GetReportDataXml()
	args.Folder = GetBaseTest().remoteFolder
	args.OutStorageName = "DropBox"

	_, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookPostWorkbookGetSmartMarkerResult(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorkbookPostWorkbookGetSmartMarkerResult - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
}

func TestCellsWorkbookPostWorkbookSettings(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorkbookPostWorkbookSettingsOpts)
	args.Name = GetBook1()
	args.Settings = new(WorkbookSettings)
	args.Settings.AutoCompressPictures = true
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookPostWorkbookSettings(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorkbookPostWorkbookSettings - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorkbookPostWorkbookSplit(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorkbookPostWorkbookSplitOpts)
	args.Name = GetBook1()
	args.Format = "png"
	args.From = 1
	args.To = 5
	args.HorizontalResolution = 96
	args.VerticalResolution = 96
	args.Folder = GetBaseTest().remoteFolder
	args.OutFolder = GetBaseTest().remoteFolder
	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookPostWorkbookSplit(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorkbookPostWorkbookSplit - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}
func TestCellsWorkbookPostWorkbookSplitToOtherStorage(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorkbookPostWorkbookSplitOpts)
	args.Name = GetBook1()
	args.Format = "png"
	args.From = 1
	args.To = 5
	args.HorizontalResolution = 96
	args.VerticalResolution = 96
	args.Folder = GetBaseTest().remoteFolder
	args.OutFolder = GetBaseTest().remoteFolder
	args.OutStorageName = "DropBox"
	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookPostWorkbookSplit(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorkbookPostWorkbookSplit - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorkbookPostWorkbooksMerge(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	name = GetMyDoc()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	args := new(CellsWorkbookPostWorkbooksMergeOpts)
	args.Name = GetBook1()
	args.MergeWith = GetMyDoc()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookPostWorkbooksMerge(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorkbookPostWorkbooksMerge - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}
func TestCellsWorkbookPostWorkbooksMergeOtherStorage(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	name = GetMyDoc()
	if err := GetBaseTest().UploadFileToOtherStorage(name, "DropBox"); err != nil {
		t.Error(err)
	}
	args := new(CellsWorkbookPostWorkbooksMergeOpts)
	args.Name = GetBook1()
	args.MergeWith = GetMyDoc()
	args.Folder = GetBaseTest().remoteFolder
	args.MergedStorageName = "DropBox"

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookPostWorkbooksMerge(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorkbookPostWorkbooksMerge - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}
func TestCellsWorkbookPostWorkbooksTextReplace(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorkbookPostWorkbooksTextReplaceOpts)
	args.Name = GetBook1()
	args.NewValue = "test"
	args.OldValue = "abc"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookPostWorkbooksTextReplace(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorkbookPostWorkbooksTextReplace - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorkbookPostWorkbooksTextSearch(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorkbookPostWorkbooksTextSearchOpts)
	args.Name = GetBook1()
	args.Text = "test"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookPostWorkbooksTextSearch(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorkbookPostWorkbooksTextSearch - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorkbookPutDocumentProtectFromChanges(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorkbookPutDocumentProtectFromChangesOpts)
	args.Name = GetBook1()
	args.Password = new(PasswordRequest)
	args.Password.Password = "123456"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookPutDocumentProtectFromChanges(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorkbookPutDocumentProtectFromChanges - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorkbookPutWorkbookCreate(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	name = GetReportDataXml()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	args := new(CellsWorkbookPutWorkbookCreateOpts)
	args.Name = GetNewXLSXFile()
	args.DataFile = ""
	args.TemplateFile = "GoTest/" + GetBook1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookPutWorkbookCreate(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorkbookPutWorkbookCreate - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorksheetValidationsDeleteWorksheetValidation(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetValidationsDeleteWorksheetValidationOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.ValidationIndex = 0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetValidationsDeleteWorksheetValidation(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetValidationsDeleteWorksheetValidation - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorksheetValidationsDeleteWorksheetValidations(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetValidationsDeleteWorksheetValidationsOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetValidationsDeleteWorksheetValidations(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetValidationsDeleteWorksheetValidations - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorksheetValidationsGetWorksheetValidations(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetValidationsGetWorksheetValidationsOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetValidationsGetWorksheetValidations(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetValidationsGetWorksheetValidations - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorksheetValidationsPostWorksheetValidation(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetValidationsPostWorksheetValidationOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.ValidationIndex = 0
	args.Validation = new(Validation)
	args.Validation.Formula1 = "=(OR(A1=\"Yes\",A1=\"No\"))"
	args.Validation.Type_ = "Custom"
	args.Validation.IgnoreBlank = true
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetValidationsPostWorksheetValidation(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetValidationsPostWorksheetValidation - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorksheetValidationsPutWorksheetValidation(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetValidationsPutWorksheetValidationOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Range_ = GetRange()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetValidationsPutWorksheetValidation(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetValidationsPutWorksheetValidation - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorksheetsDeleteUnprotectWorksheet(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsDeleteUnprotectWorksheetOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.ProtectParameter = new(ProtectSheetParameter)
	args.ProtectParameter.Password = "123456"
	args.ProtectParameter.ProtectionType = "All"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsDeleteUnprotectWorksheet(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetsDeleteUnprotectWorksheet - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorksheetsDeleteWorksheet(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsDeleteWorksheetOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsDeleteWorksheet(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetsDeleteWorksheet - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorksheetsDeleteWorksheetBackground(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsDeleteWorksheetBackgroundOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsDeleteWorksheetBackground(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetsDeleteWorksheetBackground - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorksheetsDeleteWorksheetComment(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsDeleteWorksheetCommentOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.CellName = "A1"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsDeleteWorksheetComment(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetsDeleteWorksheetComment - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorksheetsDeleteWorksheetComments(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsDeleteWorksheetCommentsOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsDeleteWorksheetComments(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetsDeleteWorksheetComments - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorksheetsDeleteWorksheetFreezePanes(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsDeleteWorksheetFreezePanesOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Row = 1
	args.Column = 1
	args.FreezedRows = (10)
	args.FreezedColumns = (10)
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsDeleteWorksheetFreezePanes(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetsDeleteWorksheetFreezePanes - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorksheetsGetNamedRanges(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsGetNamedRangesOpts)
	args.Name = GetBook1()

	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsGetNamedRanges(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetsGetNamedRanges - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorksheetsGetWorksheet(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsGetWorksheetOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Format = "pdf"
	args.Folder = GetBaseTest().remoteFolder

	_, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsGetWorksheet(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetsGetWorksheet - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
}

func TestCellsWorksheetsGetWorksheetForArea(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsGetWorksheetOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Area = "B3:K8"
	args.Format = "pdf"
	args.Folder = GetBaseTest().remoteFolder

	_, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsGetWorksheet(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetsGetWorksheet - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
}
func TestCellsWorksheetsGetWorksheetForPageIndex(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsGetWorksheetOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.PageIndex = 1
	args.Format = "pdf"
	args.Folder = GetBaseTest().remoteFolder

	_, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsGetWorksheet(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetsGetWorksheet - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
}
func TestCellsWorksheetsGetWorksheetCalculateFormula(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsGetWorksheetCalculateFormulaOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Formula = "Now()"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsGetWorksheetCalculateFormula(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetsGetWorksheetCalculateFormula - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorksheetsGetWorksheetComments(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsGetWorksheetCommentsOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsGetWorksheetComments(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetsGetWorksheetComments - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorksheetsGetWorksheetMergedCell(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsGetWorksheetMergedCellOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.MergedCellIndex = 0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsGetWorksheetMergedCell(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetsGetWorksheetMergedCell - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorksheetsGetWorksheetMergedCells(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsGetWorksheetMergedCellsOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsGetWorksheetMergedCells(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetsGetWorksheetMergedCells - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorksheetsGetWorksheetTextItems(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsGetWorksheetTextItemsOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsGetWorksheetTextItems(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetsGetWorksheetTextItems - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorksheetsGetWorksheets(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsGetWorksheetsOpts)
	args.Name = GetBook1()

	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsGetWorksheets(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetsGetWorksheets - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorksheetsPostAutofitWorksheetColumns(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsPostAutofitWorksheetColumnsOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.FirstColumn = 1
	args.LastColumn = (11)
	args.FirstRow = 1
	args.LastRow = (10)
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsPostAutofitWorksheetColumns(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetsPostAutofitWorksheetColumns - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorksheetsPostAutofitWorksheetRow(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsPostAutofitWorksheetRowOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.FirstColumn = 1
	args.LastColumn = (11)
	args.RowIndex = 1
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsPostAutofitWorksheetRow(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetsPostAutofitWorksheetRow - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorksheetsPostAutofitWorksheetRows(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsPostAutofitWorksheetRowsOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.StartRow = 1
	args.EndRow = (11)
	args.OnlyAuto = true
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsPostAutofitWorksheetRows(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetsPostAutofitWorksheetRows - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorksheetValidationsGetWorksheetValidation(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetValidationsGetWorksheetValidationOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.ValidationIndex = 0
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetValidationsGetWorksheetValidation(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetValidationsGetWorksheetValidation - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorksheetsGetWorksheetComment(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsGetWorksheetCommentOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.CellName = "B3"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsGetWorksheetComment(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetsGetWorksheetComment - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}
func TestCellsWorksheetsPostCopyWorksheet(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsPostCopyWorksheetOpts)
	args.Name = GetBook1()
	args.SheetName = "GetSheet1"
	args.SourceSheet = GetSheet3()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsPostCopyWorksheet(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetsPostCopyWorksheet - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorksheetsPostMoveWorksheet(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsPostMoveWorksheetOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Moving = new(WorksheetMovingRequest)
	args.Moving.Position = "after"
	args.Moving.DestinationWorksheet = GetSheet3()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsPostMoveWorksheet(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetsPostMoveWorksheet - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorksheetsPostRenameWorksheet(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsPostRenameWorksheetOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Newname = "NewSheet"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsPostRenameWorksheet(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetsPostRenameWorksheet - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorksheetsPostUpdateWorksheetProperty(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsPostUpdateWorksheetPropertyOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Sheet = new(Worksheet)
	args.Sheet.IsSelected = true
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsPostUpdateWorksheetProperty(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetsPostUpdateWorksheetProperty - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorksheetsPostUpdateWorksheetZoom(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsPostUpdateWorksheetZoomOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Value = (12)
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsPostUpdateWorksheetZoom(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetsPostUpdateWorksheetZoom - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorksheetsPostWorksheetRangeSort(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsPostWorksheetRangeSortOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.CellArea = GetRange()
	args.DataSorter = new(DataSorter)
	args.DataSorter.CaseSensitive = true
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsPostWorksheetRangeSort(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetsPostWorksheetRangeSort - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorksheetsPostWorksheetTextSearch(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsPostWorksheetTextSearchOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Text = "abc"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsPostWorksheetTextSearch(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetsPostWorksheetTextSearch - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorksheetsPostWorsheetTextReplace(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsPostWorsheetTextReplaceOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.OldValue = "test"
	args.NewValue = "abc"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsPostWorsheetTextReplace(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetsPostWorsheetTextReplace - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorksheetsPutAddNewWorksheet(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsPutAddNewWorksheetOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Position = 1
	args.Sheettype = "VB"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsPutAddNewWorksheet(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetsPutAddNewWorksheet - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorksheetsPutChangeVisibilityWorksheet(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsPutChangeVisibilityWorksheetOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.IsVisible = true
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsPutChangeVisibilityWorksheet(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetsPutChangeVisibilityWorksheet - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorksheetsPutProtectWorksheet(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsPutProtectWorksheetOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.ProtectParameter = new(ProtectSheetParameter)
	args.ProtectParameter.Password = "123456"
	args.ProtectParameter.ProtectionType = "All"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsPutProtectWorksheet(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetsPutProtectWorksheet - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorksheetsPutWorksheetBackground(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsPutWorksheetBackgroundOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Png = GetPng()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsPutWorksheetBackground(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetsPutWorksheetBackground - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorksheetsPutWorksheetComment(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsPutWorksheetCommentOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.CellName = "A1"
	args.Comment = new(Comment)
	args.Comment.Note = "test"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsPutWorksheetComment(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetsPutWorksheetComment - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorksheetsPutWorksheetFreezePanes(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsPutWorksheetFreezePanesOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Row = 1
	args.Column = 1
	args.FreezedRows = 1
	args.FreezedColumns = 1
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsPutWorksheetFreezePanes(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetsPutWorksheetFreezePanes - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPutConvertWorkbook(t *testing.T) {

	args := new(CellsWorkbookPutConvertWorkbookOpts)
	args.Format = "PDF"
	// args.OutPath = "book1.xlsx.pdf"
	file, err := os.Open("../TestData/" + GetBook1())
	if err != nil {
		return
	}
	localVarReturnValue, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookPutConvertWorkbook(file, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t TestCellsPutConvertWorkbook - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
		file1, err2 := os.Create("../Book1.pdf")
		if err2 != nil {
			return
		}
		if _, err3 := file1.Write(localVarReturnValue); err3 != nil {
			fmt.Println(err3)
		}
		file1.Close()
	}
}
func TestCellsPutConvertWorkbookToOtherStorage(t *testing.T) {

	args := new(CellsWorkbookPutConvertWorkbookOpts)
	args.Format = "PDF"
	args.OutPath = "book1.xlsx.pdf"
	args.StorageName = "DropBox"
	file, err := os.Open("../TestData/" + GetBook1())
	if err != nil {
		return
	}
	localVarReturnValue, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookPutConvertWorkbook(file, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t TestCellsPutConvertWorkbookToOtherStorage - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
		file1, err2 := os.Create("../Book1withother.pdf")
		if err2 != nil {
			fmt.Println(err2)
			return
		}
		if _, err3 := file1.Write(localVarReturnValue); err3 != nil {
			fmt.Println(err2)
			fmt.Println(err3)
		}
		file1.Close()
	}
}
func TestCellsImportData(t *testing.T) {
	name := GetImportData()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorkbookPostImportDataOpts)
	args.Name = GetImportData()
	importIntArrayOption := new(ImportIntArrayOption)
	importIntArrayOption.Data = []int64{1, 2, 3}
	importIntArrayOption.DestinationWorksheet = GetSheet1()
	importIntArrayOption.FirstColumn = 2
	importIntArrayOption.FirstRow = 1
	importIntArrayOption.IsVertical = true
	importIntArrayOption.ImportDataType = "IntArray"
	args.ImportData = &importIntArrayOption

	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookPostImportData(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsImportData - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCopyFile(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CopyFileOpts)
	args.SrcPath = "GoTest/" + GetBook1()
	args.DestPath = "GoTest/1" + GetBook1()

	httpResponse, err := GetBaseTest().CellsAPI.CopyFile(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCopyFile - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
}

func TestCellsWorksheetsPostWorksheetComment(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsPostWorksheetCommentOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.CellName = "B3"
	args.Comment = new(Comment)
	args.Comment.Note = "string"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsPostWorksheetComment(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetsPostWorksheetComment - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsTaskPostRunTasks(t *testing.T) {
	args := new(CellsTaskPostRunTaskOpts)
	cellsTaskData := new(TaskData)
	cellsTaskDescription := new(TaskDescription)
	cellsSplitWorkbookTaskParameter := new(SplitWorkbookTaskParameter)
	cellsSplitWorkbookTaskParameter.DestinationFileFormat = "xlsx"
	cellsSplitWorkbookTaskParameter.SplitNameRule = "sheetname"
	cellsFileSource1 := new(FileSource)
	cellsFileSource1.FilePath = "GoTest"
	cellsFileSource1.FileSourceType = "CloudFileSystem"
	cellsFileSource2 := new(FileSource)
	cellsFileSource2.FilePath = "GoTest/Book1.xlsx"
	cellsFileSource2.FileSourceType = "CloudFileSystem"
	cellsSplitWorkbookTaskParameter.DestinationFilePosition = &cellsFileSource1
	cellsSplitWorkbookTaskParameter.Workbook = &cellsFileSource2

	cellsTaskDescription.TaskParameter = &cellsSplitWorkbookTaskParameter
	cellsTaskDescription.TaskType = "SplitWorkbook"

	cellsTaskData.Tasks = []TaskDescription{*cellsTaskDescription}
	args.TaskData = &cellsTaskData
	_, httpResponse, err := GetBaseTest().CellsAPI.CellsTaskPostRunTask(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t TestCellsTaskPostRunTasks - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
}

func TestCopyFolder(t *testing.T) {
	argsCF := new(CreateFolderOpts)
	argsCF.Path = "CreatedFolder"

	httpResponseCF, errCF := GetBaseTest().CellsAPI.CreateFolder(argsCF)
	if errCF != nil {
		t.Error(errCF)
	} else if httpResponseCF.StatusCode < 200 || httpResponseCF.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCreateFolder - %d\n", GetBaseTest().GetTestNumber(), httpResponseCF.StatusCode)
	}

	args := new(CopyFolderOpts)
	args.SrcPath = "CreatedFolder"
	args.DestPath = "GoTest"

	httpResponse, err := GetBaseTest().CellsAPI.CopyFolder(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCopyFolder - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
}

func TestDeleteFile(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(DeleteFileOpts)
	args.Path = "GoTest/" + GetBook1()

	httpResponse, err := GetBaseTest().CellsAPI.DeleteFile(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestDeleteFile - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
}

func TestDeleteFolder(t *testing.T) {
	args := new(DeleteFolderOpts)
	args.Path = "GetBookFolder"
	httpResponse, err := GetBaseTest().CellsAPI.DeleteFolder(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestDeleteFile - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
}

func TestDownloadFile(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(DownloadFileOpts)
	args.Path = "GoTest/" + GetBook1()

	_, httpResponse, err := GetBaseTest().CellsAPI.DownloadFile(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestDownloadFile - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
}

func TestGetDiscUsage(t *testing.T) {
	args := new(GetDiscUsageOpts)

	_, httpResponse, err := GetBaseTest().CellsAPI.GetDiscUsage(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetDiscUsage - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
}

func TestGetFileVersions(t *testing.T) {
	if isDockerSDK {
		return
	}
	args := new(GetFileVersionsOpts)
	args.Path = "GoTest"

	_, httpResponse, err := GetBaseTest().CellsAPI.GetFileVersions(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetFileVersions - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
}

func TestGetFilesList(t *testing.T) {

	args := new(GetFilesListOpts)
	args.Path = "GoTest"

	_, httpResponse, err := GetBaseTest().CellsAPI.GetFilesList(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestGetFilesList - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
}

func TestMoveFile(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	args := new(DeleteFileOpts)
	args.Path = "GoTest/MvTest.xlsx"

	httpResponse, err := GetBaseTest().CellsAPI.DeleteFile(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestDeleteFile - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
	args1 := new(MoveFileOpts)
	args1.SrcPath = "GoTest/" + GetBook1()
	args1.DestPath = "GoTest/MvTest.xlsx"

	httpResponse, err = GetBaseTest().CellsAPI.MoveFile(args1)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestMoveFile - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
}

func TestMoveFolder(t *testing.T) {

	args := new(CreateFolderOpts)
	args.Path = "MovedFolder"

	httpResponse, err := GetBaseTest().CellsAPI.CreateFolder(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCreateFolder - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
	argsMF := new(MoveFolderOpts)
	argsMF.SrcPath = "MovedFolder"
	argsMF.DestPath = "DeletedFolder"

	httpResponseMF, errMF := GetBaseTest().CellsAPI.MoveFolder(argsMF)
	if errMF != nil {
		t.Error(errMF)
	} else if httpResponseMF.StatusCode < 200 || httpResponseMF.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestMoveFolder - %d\n", GetBaseTest().GetTestNumber(), httpResponseMF.StatusCode)
	}
	argsDF := new(DeleteFolderOpts)
	argsDF.Path = "DeletedFolder"
	httpResponseDF, errDF := GetBaseTest().CellsAPI.DeleteFolder(argsDF)
	if errDF != nil {
		t.Error(errDF)
	} else if httpResponseDF.StatusCode < 200 || httpResponseDF.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestDeleteFile - %d\n", GetBaseTest().GetTestNumber(), httpResponseDF.StatusCode)
	}

}

func TestObjectExists(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(ObjectExistsOpts)
	args.Path = "GoTest/" + GetBook1()

	_, httpResponse, err := GetBaseTest().CellsAPI.ObjectExists(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestObjectExists - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
}

func TestStorageExists(t *testing.T) {
	args := new(StorageExistsOpts)
	args.StorageName = "Default"
	_, httpResponse, err := GetBaseTest().CellsAPI.StorageExists(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestStorageExists - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
}

func TestCellsWorkbookDeleteWorkbookBackground(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorkbookDeleteWorkbookBackgroundOpts)
	args.Name = GetBook1()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookDeleteWorkbookBackground(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t CellsWorkbookDeleteWorkbookBackground - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorkbookPutWorkbookBackground(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorkbookPutWorkbookBackgroundOpts)
	args.Name = GetBook1()
	args.Png = GetPng()
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookPutWorkbookBackground(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t CellsWorkbookPutWorkbookBackground - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsWorkbookBatchFilesConvert(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(PostBatchConvertOpts)
	args.BatchConvertRequest = new(BatchConvertRequest)
	args.BatchConvertRequest.Format = "pdf"
	args.BatchConvertRequest.MatchCondition = new(MatchConditionRequest)
	args.BatchConvertRequest.MatchCondition.FullMatchConditions = []string{"Sheet1", "Sheet2"}
	_, httpResponse, err := GetBaseTest().CellsAPI.PostBatchConvert(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t PostBatchConvert - %d\n", GetBaseTest().GetTestNumber(), 200)
	}
}

func TestCellsWorkbookDeleteWorksheets(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsDeleteWorksheetsOpts)
	args.MatchCondition = new(MatchConditionRequest)

	args.MatchCondition = new(MatchConditionRequest)
	args.MatchCondition.FullMatchConditions = []string{"Sheet1", "Sheet2"}
	args.Folder = "GoTest"
	args.Name = GetBook1()
	_, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsDeleteWorksheets(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t PostBatchConvert - %d\n", GetBaseTest().GetTestNumber(), 200)
	}
}

func TestCellsWorkbookPageCount(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorkbookGetPageCountOpts)
	args.Folder = "GoTest"
	args.Name = GetBook1()
	_, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookGetPageCount(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t PostBatchConvert - %d\n", GetBaseTest().GetTestNumber(), 200)
	}
}

func TestCellsWorksheetsPageCount(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsGetPageCountOpts)
	args.Folder = "GoTest"
	args.Name = GetBook1()
	args.SheetName = "Sheet1"
	_, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsGetPageCount(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t PostBatchConvert - %d\n", GetBaseTest().GetTestNumber(), 200)
	}
}

func TestCellsListObjectsPostWorksheetListColumnsTotal(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsListObjectsPostWorksheetListColumnsTotalOpts)
	args.Folder = "GoTest"
	args.Name = GetBook1()
	args.SheetName = "Sheet7"
	args.ListObjectIndex = 0
	tableTotalRequest := new(TableTotalRequest)
	tableTotalRequest.ListColumnIndex = 1
	tableTotalRequest.TotalsCalculation = "Average"
	args.TableTotalRequests = &[]TableTotalRequest{*tableTotalRequest}

	_, httpResponse, err := GetBaseTest().CellsAPI.CellsListObjectsPostWorksheetListColumnsTotal(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t PostBatchConvert - %d\n", GetBaseTest().GetTestNumber(), 200)
	}
}

func TestCellsListObjectsPostWorksheetListColumn(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsListObjectsPostWorksheetListColumnOpts)
	args.Folder = "GoTest"
	args.Name = GetBook1()
	args.SheetName = "Sheet7"
	args.ListObjectIndex = 0
	args.ListColumn = new(ListColumn)
	args.ListColumn.Name = "test cloumn"
	args.ListColumn.TotalsCalculation = "Average"

	_, httpResponse, err := GetBaseTest().CellsAPI.CellsListObjectsPostWorksheetListColumn(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t PostBatchConvert - %d\n", GetBaseTest().GetTestNumber(), 200)
	}
}

func TestCellsShapesPostWorksheetGroupShape(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsShapesPostWorksheetGroupShapeOpts)
	args.Folder = "GoTest"
	args.Name = GetBook1()
	args.SheetName = "Sheet6"
	args.ListShape = []int64{0, 2}

	_, httpResponse, err := GetBaseTest().CellsAPI.CellsShapesPostWorksheetGroupShape(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t PostBatchConvert - %d\n", GetBaseTest().GetTestNumber(), 200)
	}
}

func TestCellsShapesPostWorksheetUngroupShape(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsShapesPostWorksheetUngroupShapeOpts)
	args.Folder = "GoTest"
	args.Name = GetBook1()
	args.SheetName = "Sheet6"
	args.Shapeindex = 0

	_, httpResponse, err := GetBaseTest().CellsAPI.CellsShapesPostWorksheetUngroupShape(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t PostBatchConvert - %d\n", GetBaseTest().GetTestNumber(), 200)
	}
}

func TestCellsWorkbookPostDigitalSignature(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	if err := GetBaseTest().UploadFile("roywang.pfx"); err != nil {
		t.Error(err)
	}
	args := new(CellsWorkbookPostDigitalSignatureOpts)
	args.Folder = "GoTest"
	args.Name = GetBook1()
	args.Digitalsignaturefile = "GoTest/roywang.pfx"
	args.Password = "123456"

	_, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookPostDigitalSignature(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t CellsWorkbookPostDigitalSignature - %d\n", GetBaseTest().GetTestNumber(), 200)
	}
}
func TestCellsPutConvertWorkbook_Extend(t *testing.T) {
	var extendedQueryParameters map[string]string
	extendedQueryParameters = make(map[string]string)
	extendedQueryParameters["OnePagePerSheet"] = "false"
	args := new(CellsWorkbookPutConvertWorkbookOpts)
	args.Format = "PDF"
	args.ExtendedQueryParameters = extendedQueryParameters
	// args.OutPath = "book1.xlsx.pdf"
	file, err := os.Open("../TestData/" + GetBook1())
	if err != nil {
		return
	}
	localVarReturnValue, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookPutConvertWorkbook(file, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t TestCellsPutConvertWorkbook - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
		file1, err2 := os.Create("../Book1_extend.pdf")
		if err2 != nil {
			return
		}
		if _, err3 := file1.Write(localVarReturnValue); err3 != nil {
			fmt.Println(err3)
		}
		file1.Close()
	}
}
func TestCellsWorkbookGetWorkbook_Extend(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	var extendedQueryParameters map[string]string
	extendedQueryParameters = make(map[string]string)
	extendedQueryParameters["OnePagePerSheet"] = "false"
	args := new(CellsWorkbookGetWorkbookOpts)
	args.Name = GetBook1()
	args.Format = "PDF"
	args.Folder = GetBaseTest().remoteFolder
	args.ExtendedQueryParameters = extendedQueryParameters

	_, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookGetWorkbook(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorkbookGetWorkbook - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
	}
}

func TestCellsSaveAsPostDocumentSaveAs_Extend(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}
	var extendedQueryParameters map[string]string
	extendedQueryParameters = make(map[string]string)
	extendedQueryParameters["OnePagePerSheet"] = "false"
	args := new(CellsSaveAsPostDocumentSaveAsOpts)
	args.Name = GetBook1()
	args.SaveOptions = nil
	args.Newfilename = GetBaseTest().remoteFolder + "/newfilego.pdf"
	args.Folder = GetBaseTest().remoteFolder
	args.ExtendedQueryParameters = extendedQueryParameters
	response, httpResponse, err := GetBaseTest().CellsAPI.CellsSaveAsPostDocumentSaveAs(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsSaveAsPostDocumentSaveAs - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
}

func TestCellsPictureGetExtractBarcodes(t *testing.T) {
	name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsPictureGetExtractBarcodesOpts)
	args.Name = GetBook1()
	args.SheetName = "Sheet8"
	args.PictureIndex = 0
	args.Folder = GetBaseTest().remoteFolder
	_, httpResponse, err := GetBaseTest().CellsAPI.CellsPictureGetExtractBarcodes(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsSaveAsPostDocumentSaveAs - %d\n", GetBaseTest().GetTestNumber(), 200)
	}
}
func TestCellsPutConvertWorkbookPDF(t *testing.T) {

	args := new(CellsWorkbookPutConvertWorkbookOpts)
	args.Format = "PDF"
	// args.OutPath = "book1.xlsx.pdf"
	file, err := os.Open("../TestData/" + GetBook1())
	if err != nil {
		return
	}
	localVarReturnValue, httpResponse, err := GetBaseTest().CellsAPI.CellsWorkbookPutConvertWorkbook(file, args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t TestCellsPutConvertWorkbook - %d\n", GetBaseTest().GetTestNumber(), httpResponse.StatusCode)
		file1, err2 := os.Create("../Book1.pdf")
		if err2 != nil {
			return
		}
		if _, err3 := file1.Write(localVarReturnValue); err3 != nil {
			fmt.Println(err3)
		}
		file1.Close()
	}
}
func TestCellsPostConvertWorkbookToPNG(t *testing.T) {

	args := new(PostConvertWorkbookToPNGOpts)

	_, httpResponse, err := GetBaseTest().CellsAPI.PostConvertWorkbookToPNG("../TestData/"+GetBook1(), args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t func TestCells PostConvertWorkbookToPNG	\n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsPostConvertWorkbookToPDF(t *testing.T) {

	args := new(PostConvertWorkbookToPDFOpts)

	_, httpResponse, err := GetBaseTest().CellsAPI.PostConvertWorkbookToPDF("../TestData/"+GetBook1(), args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t func TestCells PostConvertWorkbookToPDF	\n", GetBaseTest().GetTestNumber())
	}
}
func TestCellsPostConvertWorkbookToDocx(t *testing.T) {

	args := new(PostConvertWorkbookToDocxOpts)
	_, httpResponse, err := GetBaseTest().CellsAPI.PostConvertWorkbookToDocx("../TestData/"+GetBook1(), args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t func TestCells PostConvertWorkbookToDocx \n", GetBaseTest().GetTestNumber())
	}
}
