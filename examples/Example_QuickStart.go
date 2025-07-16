package main

import (
	"fmt"
	"os"

	. "github.com/aspose-cells-cloud/aspose-cells-cloud-go/v25"
)

func main() {
	instance := NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"))
	data, httpResponse, err := instance.ConvertSpreadsheet(&ConvertSpreadsheetRequest{Spreadsheet: "EmployeeSalesSummary.xlsx", Format: "pdf"}, &CellsCloudOptio{OptionName: "LocalOutPath", OptionValue: "EmployeeSalesSummary.pdf"})
	if err != nil {
		fmt.Print(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Print("E")
	} else {
		file, _ := os.OpenFile("EmployeeSalesSummary.pdf", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		file.Write(data)
		defer file.Close()
	}
	Version()
}
