package asposecellscloudtest

import (
	"fmt"
	"testing"

    . "asposecellscloud"
)

func TestLightCells_PostSplit_csv(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     outFormat := "csv"

     
    request := new (PostSplitRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.OutFormat =         outFormat    

    _, httpResponse, err := GetBaseTest().CellsApi.PostSplit(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostSplit_csv \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostSplit_xls(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     outFormat := "xls"

     
    request := new (PostSplitRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.OutFormat =         outFormat    

    _, httpResponse, err := GetBaseTest().CellsApi.PostSplit(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostSplit_xls \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostSplit_html(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     outFormat := "html"

     
    request := new (PostSplitRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.OutFormat =         outFormat    

    _, httpResponse, err := GetBaseTest().CellsApi.PostSplit(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostSplit_html \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostSplit_mhtml(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     outFormat := "mhtml"

     
    request := new (PostSplitRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.OutFormat =         outFormat    

    _, httpResponse, err := GetBaseTest().CellsApi.PostSplit(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostSplit_mhtml \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostSplit_ods(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     outFormat := "ods"

     
    request := new (PostSplitRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.OutFormat =         outFormat    

    _, httpResponse, err := GetBaseTest().CellsApi.PostSplit(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostSplit_ods \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostSplit_pdf(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     outFormat := "pdf"

     
    request := new (PostSplitRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.OutFormat =         outFormat    

    _, httpResponse, err := GetBaseTest().CellsApi.PostSplit(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostSplit_pdf \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostSplit_xml(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     outFormat := "xml"

     
    request := new (PostSplitRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.OutFormat =         outFormat    

    _, httpResponse, err := GetBaseTest().CellsApi.PostSplit(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostSplit_xml \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostSplit_txt(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     outFormat := "txt"

     
    request := new (PostSplitRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.OutFormat =         outFormat    

    _, httpResponse, err := GetBaseTest().CellsApi.PostSplit(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostSplit_txt \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostSplit_tif(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     outFormat := "tif"

     
    request := new (PostSplitRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.OutFormat =         outFormat    

    _, httpResponse, err := GetBaseTest().CellsApi.PostSplit(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostSplit_tif \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostSplit_xlsb(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     outFormat := "xlsb"

     
    request := new (PostSplitRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.OutFormat =         outFormat    

    _, httpResponse, err := GetBaseTest().CellsApi.PostSplit(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostSplit_xlsb \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostSplit_xlsx(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     outFormat := "xlsx"

     
    request := new (PostSplitRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.OutFormat =         outFormat    

    _, httpResponse, err := GetBaseTest().CellsApi.PostSplit(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostSplit_xlsx \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostSplit_xps(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     outFormat := "xps"

     
    request := new (PostSplitRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.OutFormat =         outFormat    

    _, httpResponse, err := GetBaseTest().CellsApi.PostSplit(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostSplit_xps \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostSplit_png(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     outFormat := "png"

     
    request := new (PostSplitRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.OutFormat =         outFormat    

    _, httpResponse, err := GetBaseTest().CellsApi.PostSplit(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostSplit_png \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostSplit_jpg(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     outFormat := "jpg"

     
    request := new (PostSplitRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.OutFormat =         outFormat    

    _, httpResponse, err := GetBaseTest().CellsApi.PostSplit(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostSplit_jpg \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostSplit_md(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     outFormat := "md"

     
    request := new (PostSplitRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.OutFormat =         outFormat    

    _, httpResponse, err := GetBaseTest().CellsApi.PostSplit(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostSplit_md \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostSplit_svg(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     outFormat := "svg"

     
    request := new (PostSplitRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.OutFormat =         outFormat    

    _, httpResponse, err := GetBaseTest().CellsApi.PostSplit(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostSplit_svg \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostSplit_docx(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     outFormat := "docx"

     
    request := new (PostSplitRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.OutFormat =         outFormat    

    _, httpResponse, err := GetBaseTest().CellsApi.PostSplit(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostSplit_docx \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostSplit_pptx(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     outFormat := "pptx"

     
    request := new (PostSplitRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.OutFormat =         outFormat    

    _, httpResponse, err := GetBaseTest().CellsApi.PostSplit(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostSplit_pptx \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostSplit_json(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     outFormat := "json"

     
    request := new (PostSplitRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.OutFormat =         outFormat    

    _, httpResponse, err := GetBaseTest().CellsApi.PostSplit(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostSplit_json \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostSplit_sql(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     outFormat := "sql"

     
    request := new (PostSplitRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.OutFormat =         outFormat    

    _, httpResponse, err := GetBaseTest().CellsApi.PostSplit(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostSplit_sql \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostAssemble_csv(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "csv"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostAssembleRequest)
    request.File =         mapFiles    


    request.Datasource =         "ds"    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostAssemble(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostAssemble_csv \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostAssemble_xls(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "xls"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostAssembleRequest)
    request.File =         mapFiles    


    request.Datasource =         "ds"    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostAssemble(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostAssemble_xls \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostAssemble_html(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "html"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostAssembleRequest)
    request.File =         mapFiles    


    request.Datasource =         "ds"    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostAssemble(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostAssemble_html \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostAssemble_mhtml(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "mhtml"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostAssembleRequest)
    request.File =         mapFiles    


    request.Datasource =         "ds"    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostAssemble(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostAssemble_mhtml \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostAssemble_ods(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "ods"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostAssembleRequest)
    request.File =         mapFiles    


    request.Datasource =         "ds"    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostAssemble(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostAssemble_ods \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostAssemble_pdf(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "pdf"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostAssembleRequest)
    request.File =         mapFiles    


    request.Datasource =         "ds"    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostAssemble(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostAssemble_pdf \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostAssemble_xml(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "xml"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostAssembleRequest)
    request.File =         mapFiles    


    request.Datasource =         "ds"    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostAssemble(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostAssemble_xml \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostAssemble_txt(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "txt"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostAssembleRequest)
    request.File =         mapFiles    


    request.Datasource =         "ds"    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostAssemble(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostAssemble_txt \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostAssemble_tif(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "tif"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostAssembleRequest)
    request.File =         mapFiles    


    request.Datasource =         "ds"    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostAssemble(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostAssemble_tif \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostAssemble_xlsb(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "xlsb"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostAssembleRequest)
    request.File =         mapFiles    


    request.Datasource =         "ds"    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostAssemble(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostAssemble_xlsb \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostAssemble_xlsm(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "xlsm"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostAssembleRequest)
    request.File =         mapFiles    


    request.Datasource =         "ds"    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostAssemble(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostAssemble_xlsm \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostAssemble_xlsx(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "xlsx"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostAssembleRequest)
    request.File =         mapFiles    


    request.Datasource =         "ds"    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostAssemble(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostAssemble_xlsx \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostAssemble_xltm(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "xltm"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostAssembleRequest)
    request.File =         mapFiles    


    request.Datasource =         "ds"    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostAssemble(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostAssemble_xltm \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostAssemble_xltx(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "xltx"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostAssembleRequest)
    request.File =         mapFiles    


    request.Datasource =         "ds"    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostAssemble(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostAssemble_xltx \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostAssemble_xps(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "xps"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostAssembleRequest)
    request.File =         mapFiles    


    request.Datasource =         "ds"    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostAssemble(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostAssemble_xps \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostAssemble_png(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "png"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostAssembleRequest)
    request.File =         mapFiles    


    request.Datasource =         "ds"    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostAssemble(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostAssemble_png \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostAssemble_jpg(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "jpg"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostAssembleRequest)
    request.File =         mapFiles    


    request.Datasource =         "ds"    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostAssemble(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostAssemble_jpg \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostAssemble_gif(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "gif"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostAssembleRequest)
    request.File =         mapFiles    


    request.Datasource =         "ds"    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostAssemble(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostAssemble_gif \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostAssemble_emf(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "emf"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostAssembleRequest)
    request.File =         mapFiles    


    request.Datasource =         "ds"    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostAssemble(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostAssemble_emf \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostAssemble_bmp(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "bmp"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostAssembleRequest)
    request.File =         mapFiles    


    request.Datasource =         "ds"    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostAssemble(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostAssemble_bmp \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostAssemble_md(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "md"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostAssembleRequest)
    request.File =         mapFiles    


    request.Datasource =         "ds"    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostAssemble(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostAssemble_md \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostAssemble_wmf(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "wmf"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostAssembleRequest)
    request.File =         mapFiles    


    request.Datasource =         "ds"    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostAssemble(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostAssemble_wmf \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostAssemble_svg(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "svg"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostAssembleRequest)
    request.File =         mapFiles    


    request.Datasource =         "ds"    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostAssemble(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostAssemble_svg \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostAssemble_docx(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "docx"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostAssembleRequest)
    request.File =         mapFiles    


    request.Datasource =         "ds"    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostAssemble(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostAssemble_docx \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostAssemble_pptx(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "pptx"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostAssembleRequest)
    request.File =         mapFiles    


    request.Datasource =         "ds"    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostAssemble(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostAssemble_pptx \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostAssemble_json(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "json"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostAssembleRequest)
    request.File =         mapFiles    


    request.Datasource =         "ds"    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostAssemble(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostAssemble_json \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostAssemble_sql(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "sql"

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostAssembleRequest)
    request.File =         mapFiles    


    request.Datasource =         "ds"    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostAssemble(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostAssemble_sql \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_csv_workbook(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "csv"
     objectType := "workbook"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_csv_workbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_xls_workbook(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "xls"
     objectType := "workbook"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_xls_workbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_html_workbook(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "html"
     objectType := "workbook"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_html_workbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_mhtml_workbook(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "mhtml"
     objectType := "workbook"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_mhtml_workbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_ods_workbook(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "ods"
     objectType := "workbook"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_ods_workbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_pdf_workbook(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "pdf"
     objectType := "workbook"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_pdf_workbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_xml_workbook(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "xml"
     objectType := "workbook"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_xml_workbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_txt_workbook(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "txt"
     objectType := "workbook"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_txt_workbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_tif_workbook(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "tif"
     objectType := "workbook"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_tif_workbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_xlsb_workbook(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "xlsb"
     objectType := "workbook"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_xlsb_workbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_xlsm_workbook(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "xlsm"
     objectType := "workbook"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_xlsm_workbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_xlsx_workbook(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "xlsx"
     objectType := "workbook"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_xlsx_workbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_xltm_workbook(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "xltm"
     objectType := "workbook"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_xltm_workbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_xltx_workbook(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "xltx"
     objectType := "workbook"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_xltx_workbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_xps_workbook(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "xps"
     objectType := "workbook"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_xps_workbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_png_workbook(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "png"
     objectType := "workbook"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_png_workbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_jpg_workbook(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "jpg"
     objectType := "workbook"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_jpg_workbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_gif_workbook(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "gif"
     objectType := "workbook"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_gif_workbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_emf_workbook(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "emf"
     objectType := "workbook"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_emf_workbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_bmp_workbook(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "bmp"
     objectType := "workbook"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_bmp_workbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_md_workbook(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "md"
     objectType := "workbook"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_md_workbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_wmf_workbook(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "wmf"
     objectType := "workbook"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_wmf_workbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_svg_workbook(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "svg"
     objectType := "workbook"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_svg_workbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_docx_workbook(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "docx"
     objectType := "workbook"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_docx_workbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_pptx_workbook(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "pptx"
     objectType := "workbook"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_pptx_workbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_json_workbook(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "json"
     objectType := "workbook"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_json_workbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_sql_workbook(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "sql"
     objectType := "workbook"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_sql_workbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_csv_worksheet(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "csv"
     objectType := "worksheet"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_csv_worksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_xls_worksheet(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "xls"
     objectType := "worksheet"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_xls_worksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_html_worksheet(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "html"
     objectType := "worksheet"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_html_worksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_mhtml_worksheet(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "mhtml"
     objectType := "worksheet"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_mhtml_worksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_ods_worksheet(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "ods"
     objectType := "worksheet"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_ods_worksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_pdf_worksheet(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "pdf"
     objectType := "worksheet"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_pdf_worksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_xml_worksheet(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "xml"
     objectType := "worksheet"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_xml_worksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_txt_worksheet(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "txt"
     objectType := "worksheet"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_txt_worksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_tif_worksheet(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "tif"
     objectType := "worksheet"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_tif_worksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_xlsb_worksheet(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "xlsb"
     objectType := "worksheet"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_xlsb_worksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_xlsm_worksheet(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "xlsm"
     objectType := "worksheet"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_xlsm_worksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_xlsx_worksheet(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "xlsx"
     objectType := "worksheet"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_xlsx_worksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_xltm_worksheet(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "xltm"
     objectType := "worksheet"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_xltm_worksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_xltx_worksheet(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "xltx"
     objectType := "worksheet"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_xltx_worksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_xps_worksheet(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "xps"
     objectType := "worksheet"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_xps_worksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_png_worksheet(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "png"
     objectType := "worksheet"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_png_worksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_jpg_worksheet(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "jpg"
     objectType := "worksheet"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_jpg_worksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_gif_worksheet(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "gif"
     objectType := "worksheet"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_gif_worksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_emf_worksheet(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "emf"
     objectType := "worksheet"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_emf_worksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_bmp_worksheet(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "bmp"
     objectType := "worksheet"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_bmp_worksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_md_worksheet(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "md"
     objectType := "worksheet"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_md_worksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_svg_worksheet(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "svg"
     objectType := "worksheet"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_svg_worksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_docx_worksheet(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "docx"
     objectType := "worksheet"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_docx_worksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_pptx_worksheet(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "pptx"
     objectType := "worksheet"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_pptx_worksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_json_worksheet(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "json"
     objectType := "worksheet"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_json_worksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_sql_worksheet(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "sql"
     objectType := "worksheet"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_sql_worksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_pdf_chart(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "pdf"
     objectType := "chart"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_pdf_chart \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_tif_chart(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "tif"
     objectType := "chart"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_tif_chart \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_png_chart(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "png"
     objectType := "chart"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_png_chart \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_jpg_chart(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "jpg"
     objectType := "chart"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_jpg_chart \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_gif_chart(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "gif"
     objectType := "chart"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_gif_chart \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_emf_chart(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "emf"
     objectType := "chart"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_emf_chart \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_bmp_chart(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "bmp"
     objectType := "chart"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_bmp_chart \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_png_picture(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "png"
     objectType := "picture"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_png_picture \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_jpg_picture(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "jpg"
     objectType := "picture"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_jpg_picture \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_gif_picture(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "gif"
     objectType := "picture"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_gif_picture \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_emf_picture(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "emf"
     objectType := "picture"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_emf_picture \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_bmp_picture(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "bmp"
     objectType := "picture"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_bmp_picture \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_csv_listobject(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "csv"
     objectType := "listobject"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_csv_listobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_xls_listobject(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "xls"
     objectType := "listobject"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_xls_listobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_html_listobject(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "html"
     objectType := "listobject"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_html_listobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_mhtml_listobject(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "mhtml"
     objectType := "listobject"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_mhtml_listobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_ods_listobject(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "ods"
     objectType := "listobject"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_ods_listobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_pdf_listobject(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "pdf"
     objectType := "listobject"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_pdf_listobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_xml_listobject(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "xml"
     objectType := "listobject"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_xml_listobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_txt_listobject(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "txt"
     objectType := "listobject"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_txt_listobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_tif_listobject(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "tif"
     objectType := "listobject"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_tif_listobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_xlsb_listobject(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "xlsb"
     objectType := "listobject"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_xlsb_listobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_xlsm_listobject(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "xlsm"
     objectType := "listobject"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_xlsm_listobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_xlsx_listobject(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "xlsx"
     objectType := "listobject"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_xlsx_listobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_xltm_listobject(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "xltm"
     objectType := "listobject"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_xltm_listobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_xltx_listobject(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "xltx"
     objectType := "listobject"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_xltx_listobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_xps_listobject(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "xps"
     objectType := "listobject"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_xps_listobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_png_listobject(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "png"
     objectType := "listobject"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_png_listobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_jpg_listobject(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "jpg"
     objectType := "listobject"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_jpg_listobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_gif_listobject(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "gif"
     objectType := "listobject"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_gif_listobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_emf_listobject(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "emf"
     objectType := "listobject"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_emf_listobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_bmp_listobject(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "bmp"
     objectType := "listobject"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_bmp_listobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_md_listobject(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "md"
     objectType := "listobject"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_md_listobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_svg_listobject(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "svg"
     objectType := "listobject"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_svg_listobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_docx_listobject(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "docx"
     objectType := "listobject"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_docx_listobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_pptx_listobject(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "pptx"
     objectType := "listobject"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_pptx_listobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_json_listobject(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "json"
     objectType := "listobject"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_json_listobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_sql_listobject(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "sql"
     objectType := "listobject"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_sql_listobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_png_oleobject(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "png"
     objectType := "oleobject"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_png_oleobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_jpg_oleobject(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "jpg"
     objectType := "oleobject"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_jpg_oleobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_gif_oleobject(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "gif"
     objectType := "oleobject"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_gif_oleobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_emf_oleobject(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "emf"
     objectType := "oleobject"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_emf_oleobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_bmp_oleobject(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "bmp"
     objectType := "oleobject"

     
    request := new (PostExportRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.ObjectType =         objectType    


    request.Format =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostExport(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_bmp_oleobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostCompress_50(t *testing.T) {
  
    dataSourceXlsx := "datasource.xlsx"

 
     compressLevel := 50

     
    request := new (PostCompressRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          dataSourceXlsx    


    request.CompressLevel =  int64(compressLevel)        

    _, httpResponse, err := GetBaseTest().CellsApi.PostCompress(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostCompress_50 \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostCompress_90(t *testing.T) {
  
    dataSourceXlsx := "datasource.xlsx"

 
     compressLevel := 90

     
    request := new (PostCompressRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          dataSourceXlsx    


    request.CompressLevel =  int64(compressLevel)        

    _, httpResponse, err := GetBaseTest().CellsApi.PostCompress(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostCompress_90 \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostMerge_csv_true(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "csv"
     mergeToOneSheet := true

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostMergeRequest)
    request.File =         mapFiles    


    request.OutFormat =         format    


    request.MergeToOneSheet =   mergeToOneSheet    

    _, httpResponse, err := GetBaseTest().CellsApi.PostMerge(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostMerge_csv_true \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostMerge_xls_true(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "xls"
     mergeToOneSheet := true

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostMergeRequest)
    request.File =         mapFiles    


    request.OutFormat =         format    


    request.MergeToOneSheet =   mergeToOneSheet    

    _, httpResponse, err := GetBaseTest().CellsApi.PostMerge(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostMerge_xls_true \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostMerge_html_true(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "html"
     mergeToOneSheet := true

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostMergeRequest)
    request.File =         mapFiles    


    request.OutFormat =         format    


    request.MergeToOneSheet =   mergeToOneSheet    

    _, httpResponse, err := GetBaseTest().CellsApi.PostMerge(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostMerge_html_true \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostMerge_mhtml_true(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "mhtml"
     mergeToOneSheet := true

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostMergeRequest)
    request.File =         mapFiles    


    request.OutFormat =         format    


    request.MergeToOneSheet =   mergeToOneSheet    

    _, httpResponse, err := GetBaseTest().CellsApi.PostMerge(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostMerge_mhtml_true \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostMerge_ods_true(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "ods"
     mergeToOneSheet := true

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostMergeRequest)
    request.File =         mapFiles    


    request.OutFormat =         format    


    request.MergeToOneSheet =   mergeToOneSheet    

    _, httpResponse, err := GetBaseTest().CellsApi.PostMerge(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostMerge_ods_true \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostMerge_pdf_true(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "pdf"
     mergeToOneSheet := true

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostMergeRequest)
    request.File =         mapFiles    


    request.OutFormat =         format    


    request.MergeToOneSheet =   mergeToOneSheet    

    _, httpResponse, err := GetBaseTest().CellsApi.PostMerge(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostMerge_pdf_true \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostMerge_xml_true(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "xml"
     mergeToOneSheet := true

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostMergeRequest)
    request.File =         mapFiles    


    request.OutFormat =         format    


    request.MergeToOneSheet =   mergeToOneSheet    

    _, httpResponse, err := GetBaseTest().CellsApi.PostMerge(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostMerge_xml_true \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostMerge_txt_true(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "txt"
     mergeToOneSheet := true

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostMergeRequest)
    request.File =         mapFiles    


    request.OutFormat =         format    


    request.MergeToOneSheet =   mergeToOneSheet    

    _, httpResponse, err := GetBaseTest().CellsApi.PostMerge(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostMerge_txt_true \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostMerge_tif_true(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "tif"
     mergeToOneSheet := true

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostMergeRequest)
    request.File =         mapFiles    


    request.OutFormat =         format    


    request.MergeToOneSheet =   mergeToOneSheet    

    _, httpResponse, err := GetBaseTest().CellsApi.PostMerge(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostMerge_tif_true \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostMerge_xlsb_true(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "xlsb"
     mergeToOneSheet := true

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostMergeRequest)
    request.File =         mapFiles    


    request.OutFormat =         format    


    request.MergeToOneSheet =   mergeToOneSheet    

    _, httpResponse, err := GetBaseTest().CellsApi.PostMerge(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostMerge_xlsb_true \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostMerge_xlsm_true(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "xlsm"
     mergeToOneSheet := true

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostMergeRequest)
    request.File =         mapFiles    


    request.OutFormat =         format    


    request.MergeToOneSheet =   mergeToOneSheet    

    _, httpResponse, err := GetBaseTest().CellsApi.PostMerge(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostMerge_xlsm_true \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostMerge_xlsx_true(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "xlsx"
     mergeToOneSheet := true

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostMergeRequest)
    request.File =         mapFiles    


    request.OutFormat =         format    


    request.MergeToOneSheet =   mergeToOneSheet    

    _, httpResponse, err := GetBaseTest().CellsApi.PostMerge(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostMerge_xlsx_true \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostMerge_xltm_true(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "xltm"
     mergeToOneSheet := true

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostMergeRequest)
    request.File =         mapFiles    


    request.OutFormat =         format    


    request.MergeToOneSheet =   mergeToOneSheet    

    _, httpResponse, err := GetBaseTest().CellsApi.PostMerge(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostMerge_xltm_true \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostMerge_xltx_true(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "xltx"
     mergeToOneSheet := true

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostMergeRequest)
    request.File =         mapFiles    


    request.OutFormat =         format    


    request.MergeToOneSheet =   mergeToOneSheet    

    _, httpResponse, err := GetBaseTest().CellsApi.PostMerge(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostMerge_xltx_true \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostMerge_xps_true(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "xps"
     mergeToOneSheet := true

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostMergeRequest)
    request.File =         mapFiles    


    request.OutFormat =         format    


    request.MergeToOneSheet =   mergeToOneSheet    

    _, httpResponse, err := GetBaseTest().CellsApi.PostMerge(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostMerge_xps_true \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostMerge_png_true(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "png"
     mergeToOneSheet := true

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostMergeRequest)
    request.File =         mapFiles    


    request.OutFormat =         format    


    request.MergeToOneSheet =   mergeToOneSheet    

    _, httpResponse, err := GetBaseTest().CellsApi.PostMerge(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostMerge_png_true \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostMerge_jpg_true(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "jpg"
     mergeToOneSheet := true

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostMergeRequest)
    request.File =         mapFiles    


    request.OutFormat =         format    


    request.MergeToOneSheet =   mergeToOneSheet    

    _, httpResponse, err := GetBaseTest().CellsApi.PostMerge(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostMerge_jpg_true \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostMerge_gif_true(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "gif"
     mergeToOneSheet := true

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostMergeRequest)
    request.File =         mapFiles    


    request.OutFormat =         format    


    request.MergeToOneSheet =   mergeToOneSheet    

    _, httpResponse, err := GetBaseTest().CellsApi.PostMerge(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostMerge_gif_true \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostMerge_emf_true(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "emf"
     mergeToOneSheet := true

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostMergeRequest)
    request.File =         mapFiles    


    request.OutFormat =         format    


    request.MergeToOneSheet =   mergeToOneSheet    

    _, httpResponse, err := GetBaseTest().CellsApi.PostMerge(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostMerge_emf_true \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostMerge_bmp_true(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "bmp"
     mergeToOneSheet := true

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostMergeRequest)
    request.File =         mapFiles    


    request.OutFormat =         format    


    request.MergeToOneSheet =   mergeToOneSheet    

    _, httpResponse, err := GetBaseTest().CellsApi.PostMerge(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostMerge_bmp_true \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostMerge_md_true(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "md"
     mergeToOneSheet := true

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostMergeRequest)
    request.File =         mapFiles    


    request.OutFormat =         format    


    request.MergeToOneSheet =   mergeToOneSheet    

    _, httpResponse, err := GetBaseTest().CellsApi.PostMerge(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostMerge_md_true \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostMerge_svg_true(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "svg"
     mergeToOneSheet := true

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostMergeRequest)
    request.File =         mapFiles    


    request.OutFormat =         format    


    request.MergeToOneSheet =   mergeToOneSheet    

    _, httpResponse, err := GetBaseTest().CellsApi.PostMerge(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostMerge_svg_true \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostMerge_docx_false(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "docx"
     mergeToOneSheet := false

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostMergeRequest)
    request.File =         mapFiles    


    request.OutFormat =         format    


    request.MergeToOneSheet =   mergeToOneSheet    

    _, httpResponse, err := GetBaseTest().CellsApi.PostMerge(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostMerge_docx_false \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostMerge_pptx_false(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "pptx"
     mergeToOneSheet := false

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostMergeRequest)
    request.File =         mapFiles    


    request.OutFormat =         format    


    request.MergeToOneSheet =   mergeToOneSheet    

    _, httpResponse, err := GetBaseTest().CellsApi.PostMerge(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostMerge_pptx_false \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostMerge_json_false(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "json"
     mergeToOneSheet := false

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostMergeRequest)
    request.File =         mapFiles    


    request.OutFormat =         format    


    request.MergeToOneSheet =   mergeToOneSheet    

    _, httpResponse, err := GetBaseTest().CellsApi.PostMerge(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostMerge_json_false \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostMerge_sql_false(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"
    dataSourceXlsx := "datasource.xlsx"

 
     format := "sql"
     mergeToOneSheet := false

    var mapFiles map[string]string       
    mapFiles = make(map[string]string)
     mapFiles[assemblyTestXlsx]= GetBaseTest().localTestDataFolder + assemblyTestXlsx 
     mapFiles[dataSourceXlsx]= GetBaseTest().localTestDataFolder + dataSourceXlsx 
    request := new (PostMergeRequest)
    request.File =         mapFiles    


    request.OutFormat =         format    


    request.MergeToOneSheet =   mergeToOneSheet    

    _, httpResponse, err := GetBaseTest().CellsApi.PostMerge(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostMerge_sql_false \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostUnlock(t *testing.T) {
  
    needUnlockXlsx := "needUnlock.xlsx"

 
     
    request := new (PostUnlockRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          needUnlockXlsx    


    request.Password =         "123456"    

    _, httpResponse, err := GetBaseTest().CellsApi.PostUnlock(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostUnlock \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostLock(t *testing.T) {
  
    needlockXlsx := "needlock.xlsx"

 
     
    request := new (PostLockRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          needlockXlsx    


    request.Password =         "123456"    

    _, httpResponse, err := GetBaseTest().CellsApi.PostLock(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostLock \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostProtect(t *testing.T) {
  
    assemblyTestXlsx := "assemblytest.xlsx"

 
    var protectWorkbookRequest = new(ProtectWorkbookRequest)
     protectWorkbookRequest.AwaysOpenReadOnly =  true      
     protectWorkbookRequest.EncryptWithPassword =        "123456"      
     
    request := new (PostProtectRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          assemblyTestXlsx    


    request.ProtectWorkbookRequest =         protectWorkbookRequest    


    request.Password =         "123456"    

    _, httpResponse, err := GetBaseTest().CellsApi.PostProtect(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostProtect \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostProtect_ProtectWorkbookRequest(t *testing.T) {
  
    dataSourceXlsx := "datasource.xlsx"

 
    var protectWorkbookRequest = new(ProtectWorkbookRequest)
     protectWorkbookRequest.AwaysOpenReadOnly =  true      
     protectWorkbookRequest.EncryptWithPassword =        "123456"      
     
    request := new (PostProtectRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          dataSourceXlsx    


    request.ProtectWorkbookRequest =         protectWorkbookRequest    

    _, httpResponse, err := GetBaseTest().CellsApi.PostProtect(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostProtect_ProtectWorkbookRequest \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostSearch(t *testing.T) {
  
    dataSourceXlsx := "datasource.xlsx"

 
     
    request := new (PostSearchRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          dataSourceXlsx    


    request.Text =         "12"    

    _, httpResponse, err := GetBaseTest().CellsApi.PostSearch(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostSearch \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostReplace(t *testing.T) {
  
    dataSourceXlsx := "datasource.xlsx"

 
     
    request := new (PostReplaceRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          dataSourceXlsx    


    request.Text =         "12"    


    request.Newtext =         "newtext"    

    _, httpResponse, err := GetBaseTest().CellsApi.PostReplace(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostReplace \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostReplaceOnlySheetname(t *testing.T) {
  
    dataSourceXlsx := "datasource.xlsx"

 
     
    request := new (PostReplaceRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          dataSourceXlsx    


    request.Text =         "12"    


    request.Newtext =         "newtext"    


    request.Sheetname =         "Sheet1"    

    _, httpResponse, err := GetBaseTest().CellsApi.PostReplace(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostReplaceOnlySheetname \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostWatermark(t *testing.T) {
  
    dataSourceXlsx := "datasource.xlsx"

 
     
    request := new (PostWatermarkRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          dataSourceXlsx    


    request.Text =         "aspose.cells cloud sdk"    


    request.Color =         "#773322"    

    _, httpResponse, err := GetBaseTest().CellsApi.PostWatermark(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostWatermark \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostClearObjects_chart(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     objecttype := "chart"

     
    request := new (PostClearObjectsRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.Objecttype =         objecttype    

    _, httpResponse, err := GetBaseTest().CellsApi.PostClearObjects(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostClearObjects_chart \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostClearObjects_comment(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     objecttype := "comment"

     
    request := new (PostClearObjectsRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.Objecttype =         objecttype    

    _, httpResponse, err := GetBaseTest().CellsApi.PostClearObjects(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostClearObjects_comment \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostClearObjects_picture(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     objecttype := "picture"

     
    request := new (PostClearObjectsRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.Objecttype =         objecttype    

    _, httpResponse, err := GetBaseTest().CellsApi.PostClearObjects(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostClearObjects_picture \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostClearObjects_shape(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     objecttype := "shape"

     
    request := new (PostClearObjectsRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.Objecttype =         objecttype    

    _, httpResponse, err := GetBaseTest().CellsApi.PostClearObjects(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostClearObjects_shape \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostClearObjects_listobject(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     objecttype := "listobject"

     
    request := new (PostClearObjectsRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.Objecttype =         objecttype    

    _, httpResponse, err := GetBaseTest().CellsApi.PostClearObjects(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostClearObjects_listobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostClearObjects_hyperlink(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     objecttype := "hyperlink"

     
    request := new (PostClearObjectsRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.Objecttype =         objecttype    

    _, httpResponse, err := GetBaseTest().CellsApi.PostClearObjects(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostClearObjects_hyperlink \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostClearObjects_oleobject(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     objecttype := "oleobject"

     
    request := new (PostClearObjectsRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.Objecttype =         objecttype    

    _, httpResponse, err := GetBaseTest().CellsApi.PostClearObjects(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostClearObjects_oleobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostClearObjects_pivottable(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     objecttype := "pivottable"

     
    request := new (PostClearObjectsRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.Objecttype =         objecttype    

    _, httpResponse, err := GetBaseTest().CellsApi.PostClearObjects(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostClearObjects_pivottable \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostClearObjects_validation(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     objecttype := "validation"

     
    request := new (PostClearObjectsRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.Objecttype =         objecttype    

    _, httpResponse, err := GetBaseTest().CellsApi.PostClearObjects(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostClearObjects_validation \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostClearObjects_Background(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     objecttype := "Background"

     
    request := new (PostClearObjectsRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.Objecttype =         objecttype    

    _, httpResponse, err := GetBaseTest().CellsApi.PostClearObjects(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostClearObjects_Background \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostRepair_xlsx(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "xlsx"

     
    request := new (PostRepairRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostRepair(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostRepair_xlsx \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostRepair_pdf(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "pdf"

     
    request := new (PostRepairRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostRepair(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostRepair_pdf \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostRepair_csv(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "csv"

     
    request := new (PostRepairRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostRepair(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostRepair_csv \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostRepair_png(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     format := "png"

     
    request := new (PostRepairRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostRepair(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostRepair_png \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostReverse_rows_pdf(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     rotateType := "rows"
     format := "pdf"

     
    request := new (PostReverseRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.RotateType =         rotateType    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostReverse(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostReverse_rows_pdf \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostReverse_cols_pdf(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     rotateType := "cols"
     format := "pdf"

     
    request := new (PostReverseRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.RotateType =         rotateType    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostReverse(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostReverse_cols_pdf \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostReverse_both_pdf(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     rotateType := "both"
     format := "pdf"

     
    request := new (PostReverseRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.RotateType =         rotateType    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostReverse(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostReverse_both_pdf \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostReverse_rows_csv(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     rotateType := "rows"
     format := "csv"

     
    request := new (PostReverseRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.RotateType =         rotateType    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostReverse(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostReverse_rows_csv \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostReverse_cols_png(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     rotateType := "cols"
     format := "png"

     
    request := new (PostReverseRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.RotateType =         rotateType    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostReverse(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostReverse_cols_png \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostReverse_both_xlsx(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     rotateType := "both"
     format := "xlsx"

     
    request := new (PostReverseRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.RotateType =         rotateType    


    request.OutFormat =         format    

    _, httpResponse, err := GetBaseTest().CellsApi.PostReverse(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostReverse_both_xlsx \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_GetMetadata(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     
    request := new (GetMetadataRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.Type_ =         "all"    

    _, httpResponse, err := GetBaseTest().CellsApi.GetMetadata(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_GetMetadata \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_DeleteMetadata(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
     
    request := new (DeleteMetadataRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.Type_ =         "all"    

    _, httpResponse, err := GetBaseTest().CellsApi.DeleteMetadata(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_DeleteMetadata \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostMetadata(t *testing.T) {
  
    book1Xlsx := "Book1.xlsx"

 
    var cellsDocumentscellsDocument0 = new(CellsDocumentProperty)
     cellsDocumentscellsDocument0.Name =        "Author"      
     cellsDocumentscellsDocument0.Value =        "roy.wang"      
    var cellsDocuments = []CellsDocumentProperty   {*       cellsDocumentscellsDocument0    }    
     
    request := new (PostMetadataRequest)
    request.LocalPath =  GetBaseTest().localTestDataFolder  +          book1Xlsx    


    request.CellsDocuments =         cellsDocuments    

    _, httpResponse, err := GetBaseTest().CellsApi.PostMetadata(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestLightCells_PostMetadata \n", GetBaseTest().GetTestNumber())
	}
}

