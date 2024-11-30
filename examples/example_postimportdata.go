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
 
    var importOptionData = []int64   {int64(1)        ,
    int64(2)        ,
    int64(3)        ,
    int64(4)        }    
    var importOption = new(ImportIntArrayOption)
     importOption.DestinationWorksheet =        "Sheet1"      
     importOption.FirstColumn = int64(1)          
     importOption.FirstRow = int64(3)          
     importOption.ImportDataType =        "IntArray"      
     importOption.IsInsert =  true      
     importOption.IsVertical =  true      
     importOption.Data =        importOptionData      

    request := new (asposecellscloud.PostImportDataRequest)
    request.Name =         remoteName    
    request.ImportOption =         importOption    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := instance.PostImportData(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
