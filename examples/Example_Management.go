package main

import (
	"os"

	. "github.com/aspose-cells-cloud/aspose-cells-cloud-go/v26"
)

func main() {
	// If no environment variables are configured, please obtain the ClientId and ClientSecret from https://dashboard.aspose.cloud/#/applications and replace the following values:
	// instance  := NewCellsApiService('YourClientId','YourClientSecret')
	instance := NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"))
	// Local file processing with cells cloud services.
	// Add new worksheet to an Excel file.
	instance.AddWorksheetToSpreadsheet(&AddWorksheetToSpreadsheetRequest{Spreadsheet: "EmployeeSalesSummary.xlsx", SheetName: "Sheet1", Position: 0}, CellsCloudOption{OptionName: "LocalOutPath", OptionValue: "EmployeeSalesSummary_Output.xlsx"})
	// Rename worksheet in an Excel file.
	instance.RenameWorksheetInSpreadsheet(&RenameWorksheetInSpreadsheetRequest{Spreadsheet: "EmployeeSalesSummary_Output.xlsx", SourceName: "Sheet1", TargetName: "NewSheet"}, CellsCloudOption{OptionName: "LocalOutPath", OptionValue: "EmployeeSalesSummary_Output.xlsx"})
	// Move worksheet to new position in an Excel file.
	instance.MoveWorksheetInSpreadsheet(&MoveWorksheetInSpreadsheetRequest{Spreadsheet: "EmployeeSalesSummary_Output.xlsx", Worksheet: "NewSheet", Position: 2}, CellsCloudOption{OptionName: "LocalOutPath", OptionValue: "EmployeeSalesSummary_Output.xlsx"})
	// Delete worksheet in an Excel file.
	instance.DeleteWorksheetFromSpreadsheet(&DeleteWorksheetFromSpreadsheetRequest{Spreadsheet: "EmployeeSalesSummary_Output.xlsx", SheetName: "NewSheet"}, CellsCloudOption{OptionName: "LocalOutPath", OptionValue: "EmployeeSalesSummary_Output.xlsx"})
	// Compress an Excel file.
	instance.CompressSpreadsheet(&CompressSpreadsheetRequest{Spreadsheet: "EmployeeSalesSummary_Output.xlsx", Level: 9}, CellsCloudOption{OptionName: "LocalOutPath", OptionValue: "EmployeeSalesSummary_Output.xlsx"})
}
