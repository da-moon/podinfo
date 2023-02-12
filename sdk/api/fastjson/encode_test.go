package fastjson_test

import (
	"bytes"
	"encoding/json"
	"testing"

	fastjson "github.com/da-moon/northern-labs-interview/sdk/api/fastjson"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncode(t *testing.T) {
	assert := assert.New(t)
	expected := map[string]interface{}{
		"validation": "process",
		"test":       "data",
	}
	actualBytes, err := fastjson.EncodeJSON(expected)
	assert.NoError(err)
	assert.NotNil(actualBytes)
	dec := json.NewDecoder(bytes.NewBuffer(actualBytes))
	actual := make(map[string]interface{})
	err = dec.Decode(&actual)
	assert.NoError(err)
	assert.Equal(expected, actual)
}

func BenchmarkEncode(b *testing.B) {
	require := require.New(b)
	input := map[string]interface{}{
		"test":       "data",
		"validation": "process",
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := fastjson.EncodeJSON(input)
		require.NoError(err)
	}
}

func BenchmarkStdLibEncode(b *testing.B) {
	require := require.New(b)
	input := map[string]interface{}{
		"test":       "data",
		"validation": "process",
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		err := enc.Encode(input)
		require.NoError(err)
	}
}
