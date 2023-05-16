package asposecellscloud

import (
	"fmt"
	"testing"
)

func TestPicturesController_GetWorksheetPictures(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(GetWorksheetPicturesRequest)
	request.Name = remoteName
	request.SheetName = "Sheet6"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetPictures(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPicturesController_GetWorksheetPictures \n", GetBaseTest().GetTestNumber())
	}
}

func TestPicturesController_GetWorksheetPictureWithFormat(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(GetWorksheetPictureWithFormatRequest)
	request.Name = remoteName
	request.SheetName = "Sheet6"
	request.PictureNumber = int64(0)
	request.Format = "png"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetPictureWithFormat(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPicturesController_GetWorksheetPictureWithFormat \n", GetBaseTest().GetTestNumber())
	}
}

func TestPicturesController_PutWorksheetAddPicture(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	waterMarkPNG := "WaterMark.png"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)
	waterMarkPNGRequest := new(UploadFileRequest)
	waterMarkPNGRequest.UploadFiles = make(map[string]string)
	waterMarkPNGRequest.UploadFiles[waterMarkPNG] = GetBaseTest().localTestDataFolder + waterMarkPNG
	waterMarkPNGRequest.Path = remoteFolder + "/WaterMark.png"
	waterMarkPNGRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(waterMarkPNGRequest)

	request := new(PutWorksheetAddPictureRequest)
	request.Name = remoteName
	request.SheetName = "Sheet6"
	request.UpperLeftRow = int64(1)
	request.UpperLeftColumn = int64(1)
	request.LowerRightRow = int64(10)
	request.LowerRightColumn = int64(10)
	request.PicturePath = remoteFolder + "/WaterMark.png"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PutWorksheetAddPicture(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPicturesController_PutWorksheetAddPicture \n", GetBaseTest().GetTestNumber())
	}
}

func TestPicturesController_PostWorksheetPicture(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var picture = new(Picture)
	picture.Left = int64(10)
	picture.Bottom = int64(10)

	request := new(PostWorksheetPictureRequest)
	request.Name = remoteName
	request.SheetName = "Sheet6"
	request.PictureIndex = int64(0)
	request.Picture = picture
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetPicture(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPicturesController_PostWorksheetPicture \n", GetBaseTest().GetTestNumber())
	}
}

func TestPicturesController_DeleteWorksheetPicture(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(DeleteWorksheetPictureRequest)
	request.Name = remoteName
	request.SheetName = "Sheet6"
	request.PictureIndex = int64(0)
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheetPicture(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPicturesController_DeleteWorksheetPicture \n", GetBaseTest().GetTestNumber())
	}
}

func TestPicturesController_DeleteWorksheetPictures(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(DeleteWorksheetPicturesRequest)
	request.Name = remoteName
	request.SheetName = "Sheet6"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheetPictures(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPicturesController_DeleteWorksheetPictures \n", GetBaseTest().GetTestNumber())
	}
}
