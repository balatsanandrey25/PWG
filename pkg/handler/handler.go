package handler

import (
	"github.com/balatsanandrey25/PWG/pkg/controllers"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	controllers *controllers.Controllers
}

func NewHandler(controllers *controllers.Controllers) *Handler {
	return &Handler{controllers: controllers}
}

func (h *Handler) InitHandler() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	{
		books := api.Group("/books")
		{
			books.GET("/", h.GetAllBooks)
			books.GET("/:id", h.GetByIDBook)
			books.POST("/", h.createBook)
			books.DELETE("/:id", h.deletBook)
			books.PATCH("/id", h.updateBook)
		}
	}

	return router
}
