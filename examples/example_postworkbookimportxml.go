package main

import (
	"os"

	asposecellscloud "github.com/aspose-cells-cloud/aspose-cells-cloud-go"
)
func main() {
	instance := asposecellscloud.NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"), "https://api.aspose.cloud", "v3.0")
    remoteFolder := "TestData/In"

    localName := "Template.xlsx"
    dataXML := "data.xml"
    remoteName := "Template.xlsx"

    localNameRequest := new(asposecellscloud.UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =   localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    instance.UploadFile(localNameRequest )
    dataXMLRequest := new(asposecellscloud.UploadFileRequest)
    dataXMLRequest.UploadFiles = make(map[string]string) 
    dataXMLRequest.UploadFiles[dataXML] =   dataXML
    dataXMLRequest.Path = remoteFolder + "/data.xml" 
    dataXMLRequest.StorageName =""
    instance.UploadFile(dataXMLRequest )
 
    var importXMLRequestXMLFileSource = new(DataSource)
     importXMLRequestXMLFileSource.DataSourceType =        "CloudFileSystem"      
     importXMLRequestXMLFileSource.DataPath =        remoteFolder + "/data.xml"      
    var importXMLRequestImportPosition = new(ImportPosition)
     importXMLRequestImportPosition.SheetName =        "Sheet1"      
     importXMLRequestImportPosition.RowIndex = int64(3)          
     importXMLRequestImportPosition.ColumnIndex = int64(4)          
    var importXMLRequest = new(ImportXMLRequest)
     importXMLRequest.XMLFileSource =        importXMLRequestXMLFileSource      
     importXMLRequest.ImportPosition =        importXMLRequestImportPosition      

    request := new (asposecellscloud.PostWorkbookImportXMLRequest)
    request.Name =         remoteName    
    request.ImportXMLRequest =         importXMLRequest    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := instance.PostWorkbookImportXML(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
