package models

import "github.com/go-sqlx/sqlx"

type BookST struct {
	db *sqlx.DB
}

func NewBookST(db *sqlx.DB) *BookST {
	return &BookST{db: db}
}

func (m *BookST) CreateBook(book *Book) error {
	query := `
        INSERT INTO books (name, author, publication)
        VALUES ($1, $2, $3)
        RETURNING id, created_at, updated_at
    `
	return m.db.QueryRow(
		query,
		book.Name,
		book.Author,
		book.Publication,
	).Scan(&book.Id, &book.CreatedAt, &book.UpdatedAt)
}

func (m *BookST) GetByID(id int) (*Book, error) {
	var book Book
	query := `SELECT * FROM books WHERE id = $1`
	err := m.db.Get(&book, query, id)
	return &book, err
}

func (m *BookST) GetAll() ([]Book, error) {
	var books []Book
	query := `SELECT * FROM books ORDER BY id`
	err := m.db.Select(&books, query)
	return books, err
}

func (m *BookST) Update(book *Book) error {
	query := `
        UPDATE books
        SET name = $1, author = $2, publication = $3, updated_at = CURRENT_TIMESTAMP
        WHERE id = $4
    `
	_, err := m.db.Exec(
		query,
		book.Name,
		book.Author,
		book.Publication,
		book.Id,
	)
	return err
}

func (m *BookST) Delete(id int) error {
	query := `DELETE FROM books WHERE id = $1`
	_, err := m.db.Exec(query, id)
	return err
}
