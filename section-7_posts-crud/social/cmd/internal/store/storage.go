package store

import (
	"context"
	"database/sql"
	"errors"
)

var (
	ErrNotFound = errors.New("resource not found")
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetByID(ctx context.Context, id int64) (*Post, error)
	}
	Users interface {
		Create(context.Context, *User) error
	}
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		Posts: &PostsStore{db},
		Users: &UsersStore{db},
	}
}
