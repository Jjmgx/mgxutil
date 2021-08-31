package mgxutil

import (
	"os"
	"path"
)

func CreateFileMust(fileName string) (*os.File, error) {
	path := path.Dir(fileName)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return os.Create(fileName)
}
