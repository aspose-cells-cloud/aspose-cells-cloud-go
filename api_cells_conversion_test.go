package asposecellscloud

import (
	"fmt"
	"testing"
)

func TestConversion_WorkbookSaveAs_csv_OutResultPostExcelSaveAscsv(t *testing.T) {
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
     newfilename := "OutResult/PostExcelSaveAs.csv"

    var saveOptions = new(PdfSaveOptions)
     saveOptions.SaveFormat =        format      

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
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_csv_OutResultPostExcelSaveAscsv \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_html_OutResultPostExcelSaveAshtml(t *testing.T) {
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
     newfilename := "OutResult/PostExcelSaveAs.html"

    var saveOptions = new(PdfSaveOptions)
     saveOptions.SaveFormat =        format      

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
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_html_OutResultPostExcelSaveAshtml \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_mhtml_OutResultPostExcelSaveAsmhtml(t *testing.T) {
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
     newfilename := "OutResult/PostExcelSaveAs.mhtml"

    var saveOptions = new(PdfSaveOptions)
     saveOptions.SaveFormat =        format      

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
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_mhtml_OutResultPostExcelSaveAsmhtml \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_ods_OutResultPostExcelSaveAsods(t *testing.T) {
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
     newfilename := "OutResult/PostExcelSaveAs.ods"

    var saveOptions = new(PdfSaveOptions)
     saveOptions.SaveFormat =        format      

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
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_ods_OutResultPostExcelSaveAsods \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_pdf_OutResultPostExcelSaveAspdf(t *testing.T) {
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
     newfilename := "OutResult/PostExcelSaveAs.pdf"

    var saveOptions = new(PdfSaveOptions)
     saveOptions.SaveFormat =        format      

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
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_pdf_OutResultPostExcelSaveAspdf \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_xml_OutResultPostExcelSaveAsxml(t *testing.T) {
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
     newfilename := "OutResult/PostExcelSaveAs.xml"

    var saveOptions = new(PdfSaveOptions)
     saveOptions.SaveFormat =        format      

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
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_xml_OutResultPostExcelSaveAsxml \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_txt_OutResultPostExcelSaveAstxt(t *testing.T) {
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
     newfilename := "OutResult/PostExcelSaveAs.txt"

    var saveOptions = new(PdfSaveOptions)
     saveOptions.SaveFormat =        format      

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
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_txt_OutResultPostExcelSaveAstxt \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_tif_OutResultPostExcelSaveAstif(t *testing.T) {
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
     newfilename := "OutResult/PostExcelSaveAs.tif"

    var saveOptions = new(PdfSaveOptions)
     saveOptions.SaveFormat =        format      

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
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_tif_OutResultPostExcelSaveAstif \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_xlsb_OutResultPostExcelSaveAsxlsb(t *testing.T) {
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
     newfilename := "OutResult/PostExcelSaveAs.xlsb"

    var saveOptions = new(PdfSaveOptions)
     saveOptions.SaveFormat =        format      

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
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_xlsb_OutResultPostExcelSaveAsxlsb \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_xps_OutResultPostExcelSaveAsxps(t *testing.T) {
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
     newfilename := "OutResult/PostExcelSaveAs.xps"

    var saveOptions = new(PdfSaveOptions)
     saveOptions.SaveFormat =        format      

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
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_xps_OutResultPostExcelSaveAsxps \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_png_OutResultPostExcelSaveAspng(t *testing.T) {
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
     newfilename := "OutResult/PostExcelSaveAs.png"

    var saveOptions = new(PdfSaveOptions)
     saveOptions.SaveFormat =        format      

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
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_png_OutResultPostExcelSaveAspng \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_md_OutResultPostExcelSaveAsmd(t *testing.T) {
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
     newfilename := "OutResult/PostExcelSaveAs.md"

    var saveOptions = new(PdfSaveOptions)
     saveOptions.SaveFormat =        format      

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
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_md_OutResultPostExcelSaveAsmd \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_numbers_OutResultPostExcelSaveAsnumbers(t *testing.T) {
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
     newfilename := "OutResult/PostExcelSaveAs.numbers"

    var saveOptions = new(PdfSaveOptions)
     saveOptions.SaveFormat =        format      

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
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_numbers_OutResultPostExcelSaveAsnumbers \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_svg_OutResultPostExcelSaveAssvg(t *testing.T) {
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
     newfilename := "OutResult/PostExcelSaveAs.svg"

    var saveOptions = new(PdfSaveOptions)
     saveOptions.SaveFormat =        format      

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
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_svg_OutResultPostExcelSaveAssvg \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_docx_OutResultPostExcelSaveAsdocx(t *testing.T) {
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
     newfilename := "OutResult/PostExcelSaveAs.docx"

    var saveOptions = new(PdfSaveOptions)
     saveOptions.SaveFormat =        format      

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
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_docx_OutResultPostExcelSaveAsdocx \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_pptx_OutResultPostExcelSaveAspptx(t *testing.T) {
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
     newfilename := "OutResult/PostExcelSaveAs.pptx"

    var saveOptions = new(PdfSaveOptions)
     saveOptions.SaveFormat =        format      

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
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_pptx_OutResultPostExcelSaveAspptx \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_json_OutResultPostExcelSaveAsjson(t *testing.T) {
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
     newfilename := "OutResult/PostExcelSaveAs.json"

    var saveOptions = new(PdfSaveOptions)
     saveOptions.SaveFormat =        format      

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
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_json_OutResultPostExcelSaveAsjson \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_sql_OutResultPostExcelSaveAssql(t *testing.T) {
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
     newfilename := "OutResult/PostExcelSaveAs.sql"

    var saveOptions = new(PdfSaveOptions)
     saveOptions.SaveFormat =        format      

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
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_sql_OutResultPostExcelSaveAssql \n", GetBaseTest().GetTestNumber())
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

func TestConversion_ConvertWorkbookSaveCloud_csv_OutResultConvertWorkbookcsv(t *testing.T) {
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
     outPath := "OutResult/ConvertWorkbook.csv"

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
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_csv_OutResultConvertWorkbookcsv \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_xls_OutResultConvertWorkbookxls(t *testing.T) {
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
     outPath := "OutResult/ConvertWorkbook.xls"

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
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_xls_OutResultConvertWorkbookxls \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_html_OutResultConvertWorkbookhtml(t *testing.T) {
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
     outPath := "OutResult/ConvertWorkbook.html"

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
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_html_OutResultConvertWorkbookhtml \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_mhtml_OutResultConvertWorkbookmhtml(t *testing.T) {
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
     outPath := "OutResult/ConvertWorkbook.mhtml"

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
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_mhtml_OutResultConvertWorkbookmhtml \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_ods_OutResultConvertWorkbookods(t *testing.T) {
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
     outPath := "OutResult/ConvertWorkbook.ods"

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
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_ods_OutResultConvertWorkbookods \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_pdf_OutResultConvertWorkbookpdf(t *testing.T) {
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
     outPath := "OutResult/ConvertWorkbook.pdf"

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
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_pdf_OutResultConvertWorkbookpdf \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_xml_OutResultConvertWorkbookxml(t *testing.T) {
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
     outPath := "OutResult/ConvertWorkbook.xml"

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
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_xml_OutResultConvertWorkbookxml \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_txt_OutResultConvertWorkbooktxt(t *testing.T) {
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
     outPath := "OutResult/ConvertWorkbook.txt"

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
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_txt_OutResultConvertWorkbooktxt \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_tif_OutResultConvertWorkbooktif(t *testing.T) {
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
     outPath := "OutResult/ConvertWorkbook.tif"

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
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_tif_OutResultConvertWorkbooktif \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_xlsb_OutResultConvertWorkbookxlsb(t *testing.T) {
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
     outPath := "OutResult/ConvertWorkbook.xlsb"

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
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_xlsb_OutResultConvertWorkbookxlsb \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_xltm_OutResultConvertWorkbookxltm(t *testing.T) {
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
     outPath := "OutResult/ConvertWorkbook.xltm"

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
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_xltm_OutResultConvertWorkbookxltm \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_xps_OutResultConvertWorkbookxps(t *testing.T) {
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
     outPath := "OutResult/ConvertWorkbook.xps"

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
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_xps_OutResultConvertWorkbookxps \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_png_OutResultConvertWorkbookpng(t *testing.T) {
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
     outPath := "OutResult/ConvertWorkbook.png"

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
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_png_OutResultConvertWorkbookpng \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_md_OutResultConvertWorkbookmd(t *testing.T) {
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
     outPath := "OutResult/ConvertWorkbook.md"

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
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_md_OutResultConvertWorkbookmd \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_numbers_OutResultConvertWorkbooknumbers(t *testing.T) {
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
     outPath := "OutResult/ConvertWorkbook.numbers"

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
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_numbers_OutResultConvertWorkbooknumbers \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_svg_OutResultConvertWorkbooksvg(t *testing.T) {
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
     outPath := "OutResult/ConvertWorkbook.svg"

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
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_svg_OutResultConvertWorkbooksvg \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_docx_OutResultConvertWorkbookdocx(t *testing.T) {
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
     outPath := "OutResult/ConvertWorkbook.docx"

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
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_docx_OutResultConvertWorkbookdocx \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_pptx_OutResultConvertWorkbookpptx(t *testing.T) {
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
     outPath := "OutResult/ConvertWorkbook.pptx"

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
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_pptx_OutResultConvertWorkbookpptx \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_json_OutResultConvertWorkbookjson(t *testing.T) {
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
     outPath := "OutResult/ConvertWorkbook.json"

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
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_json_OutResultConvertWorkbookjson \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_sql_OutResultConvertWorkbooksql(t *testing.T) {
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
     outPath := "OutResult/ConvertWorkbook.sql"

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
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_sql_OutResultConvertWorkbooksql \n", GetBaseTest().GetTestNumber())
	}
}

