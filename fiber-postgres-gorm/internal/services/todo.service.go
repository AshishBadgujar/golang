package services

import (
	"todo-app/internal/models"

	"gorm.io/gorm"
)

type TodoService struct {
	db *gorm.DB
}

func NewTodoService(db *gorm.DB) *TodoService {
	return &TodoService{db: db}
}

func (s *TodoService) Create(todo *models.Todo) error {
	return s.db.Create(todo).Error
}

func (s *TodoService) GetAll() ([]models.Todo, error) {
	var todos []models.Todo
	err := s.db.Find(&todos).Error
	return todos, err
}

func (s *TodoService) GetByID(id uint) (*models.Todo, error) {
	var todo models.Todo
	err := s.db.First(&todo, id).Error
	return &todo, err
}

func (s *TodoService) Update(todo *models.Todo) error {
	return s.db.Save(todo).Error
}

func (s *TodoService) Delete(id uint) error {
	return s.db.Delete(&models.Todo{}, id).Error
}
