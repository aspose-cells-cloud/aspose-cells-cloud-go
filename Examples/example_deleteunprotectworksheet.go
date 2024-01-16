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
 
    var protectParameter = new(ProtectSheetParameter)
     protectParameter.ProtectionType =        "ALL"      
     protectParameter.Password =        "123"      

    request := new (asposecellscloud.DeleteUnprotectWorksheetRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.ProtectParameter =         protectParameter    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := instance.DeleteUnprotectWorksheet(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
