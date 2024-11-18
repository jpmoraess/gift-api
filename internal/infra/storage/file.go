package storage

import (
	"github.com/google/uuid"
)

type File struct {
	ID        uuid.UUID
	Name      string
	Extension string
	Size      int64
	Path      string
}
