package main

import (
	"os"

	asposecellscloud "github.com/aspose-cells-cloud/aspose-cells-cloud-go"
)
func main() {
	instance := asposecellscloud.NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"), "https://api.aspose.cloud", "v3.0")
    remoteFolder := "TestData/In"

    localName := "TestCase.xlsx"
    remoteName := "TestCase.xlsx"

    localNameRequest := new(asposecellscloud.UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =   localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    instance.UploadFile(localNameRequest )
 

    request := new (asposecellscloud.PutWorksheetPivotTableRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet4"    
    request.Folder =         remoteFolder    
    request.SourceData =         "=Sheet1!C6:E13"    
    request.DestCellName =         "C1"    
    request.TableName =         "TestPivot"    
    request.UseSameSource =   true    
    request.StorageName =         ""    
    _, httpResponse, err := instance.PutWorksheetPivotTable(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
