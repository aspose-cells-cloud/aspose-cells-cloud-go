package searcher

// TODO(v4): operation below was removed when the feature layer migrated to
// v4.0. Re-add it once the corresponding v4.0 endpoint works.
//
//   - SearchWorkbook — search text across a whole cloud workbook.
//     v4.0: PUT /v4.0/cells/{name}/search/content. The live endpoint currently
//     returns HTTP 400 ("Object reference not set to an instance of an
//     object"), so the workbook-scoped remote search is unusable. Use
//     SearchWorksheet or SearchRange instead; the worksheet- and range-scoped
//     endpoints work.
//
// Notes on the migration:
//   - The v3.0 findText/replaceText endpoints
//     (POST /cells/{name}/findText, POST /cells/{name}/replaceText) were
//     replaced by the v4.0 content endpoints (PUT /v4.0/cells/search/content,
//     PUT /v4.0/cells/replace/content, plus the remote worksheet/range
//     variants). They do not exist under v4.0.
//   - Local Search now requires a worksheet name (the live endpoint returns
//     400 without one), and local Replace returns the modified file directly
//     instead of round-tripping through cloud storage.
