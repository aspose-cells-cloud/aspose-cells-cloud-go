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
 
    var colorFilterForegroundColorColor = new(Color)
     colorFilterForegroundColorColor.R = int64(48)          
     colorFilterForegroundColorColor.G = int64(48)          
     colorFilterForegroundColorColor.B = int64(48)          
    var colorFilterForegroundColor = new(CellsColor)
     colorFilterForegroundColor.Type_ =        "Automatic"      
     colorFilterForegroundColor.Color =        colorFilterForegroundColorColor      
    var colorFilter = new(ColorFilterRequest)
     colorFilter.Pattern =        "Solid"      
     colorFilter.ForegroundColor =        colorFilterForegroundColor      

    request := new (asposecellscloud.PutWorksheetColorFilterRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.Range_ =         "A1:B1"    
    request.FieldIndex =  int64(0)        
    request.ColorFilter =         colorFilter    
    request.MatchBlanks =   true    
    request.Refresh =   true    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := instance.PutWorksheetColorFilter(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
