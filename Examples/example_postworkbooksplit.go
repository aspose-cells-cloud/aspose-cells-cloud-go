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
 

    request := new (asposecellscloud.PostWorkbookSplitRequest)
    request.Name =         remoteName    
    request.Format =         "png"    
    request.OutFolder =         "OutResult"    
    request.From =  int64(1)        
    request.To =  int64(5)        
    request.HorizontalResolution =  int64(96)        
    request.VerticalResolution =  int64(96)        
    request.SplitNameRule =         "sheetname"    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    request.OutStorageName =         ""    
    _, httpResponse, err := instance.PostWorkbookSplit(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
