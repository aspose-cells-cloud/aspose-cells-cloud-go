package main

import (
	"encoding/base64"
	"fmt"
	"os"

	. "github.com/aspose-cells-cloud/aspose-cells-cloud-go/v26"
)

func main() {
	EmployeeSalesSummaryXlsx := "EmployeeSalesSummary.xlsx"
	CompanySalesXlsx := "CompanySales.xlsx"
	RemoteFolder := "GoSDK"
	// If no environment variables are configured, please obtain the ClientId and ClientSecret from https://dashboard.aspose.cloud/#/applications and replace the following values:
	// instance  := NewCellsApiService('YourClientId','YourClientSecret')
	instance := NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"))
	//Merge an Excel or other spreadsheet file into an existing Excel file online.
	_, httpResponse, err := instance.UploadFile(&UploadFileRequest{UploadFiles: CompanySalesXlsx, Path: RemoteFolder + "/" + CompanySalesXlsx})
	_, httpResponse, err = instance.UploadFile(&UploadFileRequest{UploadFiles: EmployeeSalesSummaryXlsx, Path: RemoteFolder + "/" + EmployeeSalesSummaryXlsx})
	_, httpResponse, err = instance.PostWorkbooksMerge(&PostWorkbooksMergeRequest{Name: CompanySalesXlsx, Folder: RemoteFolder, MergeWith: RemoteFolder + "/" + EmployeeSalesSummaryXlsx})
	if err != nil {
		fmt.Print(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Print(httpResponse.StatusCode)
	} else {
		println("Successful merge Excel files.")
	}
	// Merged multiple local Excel files, each containing unique data or sheets, into a single consolidated Excel file.
	fileInfo, mergedHttpResponse, mergedError := instance.PostMerge(&PostMergeRequest{File: map[string]string{CompanySalesXlsx: CompanySalesXlsx, EmployeeSalesSummaryXlsx: EmployeeSalesSummaryXlsx}, MergeToOneSheet: true})
	if mergedError != nil {
		fmt.Print(mergedError)
	} else if mergedHttpResponse.StatusCode < 200 || mergedHttpResponse.StatusCode > 299 {
		fmt.Print(mergedHttpResponse.StatusCode)
	} else {
		decodedData, _ := base64.StdEncoding.DecodeString(string(fileInfo.FileContent))
		file, _ := os.OpenFile("MergedExcel.xlsx", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		file.Write(decodedData)
		defer file.Close()
	}
}
