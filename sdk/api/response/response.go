package response

import (
	"bytes"
	"net/http"
	"strconv"
	"strings"

	fastjson "github.com/da-moon/northern-labs-interview/sdk/api/fastjson"
	spew "github.com/davecgh/go-spew/spew"
	stacktrace "github.com/palantir/stacktrace"
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
func WriteJSON(
	w http.ResponseWriter,
	r *http.Request,
	code int,
	headers map[string]string,
	body interface{},
) {
	var internalErr error
	defer func() {
		if internalErr != nil {
			LogErrorResponse(r, internalErr, "")
		}
	}()
	if headers != nil && len(headers) > 0 {
		for k, v := range headers {
			if strings.ToLower(k) != "content-length" || strings.ToLower(k) != "content-type" {
				w.Header().Set(k, v)
			}
		}
	}
	// writing JSON body if it was not nil
	if body != nil {
		resp, err := fastjson.EncodeJSON(body)
		if err != nil {
			internalErr = stacktrace.Propagate(err, ErrInternalServerError().Error())
			w.WriteHeader(int(ErrInternalServerError().StatusCode()))
			return
		}
		// TODO: see if this 'if' statement is needed
		if resp != nil {
			_, err = w.Write(resp)
			if err != nil {
				internalErr = stacktrace.Propagate(err, ErrInternalServerError().Error())
				return
			}
			w.(http.Flusher).Flush()
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Content-Length", strconv.Itoa(len(resp)))
	}

	// Setting headers if it was not nil
	w.WriteHeader(code)
	return
}
