/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="configuration.go">
*   Copyright (c) 2025 Aspose.Cells Cloud
* </copyright>
* <summary>
*   Permission is hereby granted, free of charge, to any person obtaining a copy
*  of this software and associated documentation files (the "Software"), to deal
*  in the Software without restriction, including without limitation the rights
*  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
*  copies of the Software, and to permit persons to whom the Software is
*  furnished to do so, subject to the following conditions:
* 
*  The above copyright notice and this permission notice shall be included in all
*  copies or substantial portions of the Software.
* 
*  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
*  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
*  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
*  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
*  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
*  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
*  SOFTWARE.
* </summary> 
-------------------------------------------------------------------------------------------------------------------- **/

package asposecellscloud

import (
	"net/http"
	"strings"
	"time"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextOAuth2 takes a oauth2.TokenSource as authentication for the request.
	ContextOAuth2 = contextKey("token")

	// ContextBasicAuth takes BasicAuth as authentication for the request.
	ContextBasicAuth = contextKey("basic")

	// ContextAccessToken takes a string oauth2 access token as authentication for the request.
	ContextAccessToken = contextKey("accesstoken")

	// ContextAPIKey takes an APIKey as authentication for the request
	ContextAPIKey = contextKey("apikey")
)

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Prefix string
}

type Configuration struct {
	BasePath      		string            `json:"basePath,omitempty"`
	Version            string            `json:"version,omitempty"`
	Host          		string            `json:"host,omitempty"`
	Scheme        		string            `json:"scheme,omitempty"`
	DefaultHeader 		map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     		string            `json:"userAgent,omitempty"`
	HTTPClient    		*http.Client
	ClientSecret   		string
    ClientId       		string
    AccessToken   		string
	GetAccessTokenTime 	time.Time
}

func NewConfiguration(clientId string, clientSecret string, basePath string, version string) *Configuration {
	cfg := &Configuration{
		BasePath:      "https://api.aspose.cloud/v3.0",
		Version:       "v3.0",
		DefaultHeader: make(map[string]string),
		UserAgent:     "Aspose Cells Cloud SDK for Go",
		ClientSecret: clientSecret,
        ClientId: clientId,
	}
	if basePath != "" {
        cfg.BasePath = basePath
    }
		if strings.HasSuffix(cfg.BasePath, "/") {
		cfg.BasePath = cfg.BasePath[0 : len(cfg.BasePath)-1]
	}

	if strings.HasSuffix(cfg.BasePath, "v3.0") {
		cfg.BasePath = cfg.BasePath[0 : len(cfg.BasePath)-5]
		cfg.Version = "v3.0"
	}

	if strings.HasSuffix(cfg.BasePath, "v1.1") {
		cfg.BasePath = cfg.BasePath[0 : len(cfg.BasePath)-5]
		cfg.Version = "v1.1"
	}

	if version != "" {
		cfg.Version = version
	}

	return cfg
}

func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}
