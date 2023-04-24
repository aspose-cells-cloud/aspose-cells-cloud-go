package asposecellscloud

import (
	"fmt"
	"testing"
)

func TestCellsBarcodesController_GetExtractBarcodes(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (GetExtractBarcodesRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet8"    
    request.PictureIndex =  int64(0)        
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.GetExtractBarcodes(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsBarcodesController_GetExtractBarcodes \n", GetBaseTest().GetTestNumber())
	}
}

