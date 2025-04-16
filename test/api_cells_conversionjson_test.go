package asposecellscloudtest

import (
	"fmt"
	"testing"

    . "asposecellscloud"
)

func TestConversionJson_ConvertWorkbook_csv(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "codegen-spec.json"
    remoteName := "codegen-spec.json"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "csv"

     
    request := new (PutConvertWorkbookRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          localName    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversionJson_ConvertWorkbook_csv \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversionJson_ConvertWorkbook_xls(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "codegen-spec.json"
    remoteName := "codegen-spec.json"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "xls"

     
    request := new (PutConvertWorkbookRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          localName    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversionJson_ConvertWorkbook_xls \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversionJson_ConvertWorkbook_html(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "codegen-spec.json"
    remoteName := "codegen-spec.json"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "html"

     
    request := new (PutConvertWorkbookRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          localName    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversionJson_ConvertWorkbook_html \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversionJson_ConvertWorkbook_mhtml(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "codegen-spec.json"
    remoteName := "codegen-spec.json"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "mhtml"

     
    request := new (PutConvertWorkbookRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          localName    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversionJson_ConvertWorkbook_mhtml \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversionJson_ConvertWorkbook_ods(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "codegen-spec.json"
    remoteName := "codegen-spec.json"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "ods"

     
    request := new (PutConvertWorkbookRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          localName    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversionJson_ConvertWorkbook_ods \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversionJson_ConvertWorkbook_pdf(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "codegen-spec.json"
    remoteName := "codegen-spec.json"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "pdf"

     
    request := new (PutConvertWorkbookRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          localName    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversionJson_ConvertWorkbook_pdf \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversionJson_ConvertWorkbook_xml(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "codegen-spec.json"
    remoteName := "codegen-spec.json"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "xml"

     
    request := new (PutConvertWorkbookRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          localName    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversionJson_ConvertWorkbook_xml \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversionJson_ConvertWorkbook_txt(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "codegen-spec.json"
    remoteName := "codegen-spec.json"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "txt"

     
    request := new (PutConvertWorkbookRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          localName    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversionJson_ConvertWorkbook_txt \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversionJson_ConvertWorkbook_xlsb(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "codegen-spec.json"
    remoteName := "codegen-spec.json"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "xlsb"

     
    request := new (PutConvertWorkbookRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          localName    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversionJson_ConvertWorkbook_xlsb \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversionJson_ConvertWorkbook_xps(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "codegen-spec.json"
    remoteName := "codegen-spec.json"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "xps"

     
    request := new (PutConvertWorkbookRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          localName    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversionJson_ConvertWorkbook_xps \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversionJson_ConvertWorkbook_md(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "codegen-spec.json"
    remoteName := "codegen-spec.json"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "md"

     
    request := new (PutConvertWorkbookRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          localName    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversionJson_ConvertWorkbook_md \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversionJson_ConvertWorkbook_svg(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "codegen-spec.json"
    remoteName := "codegen-spec.json"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "svg"

     
    request := new (PutConvertWorkbookRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          localName    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversionJson_ConvertWorkbook_svg \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversionJson_ConvertWorkbook_docx(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "codegen-spec.json"
    remoteName := "codegen-spec.json"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "docx"

     
    request := new (PutConvertWorkbookRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          localName    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversionJson_ConvertWorkbook_docx \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversionJson_ConvertWorkbook_pptx(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "codegen-spec.json"
    remoteName := "codegen-spec.json"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "pptx"

     
    request := new (PutConvertWorkbookRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          localName    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversionJson_ConvertWorkbook_pptx \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversionJson_ConvertWorkbook_json(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "codegen-spec.json"
    remoteName := "codegen-spec.json"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "json"

     
    request := new (PutConvertWorkbookRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          localName    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversionJson_ConvertWorkbook_json \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversionJson_ConvertWorkbook_sql(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "codegen-spec.json"
    remoteName := "codegen-spec.json"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "sql"

     
    request := new (PutConvertWorkbookRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          localName    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversionJson_ConvertWorkbook_sql \n", GetBaseTest().GetTestNumber())
	}
}

