package main

import (
	"os"

	asposecellscloud "github.com/aspose-cells-cloud/aspose-cells-cloud-go"
)
func main() {
	instance := asposecellscloud.NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"), "https://api.aspose.cloud", "v3.0")
    remoteFolder := "TestData/In"

    localName := "BookText.xlsx"
    remoteName := "BookText.xlsx"

    localNameRequest := new(asposecellscloud.UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =   localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    instance.UploadFile(localNameRequest )
 
    var addTextOptionsDataSource = new(DataSource)
     addTextOptionsDataSource.DataSourceType =        "CloudFileSystem"      
     addTextOptionsDataSource.DataPath =        "BookText.xlsx"      
    var addTextOptions = new(AddTextOptions)
     addTextOptions.DataSource =        addTextOptionsDataSource      
     addTextOptions.Text =        "Aspose.Cells Cloud is an excellent product."      
     addTextOptions.Worksheet =        "202401"      
     addTextOptions.SelectPoistion =        "AtTheBeginning"      
     addTextOptions.SkipEmptyCells =  true      

    request := new (asposecellscloud.PostAddTextContentRequest)
    request.AddTextOptions =         addTextOptions    
    _, httpResponse, err := instance.PostAddTextContent(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
