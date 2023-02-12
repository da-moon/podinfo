package files

import (
	"os"
	"path/filepath"

	stacktrace "github.com/palantir/stacktrace"
)

// TODO: maybe make this linux only
// TODO: add direct io opener
// PathExist returns whether the given path exists.
func PathExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func OpenPath(path string) (*os.File, os.FileInfo, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	fi, err := f.Stat()
	if err != nil {
		f.Close()
		return nil, nil, err
	}
	return f, fi, nil
}

// SafeOpenPath removes empty files after opening
// empty files are most often result of a failed io
func SafeOpenPath(path string) (*os.File, os.FileInfo, error) {
	if !PathExist(path) {
		err := os.ErrNotExist
		return nil, nil, err
	}
	f, fi, err := OpenPath(path)
	if err != nil {
		return nil, nil, err
	}
	if !fi.IsDir() && fi.Size() == 0 {
		err = os.Remove(path)
		if err != nil {
			return nil, nil, err
		}
	}

	return f, fi, err
}

// SyncPath makes sure file at a certain path is synced with physical disk
func SyncPath(path string) error {
	var err error
	f, _, err := OpenPath(path)
	if err != nil {
		err = stacktrace.Propagate(err, "could not sync path '%s'", path)
		return err
	}
	err = f.Sync()
	if err != nil {
		_ = f.Close()
		err = stacktrace.Propagate(err, "could not flush path '%s' to disk", path)
		return err
	}
	err = f.Close()
	if err != nil {
		err = stacktrace.Propagate(err, "could not close file descriptor at '%s' to disk", path)
		return err
	}
	return nil
}

// SubdirGlob returns the actual subdir with globbing processed.
//
// dst should be a destination directory that is already populated (the
// download is complete) and subDir should be the set subDir. If subDir
// is an empty string, this returns an empty string.
//
// The returned path is the full absolute path.
func SubdirGlob(dst, subDir string) (string, error) {
	matches, err := filepath.Glob(filepath.Join(dst, subDir))
	if err != nil {
		err = stacktrace.Propagate(err, "could not return the actual subdir with globbing processed")
		return "", err
	}

	if len(matches) == 0 {
		return "", stacktrace.NewError("subdir %q not found", subDir)
	}

	if len(matches) > 1 {
		return "", stacktrace.NewError("subdir %q matches multiple paths", subDir)
	}

	return matches[0], nil
}
