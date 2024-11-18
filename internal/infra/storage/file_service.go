package storage

import (
	"context"
	"mime/multipart"

	"github.com/google/uuid"
)

type FileService struct {
	Storage        Storage
	FileRepository *FileRepository
}

func NewFileService(storage Storage, fileRepository *FileRepository) *FileService {
	return &FileService{Storage: storage, FileRepository: fileRepository}
}

func (fs *FileService) Upload(ctx context.Context, file *multipart.FileHeader, destPath string) (f *File, err error) {
	path, err := fs.Storage.Store(file, destPath)
	if err != nil {
		return
	}

	f = &File{
		ID:        uuid.New(),
		Name:      file.Filename,
		Extension: destPath,
		Size:      file.Size,
		Path:      path,
	}

	if err = fs.FileRepository.Save(ctx, f); err != nil {
		return
	}

	return
}

func (fs *FileService) Download(ctx context.Context, id uuid.UUID) ([]byte, error) {
	file, err := fs.FileRepository.GetFile(ctx, id)
	if err != nil {
		return nil, err
	}

	return fs.Storage.Retrieve(file.Path)
}

func (fs *FileService) Delete(ctx context.Context, id uuid.UUID) (err error) {
	file, err := fs.FileRepository.GetFile(ctx, id)
	if err != nil {
		return
	}

	if err = fs.Storage.Delete(file.Path); err != nil {
		return
	}

	return fs.FileRepository.Delete(ctx, id)
}
