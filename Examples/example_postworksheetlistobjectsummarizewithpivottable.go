package main

import (
	"os"

	asposecellscloud "github.com/aspose-cells-cloud/aspose-cells-cloud-go"
)
func main() {
	instance := asposecellscloud.NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"), "https://api.aspose.cloud", "v3.0")
    remoteFolder := "TestData/In"

    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(asposecellscloud.UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =   localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    instance.UploadFile(localNameRequest )
 
    var createPivotTableRequestPivotFieldColumns = []int64   {int64(2)        }    
    var createPivotTableRequestPivotFieldData = []int64   {int64(1)        }    
    var createPivotTableRequestPivotFieldRows = []int64   {int64(0)        }    
    var createPivotTableRequest = new(CreatePivotTableRequest)
     createPivotTableRequest.DestCellName =        "C1"      
     createPivotTableRequest.Name =        "testp"      
     createPivotTableRequest.SourceData =        "=Sheet2!A1:E8"      
     createPivotTableRequest.UseSameSource =  true      
     createPivotTableRequest.PivotFieldColumns =        createPivotTableRequestPivotFieldColumns      
     createPivotTableRequest.PivotFieldData =        createPivotTableRequestPivotFieldData      
     createPivotTableRequest.PivotFieldRows =        createPivotTableRequestPivotFieldRows      

    request := new (asposecellscloud.PostWorksheetListObjectSummarizeWithPivotTableRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet7"    
    request.ListObjectIndex =  int64(0)        
    request.DestsheetName =         "Sheet2"    
    request.CreatePivotTableRequest =         createPivotTableRequest    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := instance.PostWorksheetListObjectSummarizeWithPivotTable(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
