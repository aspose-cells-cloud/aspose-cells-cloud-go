package asposecellscloudtest

import (
	"fmt"
	"testing"

    . "asposecellscloud"
)

func TestCellsStatusController_GetCellsCloudServicesHealthCheck(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (GetCellsCloudServicesHealthCheckRequest)
    _, httpResponse, err := GetBaseTest().CellsApi.GetCellsCloudServicesHealthCheck(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsStatusController_GetCellsCloudServicesHealthCheck \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsStatusController_GetCellsCloudServiceStatus(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (GetCellsCloudServiceStatusRequest)
    _, httpResponse, err := GetBaseTest().CellsApi.GetCellsCloudServiceStatus(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsStatusController_GetCellsCloudServiceStatus \n", GetBaseTest().GetTestNumber())
	}
}

