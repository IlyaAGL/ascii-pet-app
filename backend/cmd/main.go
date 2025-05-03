package main

import (
	"ascii_store/backend/internal/application"
	"ascii_store/backend/internal/infrastructure"
	"ascii_store/backend/internal/logger"
	"ascii_store/backend/internal/presentation/controllers"
	"database/sql"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../.env")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	db, err := sql.Open("pgx", DATABASE_URL)

	if err != nil {
		logger.Log.Info("Could not set database driver", "err", err)

		return
	}
	defer db.Close()

	logger.Log.Info("Successfully set database driver")

	err = db.Ping()
	if err != nil {
		logger.Log.Info("Could not ping the database", "err", err)

		return
	}

	logger.Log.Info("Connection was set")

	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		logger.Log.Info("Could not create migration driver")

		return
	}

	m, err := migrate.NewWithDatabaseInstance(
		os.Getenv("MIGRATIONS_PATH"),
		"postgres", driver)

	if err != nil {
		logger.Log.Info("Could not create migrate instance", "err", err)

		return
	}

	err = m.Up()

	if err != nil && err != migrate.ErrNoChange {
		logger.Log.Info("Could not apply migrations", "err", err.Error())

		return
	}

	logger.Log.Info("Migrations applied successfully!")
	logger.Log.Info("Successfully connected to db")

	repo := infrastructure.NewAsciiRepository(db)
	service := application.NewAsciiService(repo)
	handler := controllers.NewAsciiHandler(service)

	handler.StartApi()
}