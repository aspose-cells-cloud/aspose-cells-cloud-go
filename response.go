package asposecellscloud

import "encoding/json"

// RichResponse unified API response structure
type RichResponse struct {
	StatusCode int                 // HTTP status code
	Headers    map[string][]string // Response Header
	Body       []byte              // Raw response body
}

// GetJSON Parse the response body into the specified struct
func (r *RichResponse) GetJSON(target interface{}) error {
	if len(r.Body) == 0 || target == nil {
		return nil
	}
	return json.Unmarshal(r.Body, target)
}

func (r *RichResponse) ToBytes() []byte {
	return r.Body
}

// ToString 返回字符串格式的数据
func (r *RichResponse) ToString() string {
	return string(r.Body)
}
