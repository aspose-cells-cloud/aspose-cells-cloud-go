# Testing Guide

## Overview

The integration test suite contains **402 test functions** across **46 test files**, generated from JSON configuration files in `TestingData/`.

## Test Structure

```
integrationtests/
├── go.mod                              # Separate module referencing parent SDK
├── cells_controller_test.go            # 40 tests for CellsController
├── conversion_test.go                  # 32 tests for ConversionController
├── workbook_controller_test.go         # 32 tests for WorkbookController
├── worksheet_controller_test.go        # 39 tests for WorksheetsController
├── light_cells_test.go                 # 19 tests for LightCellsController
├── pivot_tables_controller_test.go     # 19 tests for PivotTablesController
├── charts_controller_test.go           # 14 tests for ChartsController
└── ...                                 # 38 more test files
```

## Test Configuration Format

Tests are defined in JSON files under `TestingData/`:

```json
{
  "Name": "CellsController",
  "Folder": "CellsCloud30",
  "Variables": {
    "RemoteFolder": "TestData/In"
  },
  "Cases": [
    {
      "Name": "PostClearContents",
      "ApiMethod": "PostClearContents",
      "Description": ["Test for PostClearContents."],
      "Variables": {
        "LocalName": "Book1.xlsx",
        "RemoteName": "Book1.xlsx"
      },
      "Files": [
        {
          "LocalPath": "%LocalName%",
          "RemotePath": "%RemoteFolder%/%RemoteName%",
          "StorageName": ""
        }
      ],
      "Parameters": [
        {
          "Name": "name",
          "DataType": { "Identifier": "String" },
          "Value": "%RemoteName%"
        }
      ],
      "Assertions": [
        {
          "Type": "EqualsInteger",
          "Expression": "Code",
          "Value": "200"
        }
      ]
    }
  ]
}
```

## Running Tests

### Prerequisites

1. Valid Aspose Cloud credentials
2. Go 1.16+
3. Test data files in `testdata/` directory

### Setup

```bash
# Set credentials
export CELLS_CLOUD_CLIENT_ID="your-client-id"
export CELLS_CLOUD_CLIENT_SECRET="your-client-secret"
export CELLS_CLOUD_API_BASE_URL="https://api.aspose.cloud"

# Navigate to test directory
cd integrationtests

# Install dependencies
go mod tidy
```

### Execute Tests

```bash
# Run all tests
go test -v ./...

# Run specific test file
go test -v -run "TestCells" ./...

# Run single test
go test -v -run "TestPostClearContents" ./...

# Run with timeout
go test -v -timeout 30m ./...

# Run tests matching pattern
go test -v -run "TestWorkbook.*" ./...
```

## Generated Test Pattern

Each test function follows this structure:

```go
func TestPostClearContents(t *testing.T) {
    // 1. Initialize client
    client := asposecellscloud.NewAsposeCellsCloudClient(
        os.Getenv("CellsCloudClientId"),
        os.Getenv("CellsCloudClientSecret"),
        os.Getenv("CellsCloudApiBaseUrl"),
    )

    // 2. Build request
    request := requests.NewPostClearContentsRequest(
        "Book1.xlsx",          // Required params directly
        "Sheet1",
        requests.WithCommonParameter("range", "A1:C10"),   // Optional params
        requests.WithCommonParameter("folder", "TestData/In"),
    )

    // 3. Execute
    response, err := client.Do(context.Background(), request)
    if err != nil {
        t.Error(err)
    } else {
        t.Log("TestPostClearContents success.")
    }

    // 4. Verify (assertions added as TODO comments for manual review)
    _ = response  // TODO: Verify Code == 200
}
```

## Test Parameter Mapping

The test generator maps testing data parameters to request constructor arguments:

| Spec Parameter | Test Data Value | Constructor Argument |
|---------------|----------------|---------------------|
| Required, Path, String | `"Book1.xlsx"` | Direct: `"Book1.xlsx"` |
| Required, FormData, File | `"testdata/file.xlsx"` | Direct: `"testdata/file.xlsx"` |
| Optional, Query, Integer | `0` | `WithCommonParameter("param", 0)` |
| Optional, Query, String | `"value"` | `WithCommonParameter("param", "value")` |
| Required, Body, Class | `{...}` | `&models.Type{}` (placeholder) |

## Variable Resolution

Test data uses `%VariableName%` patterns that resolve against:

1. **Group variables** — Defined at top level (e.g., `%RemoteFolder%`)
2. **Case variables** — Defined per test case (e.g., `%LocalName%`)

```json
{
  "Variables": { "RemoteFolder": "TestData/In" },
  "Cases": [{
    "Variables": { "LocalName": "Book1.xlsx" },
    "Parameters": [{
      "Value": "%RemoteFolder%/%LocalName%"
    }]
  }]
}
```

## Known Limitations

1. **Complex object initialization**: Class/Container type parameters are initialized with empty structs (`&models.Type{}`). Manual adjustment may be needed for tests that require specific field values.

2. **Unresolved variables**: Variables without definitions in the test configuration produce empty string values.

3. **Unmatched APIs**: Test cases referencing API methods not in the spec are skipped with a warning.

4. **File uploads**: Tests requiring pre-uploaded files on cloud storage need manual file upload before execution.

## Adding New Tests

1. Create or update a JSON file in `TestingData/`
2. Follow the test configuration format
3. Run `python3 generate_tests.py`
4. Verify with `go build ./... && go vet ./...`
5. Add required test data files to `testdata/`

## Continuous Integration

Tests require valid cloud credentials and network access. They are designed for:

- **Local development**: Full suite with real API calls
- **CI/CD**: Run with service principal credentials
- **Smoke testing**: Run a subset of critical path tests
