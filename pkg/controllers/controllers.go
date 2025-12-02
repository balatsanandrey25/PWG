package controllers

import (
	"github.com/balatsanandrey25/PWG/pkg/models"
	"github.com/balatsanandrey25/PWG/pkg/schema"
)

type BookInter interface {
	CreateBook(book *schema.Book) error
	GetAllBooks() ([]schema.Book, error)
	GetByIDBook(id int) (*schema.Book, error)
	DeletBook(id int) error
	UpdateBook(book *schema.Book) error
}

type Controllers struct {
	BookInter
}

func NewControllers(models *models.Models) *Controllers {
	return &Controllers{
		BookInter: NewBook(models.BookInter),
	}
}

// func (c *Controllers) Create(userId int, list todo.TodoList) (int, error) {
// 	return s.repo.Create(userId, list)
// }
