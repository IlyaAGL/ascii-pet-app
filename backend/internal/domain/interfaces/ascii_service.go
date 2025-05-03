package interfaces

import "ascii_store/backend/internal/domain/entities"

type AsciiService interface {
	GetAscii() (entities.Ascii, error)
	UploadAscii(entities.Ascii) error
}