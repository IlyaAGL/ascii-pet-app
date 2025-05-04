package interfaces

import "ascii_store/backend/internal/domain/entities"

type AsciiRepository interface {
	GetAscii() (entities.Ascii, error)
	UploadAscii(entities.Ascii) error
	DeleteAscii() error
}
