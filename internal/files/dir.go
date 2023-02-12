package files

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/da-moon/northern-labs-interview/internal/primitives"
	stacktrace "github.com/palantir/stacktrace"
)

func IsDir(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fi.IsDir()
}

// DirSize returns size of a target directory
func DirSize(dir string) (int64, error) {
	var size int64
	err := filepath.Walk(dir, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		size += info.Size()
		return nil
	})
	return size, err
}

const (
	gitDir = ".git"
)

// ReadDirFiles searches a root directory recursively for files with a pattern
func ReadDirFiles(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if pattern != "" {
			matched, err := filepath.Match(pattern, filepath.Base(path))
			if err != nil {
				err = stacktrace.Propagate(err, "filepath did not match pattern")
				return err
			} else if matched {
				relativePath := strings.TrimPrefix(path, root)
				if len(relativePath) > 0 {
					matches = append(matches, relativePath)
				}
			}
			return nil
		}
		relativePath := strings.TrimPrefix(path, root)
		if len(relativePath) > 0 {
			matches = append(matches, relativePath)
		}
		return nil
	})
	if err != nil {
		err = stacktrace.Propagate(err, "could not find file with root path '%s' and pattern '%s'", root, pattern)
		return nil, err
	}
	return matches, nil
}

// SafeMkdirAll creates directory tree in case it doesn't exist
// if it exists, it would fail
func SafeMkdirAll(path string) error {
	var err error
	if PathExist(path) {
		err = stacktrace.NewError("'%s' already exists", path)
		return err
	}
	err = mkdirSync(path)
	if err != nil {
		err = stacktrace.Propagate(err, "could not safely create directory tree at '%s' and sync it with disk", path)
		return err
	}
	return nil
}

// MkdirAll creates directory tree and won't return error if it exists
func MkdirAll(path string) error {
	var err error
	err = mkdirSync(path)
	if err != nil {
		if !os.IsExist(err) {
			err = stacktrace.Propagate(err, "could not create directory at '%s'", path)
			return err
		}
		return err
	}
	return nil
}
func mkdirSync(path string) error {
	// Directory already exists. exit
	if PathExist(path) {
		return nil
	}
	// [ NOTE ] => fuckup ?
	// _, err := os.Stat(path)
	// if err != nil {
	// 	return err
	// }
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return err
	}
	// Ensuring directory at target is created
	err = SyncPath(path)
	if err != nil {
		err = stacktrace.Propagate(err, "could not create directory at '%s'", path)
		return err
	}
	return nil
}

func CopyDir(src, dest string) error {
	var err error
	src, err = filepath.Abs(src)
	if err != nil {
		err = stacktrace.Propagate(err, "could not copy src directory '%s' to destination '%s'", src, dest)
		return err
	}
	dest, err = filepath.Abs(dest)
	if err != nil {
		err = stacktrace.Propagate(err, "could not copy src directory '%s' to destination '%s'", src, dest)
		return err
	}
	files, err := ReadDirFiles(src, "")
	if err != nil {
		err = stacktrace.Propagate(err, "could not copy src directory '%s' to destination '%s'", src, dest)
		return err
	}
	for _, f := range files {
		dp := primitives.PathJoin(dest, f)
		sp := primitives.PathJoin(src, f)
		err = os.MkdirAll(filepath.Dir(dp), 0777)
		if err != nil {
			err = stacktrace.Propagate(err, "could not copy src directory '%s' to destination '%s'", src, dest)
			return err
		}
		err := CopyFile(sp, dp)
		if err != nil {
			err = stacktrace.Propagate(err, "could not copy src directory '%s' to destination '%s'", src, dest)
			return err
		}
	}
	return nil
}

// ListFiles returns a list of files from the specified path.
// TODO: test this and use this instead of readdirfiles
func ListAllFilesInDir(path string) ([]string, error) {
	result := []string{}
	err := filepath.Walk(path, func(dir string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && info.Name() == gitDir {
			return filepath.SkipDir
		}
		result = append(result, dir)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
