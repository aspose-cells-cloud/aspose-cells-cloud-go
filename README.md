![](https://img.shields.io/badge/REST%20API-v3.0-lightgrey) [![GitHub license](https://img.shields.io/github/license/aspose-cells-cloud/aspose-cells-cloud-go)](https://github.com/aspose-cells-cloud/aspose-cells-cloud-go/blob/master/LICENSE) 

# Go API Client for Aspose.Cells Cloud

[Aspose.Cells Cloud SDK for Go](https://products.aspose.cloud/cells/go) empowers your Go applications to connect with excel document formats. The APIs let engineers read, convert, build, alter and control the substance of the [excel document formats](https://docs.aspose.cloud/cells/supported-file-formats/) without any office software installed on the machine.

## Excel® File Manipulation in the Cloud

- Create Excel files from scratch via API or [Smart Markers](https://docs.aspose.cloud/cells/create-excel-workbook-from-a-smartmarker-template/).
- Load, process & [convert Excel files](https://docs.aspose.cloud/cells/convert-excel-workbook-to-different-file-formats/) via Cloud SDK.
- Add, update or delete worksheet, charts, pictures, shapes, hyperlinks & validations.
- Add or remove cells area for conditional formatting from Excel worksheets.
- Insert or delete, horizontal or vertical page breaks.
- Add ListObject or convert ListObjects to a range of cells.
- Summarize data with [Pivot Tables](https://docs.aspose.cloud/cells/working-with-pivot-tables/) & Excel charts.
- Apply custom criteria to list filters of various types.
- Get, update, show or hide chart legend & titles.
- Manipulate page setup, header & footer.
- Create, update, fetch or delete document properties.
- Fetch the required shape from worksheet.
- Leverage the power of named ranges.

## Enhancements in Version 20.9

- Enhanced chart API.
- Added API to update Pivot Fields.
- Supported sparkline groups.

## Read & Write Spreadsheet Formats

**Microsoft Excel:** XLS, XLSX, XLSB, XLSM, XLT, XLTX, XLTM
**OpenOffice:** ODS
**SpreadsheetML:** XML
**Text:** CSV, TSV, TXT (TabDelimited)
**Web:** HTML, MHTML

## Save Spreadsheets As

**Microsoft Excel:** XLS, XLSX, XLSB
**OpenOffice:** ODS
**SpreadsheetML:** XML
**Text:** CSV, TSV, TXT (TabDelimited)
**Web:** HTML, MHTML
**Fixed Layout:** PDF, XPS
**Images:** PNG, JPG, TIFF, SVG
**Markdown:** MD
**Other:** DIF

## Read Other Formats

SXC, FODS

## Get Started with Aspose.Cells Cloud SDK for Go

First, create an account at [Aspose for Cloud](https://dashboard.aspose.cloud/#/apps) and get your application information. Then, follow these steps.

- Download the code and Add/Modify your application and refer to cells_cloud_test.go.
- If you want to download [Go Module](https://pkg.go.dev),please use `import "github.com/aspose-cells-cloud/aspose-cells-cloud-go/v20"` in your code.

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

## Add Worksheet to Excel File via Go

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

## Using Go to Convert an Excel File in the Cloud

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

## Aspose.Cells Cloud SDKs in Popular Languages

| .NET | Java | PHP | Python | Ruby | Node.js | Android | Swift | Perl |
|---|---|---|---|---|---|---|---|---|
| [GitHub](https://github.com/aspose-cells-cloud/aspose-cells-cloud-dotnet) | [GitHub](https://github.com/aspose-cells-cloud/aspose-cells-cloud-java) | [GitHub](https://github.com/aspose-cells-cloud/aspose-cells-cloud-php) | [GitHub](https://github.com/aspose-cells-cloud/aspose-cells-cloud-python)  | [GitHub](https://github.com/aspose-cells-cloud/aspose-cells-cloud-ruby) | [GitHub](https://github.com/aspose-cells-cloud/aspose-cells-cloud-node)  | [GitHub](https://github.com/aspose-cells-cloud/aspose-cells-cloud-android) | [GitHub](https://github.com/aspose-cells-cloud/aspose-cells-cloud-swift) | [GitHub](https://github.com/aspose-cells-cloud/aspose-cells-cloud-perl) |
| [NuGet](https://www.nuget.org/packages/Aspose.Cells-Cloud/) | [Maven](https://repository.aspose.cloud/webapp/#/artifacts/browse/tree/General/repo/com/aspose/aspose-cells-cloud) | [Composer](https://packagist.org/packages/aspose/cells-sdk-php) | [PIP](https://pypi.org/project/asposecellscloud/)  | [GEM](https://rubygems.org/gems/aspose_cells_cloud) | [NPM](https://www.npmjs.com/package/asposecellscloud) | [Maven](https://repository.aspose.cloud/webapp/#/artifacts/browse/tree/General/repo/com/aspose/aspose-cells-cloud-android) |  [POD](https://cocoapods.org/pods/AsposeCellsCloud) | [CPAN](https://metacpan.org/release/AsposeCellsCloud-CellsApi) |

[Product Page](https://products.aspose.cloud/cells/go) | [Documentation](https://docs.aspose.cloud/cells/) | [Live Demo](https://products.aspose.app/cells/family) | [API Reference](https://apireference.aspose.cloud/cells/) | [Code Samples](https://github.com/aspose-cells-cloud/aspose-cells-cloud-go) | [Blog](https://blog.aspose.cloud/category/cells/) | [Free Support](https://forum.aspose.cloud/c/cells) | [Free Trial](https://dashboard.aspose.cloud/#/apps)
