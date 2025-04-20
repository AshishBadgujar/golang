package storage

import "github.com/ashishbadgujar/golang/students-api/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentById(id int64) (types.Student, error)
	GetAllStudents() ([]types.Student, error)
	UpdateStudent(id int64, name string, email string, age int) error
	DeleteStudent(id int64) error
}
