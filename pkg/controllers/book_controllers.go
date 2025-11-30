package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	ErrBookNotFound = errors.New("Book not found")
	ErrInvalidID    = errors.New("invalid Book ID")
)

// func NewBool(mod ) *AuthService {
// 	return &AuthService{repo: repo}
// }

func (с *Controllers) GetAllBooks(c *gin.Context) {

}
func (с *Controllers) GetByIDBook(c *gin.Context) {

}

func (с *Controllers) CreateBook(c *gin.Context) {
}

func (с *Controllers) DeletBook(c *gin.Context) {

}

func (с *Controllers) UpdateBook(c *gin.Context) {

}

// func newErrorResponse(c *gin.Context, statusCode int, message string) {
// 	c.AbortWithStatusJSON(statusCode, errorResponse{message})
// }
