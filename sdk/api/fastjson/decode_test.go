package fastjson_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	fastjson "github.com/da-moon/northern-labs-interview/sdk/api/fastjson"
	assert "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecode(t *testing.T) {
	assert := assert.New(t)
	expected := map[string]interface{}{
		"test":       "data",
		"validation": "process",
	}
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	err := enc.Encode(expected)

	t.Run("bytes", func(t *testing.T) {
		input := buf.Bytes()
		var actual map[string]interface{}
		err = fastjson.DecodeJSON(input, &actual)
		assert.NoError(err)
		assert.NotEmpty(actual)
		assert.Equal(expected, actual)
	})
	t.Run("io_reader", func(t *testing.T) {
		input := buf
		var actual map[string]interface{}
		err := fastjson.DecodeJSONFromReader(input, &actual)
		if err != nil {
			fmt.Printf("decoding err: %v\n", err)
		}
	})
}

func BenchmarkDecode(b *testing.B) {
	require := require.New(b)
	input := `{"test":"data","validation":"process"}`
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var actual map[string]interface{}
		err := fastjson.DecodeJSON([]byte(input), &actual)
		require.NoError(err)
	}
}
func BenchmarkStdLibDecode(b *testing.B) {
	require := require.New(b)
	input := `{"test":"data","validation":"process"}`
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var out map[string]interface{}
		err := json.Unmarshal([]byte(input), &out)
		require.NoError(err)
	}
}
func BenchmarkDecodeFromReader(b *testing.B) {
	require := require.New(b)
	input := `{"test":"data","validation":"process"}`
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var actual map[string]interface{}
		err := fastjson.DecodeJSONFromReader(bytes.NewReader([]byte(input)), &actual)
		require.NoError(err)
	}
}

func BenchmarkStdlibDecodeFromReader(b *testing.B) {
	require := require.New(b)
	input := `{"test":"data","validation":"process"}`
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var out map[string]interface{}
		r := bytes.NewReader([]byte(input))
		dec := json.NewDecoder(r)
		dec.UseNumber()
		err := dec.Decode(&out)
		require.NoError(err)
	}
}
