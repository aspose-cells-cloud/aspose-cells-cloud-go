# Aspose.Cells Cloud
- API version: 3.0
- Package version: 20.1
[Aspose.Cells Cloud](https://products.aspose.cloud/cells) is a true REST API that enables you to perform a wide range of document processing operations including creation, manipulation, conversion and rendering of Excel documents in the cloud.

Our Cloud SDKs are wrappers around REST API in various programming languages, allowing you to process documents in language of your choice quickly and easily, gaining all benefits of strong types and IDE highlights. This repository contains new generation SDKs for Aspose.Cells Cloud and examples.

These SDKs are now fully supported. If you have any questions, see any bugs or have enhancement request, feel free to reach out to us at [Free Support Forums](https://forum.aspose.cloud/c/cells).

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./asposecellscloud"
```
## Getting Started

Please follow the [installation](#installation) instruction and execute the following Go code:

```go
func GetDocumentCircleAnnotations() (CircleAnnotationsResponse, *http.Response, error) {
    cellsAPI := NewCellsApiService("AppSid", "AppKey", "")
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

## Unit Tests
Aspose Cells Cloud SDK includes a suite of unit tests within the "test" subdirectory. These Unit Tests also serves as examples of how to use the Aspose Cells Cloud SDK.

## Licensing
All Aspose.Cells Cloud SDKs are licensed under [MIT License](LICENSE).

## Resources
+ [**Website**](https://www.aspose.cloud)
+ [**Product Home**](https://products.aspose.cloud/cells/cloud)
+ [**Documentation**](https://docs.aspose.cloud/display/cellscloud/Home)
+ [**Free Support Forum**](https://forum.aspose.cloud/c/cells)
+ [**Paid Support Helpdesk**](https://helpdesk.aspose.cloud/)
+ [**Blog**](https://blog.aspose.cloud/category/aspose-products/aspose-cells-product-family/)

## Documentation for API Endpoints

All URIs are relative to *https://api.aspose.cloud/v3.0*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CellsApi* | [**CellsAutoFilterDeleteWorksheetDateFilter**](docs/CellsApi.md#cellsautofilterdeleteworksheetdatefilter) | **Delete** /cells/{name}/worksheets/{sheetName}/autoFilter/dateFilter | Removes a date filter.             
*CellsApi* | [**CellsAutoFilterDeleteWorksheetFilter**](docs/CellsApi.md#cellsautofilterdeleteworksheetfilter) | **Delete** /cells/{name}/worksheets/{sheetName}/autoFilter/filter | Delete a filter for a filter column.             
*CellsApi* | [**CellsAutoFilterGetWorksheetAutoFilter**](docs/CellsApi.md#cellsautofiltergetworksheetautofilter) | **Get** /cells/{name}/worksheets/{sheetName}/autoFilter | Get Auto filter Description
*CellsApi* | [**CellsAutoFilterPostWorksheetAutoFilterRefresh**](docs/CellsApi.md#cellsautofilterpostworksheetautofilterrefresh) | **Post** /cells/{name}/worksheets/{sheetName}/autoFilter/refresh | 
*CellsApi* | [**CellsAutoFilterPostWorksheetMatchBlanks**](docs/CellsApi.md#cellsautofilterpostworksheetmatchblanks) | **Post** /cells/{name}/worksheets/{sheetName}/autoFilter/matchBlanks | Match all blank cell in the list.
*CellsApi* | [**CellsAutoFilterPostWorksheetMatchNonBlanks**](docs/CellsApi.md#cellsautofilterpostworksheetmatchnonblanks) | **Post** /cells/{name}/worksheets/{sheetName}/autoFilter/matchNonBlanks | Match all not blank cell in the list.             
*CellsApi* | [**CellsAutoFilterPutWorksheetColorFilter**](docs/CellsApi.md#cellsautofilterputworksheetcolorfilter) | **Put** /cells/{name}/worksheets/{sheetName}/autoFilter/colorFilter | 
*CellsApi* | [**CellsAutoFilterPutWorksheetCustomFilter**](docs/CellsApi.md#cellsautofilterputworksheetcustomfilter) | **Put** /cells/{name}/worksheets/{sheetName}/autoFilter/custom | Filters a list with a custom criteria.             
*CellsApi* | [**CellsAutoFilterPutWorksheetDateFilter**](docs/CellsApi.md#cellsautofilterputworksheetdatefilter) | **Put** /cells/{name}/worksheets/{sheetName}/autoFilter/dateFilter | add date filter in worksheet 
*CellsApi* | [**CellsAutoFilterPutWorksheetDynamicFilter**](docs/CellsApi.md#cellsautofilterputworksheetdynamicfilter) | **Put** /cells/{name}/worksheets/{sheetName}/autoFilter/dynamicFilter | 
*CellsApi* | [**CellsAutoFilterPutWorksheetFilter**](docs/CellsApi.md#cellsautofilterputworksheetfilter) | **Put** /cells/{name}/worksheets/{sheetName}/autoFilter/filter | Adds a filter for a filter column.             
*CellsApi* | [**CellsAutoFilterPutWorksheetFilterTop10**](docs/CellsApi.md#cellsautofilterputworksheetfiltertop10) | **Put** /cells/{name}/worksheets/{sheetName}/autoFilter/filterTop10 | Filter the top 10 item in the list
*CellsApi* | [**CellsAutoFilterPutWorksheetIconFilter**](docs/CellsApi.md#cellsautofilterputworksheeticonfilter) | **Put** /cells/{name}/worksheets/{sheetName}/autoFilter/iconFilter | Adds an icon filter.
*CellsApi* | [**CellsAutoshapesGetWorksheetAutoshape**](docs/CellsApi.md#cellsautoshapesgetworksheetautoshape) | **Get** /cells/{name}/worksheets/{sheetName}/autoshapes/{autoshapeNumber} | Get autoshape info.
*CellsApi* | [**CellsAutoshapesGetWorksheetAutoshapes**](docs/CellsApi.md#cellsautoshapesgetworksheetautoshapes) | **Get** /cells/{name}/worksheets/{sheetName}/autoshapes | Get worksheet autoshapes info.
*CellsApi* | [**CellsChartAreaGetChartArea**](docs/CellsApi.md#cellschartareagetchartarea) | **Get** /cells/{name}/worksheets/{sheetName}/charts/{chartIndex}/chartArea | Get chart area info.
*CellsApi* | [**CellsChartAreaGetChartAreaBorder**](docs/CellsApi.md#cellschartareagetchartareaborder) | **Get** /cells/{name}/worksheets/{sheetName}/charts/{chartIndex}/chartArea/border | Get chart area border info.
*CellsApi* | [**CellsChartAreaGetChartAreaFillFormat**](docs/CellsApi.md#cellschartareagetchartareafillformat) | **Get** /cells/{name}/worksheets/{sheetName}/charts/{chartIndex}/chartArea/fillFormat | Get chart area fill format info.
*CellsApi* | [**CellsChartsDeleteWorksheetChartLegend**](docs/CellsApi.md#cellschartsdeleteworksheetchartlegend) | **Delete** /cells/{name}/worksheets/{sheetName}/charts/{chartIndex}/legend | Hide legend in chart
*CellsApi* | [**CellsChartsDeleteWorksheetChartTitle**](docs/CellsApi.md#cellschartsdeleteworksheetcharttitle) | **Delete** /cells/{name}/worksheets/{sheetName}/charts/{chartIndex}/title | Hide title in chart
*CellsApi* | [**CellsChartsDeleteWorksheetClearCharts**](docs/CellsApi.md#cellschartsdeleteworksheetclearcharts) | **Delete** /cells/{name}/worksheets/{sheetName}/charts | Clear the charts.
*CellsApi* | [**CellsChartsDeleteWorksheetDeleteChart**](docs/CellsApi.md#cellschartsdeleteworksheetdeletechart) | **Delete** /cells/{name}/worksheets/{sheetName}/charts/{chartIndex} | Delete worksheet chart by index.
*CellsApi* | [**CellsChartsGetWorksheetChart**](docs/CellsApi.md#cellschartsgetworksheetchart) | **Get** /cells/{name}/worksheets/{sheetName}/charts/{chartNumber} | Get chart info.
*CellsApi* | [**CellsChartsGetWorksheetChartLegend**](docs/CellsApi.md#cellschartsgetworksheetchartlegend) | **Get** /cells/{name}/worksheets/{sheetName}/charts/{chartIndex}/legend | Get chart legend
*CellsApi* | [**CellsChartsGetWorksheetChartTitle**](docs/CellsApi.md#cellschartsgetworksheetcharttitle) | **Get** /cells/{name}/worksheets/{sheetName}/charts/{chartIndex}/title | Get chart title
*CellsApi* | [**CellsChartsGetWorksheetCharts**](docs/CellsApi.md#cellschartsgetworksheetcharts) | **Get** /cells/{name}/worksheets/{sheetName}/charts | Get worksheet charts info.
*CellsApi* | [**CellsChartsPostWorksheetChart**](docs/CellsApi.md#cellschartspostworksheetchart) | **Post** /cells/{name}/worksheets/{sheetName}/charts/{chartIndex} | Update chart propreties
*CellsApi* | [**CellsChartsPostWorksheetChartLegend**](docs/CellsApi.md#cellschartspostworksheetchartlegend) | **Post** /cells/{name}/worksheets/{sheetName}/charts/{chartIndex}/legend | Update chart legend
*CellsApi* | [**CellsChartsPostWorksheetChartTitle**](docs/CellsApi.md#cellschartspostworksheetcharttitle) | **Post** /cells/{name}/worksheets/{sheetName}/charts/{chartIndex}/title | Update chart title
*CellsApi* | [**CellsChartsPutWorksheetAddChart**](docs/CellsApi.md#cellschartsputworksheetaddchart) | **Put** /cells/{name}/worksheets/{sheetName}/charts | Add new chart to worksheet.
*CellsApi* | [**CellsChartsPutWorksheetChartLegend**](docs/CellsApi.md#cellschartsputworksheetchartlegend) | **Put** /cells/{name}/worksheets/{sheetName}/charts/{chartIndex}/legend | Show legend in chart
*CellsApi* | [**CellsChartsPutWorksheetChartTitle**](docs/CellsApi.md#cellschartsputworksheetcharttitle) | **Put** /cells/{name}/worksheets/{sheetName}/charts/{chartIndex}/title | Add chart title / Set chart title visible
*CellsApi* | [**CellsConditionalFormattingsDeleteWorksheetConditionalFormatting**](docs/CellsApi.md#cellsconditionalformattingsdeleteworksheetconditionalformatting) | **Delete** /cells/{name}/worksheets/{sheetName}/conditionalFormattings/{index} | Remove conditional formatting
*CellsApi* | [**CellsConditionalFormattingsDeleteWorksheetConditionalFormattingArea**](docs/CellsApi.md#cellsconditionalformattingsdeleteworksheetconditionalformattingarea) | **Delete** /cells/{name}/worksheets/{sheetName}/conditionalFormattings/area | Remove cell area from conditional formatting.
*CellsApi* | [**CellsConditionalFormattingsDeleteWorksheetConditionalFormattings**](docs/CellsApi.md#cellsconditionalformattingsdeleteworksheetconditionalformattings) | **Delete** /cells/{name}/worksheets/{sheetName}/conditionalFormattings | Clear all condition formattings
*CellsApi* | [**CellsConditionalFormattingsGetWorksheetConditionalFormatting**](docs/CellsApi.md#cellsconditionalformattingsgetworksheetconditionalformatting) | **Get** /cells/{name}/worksheets/{sheetName}/conditionalFormattings/{index} | Get conditional formatting
*CellsApi* | [**CellsConditionalFormattingsGetWorksheetConditionalFormattings**](docs/CellsApi.md#cellsconditionalformattingsgetworksheetconditionalformattings) | **Get** /cells/{name}/worksheets/{sheetName}/conditionalFormattings | Get conditional formattings 
*CellsApi* | [**CellsConditionalFormattingsPutWorksheetConditionalFormatting**](docs/CellsApi.md#cellsconditionalformattingsputworksheetconditionalformatting) | **Put** /cells/{name}/worksheets/{sheetName}/conditionalFormattings | Add a condition formatting.
*CellsApi* | [**CellsConditionalFormattingsPutWorksheetFormatCondition**](docs/CellsApi.md#cellsconditionalformattingsputworksheetformatcondition) | **Put** /cells/{name}/worksheets/{sheetName}/conditionalFormattings/{index} | Add a format condition.
*CellsApi* | [**CellsConditionalFormattingsPutWorksheetFormatConditionArea**](docs/CellsApi.md#cellsconditionalformattingsputworksheetformatconditionarea) | **Put** /cells/{name}/worksheets/{sheetName}/conditionalFormattings/{index}/area | add a cell area for format condition             
*CellsApi* | [**CellsConditionalFormattingsPutWorksheetFormatConditionCondition**](docs/CellsApi.md#cellsconditionalformattingsputworksheetformatconditioncondition) | **Put** /cells/{name}/worksheets/{sheetName}/conditionalFormattings/{index}/condition | Add a condition for format condition.
*CellsApi* | [**CellsDeleteWorksheetColumns**](docs/CellsApi.md#cellsdeleteworksheetcolumns) | **Delete** /cells/{name}/worksheets/{sheetName}/cells/columns/{columnIndex} | Delete worksheet columns.
*CellsApi* | [**CellsDeleteWorksheetRow**](docs/CellsApi.md#cellsdeleteworksheetrow) | **Delete** /cells/{name}/worksheets/{sheetName}/cells/rows/{rowIndex} | Delete worksheet row.
*CellsApi* | [**CellsDeleteWorksheetRows**](docs/CellsApi.md#cellsdeleteworksheetrows) | **Delete** /cells/{name}/worksheets/{sheetName}/cells/rows | Delete several worksheet rows.
*CellsApi* | [**CellsGetCellHtmlString**](docs/CellsApi.md#cellsgetcellhtmlstring) | **Get** /cells/{name}/worksheets/{sheetName}/cells/{cellName}/htmlstring | Read cell data by cell&#39;s name.
*CellsApi* | [**CellsGetWorksheetCell**](docs/CellsApi.md#cellsgetworksheetcell) | **Get** /cells/{name}/worksheets/{sheetName}/cells/{cellOrMethodName} | Read cell data by cell&#39;s name.
*CellsApi* | [**CellsGetWorksheetCellStyle**](docs/CellsApi.md#cellsgetworksheetcellstyle) | **Get** /cells/{name}/worksheets/{sheetName}/cells/{cellName}/style | Read cell&#39;s style info.
*CellsApi* | [**CellsGetWorksheetCells**](docs/CellsApi.md#cellsgetworksheetcells) | **Get** /cells/{name}/worksheets/{sheetName}/cells | Get cells info.
*CellsApi* | [**CellsGetWorksheetColumn**](docs/CellsApi.md#cellsgetworksheetcolumn) | **Get** /cells/{name}/worksheets/{sheetName}/cells/columns/{columnIndex} | Read worksheet column data by column&#39;s index.
*CellsApi* | [**CellsGetWorksheetColumns**](docs/CellsApi.md#cellsgetworksheetcolumns) | **Get** /cells/{name}/worksheets/{sheetName}/cells/columns | Read worksheet columns info.
*CellsApi* | [**CellsGetWorksheetRow**](docs/CellsApi.md#cellsgetworksheetrow) | **Get** /cells/{name}/worksheets/{sheetName}/cells/rows/{rowIndex} | Read worksheet row data by row&#39;s index.
*CellsApi* | [**CellsGetWorksheetRows**](docs/CellsApi.md#cellsgetworksheetrows) | **Get** /cells/{name}/worksheets/{sheetName}/cells/rows | Read worksheet rows info.
*CellsApi* | [**CellsHypelinksDeleteWorksheetHyperlink**](docs/CellsApi.md#cellshypelinksdeleteworksheethyperlink) | **Delete** /cells/{name}/worksheets/{sheetName}/hyperlinks/{hyperlinkIndex} | Delete worksheet hyperlink by index.
*CellsApi* | [**CellsHypelinksDeleteWorksheetHyperlinks**](docs/CellsApi.md#cellshypelinksdeleteworksheethyperlinks) | **Delete** /cells/{name}/worksheets/{sheetName}/hyperlinks | Delete all hyperlinks in worksheet.
*CellsApi* | [**CellsHypelinksGetWorksheetHyperlink**](docs/CellsApi.md#cellshypelinksgetworksheethyperlink) | **Get** /cells/{name}/worksheets/{sheetName}/hyperlinks/{hyperlinkIndex} | Get worksheet hyperlink by index.
*CellsApi* | [**CellsHypelinksGetWorksheetHyperlinks**](docs/CellsApi.md#cellshypelinksgetworksheethyperlinks) | **Get** /cells/{name}/worksheets/{sheetName}/hyperlinks | Get worksheet hyperlinks.
*CellsApi* | [**CellsHypelinksPostWorksheetHyperlink**](docs/CellsApi.md#cellshypelinkspostworksheethyperlink) | **Post** /cells/{name}/worksheets/{sheetName}/hyperlinks/{hyperlinkIndex} | Update worksheet hyperlink by index.
*CellsApi* | [**CellsHypelinksPutWorksheetHyperlink**](docs/CellsApi.md#cellshypelinksputworksheethyperlink) | **Put** /cells/{name}/worksheets/{sheetName}/hyperlinks | Add worksheet hyperlink.
*CellsApi* | [**CellsListObjectsDeleteWorksheetListObject**](docs/CellsApi.md#cellslistobjectsdeleteworksheetlistobject) | **Delete** /cells/{name}/worksheets/{sheetName}/listobjects/{listObjectIndex} | Delete worksheet list object by index
*CellsApi* | [**CellsListObjectsDeleteWorksheetListObjects**](docs/CellsApi.md#cellslistobjectsdeleteworksheetlistobjects) | **Delete** /cells/{name}/worksheets/{sheetName}/listobjects | Delete worksheet list objects
*CellsApi* | [**CellsListObjectsGetWorksheetListObject**](docs/CellsApi.md#cellslistobjectsgetworksheetlistobject) | **Get** /cells/{name}/worksheets/{sheetName}/listobjects/{listobjectindex} | Get worksheet list object info by index.
*CellsApi* | [**CellsListObjectsGetWorksheetListObjects**](docs/CellsApi.md#cellslistobjectsgetworksheetlistobjects) | **Get** /cells/{name}/worksheets/{sheetName}/listobjects | Get worksheet listobjects info.
*CellsApi* | [**CellsListObjectsPostWorksheetListObject**](docs/CellsApi.md#cellslistobjectspostworksheetlistobject) | **Post** /cells/{name}/worksheets/{sheetName}/listobjects/{listObjectIndex} | Update  list object 
*CellsApi* | [**CellsListObjectsPostWorksheetListObjectConvertToRange**](docs/CellsApi.md#cellslistobjectspostworksheetlistobjectconverttorange) | **Post** /cells/{name}/worksheets/{sheetName}/listobjects/{listObjectIndex}/ConvertToRange | 
*CellsApi* | [**CellsListObjectsPostWorksheetListObjectSortTable**](docs/CellsApi.md#cellslistobjectspostworksheetlistobjectsorttable) | **Post** /cells/{name}/worksheets/{sheetName}/listobjects/{listObjectIndex}/sort | 
*CellsApi* | [**CellsListObjectsPostWorksheetListObjectSummarizeWithPivotTable**](docs/CellsApi.md#cellslistobjectspostworksheetlistobjectsummarizewithpivottable) | **Post** /cells/{name}/worksheets/{sheetName}/listobjects/{listObjectIndex}/SummarizeWithPivotTable | 
*CellsApi* | [**CellsListObjectsPutWorksheetListObject**](docs/CellsApi.md#cellslistobjectsputworksheetlistobject) | **Put** /cells/{name}/worksheets/{sheetName}/listobjects | Add a list object into worksheet.
*CellsApi* | [**CellsOleObjectsDeleteWorksheetOleObject**](docs/CellsApi.md#cellsoleobjectsdeleteworksheetoleobject) | **Delete** /cells/{name}/worksheets/{sheetName}/oleobjects/{oleObjectIndex} | Delete OLE object.
*CellsApi* | [**CellsOleObjectsDeleteWorksheetOleObjects**](docs/CellsApi.md#cellsoleobjectsdeleteworksheetoleobjects) | **Delete** /cells/{name}/worksheets/{sheetName}/oleobjects | Delete all OLE objects.
*CellsApi* | [**CellsOleObjectsGetWorksheetOleObject**](docs/CellsApi.md#cellsoleobjectsgetworksheetoleobject) | **Get** /cells/{name}/worksheets/{sheetName}/oleobjects/{objectNumber} | Get OLE object info.
*CellsApi* | [**CellsOleObjectsGetWorksheetOleObjects**](docs/CellsApi.md#cellsoleobjectsgetworksheetoleobjects) | **Get** /cells/{name}/worksheets/{sheetName}/oleobjects | Get worksheet OLE objects info.
*CellsApi* | [**CellsOleObjectsPostUpdateWorksheetOleObject**](docs/CellsApi.md#cellsoleobjectspostupdateworksheetoleobject) | **Post** /cells/{name}/worksheets/{sheetName}/oleobjects/{oleObjectIndex} | Update OLE object.
*CellsApi* | [**CellsOleObjectsPutWorksheetOleObject**](docs/CellsApi.md#cellsoleobjectsputworksheetoleobject) | **Put** /cells/{name}/worksheets/{sheetName}/oleobjects | Add OLE object
*CellsApi* | [**CellsPageBreaksDeleteHorizontalPageBreak**](docs/CellsApi.md#cellspagebreaksdeletehorizontalpagebreak) | **Delete** /cells/{name}/worksheets/{sheetName}/horizontalpagebreaks/{index} | 
*CellsApi* | [**CellsPageBreaksDeleteHorizontalPageBreaks**](docs/CellsApi.md#cellspagebreaksdeletehorizontalpagebreaks) | **Delete** /cells/{name}/worksheets/{sheetName}/horizontalpagebreaks | 
*CellsApi* | [**CellsPageBreaksDeleteVerticalPageBreak**](docs/CellsApi.md#cellspagebreaksdeleteverticalpagebreak) | **Delete** /cells/{name}/worksheets/{sheetName}/verticalpagebreaks/{index} | 
*CellsApi* | [**CellsPageBreaksDeleteVerticalPageBreaks**](docs/CellsApi.md#cellspagebreaksdeleteverticalpagebreaks) | **Delete** /cells/{name}/worksheets/{sheetName}/verticalpagebreaks | 
*CellsApi* | [**CellsPageBreaksGetHorizontalPageBreak**](docs/CellsApi.md#cellspagebreaksgethorizontalpagebreak) | **Get** /cells/{name}/worksheets/{sheetName}/horizontalpagebreaks/{index} | 
*CellsApi* | [**CellsPageBreaksGetHorizontalPageBreaks**](docs/CellsApi.md#cellspagebreaksgethorizontalpagebreaks) | **Get** /cells/{name}/worksheets/{sheetName}/horizontalpagebreaks | 
*CellsApi* | [**CellsPageBreaksGetVerticalPageBreak**](docs/CellsApi.md#cellspagebreaksgetverticalpagebreak) | **Get** /cells/{name}/worksheets/{sheetName}/verticalpagebreaks/{index} | 
*CellsApi* | [**CellsPageBreaksGetVerticalPageBreaks**](docs/CellsApi.md#cellspagebreaksgetverticalpagebreaks) | **Get** /cells/{name}/worksheets/{sheetName}/verticalpagebreaks | 
*CellsApi* | [**CellsPageBreaksPutHorizontalPageBreak**](docs/CellsApi.md#cellspagebreaksputhorizontalpagebreak) | **Put** /cells/{name}/worksheets/{sheetName}/horizontalpagebreaks | 
*CellsApi* | [**CellsPageBreaksPutVerticalPageBreak**](docs/CellsApi.md#cellspagebreaksputverticalpagebreak) | **Put** /cells/{name}/worksheets/{sheetName}/verticalpagebreaks | 
*CellsApi* | [**CellsPageSetupDeleteHeaderFooter**](docs/CellsApi.md#cellspagesetupdeleteheaderfooter) | **Delete** /cells/{name}/worksheets/{sheetName}/pagesetup/clearheaderfooter | clear header footer
*CellsApi* | [**CellsPageSetupGetFooter**](docs/CellsApi.md#cellspagesetupgetfooter) | **Get** /cells/{name}/worksheets/{sheetName}/pagesetup/footer | get page footer information
*CellsApi* | [**CellsPageSetupGetHeader**](docs/CellsApi.md#cellspagesetupgetheader) | **Get** /cells/{name}/worksheets/{sheetName}/pagesetup/header | get page header information
*CellsApi* | [**CellsPageSetupGetPageSetup**](docs/CellsApi.md#cellspagesetupgetpagesetup) | **Get** /cells/{name}/worksheets/{sheetName}/pagesetup | Get Page Setup information.             
*CellsApi* | [**CellsPageSetupPostFooter**](docs/CellsApi.md#cellspagesetuppostfooter) | **Post** /cells/{name}/worksheets/{sheetName}/pagesetup/footer | update  page footer information 
*CellsApi* | [**CellsPageSetupPostHeader**](docs/CellsApi.md#cellspagesetuppostheader) | **Post** /cells/{name}/worksheets/{sheetName}/pagesetup/header | update  page header information 
*CellsApi* | [**CellsPageSetupPostPageSetup**](docs/CellsApi.md#cellspagesetuppostpagesetup) | **Post** /cells/{name}/worksheets/{sheetName}/pagesetup | Update Page Setup information.
*CellsApi* | [**CellsPicturesDeleteWorksheetPicture**](docs/CellsApi.md#cellspicturesdeleteworksheetpicture) | **Delete** /cells/{name}/worksheets/{sheetName}/pictures/{pictureIndex} | Delete a picture object in worksheet
*CellsApi* | [**CellsPicturesDeleteWorksheetPictures**](docs/CellsApi.md#cellspicturesdeleteworksheetpictures) | **Delete** /cells/{name}/worksheets/{sheetName}/pictures | Delete all pictures in worksheet.
*CellsApi* | [**CellsPicturesGetWorksheetPicture**](docs/CellsApi.md#cellspicturesgetworksheetpicture) | **Get** /cells/{name}/worksheets/{sheetName}/pictures/{pictureIndex} | GRead worksheet picture by number.
*CellsApi* | [**CellsPicturesGetWorksheetPictures**](docs/CellsApi.md#cellspicturesgetworksheetpictures) | **Get** /cells/{name}/worksheets/{sheetName}/pictures | Read worksheet pictures.
*CellsApi* | [**CellsPicturesPostWorksheetPicture**](docs/CellsApi.md#cellspicturespostworksheetpicture) | **Post** /cells/{name}/worksheets/{sheetName}/pictures/{pictureIndex} | Update worksheet picture by index.
*CellsApi* | [**CellsPicturesPutWorksheetAddPicture**](docs/CellsApi.md#cellspicturesputworksheetaddpicture) | **Put** /cells/{name}/worksheets/{sheetName}/pictures | Add a new worksheet picture.
*CellsApi* | [**CellsPivotTablesDeletePivotTableField**](docs/CellsApi.md#cellspivottablesdeletepivottablefield) | **Delete** /cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/PivotField | Delete pivot field into into pivot table
*CellsApi* | [**CellsPivotTablesDeleteWorksheetPivotTable**](docs/CellsApi.md#cellspivottablesdeleteworksheetpivottable) | **Delete** /cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex} | Delete worksheet pivot table by index
*CellsApi* | [**CellsPivotTablesDeleteWorksheetPivotTableFilter**](docs/CellsApi.md#cellspivottablesdeleteworksheetpivottablefilter) | **Delete** /cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/PivotFilters/{fieldIndex} | delete  pivot filter for piovt table             
*CellsApi* | [**CellsPivotTablesDeleteWorksheetPivotTableFilters**](docs/CellsApi.md#cellspivottablesdeleteworksheetpivottablefilters) | **Delete** /cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/PivotFilters | delete all pivot filters for piovt table
*CellsApi* | [**CellsPivotTablesDeleteWorksheetPivotTables**](docs/CellsApi.md#cellspivottablesdeleteworksheetpivottables) | **Delete** /cells/{name}/worksheets/{sheetName}/pivottables | Delete worksheet pivot tables
*CellsApi* | [**CellsPivotTablesGetPivotTableField**](docs/CellsApi.md#cellspivottablesgetpivottablefield) | **Get** /cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/PivotField | Get pivot field into into pivot table
*CellsApi* | [**CellsPivotTablesGetWorksheetPivotTable**](docs/CellsApi.md#cellspivottablesgetworksheetpivottable) | **Get** /cells/{name}/worksheets/{sheetName}/pivottables/{pivottableIndex} | Get worksheet pivottable info by index.
*CellsApi* | [**CellsPivotTablesGetWorksheetPivotTableFilter**](docs/CellsApi.md#cellspivottablesgetworksheetpivottablefilter) | **Get** /cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/PivotFilters/{filterIndex} | 
*CellsApi* | [**CellsPivotTablesGetWorksheetPivotTableFilters**](docs/CellsApi.md#cellspivottablesgetworksheetpivottablefilters) | **Get** /cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/PivotFilters | 
*CellsApi* | [**CellsPivotTablesGetWorksheetPivotTables**](docs/CellsApi.md#cellspivottablesgetworksheetpivottables) | **Get** /cells/{name}/worksheets/{sheetName}/pivottables | Get worksheet pivottables info.
*CellsApi* | [**CellsPivotTablesPostPivotTableCellStyle**](docs/CellsApi.md#cellspivottablespostpivottablecellstyle) | **Post** /cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/Format | Update cell style for pivot table
*CellsApi* | [**CellsPivotTablesPostPivotTableFieldHideItem**](docs/CellsApi.md#cellspivottablespostpivottablefieldhideitem) | **Post** /cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/PivotField/Hide | 
*CellsApi* | [**CellsPivotTablesPostPivotTableFieldMoveTo**](docs/CellsApi.md#cellspivottablespostpivottablefieldmoveto) | **Post** /cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/PivotField/Move | 
*CellsApi* | [**CellsPivotTablesPostPivotTableStyle**](docs/CellsApi.md#cellspivottablespostpivottablestyle) | **Post** /cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/FormatAll | Update style for pivot table
*CellsApi* | [**CellsPivotTablesPostWorksheetPivotTableCalculate**](docs/CellsApi.md#cellspivottablespostworksheetpivottablecalculate) | **Post** /cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/Calculate | Calculates pivottable&#39;s data to cells.
*CellsApi* | [**CellsPivotTablesPostWorksheetPivotTableMove**](docs/CellsApi.md#cellspivottablespostworksheetpivottablemove) | **Post** /cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/Move | 
*CellsApi* | [**CellsPivotTablesPutPivotTableField**](docs/CellsApi.md#cellspivottablesputpivottablefield) | **Put** /cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/PivotField | Add pivot field into into pivot table
*CellsApi* | [**CellsPivotTablesPutWorksheetPivotTable**](docs/CellsApi.md#cellspivottablesputworksheetpivottable) | **Put** /cells/{name}/worksheets/{sheetName}/pivottables | Add a pivot table into worksheet.
*CellsApi* | [**CellsPivotTablesPutWorksheetPivotTableFilter**](docs/CellsApi.md#cellspivottablesputworksheetpivottablefilter) | **Put** /cells/{name}/worksheets/{sheetName}/pivottables/{pivotTableIndex}/PivotFilters | Add pivot filter for piovt table index
*CellsApi* | [**CellsPostCellCalculate**](docs/CellsApi.md#cellspostcellcalculate) | **Post** /cells/{name}/worksheets/{sheetName}/cells/{cellName}/calculate | Cell calculate formula
*CellsApi* | [**CellsPostCellCharacters**](docs/CellsApi.md#cellspostcellcharacters) | **Post** /cells/{name}/worksheets/{sheetName}/cells/{cellName}/characters | Set cell characters 
*CellsApi* | [**CellsPostClearContents**](docs/CellsApi.md#cellspostclearcontents) | **Post** /cells/{name}/worksheets/{sheetName}/cells/clearcontents | Clear cells contents.
*CellsApi* | [**CellsPostClearFormats**](docs/CellsApi.md#cellspostclearformats) | **Post** /cells/{name}/worksheets/{sheetName}/cells/clearformats | Clear cells contents.
*CellsApi* | [**CellsPostColumnStyle**](docs/CellsApi.md#cellspostcolumnstyle) | **Post** /cells/{name}/worksheets/{sheetName}/cells/columns/{columnIndex}/style | Set column style
*CellsApi* | [**CellsPostCopyCellIntoCell**](docs/CellsApi.md#cellspostcopycellintocell) | **Post** /cells/{name}/worksheets/{sheetName}/cells/{destCellName}/copy | Copy cell into cell
*CellsApi* | [**CellsPostCopyWorksheetColumns**](docs/CellsApi.md#cellspostcopyworksheetcolumns) | **Post** /cells/{name}/worksheets/{sheetName}/cells/columns/copy | Copy worksheet columns.
*CellsApi* | [**CellsPostCopyWorksheetRows**](docs/CellsApi.md#cellspostcopyworksheetrows) | **Post** /cells/{name}/worksheets/{sheetName}/cells/rows/copy | Copy worksheet rows.
*CellsApi* | [**CellsPostGroupWorksheetColumns**](docs/CellsApi.md#cellspostgroupworksheetcolumns) | **Post** /cells/{name}/worksheets/{sheetName}/cells/columns/group | Group worksheet columns.
*CellsApi* | [**CellsPostGroupWorksheetRows**](docs/CellsApi.md#cellspostgroupworksheetrows) | **Post** /cells/{name}/worksheets/{sheetName}/cells/rows/group | Group worksheet rows.
*CellsApi* | [**CellsPostHideWorksheetColumns**](docs/CellsApi.md#cellsposthideworksheetcolumns) | **Post** /cells/{name}/worksheets/{sheetName}/cells/columns/hide | Hide worksheet columns.
*CellsApi* | [**CellsPostHideWorksheetRows**](docs/CellsApi.md#cellsposthideworksheetrows) | **Post** /cells/{name}/worksheets/{sheetName}/cells/rows/hide | Hide worksheet rows.
*CellsApi* | [**CellsPostRowStyle**](docs/CellsApi.md#cellspostrowstyle) | **Post** /cells/{name}/worksheets/{sheetName}/cells/rows/{rowIndex}/style | Set row style.
*CellsApi* | [**CellsPostSetCellHtmlString**](docs/CellsApi.md#cellspostsetcellhtmlstring) | **Post** /cells/{name}/worksheets/{sheetName}/cells/{cellName}/htmlstring | Set htmlstring value into cell
*CellsApi* | [**CellsPostSetCellRangeValue**](docs/CellsApi.md#cellspostsetcellrangevalue) | **Post** /cells/{name}/worksheets/{sheetName}/cells | Set cell range value 
*CellsApi* | [**CellsPostSetWorksheetColumnWidth**](docs/CellsApi.md#cellspostsetworksheetcolumnwidth) | **Post** /cells/{name}/worksheets/{sheetName}/cells/columns/{columnIndex} | Set worksheet column width.
*CellsApi* | [**CellsPostUngroupWorksheetColumns**](docs/CellsApi.md#cellspostungroupworksheetcolumns) | **Post** /cells/{name}/worksheets/{sheetName}/cells/columns/ungroup | Ungroup worksheet columns.
*CellsApi* | [**CellsPostUngroupWorksheetRows**](docs/CellsApi.md#cellspostungroupworksheetrows) | **Post** /cells/{name}/worksheets/{sheetName}/cells/rows/ungroup | Ungroup worksheet rows.
*CellsApi* | [**CellsPostUnhideWorksheetColumns**](docs/CellsApi.md#cellspostunhideworksheetcolumns) | **Post** /cells/{name}/worksheets/{sheetName}/cells/columns/unhide | Unhide worksheet columns.
*CellsApi* | [**CellsPostUnhideWorksheetRows**](docs/CellsApi.md#cellspostunhideworksheetrows) | **Post** /cells/{name}/worksheets/{sheetName}/cells/rows/unhide | Unhide worksheet rows.
*CellsApi* | [**CellsPostUpdateWorksheetCellStyle**](docs/CellsApi.md#cellspostupdateworksheetcellstyle) | **Post** /cells/{name}/worksheets/{sheetName}/cells/{cellName}/style | Update cell&#39;s style.
*CellsApi* | [**CellsPostUpdateWorksheetRangeStyle**](docs/CellsApi.md#cellspostupdateworksheetrangestyle) | **Post** /cells/{name}/worksheets/{sheetName}/cells/style | Update cell&#39;s range style.
*CellsApi* | [**CellsPostUpdateWorksheetRow**](docs/CellsApi.md#cellspostupdateworksheetrow) | **Post** /cells/{name}/worksheets/{sheetName}/cells/rows/{rowIndex} | Update worksheet row.
*CellsApi* | [**CellsPostWorksheetCellSetValue**](docs/CellsApi.md#cellspostworksheetcellsetvalue) | **Post** /cells/{name}/worksheets/{sheetName}/cells/{cellName} | Set cell value.
*CellsApi* | [**CellsPostWorksheetMerge**](docs/CellsApi.md#cellspostworksheetmerge) | **Post** /cells/{name}/worksheets/{sheetName}/cells/merge | Merge cells.
*CellsApi* | [**CellsPostWorksheetUnmerge**](docs/CellsApi.md#cellspostworksheetunmerge) | **Post** /cells/{name}/worksheets/{sheetName}/cells/unmerge | Unmerge cells.
*CellsApi* | [**CellsPropertiesDeleteDocumentProperties**](docs/CellsApi.md#cellspropertiesdeletedocumentproperties) | **Delete** /cells/{name}/documentproperties | Delete all custom document properties and clean built-in ones.
*CellsApi* | [**CellsPropertiesDeleteDocumentProperty**](docs/CellsApi.md#cellspropertiesdeletedocumentproperty) | **Delete** /cells/{name}/documentproperties/{propertyName} | Delete document property.
*CellsApi* | [**CellsPropertiesGetDocumentProperties**](docs/CellsApi.md#cellspropertiesgetdocumentproperties) | **Get** /cells/{name}/documentproperties | Read document properties.
*CellsApi* | [**CellsPropertiesGetDocumentProperty**](docs/CellsApi.md#cellspropertiesgetdocumentproperty) | **Get** /cells/{name}/documentproperties/{propertyName} | Read document property by name.
*CellsApi* | [**CellsPropertiesPutDocumentProperty**](docs/CellsApi.md#cellspropertiesputdocumentproperty) | **Put** /cells/{name}/documentproperties/{propertyName} | Set/create document property.
*CellsApi* | [**CellsPutInsertWorksheetColumns**](docs/CellsApi.md#cellsputinsertworksheetcolumns) | **Put** /cells/{name}/worksheets/{sheetName}/cells/columns/{columnIndex} | Insert worksheet columns.
*CellsApi* | [**CellsPutInsertWorksheetRow**](docs/CellsApi.md#cellsputinsertworksheetrow) | **Put** /cells/{name}/worksheets/{sheetName}/cells/rows/{rowIndex} | Insert new worksheet row.
*CellsApi* | [**CellsPutInsertWorksheetRows**](docs/CellsApi.md#cellsputinsertworksheetrows) | **Put** /cells/{name}/worksheets/{sheetName}/cells/rows | Insert several new worksheet rows.
*CellsApi* | [**CellsRangesGetWorksheetCellsRangeValue**](docs/CellsApi.md#cellsrangesgetworksheetcellsrangevalue) | **Get** /cells/{name}/worksheets/{sheetName}/ranges/value | Get cells list in a range by range name or row column indexes  
*CellsApi* | [**CellsRangesPostWorksheetCellsRangeColumnWidth**](docs/CellsApi.md#cellsrangespostworksheetcellsrangecolumnwidth) | **Post** /cells/{name}/worksheets/{sheetName}/ranges/columnWidth | Set column width of range
*CellsApi* | [**CellsRangesPostWorksheetCellsRangeMerge**](docs/CellsApi.md#cellsrangespostworksheetcellsrangemerge) | **Post** /cells/{name}/worksheets/{sheetName}/ranges/merge | Combines a range of cells into a single cell.              
*CellsApi* | [**CellsRangesPostWorksheetCellsRangeMoveTo**](docs/CellsApi.md#cellsrangespostworksheetcellsrangemoveto) | **Post** /cells/{name}/worksheets/{sheetName}/ranges/moveto | Move the current range to the dest range.             
*CellsApi* | [**CellsRangesPostWorksheetCellsRangeOutlineBorder**](docs/CellsApi.md#cellsrangespostworksheetcellsrangeoutlineborder) | **Post** /cells/{name}/worksheets/{sheetName}/ranges/outlineBorder | Sets outline border around a range of cells.
*CellsApi* | [**CellsRangesPostWorksheetCellsRangeRowHeight**](docs/CellsApi.md#cellsrangespostworksheetcellsrangerowheight) | **Post** /cells/{name}/worksheets/{sheetName}/ranges/rowHeight | set row height of range
*CellsApi* | [**CellsRangesPostWorksheetCellsRangeStyle**](docs/CellsApi.md#cellsrangespostworksheetcellsrangestyle) | **Post** /cells/{name}/worksheets/{sheetName}/ranges/style | Sets the style of the range.             
*CellsApi* | [**CellsRangesPostWorksheetCellsRangeUnmerge**](docs/CellsApi.md#cellsrangespostworksheetcellsrangeunmerge) | **Post** /cells/{name}/worksheets/{sheetName}/ranges/unmerge | Unmerges merged cells of this range.             
*CellsApi* | [**CellsRangesPostWorksheetCellsRangeValue**](docs/CellsApi.md#cellsrangespostworksheetcellsrangevalue) | **Post** /cells/{name}/worksheets/{sheetName}/ranges/value | Puts a value into the range, if appropriate the value will be converted to other data type and cell&#39;s number format will be reset.             
*CellsApi* | [**CellsRangesPostWorksheetCellsRanges**](docs/CellsApi.md#cellsrangespostworksheetcellsranges) | **Post** /cells/{name}/worksheets/{sheetName}/ranges | copy range in the worksheet
*CellsApi* | [**CellsSaveAsPostDocumentSaveAs**](docs/CellsApi.md#cellssaveaspostdocumentsaveas) | **Post** /cells/{name}/SaveAs | Convert document and save result to storage.
*CellsApi* | [**CellsShapesDeleteWorksheetShape**](docs/CellsApi.md#cellsshapesdeleteworksheetshape) | **Delete** /cells/{name}/worksheets/{sheetName}/shapes/{shapeindex} | Delete a shape in worksheet
*CellsApi* | [**CellsShapesDeleteWorksheetShapes**](docs/CellsApi.md#cellsshapesdeleteworksheetshapes) | **Delete** /cells/{name}/worksheets/{sheetName}/shapes | delete all shapes in worksheet
*CellsApi* | [**CellsShapesGetWorksheetShape**](docs/CellsApi.md#cellsshapesgetworksheetshape) | **Get** /cells/{name}/worksheets/{sheetName}/shapes/{shapeindex} | Get worksheet shape
*CellsApi* | [**CellsShapesGetWorksheetShapes**](docs/CellsApi.md#cellsshapesgetworksheetshapes) | **Get** /cells/{name}/worksheets/{sheetName}/shapes | Get worksheet shapes 
*CellsApi* | [**CellsShapesPostWorksheetShape**](docs/CellsApi.md#cellsshapespostworksheetshape) | **Post** /cells/{name}/worksheets/{sheetName}/shapes/{shapeindex} | Update a shape in worksheet
*CellsApi* | [**CellsShapesPutWorksheetShape**](docs/CellsApi.md#cellsshapesputworksheetshape) | **Put** /cells/{name}/worksheets/{sheetName}/shapes | Add shape in worksheet
*CellsApi* | [**CellsTaskPostRunTask**](docs/CellsApi.md#cellstaskpostruntask) | **Post** /cells/task/runtask | Run tasks  
*CellsApi* | [**CellsWorkbookDeleteDecryptDocument**](docs/CellsApi.md#cellsworkbookdeletedecryptdocument) | **Delete** /cells/{name}/encryption | Decrypt document.
*CellsApi* | [**CellsWorkbookDeleteDocumentUnprotectFromChanges**](docs/CellsApi.md#cellsworkbookdeletedocumentunprotectfromchanges) | **Delete** /cells/{name}/writeProtection | Unprotect document from changes.
*CellsApi* | [**CellsWorkbookDeleteUnprotectDocument**](docs/CellsApi.md#cellsworkbookdeleteunprotectdocument) | **Delete** /cells/{name}/protection | Unprotect document.
*CellsApi* | [**CellsWorkbookDeleteWorkbookName**](docs/CellsApi.md#cellsworkbookdeleteworkbookname) | **Delete** /cells/{name}/names/{nameName} | Clean workbook&#39;s names.
*CellsApi* | [**CellsWorkbookDeleteWorkbookNames**](docs/CellsApi.md#cellsworkbookdeleteworkbooknames) | **Delete** /cells/{name}/names | Clean workbook&#39;s names.
*CellsApi* | [**CellsWorkbookGetWorkbook**](docs/CellsApi.md#cellsworkbookgetworkbook) | **Get** /cells/{name} | Read workbook info or export.
*CellsApi* | [**CellsWorkbookGetWorkbookDefaultStyle**](docs/CellsApi.md#cellsworkbookgetworkbookdefaultstyle) | **Get** /cells/{name}/defaultstyle | Read workbook default style info.
*CellsApi* | [**CellsWorkbookGetWorkbookName**](docs/CellsApi.md#cellsworkbookgetworkbookname) | **Get** /cells/{name}/names/{nameName} | Read workbook&#39;s name.
*CellsApi* | [**CellsWorkbookGetWorkbookNameValue**](docs/CellsApi.md#cellsworkbookgetworkbooknamevalue) | **Get** /cells/{name}/names/{nameName}/value | Get workbook&#39;s name value.
*CellsApi* | [**CellsWorkbookGetWorkbookNames**](docs/CellsApi.md#cellsworkbookgetworkbooknames) | **Get** /cells/{name}/names | Read workbook&#39;s names.
*CellsApi* | [**CellsWorkbookGetWorkbookSettings**](docs/CellsApi.md#cellsworkbookgetworkbooksettings) | **Get** /cells/{name}/settings | Get Workbook Settings DTO
*CellsApi* | [**CellsWorkbookGetWorkbookTextItems**](docs/CellsApi.md#cellsworkbookgetworkbooktextitems) | **Get** /cells/{name}/textItems | Read workbook&#39;s text items.
*CellsApi* | [**CellsWorkbookPostAutofitWorkbookRows**](docs/CellsApi.md#cellsworkbookpostautofitworkbookrows) | **Post** /cells/{name}/autofitrows | Autofit workbook rows.
*CellsApi* | [**CellsWorkbookPostEncryptDocument**](docs/CellsApi.md#cellsworkbookpostencryptdocument) | **Post** /cells/{name}/encryption | Encript document.
*CellsApi* | [**CellsWorkbookPostImportData**](docs/CellsApi.md#cellsworkbookpostimportdata) | **Post** /cells/{name}/importdata | 
*CellsApi* | [**CellsWorkbookPostProtectDocument**](docs/CellsApi.md#cellsworkbookpostprotectdocument) | **Post** /cells/{name}/protection | Protect document.
*CellsApi* | [**CellsWorkbookPostWorkbookCalculateFormula**](docs/CellsApi.md#cellsworkbookpostworkbookcalculateformula) | **Post** /cells/{name}/calculateformula | Calculate all formulas in workbook.
*CellsApi* | [**CellsWorkbookPostWorkbookGetSmartMarkerResult**](docs/CellsApi.md#cellsworkbookpostworkbookgetsmartmarkerresult) | **Post** /cells/{name}/smartmarker | Smart marker processing result.
*CellsApi* | [**CellsWorkbookPostWorkbookSettings**](docs/CellsApi.md#cellsworkbookpostworkbooksettings) | **Post** /cells/{name}/settings | Update Workbook setting 
*CellsApi* | [**CellsWorkbookPostWorkbookSplit**](docs/CellsApi.md#cellsworkbookpostworkbooksplit) | **Post** /cells/{name}/split | Split workbook.
*CellsApi* | [**CellsWorkbookPostWorkbooksMerge**](docs/CellsApi.md#cellsworkbookpostworkbooksmerge) | **Post** /cells/{name}/merge | Merge workbooks.
*CellsApi* | [**CellsWorkbookPostWorkbooksTextReplace**](docs/CellsApi.md#cellsworkbookpostworkbookstextreplace) | **Post** /cells/{name}/replaceText | Replace text.
*CellsApi* | [**CellsWorkbookPostWorkbooksTextSearch**](docs/CellsApi.md#cellsworkbookpostworkbookstextsearch) | **Post** /cells/{name}/findText | Search text.
*CellsApi* | [**CellsWorkbookPutConvertWorkbook**](docs/CellsApi.md#cellsworkbookputconvertworkbook) | **Put** /cells/convert | Convert workbook from request content to some format.
*CellsApi* | [**CellsWorkbookPutDocumentProtectFromChanges**](docs/CellsApi.md#cellsworkbookputdocumentprotectfromchanges) | **Put** /cells/{name}/writeProtection | Protect document from changes.
*CellsApi* | [**CellsWorkbookPutWorkbookCreate**](docs/CellsApi.md#cellsworkbookputworkbookcreate) | **Put** /cells/{name} | Create new workbook using deferent methods.
*CellsApi* | [**CellsWorksheetValidationsDeleteWorksheetValidation**](docs/CellsApi.md#cellsworksheetvalidationsdeleteworksheetvalidation) | **Delete** /cells/{name}/worksheets/{sheetName}/validations/{validationIndex} | Delete worksheet validation by index.
*CellsApi* | [**CellsWorksheetValidationsDeleteWorksheetValidations**](docs/CellsApi.md#cellsworksheetvalidationsdeleteworksheetvalidations) | **Delete** /cells/{name}/worksheets/{sheetName}/validations | Clear all validation in worksheet.
*CellsApi* | [**CellsWorksheetValidationsGetWorksheetValidation**](docs/CellsApi.md#cellsworksheetvalidationsgetworksheetvalidation) | **Get** /cells/{name}/worksheets/{sheetName}/validations/{validationIndex} | Get worksheet validation by index.
*CellsApi* | [**CellsWorksheetValidationsGetWorksheetValidations**](docs/CellsApi.md#cellsworksheetvalidationsgetworksheetvalidations) | **Get** /cells/{name}/worksheets/{sheetName}/validations | Get worksheet validations.
*CellsApi* | [**CellsWorksheetValidationsPostWorksheetValidation**](docs/CellsApi.md#cellsworksheetvalidationspostworksheetvalidation) | **Post** /cells/{name}/worksheets/{sheetName}/validations/{validationIndex} | Update worksheet validation by index.
*CellsApi* | [**CellsWorksheetValidationsPutWorksheetValidation**](docs/CellsApi.md#cellsworksheetvalidationsputworksheetvalidation) | **Put** /cells/{name}/worksheets/{sheetName}/validations | Add worksheet validation at index.
*CellsApi* | [**CellsWorksheetsDeleteUnprotectWorksheet**](docs/CellsApi.md#cellsworksheetsdeleteunprotectworksheet) | **Delete** /cells/{name}/worksheets/{sheetName}/protection | Unprotect worksheet.
*CellsApi* | [**CellsWorksheetsDeleteWorksheet**](docs/CellsApi.md#cellsworksheetsdeleteworksheet) | **Delete** /cells/{name}/worksheets/{sheetName} | Delete worksheet.
*CellsApi* | [**CellsWorksheetsDeleteWorksheetBackground**](docs/CellsApi.md#cellsworksheetsdeleteworksheetbackground) | **Delete** /cells/{name}/worksheets/{sheetName}/background | Set worksheet background image.
*CellsApi* | [**CellsWorksheetsDeleteWorksheetComment**](docs/CellsApi.md#cellsworksheetsdeleteworksheetcomment) | **Delete** /cells/{name}/worksheets/{sheetName}/comments/{cellName} | Delete worksheet&#39;s cell comment.
*CellsApi* | [**CellsWorksheetsDeleteWorksheetComments**](docs/CellsApi.md#cellsworksheetsdeleteworksheetcomments) | **Delete** /cells/{name}/worksheets/{sheetName}/comments | Delete all comments for worksheet.
*CellsApi* | [**CellsWorksheetsDeleteWorksheetFreezePanes**](docs/CellsApi.md#cellsworksheetsdeleteworksheetfreezepanes) | **Delete** /cells/{name}/worksheets/{sheetName}/freezepanes | Unfreeze panes
*CellsApi* | [**CellsWorksheetsGetNamedRanges**](docs/CellsApi.md#cellsworksheetsgetnamedranges) | **Get** /cells/{name}/worksheets/ranges | Read worksheets ranges info.
*CellsApi* | [**CellsWorksheetsGetWorksheet**](docs/CellsApi.md#cellsworksheetsgetworksheet) | **Get** /cells/{name}/worksheets/{sheetName} | Read worksheet info or export.
*CellsApi* | [**CellsWorksheetsGetWorksheetCalculateFormula**](docs/CellsApi.md#cellsworksheetsgetworksheetcalculateformula) | **Get** /cells/{name}/worksheets/{sheetName}/formulaResult | Calculate formula value.
*CellsApi* | [**CellsWorksheetsGetWorksheetComment**](docs/CellsApi.md#cellsworksheetsgetworksheetcomment) | **Get** /cells/{name}/worksheets/{sheetName}/comments/{cellName} | Get worksheet comment by cell name.
*CellsApi* | [**CellsWorksheetsGetWorksheetComments**](docs/CellsApi.md#cellsworksheetsgetworksheetcomments) | **Get** /cells/{name}/worksheets/{sheetName}/comments | Get worksheet comments.
*CellsApi* | [**CellsWorksheetsGetWorksheetMergedCell**](docs/CellsApi.md#cellsworksheetsgetworksheetmergedcell) | **Get** /cells/{name}/worksheets/{sheetName}/mergedCells/{mergedCellIndex} | Get worksheet merged cell by its index.
*CellsApi* | [**CellsWorksheetsGetWorksheetMergedCells**](docs/CellsApi.md#cellsworksheetsgetworksheetmergedcells) | **Get** /cells/{name}/worksheets/{sheetName}/mergedCells | Get worksheet merged cells.
*CellsApi* | [**CellsWorksheetsGetWorksheetTextItems**](docs/CellsApi.md#cellsworksheetsgetworksheettextitems) | **Get** /cells/{name}/worksheets/{sheetName}/textItems | Get worksheet text items.
*CellsApi* | [**CellsWorksheetsGetWorksheets**](docs/CellsApi.md#cellsworksheetsgetworksheets) | **Get** /cells/{name}/worksheets | Read worksheets info.
*CellsApi* | [**CellsWorksheetsPostAutofitWorksheetColumns**](docs/CellsApi.md#cellsworksheetspostautofitworksheetcolumns) | **Post** /cells/{name}/worksheets/{sheetName}/autofitcolumns | 
*CellsApi* | [**CellsWorksheetsPostAutofitWorksheetRow**](docs/CellsApi.md#cellsworksheetspostautofitworksheetrow) | **Post** /cells/{name}/worksheets/{sheetName}/autofitrow | 
*CellsApi* | [**CellsWorksheetsPostAutofitWorksheetRows**](docs/CellsApi.md#cellsworksheetspostautofitworksheetrows) | **Post** /cells/{name}/worksheets/{sheetName}/autofitrows | Autofit worksheet rows.
*CellsApi* | [**CellsWorksheetsPostCopyWorksheet**](docs/CellsApi.md#cellsworksheetspostcopyworksheet) | **Post** /cells/{name}/worksheets/{sheetName}/copy | 
*CellsApi* | [**CellsWorksheetsPostMoveWorksheet**](docs/CellsApi.md#cellsworksheetspostmoveworksheet) | **Post** /cells/{name}/worksheets/{sheetName}/position | Move worksheet.
*CellsApi* | [**CellsWorksheetsPostRenameWorksheet**](docs/CellsApi.md#cellsworksheetspostrenameworksheet) | **Post** /cells/{name}/worksheets/{sheetName}/rename | Rename worksheet
*CellsApi* | [**CellsWorksheetsPostUpdateWorksheetProperty**](docs/CellsApi.md#cellsworksheetspostupdateworksheetproperty) | **Post** /cells/{name}/worksheets/{sheetName} | Update worksheet property
*CellsApi* | [**CellsWorksheetsPostUpdateWorksheetZoom**](docs/CellsApi.md#cellsworksheetspostupdateworksheetzoom) | **Post** /cells/{name}/worksheets/{sheetName}/zoom | 
*CellsApi* | [**CellsWorksheetsPostWorksheetComment**](docs/CellsApi.md#cellsworksheetspostworksheetcomment) | **Post** /cells/{name}/worksheets/{sheetName}/comments/{cellName} | Update worksheet&#39;s cell comment.
*CellsApi* | [**CellsWorksheetsPostWorksheetRangeSort**](docs/CellsApi.md#cellsworksheetspostworksheetrangesort) | **Post** /cells/{name}/worksheets/{sheetName}/sort | Sort worksheet range.
*CellsApi* | [**CellsWorksheetsPostWorksheetTextSearch**](docs/CellsApi.md#cellsworksheetspostworksheettextsearch) | **Post** /cells/{name}/worksheets/{sheetName}/findText | Search text.
*CellsApi* | [**CellsWorksheetsPostWorsheetTextReplace**](docs/CellsApi.md#cellsworksheetspostworsheettextreplace) | **Post** /cells/{name}/worksheets/{sheetName}/replaceText | Replace text.
*CellsApi* | [**CellsWorksheetsPutAddNewWorksheet**](docs/CellsApi.md#cellsworksheetsputaddnewworksheet) | **Put** /cells/{name}/worksheets/{sheetName} | Add new worksheet.
*CellsApi* | [**CellsWorksheetsPutChangeVisibilityWorksheet**](docs/CellsApi.md#cellsworksheetsputchangevisibilityworksheet) | **Put** /cells/{name}/worksheets/{sheetName}/visible | Change worksheet visibility.
*CellsApi* | [**CellsWorksheetsPutProtectWorksheet**](docs/CellsApi.md#cellsworksheetsputprotectworksheet) | **Put** /cells/{name}/worksheets/{sheetName}/protection | Protect worksheet.
*CellsApi* | [**CellsWorksheetsPutWorksheetBackground**](docs/CellsApi.md#cellsworksheetsputworksheetbackground) | **Put** /cells/{name}/worksheets/{sheetName}/background | Set worksheet background image.
*CellsApi* | [**CellsWorksheetsPutWorksheetComment**](docs/CellsApi.md#cellsworksheetsputworksheetcomment) | **Put** /cells/{name}/worksheets/{sheetName}/comments/{cellName} | Add worksheet&#39;s cell comment.
*CellsApi* | [**CellsWorksheetsPutWorksheetFreezePanes**](docs/CellsApi.md#cellsworksheetsputworksheetfreezepanes) | **Put** /cells/{name}/worksheets/{sheetName}/freezepanes | Set freeze panes
*CellsApi* | [**CopyFile**](docs/CellsApi.md#copyfile) | **Put** /cells/storage/file/copy/{srcPath} | Copy file
*CellsApi* | [**CopyFolder**](docs/CellsApi.md#copyfolder) | **Put** /cells/storage/folder/copy/{srcPath} | Copy folder
*CellsApi* | [**CreateFolder**](docs/CellsApi.md#createfolder) | **Put** /cells/storage/folder/{path} | Create the folder
*CellsApi* | [**DeleteFile**](docs/CellsApi.md#deletefile) | **Delete** /cells/storage/file/{path} | Delete file
*CellsApi* | [**DeleteFolder**](docs/CellsApi.md#deletefolder) | **Delete** /cells/storage/folder/{path} | Delete folder
*CellsApi* | [**DownloadFile**](docs/CellsApi.md#downloadfile) | **Get** /cells/storage/file/{path} | Download file
*CellsApi* | [**GetDiscUsage**](docs/CellsApi.md#getdiscusage) | **Get** /cells/storage/disc | Get disc usage
*CellsApi* | [**GetFileVersions**](docs/CellsApi.md#getfileversions) | **Get** /cells/storage/version/{path} | Get file versions
*CellsApi* | [**GetFilesList**](docs/CellsApi.md#getfileslist) | **Get** /cells/storage/folder/{path} | Get all files and folders within a folder
*CellsApi* | [**MoveFile**](docs/CellsApi.md#movefile) | **Put** /cells/storage/file/move/{srcPath} | Move file
*CellsApi* | [**MoveFolder**](docs/CellsApi.md#movefolder) | **Put** /cells/storage/folder/move/{srcPath} | Move folder
*CellsApi* | [**OAuthPost**](docs/CellsApi.md#oauthpost) | **Post** /connect/token | Get Access token
*CellsApi* | [**ObjectExists**](docs/CellsApi.md#objectexists) | **Get** /cells/storage/exist/{path} | Check if file or folder exists
*CellsApi* | [**StorageExists**](docs/CellsApi.md#storageexists) | **Get** /cells/storage/{storageName}/exist | Check if storage exists
*CellsApi* | [**UploadFile**](docs/CellsApi.md#uploadfile) | **Put** /cells/storage/file/{path} | Upload file


## Documentation For Models

 - [AboveAverage](docs/AboveAverage.md)
 - [AccessTokenResponse](docs/AccessTokenResponse.md)
 - [Area](docs/Area.md)
 - [AutoFitterOptions](docs/AutoFitterOptions.md)
 - [Border](docs/Border.md)
 - [CalculationOptions](docs/CalculationOptions.md)
 - [CellArea](docs/CellArea.md)
 - [CellValue](docs/CellValue.md)
 - [CellsCloudResponse](docs/CellsCloudResponse.md)
 - [CellsColor](docs/CellsColor.md)
 - [Color](docs/Color.md)
 - [ColorFilter](docs/ColorFilter.md)
 - [ColorFilterRequest](docs/ColorFilterRequest.md)
 - [ColorScale](docs/ColorScale.md)
 - [ConditionalFormattingIcon](docs/ConditionalFormattingIcon.md)
 - [ConditionalFormattingValue](docs/ConditionalFormattingValue.md)
 - [CopyOptions](docs/CopyOptions.md)
 - [CreatePivotTableRequest](docs/CreatePivotTableRequest.md)
 - [CustomFilter](docs/CustomFilter.md)
 - [CustomParserConfig](docs/CustomParserConfig.md)
 - [DataBar](docs/DataBar.md)
 - [DataBarBorder](docs/DataBarBorder.md)
 - [DataSorter](docs/DataSorter.md)
 - [DiscUsage](docs/DiscUsage.md)
 - [DynamicFilter](docs/DynamicFilter.md)
 - [ErrorDetails](docs/ErrorDetails.md)
 - [FileSource](docs/FileSource.md)
 - [FileVersions](docs/FileVersions.md)
 - [FilesList](docs/FilesList.md)
 - [FilesUploadResult](docs/FilesUploadResult.md)
 - [FillFormat](docs/FillFormat.md)
 - [FilterColumn](docs/FilterColumn.md)
 - [Font](docs/Font.md)
 - [FontSetting](docs/FontSetting.md)
 - [GradientFill](docs/GradientFill.md)
 - [GradientFillStop](docs/GradientFillStop.md)
 - [HorizontalPageBreak](docs/HorizontalPageBreak.md)
 - [IconFilter](docs/IconFilter.md)
 - [IconSet](docs/IconSet.md)
 - [ImportOption](docs/ImportOption.md)
 - [Line](docs/Line.md)
 - [Link](docs/Link.md)
 - [LinkElement](docs/LinkElement.md)
 - [ListColumn](docs/ListColumn.md)
 - [ModelError](docs/ModelError.md)
 - [ModelRange](docs/ModelRange.md)
 - [MultipleFilter](docs/MultipleFilter.md)
 - [MultipleFilters](docs/MultipleFilters.md)
 - [NegativeBarFormat](docs/NegativeBarFormat.md)
 - [ObjectExist](docs/ObjectExist.md)
 - [OperateObject](docs/OperateObject.md)
 - [OperateObjectPosition](docs/OperateObjectPosition.md)
 - [OperateParameter](docs/OperateParameter.md)
 - [PageSection](docs/PageSection.md)
 - [PasswordRequest](docs/PasswordRequest.md)
 - [PasteOptions](docs/PasteOptions.md)
 - [PatternFill](docs/PatternFill.md)
 - [PdfSecurityOptions](docs/PdfSecurityOptions.md)
 - [PicFormatOption](docs/PicFormatOption.md)
 - [PivotField](docs/PivotField.md)
 - [PivotFilter](docs/PivotFilter.md)
 - [PivotItem](docs/PivotItem.md)
 - [PivotTableFieldRequest](docs/PivotTableFieldRequest.md)
 - [ProtectSheetParameter](docs/ProtectSheetParameter.md)
 - [RangeCopyRequest](docs/RangeCopyRequest.md)
 - [RangeSetOutlineBorderRequest](docs/RangeSetOutlineBorderRequest.md)
 - [RangeSetStyleRequest](docs/RangeSetStyleRequest.md)
 - [Ranges](docs/Ranges.md)
 - [ResultDestination](docs/ResultDestination.md)
 - [SaveOptions](docs/SaveOptions.md)
 - [SaveResult](docs/SaveResult.md)
 - [ShadowEffect](docs/ShadowEffect.md)
 - [SingleValue](docs/SingleValue.md)
 - [SolidFill](docs/SolidFill.md)
 - [SortKey](docs/SortKey.md)
 - [SplitResult](docs/SplitResult.md)
 - [StorageExist](docs/StorageExist.md)
 - [StorageFile](docs/StorageFile.md)
 - [TaskData](docs/TaskData.md)
 - [TaskDescription](docs/TaskDescription.md)
 - [TaskParameter](docs/TaskParameter.md)
 - [TextureFill](docs/TextureFill.md)
 - [ThemeColor](docs/ThemeColor.md)
 - [TilePicOption](docs/TilePicOption.md)
 - [Top10](docs/Top10.md)
 - [Top10Filter](docs/Top10Filter.md)
 - [ValueType](docs/ValueType.md)
 - [VerticalPageBreak](docs/VerticalPageBreak.md)
 - [Workbook](docs/Workbook.md)
 - [WorkbookEncryptionRequest](docs/WorkbookEncryptionRequest.md)
 - [WorkbookProtectionRequest](docs/WorkbookProtectionRequest.md)
 - [WorkbookSettings](docs/WorkbookSettings.md)
 - [Worksheet](docs/Worksheet.md)
 - [WorksheetMovingRequest](docs/WorksheetMovingRequest.md)
 - [AutoFilter](docs/AutoFilter.md)
 - [AutoFilterResponse](docs/AutoFilterResponse.md)
 - [AutoShapeResponse](docs/AutoShapeResponse.md)
 - [AutoShapes](docs/AutoShapes.md)
 - [AutoShapesResponse](docs/AutoShapesResponse.md)
 - [Cell](docs/Cell.md)
 - [CellResponse](docs/CellResponse.md)
 - [Cells](docs/Cells.md)
 - [CellsDocumentProperties](docs/CellsDocumentProperties.md)
 - [CellsDocumentPropertiesResponse](docs/CellsDocumentPropertiesResponse.md)
 - [CellsDocumentProperty](docs/CellsDocumentProperty.md)
 - [CellsDocumentPropertyResponse](docs/CellsDocumentPropertyResponse.md)
 - [CellsObjectOperateTaskParameter](docs/CellsObjectOperateTaskParameter.md)
 - [CellsResponse](docs/CellsResponse.md)
 - [Chart](docs/Chart.md)
 - [ChartAreaResponse](docs/ChartAreaResponse.md)
 - [ChartFrame](docs/ChartFrame.md)
 - [ChartOperateParameter](docs/ChartOperateParameter.md)
 - [Charts](docs/Charts.md)
 - [ChartsResponse](docs/ChartsResponse.md)
 - [Column](docs/Column.md)
 - [ColumnResponse](docs/ColumnResponse.md)
 - [Columns](docs/Columns.md)
 - [ColumnsResponse](docs/ColumnsResponse.md)
 - [Comment](docs/Comment.md)
 - [CommentResponse](docs/CommentResponse.md)
 - [Comments](docs/Comments.md)
 - [CommentsResponse](docs/CommentsResponse.md)
 - [ConditionalFormatting](docs/ConditionalFormatting.md)
 - [ConditionalFormattingResponse](docs/ConditionalFormattingResponse.md)
 - [ConditionalFormattings](docs/ConditionalFormattings.md)
 - [ConditionalFormattingsResponse](docs/ConditionalFormattingsResponse.md)
 - [ConvertTaskParameter](docs/ConvertTaskParameter.md)
 - [DifSaveOptions](docs/DifSaveOptions.md)
 - [FileVersion](docs/FileVersion.md)
 - [FillFormatResponse](docs/FillFormatResponse.md)
 - [FormatCondition](docs/FormatCondition.md)
 - [HorizontalPageBreakResponse](docs/HorizontalPageBreakResponse.md)
 - [HorizontalPageBreaks](docs/HorizontalPageBreaks.md)
 - [HorizontalPageBreaksResponse](docs/HorizontalPageBreaksResponse.md)
 - [HtmlSaveOptions](docs/HtmlSaveOptions.md)
 - [Hyperlink](docs/Hyperlink.md)
 - [HyperlinkResponse](docs/HyperlinkResponse.md)
 - [Hyperlinks](docs/Hyperlinks.md)
 - [HyperlinksResponse](docs/HyperlinksResponse.md)
 - [ImageSaveOptions](docs/ImageSaveOptions.md)
 - [ImportBatchDataOption](docs/ImportBatchDataOption.md)
 - [ImportCsvDataOption](docs/ImportCsvDataOption.md)
 - [ImportDataTaskParameter](docs/ImportDataTaskParameter.md)
 - [ImportDoubleArrayOption](docs/ImportDoubleArrayOption.md)
 - [ImportIntArrayOption](docs/ImportIntArrayOption.md)
 - [ImportStringArrayOption](docs/ImportStringArrayOption.md)
 - [LegendResponse](docs/LegendResponse.md)
 - [LineFormat](docs/LineFormat.md)
 - [LineResponse](docs/LineResponse.md)
 - [ListObject](docs/ListObject.md)
 - [ListObjectOperateParameter](docs/ListObjectOperateParameter.md)
 - [ListObjectResponse](docs/ListObjectResponse.md)
 - [ListObjects](docs/ListObjects.md)
 - [ListObjectsResponse](docs/ListObjectsResponse.md)
 - [MHtmlSaveOptions](docs/MHtmlSaveOptions.md)
 - [MarkdownSaveOptions](docs/MarkdownSaveOptions.md)
 - [MergedCell](docs/MergedCell.md)
 - [MergedCellResponse](docs/MergedCellResponse.md)
 - [MergedCells](docs/MergedCells.md)
 - [MergedCellsResponse](docs/MergedCellsResponse.md)
 - [Name](docs/Name.md)
 - [NameResponse](docs/NameResponse.md)
 - [Names](docs/Names.md)
 - [NamesResponse](docs/NamesResponse.md)
 - [OdsSaveOptions](docs/OdsSaveOptions.md)
 - [OleObjectResponse](docs/OleObjectResponse.md)
 - [OleObjects](docs/OleObjects.md)
 - [OleObjectsResponse](docs/OleObjectsResponse.md)
 - [OoxmlSaveOptions](docs/OoxmlSaveOptions.md)
 - [PageBreakOperateParameter](docs/PageBreakOperateParameter.md)
 - [PageSectionsResponse](docs/PageSectionsResponse.md)
 - [PageSetup](docs/PageSetup.md)
 - [PageSetupOperateParameter](docs/PageSetupOperateParameter.md)
 - [PageSetupResponse](docs/PageSetupResponse.md)
 - [PdfSaveOptions](docs/PdfSaveOptions.md)
 - [PictureResponse](docs/PictureResponse.md)
 - [Pictures](docs/Pictures.md)
 - [PicturesResponse](docs/PicturesResponse.md)
 - [PivotFieldResponse](docs/PivotFieldResponse.md)
 - [PivotFilterResponse](docs/PivotFilterResponse.md)
 - [PivotFiltersResponse](docs/PivotFiltersResponse.md)
 - [PivotTable](docs/PivotTable.md)
 - [PivotTableOperateParameter](docs/PivotTableOperateParameter.md)
 - [PivotTableResponse](docs/PivotTableResponse.md)
 - [PivotTables](docs/PivotTables.md)
 - [PivotTablesResponse](docs/PivotTablesResponse.md)
 - [RangeValueResponse](docs/RangeValueResponse.md)
 - [RangesResponse](docs/RangesResponse.md)
 - [Row](docs/Row.md)
 - [RowResponse](docs/RowResponse.md)
 - [Rows](docs/Rows.md)
 - [RowsResponse](docs/RowsResponse.md)
 - [SaveResponse](docs/SaveResponse.md)
 - [SaveResultTaskParameter](docs/SaveResultTaskParameter.md)
 - [Shape](docs/Shape.md)
 - [ShapeOperateParameter](docs/ShapeOperateParameter.md)
 - [ShapeResponse](docs/ShapeResponse.md)
 - [Shapes](docs/Shapes.md)
 - [ShapesResponse](docs/ShapesResponse.md)
 - [SingleValueResponse](docs/SingleValueResponse.md)
 - [SmartMarkerTaskParameter](docs/SmartMarkerTaskParameter.md)
 - [SplitResultDocument](docs/SplitResultDocument.md)
 - [SplitResultResponse](docs/SplitResultResponse.md)
 - [SplitWorkbookTaskParameter](docs/SplitWorkbookTaskParameter.md)
 - [SpreadsheetMl2003SaveOptions](docs/SpreadsheetMl2003SaveOptions.md)
 - [Style](docs/Style.md)
 - [StyleResponse](docs/StyleResponse.md)
 - [SvgSaveOptions](docs/SvgSaveOptions.md)
 - [TextItem](docs/TextItem.md)
 - [TextItems](docs/TextItems.md)
 - [TextItemsResponse](docs/TextItemsResponse.md)
 - [TextOptions](docs/TextOptions.md)
 - [TitleResponse](docs/TitleResponse.md)
 - [TxtSaveOptions](docs/TxtSaveOptions.md)
 - [Validation](docs/Validation.md)
 - [ValidationResponse](docs/ValidationResponse.md)
 - [Validations](docs/Validations.md)
 - [ValidationsResponse](docs/ValidationsResponse.md)
 - [VerticalPageBreakResponse](docs/VerticalPageBreakResponse.md)
 - [VerticalPageBreaks](docs/VerticalPageBreaks.md)
 - [VerticalPageBreaksResponse](docs/VerticalPageBreaksResponse.md)
 - [WorkbookOperateParameter](docs/WorkbookOperateParameter.md)
 - [WorkbookReplaceResponse](docs/WorkbookReplaceResponse.md)
 - [WorkbookResponse](docs/WorkbookResponse.md)
 - [WorkbookSettingsOperateParameter](docs/WorkbookSettingsOperateParameter.md)
 - [WorkbookSettingsResponse](docs/WorkbookSettingsResponse.md)
 - [WorksheetReplaceResponse](docs/WorksheetReplaceResponse.md)
 - [WorksheetResponse](docs/WorksheetResponse.md)
 - [Worksheets](docs/Worksheets.md)
 - [WorksheetsResponse](docs/WorksheetsResponse.md)
 - [XlsSaveOptions](docs/XlsSaveOptions.md)
 - [XlsbSaveOptions](docs/XlsbSaveOptions.md)
 - [XpsSaveOptions](docs/XpsSaveOptions.md)
 - [AutoShape](docs/AutoShape.md)
 - [ChartArea](docs/ChartArea.md)
 - [Legend](docs/Legend.md)
 - [OleObject](docs/OleObject.md)
 - [Picture](docs/Picture.md)
 - [Title](docs/Title.md)


