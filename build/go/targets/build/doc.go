// Package build include Mage targets for building project
// binaries. It recursively searches for directories
// under `cmd/` . Each directory is supposed to contain
// go files under `main` package (i.e. a binary).
// - It cross compiles for Mac, Linux and Windows.
// - Build targets ensure the linking is fully static
// without any dependency on OS's underlying `libc` implementation
package build
