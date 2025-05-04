package application

import (
	"ascii_store/backend/internal/domain/entities"
	"ascii_store/backend/internal/domain/interfaces"
)

type AsciiService struct {
	repo interfaces.AsciiRepository
}

func NewAsciiService(repo interfaces.AsciiRepository) *AsciiService {
	return &AsciiService{repo: repo}
}

func (s *AsciiService) GetAscii() (entities.Ascii, error) {
	return s.repo.GetAscii()
}

func (s *AsciiService) UploadAscii(ascii entities.Ascii) error {
	return s.repo.UploadAscii(ascii)
}

func (s *AsciiService) DeleteAscii() error {
	return s.repo.DeleteAscii()
}