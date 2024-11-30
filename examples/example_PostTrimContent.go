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
 
    var trimContentOptionsDataSource = new(DataSource)
     trimContentOptionsDataSource.DataSourceType =        "CloudFileSystem"      
     trimContentOptionsDataSource.DataPath =        "BookText.xlsx"      
    var trimContentOptionsScopeOptions = new(ScopeOptions)
     trimContentOptionsScopeOptions.Scope =        "EntireWorkbook"      
    var trimContentOptions = new(TrimContentOptions)
     trimContentOptions.DataSource =        trimContentOptionsDataSource      
     trimContentOptions.TrimLeading =  true      
     trimContentOptions.TrimTrailing =  true      
     trimContentOptions.TrimSpaceBetweenWordTo1 =  true      
     trimContentOptions.RemoveAllLineBreaks =  true      
     trimContentOptions.ScopeOptions =        trimContentOptionsScopeOptions      

    request := new (asposecellscloud.PostTrimContentRequest)
    request.TrimContentOptions =         trimContentOptions    
    _, httpResponse, err := instance.PostTrimContent(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
