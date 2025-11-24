package handler

import "github.com/gin-gonic/gin"

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitHandler() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	{
		movies := api.Group("/movies")
		{
			movies.GET("/", h.getAllMovies)
			movies.GET("/:id", h.getByID)
			movies.POST("/", h.createMovie)
			movies.DELETE("/:id", h.deletMovie)
			// movies.PATCH("/id", h.updateMovie)
		}
	}

	return router
}
