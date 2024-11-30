package main

import (
	"os"

	asposecellscloud "github.com/aspose-cells-cloud/aspose-cells-cloud-go"
)
func main() {
	instance := asposecellscloud.NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"), "https://api.aspose.cloud", "v3.0")
    remoteFolder := "TestData/In"

    localName := "Book1.xlsx"
    waterMarkPNG := "WaterMark.png"
    remoteName := "Book1.xlsx"

    localNameRequest := new(asposecellscloud.UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =   localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    instance.UploadFile(localNameRequest )
    waterMarkPNGRequest := new(asposecellscloud.UploadFileRequest)
    waterMarkPNGRequest.UploadFiles = make(map[string]string) 
    waterMarkPNGRequest.UploadFiles[waterMarkPNG] =   waterMarkPNG
    waterMarkPNGRequest.Path = remoteFolder + "/WaterMark.png" 
    waterMarkPNGRequest.StorageName =""
    instance.UploadFile(waterMarkPNGRequest )
 

    request := new (asposecellscloud.PutWorkbookBackgroundRequest)
    request.Name =         remoteName    
    request.PicPath =         remoteFolder + "/WaterMark.png"    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := instance.PutWorkbookBackground(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
