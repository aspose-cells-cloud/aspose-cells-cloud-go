package dataprocessing

// TODO(v4): operations below existed only in the v3.0 API and were removed when
// the feature layer migrated to v4.0. The v4.0 API has no equivalent interface
// for them (verified against the live API, August 2026). Re-add them once the
// corresponding v4.0 endpoints become available.
//
//   - ImportIntoWorkbook — import data into an existing cloud workbook via
//     ImportOption. v3.0: POST /cells/{name}/importdata. Under v4.0 the
//     endpoint returns 404. Local template-based import (ImportCSV/ImportJSON/
//     ImportXML/ImportData) remains available via the v4.0
//     PUT /v4.0/cells/import/data[/csv|/json|/xml] endpoints.
//   - ImportJSONIntoWorkbook — import raw JSON into a cloud workbook.
//     v3.0: POST /cells/{name}/importjson. Under v4.0 the endpoint returns 404.
//   - ImportXMLIntoWorkbook — import raw XML into a cloud workbook.
//     v3.0: POST /cells/{name}/importxml. Under v4.0 the endpoint returns 404.
