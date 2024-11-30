package asposecellscloud

import (
	. "asposecellscloud"
	"fmt"
	"testing"
)

func TestDataProcessingController_PostWorkbookDataCleansing(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "BookCsvDuplicateData.csv"
	remoteName := "BookCsvDuplicateData.csv"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var dataCleansingDataFillDataFillDefaultValue = new(DataFillValue)
	dataCleansingDataFillDataFillDefaultValue.DefaultDate = "2024-01-01"
	dataCleansingDataFillDataFillDefaultValue.DefaultNumber = int64(0)
	dataCleansingDataFillDataFillDefaultValue.DefaultBoolean = false
	var dataCleansingDataFill = new(DataFill)
	dataCleansingDataFill.DataFillDefaultValue = dataCleansingDataFillDataFillDefaultValue
	var dataCleansing = new(DataCleansing)
	dataCleansing.NeedFillData = true
	dataCleansing.DataFill = dataCleansingDataFill

	request := new(PostWorkbookDataCleansingRequest)
	request.Name = remoteName
	request.DataCleansing = dataCleansing
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookDataCleansing(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestDataProcessingController_PostWorkbookDataCleansing \n", GetBaseTest().GetTestNumber())
	}
}

func TestDataProcessingController_PostWorkbookDataDeduplication(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "BookCsvDuplicateData.csv"
	remoteName := "BookCsvDuplicateData.csv"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var deduplicationRegionRanges = []Range{}
	var deduplicationRegion = new(DeduplicationRegion)
	deduplicationRegion.Ranges = deduplicationRegionRanges

	request := new(PostWorkbookDataDeduplicationRequest)
	request.Name = remoteName
	request.DeduplicationRegion = deduplicationRegion
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookDataDeduplication(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestDataProcessingController_PostWorkbookDataDeduplication \n", GetBaseTest().GetTestNumber())
	}
}

func TestDataProcessingController_PostWorkbookDataFill(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "BookCsvDuplicateData.csv"
	remoteName := "BookCsvDuplicateData.csv"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var dataFillDataFillDefaultValue = new(DataFillValue)
	dataFillDataFillDefaultValue.DefaultDate = "2024-01-01"
	dataFillDataFillDefaultValue.DefaultNumber = int64(0)
	dataFillDataFillDefaultValue.DefaultBoolean = false
	var dataFill = new(DataFill)
	dataFill.DataFillDefaultValue = dataFillDataFillDefaultValue

	request := new(PostWorkbookDataFillRequest)
	request.Name = remoteName
	request.DataFill = dataFill
	request.Folder = remoteFolder
	request.StorageName = ""
	_, httpResponse, err := GetBaseTest().CellsApi.PostWorkbookDataFill(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestDataProcessingController_PostWorkbookDataFill \n", GetBaseTest().GetTestNumber())
	}
}

func TestDataProcessingController_PostDataTransformation(t *testing.T) {
	remoteFolder := "TestData/In"

	localName := "BookTableL2W.xlsx"
	remoteName := "BookTableL2W.xlsx"

	localNameRequest := new(UploadFileRequest)
	localNameRequest.UploadFiles = make(map[string]string)
	localNameRequest.UploadFiles[localName] = GetBaseTest().localTestDataFolder + localName
	localNameRequest.Path = remoteFolder + "/" + remoteName
	localNameRequest.StorageName = ""
	GetBaseTest().CellsApi.UploadFile(localNameRequest)

	var dataTransformationRequestLoadDataLoadTo = new(LoadTo)
	dataTransformationRequestLoadDataLoadTo.BeginColumnIndex = int64(2)
	dataTransformationRequestLoadDataLoadTo.BeginRowIndex = int64(3)
	dataTransformationRequestLoadDataLoadTo.Worksheet = "L2W"
	var dataTransformationRequestLoadDataDataQueryDataItem = new(DataItem)
	dataTransformationRequestLoadDataDataQueryDataItem.DataItemType = "Table"
	dataTransformationRequestLoadDataDataQueryDataItem.Value = "Table1"
	var dataTransformationRequestLoadDataDataQueryDataSource = new(DataSource)
	dataTransformationRequestLoadDataDataQueryDataSource.DataSourceType = "CloudFileSystem"
	dataTransformationRequestLoadDataDataQueryDataSource.DataPath = "BookTableL2W.xlsx"
	var dataTransformationRequestLoadDataDataQuery = new(DataQuery)
	dataTransformationRequestLoadDataDataQuery.Name = "DataQuery"
	dataTransformationRequestLoadDataDataQuery.DataItem = dataTransformationRequestLoadDataDataQueryDataItem
	dataTransformationRequestLoadDataDataQuery.DataSource = dataTransformationRequestLoadDataDataQueryDataSource
	dataTransformationRequestLoadDataDataQuery.DataSourceDataType = "ListObject"
	var dataTransformationRequestLoadData = new(LoadData)
	dataTransformationRequestLoadData.LoadTo = dataTransformationRequestLoadDataLoadTo
	dataTransformationRequestLoadData.DataQuery = dataTransformationRequestLoadDataDataQuery
	var dataTransformationRequestAppliedStepsAppliedStep0AppliedOperateUnpivotColumnNames = []string{"2017",
		"2018",
		"2019"}
	var dataTransformationRequestAppliedStepsAppliedStep0AppliedOperate = new(UnpivotColumn)
	dataTransformationRequestAppliedStepsAppliedStep0AppliedOperate.AppliedOperateType = "UnpivotColumn"
	dataTransformationRequestAppliedStepsAppliedStep0AppliedOperate.ValueMapName = "Count"
	dataTransformationRequestAppliedStepsAppliedStep0AppliedOperate.ColumnMapName = "Date"
	dataTransformationRequestAppliedStepsAppliedStep0AppliedOperate.UnpivotColumnNames = dataTransformationRequestAppliedStepsAppliedStep0AppliedOperateUnpivotColumnNames
	var dataTransformationRequestAppliedStepsAppliedStep0 = new(AppliedStep)
	dataTransformationRequestAppliedStepsAppliedStep0.StepName = "UnpivotColumn"
	dataTransformationRequestAppliedStepsAppliedStep0.AppliedOperate = dataTransformationRequestAppliedStepsAppliedStep0AppliedOperate
	var dataTransformationRequestAppliedSteps = []AppliedStep{*dataTransformationRequestAppliedStepsAppliedStep0}
	var dataTransformationRequest = new(DataTransformationRequest)
	dataTransformationRequest.LoadData = dataTransformationRequestLoadData
	dataTransformationRequest.AppliedSteps = dataTransformationRequestAppliedSteps

	request := new(PostDataTransformationRequest)
	request.DataTransformationRequest = dataTransformationRequest
	_, httpResponse, err := GetBaseTest().CellsApi.PostDataTransformation(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestDataProcessingController_PostDataTransformation \n", GetBaseTest().GetTestNumber())
	}
}
