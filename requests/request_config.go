package requests

type requestConfig struct {
	Params map[string]interface{}
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
