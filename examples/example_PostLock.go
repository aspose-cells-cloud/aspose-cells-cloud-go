package main

import (
	"os"

	asposecellscloud "github.com/aspose-cells-cloud/aspose-cells-cloud-go"
)
func main() {
	instance := asposecellscloud.NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"), "https://api.aspose.cloud", "v3.0")
    needlockXlsx := "needlock.xlsx"

 
    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[needlockXlsx]= GetBaseTest().localTestDataFolder + needlockXlsx 

    request := new (asposecellscloud.PostLockRequest)
    request.File =         mapFiles    
    request.Password =         "123456"    
    _, httpResponse, err := instance.PostLock(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
