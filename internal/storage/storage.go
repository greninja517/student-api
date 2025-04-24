package storage

import "github.com/greninja517/student-api/internal/http/types"

type Storage interface {
	CreateStudent(name string, email string) (int64, error)
	GetStudent(id int64) (types.Student, error)
}
