package main

import (
	"os"

	asposecellscloud "github.com/aspose-cells-cloud/aspose-cells-cloud-go"
)
func main() {
	instance := asposecellscloud.NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"), "https://api.aspose.cloud", "v3.0")
    remoteFolder := "TestData/In"

    localName := "Book1.xlsx"
    myDocumentXLSX := "myDocument.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(asposecellscloud.UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =   localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    instance.UploadFile(localNameRequest )
    myDocumentXLSXRequest := new(asposecellscloud.UploadFileRequest)
    myDocumentXLSXRequest.UploadFiles = make(map[string]string) 
    myDocumentXLSXRequest.UploadFiles[myDocumentXLSX] =   myDocumentXLSX
    myDocumentXLSXRequest.Path = remoteFolder + "/myDocument.xlsx" 
    myDocumentXLSXRequest.StorageName =""
    instance.UploadFile(myDocumentXLSXRequest )
 

    request := new (asposecellscloud.PostWorkbooksMergeRequest)
    request.Name =         remoteName    
    request.MergeWith =         remoteFolder + "/myDocument.xlsx"    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    request.MergedStorageName =         ""    
    _, httpResponse, err := instance.PostWorkbooksMerge(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
