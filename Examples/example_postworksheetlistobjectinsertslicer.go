package main

import (
	"os"

	asposecellscloud "github.com/aspose-cells-cloud/aspose-cells-cloud-go"
)
func main() {
	instance := asposecellscloud.NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"), "https://api.aspose.cloud", "v3.0")
    remoteFolder := "TestData/In"

    localName := "TestTables.xlsx"
    remoteName := "TestTables.xlsx"

    localNameRequest := new(asposecellscloud.UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =   localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    instance.UploadFile(localNameRequest )
 

    request := new (asposecellscloud.PostWorksheetListObjectInsertSlicerRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.ListObjectIndex =  int64(0)        
    request.ColumnIndex =  int64(2)        
    request.DestCellName =         "j9"    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := instance.PostWorksheetListObjectInsertSlicer(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
