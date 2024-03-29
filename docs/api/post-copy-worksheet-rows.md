# **postCopyWorksheetRows API**

Copy data and formats from specific entire rows in the worksheet. 

```bash

POST http://api.aspose.cloud/v3.0//cells/{name}/worksheets/{sheetName}/cells/rows/copy

```

## The request parameters of **postCopyWorksheetRows** API are: 

| Parameter Name | Type | Path/Query String/HTTPBody | Description | 
| :- | :- | :- |:- | 
|name|String|Path|The file name.|
|sheetName|String|Path|The worksheet name.|
|sourceRowIndex|Integer|Query|Source row index|
|destinationRowIndex|Integer|Query|Destination row index|
|rowNumber|Integer|Query|The copied row number|
|worksheet|String|Query|The worksheet name.|
|folder|String|Query|The folder where the file is situated.|
|storageName|String|Query|The storage name where the file is situated.|


The [OpenAPI Specification](https://reference.aspose.cloud/cells/#/CellsController/PostCopyWorksheetRows) defines a publicly accessible programming interface and lets you carry out REST interactions directly from a web browser.
