package store

import "database/sql"

type PostStore struct {
	db *sql.DB
}
