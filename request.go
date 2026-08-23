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
}
