package main

import (
	"encoding/base64"
	"fmt"
	. "github.com/aspose-cells-cloud/aspose-cells-cloud-go/v25"
	"os"
)

func main() {
	employeeSalesSummaryXlsx := "EmployeeSalesSummary.xlsx"
	remoteFolder := "PythonSDK"
	instance := NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"))
	// Convert a local Excel file to another format file directly.
	convertedData, httpResponse, err := instance.PutConvertWorkbook(&PutConvertWorkbookRequest{LocalPath: employeeSalesSummaryXlsx, Format: "pdf"})
	if err != nil {
		fmt.Print(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Print(httpResponse.StatusCode)
	} else {
		file, _ := os.OpenFile("EmployeeSalesSummary1.pdf", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		file.Write(convertedData)
		defer file.Close()
	}

	// Upload a local Excel file to Cells Cloud Storage.
	_, httpResponse, err = instance.UploadFile(&UploadFileRequest{UploadFile: employeeSalesSummaryXlsx, Path: remoteFolder + "/" + employeeSalesSummaryXlsx})
	// Convert an Excel file of Cells Cloud to another format file.
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
	//Save an Excel file of Cells Cloud as another format file of Cells Cloud.
	_, httpResponse, err = instance.PostWorkbookSaveAs(&PostWorkbookSaveAsRequest{Name: employeeSalesSummaryXlsx, Newfilename: "EmployeeSalesSummary.pdf", Folder: remoteFolder, SaveOptions: &SaveOptions{SaveFormat: "pdf"}})
	if err != nil {
		println("Save as")
		fmt.Print(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Print(httpResponse.StatusCode)
	}
	// Convert a local Excel file to pdf file directly. The return value is a file info object, which includes of file name, file content(base64string), and file size.
	var fileInfo FileInfo
	fileInfo, httpResponse, err = instance.PostConvertWorkbookToPDF(&PostConvertWorkbookToPDFRequest{LocalPath: employeeSalesSummaryXlsx})
	decodedData, err := base64.StdEncoding.DecodeString(string(fileInfo.FileContent))
	file, _ := os.OpenFile("EmployeeSalesSummary3.pdf", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
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
	// Convert a worksheet of a local Excel file to another format file directly. Set query parameters : print_headings, one_page_per_sheet
	convertedData, httpResponse, err = instance.GetWorksheetWithFormat(&GetWorksheetWithFormatRequest{Name: employeeSalesSummaryXlsx, SheetName: "Sales", Folder: remoteFolder, Format: "png", PrintHeadings: true, OnePagePerSheet: false})
	if err != nil {
		fmt.Print(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Print(httpResponse.StatusCode)
	} else {
		file, _ := os.OpenFile("EmployeeSalesSummary_Sale.png", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		file.Write(convertedData)
		defer file.Close()
	}
	// Convert a local Excel file's specified worksheet page index directly to another format file. Set query parameters : print_headings, one_page_per_sheet
	convertedData, _, _ = instance.GetWorksheetWithFormat(&GetWorksheetWithFormatRequest{Name: employeeSalesSummaryXlsx, SheetName: "Sales", Folder: remoteFolder, Format: "png", PageIndex: 0, PrintHeadings: true, OnePagePerSheet: false})
	file, _ = os.OpenFile("EmployeeSalesSummary_Sale_PageIndex0.png", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	file.Write(convertedData)
	defer file.Close()
	//Convert a local Excel file's specified worksheet cells area directly to another format file. Set query parameters : print_headings, one_page_per_sheet
	convertedData, _, _ = instance.GetWorksheetWithFormat(&GetWorksheetWithFormatRequest{Name: employeeSalesSummaryXlsx, SheetName: "Sales", Folder: remoteFolder, Format: "png", Area: "B5:L36", PrintHeadings: true, OnePagePerSheet: false})
	file, _ = os.OpenFile("EmployeeSalesSummary_Sale_Area.png", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	file.Write(convertedData)
	defer file.Close()

}
