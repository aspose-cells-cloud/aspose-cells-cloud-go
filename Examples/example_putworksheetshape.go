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
 
    var shapeDTO = new(Shape)


    request := new (asposecellscloud.PutWorksheetShapeRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.ShapeDTO =         shapeDTO    
    request.DrawingType =         "arc"    
    request.UpperLeftRow =  int64(1)        
    request.UpperLeftColumn =  int64(1)        
    request.Top =  int64(10)        
    request.Left =  int64(10)        
    request.Width =  int64(100)        
    request.Height =  int64(100)        
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := instance.PutWorksheetShape(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
