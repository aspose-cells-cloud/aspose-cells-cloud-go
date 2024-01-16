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
 
    var optionsvalue0Font = new(Font)
     optionsvalue0Font.IsBold =  true      
     optionsvalue0Font.Size = int64(16)          
    var optionsvalue0 = new(FontSetting)
     optionsvalue0.Length = int64(5)          
     optionsvalue0.StartIndex = int64(0)          
     optionsvalue0.Font =        optionsvalue0Font      
    var options = []FontSetting   {*       optionsvalue0    }    

    request := new (asposecellscloud.PostCellCharactersRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.CellName =         "E36"    
    request.Options =         options    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := instance.PostCellCharacters(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
