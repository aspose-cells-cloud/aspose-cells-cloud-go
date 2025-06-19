package main

import (
	"fmt"
	"os"

	. "github.com/aspose-cells-cloud/aspose-cells-cloud-go/v25"
)

func main() {
	EmployeeSalesSummaryXlsx := "EmployeeSalesSummary.xlsx"
	CompanySalesXlsx := "CompanySales.xlsx"
	RemoteFolder := "GoSDK"
	RemoteOutputFolder := "GoSDK/Output"
	instance := NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"))
	_, httpResponse, err := instance.UploadFile(&UploadFileRequest{UploadFiles: CompanySalesXlsx, Path: RemoteFolder + "/" + CompanySalesXlsx})
	_, httpResponse, err = instance.UploadFile(&UploadFileRequest{UploadFiles: EmployeeSalesSummaryXlsx, Path: RemoteFolder + "/" + EmployeeSalesSummaryXlsx})
	// Batch convert Excel files.
	_, httpResponse, err = instance.PostBatchConvert(&PostBatchConvertRequest{BatchConvertRequest: &BatchConvertRequest{SourceFolder: RemoteFolder, Format: "Pdf", OutFolder: RemoteOutputFolder, MatchCondition: &MatchConditionRequest{RegexPattern: "xlsx"}}})
	if err != nil {
		fmt.Print(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Print(httpResponse.StatusCode)
	} else {
		println("Successful batch conversion of Excel files.")
	}
	// Batch lock Excel files.
	_, httpResponse, err = instance.PostBatchLock(&PostBatchLockRequest{BatchLockRequest: &BatchLockRequest{SourceFolder: RemoteFolder, Password: "123456", OutFolder: RemoteOutputFolder, MatchCondition: &MatchConditionRequest{RegexPattern: "xlsx"}}})
	if err != nil {
		fmt.Print(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Print(httpResponse.StatusCode)
	} else {
		println("Successful batch lock of Excel files.")
	}
	// Batch unlock Excel files.
	_, httpResponse, err = instance.PostBatchUnlock(&PostBatchUnlockRequest{BatchLockRequest: &BatchLockRequest{SourceFolder: RemoteFolder, Password: "123456", OutFolder: RemoteOutputFolder, MatchCondition: &MatchConditionRequest{RegexPattern: "xlsx"}}})
	if err != nil {
		fmt.Print(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Print(httpResponse.StatusCode)
	} else {
		println("Successful batch unlock of Excel files.")
	}

	// Batch split Excel files.
	_, httpResponse, err = instance.PostBatchSplit(&PostBatchSplitRequest{BatchSplitRequest: &BatchSplitRequest{SourceFolder: RemoteFolder, Format: "pdf", OutFolder: RemoteOutputFolder, MatchCondition: &MatchConditionRequest{RegexPattern: "xlsx"}}})
	if err != nil {
		fmt.Print(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Print(httpResponse.StatusCode)
	} else {
		println("Successful batch split of Excel files.")
	}

}
