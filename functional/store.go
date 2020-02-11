package main

import (
	"io/ioutil"
	"path"
)

type Store interface {
	ReadRecipe(name string) ([]byte, error)
	WriteRecipe(name string, data []byte) error
}

type fileStore struct {
	dir string
}

func NewStore() (Store, error) {
	tmpdir, err := ioutil.TempDir("", "recipes")
	if err != nil {
		return nil, err
	}

	return &fileStore{tmpdir}, nil
}

func (f *fileStore) ReadRecipe(name string) ([]byte, error) {
	return ioutil.ReadFile(path.Join(f.dir, name))
}

func (f *fileStore) WriteRecipe(name string, data []byte) error {
	return ioutil.WriteFile(path.Join(f.dir, name), data, 0644)
}
