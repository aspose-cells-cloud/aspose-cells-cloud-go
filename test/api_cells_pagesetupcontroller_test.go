package asposecellscloudtest

import (
	"fmt"
	"testing"

	. "github.com/aspose-cells-cloud/aspose-cells-cloud-go/v24"
)

func TestPageSetupController_GetPageSetup(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(GetPageSetupRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.GetPageSetup(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPageSetupController_GetPageSetup \n", GetBaseTest().GetTestNumber())
	}
}

func TestPageSetupController_PostPageSetup(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var pageSetup = new(PageSetup)
	pageSetup.BlackAndWhite = true
	pageSetup.CenterHorizontally = true
	pageSetup.CenterVertically = true

	request := new(PostPageSetupRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.PageSetup = pageSetup
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostPageSetup(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPageSetupController_PostPageSetup \n", GetBaseTest().GetTestNumber())
	}
}

func TestPageSetupController_DeleteHeaderFooter(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(DeleteHeaderFooterRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.DeleteHeaderFooter(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPageSetupController_DeleteHeaderFooter \n", GetBaseTest().GetTestNumber())
	}
}

func TestPageSetupController_GetHeader(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(GetHeaderRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.GetHeader(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPageSetupController_GetHeader \n", GetBaseTest().GetTestNumber())
	}
}

func TestPageSetupController_PostHeader(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PostHeaderRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.Section = int64(1)
	request.Script = "Update add header"
	request.IsFirstPage = true
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostHeader(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPageSetupController_PostHeader \n", GetBaseTest().GetTestNumber())
	}
}

func TestPageSetupController_GetFooter(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(GetFooterRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.GetFooter(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPageSetupController_GetFooter \n", GetBaseTest().GetTestNumber())
	}
}

func TestPageSetupController_PostFooter(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PostFooterRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.Section = int64(1)
	request.Script = "add footer script"
	request.IsFirstPage = true
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostFooter(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPageSetupController_PostFooter \n", GetBaseTest().GetTestNumber())
	}
}

func TestPageSetupController_PostFitWideToPages(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PostFitWideToPagesRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostFitWideToPages(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPageSetupController_PostFitWideToPages \n", GetBaseTest().GetTestNumber())
	}
}

func TestPageSetupController_PostFitTallToPages(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	request := new(PostFitTallToPagesRequest)
	request.Name = remoteName
	request.SheetName = "Sheet1"
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostFitTallToPages(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestPageSetupController_PostFitTallToPages \n", GetBaseTest().GetTestNumber())
	}
}
