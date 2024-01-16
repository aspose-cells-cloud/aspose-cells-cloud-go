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
 
    var rangeOperateborderColor = new(Color)
     rangeOperateborderColor.R = int64(48)          
     rangeOperateborderColor.G = int64(48)          
     rangeOperateborderColor.B = int64(48)          
    var rangeOperateRange = new(Range)
     rangeOperateRange.ColumnCount = int64(1)          
     rangeOperateRange.ColumnWidth = 10.0      
     rangeOperateRange.FirstRow = int64(1)          
     rangeOperateRange.RowCount = int64(10)          
    var rangeOperate = new(RangeSetOutlineBorderRequest)
     rangeOperate.BorderEdge =        "LeftBorder"      
     rangeOperate.BorderStyle =        "Dotted"      
     rangeOperate.BorderColor =        rangeOperateborderColor      
     rangeOperate.Range_ =        rangeOperateRange      

    request := new (asposecellscloud.PostWorksheetCellsRangeOutlineBorderRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.RangeOperate =         rangeOperate    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := instance.PostWorksheetCellsRangeOutlineBorder(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
