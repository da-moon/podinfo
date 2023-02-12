package fastjson

import (
	"io"

	jsoniter "github.com/json-iterator/go"
	stacktrace "github.com/palantir/stacktrace"
)

// DecodeJSON decodes the json payload
func DecodeJSON(data []byte, out interface{}) error {
	if len(data) == 0 {
		return stacktrace.NewError("'data' being decoded is nil")
	}
	if out == nil {
		return stacktrace.NewError("output parameter 'out' is nil")
	}
	iter := jsoniter.ConfigFastest.BorrowIterator(data)
	defer jsoniter.ConfigFastest.ReturnIterator(iter)
	iter.ReadVal(&out)
	if iter.Error != nil {
		return stacktrace.Propagate(iter.Error, "Failed to decode JSON Blob")
	}
	return nil
}

// DecodeJSONFromReader Decodes/Unmarshals the given
// io.Reader pointing to a JSON, into a desired object
func DecodeJSONFromReader(r io.Reader, out interface{}) error {
	if r == nil {
		return stacktrace.NewError("'io.Reader' being decoded is nil")
	}
	if out == nil {
		return stacktrace.NewError("output parameter 'out' is nil")
	}
	dec := jsoniter.NewDecoder(r)
	dec.UseNumber()
	return dec.Decode(out)
}
