package main

import (
	"os"

	asposecellscloud "github.com/aspose-cells-cloud/aspose-cells-cloud-go"
)
func main() {
	instance := asposecellscloud.NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"), "https://api.aspose.cloud", "v3.0")
    remoteFolder := "TestData/In"

    localBook1 := "Book1.xlsx"
    remoteBook1 := "Book1.xlsx"
    localMyDoc := "myDocument.xlsx"
    remoteMyDoc := "myDocument.xlsx"

    localBook1Request := new(asposecellscloud.UploadFileRequest)
    localBook1Request.UploadFiles = make(map[string]string) 
    localBook1Request.UploadFiles[localBook1] =   localBook1
    localBook1Request.Path = remoteFolder + "/" + remoteBook1 
    localBook1Request.StorageName =""
    instance.UploadFile(localBook1Request )
    localMyDocRequest := new(asposecellscloud.UploadFileRequest)
    localMyDocRequest.UploadFiles = make(map[string]string) 
    localMyDocRequest.UploadFiles[localMyDoc] =   localMyDoc
    localMyDocRequest.Path = remoteFolder + "/" + remoteMyDoc 
    localMyDocRequest.StorageName =""
    instance.UploadFile(localMyDocRequest )
 
    var batchProtectRequestMatchCondition = new(MatchConditionRequest)
     batchProtectRequestMatchCondition.RegexPattern =        "(^Book)(.+)(xlsx$)"      
    var batchProtectRequest = new(BatchProtectRequest)
     batchProtectRequest.SourceFolder =        remoteFolder      
     batchProtectRequest.ProtectionType =        "All"      
     batchProtectRequest.Password =        "123456"      
     batchProtectRequest.OutFolder =        "OutResult"      
     batchProtectRequest.MatchCondition =        batchProtectRequestMatchCondition      

    request := new (asposecellscloud.PostBatchProtectRequest)
    request.BatchProtectRequest =         batchProtectRequest    
    _, httpResponse, err := instance.PostBatchProtect(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
