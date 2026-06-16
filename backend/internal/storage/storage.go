// Package storage 提供文件存储接口及其本地文件系统实现。
package storage

import (
	"context"
	"io"
	"os"
	"path/filepath"
)

// Store 定义文件存储接口，支持文件的保存与打开操作。
type Store interface {
	Save(ctx context.Context, key string, r io.Reader) (string, error)
	Open(ctx context.Context, key string) (io.ReadCloser, error)
}

// LocalStore 基于本地文件系统的存储实现，将文件保存到指定根目录下。
type LocalStore struct{ root string }

// NewLocalStore 创建指定根目录的 LocalStore 实例。
func NewLocalStore(root string) LocalStore { return LocalStore{root: root} }

// Save 将读取器中的内容保存到存储中，返回文件的完整路径。
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

// Open 打开存储中指定键对应的文件，返回读取流。
func (s LocalStore) Open(ctx context.Context, key string) (io.ReadCloser, error) {
	return os.Open(filepath.Join(s.root, filepath.Clean(key)))
}
