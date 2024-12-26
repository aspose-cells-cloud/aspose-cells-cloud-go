package asposecellscloudtest

import (
	"fmt"
	"testing"

	. "github.com/aspose-cells-cloud/aspose-cells-cloud-go/v24"
)

func TestWorksheetValidationsController_GetWorksheetValidations(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(GetWorksheetValidationsRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetValidations(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetValidationsController_GetWorksheetValidations \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetValidationsController_GetWorksheetValidation(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(GetWorksheetValidationRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.ValidationIndex = int64(0)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetValidation(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetValidationsController_GetWorksheetValidation \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetValidationsController_PutWorksheetValidation(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PutWorksheetValidationRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.Range_ = "A1:C10"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PutWorksheetValidation(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetValidationsController_PutWorksheetValidation \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetValidationsController_PostWorksheetValidation(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var validation = new(Validation)
	validation.Formula1 = "=A1"
	validation.Type_ = "Custom"

	request := new(PostWorksheetValidationRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.ValidationIndex = int64(0)
	request.Validation = validation
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetValidation(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetValidationsController_PostWorksheetValidation \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetValidationsController_DeleteWorksheetValidation(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(DeleteWorksheetValidationRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.ValidationIndex = int64(0)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheetValidation(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetValidationsController_DeleteWorksheetValidation \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetValidationsController_DeleteWorksheetValidations(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(DeleteWorksheetValidationsRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheetValidations(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetValidationsController_DeleteWorksheetValidations \n", GetBaseTest().GetTestNumber())
	}
}
