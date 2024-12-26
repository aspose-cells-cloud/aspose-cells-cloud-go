package asposecellscloudtest

import (
	"fmt"
	"testing"

	. "github.com/aspose-cells-cloud/aspose-cells-cloud-go/v24"
)

func TestCellsController_PostClearContents(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PostClearContentsRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.Range_ = "A1:C10"
	request.StartRow = int64(1)
	request.StartColumn = int64(1)
	request.EndRow = int64(3)
	request.EndColumn = int64(3)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostClearContents(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_PostClearContents \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostClearFormats(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PostClearFormatsRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.Range_ = "A1:C10"
	request.StartRow = int64(1)
	request.StartColumn = int64(1)
	request.EndRow = int64(3)
	request.EndColumn = int64(3)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostClearFormats(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_PostClearFormats \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostUpdateWorksheetRangeStyle(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var styleFont = new(Font)
	styleFont.Size = int64(16)
	var style = new(Style)
	style.Font = styleFont

	request := new(PostUpdateWorksheetRangeStyleRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.Range_ = "A1:C10"
	request.Style = style
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostUpdateWorksheetRangeStyle(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_PostUpdateWorksheetRangeStyle \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostWorksheetMerge(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PostWorksheetMergeRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.StartRow = int64(1)
	request.StartColumn = int64(1)
	request.TotalRows = int64(4)
	request.TotalColumns = int64(4)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetMerge(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_PostWorksheetMerge \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostWorksheetUnmerge(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PostWorksheetUnmergeRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.StartRow = int64(1)
	request.StartColumn = int64(1)
	request.TotalRows = int64(4)
	request.TotalColumns = int64(4)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetUnmerge(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_PostWorksheetUnmerge \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_GetWorksheetCells(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(GetWorksheetCellsRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.Offest = int64(1)
	request.Count = int64(10)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetCells(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_GetWorksheetCells \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_GetWorksheetCell(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(GetWorksheetCellRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.CellOrMethodName = "A1"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetCell(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_GetWorksheetCell \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_GetWorksheetCellStyle(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(GetWorksheetCellStyleRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.CellName = "A1"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetCellStyle(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_GetWorksheetCellStyle \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostWorksheetCellSetValue(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PostWorksheetCellSetValueRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.CellName = "A1"
	request.Value = "1"
	request.Type_ = "int"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetCellSetValue(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_PostWorksheetCellSetValue \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostUpdateWorksheetCellStyle(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var styleFont = new(Font)
	styleFont.Size = int64(16)
	var style = new(Style)
	style.Font = styleFont

	request := new(PostUpdateWorksheetCellStyleRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.CellName = "A1"
	request.Style = style
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostUpdateWorksheetCellStyle(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_PostUpdateWorksheetCellStyle \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostSetCellRangeValue(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PostSetCellRangeValueRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.Cellarea = "A1:C10"
	request.Value = "Test"
	request.Type_ = "string"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostSetCellRangeValue(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_PostSetCellRangeValue \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostCopyCellIntoCell(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PostCopyCellIntoCellRequest)
	request.Name = remoteName
	request.DestCellName = "C1"
	request.SheetName = "Sheet1"
	request.Worksheet = "Sheet2"
	request.Cellname = "A1"
	request.Row = int64(1)
	request.Column = int64(1)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostCopyCellIntoCell(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_PostCopyCellIntoCell \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_GetCellHtmlString(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(GetCellHtmlStringRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.CellName = "A1"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.GetCellHtmlString(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_GetCellHtmlString \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostSetCellHtmlString(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PostSetCellHtmlStringRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.CellName = "A1"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostSetCellHtmlString(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_PostSetCellHtmlString \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostCellCalculate(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var options = new(CalculationOptions)
	options.Recursive = true
	options.IgnoreError = true

	request := new(PostCellCalculateRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.CellName = "A1"
	request.Options = options
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostCellCalculate(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_PostCellCalculate \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostCellCharacters(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var optionsvalue0Font = new(Font)
	optionsvalue0Font.IsBold = true
	optionsvalue0Font.Size = int64(16)
	var optionsvalue0 = new(FontSetting)
	optionsvalue0.Length = int64(5)
	optionsvalue0.StartIndex = int64(0)
	optionsvalue0.Font = optionsvalue0Font
	var options = []FontSetting{*optionsvalue0}

	request := new(PostCellCharactersRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.CellName = "E36"
	request.Options = options
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostCellCharacters(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_PostCellCharacters \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_GetWorksheetColumns(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(GetWorksheetColumnsRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.Offset = int64(1)
	request.Count = int64(10)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetColumns(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_GetWorksheetColumns \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostSetWorksheetColumnWidth(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PostSetWorksheetColumnWidthRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.ColumnIndex = int64(1)
	request.Width = 10.9
	request.Count = int64(10)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostSetWorksheetColumnWidth(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_PostSetWorksheetColumnWidth \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_GetWorksheetColumn(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(GetWorksheetColumnRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.ColumnIndex = int64(1)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetColumn(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_GetWorksheetColumn \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PutInsertWorksheetColumns(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PutInsertWorksheetColumnsRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.ColumnIndex = int64(1)
	request.Columns = int64(10)
	request.UpdateReference = true
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PutInsertWorksheetColumns(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_PutInsertWorksheetColumns \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_DeleteWorksheetColumns(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(DeleteWorksheetColumnsRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.ColumnIndex = int64(1)
	request.Columns = int64(10)
	request.UpdateReference = true
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheetColumns(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_DeleteWorksheetColumns \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostHideWorksheetColumns(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PostHideWorksheetColumnsRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.StartColumn = int64(1)
	request.TotalColumns = int64(10)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostHideWorksheetColumns(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_PostHideWorksheetColumns \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostUnhideWorksheetColumns(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PostUnhideWorksheetColumnsRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.StartColumn = int64(1)
	request.TotalColumns = int64(10)
	request.Width = 10.9
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostUnhideWorksheetColumns(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_PostUnhideWorksheetColumns \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostGroupWorksheetColumns(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PostGroupWorksheetColumnsRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.FirstIndex = int64(1)
	request.LastIndex = int64(9)
	request.Hide = true
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostGroupWorksheetColumns(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_PostGroupWorksheetColumns \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostUngroupWorksheetColumns(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PostUngroupWorksheetColumnsRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.FirstIndex = int64(1)
	request.LastIndex = int64(9)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostUngroupWorksheetColumns(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_PostUngroupWorksheetColumns \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostCopyWorksheetColumns(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PostCopyWorksheetColumnsRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.SourceColumnIndex = int64(1)
	request.DestinationColumnIndex = int64(19)
	request.ColumnNumber = int64(8)
	request.Worksheet = "Sheet2"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostCopyWorksheetColumns(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_PostCopyWorksheetColumns \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostColumnStyle(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var styleFont = new(Font)
	styleFont.Size = int64(16)
	var style = new(Style)
	style.Font = styleFont

	request := new(PostColumnStyleRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.ColumnIndex = int64(1)
	request.Style = style
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostColumnStyle(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_PostColumnStyle \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_GetWorksheetRows(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(GetWorksheetRowsRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.Offset = int64(1)
	request.Count = int64(10)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetRows(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_GetWorksheetRows \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_GetWorksheetRow(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(GetWorksheetRowRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.RowIndex = int64(1)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetRow(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_GetWorksheetRow \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_DeleteWorksheetRow(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(DeleteWorksheetRowRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.RowIndex = int64(1)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheetRow(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_DeleteWorksheetRow \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_DeleteWorksheetRows(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(DeleteWorksheetRowsRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.Startrow = int64(1)
	request.TotalRows = int64(10)
	request.UpdateReference = true
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheetRows(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_DeleteWorksheetRows \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PutInsertWorksheetRows(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PutInsertWorksheetRowsRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.Startrow = int64(1)
	request.TotalRows = int64(10)
	request.UpdateReference = true
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PutInsertWorksheetRows(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_PutInsertWorksheetRows \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PutInsertWorksheetRow(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PutInsertWorksheetRowRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.RowIndex = int64(1)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PutInsertWorksheetRow(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_PutInsertWorksheetRow \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostUpdateWorksheetRow(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PostUpdateWorksheetRowRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.RowIndex = int64(1)
	request.Height = 10.8
	request.Count = int64(9)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostUpdateWorksheetRow(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_PostUpdateWorksheetRow \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostHideWorksheetRows(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PostHideWorksheetRowsRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.Startrow = int64(1)
	request.TotalRows = int64(6)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostHideWorksheetRows(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_PostHideWorksheetRows \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostUnhideWorksheetRows(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PostUnhideWorksheetRowsRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.Startrow = int64(1)
	request.TotalRows = int64(8)
	request.Height = 10.9
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostUnhideWorksheetRows(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_PostUnhideWorksheetRows \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostGroupWorksheetRows(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PostGroupWorksheetRowsRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.FirstIndex = int64(1)
	request.LastIndex = int64(9)
	request.Hide = true
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostGroupWorksheetRows(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_PostGroupWorksheetRows \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostUngroupWorksheetRows(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PostUngroupWorksheetRowsRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.FirstIndex = int64(1)
	request.LastIndex = int64(9)
	request.IsAll = true
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostUngroupWorksheetRows(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_PostUngroupWorksheetRows \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostCopyWorksheetRows(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PostCopyWorksheetRowsRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.SourceRowIndex = int64(1)
	request.DestinationRowIndex = int64(12)
	request.RowNumber = int64(5)
	request.Worksheet = "Sheet2"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostCopyWorksheetRows(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_PostCopyWorksheetRows \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostRowStyle(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var styleFont = new(Font)
	styleFont.Size = int64(16)
	var style = new(Style)
	style.Font = styleFont

	request := new(PostRowStyleRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.RowIndex = int64(1)
	request.Style = style
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostRowStyle(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsController_PostRowStyle \n", GetBaseTest().GetTestNumber())
	}
}
