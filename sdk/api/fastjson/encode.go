package fastjson

import (
	"bytes"
	"io"

	jsoniter "github.com/json-iterator/go"
	stacktrace "github.com/palantir/stacktrace"
)

// EncodeJSONWithoutErr Encodes/Marshals the given
// object into JSON but does not return an err
func EncodeJSONWithoutErr(in interface{}) []byte {
	res, _ := EncodeJSON(in)
	return res
}

// EncodeJSON Encodes/Marshals the given object into JSON
func EncodeJSON(in interface{}) ([]byte, error) {
	if in == nil {
		return nil, stacktrace.NewError("input for encoding is nil")
	}
	stream := jsoniter.ConfigFastest.BorrowStream(nil)
	defer jsoniter.ConfigFastest.ReturnStream(stream)
	stream.WriteVal(in)
	if stream.Error != nil {
		return nil, stacktrace.Propagate(stream.Error, "Failed to encode JSON")
	}
	return stream.Buffer(), nil
}

// EncodeJSONWithIndentation Encodes/Marshals the given object into JSON
// DEPRECATED
func EncodeJSONWithIndentation(in interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := EncodeJSONToWriter(buf, in, "", "    ")
	if err != nil {
		return nil, stacktrace.Propagate(err, "Failed to encode JSON with indentation")
	}
	return buf.Bytes(), nil
}

// EncodeJSONToWriter encodes/marshals a given interface
// to an io writer. it can also indent the output
func EncodeJSONToWriter(w io.Writer, in interface{}, prefix, indent string) error {
	if w == nil {
		return stacktrace.NewError("io.Writer is nil")
	}
	// TODO: ensure this does not cause any errors
	if in == nil {
		return stacktrace.NewError("input for encoding is nil")
	}
	enc := jsoniter.NewEncoder(w)
	enc.SetEscapeHTML(true)
	if prefix == "" && indent != "" {
		enc.SetIndent(prefix, indent)
	}
	return enc.Encode(in)
}
