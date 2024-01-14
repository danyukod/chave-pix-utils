package uow

import (
	"context"
	"database/sql"
)

type RepositoryFactory func(tx *sql.Tx) interface{}

type UowInterface interface {
	Register(name string, fc RepositoryFactory)
	GetRepository(ctx context.Context, name string) (interface{}, error)
	Do(ctx context.Context, fn func(uow *Uow) error) error
	CommitOrRollback() error
	Rollback() error
	UnRegister(name string)
}

func NewUow(ctx context.Context, db *sql.DB) *Uow {
	return &Uow{
		ctx:          ctx,
		Db:           db,
		repositories: make(map[string]RepositoryFactory),
	}
}

type Uow struct {
	ctx          context.Context
	Db           *sql.DB
	Tx           *sql.Tx
	repositories map[string]RepositoryFactory
}
