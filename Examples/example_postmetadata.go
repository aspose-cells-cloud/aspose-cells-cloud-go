package main

import (
	"os"

	asposecellscloud "github.com/aspose-cells-cloud/aspose-cells-cloud-go"
)
func main() {
	instance := asposecellscloud.NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"), "https://api.aspose.cloud", "v3.0")
    book1Xlsx := "Book1.xlsx"

 
    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
    var cellsDocumentscellsDocument0 = new(CellsDocumentProperty)
     cellsDocumentscellsDocument0.Name =        "Author"      
     cellsDocumentscellsDocument0.Value =        "roy.wang"      
    var cellsDocuments = []CellsDocumentProperty   {*       cellsDocumentscellsDocument0    }    
     mapFiles[book1Xlsx]= GetBaseTest().localTestDataFolder + book1Xlsx 

    request := new (asposecellscloud.PostMetadataRequest)
    request.File =         mapFiles    
    request.CellsDocuments =         cellsDocuments    
    _, httpResponse, err := instance.PostMetadata(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
