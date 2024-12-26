package asposecellscloudtest

import (
	"fmt"
	"testing"

	. "github.com/aspose-cells-cloud/aspose-cells-cloud-go/v24"
)

func TestChartsController_GetWorksheetCharts(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(GetWorksheetChartsRequest)
	request.Name = remoteName
	request.SheetName = "Sheet4"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetCharts(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestChartsController_GetWorksheetCharts \n", GetBaseTest().GetTestNumber())
	}
}

func TestChartsController_GetWorksheetChart(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(GetWorksheetChartRequest)
	request.Name = remoteName
	request.SheetName = "Sheet4"
	request.ChartNumber = int64(0)
	request.Format = "png"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetChart(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestChartsController_GetWorksheetChart \n", GetBaseTest().GetTestNumber())
	}
}

func TestChartsController_PutWorksheetChart(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PutWorksheetChartRequest)
	request.Name = remoteName
	request.SheetName = "Sheet4"
	request.ChartType = "Pie"
	request.UpperLeftRow = int64(5)
	request.UpperLeftColumn = int64(5)
	request.LowerRightRow = int64(10)
	request.LowerRightColumn = int64(10)
	request.Area = "C7:D11"
	request.IsVertical = true
	request.Title = "Aspose Chart"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PutWorksheetChart(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestChartsController_PutWorksheetChart \n", GetBaseTest().GetTestNumber())
	}
}

func TestChartsController_DeleteWorksheetChart(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(DeleteWorksheetChartRequest)
	request.Name = remoteName
	request.SheetName = "Sheet4"
	request.ChartIndex = int64(0)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheetChart(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestChartsController_DeleteWorksheetChart \n", GetBaseTest().GetTestNumber())
	}
}

func TestChartsController_PostWorksheetChart(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var chart = new(Chart)
	chart.ShowLegend = true
	chart.ShowDataTable = true

	request := new(PostWorksheetChartRequest)
	request.Name = remoteName
	request.SheetName = "Sheet4"
	request.ChartIndex = int64(0)
	request.Chart = chart
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetChart(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestChartsController_PostWorksheetChart \n", GetBaseTest().GetTestNumber())
	}
}

func TestChartsController_GetWorksheetChartLegend(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(GetWorksheetChartLegendRequest)
	request.Name = remoteName
	request.SheetName = "Sheet4"
	request.ChartIndex = int64(0)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetChartLegend(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestChartsController_GetWorksheetChartLegend \n", GetBaseTest().GetTestNumber())
	}
}

func TestChartsController_PostWorksheetChartLegend(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var legend = new(Legend)
	legend.Position = "Top"

	request := new(PostWorksheetChartLegendRequest)
	request.Name = remoteName
	request.SheetName = "Sheet4"
	request.ChartIndex = int64(0)
	request.Legend = legend
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetChartLegend(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestChartsController_PostWorksheetChartLegend \n", GetBaseTest().GetTestNumber())
	}
}

func TestChartsController_PutWorksheetChartLegend(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PutWorksheetChartLegendRequest)
	request.Name = remoteName
	request.SheetName = "Sheet4"
	request.ChartIndex = int64(0)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PutWorksheetChartLegend(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestChartsController_PutWorksheetChartLegend \n", GetBaseTest().GetTestNumber())
	}
}

func TestChartsController_DeleteWorksheetChartLegend(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(DeleteWorksheetChartLegendRequest)
	request.Name = remoteName
	request.SheetName = "Sheet4"
	request.ChartIndex = int64(0)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheetChartLegend(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestChartsController_DeleteWorksheetChartLegend \n", GetBaseTest().GetTestNumber())
	}
}

func TestChartsController_DeleteWorksheetCharts(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(DeleteWorksheetChartsRequest)
	request.Name = remoteName
	request.SheetName = "Sheet4"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheetCharts(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestChartsController_DeleteWorksheetCharts \n", GetBaseTest().GetTestNumber())
	}
}

func TestChartsController_GetWorksheetChartTitle(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(GetWorksheetChartTitleRequest)
	request.Name = remoteName
	request.SheetName = "Sheet4"
	request.ChartIndex = int64(0)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetChartTitle(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestChartsController_GetWorksheetChartTitle \n", GetBaseTest().GetTestNumber())
	}
}

func TestChartsController_PostWorksheetChartTitle(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var title = new(Title)
	title.IsVisible = true

	request := new(PostWorksheetChartTitleRequest)
	request.Name = remoteName
	request.SheetName = "Sheet4"
	request.ChartIndex = int64(0)
	request.Title = title
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetChartTitle(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestChartsController_PostWorksheetChartTitle \n", GetBaseTest().GetTestNumber())
	}
}

func TestChartsController_PutWorksheetChartTitle(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var title = new(Title)
	title.IsVisible = true

	request := new(PutWorksheetChartTitleRequest)
	request.Name = remoteName
	request.SheetName = "Sheet4"
	request.ChartIndex = int64(0)
	request.Title = title
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PutWorksheetChartTitle(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestChartsController_PutWorksheetChartTitle \n", GetBaseTest().GetTestNumber())
	}
}

func TestChartsController_DeleteWorksheetChartTitle(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(DeleteWorksheetChartTitleRequest)
	request.Name = remoteName
	request.SheetName = "Sheet4"
	request.ChartIndex = int64(0)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheetChartTitle(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestChartsController_DeleteWorksheetChartTitle \n", GetBaseTest().GetTestNumber())
	}
}
