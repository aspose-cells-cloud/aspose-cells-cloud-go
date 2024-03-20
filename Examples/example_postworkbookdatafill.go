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
 
    var dataFillDataFillDefaultValue = new(DataFillValue)
     dataFillDataFillDefaultValue.DefaultDate =        "2024-01-01"      
     dataFillDataFillDefaultValue.DefaultNumber = int64(0)          
     dataFillDataFillDefaultValue.DefaultBoolean =  false      
    var dataFill = new(DataFill)
     dataFill.DataFillDefaultValue =        dataFillDataFillDefaultValue      

    request := new (asposecellscloud.PostWorkbookDataFillRequest)
    request.Name =         remoteName    
    request.DataFill =         dataFill    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := instance.PostWorkbookDataFill(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
