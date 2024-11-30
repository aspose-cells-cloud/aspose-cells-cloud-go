package main

import (
	"os"

	asposecellscloud "github.com/aspose-cells-cloud/aspose-cells-cloud-go"
)
func main() {
	instance := asposecellscloud.NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"), "https://api.aspose.cloud", "v3.0")
    remoteFolder := "TestData/In"

    localName := "BookTableL2W.xlsx"
    remoteName := "BookTableL2W.xlsx"

    localNameRequest := new(asposecellscloud.UploadFileRequest)
    localNameRequest.UploadFiles = make(map[string]string) 
    localNameRequest.UploadFiles[localName] =   localName
    localNameRequest.Path = remoteFolder + "/" + remoteName 
    localNameRequest.StorageName =""
    instance.UploadFile(localNameRequest )
 
    var dataTransformationRequestLoadDataLoadTo = new(LoadTo)
     dataTransformationRequestLoadDataLoadTo.BeginColumnIndex = int64(2)          
     dataTransformationRequestLoadDataLoadTo.BeginRowIndex = int64(3)          
     dataTransformationRequestLoadDataLoadTo.Worksheet =        "L2W"      
    var dataTransformationRequestLoadDataDataQueryDataItem = new(DataItem)
     dataTransformationRequestLoadDataDataQueryDataItem.DataItemType =        "Table"      
     dataTransformationRequestLoadDataDataQueryDataItem.Value =        "Table1"      
    var dataTransformationRequestLoadDataDataQueryDataSource = new(DataSource)
     dataTransformationRequestLoadDataDataQueryDataSource.DataSourceType =        "CloudFileSystem"      
     dataTransformationRequestLoadDataDataQueryDataSource.DataPath =        "BookTableL2W.xlsx"      
    var dataTransformationRequestLoadDataDataQuery = new(DataQuery)
     dataTransformationRequestLoadDataDataQuery.Name =        "DataQuery"      
     dataTransformationRequestLoadDataDataQuery.DataItem =        dataTransformationRequestLoadDataDataQueryDataItem      
     dataTransformationRequestLoadDataDataQuery.DataSource =        dataTransformationRequestLoadDataDataQueryDataSource      
     dataTransformationRequestLoadDataDataQuery.DataSourceDataType =        "ListObject"      
    var dataTransformationRequestLoadData = new(LoadData)
     dataTransformationRequestLoadData.LoadTo =        dataTransformationRequestLoadDataLoadTo      
     dataTransformationRequestLoadData.DataQuery =        dataTransformationRequestLoadDataDataQuery      
    var dataTransformationRequestAppliedStepsAppliedStep0AppliedOperateUnpivotColumnNames = []string   {       "2017"    ,
           "2018"    ,
           "2019"    }    
    var dataTransformationRequestAppliedStepsAppliedStep0AppliedOperate = new(UnpivotColumn)
     dataTransformationRequestAppliedStepsAppliedStep0AppliedOperate.AppliedOperateType =        "UnpivotColumn"      
     dataTransformationRequestAppliedStepsAppliedStep0AppliedOperate.ValueMapName =        "Count"      
     dataTransformationRequestAppliedStepsAppliedStep0AppliedOperate.ColumnMapName =        "Date"      
     dataTransformationRequestAppliedStepsAppliedStep0AppliedOperate.UnpivotColumnNames =        dataTransformationRequestAppliedStepsAppliedStep0AppliedOperateUnpivotColumnNames      
    var dataTransformationRequestAppliedStepsAppliedStep0 = new(AppliedStep)
     dataTransformationRequestAppliedStepsAppliedStep0.StepName =        "UnpivotColumn"      
     dataTransformationRequestAppliedStepsAppliedStep0.AppliedOperate =        dataTransformationRequestAppliedStepsAppliedStep0AppliedOperate      
    var dataTransformationRequestAppliedSteps = []AppliedStep   {*       dataTransformationRequestAppliedStepsAppliedStep0    }    
    var dataTransformationRequest = new(DataTransformationRequest)
     dataTransformationRequest.LoadData =        dataTransformationRequestLoadData      
     dataTransformationRequest.AppliedSteps =        dataTransformationRequestAppliedSteps      

    request := new (asposecellscloud.PostDataTransformationRequest)
    request.DataTransformationRequest =         dataTransformationRequest    
    _, httpResponse, err := instance.PostDataTransformation(request)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	}
}
