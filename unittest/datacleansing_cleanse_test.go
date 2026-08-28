package unittest_test

import (
	"context"
	"errors"
	"testing"

	"asposecellscloud"
	"asposecellscloud/datacleansing"
	"asposecellscloud/datasource"
	"asposecellscloud/internal/testutil"
)

func TestCleanse(t *testing.T) {
	tests := []struct {
		name string
		fn   func(context.Context, *asposecellscloud.AsposeCellsCloudClient, datasource.DataSource, datasource.DataSink, ...datacleansing.Option) error
		path string
	}{
		{"datacleansing.RemoveBlankRows", datacleansing.RemoveBlankRows, "/v4.0/cells/remove/blank-rows"},
		{"datacleansing.RemoveBlankColumns", datacleansing.RemoveBlankColumns, "/v4.0/cells/remove/blank-columns"},
		{"datacleansing.RemoveBlankWorksheets", datacleansing.RemoveBlankWorksheets, "/v4.0/cells/remove/blank-worksheets"},
		{"datacleansing.RemoveDuplicates", datacleansing.RemoveDuplicates, "/v4.0/cells/remove/duplicates"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, capture := testutil.NewServer(t, "cleaned")
			sink := &datasource.BytesSink{}

			var opts []datacleansing.Option
			if tt.name == "datacleansing.RemoveDuplicates" {
				opts = append(opts, datacleansing.WithWorksheet("Sheet1"))
			}
			if err := tt.fn(context.Background(), client, datasource.BytesSource([]byte("xlsx")), sink, opts...); err != nil {
				t.Fatalf("%s failed: %v", tt.name, err)
			}
			c := capture()
			if c.Method != "PUT" || c.Path != tt.path {
				t.Errorf("request = %s %s, want PUT %s", c.Method, c.Path, tt.path)
			}
			if got := string(c.Files["Spreadsheet"]); got != "xlsx" {
				t.Errorf("Spreadsheet part = %q, want xlsx", got)
			}
			if tt.name == "datacleansing.RemoveDuplicates" {
				if got := c.Query.Get("worksheet"); got != "Sheet1" {
					t.Errorf("worksheet = %q, want Sheet1", got)
				}
			}
			if got := string(sink.Bytes()); got != "cleaned" {
				t.Errorf("sink = %q, want cleaned", got)
			}
		})
	}
}

func TestCleanseValidation(t *testing.T) {
	client, _ := testutil.NewServer(t, "")
	ctx := context.Background()
	sink := &datasource.BytesSink{}
	src := datasource.BytesSource([]byte("x"))

	if err := datacleansing.RemoveBlankRows(ctx, client, nil, sink); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("datacleansing.RemoveBlankRows nil source: got %v, want ErrInvalidParam", err)
	}
	if err := datacleansing.RemoveBlankRows(ctx, client, src, nil); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("datacleansing.RemoveBlankRows nil sink: got %v, want ErrInvalidParam", err)
	}
	if err := datacleansing.RemoveBlankColumns(ctx, client, nil, sink); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("datacleansing.RemoveBlankColumns nil source: got %v, want ErrInvalidParam", err)
	}
	if err := datacleansing.RemoveBlankWorksheets(ctx, client, nil, sink); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("datacleansing.RemoveBlankWorksheets nil source: got %v, want ErrInvalidParam", err)
	}
	if err := datacleansing.RemoveDuplicates(ctx, client, nil, sink); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("datacleansing.RemoveDuplicates nil source: got %v, want ErrInvalidParam", err)
	}
	if err := datacleansing.RemoveDuplicates(ctx, client, src, nil); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("datacleansing.RemoveDuplicates nil sink: got %v, want ErrInvalidParam", err)
	}
}
