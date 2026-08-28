package unittest_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	asposecellscloud "asposecellscloud"
	"asposecellscloud/requests"
)

const tokenBody = `{"access_token":"tok123","token_type":"Bearer","expires_in":3600}`

// flakyTransport fails the first N RoundTrip calls, then forwards to base.
type flakyTransport struct {
	base     http.RoundTripper
	failures int
	attempts int
}

func (f *flakyTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	f.attempts++
	if f.attempts <= f.failures {
		return nil, errors.New("connection reset by peer")
	}
	return f.base.RoundTrip(req)
}

func newTokenServer(t *testing.T, apiBody string) *httptest.Server {
	t.Helper()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/cells/connect/token") {
			io.WriteString(w, tokenBody)
			return
		}
		io.WriteString(w, apiBody)
	}))
	t.Cleanup(ts.Close)
	return ts
}

// TestWithHTTPClient makes the injected *http.Client drive the request, and
// verifies its own timeout is honored rather than overwritten by the default.
func TestWithHTTPClient(t *testing.T) {
	// 1. A custom client with a short timeout must abort a slow server response.
	slow := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/cells/connect/token") {
			io.WriteString(w, tokenBody)
			return
		}
		time.Sleep(500 * time.Millisecond)
		io.WriteString(w, "slow")
	}))
	t.Cleanup(slow.Close)

	client := asposecellscloud.NewAsposeCellsCloudClient("id", "secret", slow.URL,
		asposecellscloud.WithHTTPClient(&http.Client{Timeout: 100 * time.Millisecond}))
	req := requests.NewConvertSpreadsheetRequest("pdf", "")
	req.SetSpreadsheetBytes([]byte("x"), "Spreadsheet")
	if _, err := asposecellscloud.DoChecked(context.Background(), client, req); err == nil {
		t.Fatal("expected the injected client timeout to abort the slow request")
	}
}

// TestRetryFailFastOnValidation verifies that a missing required parameter is
// rejected before the retry loop, so no network attempt is ever made even when
// retries are configured.
func TestRetryFailFastOnValidation(t *testing.T) {
	var attempts int
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/cells/connect/token") {
			io.WriteString(w, tokenBody)
			return
		}
		attempts++
		io.WriteString(w, "ok")
	}))
	t.Cleanup(ts.Close)

	client := asposecellscloud.NewAsposeCellsCloudClient("id", "secret", ts.URL,
		asposecellscloud.WithHTTPClient(&http.Client{Transport: &flakyTransport{base: http.DefaultTransport, failures: 100}}),
		asposecellscloud.WithRetries(5))

	// Missing required file parameter.
	req := requests.NewConvertSpreadsheetRequest("pdf", "")
	if _, err := asposecellscloud.DoChecked(context.Background(), client, req); err == nil {
		t.Fatal("expected Validate to reject the request")
	}
	if attempts != 0 {
		t.Errorf("expected 0 network attempts (fail-fast), got %d", attempts)
	}
}

// TestRetrySucceedsAfterTransientFailure verifies the exponential backoff loop
// retries a transport failure and eventually succeeds. The OAuth token request
// goes through the same injected transport, so the first attempt consumes it:
// attempt 1 fails on the token call, attempt 2 refreshes the token and runs
// the data call to completion (3 total RoundTrips).
func TestRetrySucceedsAfterTransientFailure(t *testing.T) {
	ts := newTokenServer(t, "ok")
	flaky := &flakyTransport{base: http.DefaultTransport, failures: 1}

	client := asposecellscloud.NewAsposeCellsCloudClient("id", "secret", ts.URL,
		asposecellscloud.WithHTTPClient(&http.Client{Transport: flaky}),
		asposecellscloud.WithRetries(2))

	req := requests.NewConvertSpreadsheetRequest("pdf", "")
	req.SetSpreadsheetBytes([]byte("x"), "Spreadsheet")

	resp, err := asposecellscloud.DoChecked(context.Background(), client, req)
	if err != nil {
		t.Fatalf("DoChecked failed after retries: %v", err)
	}
	if resp.ToString() != "ok" {
		t.Errorf("body = %q, want ok", resp.ToString())
	}
	if flaky.attempts != 3 {
		t.Errorf("expected 3 attempts (token retry + data), got %d", flaky.attempts)
	}
}

// TestRetryContextCancellation verifies that a canceled context interrupts the
// backoff sleep instead of waiting out the full schedule.
func TestRetryContextCancellation(t *testing.T) {
	ts := newTokenServer(t, "ok")
	always := &flakyTransport{base: http.DefaultTransport, failures: 100}

	client := asposecellscloud.NewAsposeCellsCloudClient("id", "secret", ts.URL,
		asposecellscloud.WithHTTPClient(&http.Client{Transport: always}),
		asposecellscloud.WithRetries(5))

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	req := requests.NewConvertSpreadsheetRequest("pdf", "")
	req.SetSpreadsheetBytes([]byte("x"), "Spreadsheet")

	start := time.Now()
	_, err := asposecellscloud.DoChecked(ctx, client, req)
	elapsed := time.Since(start)

	if err == nil {
		t.Fatal("expected an error after context cancellation")
	}
	if !strings.Contains(err.Error(), "context deadline exceeded") {
		t.Errorf("expected context deadline error, got %v", err)
	}
	// Uninterrupted backoff would take 1+2+4+8+16 = 31s; cancellation must
	// return well before that.
	if elapsed > 3*time.Second {
		t.Errorf("cancellation did not interrupt backoff, took %v", elapsed)
	}
}
