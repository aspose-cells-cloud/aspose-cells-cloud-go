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
 
    var rangeOperateStyleFont = new(Font)
     rangeOperateStyleFont.Size = int64(16)          
    var rangeOperateStyle = new(Style)
     rangeOperateStyle.Font =        rangeOperateStyleFont      
    var rangeOperateRange = new(Range)
     rangeOperateRange.ColumnCount = int64(1)          
     rangeOperateRange.ColumnWidth = 10.0      
     rangeOperateRange.FirstRow = int64(1)          
     rangeOperateRange.RowCount = int64(10)          
    var rangeOperate = new(RangeSetStyleRequest)
     rangeOperate.Style =        rangeOperateStyle      
     rangeOperate.Range_ =        rangeOperateRange      

    request := new (asposecellscloud.PostWorksheetCellsRangeStyleRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.RangeOperate =         rangeOperate    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := instance.PostWorksheetCellsRangeStyle(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
