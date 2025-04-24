package sqlite

import (
	"database/sql"

	"github.com/greninja517/student-api/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteDB struct {
	DB *sql.DB
}

func New(cfg *config.Config) (*SqliteDB, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students(
	ID INTEGER PRIMARY KEY AUTOINCREMENT,
	Name TEXT,
	Email TEXT
	)`)
	if err != nil {
		return nil, err
	}

	sqlitedb := &SqliteDB{
		DB: db,
	}

	return sqlitedb, nil
}

func (s *SqliteDB) CreateStudent(name string, email string) (int64, error) {
	// prepare the sql query statement
	stmt, err := s.DB.Prepare(`INSERT INTO students (Name, Email) VALUES (?, ?)`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	// execute the statement
	result, err := stmt.Exec(name, email)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil

}
