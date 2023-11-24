package asposecellscloud

import (
	"fmt"
	"testing"
)

func TestFolderController_GetFilesList(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (GetFilesListRequest)
    request.Path =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.GetFilesList(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestFolderController_GetFilesList \n", GetBaseTest().GetTestNumber())
	}
}

func TestFolderController_CreateFolder(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (CreateFolderRequest)
    request.Path =         "OutResult/NewFolder"    
    request.StorageName =         ""    
    httpResponse, err := GetBaseTest().CellsApi.CreateFolder(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestFolderController_CreateFolder \n", GetBaseTest().GetTestNumber())
	}
}

func TestFolderController_CopyFolder(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (CopyFolderRequest)
    request.SrcPath =         remoteFolder    
    request.DestPath =         "OutResult/Create"    
    request.SrcStorageName =         ""    
    request.DestStorageName =         ""    
    httpResponse, err := GetBaseTest().CellsApi.CopyFolder(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestFolderController_CopyFolder \n", GetBaseTest().GetTestNumber())
	}
}

func TestFolderController_MoveFolder(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (MoveFolderRequest)
    request.SrcPath =         "OutResult/Create"    
    request.DestPath =         "OutResult/Move"    
    request.SrcStorageName =         ""    
    request.DestStorageName =         ""    
    httpResponse, err := GetBaseTest().CellsApi.MoveFolder(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestFolderController_MoveFolder \n", GetBaseTest().GetTestNumber())
	}
}

func TestFolderController_DeleteFolder(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (DeleteFolderRequest)
    request.Path =         "OutResult/Create"    
    request.StorageName =         ""    
    request.Recursive =   true    
    httpResponse, err := GetBaseTest().CellsApi.DeleteFolder(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestFolderController_DeleteFolder \n", GetBaseTest().GetTestNumber())
	}
}

