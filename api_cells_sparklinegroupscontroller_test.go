package asposecellscloud

import (
	"fmt"
	"testing"
)

func TestSparklineGroupsController_GetWorksheetSparklineGroups(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "TestCase.xlsx"
    remoteName := "TestCase.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (GetWorksheetSparklineGroupsRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetSparklineGroups(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestSparklineGroupsController_GetWorksheetSparklineGroups \n", GetBaseTest().GetTestNumber())
	}
}

func TestSparklineGroupsController_GetWorksheetSparklineGroup(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "TestCase.xlsx"
    remoteName := "TestCase.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (GetWorksheetSparklineGroupRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.SparklineIndex =  int64(0)        
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetSparklineGroup(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestSparklineGroupsController_GetWorksheetSparklineGroup \n", GetBaseTest().GetTestNumber())
	}
}

func TestSparklineGroupsController_DeleteWorksheetSparklineGroups(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "TestCase.xlsx"
    remoteName := "TestCase.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (DeleteWorksheetSparklineGroupsRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheetSparklineGroups(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestSparklineGroupsController_DeleteWorksheetSparklineGroups \n", GetBaseTest().GetTestNumber())
	}
}

func TestSparklineGroupsController_DeleteWorksheetSparklineGroup(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "TestCase.xlsx"
    remoteName := "TestCase.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (DeleteWorksheetSparklineGroupRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.SparklineIndex =  int64(0)        
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheetSparklineGroup(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestSparklineGroupsController_DeleteWorksheetSparklineGroup \n", GetBaseTest().GetTestNumber())
	}
}

func TestSparklineGroupsController_PutWorksheetSparklineGroup(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "TestCase.xlsx"
    remoteName := "TestCase.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (PutWorksheetSparklineGroupRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.Type_ =         "Line"    
    request.DataRange =         "C6:E13"    
    request.IsVertical =   false    
    request.LocationRange =         "G6:G13"    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.PutWorksheetSparklineGroup(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestSparklineGroupsController_PutWorksheetSparklineGroup \n", GetBaseTest().GetTestNumber())
	}
}

func TestSparklineGroupsController_PostWorksheetSparklineGroup(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "TestCase.xlsx"
    remoteName := "TestCase.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    var sparklineGroup = new(SparklineGroup)
     sparklineGroup.DisplayHidden =  true      
     sparklineGroup.PlotRightToLeft =  true      

    request := new (PostWorksheetSparklineGroupRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.SparklineGroupIndex =  int64(0)        
    request.SparklineGroup =         sparklineGroup    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetSparklineGroup(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestSparklineGroupsController_PostWorksheetSparklineGroup \n", GetBaseTest().GetTestNumber())
	}
}

