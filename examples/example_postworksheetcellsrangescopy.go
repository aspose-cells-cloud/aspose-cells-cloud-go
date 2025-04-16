package main

import (
	"os"

	asposecellscloud "github.com/aspose-cells-cloud/aspose-cells-cloud-go"
)
func main() {
	instance := asposecellscloud.NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"), "https://api.aspose.cloud", "v3.0")
    remoteFolder := "TestData/In"

    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(asposecellscloud.UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =   localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    instance.UploadFile(localNameRequest )
 
    var rangeOperateSource = new(Range)
     rangeOperateSource.ColumnCount = int64(3)          
     rangeOperateSource.FirstColumn = int64(8)          
     rangeOperateSource.FirstRow = int64(3)          
     rangeOperateSource.RowCount = int64(2)          
    var rangeOperateTarget = new(Range)
     rangeOperateTarget.ColumnCount = int64(3)          
     rangeOperateTarget.FirstColumn = int64(8)          
     rangeOperateTarget.FirstRow = int64(13)          
     rangeOperateTarget.RowCount = int64(2)          
    var rangeOperate = new(RangeCopyRequest)
     rangeOperate.Operate =        "copydata"      
     rangeOperate.Source =        rangeOperateSource      
     rangeOperate.Target =        rangeOperateTarget      

    request := new (asposecellscloud.PostWorksheetCellsRangesCopyRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.RangeOperate =         rangeOperate    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := instance.PostWorksheetCellsRangesCopy(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
