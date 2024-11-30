package asposecellscloud

import (
	. "asposecellscloud"
	"fmt"
	"testing"
)

func TestListObjectsController_GetWorksheetListObjects(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(GetWorksheetListObjectsRequest)
	request.Name = remoteName
	request.SheetName = "Sheet7"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetListObjects(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestListObjectsController_GetWorksheetListObjects \n", GetBaseTest().GetTestNumber())
	}
}

func TestListObjectsController_GetWorksheetListObject(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(GetWorksheetListObjectRequest)
	request.Name = remoteName
	request.SheetName = "Sheet7"
	request.Listobjectindex = int64(0)
	request.Format = "pdf"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetListObject(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestListObjectsController_GetWorksheetListObject \n", GetBaseTest().GetTestNumber())
	}
}

func TestListObjectsController_PutWorksheetListObject(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PutWorksheetListObjectRequest)
	request.Name = remoteName
	request.SheetName = "Sheet7"
	request.StartRow = int64(1)
	request.StartColumn = int64(1)
	request.EndRow = int64(6)
	request.EndColumn = int64(6)
	request.Folder = remoteFolder
	request.HasHeaders = true
	request.DisplayName = "true"
	request.ShowTotals = false
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PutWorksheetListObject(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestListObjectsController_PutWorksheetListObject \n", GetBaseTest().GetTestNumber())
	}
}

func TestListObjectsController_DeleteWorksheetListObjects(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(DeleteWorksheetListObjectsRequest)
	request.Name = remoteName
	request.SheetName = "Sheet7"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheetListObjects(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestListObjectsController_DeleteWorksheetListObjects \n", GetBaseTest().GetTestNumber())
	}
}

func TestListObjectsController_DeleteWorksheetListObject(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(DeleteWorksheetListObjectRequest)
	request.Name = remoteName
	request.SheetName = "Sheet7"
	request.ListObjectIndex = int64(0)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheetListObject(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestListObjectsController_DeleteWorksheetListObject \n", GetBaseTest().GetTestNumber())
	}
}

func TestListObjectsController_PostWorksheetListObject(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var listObject = new(ListObject)
	listObject.ShowHeaderRow = true

	request := new(PostWorksheetListObjectRequest)
	request.Name = remoteName
	request.SheetName = "Sheet7"
	request.ListObjectIndex = int64(0)
	request.ListObject = listObject
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetListObject(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestListObjectsController_PostWorksheetListObject \n", GetBaseTest().GetTestNumber())
	}
}

func TestListObjectsController_PostWorksheetListObjectConvertToRange(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PostWorksheetListObjectConvertToRangeRequest)
	request.Name = remoteName
	request.SheetName = "Sheet7"
	request.ListObjectIndex = int64(0)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetListObjectConvertToRange(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestListObjectsController_PostWorksheetListObjectConvertToRange \n", GetBaseTest().GetTestNumber())
	}
}

func TestListObjectsController_PostWorksheetListObjectSummarizeWithPivotTable(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var createPivotTableRequestPivotFieldColumns = []int64{int64(2)}
	var createPivotTableRequestPivotFieldData = []int64{int64(1)}
	var createPivotTableRequestPivotFieldRows = []int64{int64(0)}
	var createPivotTableRequest = new(CreatePivotTableRequest)
	createPivotTableRequest.DestCellName = "C1"
	createPivotTableRequest.Name = "testp"
	createPivotTableRequest.SourceData = "=Sheet2!A1:E8"
	createPivotTableRequest.UseSameSource = true
	createPivotTableRequest.PivotFieldColumns = createPivotTableRequestPivotFieldColumns
	createPivotTableRequest.PivotFieldData = createPivotTableRequestPivotFieldData
	createPivotTableRequest.PivotFieldRows = createPivotTableRequestPivotFieldRows

	request := new(PostWorksheetListObjectSummarizeWithPivotTableRequest)
	request.Name = remoteName
	request.SheetName = "Sheet7"
	request.ListObjectIndex = int64(0)
	request.DestsheetName = "Sheet2"
	request.CreatePivotTableRequest = createPivotTableRequest
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetListObjectSummarizeWithPivotTable(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestListObjectsController_PostWorksheetListObjectSummarizeWithPivotTable \n", GetBaseTest().GetTestNumber())
	}
}

func TestListObjectsController_PostWorksheetListObjectSortTable(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var dataSorter = new(DataSorter)
	dataSorter.CaseSensitive = true

	request := new(PostWorksheetListObjectSortTableRequest)
	request.Name = remoteName
	request.SheetName = "Sheet7"
	request.ListObjectIndex = int64(0)
	request.DataSorter = dataSorter
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetListObjectSortTable(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestListObjectsController_PostWorksheetListObjectSortTable \n", GetBaseTest().GetTestNumber())
	}
}

func TestListObjectsController_PostWorksheetListColumn(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var listColumn = new(ListColumn)
	listColumn.Name = "test cloumn"

	request := new(PostWorksheetListColumnRequest)
	request.Name = remoteName
	request.SheetName = "Sheet7"
	request.ListObjectIndex = int64(0)
	request.ColumnIndex = int64(0)
	request.ListColumn = listColumn
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetListColumn(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestListObjectsController_PostWorksheetListColumn \n", GetBaseTest().GetTestNumber())
	}
}

func TestListObjectsController_PostWorksheetListColumnsTotal(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var tableTotalRequeststableTotalRequest0 = new(TableTotalRequest)
	tableTotalRequeststableTotalRequest0.ListColumnIndex = int64(1)
	tableTotalRequeststableTotalRequest0.TotalsCalculation = "Average"
	var tableTotalRequests = []TableTotalRequest{*tableTotalRequeststableTotalRequest0}

	request := new(PostWorksheetListColumnsTotalRequest)
	request.Name = remoteName
	request.SheetName = "Sheet7"
	request.ListObjectIndex = int64(0)
	request.TableTotalRequests = tableTotalRequests
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetListColumnsTotal(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestListObjectsController_PostWorksheetListColumnsTotal \n", GetBaseTest().GetTestNumber())
	}
}

func TestListObjectsController_PostWorksheetListObjectRemoveDuplicates(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "TestTables.xlsx"
	remoteName := "TestTables.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PostWorksheetListObjectRemoveDuplicatesRequest)
	request.Name = remoteName
	request.SheetName = "Sheet2"
	request.ListObjectIndex = int64(0)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetListObjectRemoveDuplicates(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestListObjectsController_PostWorksheetListObjectRemoveDuplicates \n", GetBaseTest().GetTestNumber())
	}
}

func TestListObjectsController_PostWorksheetListObjectInsertSlicer(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "TestTables.xlsx"
	remoteName := "TestTables.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PostWorksheetListObjectInsertSlicerRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.ListObjectIndex = int64(0)
	request.ColumnIndex = int64(2)
	request.DestCellName = "j9"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetListObjectInsertSlicer(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestListObjectsController_PostWorksheetListObjectInsertSlicer \n", GetBaseTest().GetTestNumber())
	}
}
