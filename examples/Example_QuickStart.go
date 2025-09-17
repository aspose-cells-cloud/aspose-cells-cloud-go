package main

import (
	"fmt"
	"os"

	. "github.com/aspose-cells-cloud/aspose-cells-cloud-go/v25"
)

func main() {
	// If no environment variables are configured, please obtain the ClientId and ClientSecret from https://dashboard.aspose.cloud/#/applications and replace the following values:
	// instance  := NewCellsApiService('YourClientId','YourClientSecret')
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
