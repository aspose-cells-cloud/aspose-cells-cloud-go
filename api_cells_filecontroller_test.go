package asposecellscloud

import (
	"fmt"
	"testing"
)

func TestFileController_DownloadFile(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (DownloadFileRequest)
    request.Path =         remoteFolder + "/" + remoteName    
    request.StorageName =         ""    
    request.VersionId =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.DownloadFile(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestFileController_DownloadFile \n", GetBaseTest().GetTestNumber())
	}
}

func TestFileController_UploadFile(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (UploadFileRequest)
    request.UploadFiles =         mapFiles    
    request.Path =         remoteFolder + "/" + remoteName    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.UploadFile(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestFileController_UploadFile \n", GetBaseTest().GetTestNumber())
	}
}

func TestFileController_CopyFile(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (CopyFileRequest)
    request.SrcPath =         remoteFolder + "/" + remoteName    
    request.DestPath =         "OutResult/" + remoteName    
    request.SrcStorageName =         ""    
    request.DestStorageName =         ""    
    request.VersionId =         ""    
    httpResponse, err := GetBaseTest().CellsApi.CopyFile(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestFileController_CopyFile \n", GetBaseTest().GetTestNumber())
	}
}

func TestFileController_MoveFile(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (MoveFileRequest)
    request.SrcPath =         remoteFolder + "/" + remoteName    
    request.DestPath =         "OutResult/" + remoteName    
    request.SrcStorageName =         ""    
    request.DestStorageName =         ""    
    request.VersionId =         ""    
    httpResponse, err := GetBaseTest().CellsApi.MoveFile(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestFileController_MoveFile \n", GetBaseTest().GetTestNumber())
	}
}

func TestFileController_DeleteFile(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (DeleteFileRequest)
    request.Path =         remoteFolder + "/" + remoteName    
    request.StorageName =         ""    
    request.VersionId =         ""    
    httpResponse, err := GetBaseTest().CellsApi.DeleteFile(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestFileController_DeleteFile \n", GetBaseTest().GetTestNumber())
	}
}

