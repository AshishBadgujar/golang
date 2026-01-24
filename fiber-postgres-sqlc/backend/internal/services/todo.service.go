package services

import (
	"context"

	"todo-app/internal/db/sqlc"
)

type TodoService struct {
	queries *sqlc.Queries
}

func NewTodoService(q *sqlc.Queries) *TodoService {
	return &TodoService{queries: q}
}

func (s *TodoService) Create(ctx context.Context, title string) (*sqlc.Todo, error) {
	todo, err := s.queries.CreateTodo(ctx, sqlc.CreateTodoParams{
		Title:     title,
		Completed: false,
	})
	return &todo, err
}

func (s *TodoService) GetAll(ctx context.Context) ([]sqlc.Todo, error) {
	todos, err := s.queries.GetTodos(ctx)
	if todos == nil {
		return []sqlc.Todo{}, err
	}
	return todos, err
}

func (s *TodoService) GetByID(ctx context.Context, id int64) (*sqlc.Todo, error) {
	todo, err := s.queries.GetTodoByID(ctx, id)
	return &todo, err
}

func (s *TodoService) Update(
	ctx context.Context,
	id int64,
	title string,
	completed bool,
) (*sqlc.Todo, error) {

	todo, err := s.queries.UpdateTodo(ctx, sqlc.UpdateTodoParams{
		ID:        id,
		Title:     title,
		Completed: completed,
	})
	return &todo, err
}

func (s *TodoService) UpdateCompleted(ctx context.Context, id int64, completed bool) (*sqlc.Todo, error) {
	todo, err := s.queries.UpdateTodoCompleted(ctx, sqlc.UpdateTodoCompletedParams{
		ID:        id,
		Completed: completed,
	})
	return &todo, err
}

func (s *TodoService) UpdateTitle(ctx context.Context, id int64, title string) (*sqlc.Todo, error) {
	todo, err := s.queries.UpdateTodoTitle(ctx, sqlc.UpdateTodoTitleParams{
		ID:    id,
		Title: title,
	})
	return &todo, err
}

func (s *TodoService) Delete(ctx context.Context, id int64) error {
	return s.queries.DeleteTodo(ctx, id)
}
