package main

import (
	"os"

	asposecellscloud "github.com/aspose-cells-cloud/aspose-cells-cloud-go"
)
func main() {
	instance := asposecellscloud.NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"), "https://api.aspose.cloud", "v3.0")
    remoteFolder := "TestData/In"

    localName := "BookCsvDuplicateData.csv"
    remoteName := "BookCsvDuplicateData.csv"

    localNameRequest := new(asposecellscloud.UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =   localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    instance.UploadFile(localNameRequest )
 
    var dataCleansingDataFillDataFillDefaultValue = new(DataFillValue)
     dataCleansingDataFillDataFillDefaultValue.DefaultDate =        "2024-01-01"      
     dataCleansingDataFillDataFillDefaultValue.DefaultNumber = int64(0)          
     dataCleansingDataFillDataFillDefaultValue.DefaultBoolean =  false      
    var dataCleansingDataFill = new(DataFill)
     dataCleansingDataFill.DataFillDefaultValue =        dataCleansingDataFillDataFillDefaultValue      
    var dataCleansing = new(DataCleansing)
     dataCleansing.NeedFillData =  true      
     dataCleansing.DataFill =        dataCleansingDataFill      

    request := new (asposecellscloud.PostWorkbookDataCleansingRequest)
    request.Name =         remoteName    
    request.DataCleansing =         dataCleansing    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := instance.PostWorkbookDataCleansing(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
