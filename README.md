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

## Feature & Enhancements in Version 24.1

Full list of issues covering all changes in this release:

- Fixed spelling mistakes for several functions.
- Add the PostFitTallToPages method for page setup controller.
- Add the PostFitWideToPages method for page setup controller.
- Optimize save options about paginated.


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
	remoteFolder := "TestData/In"
	localFolder := "testdata/"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	instance := asposecellscloud.NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"), "https://api.aspose.cloud", "v3.0")

	localNameRequest := new(asposecellscloud.UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = localFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	instance.UploadFile(localNameRequest)
}
```

## Add Worksheet to Excel File via Go

```golang
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 

    request := new (PutAddNewWorksheetRequest)
    request.Name =         remoteName    
    request.SheetName =         "Sheet1"    
    request.Position =  int64(0)        
    request.Sheettype =         "VB"    
    request.Folder =         remoteFolder    
    request.StorageName =         ""    
    _, httpResponse, err := GetBaseTest().CellsApi.PutAddNewWorksheet(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestWorksheetController_PutAddNewWorksheet \n", GetBaseTest().GetTestNumber())
	}
```

## Using Go to Convert an Excel File in the Cloud

```golang
// Upload source file to aspose cloud storage
	remoteFolder := "TestData/In"
	localFolder := "testdata/"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	instance := asposecellscloud.NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"), "https://api.aspose.cloud", "v3.0")

	localNameRequest := new(asposecellscloud.UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = localFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	instance.UploadFile(localNameRequest)

	newfilename := "TestData/OutResult/PostExcelSaveAs.pdf"

	var saveOptions = new(asposecellscloud.PdfSaveOptions)
	saveOptions.SaveFormat = "pdf"

	request := new(asposecellscloud.PostWorkbookSaveAsRequest)
	request.Name = remoteName
	request.Newfilename = newfilename
	request.SaveOptions = saveOptions
	request.Folder = remoteFolder
	_, httpResponse, err := instance.PostWorkbookSaveAs(request)
	if err != nil {
		println(err)
	}
	println(httpResponse.StatusCode)
```

## Aspose.Cells Cloud SDKs in Popular Languages

| .NET | Java | PHP | Python | Ruby | Node.js | Android | Swift | Perl |
|---|---|---|---|---|---|---|---|---|
| [GitHub](https://github.com/aspose-cells-cloud/aspose-cells-cloud-dotnet) | [GitHub](https://github.com/aspose-cells-cloud/aspose-cells-cloud-java) | [GitHub](https://github.com/aspose-cells-cloud/aspose-cells-cloud-php) | [GitHub](https://github.com/aspose-cells-cloud/aspose-cells-cloud-python)  | [GitHub](https://github.com/aspose-cells-cloud/aspose-cells-cloud-ruby) | [GitHub](https://github.com/aspose-cells-cloud/aspose-cells-cloud-node)  | [GitHub](https://github.com/aspose-cells-cloud/aspose-cells-cloud-android) | [GitHub](https://github.com/aspose-cells-cloud/aspose-cells-cloud-swift) | [GitHub](https://github.com/aspose-cells-cloud/aspose-cells-cloud-perl) |
| [NuGet](https://www.nuget.org/packages/Aspose.Cells-Cloud/) | [Maven](https://repository.aspose.cloud/webapp/#/artifacts/browse/tree/General/repo/com/aspose/aspose-cells-cloud) | [Composer](https://packagist.org/packages/aspose/cells-sdk-php) | [PIP](https://pypi.org/project/asposecellscloud/)  | [GEM](https://rubygems.org/gems/aspose_cells_cloud) | [NPM](https://www.npmjs.com/package/asposecellscloud) | [Maven](https://repository.aspose.cloud/webapp/#/artifacts/browse/tree/General/repo/com/aspose/aspose-cells-cloud-android) |  [POD](https://cocoapods.org/pods/AsposeCellsCloud) | [CPAN](https://metacpan.org/release/AsposeCellsCloud-CellsApi) |

[Product Page](https://products.aspose.cloud/cells/go) | [Documentation](https://docs.aspose.cloud/cells/) | [Live Demo](https://products.aspose.app/cells/family) | [API Reference](https://apireference.aspose.cloud/cells/) | [Code Samples](https://github.com/aspose-cells-cloud/aspose-cells-cloud-go) | [Blog](https://blog.aspose.cloud/category/cells/) | [Free Support](https://forum.aspose.cloud/c/cells) | [Free Trial](https://dashboard.aspose.cloud/#/apps)
