package storage

import "mime/multipart"

type Storage interface {
	Store(file *multipart.FileHeader, destPath string) (string, error)
	Retrieve(filePath string) ([]byte, error)
	Delete(filePath string) error
}
