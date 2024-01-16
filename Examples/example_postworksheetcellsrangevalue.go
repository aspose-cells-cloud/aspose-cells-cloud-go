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
 
    var range_ = new(Range)
     range_.ColumnCount = int64(1)          
     range_.ColumnWidth = 10.0      
     range_.FirstRow = int64(1)          
     range_.RowCount = int64(10)          

    request := new (asposecellscloud.PostWorksheetCellsRangeValueRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.Range_ =         range_    
    request.Value =         "100"    
    request.IsConverted =   true    
    request.SetStyle =   true    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := instance.PostWorksheetCellsRangeValue(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
