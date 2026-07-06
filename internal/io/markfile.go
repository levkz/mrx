package io

import (
	"encoding/json"
	"mrx/internal/marks"
	"os"
	"path/filepath"
)

const BOOKMARKDIR = ".bookmarkz"

func getBookmarkdir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	absPath, err := filepath.Abs(filepath.Join(home, BOOKMARKDIR))
	if err != nil {
		return "", err
	}
	return absPath, err
}
func EnsureBookmarkdirExists() error {
	bmd, err := getBookmarkdir()
	if err != nil {
		return err
	}
	return os.MkdirAll(bmd, 0755)
}

func OpenMarksFile() (*os.File, error) {
	bmd, err := getBookmarkdir()
	if err != nil {
		return nil, err
	}

	return os.OpenFile(
		bmd+"/marks.json",
		os.O_CREATE|os.O_RDWR,
		0644,
	)
}

func DecodeMarkFile(f *os.File) []marks.Mark {
	var marks []marks.Mark

	dec := json.NewDecoder(f)
	dec.Decode(&marks)

	return marks
}

func EncodeMarkFile(f *os.File, marks []marks.Mark) error {
	if err := f.Truncate(0); err != nil {
		return err
	}

	if _, err := f.Seek(0, 0); err != nil {
		return err
	}
	enc := json.NewEncoder(f)
	return enc.Encode(marks)
}
