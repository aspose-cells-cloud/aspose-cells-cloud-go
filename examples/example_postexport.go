package main

import (
	"os"

	asposecellscloud "github.com/aspose-cells-cloud/aspose-cells-cloud-go"
)
func main() {
	instance := asposecellscloud.NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"), "https://api.aspose.cloud", "v3.0")
    assemblyTestXlsx := "assemblytest.xlsx"
    book1Xlsx := "Book1.xlsx"

 
     format := "csv"
     objectType := "workbook"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[book1Xlsx]= GetBaseTest().localTestDataFolder + book1Xlsx 

    request := new (asposecellscloud.PostExportRequest)
    request.File =         mapFiles    
    request.ObjectType =         objectType    
    request.Format =         format    
    _, httpResponse, err := instance.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
