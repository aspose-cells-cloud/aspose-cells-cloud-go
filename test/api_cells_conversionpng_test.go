package asposecellscloudtest

import (
	"fmt"
	"testing"

    . "asposecellscloud"
)

func TestConversionPng_ConvertWorkbook_html(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "cloud.png"
    remoteName := "cloud.png"

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
		fmt.Printf("%d\tTestConversionPng_ConvertWorkbook_html \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversionPng_ConvertWorkbook_mhtml(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "cloud.png"
    remoteName := "cloud.png"

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
		fmt.Printf("%d\tTestConversionPng_ConvertWorkbook_mhtml \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversionPng_ConvertWorkbook_ods(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "cloud.png"
    remoteName := "cloud.png"

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
		fmt.Printf("%d\tTestConversionPng_ConvertWorkbook_ods \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversionPng_ConvertWorkbook_pdf(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "cloud.png"
    remoteName := "cloud.png"

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
		fmt.Printf("%d\tTestConversionPng_ConvertWorkbook_pdf \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversionPng_ConvertWorkbook_xml(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "cloud.png"
    remoteName := "cloud.png"

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
		fmt.Printf("%d\tTestConversionPng_ConvertWorkbook_xml \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversionPng_ConvertWorkbook_tif(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "cloud.png"
    remoteName := "cloud.png"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "tif"

     
    request := new (PutConvertWorkbookRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          localName    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversionPng_ConvertWorkbook_tif \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversionPng_ConvertWorkbook_xps(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "cloud.png"
    remoteName := "cloud.png"

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
		fmt.Printf("%d\tTestConversionPng_ConvertWorkbook_xps \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversionPng_ConvertWorkbook_jpg(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "cloud.png"
    remoteName := "cloud.png"

    localNameRequest := new(UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =  GetBaseTest().localTestDataFolder  + localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    GetBaseTest().CellsApi.UploadFile(localNameRequest )
 
     format := "jpg"

     
    request := new (PutConvertWorkbookRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          localName    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PutConvertWorkbook(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestConversionPng_ConvertWorkbook_jpg \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversionPng_ConvertWorkbook_md(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "cloud.png"
    remoteName := "cloud.png"

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
		fmt.Printf("%d\tTestConversionPng_ConvertWorkbook_md \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversionPng_ConvertWorkbook_svg(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "cloud.png"
    remoteName := "cloud.png"

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
		fmt.Printf("%d\tTestConversionPng_ConvertWorkbook_svg \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversionPng_ConvertWorkbook_docx(t *testing.T) {
    remoteFolder := "TestData/In"
  
    localName := "cloud.png"
    remoteName := "cloud.png"

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
		fmt.Printf("%d\tTestConversionPng_ConvertWorkbook_docx \n", GetBaseTest().GetTestNumber())
	}
}

