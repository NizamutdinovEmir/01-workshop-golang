package storage

import (
	"fmt"
	"github.com/NizamutdinovEmir/01-workshop-golang/storage/internal/file"
	"github.com/google/uuid"
)

type Storage struct {
	Files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		Files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	newfile, err := file.NewFile(filename, blob)
	if err != nil {
		return nil, err
	}

	s.Files[newfile.ID] = newfile

	return newfile, nil
}

func (s *Storage) GetById(fileId uuid.UUID) (*file.File, error) {
	foundFile, ok := s.Files[fileId]
	if !ok {
		return nil, fmt.Errorf("file %v not found", fileId)
	}

	return foundFile, nil
}
