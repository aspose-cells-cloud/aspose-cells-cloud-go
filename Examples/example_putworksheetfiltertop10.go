package main

import (
	"os"

	asposecellscloud "github.com/aspose-cells-cloud/aspose-cells-cloud-go"
)
func main() {
	instance := asposecellscloud.NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"), "https://api.aspose.cloud", "v3.0")
    remoteFolder := "TestData/In"

    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(asposecellscloud.UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =   localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    instance.UploadFile(localNameRequest )
 

    request := new (asposecellscloud.PutWorksheetFilterTop10Request)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.Range_ =         "A1:B1"    
    request.FieldIndex =  int64(0)        
    request.IsTop =   true    
    request.IsPercent =   true    
    request.ItemCount =  int64(1)        
    request.MatchBlanks =   false    
    request.Refresh =   true    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := instance.PutWorksheetFilterTop10(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
