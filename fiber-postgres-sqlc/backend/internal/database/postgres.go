package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(
	host, port, user, password, dbname string,
) *pgxpool.Pool {

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		user, password, host, port, dbname,
	)

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}

	return pool
}
