package asposecellscloud

import (
	"fmt"
	"testing"
)

func TestPivotTablesController_GetWorksheetPivotTables(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(GetWorksheetPivotTablesRequest)
	request.Name = remoteName
	request.SheetName = "Sheet4"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetPivotTables(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPivotTablesController_GetWorksheetPivotTables \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_GetWorksheetPivotTable(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(GetWorksheetPivotTableRequest)
	request.Name = remoteName
	request.SheetName = "Sheet4"
	request.PivottableIndex = int64(0)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetPivotTable(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPivotTablesController_GetWorksheetPivotTable \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_GetPivotTableField(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(GetPivotTableFieldRequest)
	request.Name = remoteName
	request.SheetName = "Sheet4"
	request.PivotTableIndex = int64(0)
	request.PivotFieldIndex = int64(0)
	request.PivotFieldType = "Row"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.GetPivotTableField(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPivotTablesController_GetPivotTableField \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_GetWorksheetPivotTableFilters(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(GetWorksheetPivotTableFiltersRequest)
	request.Name = remoteName
	request.SheetName = "Sheet3"
	request.PivotTableIndex = int64(0)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetPivotTableFilters(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPivotTablesController_GetWorksheetPivotTableFilters \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_GetWorksheetPivotTableFilter(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(GetWorksheetPivotTableFilterRequest)
	request.Name = remoteName
	request.SheetName = "Sheet3"
	request.PivotTableIndex = int64(0)
	request.FilterIndex = int64(0)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetPivotTableFilter(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPivotTablesController_GetWorksheetPivotTableFilter \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_PutWorksheetPivotTable(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PutWorksheetPivotTableRequest)
	request.Name = remoteName
	request.SheetName = "Sheet3"
	request.Folder = remoteFolder
	request.SourceData = "=Sheet1!C6:E13"
	request.DestCellName = "C1"
	request.TableName = "TestPivot"
	request.UseSameSource = true
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PutWorksheetPivotTable(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPivotTablesController_PutWorksheetPivotTable \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_PutPivotTableField(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var pivotTableFieldRequestData = []int64{int64(0)}
	var pivotTableFieldRequest = new(PivotTableFieldRequest)
	pivotTableFieldRequest.Data = pivotTableFieldRequestData

	request := new(PutPivotTableFieldRequest)
	request.Name = remoteName
	request.SheetName = "Sheet4"
	request.PivotTableIndex = int64(0)
	request.PivotFieldType = "Row"
	request.PivotTableFieldRequest = pivotTableFieldRequest
	request.NeedReCalculate = true
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PutPivotTableField(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPivotTablesController_PutPivotTableField \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_PutWorksheetPivotTableFilter(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var filter = new(PivotFilter)
	filter.FieldIndex = int64(1)
	filter.FilterType = "Count"

	request := new(PutWorksheetPivotTableFilterRequest)
	request.Name = remoteName
	request.SheetName = "Sheet4"
	request.PivotTableIndex = int64(0)
	request.Filter = filter
	request.NeedReCalculate = true
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PutWorksheetPivotTableFilter(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPivotTablesController_PutWorksheetPivotTableFilter \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_PostPivotTableFieldHideItem(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PostPivotTableFieldHideItemRequest)
	request.Name = remoteName
	request.SheetName = "Sheet4"
	request.PivotTableIndex = int64(0)
	request.PivotFieldType = "Row"
	request.FieldIndex = int64(0)
	request.ItemIndex = int64(1)
	request.IsHide = true
	request.NeedReCalculate = true
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostPivotTableFieldHideItem(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPivotTablesController_PostPivotTableFieldHideItem \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_PostPivotTableFieldMoveTo(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PostPivotTableFieldMoveToRequest)
	request.Name = remoteName
	request.SheetName = "Sheet4"
	request.PivotTableIndex = int64(0)
	request.FieldIndex = int64(0)
	request.From = "Row"
	request.To = "Column"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostPivotTableFieldMoveTo(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPivotTablesController_PostPivotTableFieldMoveTo \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_PostPivotTableCellStyle(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"

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

	request := new(PostPivotTableCellStyleRequest)
	request.Name = remoteName
	request.SheetName = "Sheet4"
	request.PivotTableIndex = int64(0)
	request.Column = int64(1)
	request.Row = int64(1)
	request.Style = style
	request.NeedReCalculate = true
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostPivotTableCellStyle(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPivotTablesController_PostPivotTableCellStyle \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_PostPivotTableStyle(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"

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

	request := new(PostPivotTableStyleRequest)
	request.Name = remoteName
	request.SheetName = "Sheet4"
	request.PivotTableIndex = int64(0)
	request.Style = style
	request.NeedReCalculate = true
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostPivotTableStyle(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPivotTablesController_PostPivotTableStyle \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_PostPivotTableUpdatePivotFields(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var pivotField = new(PivotField)
	pivotField.ShowCompact = true

	request := new(PostPivotTableUpdatePivotFieldsRequest)
	request.Name = remoteName
	request.SheetName = "Sheet4"
	request.PivotTableIndex = int64(0)
	request.PivotFieldType = "Row"
	request.PivotField = pivotField
	request.NeedReCalculate = true
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostPivotTableUpdatePivotFields(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPivotTablesController_PostPivotTableUpdatePivotFields \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_PostPivotTableUpdatePivotField(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var pivotField = new(PivotField)
	pivotField.ShowCompact = true

	request := new(PostPivotTableUpdatePivotFieldRequest)
	request.Name = remoteName
	request.SheetName = "Sheet4"
	request.PivotTableIndex = int64(0)
	request.PivotFieldIndex = int64(0)
	request.PivotFieldType = "Row"
	request.PivotField = pivotField
	request.NeedReCalculate = true
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostPivotTableUpdatePivotField(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPivotTablesController_PostPivotTableUpdatePivotField \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_PostWorksheetPivotTableCalculate(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PostWorksheetPivotTableCalculateRequest)
	request.Name = remoteName
	request.SheetName = "Sheet4"
	request.PivotTableIndex = int64(0)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetPivotTableCalculate(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPivotTablesController_PostWorksheetPivotTableCalculate \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_PostWorksheetPivotTableMove(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PostWorksheetPivotTableMoveRequest)
	request.Name = remoteName
	request.SheetName = "Sheet4"
	request.PivotTableIndex = int64(0)
	request.Row = int64(1)
	request.Column = int64(1)
	request.DestCellName = "C10"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetPivotTableMove(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPivotTablesController_PostWorksheetPivotTableMove \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_DeleteWorksheetPivotTables(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(DeleteWorksheetPivotTablesRequest)
	request.Name = remoteName
	request.SheetName = "Sheet4"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheetPivotTables(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPivotTablesController_DeleteWorksheetPivotTables \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_DeleteWorksheetPivotTable(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(DeleteWorksheetPivotTableRequest)
	request.Name = remoteName
	request.SheetName = "Sheet4"
	request.PivotTableIndex = int64(0)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheetPivotTable(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPivotTablesController_DeleteWorksheetPivotTable \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_DeletePivotTableField(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var pivotTableFieldRequestData = []int64{int64(0)}
	var pivotTableFieldRequest = new(PivotTableFieldRequest)
	pivotTableFieldRequest.Data = pivotTableFieldRequestData

	request := new(DeletePivotTableFieldRequest)
	request.Name = remoteName
	request.SheetName = "Sheet4"
	request.PivotTableIndex = int64(0)
	request.PivotFieldType = "Row"
	request.PivotTableFieldRequest = pivotTableFieldRequest
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.DeletePivotTableField(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPivotTablesController_DeletePivotTableField \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_DeleteWorksheetPivotTableFilters(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(DeleteWorksheetPivotTableFiltersRequest)
	request.Name = remoteName
	request.SheetName = "Sheet3"
	request.PivotTableIndex = int64(0)
	request.NeedReCalculate = true
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheetPivotTableFilters(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPivotTablesController_DeleteWorksheetPivotTableFilters \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_DeleteWorksheetPivotTableFilter(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(DeleteWorksheetPivotTableFilterRequest)
	request.Name = remoteName
	request.SheetName = "Sheet3"
	request.PivotTableIndex = int64(0)
	request.FieldIndex = int64(0)
	request.NeedReCalculate = true
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheetPivotTableFilter(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPivotTablesController_DeleteWorksheetPivotTableFilter \n", GetBaseTest().GetTestNumber())
	}
}
