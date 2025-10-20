package main

import (
	"archive/zip"
	"bytes"
	"io"
	"os"

	. "github.com/aspose-cells-cloud/aspose-cells-cloud-go/v25"
)

func main() {
	EmployeeSalesSummaryXlsx := "EmployeeSalesSummary.xlsx"
	// If no environment variables are configured, please obtain the ClientId and ClientSecret from https://dashboard.aspose.cloud/#/applications and replace the following values:
	// instance  := NewCellsApiService('YourClientId','YourClientSecret')
	instance := NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"))

	//Splitting an Excel or other spreadsheet file to another format files online.
	zipData, _, err := instance.SplitSpreadsheet(&SplitSpreadsheetRequest{Spreadsheet: EmployeeSalesSummaryXlsx, From: 0, To: 2}, CellsCloudOption{OptionName: "LocalOutPath", OptionValue: "splittedFiles.zip"})
	if err != nil {
		print("Error:", err.Error())
	} else {
		// err = os.WriteFile("splittedFiles.zip", zipData, 0644)
		// if err != nil {
		// 	print("Error:", err.Error())
		// 	return
		// }
		println("Successful split Excel files.")
	}

	reader := bytes.NewReader(zipData)
	zipReader, err := zip.NewReader(reader, int64(len(zipData)))
	if err != nil {
		return
	}

	for _, file := range zipReader.File {
		rc, err := file.Open()
		if err != nil {
			print("Error:", err.Error())
			return
		}

		content, err := io.ReadAll(rc)
		rc.Close()
		if err != nil {
			print("Error:", err.Error())
			return
		}
		err = os.WriteFile(file.Name, content, file.Mode())
		if err != nil {
			print("Error:", err.Error())
			return
		}
	}

}
