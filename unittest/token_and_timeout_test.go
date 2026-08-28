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

// errInjectedTransportHit is the sentinel returned by alwaysErrorTransport; its
// presence in a wrapped error proves the request went through the injected
// transport rather than the package-level http.DefaultClient.
var errInjectedTransportHit = errors.New("injected transport reached")

type alwaysErrorTransport struct{}

func (alwaysErrorTransport) RoundTrip(*http.Request) (*http.Response, error) {
	return nil, errInjectedTransportHit
}

// TestRequestOauthTokenUsesInjectedClient proves the OAuth token endpoint is
// reached through the injected *http.Client/Transport. The token call runs
// before any data call, so the only RoundTrip observed is the token request.
func TestRequestOauthTokenUsesInjectedClient(t *testing.T) {
	client := asposecellscloud.NewAsposeCellsCloudClient("id", "secret", "https://example.invalid",
		asposecellscloud.WithHTTPClient(&http.Client{Transport: alwaysErrorTransport{}}))

	req := requests.NewConvertSpreadsheetRequest("pdf", "")
	req.SetSpreadsheetBytes([]byte("x"), "Spreadsheet")

	_, err := asposecellscloud.DoChecked(context.Background(), client, req)
	if err == nil {
		t.Fatal("expected the token request to fail through the injected transport")
	}
	if !errors.Is(err, errInjectedTransportHit) {
		t.Errorf("token request did not go through the injected transport: %v", err)
	}
}

// TestTimeoutResolutionOrderIndependent verifies the effective request timeout
// does not depend on the order of WithTimeout and WithHTTPClient: an injected
// client's non-zero Timeout wins, and WithTimeout alone is applied when the
// client sets none.
func TestTimeoutResolutionOrderIndependent(t *testing.T) {
	cases := []struct {
		name string
		opts []asposecellscloud.AsposeCellsCloudClientOption
	}{
		{"timeout-alone", []asposecellscloud.AsposeCellsCloudClientOption{
			asposecellscloud.WithTimeout(50 * time.Millisecond),
		}},
		{"client-then-timeout", []asposecellscloud.AsposeCellsCloudClientOption{
			asposecellscloud.WithHTTPClient(&http.Client{Timeout: 50 * time.Millisecond}),
			asposecellscloud.WithTimeout(5 * time.Second),
		}},
		{"timeout-then-client", []asposecellscloud.AsposeCellsCloudClientOption{
			asposecellscloud.WithTimeout(5 * time.Second),
			asposecellscloud.WithHTTPClient(&http.Client{Timeout: 50 * time.Millisecond}),
		}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			slow := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if strings.HasSuffix(r.URL.Path, "/cells/connect/token") {
					io.WriteString(w, tokenBody)
					return
				}
				time.Sleep(200 * time.Millisecond)
				io.WriteString(w, "slow")
			}))
			t.Cleanup(slow.Close)

			client := asposecellscloud.NewAsposeCellsCloudClient("id", "secret", slow.URL, tc.opts...)
			req := requests.NewConvertSpreadsheetRequest("pdf", "")
			req.SetSpreadsheetBytes([]byte("x"), "Spreadsheet")

			start := time.Now()
			_, err := asposecellscloud.DoChecked(context.Background(), client, req)
			elapsed := time.Since(start)

			if err == nil {
				t.Fatal("expected the 50ms timeout to abort the 200ms request")
			}
			if elapsed > 150*time.Millisecond {
				t.Errorf("effective timeout leaked toward 5s: request took %v", elapsed)
			}
		})
	}
}

// TestCheckResponseStatusWrapsErrRequestFailed verifies the ErrRequestFailed
// sentinel is reachable via errors.Is for non-2xx HTTP statuses.
func TestCheckResponseStatusWrapsErrRequestFailed(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/cells/connect/token") {
			io.WriteString(w, tokenBody)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"Status":"InternalServerError"}`)
	}))
	t.Cleanup(ts.Close)

	client := asposecellscloud.NewAsposeCellsCloudClient("id", "secret", ts.URL)
	req := requests.NewConvertSpreadsheetRequest("pdf", "")
	req.SetSpreadsheetBytes([]byte("x"), "Spreadsheet")

	_, err := asposecellscloud.DoChecked(context.Background(), client, req)
	if err == nil {
		t.Fatal("expected HTTP 500 to fail")
	}
	if !errors.Is(err, asposecellscloud.ErrRequestFailed) {
		t.Errorf("expected errors.Is(err, ErrRequestFailed), got %v", err)
	}
}
