package asposecellscloudtest

import (
	"fmt"
	"testing"

	. "github.com/aspose-cells-cloud/aspose-cells-cloud-go/v25"
)

func TestOleObjectsController_GetWorksheetOleObjects(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(GetWorksheetOleObjectsRequest)
	request.Name = remoteName
	request.SheetName = "Sheet6"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetOleObjects(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestOleObjectsController_GetWorksheetOleObjects \n", GetBaseTest().GetTestNumber())
	}
}

func TestOleObjectsController_GetWorksheetOleObject(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(GetWorksheetOleObjectRequest)
	request.Name = remoteName
	request.SheetName = "Sheet6"
	request.ObjectNumber = int64(0)
	request.Format = "png"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetOleObject(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestOleObjectsController_GetWorksheetOleObject \n", GetBaseTest().GetTestNumber())
	}
}

func TestOleObjectsController_DeleteWorksheetOleObjects(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(DeleteWorksheetOleObjectsRequest)
	request.Name = remoteName
	request.SheetName = "Sheet6"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheetOleObjects(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestOleObjectsController_DeleteWorksheetOleObjects \n", GetBaseTest().GetTestNumber())
	}
}

func TestOleObjectsController_DeleteWorksheetOleObject(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(DeleteWorksheetOleObjectRequest)
	request.Name = remoteName
	request.SheetName = "Sheet6"
	request.OleObjectIndex = int64(0)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheetOleObject(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestOleObjectsController_DeleteWorksheetOleObject \n", GetBaseTest().GetTestNumber())
	}
}

func TestOleObjectsController_PostUpdateWorksheetOleObject(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var ole = new(OleObject)
	ole.Left = int64(10)
	ole.Right = int64(10)
	ole.Height = int64(90)
	ole.Width = int64(78)

	request := new(PostUpdateWorksheetOleObjectRequest)
	request.Name = remoteName
	request.SheetName = "Sheet6"
	request.OleObjectIndex = int64(0)
	request.Ole = ole
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostUpdateWorksheetOleObject(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestOleObjectsController_PostUpdateWorksheetOleObject \n", GetBaseTest().GetTestNumber())
	}
}

func TestOleObjectsController_PutWorksheetOleObject(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	oLEDoc := "OLEDoc.docx"
	wordJPG := "word.jpg"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)
	oLEDocRequest := new(UploadFileRequest)
	oLEDocRequest.UploadFiles = make(map[string]string)
	oLEDocRequest.UploadFiles[oLEDoc] = GetBaseTest().localTestDataFolder + oLEDoc
	oLEDocRequest.Path = "OLEDoc.docx"
	oLEDocRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(oLEDocRequest)
	wordJPGRequest := new(UploadFileRequest)
	wordJPGRequest.UploadFiles = make(map[string]string)
	wordJPGRequest.UploadFiles[wordJPG] = GetBaseTest().localTestDataFolder + wordJPG
	wordJPGRequest.Path = "word.jpg"
	wordJPGRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(wordJPGRequest)

	request := new(PutWorksheetOleObjectRequest)
	request.Name = remoteName
	request.SheetName = "Sheet6"
	request.UpperLeftRow = int64(1)
	request.UpperLeftColumn = int64(1)
	request.Height = int64(100)
	request.Width = int64(80)
	request.OleFile = "OLEDoc.docx"
	request.ImageFile = "word.jpg"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PutWorksheetOleObject(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestOleObjectsController_PutWorksheetOleObject \n", GetBaseTest().GetTestNumber())
	}
}
