package asposecellscloud

import (
	"fmt"
	"testing"
)

func TestHypelinksController_GetWorkSheetHyperlinks(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (GetWorkSheetHyperlinksRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.GetWorkSheetHyperlinks(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestHypelinksController_GetWorkSheetHyperlinks \n", GetBaseTest().GetTestNumber())
	}
}

func TestHypelinksController_GetWorkSheetHyperlink(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (GetWorkSheetHyperlinkRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.HyperlinkIndex =  int64(0)        
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.GetWorkSheetHyperlink(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestHypelinksController_GetWorkSheetHyperlink \n", GetBaseTest().GetTestNumber())
	}
}

func TestHypelinksController_DeleteWorkSheetHyperlink(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (DeleteWorkSheetHyperlinkRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.HyperlinkIndex =  int64(0)        
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.DeleteWorkSheetHyperlink(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestHypelinksController_DeleteWorkSheetHyperlink \n", GetBaseTest().GetTestNumber())
	}
}

func TestHypelinksController_PostWorkSheetHyperlink(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    var hyperlink = new(Hyperlink);
     hyperlink.Address =        "https://products.aspose.cloud/cells/"      ;

    request := new (PostWorkSheetHyperlinkRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.HyperlinkIndex =  int64(0)        
    request.Hyperlink =         hyperlink    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkSheetHyperlink(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestHypelinksController_PostWorkSheetHyperlink \n", GetBaseTest().GetTestNumber())
	}
}

func TestHypelinksController_PutWorkSheetHyperlink(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (PutWorkSheetHyperlinkRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.FirstRow =  int64(1)        
    request.FirstColumn =  int64(1)        
    request.TotalRows =  int64(2)        
    request.TotalColumns =  int64(3)        
    request.Address =         "https://products.aspose.cloud/cells/"    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.PutWorkSheetHyperlink(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestHypelinksController_PutWorkSheetHyperlink \n", GetBaseTest().GetTestNumber())
	}
}

func TestHypelinksController_DeleteWorkSheetHyperlinks(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (DeleteWorkSheetHyperlinksRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.DeleteWorkSheetHyperlinks(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestHypelinksController_DeleteWorkSheetHyperlinks \n", GetBaseTest().GetTestNumber())
	}
}

