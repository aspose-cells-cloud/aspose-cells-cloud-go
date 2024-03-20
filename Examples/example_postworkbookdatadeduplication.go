package main

import (
	"os"

	asposecellscloud "github.com/aspose-cells-cloud/aspose-cells-cloud-go"
)
func main() {
	instance := asposecellscloud.NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"), "https://api.aspose.cloud", "v3.0")
    remoteFolder := "TestData/In"

    localName := "BookCsvDuplicateData.csv"
    remoteName := "BookCsvDuplicateData.csv"

    localNameRequest := new(asposecellscloud.UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =   localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    instance.UploadFile(localNameRequest )
 
    var deduplicationRegion = new(DeduplicationRegion)


    request := new (asposecellscloud.PostWorkbookDataDeduplicationRequest)
    request.Name =         remoteName    
    request.DeduplicationRegion =         deduplicationRegion    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := instance.PostWorkbookDataDeduplication(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
