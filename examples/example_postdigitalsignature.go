package main

import (
	"os"

	asposecellscloud "github.com/aspose-cells-cloud/aspose-cells-cloud-go"
)
func main() {
	instance := asposecellscloud.NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"), "https://api.aspose.cloud", "v3.0")
    remoteFolder := "TestData/In"

    localName := "Book1.xlsx"
    roywangPFX := "roywang.pfx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(asposecellscloud.UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =   localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    instance.UploadFile(localNameRequest )
    roywangPFXRequest := new(asposecellscloud.UploadFileRequest)
    roywangPFXRequest.UploadFiles = make(map[string]string) 
    roywangPFXRequest.UploadFiles[roywangPFX] =   roywangPFX
    roywangPFXRequest.Path = remoteFolder + "/roywang.pfx" 
    roywangPFXRequest.StorageName =""
    instance.UploadFile(roywangPFXRequest )
 

    request := new (asposecellscloud.PostDigitalSignatureRequest)
    request.Name =         remoteName    
    request.Digitalsignaturefile =         remoteFolder + "/roywang.pfx"    
    request.Password =         "123456"    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := instance.PostDigitalSignature(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
