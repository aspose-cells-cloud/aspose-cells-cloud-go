package main

import (
	"encoding/base64"
	"fmt"
	"os"

	. "github.com/aspose-cells-cloud/aspose-cells-cloud-go/v25"
)

func main() {
	employeeSalesSummaryXlsx := "EmployeeSalesSummary.xlsx"
	remoteFolder := "GoSDK"
	// If no environment variables are configured, please obtain the ClientId and ClientSecret from https://dashboard.aspose.cloud/#/applications and replace the following values:
	// instance  := NewCellsApiService('YourClientId','YourClientSecret')
	instance := NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"))

	// Local file processing
	// Convert a local Excel file to another format file directly, save to local file.
	// Complete in one step.
	_, httpResponse, err := instance.ConvertSpreadsheet(&ConvertSpreadsheetRequest{Spreadsheet: "EmployeeSalesSummary.xlsx", Format: "pdf"}, []CellsCloudOption{{OptionName: "LocalOutPath", OptionValue: "EmployeeSalesSummary.pdf"}}...)
	// Completed in two steps
	// Convert a local Excel file to byte data.
	convertedData, httpResponse, err := instance.ConvertSpreadsheet(&ConvertSpreadsheetRequest{Spreadsheet: employeeSalesSummaryXlsx, Format: "pdf"})
	if err != nil {
		fmt.Print(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Print(httpResponse.StatusCode)
	} else {
		// 2 Save result to specified path.
		file, _ := os.OpenFile("EmployeeSalesSummary1.pdf", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		file.Write(convertedData)
		defer file.Close()
	}

	// Cloud file processing
	// Convert Cloud file to PDF
	// 1. Upload a local Excel file to Cells Cloud Storage.
	_, httpResponse, err = instance.UploadFile(&UploadFileRequest{UploadFiles: employeeSalesSummaryXlsx, Path: remoteFolder + "/" + employeeSalesSummaryXlsx})
	// 2. Convert an Excel file of Cells Cloud to another format file.
	// 2.1 Save an Excel file of Cells Cloud as another format file of Cells Cloud.
	_, httpResponse, err = instance.SaveSpreadsheetAs(&SaveSpreadsheetAsRequest{Name: employeeSalesSummaryXlsx, Format: "pdf", Folder: remoteFolder})
	if err != nil {
		println("Save as")
		fmt.Print(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Print(httpResponse.StatusCode)
	}

	// 2.2 Export a worksheet of a cloud Excel file to another format file directly. Set query parameters : print_headings, one_page_per_sheet
	convertedData, httpResponse, err = instance.ExportWorksheetAsFormat(&ExportWorksheetAsFormatRequest{Name: employeeSalesSummaryXlsx, Worksheet: "Sales", Folder: remoteFolder, Format: "pdf"}, []CellsCloudOption{{OptionName: "LocalOutPath", OptionValue: "EmployeeSalesSummary_Sales.pdf"}}...)
	if err != nil {
		fmt.Print(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Print(httpResponse.StatusCode)
	} else {
		file, _ := os.OpenFile("EmployeeSalesSummary_Sale.png", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		file.Write(convertedData)
		defer file.Close()
	}

	// 2.3 Export a cloud Excel file's specified worksheet page index directly to another format file. Set query parameters : print_headings, one_page_per_sheet
	convertedData, _, _ = instance.ExportWorksheetAsFormat(&ExportWorksheetAsFormatRequest{Name: employeeSalesSummaryXlsx, Worksheet: "Sales", Folder: remoteFolder, Format: "png"})
	file, _ := os.OpenFile("EmployeeSalesSummary_Sale_1.png", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	file.Write(convertedData)
	defer file.Close()
	// 2.4 Export a cloud Excel file's specified worksheet cells area directly to another format file. Set query parameters : print_headings, one_page_per_sheet
	convertedData, _, _ = instance.ExportRangeAsFormat(&ExportRangeAsFormatRequest{Name: employeeSalesSummaryXlsx, Worksheet: "Sales", Folder: remoteFolder, Format: "png", Range_: "B5:L36"})
	file, _ = os.OpenFile("EmployeeSalesSummary_Sale_Area.png", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	file.Write(convertedData)
	defer file.Close()

	// Get convert result
	convertedData, httpResponse, err = instance.GetWorkbook(&GetWorkbookRequest{Name: employeeSalesSummaryXlsx, Format: "pdf", Folder: remoteFolder})
	if err != nil {
		fmt.Print(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Print(httpResponse.StatusCode)
	} else {
		file, _ := os.OpenFile("EmployeeSalesSummary2.pdf", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		file.Write(convertedData)
		defer file.Close()
	}
	// Convert a local Excel file to pdf file directly. The return value is a file info object, which includes of file name, file content(base64string), and file size.
	var fileInfo FileInfo
	fileInfo, httpResponse, err = instance.PostConvertWorkbookToPDF(&PostConvertWorkbookToPDFRequest{LocalPath: employeeSalesSummaryXlsx})
	decodedData, err := base64.StdEncoding.DecodeString(string(fileInfo.FileContent))
	file, _ = os.OpenFile("EmployeeSalesSummary3.pdf", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	file.Write(decodedData)
	defer file.Close()
	//Convert a local Excel file to pdf file directly. The return value is a file info object, which includes of file name, file content(base64string), and file size.
	fileInfo, httpResponse, err = instance.PostConvertWorkbookToJson(&PostConvertWorkbookToJsonRequest{LocalPath: employeeSalesSummaryXlsx})
	decodedData, err = base64.StdEncoding.DecodeString(string(fileInfo.FileContent))
	file, _ = os.OpenFile("EmployeeSalesSummary.json", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	file.Write(decodedData)
	defer file.Close()
	//Convert a local Excel file to docx file directly. The return value is a file info object, which includes of file name, file content(base64string), and file size.
	fileInfo, httpResponse, err = instance.PostConvertWorkbookToDocx(&PostConvertWorkbookToDocxRequest{LocalPath: employeeSalesSummaryXlsx})
	decodedData, err = base64.StdEncoding.DecodeString(string(fileInfo.FileContent))
	file, _ = os.OpenFile("EmployeeSalesSummary.docx", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	file.Write(decodedData)
	defer file.Close()
	//Convert a local Excel file to png file directly. The return value is a file info object, which includes of file name, file content(base64string), and file size.
	fileInfo, httpResponse, err = instance.PostConvertWorkbookToPNG(&PostConvertWorkbookToPNGRequest{LocalPath: employeeSalesSummaryXlsx})
	decodedData, err = base64.StdEncoding.DecodeString(string(fileInfo.FileContent))
	file, _ = os.OpenFile("EmployeeSalesSummary.png", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	file.Write(decodedData)
	defer file.Close()
	// Convert a local Excel file to pptx file directly. The return value is a file info object, which includes of file name, file content(base64string), and file size.
	fileInfo, httpResponse, err = instance.PostConvertWorkbookToPptx(&PostConvertWorkbookToPptxRequest{LocalPath: employeeSalesSummaryXlsx})
	decodedData, err = base64.StdEncoding.DecodeString(string(fileInfo.FileContent))
	file, _ = os.OpenFile("EmployeeSalesSummary.pptx", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	file.Write(decodedData)
	defer file.Close()
	// Convert a local Excel file to html file directly. The return value is a file info object, which includes of file name, file content(base64string), and file size.
	fileInfo, httpResponse, err = instance.PostConvertWorkbookToHtml(&PostConvertWorkbookToHtmlRequest{LocalPath: employeeSalesSummaryXlsx})
	decodedData, err = base64.StdEncoding.DecodeString(string(fileInfo.FileContent))
	file, _ = os.OpenFile("EmployeeSalesSummary.html", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	file.Write(decodedData)
	defer file.Close()

}
