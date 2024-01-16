package main

import (
	"os"

	asposecellscloud "github.com/aspose-cells-cloud/aspose-cells-cloud-go"
)
func main() {
	instance := asposecellscloud.NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"), "https://api.aspose.cloud", "v3.0")
    remoteFolder := "TestData/In"

    localName := "TestCase.xlsx"
    remoteName := "TestCase.xlsx"

    localNameRequest := new(asposecellscloud.UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =   localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    instance.UploadFile(localNameRequest )
 
    var sparklineGroup = new(SparklineGroup)
     sparklineGroup.DisplayHidden =  true      
     sparklineGroup.PlotRightToLeft =  true      

    request := new (asposecellscloud.PostWorksheetSparklineGroupRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.SparklineGroupIndex =  int64(0)        
    request.SparklineGroup =         sparklineGroup    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := instance.PostWorksheetSparklineGroup(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
