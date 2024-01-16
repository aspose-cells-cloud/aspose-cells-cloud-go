package main

import (
	"os"

	asposecellscloud "github.com/aspose-cells-cloud/aspose-cells-cloud-go"
)
func main() {
	instance := asposecellscloud.NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"), "https://api.aspose.cloud", "v3.0")
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     compressLevel := 50

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 

    request := new (asposecellscloud.PostCompressRequest)
    request.File =         mapFiles    
    request.CompressLevel =  int64(compressLevel)        
    _, httpResponse, err := instance.PostCompress(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
