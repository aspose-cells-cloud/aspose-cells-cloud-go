package asposecellscloud

import "net/url"

// RequestOption Interfaces that all Requests must implement
type RequestOption interface {
	GetMethod() string
	GetHeaderParameters() map[string]string
	GetPath() string
	GetQueryParameters() url.Values
	GetJSONBody() interface{}
	GetMultipartForm() map[string]interface{}
	// Validate returns a non-nil error when a required parameter is missing.
	// Constructors never return nil; validation happens at execution time here.
	Validate() error
}
