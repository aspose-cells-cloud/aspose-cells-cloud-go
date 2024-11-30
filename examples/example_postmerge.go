package main

import (
	"os"

	asposecellscloud "github.com/aspose-cells-cloud/aspose-cells-cloud-go"
)
func main() {
	instance := asposecellscloud.NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"), "https://api.aspose.cloud", "v3.0")
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "csv"
     mergeToOneSheet := true

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 

    request := new (asposecellscloud.PostMergeRequest)
    request.File =         mapFiles    
    request.OutFormat =         format    
    request.MergeToOneSheet =   mergeToOneSheet    
    _, httpResponse, err := instance.PostMerge(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
