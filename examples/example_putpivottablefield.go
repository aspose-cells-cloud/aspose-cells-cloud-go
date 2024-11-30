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
 
    var pivotTableFieldRequestData = []int64   {int64(0)        }    
    var pivotTableFieldRequest = new(PivotTableFieldRequest)
     pivotTableFieldRequest.Data =        pivotTableFieldRequestData      

    request := new (asposecellscloud.PutPivotTableFieldRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet4"    
    request.PivotTableIndex =  int64(0)        
    request.PivotFieldType =         "Row"    
    request.PivotTableFieldRequest =         pivotTableFieldRequest    
    request.NeedReCalculate =   true    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := instance.PutPivotTableField(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
