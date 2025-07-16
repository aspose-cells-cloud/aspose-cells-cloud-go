package main

import (
	"fmt"
	"os"

	. "github.com/aspose-cells-cloud/aspose-cells-cloud-go/v25"
)

func main() {
	instance := NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"))
	_, httpResponse, err := instance.ConvertSpreadsheet(&ConvertSpreadsheetRequest{Spreadsheet: "EmployeeSalesSummary.xlsx", Format: "pdf"}, []CellsCloudOption{{OptionName: "LocalOutPath", OptionValue: "EmployeeSalesSummary.pdf"}}...)
	if err != nil {
		fmt.Print(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Print("E")
	} else {
		fmt.Print("T")
	}
	Version()
}
