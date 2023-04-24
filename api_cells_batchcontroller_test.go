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
 
    var batchConvertRequestMatchCondition = new(MatchConditionRequest);
     batchConvertRequestMatchCondition.RegexPattern =        "(^Book)(.+)(xlsx$)"      ;
    var batchConvertRequest = new(BatchConvertRequest);
     batchConvertRequest.SourceFolder =        remoteFolder      ;
     batchConvertRequest.Format =        "pdf"      ;
     batchConvertRequest.OutFolder =        "TestResult"      ;
     batchConvertRequest.MatchCondition =        batchConvertRequestMatchCondition      ;

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

