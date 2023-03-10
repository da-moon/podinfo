package test

import (
	mg "github.com/magefile/mage/mg"
	sh "github.com/magefile/mage/sh"
)

// Target uses `gorc` to recursively
// run all project's unit tests.
func Target() error {
	args := []string{"test"}
	if mg.Verbose() {
		args = append(args, "-v")
	}
	return sh.RunV("gorc", args...)
}
