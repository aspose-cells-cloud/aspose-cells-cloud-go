# Go  API client for asposecellscloud

[Aspose.Cells Cloud SDK for Go](https://products.aspose.cloud/cells/go) empowers your Go applications to connect with excel document formats. The APIs let engineers read, convert, build, alter and control the substance of the [excel document formats](https://docs.aspose.cloud/display/cellscloud/Supported+File+Formats) without any office software installed on the machine..

## Spreadsheet Processing Features

- Add, update or delete charts, worksheet pictures, shapes, hyperlinks & validations.
- Add or remove cells area for conditional formatting, or OleObjects from Excel worksheets.
- Insert or delete, horizontal or vertical page breaks
- Add ListObject at a specific place within an Excel file & convert them to a range of cells.
- Delete specific or all ListObjects in a worksheet or summarize its data with pivot table.
- Apply custom criteria to list filters of various types.
- Get, update, show or hide chart legend & titles.
- Manipulate page setup, header & footer.
- Create, update, fetch or delete document properties.
- Fetch the required shape from worksheet.
- Load & Process Excel Spreadsheets via Cloud SDK.
- Cloud SDK to Read & Process Excel Worksheets.
- Leverage the Power of Pivot Tables & Ranges.

## Enhancements in Version 20.8

- Aspose.Cells Cloud API calls are not working with explicit storage name but only default storage.
- Get output file size without downloading during conversion.
- Enhancement for CellsShapesPutWorksheetShape API.

## Read & Write Spreadsheet Formats

**Microsoft Excel:** XLS, XLSX, XLSB, XLSM, XLT, XLTX, XLTM
**OpenOffice:** ODS
**SpreadsheetML:** XML
**Text:** CSV, TSV, TXT (TabDelimited)
**Web:** HTML, MHTML
**PDF**

## Save Spreadsheet As

DIF, HTML, MHTML,PNG,JPG, TIFF, XPS, SVG, MD (Markdown), ODS ,xlsx,xls,xlsb, PDF,XML,TXT,CSV

## Read Spreadsheet Formats

SXC, FODS

## Getting Started with Aspose.Cells Cloud SDK for Go

You do not need to install anything to get started with Aspose.Cells Cloud SDK for Go. Just create an account at [Aspose for Cloud](https://dashboard.aspose.cloud/#/apps) and get your application information.

### Installation
## How to use the SDK

- 1:Download the code and Add/Modify your application refer to cells_cloud_test.go
- 2:If you want to download by [Go Module](https://pkg.go.dev),please use
  import "github.com/aspose-cells-cloud/aspose-cells-cloud-go/v20" in your code
```
### Getting Started

Please follow the [installation](#installation) instruction and execute the following Go code:

```golang
func GetDocumentCircleAnnotations() (CircleAnnotationsResponse, *http.Response, error) {
    cellsAPI := NewCellsApiService("AppSid", "AppKey", "https://api.aspose.cloud","v3.0")
	name := "Book1.xlsx"	

	args := new(UploadFileOpts)
	args.Path = "GoTest/Booka1.xlsx"
	file, err := os.Open( "test_data/" + name)
	if err != nil {
		return err
	}

	_, _, err =cellsAPI.UploadFile(file, args)
	return err
}
```

Please check the [GitHub Repository](https://github.com/aspose-cells-cloud/aspose-cells-cloud-go) for other common usage scenarios.

## Using Go to Add a New Worksheet to an Excel File

The following code snippet demonstrates how to add a new worksheet to a Microsoft Excel document using C# code:

```golang
   name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsWorksheetsPutAddNewWorksheetOpts)
	args.Name = GetBook1()
	args.SheetName = GetSheet1()
	args.Position = 1
	args.Sheettype = "VB"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsWorksheetsPutAddNewWorksheet(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsWorksheetsPutAddNewWorksheet - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
```

## Using Go to Convert an Excel File to another File Format

The following code example elaborates how you can use C# code to convert an Excel document to another file format in the cloud:

```golang
// Upload source file to aspose cloud storage
name := GetBook1()
	if err := GetBaseTest().UploadFile(name); err != nil {
		t.Error(err)
	}

	args := new(CellsSaveAsPostDocumentSaveAsOpts)
	args.Name = GetBook1()
	args.Newfilename = "GoTest/newfilego.pdf"
	args.Folder = GetBaseTest().remoteFolder

	response, httpResponse, err := GetBaseTest().CellsAPI.CellsSaveAsPostDocumentSaveAs(args)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsSaveAsPostDocumentSaveAs - %d\n", GetBaseTest().GetTestNumber(), response.Code)
	}
```

## Licensing

All Aspose.Cells Cloud SDKs are licensed under [MIT License](https://github.com/aspose-cells-cloud/aspose-cells-cloud-go/blob/master/LICENSE).

[Product Page](https://products.aspose.cloud/cells/go) | [Documentation](https://docs.aspose.cloud/display/cellscloud/Home) | [Live Demo](https://products.aspose.app/cells/family) | [API Reference](https://apireference.aspose.cloud/cells/) | [Code Samples](https://github.com/aspose-cells-cloud/aspose-cells-cloud-go) | [Blog](https://blog.aspose.cloud/category/cells/) | [Free Support](https://forum.aspose.cloud/c/cells) | [Free Trial](https://dashboard.aspose.cloud/#/apps)




