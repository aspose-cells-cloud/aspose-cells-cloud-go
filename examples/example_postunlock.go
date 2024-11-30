package main

import (
	"os"

	asposecellscloud "github.com/aspose-cells-cloud/aspose-cells-cloud-go"
)
func main() {
	instance := asposecellscloud.NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"), "https://api.aspose.cloud", "v3.0")
    needUnlockXlsx := "needUnlock.xlsx"

 
    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[needUnlockXlsx]= GetBaseTest().localTestDataFolder + needUnlockXlsx 

    request := new (asposecellscloud.PostUnlockRequest)
    request.File =         mapFiles    
    request.Password =         "123456"    
    _, httpResponse, err := instance.PostUnlock(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
