package converter

// TODO(v4): operations below existed only in the v3.0 API and were removed when
// the feature layer migrated to v4.0. The v4.0 API has no equivalent interface
// for them (verified against the live API, August 2026). Re-add them once the
// corresponding v4.0 endpoints become available.
//
// Component-level converts (all v3.0 cloud WorkbookRef operations):
//   - ConvertWorksheet* — convert a worksheet to another format.
//     v3.0: POST /cells/convert/worksheet. Under v4.0 the endpoint returns 404.
//   - ConvertRangeToPDF / ConvertChartToPDF / ConvertTableToPDF — component
//     converts. v3.0: POST /cells/convert/range|chart|table. Under v4.0 these
//     endpoints return 404.
//
// Whole-workbook conversion remains available: Convert / ConvertTo*
// (PUT /v4.0/cells/convert/spreadsheet?format=...) plus the cloud export
// helpers Workbook / Worksheet
// (GET /v4.0/cells/{name}?format= and GET /v4.0/cells/{name}/worksheets/{ws}?format=).
//
// Export operations with no v4.0 equivalent:
//   - WorksheetXML — export a workbook to its XML representation.
//     v3.0: POST /cells/{name}/exportxml. Under v4.0 the endpoint returns 404.
//   - RangeValues — read the cell values of a range as a string.
//     v3.0: GET /cells/{name}/worksheets/ranges/{namerange}/value.
//     Under v4.0 the endpoint returns 404.
