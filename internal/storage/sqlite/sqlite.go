package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/greninja517/student-api/internal/config"
	"github.com/greninja517/student-api/internal/http/types"
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

func (s *SqliteDB) GetStudent(id int64) (types.Student, error) {
	var student types.Student

	// prepare the query statement
	stmt, err := s.DB.Prepare(`SELECT ID, Name,Email FROM students WHERE ID=?`)
	if err != nil {
		return types.Student{}, err
	}
	defer stmt.Close()

	// execute the Statement
	err = stmt.QueryRow(id).Scan(&student.ID, &student.Name, &student.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("no student found with id: %v", id)
		}
		return student, fmt.Errorf("query error for id: %v. Error: %v", id, err)
	}

	return student, nil
}
