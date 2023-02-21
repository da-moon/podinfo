package chroot

import (
	"strings"

	physical "github.com/da-moon/podinfo/sdk/physical"
	stacktrace "github.com/palantir/stacktrace"
)

// sanityCheck is used to perform a sanity check on a key
func (v *View) sanityCheck(key string) error {
	if strings.Contains(key, "..") {
		errMsg := physical.RelativePathErrorMessage
		err := stacktrace.NewError(errMsg)
		return err
	}
	return nil
}

// expandKey is used to expand to the full key path with the prefix
func (v *View) expandKey(suffix string) string {
	return v.prefix + suffix
}

// truncateKey is used to remove the prefix of the key
func (v *View) truncateKey(full string) string {
	return strings.TrimPrefix(full, v.prefix)
}
