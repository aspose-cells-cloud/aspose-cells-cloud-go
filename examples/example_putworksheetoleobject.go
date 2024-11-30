package main

import (
	"os"

	asposecellscloud "github.com/aspose-cells-cloud/aspose-cells-cloud-go"
)
func main() {
	instance := asposecellscloud.NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"), "https://api.aspose.cloud", "v3.0")
    remoteFolder := "TestData/In"

    localName := "Book1.xlsx"
    oLEDoc := "OLEDoc.docx"
    wordJPG := "word.jpg"
    remoteName := "Book1.xlsx"

    localNameRequest := new(asposecellscloud.UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =   localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    instance.UploadFile(localNameRequest )
    oLEDocRequest := new(asposecellscloud.UploadFileRequest)
    oLEDocRequest.UploadFiles = make(map[string]string) 
    oLEDocRequest.UploadFiles[oLEDoc] =   oLEDoc
    oLEDocRequest.Path = "OLEDoc.docx" 
    oLEDocRequest.StorageName =""
    instance.UploadFile(oLEDocRequest )
    wordJPGRequest := new(asposecellscloud.UploadFileRequest)
    wordJPGRequest.UploadFiles = make(map[string]string) 
    wordJPGRequest.UploadFiles[wordJPG] =   wordJPG
    wordJPGRequest.Path = "word.jpg" 
    wordJPGRequest.StorageName =""
    instance.UploadFile(wordJPGRequest )
 

    request := new (asposecellscloud.PutWorksheetOleObjectRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet6"    
    request.UpperLeftRow =  int64(1)        
    request.UpperLeftColumn =  int64(1)        
    request.Height =  int64(100)        
    request.Width =  int64(80)        
    request.OleFile =         "OLEDoc.docx"    
    request.ImageFile =         "word.jpg"    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := instance.PutWorksheetOleObject(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
