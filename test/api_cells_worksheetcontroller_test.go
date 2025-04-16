package asposecellscloudtest

import (
	"fmt"
	"testing"

    . "asposecellscloud"
)

func TestWorksheetController_GetWorksheets(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (GetWorksheetsRequest)
    request.Name =         remoteName    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.GetWorksheets(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_GetWorksheets \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_GetWorksheetWithFormat(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (GetWorksheetWithFormatRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.Format =         "png"    


    request.PageIndex =  int64(0)        


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetWithFormat(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_GetWorksheetWithFormat \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_PutChangeVisibilityWorksheet(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (PutChangeVisibilityWorksheetRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.IsVisible =   true    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PutChangeVisibilityWorksheet(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_PutChangeVisibilityWorksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_PutActiveWorksheet(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (PutActiveWorksheetRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PutActiveWorksheet(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_PutActiveWorksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_PutInsertNewWorksheet(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (PutInsertNewWorksheetRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.Index =  int64(1)        


    request.Sheettype =         "VB"    


    request.Newsheetname =         "VBASheet"    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PutInsertNewWorksheet(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_PutInsertNewWorksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_PutAddNewWorksheet(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (PutAddNewWorksheetRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.Position =  int64(0)        


    request.Sheettype =         "VB"    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PutAddNewWorksheet(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_PutAddNewWorksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_DeleteWorksheet(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (DeleteWorksheetRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheet(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_DeleteWorksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_DeleteWorksheets(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    var matchCondition = new(MatchConditionRequest)
     matchCondition.RegexPattern =        "{*}"      
    request := new (DeleteWorksheetsRequest)
    request.Name =         remoteName    


    request.MatchCondition =         matchCondition    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheets(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_DeleteWorksheets \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_PostMoveWorksheet(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    var moving = new(WorksheetMovingRequest)
     moving.DestinationWorksheet =        "Sheet4"      
     moving.Position =        "After"      
    request := new (PostMoveWorksheetRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.Moving =         moving    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PostMoveWorksheet(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_PostMoveWorksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_PutProtectWorksheet(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    var protectParameter = new(ProtectSheetParameter)
     protectParameter.ProtectionType =        "ALL"      
     protectParameter.Password =        "123"      
    request := new (PutProtectWorksheetRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.ProtectParameter =         protectParameter    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PutProtectWorksheet(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_PutProtectWorksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_DeleteUnprotectWorksheet(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    var protectParameter = new(ProtectSheetParameter)
     protectParameter.ProtectionType =        "ALL"      
     protectParameter.Password =        "123"      
    request := new (DeleteUnprotectWorksheetRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.ProtectParameter =         protectParameter    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.DeleteUnprotectWorksheet(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_DeleteUnprotectWorksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_GetWorksheetTextItems(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (GetWorksheetTextItemsRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetTextItems(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_GetWorksheetTextItems \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_GetWorksheetComments(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (GetWorksheetCommentsRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetComments(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_GetWorksheetComments \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_GetWorksheetComment(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (GetWorksheetCommentRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.CellName =         "B3"    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetComment(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_GetWorksheetComment \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_PutWorksheetComment(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    var comment = new(Comment)
     comment.Author =        "aspose cells developer"      
     comment.Note =        "aspose cells cloud api add comment."      
    request := new (PutWorksheetCommentRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.CellName =         "C1"    


    request.Comment =         comment    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PutWorksheetComment(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_PutWorksheetComment \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_PostWorksheetComment(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    var comment = new(Comment)
     comment.Author =        "aspose cells developer"      
     comment.Note =        "aspose cells cloud api update comment."      
    request := new (PostWorksheetCommentRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.CellName =         "B3"    


    request.Comment =         comment    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetComment(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_PostWorksheetComment \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_DeleteWorksheetComment(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (DeleteWorksheetCommentRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.CellName =         "B3"    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheetComment(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_DeleteWorksheetComment \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_DeleteWorksheetComments(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (DeleteWorksheetCommentsRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheetComments(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_DeleteWorksheetComments \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_GetWorksheetMergedCells(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (GetWorksheetMergedCellsRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetMergedCells(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_GetWorksheetMergedCells \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_GetWorksheetMergedCell(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (GetWorksheetMergedCellRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.MergedCellIndex =  int64(0)        


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetMergedCell(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_GetWorksheetMergedCell \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_GetWorksheetCalculateFormula(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (GetWorksheetCalculateFormulaRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.Formula =         "=NOW()"    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetCalculateFormula(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_GetWorksheetCalculateFormula \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_PostWorksheetCalculateFormula(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (PostWorksheetCalculateFormulaRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.Formula =         "=NOW()"    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetCalculateFormula(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_PostWorksheetCalculateFormula \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_PostWorksheetTextSearch(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (PostWorksheetTextSearchRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.Text =         "123"    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetTextSearch(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_PostWorksheetTextSearch \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_PostWorksheetTextReplace(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (PostWorksheetTextReplaceRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.OldValue =         "123"    


    request.NewValue =         "456"    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetTextReplace(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_PostWorksheetTextReplace \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_PostWorksheetRangeSort(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    var dataSorter = new(DataSorter)
     dataSorter.CaseSensitive =  true      
    request := new (PostWorksheetRangeSortRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.CellArea =         "A1:C10"    


    request.DataSorter =         dataSorter    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetRangeSort(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_PostWorksheetRangeSort \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_PostAutofitWorksheetRow(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (PostAutofitWorksheetRowRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.RowIndex =  int64(1)        


    request.FirstColumn =  int64(1)        


    request.LastColumn =  int64(8)        


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PostAutofitWorksheetRow(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_PostAutofitWorksheetRow \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_PostAutofitWorksheetRows(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (PostAutofitWorksheetRowsRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.StartRow =  int64(1)        


    request.EndRow =  int64(9)        


    request.OnlyAuto =   true    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PostAutofitWorksheetRows(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_PostAutofitWorksheetRows \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_PostAutofitWorksheetColumns(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (PostAutofitWorksheetColumnsRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.StartColumn =  int64(1)        


    request.EndColumn =  int64(9)        


    request.OnlyAuto =   true    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PostAutofitWorksheetColumns(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_PostAutofitWorksheetColumns \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_PutWorksheetBackground(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    waterMarkPNG := "WaterMark.png"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
    waterMarkPNGRequest := new(UploadFileRequest)
    waterMarkPNGRequest.UploadFiles = make(map[string]string) 
    waterMarkPNGRequest.UploadFiles[waterMarkPNG] =  GetBaseTest().localTestDataFolder  + waterMarkPNG
    waterMarkPNGRequest.Path = remoteFolder + "/WaterMark.png" 
    waterMarkPNGRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(waterMarkPNGRequest )
 
    request := new (PutWorksheetBackgroundRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.PicPath =         remoteFolder + "/WaterMark.png"    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PutWorksheetBackground(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_PutWorksheetBackground \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_DeleteWorksheetBackground(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (DeleteWorksheetBackgroundRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheetBackground(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_DeleteWorksheetBackground \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_PutWorksheetFreezePanes(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (PutWorksheetFreezePanesRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.Row =  int64(1)        


    request.Column =  int64(1)        


    request.FreezedRows =  int64(4)        


    request.FreezedColumns =  int64(5)        


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PutWorksheetFreezePanes(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_PutWorksheetFreezePanes \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_DeleteWorksheetFreezePanes(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (DeleteWorksheetFreezePanesRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.Row =  int64(1)        


    request.Column =  int64(1)        


    request.FreezedRows =  int64(4)        


    request.FreezedColumns =  int64(5)        


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheetFreezePanes(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_DeleteWorksheetFreezePanes \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_PostCopyWorksheet(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    var options = new(CopyOptions)
     options.ColumnCharacterWidth =  true      
    request := new (PostCopyWorksheetRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet15"    


    request.SourceSheet =         "Sheet6"    


    request.Options =         options    


    request.SourceWorkbook =         ""    


    request.SourceFolder =         ""    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PostCopyWorksheet(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_PostCopyWorksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_PostRenameWorksheet(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (PostRenameWorksheetRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet5"    


    request.Newname =         "Sheet55"    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PostRenameWorksheet(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_PostRenameWorksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_PostUpdateWorksheetProperty(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    var sheet = new(Worksheet)
     sheet.Name =        "sheet65"      
     sheet.IsGridlinesVisible =  true      
    request := new (PostUpdateWorksheetPropertyRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet5"    


    request.Sheet =         sheet    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PostUpdateWorksheetProperty(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_PostUpdateWorksheetProperty \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_GetNamedRanges(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (GetNamedRangesRequest)
    request.Name =         remoteName    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.GetNamedRanges(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_GetNamedRanges \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_GetNamedRangeValue(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (GetNamedRangeValueRequest)
    request.Name =         remoteName    


    request.Namerange =         "Name_2"    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.GetNamedRangeValue(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_GetNamedRangeValue \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_PostUpdateWorksheetZoom(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (PostUpdateWorksheetZoomRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.Value =  int64(90)        


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PostUpdateWorksheetZoom(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_PostUpdateWorksheetZoom \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorksheetController_GetWorksheetPageCount(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (GetWorksheetPageCountRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetPageCount(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_GetWorksheetPageCount \n", GetBaseTest().GetTestNumber())
	}
}

