package models

import (
	"github.com/balatsanandrey25/PWG/pkg/schema"
	"github.com/go-sqlx/sqlx"
)

type BookInter interface {
	CreateBook(book *schema.Book) error
	GetAllBooks() ([]schema.Book, error)
	GetByIDBook(id int) (*schema.Book, error)
	DeletBook(id int) error
	UpdateBook(book *schema.Book) error
}

type Models struct {
	BookInter
}

func NewModels(db *sqlx.DB) *Models {
	return &Models{BookInter: NewBookST(db)}
}
