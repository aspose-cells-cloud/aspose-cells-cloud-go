package main

import (
	"os"

	asposecellscloud "github.com/aspose-cells-cloud/aspose-cells-cloud-go"
)
func main() {
	instance := asposecellscloud.NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"), "https://api.aspose.cloud", "v3.0")
    remoteFolder := "TestData/In"

    localName := "Book1.xlsx"
    reportDataXML := "ReportData.xml"
    remoteName := "Book1.xlsx"

    localNameRequest := new(asposecellscloud.UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =   localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    instance.UploadFile(localNameRequest )
    reportDataXMLRequest := new(asposecellscloud.UploadFileRequest)
    reportDataXMLRequest.UploadFiles = make(map[string]string) 
    reportDataXMLRequest.UploadFiles[reportDataXML] =   reportDataXML
    reportDataXMLRequest.Path = remoteFolder + "/ReportData.xml" 
    reportDataXMLRequest.StorageName =""
    instance.UploadFile(reportDataXMLRequest )
 

    request := new (asposecellscloud.PostWorkbookGetSmartMarkerResultRequest)
    request.Name =         remoteName    
    request.XmlFile =         remoteFolder + "/ReportData.xml"    
    request.Folder =         remoteFolder    
    request.OutPath =         "OutResult/SmartMarkerResult.xlsx"    
    request.StorageName =         ""    
    request.OutStorageName =         ""    
    _, httpResponse, err := instance.PostWorkbookGetSmartMarkerResult(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
