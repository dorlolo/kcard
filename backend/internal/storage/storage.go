package storage

import (
	"context"
	"io"
	"os"
	"path/filepath"
)

type Store interface {
	Save(ctx context.Context, key string, r io.Reader) (string, error)
	Open(ctx context.Context, key string) (io.ReadCloser, error)
}

type LocalStore struct{ root string }

func NewLocalStore(root string) LocalStore { return LocalStore{root: root} }

func (s LocalStore) Save(ctx context.Context, key string, r io.Reader) (string, error) {
	path := filepath.Join(s.root, filepath.Clean(key))
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return "", err
	}
	file, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	_, err = io.Copy(file, r)
	return path, err
}

func (s LocalStore) Open(ctx context.Context, key string) (io.ReadCloser, error) {
	return os.Open(filepath.Join(s.root, filepath.Clean(key)))
}
