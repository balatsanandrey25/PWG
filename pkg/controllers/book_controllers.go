package controllers

import (
	"errors"

	"github.com/balatsanandrey25/PWG/pkg/models"
	"github.com/balatsanandrey25/PWG/pkg/schema"
)

var (
	ErrBookNotFound = errors.New("Book not found")
	ErrInvalidID    = errors.New("invalid Book ID")
)

type BookServe struct {
	mod models.BookInter
}

func NewBook(mod models.BookInter) *BookServe {
	return &BookServe{mod: mod}
}

func (c *BookServe) GetAllBooks() ([]schema.Book, error) {
	return c.mod.GetAllBooks()
}
func (c *BookServe) GetByIDBook(id int) (*schema.Book, error) {
	return c.mod.GetByIDBook(id)
}

func (c *BookServe) CreateBook(book *schema.Book) error {
	return c.mod.CreateBook(book)
}

func (c *BookServe) DeletBook(id int) error {
	return c.DeletBook(id)
}

func (c *BookServe) UpdateBook(book *schema.Book) error {
	return c.mod.UpdateBook(book)
}
