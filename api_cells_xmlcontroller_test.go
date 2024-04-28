package asposecellscloud

import (
	"fmt"
	"testing"
)

func TestXmlController_PostWorkbookExportXML(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Template.xlsx"
    remoteName := "Template.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (PostWorkbookExportXMLRequest)
    request.Name =         remoteName    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookExportXML(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestXmlController_PostWorkbookExportXML \n", GetBaseTest().GetTestNumber())
	}
}

func TestXmlController_PostWorkbookImportXML(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Template.xlsx"
    dataXML := "data.xml"
    remoteName := "Template.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
    dataXMLRequest := new(UploadFileRequest)
    dataXMLRequest.UploadFiles = make(map[string]string) 
    dataXMLRequest.UploadFiles[dataXML] =  GetBaseTest().localTestDataFolder  + dataXML
    dataXMLRequest.Path = remoteFolder + "/data.xml" 
    dataXMLRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(dataXMLRequest )
 
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

    request := new (PostWorkbookImportXMLRequest)
    request.Name =         remoteName    
    request.ImportXMLRequest =         importXMLRequest    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookImportXML(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestXmlController_PostWorkbookImportXML \n", GetBaseTest().GetTestNumber())
	}
}

