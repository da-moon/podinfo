package files_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	files "github.com/da-moon/podinfo/internal/files"
	"github.com/stretchr/testify/assert"
)

func TestSubdirGlob(t *testing.T) {
	assert := assert.New(t)
	td, err := ioutil.TempDir("", "subdir-glob")
	assert.NoError(err)
	assert.NotEmpty(td)
	defer os.RemoveAll(td)

	err = os.Mkdir(filepath.Join(td, "subdir"), 0755)
	assert.NoError(err)
	err = os.Mkdir(filepath.Join(td, "subdir/one"), 0755)
	assert.NoError(err)

	err = os.Mkdir(filepath.Join(td, "subdir/two"), 0755)
	assert.NoError(err)
	expected := filepath.Join(td, "subdir")

	t.Run("march_exact_dir", func(t *testing.T) {
		actual, err := files.SubdirGlob(td, "subdir")
		if err != nil {
			t.Fatal(err)
		}
		assert.NoError(err)
		assert.Equal(expected, actual)
	})
	t.Run("single_match_wildcard", func(t *testing.T) {
		actual, err := files.SubdirGlob(td, "*")
		assert.NoError(err)
		assert.Equal(expected, actual)
	})
	t.Run("multiple_matches", func(t *testing.T) {
		actual, err := files.SubdirGlob(td, "subdir/*")
		assert.Error(err)
		assert.Empty(actual)
	})
	t.Run("non_existent", func(t *testing.T) {
		actual, err := files.SubdirGlob(td, "foo")
		assert.Error(err)
		assert.Empty(actual)
	})
}
