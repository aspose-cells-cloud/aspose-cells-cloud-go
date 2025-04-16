package asposecellscloudtest

import (
	"fmt"
	"testing"

    . "asposecellscloud"
)

func TestShapesController_GetWorksheetShapes(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (GetWorksheetShapesRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetShapes(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestShapesController_GetWorksheetShapes \n", GetBaseTest().GetTestNumber())
	}
}

func TestShapesController_GetWorksheetShape(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (GetWorksheetShapeRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.Shapeindex =  int64(0)        


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.GetWorksheetShape(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestShapesController_GetWorksheetShape \n", GetBaseTest().GetTestNumber())
	}
}

func TestShapesController_PutWorksheetShape(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    var shapeDTO = new(Shape)

    request := new (PutWorksheetShapeRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.ShapeDTO =         shapeDTO    


    request.DrawingType =         "arc"    


    request.UpperLeftRow =  int64(1)        


    request.UpperLeftColumn =  int64(1)        


    request.Top =  int64(10)        


    request.Left =  int64(10)        


    request.Width =  int64(100)        


    request.Height =  int64(100)        


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PutWorksheetShape(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestShapesController_PutWorksheetShape \n", GetBaseTest().GetTestNumber())
	}
}

func TestShapesController_DeleteWorksheetShapes(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (DeleteWorksheetShapesRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheetShapes(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestShapesController_DeleteWorksheetShapes \n", GetBaseTest().GetTestNumber())
	}
}

func TestShapesController_DeleteWorksheetShape(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (DeleteWorksheetShapeRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.Shapeindex =  int64(0)        


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.DeleteWorksheetShape(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestShapesController_DeleteWorksheetShape \n", GetBaseTest().GetTestNumber())
	}
}

func TestShapesController_PostWorksheetShape(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    var dto = new(Shape)
     dto.LowerRightColumn = int64(10)          
    request := new (PostWorksheetShapeRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.Shapeindex =  int64(0)        


    request.Dto =         dto    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetShape(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestShapesController_PostWorksheetShape \n", GetBaseTest().GetTestNumber())
	}
}

func TestShapesController_PostWorksheetGroupShape(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    var listShape = []int64   {int64(0)        ,
    int64(1)        }    
    request := new (PostWorksheetGroupShapeRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet6"    


    request.ListShape =         listShape    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetGroupShape(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestShapesController_PostWorksheetGroupShape \n", GetBaseTest().GetTestNumber())
	}
}

func TestShapesController_PostWorksheetUngroupShape(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (PostWorksheetUngroupShapeRequest)
    request.Name =         remoteName    


    request.SheetName =         "Sheet1"    


    request.Shapeindex =  int64(0)        


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PostWorksheetUngroupShape(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestShapesController_PostWorksheetUngroupShape \n", GetBaseTest().GetTestNumber())
	}
}

