package reporting

// TODO(v4): operations below existed only in the v3.0 API and were removed when
// the feature layer migrated to v4.0. The v4.0 API has no equivalent interface
// for them (verified against the live API, August 2026). Re-add them once the
// corresponding v4.0 endpoints become available.
//
//   - Assemble — merge an Excel template with an XML data source into a report
//     file. v3.0: POST /cells/assemble. Under v4.0 the endpoint returns 405.
//   - Analyze — analyze a set of cloud files. v3.0: POST /cells/analyze.
//     Under v4.0 the endpoint returns 405.
//   - DataTransformation — run a data-transformation pipeline. v3.0:
//     POST /cells/datatransformation. Under v4.0 the endpoint returns 405.
//
// The v4.0 AI analysis functions remain available: ReportAnalysis
// (PUT /v4.0/cells/ai/report/analysis), Summarize
// (PUT /v4.0/cells/ai/summarize/spreadsheet), AggregateByColor
// (PUT /v4.0/cells/calculate/aggergate/color) and MathCalculate
// (PUT /v4.0/cells/calculate/math).
