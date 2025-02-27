package asposecellscloudtest

import (
	"fmt"
	"testing"

    . "asposecellscloud"
)

func TestChartAreaController_GetChartArea(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (GetChartAreaRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet4"    
    request.ChartIndex =  int64(0)        
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.GetChartArea(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestChartAreaController_GetChartArea \n", GetBaseTest().GetTestNumber())
	}
}

func TestChartAreaController_GetChartAreaFillFormat(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (GetChartAreaFillFormatRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet4"    
    request.ChartIndex =  int64(0)        
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.GetChartAreaFillFormat(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestChartAreaController_GetChartAreaFillFormat \n", GetBaseTest().GetTestNumber())
	}
}

func TestChartAreaController_GetChartAreaBorder(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (GetChartAreaBorderRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet4"    
    request.ChartIndex =  int64(0)        
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.GetChartAreaBorder(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestChartAreaController_GetChartAreaBorder \n", GetBaseTest().GetTestNumber())
	}
}

