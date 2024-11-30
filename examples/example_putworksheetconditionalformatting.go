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
 
    var formatcondition = new(FormatCondition)
     formatcondition.Type_ =        "CellValue"      
     formatcondition.Operator =        "Between"      
     formatcondition.Formula1 =        "v1"      
     formatcondition.Formula2 =        "v2"      

    request := new (asposecellscloud.PutWorksheetConditionalFormattingRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.Formatcondition =         formatcondition    
    request.CellArea =         "A1:C10"    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := instance.PutWorksheetConditionalFormatting(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
