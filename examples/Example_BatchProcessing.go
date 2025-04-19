package main

import (
	"fmt"
	. "github.com/aspose-cells-cloud/aspose-cells-cloud-go/v25"
	"os"
)

func main() {
	EmployeeSalesSummaryXlsx := "EmployeeSalesSummary.xlsx"
	CompanySalesXlsx := "CompanySales.xlsx"
	RemoteFolder := "GoSDK"
	RemoteOutputFolder := "GoSDKOut"
	instance := NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"))
	_, httpResponse, err := instance.UploadFile(&UploadFileRequest{UploadFile: CompanySalesXlsx, Path: RemoteFolder + "/" + CompanySalesXlsx})
	_, httpResponse, err = instance.UploadFile(&UploadFileRequest{UploadFile: EmployeeSalesSummaryXlsx, Path: RemoteFolder + "/" + EmployeeSalesSummaryXlsx})
	_, httpResponse, err = instance.PostBatchConvert(&PostBatchConvertRequest{BatchConvertRequest: &BatchConvertRequest{SourceFolder: RemoteFolder, Format: "Pdf", OutFolder: RemoteOutputFolder, MatchCondition: &MatchConditionRequest{RegexPattern: "xlsx"}}})
	if err != nil {
		fmt.Print(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Print(httpResponse.StatusCode)
	} else {
	}
}
