package storage

import (
	"context"

	"github.com/google/uuid"
	db "github.com/jpmoraess/gift-api/db/sqlc"
)

type FileRepository struct {
	db db.Store
}

func NewFileRepository(db db.Store) *FileRepository {
	return &FileRepository{db: db}
}

func (r *FileRepository) Save(ctx context.Context, file *File) error {
	_, err := r.db.InsertFile(ctx, db.InsertFileParams{
		ID:        file.ID,
		Name:      file.Name,
		Extension: file.Extension,
		Size:      int32(file.Size),
		Path:      file.Path,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *FileRepository) GetFile(ctx context.Context, id uuid.UUID) (file *File, err error) {
	data, err := r.db.GetFile(ctx, id)
	if err != nil {
		return
	}

	return &File{
		ID:        data.ID,
		Name:      data.Name,
		Extension: data.Extension,
		Size:      int64(data.Size),
		Path:      data.Path,
	}, nil
}

func (r *FileRepository) Delete(ctx context.Context, id uuid.UUID) (err error) {
	return r.db.DeleteFile(ctx, id)
}
