package asposecellscloud

import (
	. "asposecellscloud"
	"fmt"
	"testing"
)

func TestRangesController_PostWorksheetCellsRangesCopy(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var rangeOperateSource = new(Range)
	rangeOperateSource.ColumnCount = int64(1)
	rangeOperateSource.ColumnWidth = 10.0
	rangeOperateSource.FirstRow = int64(1)
	rangeOperateSource.RowCount = int64(10)
	var rangeOperateTarget = new(Range)
	rangeOperateTarget.ColumnCount = int64(1)
	rangeOperateTarget.ColumnWidth = 10.0
	rangeOperateTarget.FirstRow = int64(10)
	rangeOperateTarget.RowCount = int64(10)
	var rangeOperate = new(RangeCopyRequest)
	rangeOperate.Operate = "copydata"
	rangeOperate.Source = rangeOperateSource
	rangeOperate.Target = rangeOperateTarget

	request := new(PostWorksheetCellsRangesCopyRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.RangeOperate = rangeOperate
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetCellsRangesCopy(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestRangesController_PostWorksheetCellsRangesCopy \n", GetBaseTest().GetTestNumber())
	}
}

func TestRangesController_PostWorksheetCellsRangeMerge(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var range_ = new(Range)
	range_.ColumnCount = int64(1)
	range_.ColumnWidth = 10.0
	range_.FirstRow = int64(1)
	range_.RowCount = int64(10)

	request := new(PostWorksheetCellsRangeMergeRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.Range_ = range_
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetCellsRangeMerge(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestRangesController_PostWorksheetCellsRangeMerge \n", GetBaseTest().GetTestNumber())
	}
}

func TestRangesController_PostWorksheetCellsRangeUnMerge(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var range_ = new(Range)
	range_.ColumnCount = int64(1)
	range_.ColumnWidth = 10.0
	range_.FirstRow = int64(1)
	range_.RowCount = int64(10)

	request := new(PostWorksheetCellsRangeUnMergeRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.Range_ = range_
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetCellsRangeUnMerge(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestRangesController_PostWorksheetCellsRangeUnMerge \n", GetBaseTest().GetTestNumber())
	}
}

func TestRangesController_PostWorksheetCellsRangeStyle(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var rangeOperateStyleFont = new(Font)
	rangeOperateStyleFont.Size = int64(16)
	var rangeOperateStyle = new(Style)
	rangeOperateStyle.Font = rangeOperateStyleFont
	var rangeOperateRange = new(Range)
	rangeOperateRange.ColumnCount = int64(1)
	rangeOperateRange.ColumnWidth = 10.0
	rangeOperateRange.FirstRow = int64(1)
	rangeOperateRange.RowCount = int64(10)
	var rangeOperate = new(RangeSetStyleRequest)
	rangeOperate.Style = rangeOperateStyle
	rangeOperate.Range_ = rangeOperateRange

	request := new(PostWorksheetCellsRangeStyleRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.RangeOperate = rangeOperate
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetCellsRangeStyle(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestRangesController_PostWorksheetCellsRangeStyle \n", GetBaseTest().GetTestNumber())
	}
}

func TestRangesController_GetWorksheetCellsRangeValue(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(GetWorksheetCellsRangeValueRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.Namerange = "Name_2"
	request.FirstRow = int64(0)
	request.FirstColumn = int64(0)
	request.RowCount = int64(3)
	request.ColumnCount = int64(2)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetCellsRangeValue(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestRangesController_GetWorksheetCellsRangeValue \n", GetBaseTest().GetTestNumber())
	}
}

func TestRangesController_PostWorksheetCellsRangeValue(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var range_ = new(Range)
	range_.ColumnCount = int64(1)
	range_.ColumnWidth = 10.0
	range_.FirstRow = int64(1)
	range_.RowCount = int64(10)

	request := new(PostWorksheetCellsRangeValueRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.Range_ = range_
	request.Value = "100"
	request.IsConverted = true
	request.SetStyle = true
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetCellsRangeValue(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestRangesController_PostWorksheetCellsRangeValue \n", GetBaseTest().GetTestNumber())
	}
}

func TestRangesController_PostWorksheetCellsRangeMoveTo(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var range_ = new(Range)
	range_.ColumnCount = int64(1)
	range_.ColumnWidth = 10.0
	range_.FirstRow = int64(1)
	range_.RowCount = int64(10)

	request := new(PostWorksheetCellsRangeMoveToRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.Range_ = range_
	request.DestRow = int64(10)
	request.DestColumn = int64(10)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetCellsRangeMoveTo(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestRangesController_PostWorksheetCellsRangeMoveTo \n", GetBaseTest().GetTestNumber())
	}
}

func TestRangesController_PostWorksheetCellsRangeOutlineBorder(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var rangeOperateborderColor = new(Color)
	rangeOperateborderColor.R = int64(48)
	rangeOperateborderColor.G = int64(48)
	rangeOperateborderColor.B = int64(48)
	var rangeOperateRange = new(Range)
	rangeOperateRange.ColumnCount = int64(1)
	rangeOperateRange.ColumnWidth = 10.0
	rangeOperateRange.FirstRow = int64(1)
	rangeOperateRange.RowCount = int64(10)
	var rangeOperate = new(RangeSetOutlineBorderRequest)
	rangeOperate.BorderEdge = "LeftBorder"
	rangeOperate.BorderStyle = "Dotted"
	rangeOperate.BorderColor = rangeOperateborderColor
	rangeOperate.Range_ = rangeOperateRange

	request := new(PostWorksheetCellsRangeOutlineBorderRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.RangeOperate = rangeOperate
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetCellsRangeOutlineBorder(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestRangesController_PostWorksheetCellsRangeOutlineBorder \n", GetBaseTest().GetTestNumber())
	}
}

func TestRangesController_PostWorksheetCellsRangeColumnWidth(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var range_ = new(Range)
	range_.ColumnCount = int64(1)
	range_.ColumnWidth = 10.0
	range_.FirstRow = int64(1)
	range_.RowCount = int64(10)

	request := new(PostWorksheetCellsRangeColumnWidthRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.Range_ = range_
	request.Value = 10.7
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetCellsRangeColumnWidth(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestRangesController_PostWorksheetCellsRangeColumnWidth \n", GetBaseTest().GetTestNumber())
	}
}

func TestRangesController_PostWorksheetCellsRangeRowHeight(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var range_ = new(Range)
	range_.ColumnCount = int64(1)
	range_.ColumnWidth = 10.0
	range_.FirstRow = int64(1)
	range_.RowCount = int64(10)

	request := new(PostWorksheetCellsRangeRowHeightRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.Range_ = range_
	request.Value = 10.9
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetCellsRangeRowHeight(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestRangesController_PostWorksheetCellsRangeRowHeight \n", GetBaseTest().GetTestNumber())
	}
}

func TestRangesController_PutWorksheetCellsRange(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PutWorksheetCellsRangeRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.Range_ = "A1:C6"
	request.Shift = "Down"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PutWorksheetCellsRange(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestRangesController_PutWorksheetCellsRange \n", GetBaseTest().GetTestNumber())
	}
}

func TestRangesController_DeleteWorksheetCellsRange(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(DeleteWorksheetCellsRangeRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.Range_ = "A1:C6"
	request.Shift = "Up"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheetCellsRange(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestRangesController_DeleteWorksheetCellsRange \n", GetBaseTest().GetTestNumber())
	}
}

func TestRangesController_PostWorksheetCellsRangeSort(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Group.xlsx"
	remoteName := "Group.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var rangeSortRequestDataSorter = new(DataSorter)
	rangeSortRequestDataSorter.CaseSensitive = true
	var rangeSortRequestCellArea = new(Range)
	rangeSortRequestCellArea.ColumnCount = int64(3)
	rangeSortRequestCellArea.FirstColumn = int64(0)
	rangeSortRequestCellArea.FirstRow = int64(0)
	rangeSortRequestCellArea.RowCount = int64(15)
	var rangeSortRequest = new(RangeSortRequest)
	rangeSortRequest.DataSorter = rangeSortRequestDataSorter
	rangeSortRequest.CellArea = rangeSortRequestCellArea

	request := new(PostWorksheetCellsRangeSortRequest)
	request.Name = remoteName
	request.SheetName = "book1"
	request.RangeSortRequest = rangeSortRequest
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetCellsRangeSort(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestRangesController_PostWorksheetCellsRangeSort \n", GetBaseTest().GetTestNumber())
	}
}
