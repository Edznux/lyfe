package storage

import (
	"context"
	"database/sql"
	"log"

	_ "embed"

	"github.com/edznux/lyfe/storage/sqliteStore"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed sqliteStore/schema.sql
var ddl string

type Store struct {
	db      *sql.DB
	Querier *sqliteStore.Queries
}

func NewStore(ctx context.Context) (*Store, error) {
	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	// create tables
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return nil, err
	}

	querier := sqliteStore.New(db)

	return &Store{
		db:      db,
		Querier: querier,
	}, nil
}
