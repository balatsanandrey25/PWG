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
			books.GET("/", h.controllers.GetAllBooks)
			books.GET("/:id", h.controllers.GetByIDBook)
			books.POST("/", h.controllers.CreateBook)
			books.DELETE("/:id", h.controllers.DeletBook)
			books.PATCH("/id", h.controllers.UpdateBook)
		}
	}

	return router
}
