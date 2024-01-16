package main

import (
	"os"

	asposecellscloud "github.com/aspose-cells-cloud/aspose-cells-cloud-go"
)
func main() {
	instance := asposecellscloud.NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"), "https://api.aspose.cloud", "v3.0")
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 

    request := new (asposecellscloud.PostSearchRequest)
    request.File =         mapFiles    
    request.Text =         "12"    
    _, httpResponse, err := instance.PostSearch(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
