# Feature Classification — Aspose.Cells Cloud SDK for Go

This document reclassifies the SDK's secondary-development (feature-layer) capabilities into broad,
conventional business categories. It supersedes the raw controller listing as the way to navigate
"what can I do with this SDK": instead of 42 controller names, you look up a capability such as **Data
Cleansing** or **Content Editing** and find the controllers, feature packages, and operations behind it.

The classification is a **grouping of intent**, not just a code layout, but the feature packages live
in directories named after these categories (see [Directory mapping](#directory-mapping)). Controllers
and feature packages are mapped to their *primary* category; cross-cutting operations are noted inline.

## Categories at a Glance

Six primary categories hold the feature-layer packages; three secondary categories cover the remaining
controllers.

| # | Category (EN) | Category (中文) | Feature package | Controllers | Operations |
|---|---------------|-----------------|------------------|-------------|-----------|
| 1 | Format Conversion | 格式转换 | `converter/` | Conversion (Convert*/Export*/SaveAs*), Batch (BatchConvert) | ~47 |
| 2 | Data Processing | 数据处理 | `dataprocessing/` | DataProcessing (Import*/Merge*/Split*), Transform (Swap/Flip/Transpose/Unpivot), Management (Compress/Repair), LightCells (Merge/Split/Compress/Repair/Assemble/Rotate/Watermark), Batch (BatchSplit) | ~35 |
| 3 | Data Cleansing | 数据清洗 | `datacleansing/` | Transform (RemoveBlank*/RemoveDuplicates), DataProcessing (Cleansing/Deduplication/Fill/DeleteIncompleteRows) | ~19 |
| 4 | Content Editing | 内容编辑 | `editor/` | Management, Cells, Worksheets, Workbook, Charts, PivotTables, Shapes, Pictures, OLE, Hyperlinks, Sparklines, AutoFilter, ConditionalFormatting, PageSetup, PageBreaks, Ranges, ListObjects, Properties, Validations, DataChecking, ChartArea, Autoshapes | ~260 |
| 5 | Search & Replace | 搜索与替换 | `searcher/` | Search, TextProcessing (AddText) | ~15 |
| 6 | Reporting & AI | 报告与智能分析 | `reporting/` | AI, Calculate, Analyse, StatisticalCharacters | ~11 |
| 7 | Text Processing | 文本处理 | — | TextProcessing | 24 |
| 8 | Security & Protection | 安全与保护 | — | Protection, CellsAuthority, Key, Batch (Protect/Lock/Unlock) | ~20 |
| 9 | System & Operations | 系统与运维 | — | CellsStatus, Storage, File, Folder, Specification, Task | ~21 |

## Directory mapping

The feature packages were reorganized into category-named directories. `export/` and `convert/` were
merged into `converter/`; `import/` moved under Data Processing; `datacleansing/` is new.

| Old directory | New directory |
|---------------|---------------|
| `convert/` | `converter/` |
| `export/` | `converter/` (merged with convert) |
| `editor/` | `editor/` (name kept) |
| `import/` | `dataprocessing/` |
| `search/` | `searcher/` |
| `report/` | `reporting/` |
| *(new)* | `datacleansing/` |

> The old directories no longer exist. Update import paths of the form
> `asposecellscloud/convert` → `asposecellscloud/converter`,
> `asposecellscloud/export` → `asposecellscloud/converter` (merged),
> `asposecellscloud/import` → `asposecellscloud/dataprocessing`,
> `asposecellscloud/search` → `asposecellscloud/searcher`,
> `asposecellscloud/report` → `asposecellscloud/reporting`.
> `asposecellscloud/editor` keeps its name.

## Feature Package → Category Mapping

| Feature package | Category | What it does |
|-----------------|----------|--------------|
| `converter/` | **Format Conversion** | Convert a local spreadsheet to 40+ formats (`Convert`, `ConvertToPDF/CSV/JSON/HTML/XLSX/PNG/DOCX/SQL`) and export a cloud workbook or worksheet to a target format (`Workbook`, `Worksheet`). |
| `dataprocessing/` | **Data Processing** | Import a data file (CSV/JSON/XML) into a template spreadsheet (`ImportData`, `ImportCSV`, `ImportJSON`, `ImportXML`) and merge/split workbooks (`MergeSpreadsheet`, `SplitSpreadsheet`, `MergeRemoteSpreadsheet`, `SplitRemoteSpreadsheet`). |
| `datacleansing/` | **Data Cleansing** | Remove blank rows/columns/worksheets and duplicate rows: `RemoveBlankRows`, `RemoveBlankColumns`, `RemoveBlankWorksheets`, `RemoveDuplicates`. |
| `editor/` | **Content Editing** | Worksheet lifecycle on a spreadsheet: `AddWorksheet`, `DeleteWorksheet`, `RenameWorksheet`, `MoveWorksheet`, `ListWorksheets`, `CreateSpreadsheet`. |
| `searcher/` | **Search & Replace** | Find and replace text locally and in cloud: `Search`, `Replace`, `SearchWorksheet`, `SearchRange`, `ReplaceWorkbook`, `ReplaceWorksheet`, `ReplaceRange`. |
| `reporting/` | **Reporting & AI** | AI report analysis and calculation: `ReportAnalysis`, `Summarize`, `AggregateByColor`, `MathCalculate`. |
| `datasource/` | **(shared)** | Plumbing for feeding data in and out, imported by all feature packages: `DataSource`/`DataSink` (`FilePathSource`, `BytesSource`, `ReaderSource`, `FilePathSink`, `BytesSink`, `UrlPathSink`). |

## Category Details

### 1. Format Conversion · 格式转换

Convert spreadsheets between 40+ formats (XLSX, XLS, CSV, PDF, HTML, JSON, ODS, …) and export cloud
workbooks/worksheets into a target format.

| Feature / Controller | Operations |
|----------------------|-----------|
| `converter/` | `Convert`, `ConvertToPDF`, `ConvertToCSV`, `ConvertToJSON`, `ConvertToHTML`, `ConvertToXlsx`, `ConvertToPNG`, `ConvertToDocx`, `ConvertToSQL` (local), `Workbook`, `Worksheet` (cloud export) |
| ConversionController | `ConvertSpreadsheet*`, `ConvertWorksheet*`, `ConvertRange*`, `ConvertChart*`, `ConvertTable*`, `SaveSpreadsheetAs`, `GetWorkbook`, `PutConvertWorkbook`, `PostWorkbookSaveAs`, `PostConvertWorkbookTo*`, `ExportSpreadsheetAsFormat`, `ExportWorksheetAsFormat`, `ExportChartAsFormat`, `ExportTableAsFormat`, `ExportRangeAsFormat` |
| BatchController | `PostBatchConvert` |
| FileController | `DownloadFile` |

### 2. Data Processing · 数据处理

Bring data into workbooks, merge/split whole documents, and run document-level transforms.

| Feature / Controller | Operations |
|----------------------|-----------|
| `dataprocessing/` | `ImportData`, `ImportCSV`, `ImportJSON`, `ImportXML` (local import), `MergeSpreadsheet`, `MergeRemoteSpreadsheet`, `SplitSpreadsheet`, `SplitRemoteSpreadsheet` (merge/split) |
| DataProcessingController | `ImportDataIntoSpreadsheet`, `ImportJSONDataIntoSpreadsheet`, `ImportXMLDataIntoSpreadsheet`, `ImportCSVDataIntoSpreadsheet`, `MergeSpreadsheets*`, `SplitSpreadsheet`, `SplitTable` |
| TransformController | `SwapRange`, `FlipData`, `TransposeData`, `UnpivotRange`, `UnpivotTable` |
| ManagementController | `CompressSpreadsheet`, `RepairSpreadsheet` |
| LightCellsController | `PostAssemble`, `PostMerge`, `PostSplit`, `PostCompress`, `PostRepair`, `PostRotate`, `PostWatermark` |
| BatchController | `PostBatchSplit` |

> **Merge & split** (`MergeSpreadsheet`, `MergeRemoteSpreadsheet`, `SplitSpreadsheet`,
> `SplitRemoteSpreadsheet`) were added to this package so that "combine workbooks" and "split a
> workbook by worksheets" live next to "import data" — the document-level data operations.

### 3. Data Cleansing · 数据清洗

Remove blank rows/columns/sheets, drop duplicates, fill gaps, delete incomplete rows, and normalize
shape.

| Feature / Controller | Operations |
|----------------------|-----------|
| `datacleansing/` | `RemoveBlankRows`, `RemoveBlankColumns`, `RemoveBlankWorksheets`, `RemoveDuplicates` (local src→sink) |
| TransformController | `RemoveSpreadsheetBlankRows`, `RemoveSpreadsheetBlankColumns`, `RemoveSpreadsheetBlankWorksheets`, `RemoveDuplicates` |
| DataProcessingController | `PostWorkbookDataCleansing`, `PostDataCleansing`, `PostWorkbookDataDeduplication`, `PostDataDeduplication`, `PostWorkbookDataFill`, `PostDataFill`, `PostDeleteIncompleteRows` |

> The high-level `datacleansing/` package wraps the local TransformController endpoints, which return
> the cleaned workbook as file bytes. The cloud `Post*Cleansing`/`Post*Deduplication`/`Post*DataFill`
> endpoints are reachable directly through the generated requests.

### 4. Content Editing · 内容编辑

The largest category: create, read, update, and delete everything *inside* a workbook — cells,
worksheets, ranges, charts, shapes, pictures, pivot tables, filters, formatting, and document
properties.

| Feature / Controller | Operations |
|----------------------|-----------|
| `editor/` | `CreateSpreadsheet`, `AddWorksheet`, `DeleteWorksheet`, `RenameWorksheet`, `MoveWorksheet`, `ListWorksheets` |
| ManagementController | `CreateSpreadsheet`, `AddWorksheetToSpreadsheet`, `DeleteWorksheetFromSpreadsheet`, `RenameWorksheetInSpreadsheet`, `MoveWorksheetInSpreadsheet`, `GetSpreadsheetStructure`, `GetMergedCells*`, `AcceptAllRevisions*` |
| CellsController | 40 cell operations (clear, merge, style, values, validations) |
| WorksheetsController | 39 worksheet operations (add, copy, delete, rename, page setup) |
| WorkbookController | 25 workbook operations (names, styles, text items, properties) |
| RangesController | 14 range operations (copy, merge, style, value) |
| ChartsController / ChartAreaController / AutoshapesController | chart + chart-area + autoshape operations |
| ShapesController / PicturesController / OleObjectsController / HyperlinksController / SparklineGroupsController | object-layer editing |
| PivotTablesController | 21 pivot table & filter operations |
| ConditionalFormattingsController / AutoFilterController | conditional formatting + auto filters |
| PageSetupController / PageBreaksController | page layout |
| ListObjectsController | table / list-object operations |
| PropertiesController | document properties |
| WorksheetValidationsController / DataCheckingController | cell validation & integrity checks |

### 5. Search & Replace · 搜索与替换

Find text, broken links, and broken references; replace in place across workbooks, worksheets, and
ranges.

| Feature / Controller | Operations |
|----------------------|-----------|
| `searcher/` | `Search`, `Replace`, `SearchWorksheet`, `SearchRange`, `ReplaceWorkbook`, `ReplaceWorksheet`, `ReplaceRange` |
| SearchController | `SearchSpreadsheetContent`, `SearchSpreadsheetAllTextItems`, `SearchSpreadsheetBrokenLinks`, `ReplaceSpreadsheetContent`, `SearchBrokenLinks*`, `ReplaceContent*` |
| LightCellsController | `PostSearch`, `PostReplace` |

### 6. Reporting & AI · 报告与智能分析

AI-assisted analysis, summarization, translation, math calculation, and character/word statistics.

| Feature / Controller | Operations |
|----------------------|-----------|
| `reporting/` | `ReportAnalysis`, `Summarize`, `AggregateByColor`, `MathCalculate` |
| AIController | `ReportAIAnalysis`, `SummarizeSpreadsheet`, `TranslateSpreadsheet`, `TranslateTextFile`, `DecomposeUserTask` |
| CalculateController | `AggregateCellsByColor`, `MathCalculate` |
| AnalyseController | `PostAnalyzeExcel` |
| StatisticalCharactersController | `PostCharacterCount`, `PostWordsCount`, `PostSpecifyWordsCount` |

### 7. Text Processing · 文本处理

Extract, add, convert, split, trim, and normalize text inside cells.

| Feature / Controller | Operations |
|----------------------|-----------|
| TextProcessingController | `ExtractText`, `AddText`, `SplitText`, `ConvertText`, `TrimCharacter*`, `UpdateWordCase*`, `RemoveCharacters*`, `RemoveCharactersByPosition*`, `RemoveDuplicateSubstrings*` |

> Cleaning-oriented text ops (`TrimCharacter`, `RemoveCharacters`, `UpdateWordCase`,
> `RemoveDuplicateSubstrings`) also serve **Data Cleansing**; editing-oriented text ops (`AddText`,
> `ExtractText`, `SplitText`, `ConvertText`) also serve **Content Editing**.

### 8. Security & Protection · 安全与保护

Encrypt, protect, lock, unlock, and digitally sign workbooks.

| Feature / Controller | Operations |
|----------------------|-----------|
| ProtectionController | `ProtectSpreadsheet`, `UnprotectSpreadsheet`, `SpreadsheetDigitalsignature`, `PostDigitalSignature`, `PostEncryptWorkbook`, `DeleteDecryptWorkbook`, `PostProtectWorkbook`, `DeleteUnProtectWorkbook`, `PutDocumentProtectFromChanges`, `DeleteDocumentUnProtectFromChanges`, `PostLock`, `PostUnlock`, `PostProtect` |
| CellsAuthorityController | `PostAccessToken` |
| KeyController | `GetPublicKey` |
| BatchController | `PostBatchProtect`, `PostBatchLock`, `PostBatchUnlock` |

### 9. System & Operations · 系统与运维

Health checks, storage/file/folder management, API specs, and task execution.

| Feature / Controller | Operations |
|----------------------|-----------|
| CellsStatusController | `GetAsposeCellsCloudStatus`, `CheckCloudServiceHealth`, `GetCellsCloudServicesHealthCheck`, `GetCellsCloudServiceStatus` |
| StorageController | `StorageExists`, `ObjectExists`, `GetDiscUsage`, `GetFileVersions` |
| FileController | `UploadFile`, `DownloadFile`, `CopyFile`, `MoveFile`, `DeleteFile` |
| FolderController | `GetFilesList`, `CreateFolder`, `CopyFolder`, `MoveFolder`, `DeleteFolder` |
| SpecificationController | `Spec`, `CodegenSpec` |
| TaskController | `PostRunTask` |

## Controller → Category Quick Map (all 42)

| Controller | Ops | Category |
|-----------|-----|----------|
| AIController | 5 | Reporting & AI |
| AnalyseController | 1 | Reporting & AI |
| AutoFilterController | 13 | Content Editing |
| AutoshapesController | 2 | Content Editing |
| BatchController | 5 | Data Processing |
| CalculateController | 2 | Reporting & AI |
| CellsAuthorityController | 1 | Security & Protection |
| CellsController | 40 | Content Editing |
| CellsStatusController | 4 | System & Operations |
| ChartAreaController | 3 | Content Editing |
| ChartsController | 24 | Content Editing |
| ConditionalFormattingsController | 9 | Content Editing |
| ConversionController | 42 | Format Conversion |
| DataCheckingController | 2 | Content Editing |
| DataProcessingController | 23 | Data Processing |
| FileController | 5 | System & Operations |
| FolderController | 5 | System & Operations |
| HypelinksController | 6 | Content Editing |
| KeyController | 1 | Security & Protection |
| LightCellsController | 15 | Data Processing |
| ListObjectsController | 13 | Content Editing |
| ManagementController | 14 | Content Editing |
| OleObjectsController | 6 | Content Editing |
| PageBreaksController | 10 | Content Editing |
| PageSetupController | 9 | Content Editing |
| PicturesController | 7 | Content Editing |
| PivotTablesController | 21 | Content Editing |
| PropertiesController | 5 | Content Editing |
| ProtectionController | 13 | Security & Protection |
| RangesController | 14 | Content Editing |
| SearchController | 14 | Search & Replace |
| ShapesController | 8 | Content Editing |
| SparklineGroupsController | 6 | Content Editing |
| SpecificationController | 2 | System & Operations |
| StatisticalCharactersController | 3 | Reporting & AI |
| StorageController | 4 | System & Operations |
| TaskController | 1 | System & Operations |
| TextProcessingController | 24 | Text Processing |
| TransformController | 9 | Data Cleansing |
| WorkbookController | 25 | Content Editing |
| WorksheetsController | 39 | Content Editing |
| WorksheetValidationsController | 6 | Content Editing |

## Notes on Cross-Cutting Controllers

- **LightCellsController** (15) is intentionally lightweight and cross-cutting: `PostAssemble` /
  `PostMerge` / `PostSplit` / `PostCompress` / `PostRepair` / `PostRotate` / `PostWatermark` →
  **Data Processing**; `PostSearch` / `PostReplace` → **Search & Replace**; `PostClearObjects` /
  `PostReverse` → **Data Cleansing / Data Processing**; `GetMetadata` / `PostMetadata` /
  `DeleteMetadata` → **Content Editing** (document properties). It is listed under Data Processing as
  its primary home.
- **TextProcessingController** (24) splits between **Text Processing** (its primary home) and the
  cleansing/editing categories as noted above.
- **ConversionController**'s `Export*` operations and **FileController**'s `DownloadFile` serve
  **Format Conversion** (cloud export) even though the controllers are primarily Format Conversion and
  System & Operations respectively.
- **DataProcessingController**'s `Import*` operations serve **Data Processing** (import) and its
  `*Cleansing*` / `*Deduplication*` / `*DataFill*` / `*DeleteIncompleteRows*` operations serve
  **Data Cleansing**, even though the controller is primarily **Data Processing**.
