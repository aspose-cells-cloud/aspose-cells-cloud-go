package asposecellscloudtest

import (
	"fmt"
	"testing"

    . "asposecellscloud"
)

func TestConditionalFormattingsController_GetWorksheetConditionalFormattings(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (GetWorksheetConditionalFormattingsRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetConditionalFormattings(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConditionalFormattingsController_GetWorksheetConditionalFormattings \n", GetBaseTest().GetTestNumber())
	}
}

func TestConditionalFormattingsController_GetWorksheetConditionalFormatting(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (GetWorksheetConditionalFormattingRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.Index =  int64(0)        
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetConditionalFormatting(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConditionalFormattingsController_GetWorksheetConditionalFormatting \n", GetBaseTest().GetTestNumber())
	}
}

func TestConditionalFormattingsController_PutWorksheetConditionalFormatting(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    var formatcondition = new(FormatCondition)
     formatcondition.Type_ =        "CellValue"      
     formatcondition.Operator =        "Between"      
     formatcondition.Formula1 =        "v1"      
     formatcondition.Formula2 =        "v2"      

    request := new (PutWorksheetConditionalFormattingRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.Formatcondition =         formatcondition    
    request.CellArea =         "A1:C10"    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.PutWorksheetConditionalFormatting(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConditionalFormattingsController_PutWorksheetConditionalFormatting \n", GetBaseTest().GetTestNumber())
	}
}

func TestConditionalFormattingsController_PutWorksheetFormatCondition(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (PutWorksheetFormatConditionRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.Index =  int64(0)        
    request.CellArea =         "A1:C10"    
    request.Type_ =         "CellValue"    
    request.OperatorType =         "Between"    
    request.Formula1 =         "v1"    
    request.Formula2 =         "v2"    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.PutWorksheetFormatCondition(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConditionalFormattingsController_PutWorksheetFormatCondition \n", GetBaseTest().GetTestNumber())
	}
}

func TestConditionalFormattingsController_PutWorksheetFormatConditionArea(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (PutWorksheetFormatConditionAreaRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.Index =  int64(0)        
    request.CellArea =         "A1:C10"    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.PutWorksheetFormatConditionArea(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConditionalFormattingsController_PutWorksheetFormatConditionArea \n", GetBaseTest().GetTestNumber())
	}
}

func TestConditionalFormattingsController_PutWorksheetFormatConditionCondition(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (PutWorksheetFormatConditionConditionRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.Index =  int64(0)        
    request.Type_ =         "CellValue"    
    request.OperatorType =         "Between"    
    request.Formula1 =         "v1"    
    request.Formula2 =         "v2"    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.PutWorksheetFormatConditionCondition(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConditionalFormattingsController_PutWorksheetFormatConditionCondition \n", GetBaseTest().GetTestNumber())
	}
}

func TestConditionalFormattingsController_DeleteWorksheetConditionalFormattings(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (DeleteWorksheetConditionalFormattingsRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheetConditionalFormattings(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConditionalFormattingsController_DeleteWorksheetConditionalFormattings \n", GetBaseTest().GetTestNumber())
	}
}

func TestConditionalFormattingsController_DeleteWorksheetConditionalFormatting(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (DeleteWorksheetConditionalFormattingRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.Index =  int64(0)        
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheetConditionalFormatting(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConditionalFormattingsController_DeleteWorksheetConditionalFormatting \n", GetBaseTest().GetTestNumber())
	}
}

func TestConditionalFormattingsController_DeleteWorksheetConditionalFormattingArea(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (DeleteWorksheetConditionalFormattingAreaRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.StartRow =  int64(1)        
    request.StartColumn =  int64(1)        
    request.TotalRows =  int64(4)        
    request.TotalColumns =  int64(6)        
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheetConditionalFormattingArea(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConditionalFormattingsController_DeleteWorksheetConditionalFormattingArea \n", GetBaseTest().GetTestNumber())
	}
}

