package main

import (
	"os"

	. "github.com/aspose-cells-cloud/aspose-cells-cloud-go/v25"
)

func main() {
	// If no environment variables are configured, please obtain the ClientId and ClientSecret from https://dashboard.aspose.cloud/#/applications and replace the following values:
	// instance  := NewCellsApiService('YourClientId','YourClientSecret')
	instance := NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"))
	// Local file processing with cells cloud services.
	// Protect an Excel file.
	instance.ProtectSpreadsheet(&ProtectSpreadsheetRequest{Spreadsheet: "EmployeeSalesSummary.xlsx", Password: "123456"}, CellsCloudOption{OptionName: "LocalOutPath", OptionValue: "EmployeeSalesSummary_Lock.xlsx"})
	// Unprotect an Excel file.
	instance.UnprotectSpreadsheet(&UnprotectSpreadsheetRequest{Spreadsheet: "EmployeeSalesSummary_Lock.xlsx", Password: "123456"}, CellsCloudOption{OptionName: "LocalOutPath", OptionValue: "EmployeeSalesSummary_unlock.xlsx"})

}
