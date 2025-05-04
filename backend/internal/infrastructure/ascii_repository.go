package infrastructure

import (
	"ascii_store/backend/internal/domain/entities"
	"ascii_store/backend/internal/logger"
	"encoding/json"
	"errors"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

const dataDir = "./data" // Место, куда будут складываться все данные

type FileAsciiRepository struct{}

func NewAsciiRepository() *FileAsciiRepository {
	if err := os.MkdirAll(dataDir, 0755); err != nil { // Убедимся, что папка существует
		logger.Log.Info("The folder ./data already exists", "ERROR", err.Error())

		return nil
	}

	return &FileAsciiRepository{}
}

func (r *FileAsciiRepository) GetAscii() (entities.Ascii, error) {
	files, err := filepath.Glob(filepath.Join(dataDir, "*.json")) // Получаем набор паттернов файлов, которые имеют расширение .json
	if err != nil {
		return entities.Ascii{}, err
	}

	if len(files) == 0 {
		return entities.Ascii{}, errors.New("no ASCII files found")
	}

	randomFile := files[rand.Intn(len(files))]

	buf, err := os.ReadFile(randomFile)
	if err != nil {
		return entities.Ascii{}, err
	}

	var a entities.Ascii
	if err := json.Unmarshal(buf, &a); err != nil {
		return entities.Ascii{}, err
	}

	logger.Log.Info("Loaded ASCII from file", "file", randomFile, "description", a.Description)
	return a, nil
}

func (r *FileAsciiRepository) UploadAscii(ascii entities.Ascii) error {
	logger.Log.Info("Inserting ASCII", "ASCII", "DESCRIPTION", ascii.Description)

	filename := uuid.New().String() + ".json" // Было решено генерировать названия файлов, присваивая им уникальный uuid (решаю проблему отправки одних и тех же данных)
	fullPath := filepath.Join(dataDir, filename)

	buf, err := json.MarshalIndent(ascii, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(fullPath, buf, 0644); err != nil {
		return err
	}

	logger.Log.Info("Saved ASCII to file", "file", fullPath, "description", ascii.Description)
	return nil
}

func (r *FileAsciiRepository) DeleteAscii() error {
	files, err := filepath.Glob(filepath.Join(dataDir, "*.json")) // Получаем набор паттернов файлов, которые имеют расширение .json
	if err != nil {
		return err
	}

	for _, f := range files {
		if err := os.Remove(f); err != nil {
			return err
		}
	}

	logger.Log.Info("Deleted all ASCII files")
	return nil
}
