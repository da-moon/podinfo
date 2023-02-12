package response

import (
	"bytes"

	fastjson "github.com/da-moon/northern-labs-interview/sdk/api/fastjson"
	spew "github.com/davecgh/go-spew/spew"
)

// Response represents http request response
type Response struct {
	Success bool        `json:"success"`
	Body    interface{} `json:"body"`
}

// EncodeJSON encode Response struct as JSON
// using fastjson library
func (r *Response) EncodeJSON() ([]byte, error) {
	result, err := fastjson.EncodeJSON(r)
	if err != nil {
		return nil, err
	}
	return result, err
}

// FromJSON decodes json bytes into Response struct
// using fastjson library
func FromJSON(in []byte) (*Response, error) {
	var r Response
	err := fastjson.DecodeJSON(in, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

// GoString method is used to print
// Response passed as an operand to a %#v format.
func (r *Response) GoString() string {
	buf := new(bytes.Buffer)
	spew.Fdump(buf, r)
	return buf.String()
}
