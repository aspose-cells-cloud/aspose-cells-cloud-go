package main

import (
	"encoding/base64"
	"fmt"
	. "github.com/aspose-cells-cloud/aspose-cells-cloud-go/v25"
	"os"
)

func main() {
	EmployeeSalesSummaryXlsx := "EmployeeSalesSummary.xlsx"
	CompanySalesXlsx := "CompanySales.xlsx"
	RemoteFolder := "GoSDK"

	instance := NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"))

	//Splitting an Excel or other spreadsheet file to another format files online.
	_, httpResponse, err := instance.UploadFile(&UploadFileRequest{UploadFile: EmployeeSalesSummaryXlsx, Path: RemoteFolder + "/" + EmployeeSalesSummaryXlsx})
	httpResponse, err = instance.CreateFolder(&CreateFolderRequest{Path: RemoteFolder + "/Output"})
	_, httpResponse, err = instance.PostWorkbookSplit(&PostWorkbookSplitRequest{Name: EmployeeSalesSummaryXlsx, Folder: RemoteFolder, OutFolder: RemoteFolder + "/Output", Format: "pdf", To: 3})
	if err != nil {
		fmt.Print(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Print(httpResponse.StatusCode)
	} else {
		println("Successful split Excel files.")
	}
	// Split two local Excel files to another format files.
	filesResult, mergedHttpResponse, mergedError := instance.PostSplit(&PostSplitRequest{File: map[string]string{CompanySalesXlsx: CompanySalesXlsx, EmployeeSalesSummaryXlsx: EmployeeSalesSummaryXlsx}, OutFormat: "png", To: 3})
	if mergedError != nil {
		fmt.Print(mergedError)
	} else if mergedHttpResponse.StatusCode < 200 || mergedHttpResponse.StatusCode > 299 {
		fmt.Print(mergedHttpResponse.StatusCode)
	} else {
		println("Successful split Excel files. Begin to write every files.")
		for _, fileInfo := range filesResult.Files {
			print(fileInfo.Filename)
			decodedData, _ := base64.StdEncoding.DecodeString(string(fileInfo.FileContent))
			file, _ := os.OpenFile(fileInfo.Filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
			file.Write(decodedData)
			defer file.Close()
		}
	}
	// Split a local Excel files to another format files.
	filesResult, mergedHttpResponse, mergedError = instance.PostSplit(&PostSplitRequest{LocalPath: CompanySalesXlsx, OutFormat: "pdf", To: 3})
	if mergedError != nil {
		fmt.Print(mergedError)
	} else if mergedHttpResponse.StatusCode < 200 || mergedHttpResponse.StatusCode > 299 {
		fmt.Print(mergedHttpResponse.StatusCode)
	} else {
		println("Successful split an Excel files. Begin to write every files.")
		for _, fileInfo := range filesResult.Files {
			decodedData, _ := base64.StdEncoding.DecodeString(string(fileInfo.FileContent))
			file, _ := os.OpenFile(fileInfo.Filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
			file.Write(decodedData)
			defer file.Close()
		}
	}
}
