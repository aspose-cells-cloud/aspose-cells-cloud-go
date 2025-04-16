package asposecellscloudtest

import (
	"fmt"
	"testing"

    . "asposecellscloud"
)

func TestAutoFilterController_GetWorksheetAutoFilter(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (GetWorksheetAutoFilterRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetAutoFilter(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestAutoFilterController_GetWorksheetAutoFilter \n", GetBaseTest().GetTestNumber())
	}
}

func TestAutoFilterController_PutWorksheetDateFilter(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (PutWorksheetDateFilterRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.Range_ =         "A1:B1"    


    request.FieldIndex =  int64(0)        


    request.DateTimeGroupingType =         "Year"    


    request.Year =  int64(1920)        


    request.MatchBlanks =   false    


    request.Refresh =   true    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PutWorksheetDateFilter(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestAutoFilterController_PutWorksheetDateFilter \n", GetBaseTest().GetTestNumber())
	}
}

func TestAutoFilterController_PutWorksheetFilter(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (PutWorksheetFilterRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.Range_ =         "A1:B1"    


    request.FieldIndex =  int64(0)        


    request.Criteria =         "Year"    


    request.MatchBlanks =   false    


    request.Refresh =   true    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PutWorksheetFilter(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestAutoFilterController_PutWorksheetFilter \n", GetBaseTest().GetTestNumber())
	}
}

func TestAutoFilterController_PutWorksheetIconFilter(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (PutWorksheetIconFilterRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.Range_ =         "A1:B1"    


    request.FieldIndex =  int64(0)        


    request.IconSetType =         "ArrowsGray3"    


    request.IconId =  int64(1)        


    request.MatchBlanks =   false    


    request.Refresh =   true    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PutWorksheetIconFilter(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestAutoFilterController_PutWorksheetIconFilter \n", GetBaseTest().GetTestNumber())
	}
}

func TestAutoFilterController_PutWorksheetCustomFilter(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (PutWorksheetCustomFilterRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.Range_ =         "A1:B1"    


    request.FieldIndex =  int64(0)        


    request.OperatorType1 =         "LessOrEqual"    


    request.Criteria1 =         "1"    


    request.MatchBlanks =   false    


    request.Refresh =   true    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PutWorksheetCustomFilter(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestAutoFilterController_PutWorksheetCustomFilter \n", GetBaseTest().GetTestNumber())
	}
}

func TestAutoFilterController_PutWorksheetDynamicFilter(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (PutWorksheetDynamicFilterRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.Range_ =         "A1:B1"    


    request.FieldIndex =  int64(0)        


    request.DynamicFilterType =         "BelowAverage"    


    request.MatchBlanks =   false    


    request.Refresh =   true    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PutWorksheetDynamicFilter(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestAutoFilterController_PutWorksheetDynamicFilter \n", GetBaseTest().GetTestNumber())
	}
}

func TestAutoFilterController_PutWorksheetFilterTop10(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (PutWorksheetFilterTop10Request)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.Range_ =         "A1:B1"    


    request.FieldIndex =  int64(0)        


    request.IsTop =   true    


    request.IsPercent =   true    


    request.ItemCount =  int64(1)        


    request.MatchBlanks =   false    


    request.Refresh =   true    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PutWorksheetFilterTop10(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestAutoFilterController_PutWorksheetFilterTop10 \n", GetBaseTest().GetTestNumber())
	}
}

func TestAutoFilterController_PutWorksheetColorFilter(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    var colorFilterForegroundColorColor = new(Color)
     colorFilterForegroundColorColor.R = int64(48)          
     colorFilterForegroundColorColor.G = int64(48)          
     colorFilterForegroundColorColor.B = int64(48)          
    var colorFilterForegroundColor = new(CellsColor)
     colorFilterForegroundColor.Type_ =        "Automatic"      
     colorFilterForegroundColor.Color =        colorFilterForegroundColorColor      
    var colorFilter = new(ColorFilterRequest)
     colorFilter.Pattern =        "Solid"      
     colorFilter.ForegroundColor =        colorFilterForegroundColor      
    request := new (PutWorksheetColorFilterRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.Range_ =         "A1:B1"    


    request.FieldIndex =  int64(0)        


    request.ColorFilter =         colorFilter    


    request.MatchBlanks =   true    


    request.Refresh =   true    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PutWorksheetColorFilter(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestAutoFilterController_PutWorksheetColorFilter \n", GetBaseTest().GetTestNumber())
	}
}

func TestAutoFilterController_PostWorksheetMatchBlanks(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (PostWorksheetMatchBlanksRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.FieldIndex =  int64(0)        


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetMatchBlanks(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestAutoFilterController_PostWorksheetMatchBlanks \n", GetBaseTest().GetTestNumber())
	}
}

func TestAutoFilterController_PostWorksheetMatchNonBlanks(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (PostWorksheetMatchNonBlanksRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.FieldIndex =  int64(0)        


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetMatchNonBlanks(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestAutoFilterController_PostWorksheetMatchNonBlanks \n", GetBaseTest().GetTestNumber())
	}
}

func TestAutoFilterController_PostWorksheetAutoFilterRefresh(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (PostWorksheetAutoFilterRefreshRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetAutoFilterRefresh(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestAutoFilterController_PostWorksheetAutoFilterRefresh \n", GetBaseTest().GetTestNumber())
	}
}

func TestAutoFilterController_DeleteWorksheetDateFilter(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (DeleteWorksheetDateFilterRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.FieldIndex =  int64(0)        


    request.DateTimeGroupingType =         "Year"    


    request.Year =  int64(1920)        


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheetDateFilter(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestAutoFilterController_DeleteWorksheetDateFilter \n", GetBaseTest().GetTestNumber())
	}
}

func TestAutoFilterController_DeleteWorksheetFilter(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (DeleteWorksheetFilterRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.FieldIndex =  int64(0)        


    request.Criteria =         "year"    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheetFilter(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestAutoFilterController_DeleteWorksheetFilter \n", GetBaseTest().GetTestNumber())
	}
}

