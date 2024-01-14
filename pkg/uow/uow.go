package uow

import (
	"context"
	"database/sql"
	"fmt"
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

type Uow struct {
	ctx          context.Context
	Db           *sql.DB
	Tx           *sql.Tx
	repositories map[string]RepositoryFactory
}

func NewUow(ctx context.Context, db *sql.DB) *Uow {
	return &Uow{
		ctx:          ctx,
		Db:           db,
		repositories: make(map[string]RepositoryFactory),
	}
}

func (u *Uow) Register(name string, fc RepositoryFactory) {
	u.repositories[name] = fc
}

func (u *Uow) GetRepository(ctx context.Context, name string) (interface{}, error) {
	// Implement the logic to get a repository by its name
	// For now, we'll just return the repository and nil error
	if repo, ok := u.repositories[name]; ok {
		return repo(u.Tx), nil
	}
	return nil, fmt.Errorf("repository not found")
}

func (u *Uow) Do(ctx context.Context, fn func(uow *Uow) error) error {
	// Implement the logic to execute a function with the Uow as a parameter
	// For now, we'll just call the function and return its result
	return fn(u)
}

func (u *Uow) CommitOrRollback() error {
	// Implement the logic to commit or rollback the transaction
	// For now, we'll just return a nil error
	return nil
}

func (u *Uow) Rollback() error {
	// Implement the logic to rollback the transaction
	// For now, we'll just return a nil error
	return nil
}

func (u *Uow) UnRegister(name string) {
	delete(u.repositories, name)
}
