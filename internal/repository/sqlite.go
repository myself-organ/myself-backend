package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(dataSourceName string) (*SQLiteRepository, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}
	repo := &SQLiteRepository{db: db}
	if err := repo.Init(); err != nil {
		return nil, err
	}
	return repo, nil
}

func (r *SQLiteRepository) Init() error {
	query := `
    CREATE TABLE IF NOT EXISTS cv (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        email TEXT NOT NULL,
        phone TEXT NOT NULL,
        address TEXT NOT NULL
    );`
	_, err := r.db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}
	return nil
}

func (r *SQLiteRepository) FindByID(id int) (*CV, error) {
	query := "SELECT * FROM cv WHERE id = ?"
	row := r.db.QueryRow(query, id)
	cv := &CV{}
	if err := row.Scan(&cv.ID, &cv.Name, &cv.Email, &cv.Phone, &cv.Address); err != nil {
		return nil, fmt.Errorf("failed to find cv: %w", err)
	}
	return cv, nil
}

// ...other methods to interact with the database...
