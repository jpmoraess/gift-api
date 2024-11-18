package storage

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type LocalStorage struct {
	BasePath string
}

func NewLocalStorage(basePath string) *LocalStorage {
	return &LocalStorage{BasePath: basePath}
}

func (ls *LocalStorage) Store(file *multipart.FileHeader, destPath string) (string, error) {
	fullPath := filepath.Join(ls.BasePath, destPath)
	dir := filepath.Dir(fullPath)

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", err
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer func(src multipart.File) {
		err = src.Close()
		if err != nil {

		}
	}(src)

	dst, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer func(dst *os.File) {
		err = dst.Close()
		if err != nil {

		}
	}(dst)

	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}

	return fullPath, nil
}

func (ls *LocalStorage) Retrieve(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

func (ls *LocalStorage) Delete(filePath string) error {
	return os.Remove(filePath)
}
