package editor

// TODO(v4): operations below existed only in the v3.0 API and were removed when
// the feature layer migrated to v4.0. The v4.0 API has no equivalent interface
// for them (verified against the live API, August 2026). Re-add them once the
// corresponding v4.0 endpoints become available.
//
// Cell / row / column editing (all v3.0 cloud WorkbookRef operations):
//   - SetCellValue, SetRangeValue
//   - SetCellHTML
//   - InsertRows, InsertColumns, DeleteRows, DeleteColumns
//   - SetRowHeight, SetColumnWidth
//
// Worksheet management with no v4.0 equivalent:
//   - CopyWorksheet
//   - SetWorksheetVisibility
//
// Pictures and charts:
//   - AddPicture, DeletePicture, AddChart, DeleteChart
//
// The v4.0 worksheet management operations that remain available are all local
// (DataSource -> DataSink): AddWorksheet, DeleteWorksheet, RenameWorksheet,
// MoveWorksheet, ListWorksheets and CreateSpreadsheet.
