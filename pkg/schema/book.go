package schema

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
