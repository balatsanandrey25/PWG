package models

import (
	"time"

	"github.com/go-sqlx/sqlx"
)

type Book struct {
	db          *sqlx.DB
	Id          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Author      string    `json:"author" db:"author"`
	Publication string    `json:"publication" db:"publication"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type BookInter interface {
	CreateBook(book *Book) error
	GetAll() ([]Book, error)
	GetByID(id int) (*Book, error)
	Delete(id int) error
	Update(book *Book) error
}

type Models struct {
	BookInter
}

func NewModels(db *sqlx.DB) *Models {
	return &Models{BookInter: NewBookST(db)}
}
