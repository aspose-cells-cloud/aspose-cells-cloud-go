package asposecellscloudtest

import (
	"fmt"
	"testing"

	. "github.com/aspose-cells-cloud/aspose-cells-cloud-go/v24"
)

func TestTextProcessingController_PostAddTextContent(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "BookText.xlsx"
	remoteName := "BookText.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var addTextOptionsDataSource = new(DataSource)
	addTextOptionsDataSource.DataSourceType = "CloudFileSystem"
	addTextOptionsDataSource.DataPath = "BookText.xlsx"
	var addTextOptions = new(AddTextOptions)
	addTextOptions.DataSource = addTextOptionsDataSource
	addTextOptions.Text = "Aspose.Cells Cloud is an excellent product."
	addTextOptions.Worksheet = "202401"
	addTextOptions.SelectPoistion = "AtTheBeginning"
	addTextOptions.SkipEmptyCells = true

	request := new(PostAddTextContentRequest)
	request.AddTextOptions = addTextOptions
	_, httpResponse, err := GetBaseTest().CellsApi.PostAddTextContent(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestTextProcessingController_PostAddTextContent \n", GetBaseTest().GetTestNumber())
	}
}

func TestTextProcessingController_PostTrimContent(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "BookText.xlsx"
	remoteName := "BookText.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var trimContentOptionsDataSource = new(DataSource)
	trimContentOptionsDataSource.DataSourceType = "CloudFileSystem"
	trimContentOptionsDataSource.DataPath = "BookText.xlsx"
	var trimContentOptionsScopeOptions = new(ScopeOptions)
	trimContentOptionsScopeOptions.Scope = "EntireWorkbook"
	var trimContentOptions = new(TrimContentOptions)
	trimContentOptions.DataSource = trimContentOptionsDataSource
	trimContentOptions.TrimLeading = true
	trimContentOptions.TrimTrailing = true
	trimContentOptions.TrimSpaceBetweenWordTo1 = true
	trimContentOptions.RemoveAllLineBreaks = true
	trimContentOptions.ScopeOptions = trimContentOptionsScopeOptions

	request := new(PostTrimContentRequest)
	request.TrimContentOptions = trimContentOptions
	_, httpResponse, err := GetBaseTest().CellsApi.PostTrimContent(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestTextProcessingController_PostTrimContent \n", GetBaseTest().GetTestNumber())
	}
}

func TestTextProcessingController_PostUpdateWordCase(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "BookText.xlsx"
	remoteName := "BookText.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var wordCaseOptionsDataSource = new(DataSource)
	wordCaseOptionsDataSource.DataSourceType = "CloudFileSystem"
	wordCaseOptionsDataSource.DataPath = "BookText.xlsx"
	var wordCaseOptionsScopeOptions = new(ScopeOptions)
	wordCaseOptionsScopeOptions.Scope = "EntireWorkbook"
	var wordCaseOptions = new(WordCaseOptions)
	wordCaseOptions.DataSource = wordCaseOptionsDataSource
	wordCaseOptions.WordCaseType = "None"
	wordCaseOptions.ScopeOptions = wordCaseOptionsScopeOptions

	request := new(PostUpdateWordCaseRequest)
	request.WordCaseOptions = wordCaseOptions
	_, httpResponse, err := GetBaseTest().CellsApi.PostUpdateWordCase(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestTextProcessingController_PostUpdateWordCase \n", GetBaseTest().GetTestNumber())
	}
}
