package db

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Db *pgxpool.Pool
}

var (
	pgInstance *Postgres
	pgOnce     sync.Once
)

func New(ctx context.Context) (*Postgres, error) {
	user := os.Getenv("DB_POSTGRES_USER")
	password := os.Getenv("DB_POSTGRES_PASSWORD")
	host := os.Getenv("DB_POSTGRES_HOST")
	port := os.Getenv("DB_POSTGRES_PORT")
	database := os.Getenv("DB_POSTGRES_DATABASE")

	url := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s",
		user,
		password,
		host,
		port,
		database,
	)

	return NewForConnection(ctx, url)
}

func NewForConnection(ctx context.Context, connectionString string) (*Postgres, error) {
	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, connectionString)
		if err != nil {
			panic(fmt.Sprintf("unable to create connection pool: %s", err))
		}

		pgInstance = &Postgres{
			Db: db,
		}
	})

	return pgInstance, nil
}

func (pg *Postgres) Exec(ctx context.Context, sql string, args ...any) (err error) {
	_, err = pg.Db.Exec(ctx, sql, args...)

	return
}

func (pg *Postgres) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return pg.Db.Query(ctx, sql, args...)
}

func (pg *Postgres) QueryRow(ctx context.Context, sql string, args ...any) (pgx.Row) {
	return pg.Db.QueryRow(ctx, sql, args...)
}

func (pg *Postgres) Ping(ctx context.Context) error {
	return pg.Db.Ping(ctx)
}

func (pg *Postgres) Close() {
	pg.Db.Close()
}
