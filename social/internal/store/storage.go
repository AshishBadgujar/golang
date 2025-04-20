package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Post interface {
		Create(context.Context) error
	}
	User interface {
		Create(context.Context) error
	}
}

func NewPostgresStorage(db *sql.DB) Storage {
	return Storage{
		Post: NewPostgresPostStore(db),
		User: NewPostgresUserStore(db),
	}
}
