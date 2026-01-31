package storage

import "github.com/nitin0409sep/students-api/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetById(id int64) (types.Student, error)
}
