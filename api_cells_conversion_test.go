package asposecellscloud

import (
	"fmt"
	"testing"
)

func TestConversion_WorkbookSaveAs_csv_DotNetSDKOutResultPostExcelSaveAscsv(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "csv"
     newfilename := "DotNetSDK/OutResult/PostExcelSaveAs.csv"

    var saveOptions = new(PdfSaveOptions);
     saveOptions.SaveFormat =        format      ;

    request := new (PostWorkbookSaveAsRequest)
    request.Name =         remoteName    
    request.Newfilename =         newfilename    
    request.SaveOptions =         saveOptions    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookSaveAs(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_csv_DotNetSDKOutResultPostExcelSaveAscsv \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_xls_DotNetSDKOutResultPostExcelSaveAsxls(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "xls"
     newfilename := "DotNetSDK/OutResult/PostExcelSaveAs.xls"

    var saveOptions = new(PdfSaveOptions);
     saveOptions.SaveFormat =        format      ;

    request := new (PostWorkbookSaveAsRequest)
    request.Name =         remoteName    
    request.Newfilename =         newfilename    
    request.SaveOptions =         saveOptions    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookSaveAs(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_xls_DotNetSDKOutResultPostExcelSaveAsxls \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_html_DotNetSDKOutResultPostExcelSaveAshtml(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "html"
     newfilename := "DotNetSDK/OutResult/PostExcelSaveAs.html"

    var saveOptions = new(PdfSaveOptions);
     saveOptions.SaveFormat =        format      ;

    request := new (PostWorkbookSaveAsRequest)
    request.Name =         remoteName    
    request.Newfilename =         newfilename    
    request.SaveOptions =         saveOptions    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookSaveAs(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_html_DotNetSDKOutResultPostExcelSaveAshtml \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_mhtml_DotNetSDKOutResultPostExcelSaveAsmhtml(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "mhtml"
     newfilename := "DotNetSDK/OutResult/PostExcelSaveAs.mhtml"

    var saveOptions = new(PdfSaveOptions);
     saveOptions.SaveFormat =        format      ;

    request := new (PostWorkbookSaveAsRequest)
    request.Name =         remoteName    
    request.Newfilename =         newfilename    
    request.SaveOptions =         saveOptions    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookSaveAs(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_mhtml_DotNetSDKOutResultPostExcelSaveAsmhtml \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_ods_DotNetSDKOutResultPostExcelSaveAsods(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "ods"
     newfilename := "DotNetSDK/OutResult/PostExcelSaveAs.ods"

    var saveOptions = new(PdfSaveOptions);
     saveOptions.SaveFormat =        format      ;

    request := new (PostWorkbookSaveAsRequest)
    request.Name =         remoteName    
    request.Newfilename =         newfilename    
    request.SaveOptions =         saveOptions    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookSaveAs(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_ods_DotNetSDKOutResultPostExcelSaveAsods \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_pdf_DotNetSDKOutResultPostExcelSaveAspdf(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "pdf"
     newfilename := "DotNetSDK/OutResult/PostExcelSaveAs.pdf"

    var saveOptions = new(PdfSaveOptions);
     saveOptions.SaveFormat =        format      ;

    request := new (PostWorkbookSaveAsRequest)
    request.Name =         remoteName    
    request.Newfilename =         newfilename    
    request.SaveOptions =         saveOptions    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookSaveAs(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_pdf_DotNetSDKOutResultPostExcelSaveAspdf \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_xml_DotNetSDKOutResultPostExcelSaveAsxml(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "xml"
     newfilename := "DotNetSDK/OutResult/PostExcelSaveAs.xml"

    var saveOptions = new(PdfSaveOptions);
     saveOptions.SaveFormat =        format      ;

    request := new (PostWorkbookSaveAsRequest)
    request.Name =         remoteName    
    request.Newfilename =         newfilename    
    request.SaveOptions =         saveOptions    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookSaveAs(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_xml_DotNetSDKOutResultPostExcelSaveAsxml \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_txt_DotNetSDKOutResultPostExcelSaveAstxt(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "txt"
     newfilename := "DotNetSDK/OutResult/PostExcelSaveAs.txt"

    var saveOptions = new(PdfSaveOptions);
     saveOptions.SaveFormat =        format      ;

    request := new (PostWorkbookSaveAsRequest)
    request.Name =         remoteName    
    request.Newfilename =         newfilename    
    request.SaveOptions =         saveOptions    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookSaveAs(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_txt_DotNetSDKOutResultPostExcelSaveAstxt \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_tif_DotNetSDKOutResultPostExcelSaveAstif(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "tif"
     newfilename := "DotNetSDK/OutResult/PostExcelSaveAs.tif"

    var saveOptions = new(PdfSaveOptions);
     saveOptions.SaveFormat =        format      ;

    request := new (PostWorkbookSaveAsRequest)
    request.Name =         remoteName    
    request.Newfilename =         newfilename    
    request.SaveOptions =         saveOptions    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookSaveAs(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_tif_DotNetSDKOutResultPostExcelSaveAstif \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_xlsb_DotNetSDKOutResultPostExcelSaveAsxlsb(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "xlsb"
     newfilename := "DotNetSDK/OutResult/PostExcelSaveAs.xlsb"

    var saveOptions = new(PdfSaveOptions);
     saveOptions.SaveFormat =        format      ;

    request := new (PostWorkbookSaveAsRequest)
    request.Name =         remoteName    
    request.Newfilename =         newfilename    
    request.SaveOptions =         saveOptions    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookSaveAs(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_xlsb_DotNetSDKOutResultPostExcelSaveAsxlsb \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_xlsm_DotNetSDKOutResultPostExcelSaveAsxlsm(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "xlsm"
     newfilename := "DotNetSDK/OutResult/PostExcelSaveAs.xlsm"

    var saveOptions = new(PdfSaveOptions);
     saveOptions.SaveFormat =        format      ;

    request := new (PostWorkbookSaveAsRequest)
    request.Name =         remoteName    
    request.Newfilename =         newfilename    
    request.SaveOptions =         saveOptions    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookSaveAs(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_xlsm_DotNetSDKOutResultPostExcelSaveAsxlsm \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_xlsx_DotNetSDKOutResultPostExcelSaveAsxlsx(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "xlsx"
     newfilename := "DotNetSDK/OutResult/PostExcelSaveAs.xlsx"

    var saveOptions = new(PdfSaveOptions);
     saveOptions.SaveFormat =        format      ;

    request := new (PostWorkbookSaveAsRequest)
    request.Name =         remoteName    
    request.Newfilename =         newfilename    
    request.SaveOptions =         saveOptions    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookSaveAs(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_xlsx_DotNetSDKOutResultPostExcelSaveAsxlsx \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_xltm_DotNetSDKOutResultPostExcelSaveAsxltm(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "xltm"
     newfilename := "DotNetSDK/OutResult/PostExcelSaveAs.xltm"

    var saveOptions = new(PdfSaveOptions);
     saveOptions.SaveFormat =        format      ;

    request := new (PostWorkbookSaveAsRequest)
    request.Name =         remoteName    
    request.Newfilename =         newfilename    
    request.SaveOptions =         saveOptions    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookSaveAs(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_xltm_DotNetSDKOutResultPostExcelSaveAsxltm \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_xltx_DotNetSDKOutResultPostExcelSaveAsxltx(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "xltx"
     newfilename := "DotNetSDK/OutResult/PostExcelSaveAs.xltx"

    var saveOptions = new(PdfSaveOptions);
     saveOptions.SaveFormat =        format      ;

    request := new (PostWorkbookSaveAsRequest)
    request.Name =         remoteName    
    request.Newfilename =         newfilename    
    request.SaveOptions =         saveOptions    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookSaveAs(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_xltx_DotNetSDKOutResultPostExcelSaveAsxltx \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_xps_DotNetSDKOutResultPostExcelSaveAsxps(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "xps"
     newfilename := "DotNetSDK/OutResult/PostExcelSaveAs.xps"

    var saveOptions = new(PdfSaveOptions);
     saveOptions.SaveFormat =        format      ;

    request := new (PostWorkbookSaveAsRequest)
    request.Name =         remoteName    
    request.Newfilename =         newfilename    
    request.SaveOptions =         saveOptions    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookSaveAs(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_xps_DotNetSDKOutResultPostExcelSaveAsxps \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_png_DotNetSDKOutResultPostExcelSaveAspng(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "png"
     newfilename := "DotNetSDK/OutResult/PostExcelSaveAs.png"

    var saveOptions = new(PdfSaveOptions);
     saveOptions.SaveFormat =        format      ;

    request := new (PostWorkbookSaveAsRequest)
    request.Name =         remoteName    
    request.Newfilename =         newfilename    
    request.SaveOptions =         saveOptions    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookSaveAs(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_png_DotNetSDKOutResultPostExcelSaveAspng \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_jpg_DotNetSDKOutResultPostExcelSaveAsjpg(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "jpg"
     newfilename := "DotNetSDK/OutResult/PostExcelSaveAs.jpg"

    var saveOptions = new(PdfSaveOptions);
     saveOptions.SaveFormat =        format      ;

    request := new (PostWorkbookSaveAsRequest)
    request.Name =         remoteName    
    request.Newfilename =         newfilename    
    request.SaveOptions =         saveOptions    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookSaveAs(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_jpg_DotNetSDKOutResultPostExcelSaveAsjpg \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_gif_DotNetSDKOutResultPostExcelSaveAsgif(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "gif"
     newfilename := "DotNetSDK/OutResult/PostExcelSaveAs.gif"

    var saveOptions = new(PdfSaveOptions);
     saveOptions.SaveFormat =        format      ;

    request := new (PostWorkbookSaveAsRequest)
    request.Name =         remoteName    
    request.Newfilename =         newfilename    
    request.SaveOptions =         saveOptions    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookSaveAs(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_gif_DotNetSDKOutResultPostExcelSaveAsgif \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_emf_DotNetSDKOutResultPostExcelSaveAsemf(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "emf"
     newfilename := "DotNetSDK/OutResult/PostExcelSaveAs.emf"

    var saveOptions = new(PdfSaveOptions);
     saveOptions.SaveFormat =        format      ;

    request := new (PostWorkbookSaveAsRequest)
    request.Name =         remoteName    
    request.Newfilename =         newfilename    
    request.SaveOptions =         saveOptions    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookSaveAs(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_emf_DotNetSDKOutResultPostExcelSaveAsemf \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_bmp_DotNetSDKOutResultPostExcelSaveAsbmp(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "bmp"
     newfilename := "DotNetSDK/OutResult/PostExcelSaveAs.bmp"

    var saveOptions = new(PdfSaveOptions);
     saveOptions.SaveFormat =        format      ;

    request := new (PostWorkbookSaveAsRequest)
    request.Name =         remoteName    
    request.Newfilename =         newfilename    
    request.SaveOptions =         saveOptions    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookSaveAs(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_bmp_DotNetSDKOutResultPostExcelSaveAsbmp \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_md_DotNetSDKOutResultPostExcelSaveAsmd(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "md"
     newfilename := "DotNetSDK/OutResult/PostExcelSaveAs.md"

    var saveOptions = new(PdfSaveOptions);
     saveOptions.SaveFormat =        format      ;

    request := new (PostWorkbookSaveAsRequest)
    request.Name =         remoteName    
    request.Newfilename =         newfilename    
    request.SaveOptions =         saveOptions    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookSaveAs(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_md_DotNetSDKOutResultPostExcelSaveAsmd \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_numbers_DotNetSDKOutResultPostExcelSaveAsnumbers(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "numbers"
     newfilename := "DotNetSDK/OutResult/PostExcelSaveAs.numbers"

    var saveOptions = new(PdfSaveOptions);
     saveOptions.SaveFormat =        format      ;

    request := new (PostWorkbookSaveAsRequest)
    request.Name =         remoteName    
    request.Newfilename =         newfilename    
    request.SaveOptions =         saveOptions    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookSaveAs(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_numbers_DotNetSDKOutResultPostExcelSaveAsnumbers \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_wmf_DotNetSDKOutResultPostExcelSaveAswmf(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "wmf"
     newfilename := "DotNetSDK/OutResult/PostExcelSaveAs.wmf"

    var saveOptions = new(PdfSaveOptions);
     saveOptions.SaveFormat =        format      ;

    request := new (PostWorkbookSaveAsRequest)
    request.Name =         remoteName    
    request.Newfilename =         newfilename    
    request.SaveOptions =         saveOptions    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookSaveAs(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_wmf_DotNetSDKOutResultPostExcelSaveAswmf \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_svg_DotNetSDKOutResultPostExcelSaveAssvg(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "svg"
     newfilename := "DotNetSDK/OutResult/PostExcelSaveAs.svg"

    var saveOptions = new(PdfSaveOptions);
     saveOptions.SaveFormat =        format      ;

    request := new (PostWorkbookSaveAsRequest)
    request.Name =         remoteName    
    request.Newfilename =         newfilename    
    request.SaveOptions =         saveOptions    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookSaveAs(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_svg_DotNetSDKOutResultPostExcelSaveAssvg \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_docx_DotNetSDKOutResultPostExcelSaveAsdocx(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "docx"
     newfilename := "DotNetSDK/OutResult/PostExcelSaveAs.docx"

    var saveOptions = new(PdfSaveOptions);
     saveOptions.SaveFormat =        format      ;

    request := new (PostWorkbookSaveAsRequest)
    request.Name =         remoteName    
    request.Newfilename =         newfilename    
    request.SaveOptions =         saveOptions    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookSaveAs(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_docx_DotNetSDKOutResultPostExcelSaveAsdocx \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_pptx_DotNetSDKOutResultPostExcelSaveAspptx(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "pptx"
     newfilename := "DotNetSDK/OutResult/PostExcelSaveAs.pptx"

    var saveOptions = new(PdfSaveOptions);
     saveOptions.SaveFormat =        format      ;

    request := new (PostWorkbookSaveAsRequest)
    request.Name =         remoteName    
    request.Newfilename =         newfilename    
    request.SaveOptions =         saveOptions    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookSaveAs(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_pptx_DotNetSDKOutResultPostExcelSaveAspptx \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_json_DotNetSDKOutResultPostExcelSaveAsjson(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "json"
     newfilename := "DotNetSDK/OutResult/PostExcelSaveAs.json"

    var saveOptions = new(PdfSaveOptions);
     saveOptions.SaveFormat =        format      ;

    request := new (PostWorkbookSaveAsRequest)
    request.Name =         remoteName    
    request.Newfilename =         newfilename    
    request.SaveOptions =         saveOptions    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookSaveAs(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_json_DotNetSDKOutResultPostExcelSaveAsjson \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_sql_DotNetSDKOutResultPostExcelSaveAssql(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "sql"
     newfilename := "DotNetSDK/OutResult/PostExcelSaveAs.sql"

    var saveOptions = new(PdfSaveOptions);
     saveOptions.SaveFormat =        format      ;

    request := new (PostWorkbookSaveAsRequest)
    request.Name =         remoteName    
    request.Newfilename =         newfilename    
    request.SaveOptions =         saveOptions    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookSaveAs(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_sql_DotNetSDKOutResultPostExcelSaveAssql \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_GetWorkbookFormat_csv(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "csv"


    request := new (GetWorkbookRequest)
    request.Name =         remoteName    
    request.Format =         format    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.GetWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_GetWorkbookFormat_csv \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_GetWorkbookFormat_xls(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "xls"


    request := new (GetWorkbookRequest)
    request.Name =         remoteName    
    request.Format =         format    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.GetWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_GetWorkbookFormat_xls \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_GetWorkbookFormat_html(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "html"


    request := new (GetWorkbookRequest)
    request.Name =         remoteName    
    request.Format =         format    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.GetWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_GetWorkbookFormat_html \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_GetWorkbookFormat_mhtml(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "mhtml"


    request := new (GetWorkbookRequest)
    request.Name =         remoteName    
    request.Format =         format    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.GetWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_GetWorkbookFormat_mhtml \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_GetWorkbookFormat_ods(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "ods"


    request := new (GetWorkbookRequest)
    request.Name =         remoteName    
    request.Format =         format    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.GetWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_GetWorkbookFormat_ods \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_GetWorkbookFormat_pdf(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "pdf"


    request := new (GetWorkbookRequest)
    request.Name =         remoteName    
    request.Format =         format    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.GetWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_GetWorkbookFormat_pdf \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_GetWorkbookFormat_xml(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "xml"


    request := new (GetWorkbookRequest)
    request.Name =         remoteName    
    request.Format =         format    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.GetWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_GetWorkbookFormat_xml \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_GetWorkbookFormat_txt(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "txt"


    request := new (GetWorkbookRequest)
    request.Name =         remoteName    
    request.Format =         format    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.GetWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_GetWorkbookFormat_txt \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_GetWorkbookFormat_tif(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "tif"


    request := new (GetWorkbookRequest)
    request.Name =         remoteName    
    request.Format =         format    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.GetWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_GetWorkbookFormat_tif \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_GetWorkbookFormat_xlsb(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "xlsb"


    request := new (GetWorkbookRequest)
    request.Name =         remoteName    
    request.Format =         format    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.GetWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_GetWorkbookFormat_xlsb \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_GetWorkbookFormat_xlsm(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "xlsm"


    request := new (GetWorkbookRequest)
    request.Name =         remoteName    
    request.Format =         format    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.GetWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_GetWorkbookFormat_xlsm \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_GetWorkbookFormat_xlsx(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "xlsx"


    request := new (GetWorkbookRequest)
    request.Name =         remoteName    
    request.Format =         format    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.GetWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_GetWorkbookFormat_xlsx \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_GetWorkbookFormat_xltm(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "xltm"


    request := new (GetWorkbookRequest)
    request.Name =         remoteName    
    request.Format =         format    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.GetWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_GetWorkbookFormat_xltm \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_GetWorkbookFormat_xltx(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "xltx"


    request := new (GetWorkbookRequest)
    request.Name =         remoteName    
    request.Format =         format    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.GetWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_GetWorkbookFormat_xltx \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_GetWorkbookFormat_xps(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "xps"


    request := new (GetWorkbookRequest)
    request.Name =         remoteName    
    request.Format =         format    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.GetWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_GetWorkbookFormat_xps \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_GetWorkbookFormat_png(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "png"


    request := new (GetWorkbookRequest)
    request.Name =         remoteName    
    request.Format =         format    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.GetWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_GetWorkbookFormat_png \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_GetWorkbookFormat_jpg(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "jpg"


    request := new (GetWorkbookRequest)
    request.Name =         remoteName    
    request.Format =         format    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.GetWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_GetWorkbookFormat_jpg \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_GetWorkbookFormat_gif(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "gif"


    request := new (GetWorkbookRequest)
    request.Name =         remoteName    
    request.Format =         format    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.GetWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_GetWorkbookFormat_gif \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_GetWorkbookFormat_emf(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "emf"


    request := new (GetWorkbookRequest)
    request.Name =         remoteName    
    request.Format =         format    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.GetWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_GetWorkbookFormat_emf \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_GetWorkbookFormat_bmp(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "bmp"


    request := new (GetWorkbookRequest)
    request.Name =         remoteName    
    request.Format =         format    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.GetWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_GetWorkbookFormat_bmp \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_GetWorkbookFormat_md(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "md"


    request := new (GetWorkbookRequest)
    request.Name =         remoteName    
    request.Format =         format    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.GetWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_GetWorkbookFormat_md \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_GetWorkbookFormat_numbers(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "numbers"


    request := new (GetWorkbookRequest)
    request.Name =         remoteName    
    request.Format =         format    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.GetWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_GetWorkbookFormat_numbers \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_GetWorkbookFormat_wmf(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "wmf"


    request := new (GetWorkbookRequest)
    request.Name =         remoteName    
    request.Format =         format    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.GetWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_GetWorkbookFormat_wmf \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_GetWorkbookFormat_svg(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "svg"


    request := new (GetWorkbookRequest)
    request.Name =         remoteName    
    request.Format =         format    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.GetWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_GetWorkbookFormat_svg \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_GetWorkbookFormat_docx(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "docx"


    request := new (GetWorkbookRequest)
    request.Name =         remoteName    
    request.Format =         format    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.GetWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_GetWorkbookFormat_docx \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_GetWorkbookFormat_pptx(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "pptx"


    request := new (GetWorkbookRequest)
    request.Name =         remoteName    
    request.Format =         format    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.GetWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_GetWorkbookFormat_pptx \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_GetWorkbookFormat_json(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "json"


    request := new (GetWorkbookRequest)
    request.Name =         remoteName    
    request.Format =         format    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.GetWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_GetWorkbookFormat_json \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_GetWorkbookFormat_sql(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "sql"


    request := new (GetWorkbookRequest)
    request.Name =         remoteName    
    request.Format =         format    
    request.Folder =         remoteFolder    
    _, httpResponse, err := GetBaseTest().CellsApi.GetWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_GetWorkbookFormat_sql \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbook_csv(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "csv"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbook_csv \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbook_xls(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "xls"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbook_xls \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbook_html(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "html"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbook_html \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbook_mhtml(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "mhtml"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbook_mhtml \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbook_ods(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "ods"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbook_ods \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbook_pdf(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "pdf"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbook_pdf \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbook_xml(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "xml"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbook_xml \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbook_txt(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "txt"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbook_txt \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbook_tif(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "tif"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbook_tif \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbook_xlsb(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "xlsb"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbook_xlsb \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbook_xlsm(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "xlsm"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbook_xlsm \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbook_xlsx(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "xlsx"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbook_xlsx \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbook_xltm(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "xltm"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbook_xltm \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbook_xltx(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "xltx"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbook_xltx \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbook_xps(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "xps"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbook_xps \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbook_png(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "png"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbook_png \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbook_jpg(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "jpg"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbook_jpg \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbook_gif(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "gif"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbook_gif \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbook_emf(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "emf"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbook_emf \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbook_bmp(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "bmp"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbook_bmp \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbook_md(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "md"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbook_md \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbook_numbers(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "numbers"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbook_numbers \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbook_wmf(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "wmf"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbook_wmf \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbook_svg(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "svg"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbook_svg \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbook_docx(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "docx"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbook_docx \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbook_pptx(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "pptx"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbook_pptx \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbook_json(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "json"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbook_json \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbook_sql(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "sql"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbook_sql \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_csv_DotNetSDKOutResultConvertWorkbookcsv(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "csv"
     outPath := "DotNetSDK/OutResult/ConvertWorkbook.csv"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    request.OutPath =         outPath    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_csv_DotNetSDKOutResultConvertWorkbookcsv \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_xls_DotNetSDKOutResultConvertWorkbookxls(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "xls"
     outPath := "DotNetSDK/OutResult/ConvertWorkbook.xls"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    request.OutPath =         outPath    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_xls_DotNetSDKOutResultConvertWorkbookxls \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_html_DotNetSDKOutResultConvertWorkbookhtml(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "html"
     outPath := "DotNetSDK/OutResult/ConvertWorkbook.html"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    request.OutPath =         outPath    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_html_DotNetSDKOutResultConvertWorkbookhtml \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_mhtml_DotNetSDKOutResultConvertWorkbookmhtml(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "mhtml"
     outPath := "DotNetSDK/OutResult/ConvertWorkbook.mhtml"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    request.OutPath =         outPath    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_mhtml_DotNetSDKOutResultConvertWorkbookmhtml \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_ods_DotNetSDKOutResultConvertWorkbookods(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "ods"
     outPath := "DotNetSDK/OutResult/ConvertWorkbook.ods"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    request.OutPath =         outPath    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_ods_DotNetSDKOutResultConvertWorkbookods \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_pdf_DotNetSDKOutResultConvertWorkbookpdf(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "pdf"
     outPath := "DotNetSDK/OutResult/ConvertWorkbook.pdf"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    request.OutPath =         outPath    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_pdf_DotNetSDKOutResultConvertWorkbookpdf \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_xml_DotNetSDKOutResultConvertWorkbookxml(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "xml"
     outPath := "DotNetSDK/OutResult/ConvertWorkbook.xml"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    request.OutPath =         outPath    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_xml_DotNetSDKOutResultConvertWorkbookxml \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_txt_DotNetSDKOutResultConvertWorkbooktxt(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "txt"
     outPath := "DotNetSDK/OutResult/ConvertWorkbook.txt"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    request.OutPath =         outPath    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_txt_DotNetSDKOutResultConvertWorkbooktxt \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_tif_DotNetSDKOutResultConvertWorkbooktif(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "tif"
     outPath := "DotNetSDK/OutResult/ConvertWorkbook.tif"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    request.OutPath =         outPath    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_tif_DotNetSDKOutResultConvertWorkbooktif \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_xlsb_DotNetSDKOutResultConvertWorkbookxlsb(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "xlsb"
     outPath := "DotNetSDK/OutResult/ConvertWorkbook.xlsb"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    request.OutPath =         outPath    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_xlsb_DotNetSDKOutResultConvertWorkbookxlsb \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_xlsm_DotNetSDKOutResultConvertWorkbookxlsm(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "xlsm"
     outPath := "DotNetSDK/OutResult/ConvertWorkbook.xlsm"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    request.OutPath =         outPath    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_xlsm_DotNetSDKOutResultConvertWorkbookxlsm \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_xlsx_DotNetSDKOutResultConvertWorkbookxlsx(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "xlsx"
     outPath := "DotNetSDK/OutResult/ConvertWorkbook.xlsx"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    request.OutPath =         outPath    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_xlsx_DotNetSDKOutResultConvertWorkbookxlsx \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_xltm_DotNetSDKOutResultConvertWorkbookxltm(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "xltm"
     outPath := "DotNetSDK/OutResult/ConvertWorkbook.xltm"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    request.OutPath =         outPath    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_xltm_DotNetSDKOutResultConvertWorkbookxltm \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_xltx_DotNetSDKOutResultConvertWorkbookxltx(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "xltx"
     outPath := "DotNetSDK/OutResult/ConvertWorkbook.xltx"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    request.OutPath =         outPath    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_xltx_DotNetSDKOutResultConvertWorkbookxltx \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_xps_DotNetSDKOutResultConvertWorkbookxps(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "xps"
     outPath := "DotNetSDK/OutResult/ConvertWorkbook.xps"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    request.OutPath =         outPath    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_xps_DotNetSDKOutResultConvertWorkbookxps \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_png_DotNetSDKOutResultConvertWorkbookpng(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "png"
     outPath := "DotNetSDK/OutResult/ConvertWorkbook.png"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    request.OutPath =         outPath    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_png_DotNetSDKOutResultConvertWorkbookpng \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_jpg_DotNetSDKOutResultConvertWorkbookjpg(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "jpg"
     outPath := "DotNetSDK/OutResult/ConvertWorkbook.jpg"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    request.OutPath =         outPath    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_jpg_DotNetSDKOutResultConvertWorkbookjpg \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_gif_DotNetSDKOutResultConvertWorkbookgif(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "gif"
     outPath := "DotNetSDK/OutResult/ConvertWorkbook.gif"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    request.OutPath =         outPath    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_gif_DotNetSDKOutResultConvertWorkbookgif \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_emf_DotNetSDKOutResultConvertWorkbookemf(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "emf"
     outPath := "DotNetSDK/OutResult/ConvertWorkbook.emf"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    request.OutPath =         outPath    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_emf_DotNetSDKOutResultConvertWorkbookemf \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_bmp_DotNetSDKOutResultConvertWorkbookbmp(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "bmp"
     outPath := "DotNetSDK/OutResult/ConvertWorkbook.bmp"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    request.OutPath =         outPath    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_bmp_DotNetSDKOutResultConvertWorkbookbmp \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_md_DotNetSDKOutResultConvertWorkbookmd(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "md"
     outPath := "DotNetSDK/OutResult/ConvertWorkbook.md"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    request.OutPath =         outPath    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_md_DotNetSDKOutResultConvertWorkbookmd \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_numbers_DotNetSDKOutResultConvertWorkbooknumbers(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "numbers"
     outPath := "DotNetSDK/OutResult/ConvertWorkbook.numbers"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    request.OutPath =         outPath    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_numbers_DotNetSDKOutResultConvertWorkbooknumbers \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_wmf_DotNetSDKOutResultConvertWorkbookwmf(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "wmf"
     outPath := "DotNetSDK/OutResult/ConvertWorkbook.wmf"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    request.OutPath =         outPath    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_wmf_DotNetSDKOutResultConvertWorkbookwmf \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_svg_DotNetSDKOutResultConvertWorkbooksvg(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "svg"
     outPath := "DotNetSDK/OutResult/ConvertWorkbook.svg"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    request.OutPath =         outPath    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_svg_DotNetSDKOutResultConvertWorkbooksvg \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_docx_DotNetSDKOutResultConvertWorkbookdocx(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "docx"
     outPath := "DotNetSDK/OutResult/ConvertWorkbook.docx"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    request.OutPath =         outPath    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_docx_DotNetSDKOutResultConvertWorkbookdocx \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_pptx_DotNetSDKOutResultConvertWorkbookpptx(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "pptx"
     outPath := "DotNetSDK/OutResult/ConvertWorkbook.pptx"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    request.OutPath =         outPath    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_pptx_DotNetSDKOutResultConvertWorkbookpptx \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_json_DotNetSDKOutResultConvertWorkbookjson(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "json"
     outPath := "DotNetSDK/OutResult/ConvertWorkbook.json"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    request.OutPath =         outPath    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_json_DotNetSDKOutResultConvertWorkbookjson \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_sql_DotNetSDKOutResultConvertWorkbooksql(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "Book1.xlsx"
    remoteName := "Book1.xlsx"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "sql"
     outPath := "DotNetSDK/OutResult/ConvertWorkbook.sql"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)

     mapFiles[localName]= GetBaseTest().localTestDataFolder + localName 

    request := new (PutConvertWorkbookRequest)
    request.File =         mapFiles    
    request.Format =         format    
    request.OutPath =         outPath    
    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_sql_DotNetSDKOutResultConvertWorkbooksql \n", GetBaseTest().GetTestNumber())
	}
}

