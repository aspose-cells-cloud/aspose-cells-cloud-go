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
     mapFiles[book1Xlsx]= GetBaseTest().localTestDataFolder + book1Xlsx 

    request := new (asposecellscloud.DeleteMetadataRequest)
    request.File =         mapFiles    
    request.Type_ =         "all"    
    _, httpResponse, err := instance.DeleteMetadata(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
