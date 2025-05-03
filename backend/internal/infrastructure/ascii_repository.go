package infrastructure

import (
	"ascii_store/backend/internal/domain/entities"
	"ascii_store/backend/internal/logger"
	"database/sql"
)

type AsciiRepository struct {
	db *sql.DB
}

func NewAsciiRepository(db *sql.DB) *AsciiRepository {
	return &AsciiRepository{db: db}
}

func (r *AsciiRepository) GetAscii() (entities.Ascii, error) {
	query := "SELECT * FROM Ascii LIMIT 1"

	row := r.db.QueryRow(query)

	var a entities.Ascii
	err := row.Scan(&a.Ascii, &a.Description)

	if err != nil {
		return entities.Ascii{}, err
	}

	logger.Log.Info("Received from DB", "DESCRIPTION", a.Description)

	return a, nil
}

func (r *AsciiRepository) UploadAscii(ascii entities.Ascii) error {
	logger.Log.Info("Inserting ASCII", "ASCII", "DESCRIPTION", ascii.Description)
	query := "INSERT INTO Ascii VALUES ($1, $2)"

	return r.db.QueryRow(query, 
		ascii.Ascii, 
		ascii.Description).Err()
}