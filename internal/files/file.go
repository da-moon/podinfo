package files

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	primitives "github.com/da-moon/podinfo/internal/primitives"
	stacktrace "github.com/palantir/stacktrace"
)

// FileSize returns file size for the given path.
// TODO: add flock
func FileSize(path string) (int64, error) {
	var err error
	fi, err := os.Stat(path)
	if err != nil {
		err = stacktrace.Propagate(err, "could not get file size at '%s'", path)
		return -1, err
	}
	if fi.IsDir() {
		err = stacktrace.NewError("'%s' is a directory", path)
		return -1, err
	}
	return fi.Size(), nil
}

// IsTemporaryFileName returns true if fn matches temporary file name pattern
func IsTemporaryFileName(fn string) bool {
	tmpFileNameRe := regexp.MustCompile(`\.tmp\.\d+$`)
	return tmpFileNameRe.MatchString(fn)
}
func NewFile(path string) (*os.File, error) {
	if !PathExist(path) {
		// [ NOTE ] => ensuring directory exists
		err := MkdirAll(filepath.Dir(path))
		if err != nil {
			err = stacktrace.Propagate(err, "could not create parent directory [%s]", filepath.Dir(path))
			return nil, err
		}
	}
	f, err := os.OpenFile(
		path,
		os.O_CREATE|os.O_TRUNC|os.O_WRONLY,
		0600)
	if err != nil {
		if f != nil {
			f.Close()
		}
		err = stacktrace.Propagate(err, "Could not create empty file at [%s]", path)
		return nil, err
	}
	return f, nil
}

func GetFilesHash(rootDir string, files []string) map[string]string {
	result := make(map[string]string)
	for _, v := range files {
		path := primitives.PathJoin(rootDir, v)
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			log.Printf("File %s does not exist", v)
			continue
		}
		result[v], _ = GetFileHash(path)
	}
	return result
}

func GetFileHash(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		err = stacktrace.Propagate(err, "Can't open %s for reading", path)
		return "", err
	}
	defer file.Close()
	hasher := sha256.New()
	_, err = io.Copy(hasher, file)
	if err != nil {
		err = stacktrace.Propagate(err, "Can't read file %s", path)
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

// CopyFile copies a file from src to destination
func CopyFile(src, dest string) error {
	in, err := os.Open(src)
	if err != nil {
		return nil
	}
	defer in.Close()

	out, err := os.Create(dest)
	if err != nil {
		return nil
	}
	defer func() {
		e := out.Close()
		if e != nil {
			err = e
		}
	}()
	_, err = io.Copy(out, in)
	if err != nil {
		return nil
	}

	err = out.Sync()
	if err != nil {
		return nil
	}

	si, err := os.Stat(src)
	if err != nil {
		return nil
	}

	err = os.Chmod(dest, si.Mode())
	if err != nil {
		return nil
	}
	return nil
}

// IsFile returns true if the specified path is file.
func IsFile(path string) bool {
	if !PathExist(path) {
		return false
	}
	file, _ := os.Stat(path)
	return file.IsDir()
}

// ListFiles returns a list of files from the specified path.
func ListFiles(path string) ([]string, error) {
	files := []string{}
	err := filepath.Walk(path, func(dir string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && info.Name() == gitDir {
			return filepath.SkipDir
		}
		files = append(files, dir)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

// ListFilesWithSuffix returns a list of files with specific suffix.
// TODO: use ListFiles
func ListFilesWithSuffix(input string, suffix string) ([]string, error) {
	if input == "" {
		wd, err := os.Getwd()
		if err != nil {
			err = stacktrace.Propagate(err, "could not list files in '%s' with '%s' prefix", wd, suffix)
			return nil, err
		}
		input = wd
	}
	if !PathExist(input) {
		err := stacktrace.NewError("input path '%s' not found", input)
		return nil, err
	}
	if !IsDir(input) {
		if !strings.HasSuffix(input, suffix) {
			// Input file does not have the required suffix.
			return nil, nil
		}
		return []string{input}, nil
	}
	fileInfos, err := ioutil.ReadDir(input)
	if err != nil {
		err = stacktrace.Propagate(err, "could not list files in '%s' with '%s' prefix", input, suffix)
		return nil, err
	}
	result := make([]string, 0)
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue
		}
		path := primitives.PathJoin(input, fileInfo.Name())
		path = primitives.NormalizePath(path)
		if strings.HasSuffix(path, suffix) {
			result = append(result, path)
		}
	}
	return result, nil
}
