package requests

type requestConfig struct {
	Params           map[string]interface{}
	extraQueryParams map[string]string
}

type RequestOption interface {
	apply(*requestConfig)
}

type optionFunc func(*requestConfig)

func (f optionFunc) apply(c *requestConfig) {
	f(c)
}
func WithCommonParameter(key string, value interface{}) RequestOption {

	return optionFunc(func(c *requestConfig) {
		if c.Params == nil {
			c.Params = make(map[string]interface{})
		}
		println(key)
		c.Params[key] = value
	})
}

// WithQueryParameter adds a single custom query parameter (name/value pair) to the
// request. This is an implicit extension on top of the operation's declared query
// parameters, used to pass extra attributes such as save options.
func WithQueryParameter(name, value string) RequestOption {
	return optionFunc(func(c *requestConfig) {
		if c.extraQueryParams == nil {
			c.extraQueryParams = make(map[string]string)
		}
		c.extraQueryParams[name] = value
	})
}

// WithQueryParameters adds multiple custom query parameters at once.
func WithQueryParameters(params map[string]string) RequestOption {
	return optionFunc(func(c *requestConfig) {
		if c.extraQueryParams == nil {
			c.extraQueryParams = make(map[string]string)
		}
		for k, v := range params {
			c.extraQueryParams[k] = v
		}
	})
}
