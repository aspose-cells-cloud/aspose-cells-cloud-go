package asposecellscloud

import "C"
import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Client SDK Core client
type CellsCloudClient struct {
	cfg        *Configuration
	httpClient *http.Client
	timeout    time.Duration
	retries    int
}

// Client configuration Option function type
type CellsCloudClientOption func(*CellsCloudClient)

// NewClient Create a new Client instance
func NewCellsCloudClient(clientId string, clientSecret string, baseURL string, opts ...CellsCloudClientOption) *CellsCloudClient {
	client := &CellsCloudClient{
		cfg:        NewConfiguration(clientId, clientSecret, baseURL, "v4.0"),
		httpClient: &http.Client{},
		timeout:    30 * time.Second,
		retries:    0,
	}

	// Apply all options
	for _, opt := range opts {
		opt(client)
	}

	// Set default timeout
	client.httpClient.Timeout = client.timeout

	return client
}

// WithTimeout Set request timeout
func WithTimeout(timeout time.Duration) CellsCloudClientOption {
	return func(c *CellsCloudClient) {
		c.timeout = timeout
		c.httpClient.Timeout = timeout
	}
}

// WithRetries Set the number of retries
func WithRetries(retries int) CellsCloudClientOption {
	return func(c *CellsCloudClient) {
		c.retries = retries
	}
}

// WithHeader Set global request headers
func WithHeader(key, value string) CellsCloudClientOption {
	return func(c *CellsCloudClient) {
		c.cfg.DefaultHeader[key] = value
	}
}

func (client *CellsCloudClient) addAuth(request *http.Request) (err error) {
	if err := client.RequestOauthToken(); err != nil {
		if err := client.RequestOauthToken(); err != nil {
			return err
		}
	}

	request.Header.Add("Authorization", "Bearer "+client.cfg.AccessToken)
	return nil
}

// RequestOauthToken function for requests OAuth token
func (client *CellsCloudClient) RequestOauthToken() error {
	var getAccessTokeUri = client.cfg.BasePath + "/v3.0/cells/connect/token"
	if client.cfg.Version == "v1.1" {
		getAccessTokeUri = client.cfg.BasePath + "/oauth2/token"
	}
	resp, err := http.PostForm(getAccessTokeUri, url.Values{
		"grant_type":    {"client_credentials"},
		"client_id":     {client.cfg.ClientId},
		"client_secret": {client.cfg.ClientSecret}})

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var tr TokenResponse
	if err = json.NewDecoder(resp.Body).Decode(&tr); err != nil {
		return err
	}
	client.cfg.GetAccessTokenTime = time.Now()
	client.cfg.AccessToken = tr.AccessToken
	return nil
}

// Do Execute one or more requests (core scheduling method)
func (client *CellsCloudClient) Do(ctx context.Context, requests ...RequestOption) ([]*RichResponse, error) {
	var responses []*RichResponse

	for _, req := range requests {
		resp, err := client.executeWithRetry(ctx, req)
		if err != nil {
			println(err.Error())
			return responses, &SDKError{
				Code:    -1,
				Message: "request execution failed",
				Err:     err,
			}
		}
		println(resp.StatusCode)
		println(resp.Body)
		responses = append(responses, resp)
	}
	println(len(responses))
	return responses, nil
}

// executeWithRetry Execution logic with retries
func (client *CellsCloudClient) executeWithRetry(ctx context.Context, req RequestOption) (*RichResponse, error) {
	var lastErr error

	for attempt := 0; attempt <= client.retries; attempt++ {
		resp, err := client.executeOnce(ctx, req)
		if err == nil {
			return resp, nil
		}
		lastErr = err
		// Simple exponential backoff (can be enhanced as needed)
		if attempt < client.retries {
			println(lastErr.Error())
			println("Sleep time.....")
			time.Sleep(time.Duration(1<<attempt) * time.Second)
		}
	}

	return nil, lastErr
}

// executeOnce Single execution logic
func (client *CellsCloudClient) executeOnce(ctx context.Context, req RequestOption) (*RichResponse, error) {
	// 1. Build URL
	u, err := url.Parse(client.cfg.BasePath + req.GetPath())
	if err != nil {
		return nil, err
	}
	// Concatenate Query Parameters
	query := u.Query()
	for k, v := range req.GetQueryParameters() {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	u.RawQuery = query.Encode()

	// 2. Build the request body
	jsonBody := req.GetJSONBody()
	multipartForm := req.GetMultipartForm()
	hasJSON := jsonBody != nil
	hasMultipart := multipartForm != nil && len(multipartForm) > 0
	println(u.Path)
	println(u.String())
	var bodyReader io.Reader
	var contentType string

	if hasJSON && hasMultipart {
		// Serialize the JSON into a string and stick it into the form as a regular text field
		jsonBytes, err := json.Marshal(jsonBody)
		if err != nil {
			return nil, err
		}
		// Agreement: Pass the entire JSON as a text field named 'request_body_data'
		multipartForm["request_body_data"] = string(jsonBytes)

		bodyReader, contentType, err = client.buildMultipartForm(multipartForm)
		if err != nil {
			return nil, err
		}
	} else if hasJSON {
		jsonBytes, err := json.Marshal(jsonBody)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(jsonBytes)
		contentType = "application/json"
	} else if hasMultipart {
		bodyReader, contentType, err = client.buildMultipartForm(multipartForm)
		if err != nil {
			return nil, err
		}
	}

	// 3. Create an HTTP request
	httpReq, err := http.NewRequestWithContext(ctx, req.GetMethod(), u.String(), bodyReader)
	if err != nil {
		return nil, err
	}

	// 4. Set request headers
	httpReq.Header.Set("Content-Type", contentType)
	for k, v := range client.cfg.DefaultHeader {
		httpReq.Header.Set(k, v)
	}
	httpReq.Header.Set("x-aspose-client", "go sdk")
	httpReq.Header.Set("x-aspose-client-version", globalCellsCloudSDKVersion)

	client.addAuth(httpReq)

	// 5. Execute request
	resp, err := client.httpClient.Do(httpReq)
	if err != nil {

		return nil, err
	}
	defer resp.Body.Close()
	// 6.Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 7. Read and build RichResponse
	return &RichResponse{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       body,
	}, nil
}

// buildMultipartForm Building a multipart/form-data request body
func (client *CellsCloudClient) buildMultipartForm(form map[string]interface{}) (io.Reader, string, error) {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	for key, value := range form {
		switch v := value.(type) {
		case string:
			if strings.HasPrefix(key, "@") { // file
				err := addFile(writer, key[1:], v)
				if err != nil {
					println("Add file exception.")
					return nil, "", err
				}
			} else { // form value
				// Plain text field
				if err := writer.WriteField(key, v); err != nil {
					return nil, "", err
				}
			}
		case []byte:
			// Binary file field
			println("[]byte")
			part, err := writer.CreateFormFile(key, key)
			if err != nil {
				return nil, "", err
			}
			if _, err := part.Write(v); err != nil {
				return nil, "", err
			}
		default:
			println("==================")
			println(v)
			println("==================")
		}
	}

	if err := writer.Close(); err != nil {
		return nil, "", err
	}

	return &buf, writer.FormDataContentType(), nil
}

func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}
