package asposecellscloud

import (
	"fmt"
	"testing"
)

func TestBatchController_PostBatchConvert(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localBook1 := "Book1.xlsx"
    remoteBook1 := "Book1.xlsx"
    localMyDoc := "myDocument.xlsx"
    remoteMyDoc := "myDocument.xlsx"

    localBook1Request := new(UploadFileRequest)
    localBook1Request.UploadFiles = make(map[string]string) 
    localBook1Request.UploadFiles[localBook1] =  GetBaseTest().localTestDataFolder  + localBook1
    localBook1Request.Path = remoteFolder + "/" + remoteBook1 
    localBook1Request.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localBook1Request )
    localMyDocRequest := new(UploadFileRequest)
    localMyDocRequest.UploadFiles = make(map[string]string) 
    localMyDocRequest.UploadFiles[localMyDoc] =  GetBaseTest().localTestDataFolder  + localMyDoc
    localMyDocRequest.Path = remoteFolder + "/" + remoteMyDoc 
    localMyDocRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localMyDocRequest )
 
    var batchConvertRequestMatchCondition = new(MatchConditionRequest)
     batchConvertRequestMatchCondition.RegexPattern =        "(^Book)(.+)(xlsx$)"      
    var batchConvertRequest = new(BatchConvertRequest)
     batchConvertRequest.SourceFolder =        remoteFolder      
     batchConvertRequest.Format =        "pdf"      
     batchConvertRequest.OutFolder =        "OutResult"      
     batchConvertRequest.MatchCondition =        batchConvertRequestMatchCondition      

    request := new (PostBatchConvertRequest)
    request.BatchConvertRequest =         batchConvertRequest    
    _, httpResponse, err := GetBaseTest().CellsApi.PostBatchConvert(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestBatchController_PostBatchConvert \n", GetBaseTest().GetTestNumber())
	}
}

func TestBatchController_PostBatchProtect(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localBook1 := "Book1.xlsx"
    remoteBook1 := "Book1.xlsx"
    localMyDoc := "myDocument.xlsx"
    remoteMyDoc := "myDocument.xlsx"

    localBook1Request := new(UploadFileRequest)
    localBook1Request.UploadFiles = make(map[string]string) 
    localBook1Request.UploadFiles[localBook1] =  GetBaseTest().localTestDataFolder  + localBook1
    localBook1Request.Path = remoteFolder + "/" + remoteBook1 
    localBook1Request.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localBook1Request )
    localMyDocRequest := new(UploadFileRequest)
    localMyDocRequest.UploadFiles = make(map[string]string) 
    localMyDocRequest.UploadFiles[localMyDoc] =  GetBaseTest().localTestDataFolder  + localMyDoc
    localMyDocRequest.Path = remoteFolder + "/" + remoteMyDoc 
    localMyDocRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localMyDocRequest )
 
    var batchProtectRequestMatchCondition = new(MatchConditionRequest)
     batchProtectRequestMatchCondition.RegexPattern =        "(^Book)(.+)(xlsx$)"      
    var batchProtectRequest = new(BatchProtectRequest)
     batchProtectRequest.SourceFolder =        remoteFolder      
     batchProtectRequest.ProtectionType =        "All"      
     batchProtectRequest.Password =        "123456"      
     batchProtectRequest.OutFolder =        "OutResult"      
     batchProtectRequest.MatchCondition =        batchProtectRequestMatchCondition      

    request := new (PostBatchProtectRequest)
    request.BatchProtectRequest =         batchProtectRequest    
    _, httpResponse, err := GetBaseTest().CellsApi.PostBatchProtect(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestBatchController_PostBatchProtect \n", GetBaseTest().GetTestNumber())
	}
}

func TestBatchController_PostBatchLock(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localBook1 := "Book1.xlsx"
    remoteBook1 := "Book1.xlsx"
    localMyDoc := "myDocument.xlsx"
    remoteMyDoc := "myDocument.xlsx"

    localBook1Request := new(UploadFileRequest)
    localBook1Request.UploadFiles = make(map[string]string) 
    localBook1Request.UploadFiles[localBook1] =  GetBaseTest().localTestDataFolder  + localBook1
    localBook1Request.Path = remoteFolder + "/" + remoteBook1 
    localBook1Request.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localBook1Request )
    localMyDocRequest := new(UploadFileRequest)
    localMyDocRequest.UploadFiles = make(map[string]string) 
    localMyDocRequest.UploadFiles[localMyDoc] =  GetBaseTest().localTestDataFolder  + localMyDoc
    localMyDocRequest.Path = remoteFolder + "/" + remoteMyDoc 
    localMyDocRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localMyDocRequest )
 
    var batchLockRequestMatchCondition = new(MatchConditionRequest)
     batchLockRequestMatchCondition.RegexPattern =        "(^Book)(.+)(xlsx$)"      
    var batchLockRequest = new(BatchLockRequest)
     batchLockRequest.SourceFolder =        remoteFolder      
     batchLockRequest.Password =        "123456"      
     batchLockRequest.OutFolder =        "OutResult"      
     batchLockRequest.MatchCondition =        batchLockRequestMatchCondition      

    request := new (PostBatchLockRequest)
    request.BatchLockRequest =         batchLockRequest    
    _, httpResponse, err := GetBaseTest().CellsApi.PostBatchLock(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestBatchController_PostBatchLock \n", GetBaseTest().GetTestNumber())
	}
}

func TestBatchController_PostBatchUnlock(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localBook1 := "Book1.xlsx"
    remoteBook1 := "Book1.xlsx"
    localMyDoc := "myDocument.xlsx"
    remoteMyDoc := "myDocument.xlsx"

    localBook1Request := new(UploadFileRequest)
    localBook1Request.UploadFiles = make(map[string]string) 
    localBook1Request.UploadFiles[localBook1] =  GetBaseTest().localTestDataFolder  + localBook1
    localBook1Request.Path = remoteFolder + "/" + remoteBook1 
    localBook1Request.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localBook1Request )
    localMyDocRequest := new(UploadFileRequest)
    localMyDocRequest.UploadFiles = make(map[string]string) 
    localMyDocRequest.UploadFiles[localMyDoc] =  GetBaseTest().localTestDataFolder  + localMyDoc
    localMyDocRequest.Path = remoteFolder + "/" + remoteMyDoc 
    localMyDocRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localMyDocRequest )
 
    var batchLockRequestMatchCondition = new(MatchConditionRequest)
     batchLockRequestMatchCondition.RegexPattern =        "(^Book)(.+)(xlsx$)"      
    var batchLockRequest = new(BatchLockRequest)
     batchLockRequest.SourceFolder =        remoteFolder      
     batchLockRequest.Password =        "123456"      
     batchLockRequest.OutFolder =        "OutResult"      
     batchLockRequest.MatchCondition =        batchLockRequestMatchCondition      

    request := new (PostBatchUnlockRequest)
    request.BatchLockRequest =         batchLockRequest    
    _, httpResponse, err := GetBaseTest().CellsApi.PostBatchUnlock(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestBatchController_PostBatchUnlock \n", GetBaseTest().GetTestNumber())
	}
}

