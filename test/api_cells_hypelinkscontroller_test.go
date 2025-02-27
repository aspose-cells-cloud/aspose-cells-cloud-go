package asposecellscloudtest

import (
	"fmt"
	"testing"

    . "asposecellscloud"
)

func TestHypelinksController_GetWorksheetHyperlinks(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (GetWorksheetHyperlinksRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetHyperlinks(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestHypelinksController_GetWorksheetHyperlinks \n", GetBaseTest().GetTestNumber())
	}
}

func TestHypelinksController_GetWorksheetHyperlink(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (GetWorksheetHyperlinkRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.HyperlinkIndex =  int64(0)        
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetHyperlink(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestHypelinksController_GetWorksheetHyperlink \n", GetBaseTest().GetTestNumber())
	}
}

func TestHypelinksController_DeleteWorksheetHyperlink(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (DeleteWorksheetHyperlinkRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.HyperlinkIndex =  int64(0)        
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheetHyperlink(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestHypelinksController_DeleteWorksheetHyperlink \n", GetBaseTest().GetTestNumber())
	}
}

func TestHypelinksController_PostWorksheetHyperlink(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    var hyperlink = new(Hyperlink)
     hyperlink.Address =        "https://products.aspose.cloud/cells/"      

    request := new (PostWorksheetHyperlinkRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.HyperlinkIndex =  int64(0)        
    request.Hyperlink =         hyperlink    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetHyperlink(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestHypelinksController_PostWorksheetHyperlink \n", GetBaseTest().GetTestNumber())
	}
}

func TestHypelinksController_PutWorksheetHyperlink(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (PutWorksheetHyperlinkRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.FirstRow =  int64(1)        
    request.FirstColumn =  int64(1)        
    request.TotalRows =  int64(2)        
    request.TotalColumns =  int64(3)        
    request.Address =         "https://products.aspose.cloud/cells/"    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.PutWorksheetHyperlink(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestHypelinksController_PutWorksheetHyperlink \n", GetBaseTest().GetTestNumber())
	}
}

func TestHypelinksController_DeleteWorksheetHyperlinks(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (DeleteWorksheetHyperlinksRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheetHyperlinks(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestHypelinksController_DeleteWorksheetHyperlinks \n", GetBaseTest().GetTestNumber())
	}
}

