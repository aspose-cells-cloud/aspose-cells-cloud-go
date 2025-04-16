package asposecellscloudtest

import (
	"fmt"
	"testing"

    . "asposecellscloud"
)

func TestWorkbookController_PostDigitalSignature(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    roywangPFX := "roywang.pfx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
    roywangPFXRequest := new(UploadFileRequest)
    roywangPFXRequest.UploadFiles = make(map[string]string) 
    roywangPFXRequest.UploadFiles[roywangPFX] =  GetBaseTest().localTestDataFolder  + roywangPFX
    roywangPFXRequest.Path = remoteFolder + "/roywang.pfx" 
    roywangPFXRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(roywangPFXRequest )
 
    request := new (PostDigitalSignatureRequest)
    request.Name =         remoteName    


    request.Digitalsignaturefile =         remoteFolder + "/roywang.pfx"    


    request.Password =         "123456"    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PostDigitalSignature(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorkbookController_PostDigitalSignature \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_PostEncryptWorkbook(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    var encryption = new(WorkbookEncryptionRequest)
     encryption.Password =        "123456"      
     encryption.EncryptionType =        "XOR"      
     encryption.KeyLength = int64(128)          
    request := new (PostEncryptWorkbookRequest)
    request.Name =         remoteName    


    request.Encryption =         encryption    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PostEncryptWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorkbookController_PostEncryptWorkbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_DeleteDecryptWorkbook(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    var encryption = new(WorkbookEncryptionRequest)
     encryption.Password =        "123456"      
     encryption.EncryptionType =        "XOR"      
     encryption.KeyLength = int64(128)          
    request := new (DeleteDecryptWorkbookRequest)
    request.Name =         remoteName    


    request.Encryption =         encryption    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.DeleteDecryptWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorkbookController_DeleteDecryptWorkbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_PostProtectWorkbook(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    var protectWorkbookRequest = new(ProtectWorkbookRequest)
     protectWorkbookRequest.EncryptWithPassword =        "123456"      
     protectWorkbookRequest.ProtectWorkbookStructure =        "ALL"      
    request := new (PostProtectWorkbookRequest)
    request.Name =         remoteName    


    request.ProtectWorkbookRequest =         protectWorkbookRequest    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PostProtectWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorkbookController_PostProtectWorkbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_DeleteUnProtectWorkbook(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (DeleteUnProtectWorkbookRequest)
    request.Name =         remoteName    


    request.Password =         remoteName    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.DeleteUnProtectWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorkbookController_DeleteUnProtectWorkbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_GetWorkbookDefaultStyle(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (GetWorkbookDefaultStyleRequest)
    request.Name =         remoteName    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.GetWorkbookDefaultStyle(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorkbookController_GetWorkbookDefaultStyle \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_GetWorkbookTextItems(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (GetWorkbookTextItemsRequest)
    request.Name =         remoteName    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.GetWorkbookTextItems(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorkbookController_GetWorkbookTextItems \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_GetWorkbookNames(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (GetWorkbookNamesRequest)
    request.Name =         remoteName    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.GetWorkbookNames(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorkbookController_GetWorkbookNames \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_PutWorkbookName(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    var newName = new(Name)
     newName.Text =        "name_1804"      
     newName.Comment =        "KeepSourceFormatting"      
     newName.RefersTo =        "=Sheet1!$I$4"      
    request := new (PutWorkbookNameRequest)
    request.Name =         remoteName    


    request.NewName =         newName    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PutWorkbookName(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorkbookController_PutWorkbookName \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_GetWorkbookName(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (GetWorkbookNameRequest)
    request.Name =         remoteName    


    request.NameName =         "Name_2"    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.GetWorkbookName(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorkbookController_GetWorkbookName \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_PostWorkbookName(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    var newName = new(Name)
     newName.Text =        "name_1804"      
     newName.Comment =        "KeepSourceFormatting"      
     newName.RefersTo =        "=Sheet1!$I$4"      
    request := new (PostWorkbookNameRequest)
    request.Name =         remoteName    


    request.NameName =         "Name_2"    


    request.NewName =         newName    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookName(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorkbookController_PostWorkbookName \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_GetWorkbookNameValue(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (GetWorkbookNameValueRequest)
    request.Name =         remoteName    


    request.NameName =         "Name_2"    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.GetWorkbookNameValue(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorkbookController_GetWorkbookNameValue \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_DeleteWorkbookNames(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (DeleteWorkbookNamesRequest)
    request.Name =         remoteName    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.DeleteWorkbookNames(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorkbookController_DeleteWorkbookNames \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_DeleteWorkbookName(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (DeleteWorkbookNameRequest)
    request.Name =         remoteName    


    request.NameName =         "Name_2"    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.DeleteWorkbookName(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorkbookController_DeleteWorkbookName \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_PutDocumentProtectFromChanges(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    var password = new(PasswordRequest)
     password.Password =        "123456"      
    request := new (PutDocumentProtectFromChangesRequest)
    request.Name =         remoteName    


    request.Password =         password    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PutDocumentProtectFromChanges(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorkbookController_PutDocumentProtectFromChanges \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_DeleteDocumentUnProtectFromChanges(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (DeleteDocumentUnProtectFromChangesRequest)
    request.Name =         remoteName    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.DeleteDocumentUnProtectFromChanges(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorkbookController_DeleteDocumentUnProtectFromChanges \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_PostWorkbooksMerge(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    myDocumentXLSX := "myDocument.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
    myDocumentXLSXRequest := new(UploadFileRequest)
    myDocumentXLSXRequest.UploadFiles = make(map[string]string) 
    myDocumentXLSXRequest.UploadFiles[myDocumentXLSX] =  GetBaseTest().localTestDataFolder  + myDocumentXLSX
    myDocumentXLSXRequest.Path = remoteFolder + "/myDocument.xlsx" 
    myDocumentXLSXRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(myDocumentXLSXRequest )
 
    request := new (PostWorkbooksMergeRequest)
    request.Name =         remoteName    


    request.MergeWith =         remoteFolder + "/myDocument.xlsx"    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    


    request.MergedStorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbooksMerge(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorkbookController_PostWorkbooksMerge \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_PostWorkbooksTextSearch(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (PostWorkbooksTextSearchRequest)
    request.Name =         remoteName    


    request.Text =         "1234"    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbooksTextSearch(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorkbookController_PostWorkbooksTextSearch \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_PostWorkbookTextReplace(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (PostWorkbookTextReplaceRequest)
    request.Name =         remoteName    


    request.OldValue =         "1234"    


    request.NewValue =         "5678"    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookTextReplace(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorkbookController_PostWorkbookTextReplace \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_PostWorkbookGetSmartMarkerResult(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    reportDataXML := "ReportData.xml"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
    reportDataXMLRequest := new(UploadFileRequest)
    reportDataXMLRequest.UploadFiles = make(map[string]string) 
    reportDataXMLRequest.UploadFiles[reportDataXML] =  GetBaseTest().localTestDataFolder  + reportDataXML
    reportDataXMLRequest.Path = remoteFolder + "/ReportData.xml" 
    reportDataXMLRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(reportDataXMLRequest )
 
    request := new (PostWorkbookGetSmartMarkerResultRequest)
    request.Name =         remoteName    


    request.XmlFile =         remoteFolder + "/ReportData.xml"    


    request.Folder =         remoteFolder    


    request.OutPath =         "OutResult/SmartMarkerResult.xlsx"    


    request.StorageName =         ""    


    request.OutStorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookGetSmartMarkerResult(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorkbookController_PostWorkbookGetSmartMarkerResult \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_PutWorkbookCreate(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    reportDataXML := "ReportData.xml"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
    reportDataXMLRequest := new(UploadFileRequest)
    reportDataXMLRequest.UploadFiles = make(map[string]string) 
    reportDataXMLRequest.UploadFiles[reportDataXML] =  GetBaseTest().localTestDataFolder  + reportDataXML
    reportDataXMLRequest.Path = remoteFolder + "/ReportData.xml" 
    reportDataXMLRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(reportDataXMLRequest )
 
    request := new (PutWorkbookCreateRequest)
    request.Name =         "PutWorkbookCreate.xlsx"    


    request.TemplateFile =         remoteFolder + "/" + remoteName    


    request.DataFile =         remoteFolder + "/ReportData.xml"    


    request.IsWriteOver =   true    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    


    request.CheckExcelRestriction =   true    

    _, httpResponse, err := GetBaseTest().CellsApi.PutWorkbookCreate(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorkbookController_PutWorkbookCreate \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_PostWorkbookSplit(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (PostWorkbookSplitRequest)
    request.Name =         remoteName    


    request.Format =         "png"    


    request.OutFolder =         "OutResult"    


    request.From =  int64(1)        


    request.To =  int64(5)        


    request.HorizontalResolution =  int64(96)        


    request.VerticalResolution =  int64(96)        


    request.SplitNameRule =         "sheetname"    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    


    request.OutStorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookSplit(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorkbookController_PostWorkbookSplit \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_PostImportData(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    var importOptionData = []int64   {int64(1)        ,
    int64(2)        ,
    int64(3)        ,
    int64(4)        }    
    var importOption = new(ImportIntArrayOption)
     importOption.DestinationWorksheet =        "Sheet1"      
     importOption.FirstColumn = int64(1)          
     importOption.FirstRow = int64(3)          
     importOption.ImportDataType =        "IntArray"      
     importOption.IsInsert =  true      
     importOption.IsVertical =  true      
     importOption.Data =        importOptionData      
    request := new (PostImportDataRequest)
    request.Name =         remoteName    


    request.ImportOption =         importOption    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PostImportData(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorkbookController_PostImportData \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_PostWorkbookCalculateFormula(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    var options = new(CalculationOptions)
     options.IgnoreError =  true      
     options.Recursive =  true      
    request := new (PostWorkbookCalculateFormulaRequest)
    request.Name =         remoteName    


    request.Options =         options    


    request.IgnoreError =   true    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookCalculateFormula(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorkbookController_PostWorkbookCalculateFormula \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_PostAutofitWorkbookRows(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (PostAutofitWorkbookRowsRequest)
    request.Name =         remoteName    


    request.StartRow =  int64(1)        


    request.EndRow =  int64(100)        


    request.OnlyAuto =   true    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PostAutofitWorkbookRows(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorkbookController_PostAutofitWorkbookRows \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_PostAutofitWorkbookColumns(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (PostAutofitWorkbookColumnsRequest)
    request.Name =         remoteName    


    request.StartColumn =  int64(1)        


    request.EndColumn =  int64(20)        


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PostAutofitWorkbookColumns(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorkbookController_PostAutofitWorkbookColumns \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_GetWorkbookSettings(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (GetWorkbookSettingsRequest)
    request.Name =         remoteName    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.GetWorkbookSettings(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorkbookController_GetWorkbookSettings \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_PostWorkbookSettings(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    var settings = new(WorkbookSettings)
     settings.AutoCompressPictures =  true      
     settings.HidePivotFieldList =  true      
    request := new (PostWorkbookSettingsRequest)
    request.Name =         remoteName    


    request.Settings =         settings    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookSettings(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorkbookController_PostWorkbookSettings \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_PutWorkbookBackground(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    waterMarkPNG := "WaterMark.png"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
    waterMarkPNGRequest := new(UploadFileRequest)
    waterMarkPNGRequest.UploadFiles = make(map[string]string) 
    waterMarkPNGRequest.UploadFiles[waterMarkPNG] =  GetBaseTest().localTestDataFolder  + waterMarkPNG
    waterMarkPNGRequest.Path = remoteFolder + "/WaterMark.png" 
    waterMarkPNGRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(waterMarkPNGRequest )
 
    request := new (PutWorkbookBackgroundRequest)
    request.Name =         remoteName    


    request.PicPath =         remoteFolder + "/WaterMark.png"    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PutWorkbookBackground(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorkbookController_PutWorkbookBackground \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_DeleteWorkbookBackground(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (DeleteWorkbookBackgroundRequest)
    request.Name =         remoteName    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.DeleteWorkbookBackground(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorkbookController_DeleteWorkbookBackground \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_PutWorkbookWaterMarker(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    var textWaterMarkerRequest = new(TextWaterMarkerRequest)
     textWaterMarkerRequest.Text =        "Aspose Cells Cloud"      
     textWaterMarkerRequest.FontSize = int64(12)          
    request := new (PutWorkbookWaterMarkerRequest)
    request.Name =         remoteName    


    request.TextWaterMarkerRequest =         textWaterMarkerRequest    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.PutWorkbookWaterMarker(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorkbookController_PutWorkbookWaterMarker \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_GetPageCount(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
    request := new (GetPageCountRequest)
    request.Name =         remoteName    


    request.Folder =         remoteFolder    


    request.StorageName =         ""    

    _, httpResponse, err := GetBaseTest().CellsApi.GetPageCount(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorkbookController_GetPageCount \n", GetBaseTest().GetTestNumber())
	}
}

