package asposecellscloud

import (
	"fmt"
	"testing"
)

func TestDataProcessingController_PostWorkbookDataCleansing(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "BookCsvDuplicateData.csv"
    remoteName := "BookCsvDuplicateData.csv"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    var dataCleansingDataFillDataFillDefaultValue = new(DataFillValue)
     dataCleansingDataFillDataFillDefaultValue.DefaultDate =        "2024-01-01"      
     dataCleansingDataFillDataFillDefaultValue.DefaultNumber = int64(0)          
     dataCleansingDataFillDataFillDefaultValue.DefaultBoolean =  false      
    var dataCleansingDataFill = new(DataFill)
     dataCleansingDataFill.DataFillDefaultValue =        dataCleansingDataFillDataFillDefaultValue      
    var dataCleansing = new(DataCleansing)
     dataCleansing.NeedFillData =  true      
     dataCleansing.DataFill =        dataCleansingDataFill      

    request := new (PostWorkbookDataCleansingRequest)
    request.Name =         remoteName    
    request.DataCleansing =         dataCleansing    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookDataCleansing(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestDataProcessingController_PostWorkbookDataCleansing \n", GetBaseTest().GetTestNumber())
	}
}

func TestDataProcessingController_PostWorkbookDataDeduplication(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "BookCsvDuplicateData.csv"
    remoteName := "BookCsvDuplicateData.csv"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    var deduplicationRegion = new(DeduplicationRegion)


    request := new (PostWorkbookDataDeduplicationRequest)
    request.Name =         remoteName    
    request.DeduplicationRegion =         deduplicationRegion    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookDataDeduplication(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestDataProcessingController_PostWorkbookDataDeduplication \n", GetBaseTest().GetTestNumber())
	}
}

func TestDataProcessingController_PostWorkbookDataFill(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "BookCsvDuplicateData.csv"
    remoteName := "BookCsvDuplicateData.csv"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    var dataFillDataFillDefaultValue = new(DataFillValue)
     dataFillDataFillDefaultValue.DefaultDate =        "2024-01-01"      
     dataFillDataFillDefaultValue.DefaultNumber = int64(0)          
     dataFillDataFillDefaultValue.DefaultBoolean =  false      
    var dataFill = new(DataFill)
     dataFill.DataFillDefaultValue =        dataFillDataFillDefaultValue      

    request := new (PostWorkbookDataFillRequest)
    request.Name =         remoteName    
    request.DataFill =         dataFill    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookDataFill(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestDataProcessingController_PostWorkbookDataFill \n", GetBaseTest().GetTestNumber())
	}
}

