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
 
    var wordCaseOptionsDataSource = new(DataSource)
     wordCaseOptionsDataSource.DataSourceType =        "CloudFileSystem"      
     wordCaseOptionsDataSource.DataPath =        "BookText.xlsx"      
    var wordCaseOptionsScopeOptions = new(ScopeOptions)
     wordCaseOptionsScopeOptions.Scope =        "EntireWorkbook"      
    var wordCaseOptions = new(WordCaseOptions)
     wordCaseOptions.DataSource =        wordCaseOptionsDataSource      
     wordCaseOptions.WordCaseType =        "None"      
     wordCaseOptions.ScopeOptions =        wordCaseOptionsScopeOptions      

    request := new (asposecellscloud.PostUpdateWordCaseRequest)
    request.WordCaseOptions =         wordCaseOptions    
    _, httpResponse, err := instance.PostUpdateWordCase(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
