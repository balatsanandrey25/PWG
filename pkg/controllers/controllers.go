package controllers

import "github.com/balatsanandrey25/PWG/pkg/models"

type BookInter interface {
	CreateBook(book *models.Book) error
	GetAll() ([]models.Book, error)
	GetByID(id int) (*models.Book, error)
	Delete(id int) error
	Update(book *models.Book) error
}
type Controllers struct {
	BookInter
}

// func NewControllers(models *models.BookInter) *Controllers {
// 	return &Controllers{mod: models}
// }

// func (c *Controllers) Create(userId int, list todo.TodoList) (int, error) {
// 	return s.repo.Create(userId, list)
// }
