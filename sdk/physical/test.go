package physical

import (
	"context"
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ExerciseBackend runs a series of tests against the given backend to ensure
// it is functioning properly. This is intended to be used in tests
// taken from
// https://github.com/hashicorp/vault/blob/master/sdk/physical/testing.go
func ExerciseBackend(t testing.TB, b Backend) {
	t.Helper()

	// Should be empty
	keys, err := b.List(context.Background(), "")
	assert.NoError(t, err, "initial list failed")
	assert.Empty(t, keys)

	// Delete should work if it does not exist
	err = b.Delete(context.Background(), "foo")
	assert.NoError(t, err, "idempotent delete")

	// Get should fail
	out, err := b.Get(context.Background(), "foo")

	assert.Error(t, err, "initial get successful")
	assert.Nil(t, out, "initial get was not nil")

	// Make an entry
	e := &Entry{Key: "foo", Value: []byte("test")}
	err = b.Put(context.Background(), e)
	assert.NoError(t, err)

	// Get should work
	out, err = b.Get(context.Background(), "foo")
	assert.NoError(t, err)
	assert.Equal(t, out, e)

	// List should not be empty
	keys, err = b.List(context.Background(), "")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(keys))
	assert.Equal(t, "foo", keys[0])
	// Delete should work
	err = b.Delete(context.Background(), "foo")
	assert.NoError(t, err)
	// Should be empty
	keys, err = b.List(context.Background(), "")

	assert.NoError(t, err, "list after delete")
	assert.Empty(t, keys, "list after delete")

	// Get should fail
	out, err = b.Get(context.Background(), "foo")
	assert.Nil(t, out, "get after delete")
	assert.Error(t, err, "get after delete")

	// ─── TODO CONTINUE ──────────────────────────────────────────────────────────────

	// Multiple Puts should work; GH-189
	e = &Entry{Key: "foo", Value: []byte("test")}
	err = b.Put(context.Background(), e)
	if err != nil {
		t.Fatalf("multi put 1 failed: %v", err)
	}
	e = &Entry{Key: "foo", Value: []byte("test")}
	err = b.Put(context.Background(), e)
	if err != nil {
		t.Fatalf("multi put 2 failed: %v", err)
	}

	// Make a nested entry
	e = &Entry{Key: "foo/bar", Value: []byte("baz")}
	err = b.Put(context.Background(), e)
	assert.Error(t, err)
	e = &Entry{Key: "bar/baz", Value: []byte("baz")}
	err = b.Put(context.Background(), e)
	if err != nil {
		t.Fatalf("nested put failed: %v", err)
	}

	// Get should work
	out, err = b.Get(context.Background(), "bar/baz")
	if err != nil {
		t.Fatalf("get failed: %v", err)
	}
	if !reflect.DeepEqual(out, e) {
		t.Errorf("bad: %v expected: %v", out, e)
	}

	keys, err = b.List(context.Background(), "")
	assert.NoError(t, err, "list multi failed")
	sort.Strings(keys)
	assert.ElementsMatch(t, []string{"bar/", "foo"}, keys)

	// Delete with children should work
	err = b.Delete(context.Background(), "foo")
	if err != nil {
		t.Fatalf("delete after multi: %v", err)
	}

	// Get should return the child
	out, err = b.Get(context.Background(), "bar/baz")
	if err != nil {
		t.Fatalf("get after multi delete: %v", err)
	}
	if out == nil {
		t.Errorf("get after multi delete not nil: %v", out)
	}

	// Removal of nested secret should not leave artifacts
	e = &Entry{Key: "foo/nested1/nested2/nested3", Value: []byte("baz")}
	err = b.Put(context.Background(), e)
	if err != nil {
		t.Fatalf("deep nest: %v", err)
	}

	err = b.Delete(context.Background(), "foo/nested1/nested2/nested3")
	if err != nil {
		t.Fatalf("failed to remove deep nest: %v", err)
	}

	keys, err = b.List(context.Background(), "bar/")

	assert.NoError(t, err)
	sort.Strings(keys)
	assert.ElementsMatch(t, []string{"baz"}, keys)

	// Make a second nested entry to test prefix removal
	e = &Entry{Key: "foo/zip", Value: []byte("zap")}
	err = b.Put(context.Background(), e)
	if err != nil {
		t.Fatalf("failed to create second nested: %v", err)
	}

	// Delete should not remove the prefix
	err = b.Delete(context.Background(), "bar/baz")
	if err != nil {
		t.Fatalf("failed to delete nested prefix: %v", err)
	}

	keys, err = b.List(context.Background(), "")
	if err != nil {
		t.Fatalf("list nested prefix: %v", err)
	}
	if len(keys) != 1 || keys[0] != "foo/" {
		t.Errorf("should be exactly 1 key == foo/: %v", keys)
	}

	// Delete should remove the prefix
	err = b.Delete(context.Background(), "foo/zip")
	if err != nil {
		t.Fatalf("failed to delete second prefix: %v", err)
	}

	keys, err = b.List(context.Background(), "")
	if err != nil {
		t.Fatalf("listing after second delete failed: %v", err)
	}
	if len(keys) != 0 {
		t.Errorf("should be empty at end: %v", keys)
	}

	// When the root path is empty, adding and removing deep nested values should not break listing
	e = &Entry{Key: "foo/nested1/nested2/value1", Value: []byte("baz")}
	err = b.Put(context.Background(), e)
	if err != nil {
		t.Fatalf("deep nest: %v", err)
	}

	e = &Entry{Key: "foo/nested1/nested2/value2", Value: []byte("baz")}
	err = b.Put(context.Background(), e)
	if err != nil {
		t.Fatalf("deep nest: %v", err)
	}

	err = b.Delete(context.Background(), "foo/nested1/nested2/value2")
	if err != nil {
		t.Fatalf("failed to remove deep nest: %v", err)
	}

	keys, err = b.List(context.Background(), "")
	if err != nil {
		t.Fatalf("listing of root failed after deletion: %v", err)
	}
	if len(keys) == 0 {
		t.Errorf("root is returning empty after deleting a single nested value, expected nested1/: %v", keys)
		keys, err = b.List(context.Background(), "foo/nested1")
		if err != nil {
			t.Fatalf("listing of expected nested path 'foo/nested1' failed: %v", err)
		}
		// prove that the root should not be empty and that foo/nested1 exists
		if len(keys) != 0 {
			t.Logf("  keys can still be listed from nested1/ so it's not empty, expected nested2/: %v", keys)
		}
	}

	// cleanup left over listing bug test value
	err = b.Delete(context.Background(), "foo/nested1/nested2/value1")
	if err != nil {
		t.Fatalf("failed to remove deep nest: %v", err)
	}

	keys, err = b.List(context.Background(), "")
	if err != nil {
		t.Fatalf("listing of root failed after delete of deep nest: %v", err)
	}
	if len(keys) != 0 {
		t.Errorf("should be empty at end: %v", keys)
	}
}

// ExerciseBackendListPrefix exercises the ListPrefix method of a backend.
func ExerciseBackendListPrefix(t testing.TB, b Backend) {
	t.Helper()
	e1 := &Entry{Key: "foo", Value: []byte("test")}
	e2 := &Entry{Key: "bar/foo", Value: []byte("test")}
	e3 := &Entry{Key: "bar/fizz/fuzz", Value: []byte("test")}

	defer func() {
		b.Delete(context.Background(), "foo")
		b.Delete(context.Background(), "bar/foo")
		b.Delete(context.Background(), "bar/fizz/fuzz")
	}()

	err := b.Put(context.Background(), e1)
	if err != nil {
		t.Fatalf("failed to put entry 1: %v", err)
	}
	err = b.Put(context.Background(), e2)
	if err != nil {
		t.Fatalf("failed to put entry 2: %v", err)
	}
	err = b.Put(context.Background(), e3)
	if err != nil {
		t.Fatalf("failed to put entry 3: %v", err)
	}

	// Scan the root
	keys, err := b.List(context.Background(), "")
	if err != nil {
		t.Fatalf("list root: %v", err)
	}
	sort.Strings(keys)
	assert.ElementsMatch(t, []string{"foo", "bar/"}, keys)

	// Scan foo/
	keys, err = b.List(context.Background(), "bar/")
	if err != nil {
		t.Fatalf("list level 1: %v", err)
	}
	sort.Strings(keys)
	assert.ElementsMatch(t, []string{"fizz/", "foo"}, keys)
}

// ExerciseTransactionalBackend exercises the transactional methods of a
// physical backend.
func ExerciseTransactionalBackend(t testing.TB, b Backend) {
	t.Helper()
	require := require.New(t)
	assert := assert.New(t)
	// Add a few keys so that we test rollback with deletion
	err := b.Put(context.Background(), &Entry{
		Key:   "foo",
		Value: []byte("bar"),
	})
	require.NoError(err)
	err = b.Put(context.Background(), &Entry{
		Key:   "zip",
		Value: []byte("zap"),
	})
	require.NoError(err)
	err = b.Put(context.Background(), &Entry{
		Key: "deleteme",
	})
	require.Error(err)
	err = b.Put(context.Background(), &Entry{
		Key:   "deleteme",
		Value: []byte("deleteme"),
	})
	require.NoError(err)
	err = b.Put(context.Background(), &Entry{
		Key: "deleteme2",
	})
	require.Error(err)
	err = b.Put(context.Background(), &Entry{
		Key:   "deleteme2",
		Value: []byte("deleteme2"),
	})
	require.NoError(err)
	txns := []*TxnEntry{
		{
			Operation: PutOperation,
			Entry: Entry{
				Key:   "foo",
				Value: []byte("bar2"),
			},
		},
		{
			Operation: DeleteOperation,
			Entry: Entry{
				Key: "deleteme",
			},
		},
		{
			Operation: PutOperation,
			Entry: Entry{
				Key:   "foo",
				Value: []byte("bar3"),
			},
		},
		{
			Operation: DeleteOperation,
			Entry: Entry{
				Key: "deleteme2",
			},
		},
		{
			Operation: PutOperation,
			Entry: Entry{
				Key:   "zip",
				Value: []byte("zap3"),
			},
		},
	}
	tb, ok := b.(Transactional)
	require.True(ok)
	txnErrs := tb.Transaction(context.Background(), txns)
	require.Empty(txnErrs)
	{
		actual, err := b.List(context.Background(), "")
		assert.NoError(err)
		expected := []string{"foo", "zip"}
		assert.Equal(expected, actual)
	}
	{
		entry, err := b.Get(context.Background(), "foo")
		assert.NoError(err)
		assert.NotEmpty(entry)
		assert.NotEmpty(entry.Value)
		actual := string(entry.Value)
		expected := "bar3"
		assert.Equal(expected, actual, "updates did not apply correctly")
	}
	{
		entry, err := b.Get(context.Background(), "zip")
		assert.NoError(err)
		assert.NotEmpty(entry)
		assert.NotEmpty(entry.Value)
		actual := string(entry.Value)
		expected := "zap3"
		assert.Equal(expected, actual, "updates did not apply correctly")
	}
}
