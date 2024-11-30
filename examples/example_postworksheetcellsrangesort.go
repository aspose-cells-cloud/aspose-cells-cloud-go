package main

import (
	"os"

	asposecellscloud "github.com/aspose-cells-cloud/aspose-cells-cloud-go"
)
func main() {
	instance := asposecellscloud.NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"), "https://api.aspose.cloud", "v3.0")
    remoteFolder := "TestData/In"

    localName := "Group.xlsx"
    remoteName := "Group.xlsx"

    localNameRequest := new(asposecellscloud.UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =   localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    instance.UploadFile(localNameRequest )
 
    var rangeSortRequestDataSorter = new(DataSorter)
     rangeSortRequestDataSorter.CaseSensitive =  true      
    var rangeSortRequestCellArea = new(Range)
     rangeSortRequestCellArea.ColumnCount = int64(3)          
     rangeSortRequestCellArea.FirstColumn = int64(0)          
     rangeSortRequestCellArea.FirstRow = int64(0)          
     rangeSortRequestCellArea.RowCount = int64(15)          
    var rangeSortRequest = new(RangeSortRequest)
     rangeSortRequest.DataSorter =        rangeSortRequestDataSorter      
     rangeSortRequest.CellArea =        rangeSortRequestCellArea      

    request := new (asposecellscloud.PostWorksheetCellsRangeSortRequest)
    request.Name =         remoteName    
    request.SheetName =         "book1"    
    request.RangeSortRequest =         rangeSortRequest    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := instance.PostWorksheetCellsRangeSort(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
